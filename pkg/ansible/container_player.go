package ansible

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

type apiClient interface {
	ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)
	ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
	ContainerCreate(
		ctx context.Context,
		config *container.Config,
		hostConfig *container.HostConfig,
		networkingConfig *network.NetworkingConfig,
		containerName string,
	) (container.ContainerCreateCreatedBody, error)
	ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error
	ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
	ContainerExecAttach(
		ctx context.Context,
		execID string,
		config types.ExecStartCheck,
	) (types.HijackedResponse, error)
	ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error)
	ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
	ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error
}

// ContainerPlayer plays Ansible playbooks via docker exec.
type ContainerPlayer struct {
	reporter *report.Reporter
	log      *logrus.Entry
	client   apiClient
}

// NewContainerPlayer returns *ContainerPlayer.
func NewContainerPlayer(reporter *report.Reporter, logFilePath string) (*ContainerPlayer, error) {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &ContainerPlayer{
		reporter: reporter,
		log:      logutil.NewFileLogger("ansible-deployer-container", logFilePath),
		client:   cli,
	}, nil
}

// StartExecuteAndRemove creates a container, executes certain command and after completion it removes that container.
// The first element of cmd is command and the rest are its parameters.
// Notice that imageRepo should already be added to docker daemon's insecure registry list by contrail-command-deployer,
// so private insecure registries don't need extra care when getting pulled.
func (p *ContainerPlayer) StartExecuteAndRemove(
	ctx context.Context, cp *ContainerParameters, cmd []string,
) (err error) {
	p.log.WithFields(logrus.Fields{
		"command":    cmd,
		"parameters": cp.String(),
	}).Info("Running cmd in container")

	containerName, err := p.ensureContainerRunning(ctx, cp)
	if err != nil {
		return errors.Wrap(err, "ensure container running")
	}

	if cp.RemoveContainer {
		defer func() {
			if removeErr := p.removeContainer(ctx, containerName); removeErr != nil {
				err = errors.Wrapf(err, "could not remove container: %v", removeErr)
			}
		}()
	}

	return p.execCmd(ctx, containerName, cp, cmd)
}

// ContainerParameters holds all the necessary parameters for executing a command in container.
type ContainerParameters struct {
	// HostVolumes is the list of volumes that will be mounted from host to container.
	// If ForceContainerRecreate is disabled and Container exists it won't load current volumes!
	HostVolumes []Volume
	// Env is a list of environment variables. Single string value should be in form KEY=VALUE.
	Env []string
	// ImageRef is the name of image that will be pulled for certain action REGISTRY/IMAGE_NAME:TAG.
	ImageRef string
	// ImageRefUsername is a username for Docker registry.
	ImageRefUsername string
	// ImageRefPassword is a password for Docker registry.
	ImageRefPassword string
	// WorkingDirectory is the directory from which a command is started.
	WorkingDirectory string
	// ContainerPrefix is the prefix that will be used for container name.
	ContainerPrefix string
	// If true instead of starting base container entrypoint it will call command to keep it alive.
	OverwriteEntrypoint bool
	// ForceContainerRecreate determines whether if there was an existing container should it be used or removed
	// and created again to use current volumes.
	ForceContainerRecreate bool
	// RemoveContainer determines whether container should be removed after its work.
	RemoveContainer bool
	// HostNetwork determines whether new container should have access to host's network configuration.
	HostNetwork bool
	// Privileged gives all capabilities to the container.
	Privileged bool
}

func (c *ContainerParameters) String() string {
	out, err := yaml.Marshal(c)
	if err != nil {
		return "Unable to marshal Container Parameters"
	}
	return string(out)
}

func (p *ContainerPlayer) ensureContainerRunning(ctx context.Context, cp *ContainerParameters) (string, error) {
	// TODO (Ignacy): Move pullImage method to createRunningContainer method.
	if err := p.pullImage(ctx, cp.ImageRef, cp.ImageRefUsername, cp.ImageRefPassword); err != nil {
		return "", errors.Wrap(err, "image pulling failed")
	}
	containerName, err := p.startContainer(ctx, cp)
	if err != nil {
		return "", errors.Wrap(err, "container Initialization failed")
	}
	return containerName, nil
}

func (p *ContainerPlayer) startContainer(ctx context.Context, cp *ContainerParameters) (string, error) {
	// Because we don't use fixed name ansible deployer container, and the user might choose to not remove the
	// container after finishing ansible deployment, we need to search for existing container created from specified
	// image registry first to use it.
	containerInstance, isFound, err := p.searchContainer(ctx, cp.ImageRef)
	if err != nil {
		return "", errors.Wrap(err, "container search failed")
	}

	if isFound && cp.ForceContainerRecreate {
		if err = p.removeContainer(ctx, getContainerName(containerInstance)); err != nil {
			return "", err
		}
	} else if isFound {
		p.log.Info("Found already existing container. Volumes won't be updated")
		return p.startExistingContainer(ctx, containerInstance)
	}

	return p.createRunningContainer(ctx, cp)
}

func (p *ContainerPlayer) execCmd(
	ctx context.Context, containerName string, cp *ContainerParameters, cmd []string,
) error {
	p.log.WithFields(logrus.Fields{
		"container-name":   containerName,
		"workingDirectory": cp.WorkingDirectory,
		"command":          cmd,
	}).Debug("Executing command")

	createResp, err := p.client.ContainerExecCreate(
		ctx,
		containerName,
		types.ExecConfig{
			AttachStdout: true,
			AttachStderr: true,
			WorkingDir:   cp.WorkingDirectory,
			Cmd:          cmd,
			Env:          cp.Env,
			Privileged:   cp.Privileged,
		},
	)
	if err != nil {
		return errors.Wrap(err, "creating container exec failed")
	}

	hijack, err := p.client.ContainerExecAttach(ctx, createResp.ID, types.ExecStartCheck{})
	if err != nil {
		return errors.Wrap(err, "attaching to container exec failed")
	}
	defer hijack.Close()

	// TODO(dji): use a better method to deal with scanner not supporting non-UTF8 control characters
	// For some reason, bufio.Scanner doesn't deal with control character in ansible output
	// Used docker package's stdcopy.Stdcopy which deals with control characters
	buffer := new(bytes.Buffer)
	_, err = stdcopy.StdCopy(buffer, buffer, hijack.Reader)
	if err != nil {
		return errors.Wrap(err, "copy ansible ouput from container to local buffer failed")
	}

	scanner := bufio.NewScanner(strings.NewReader(buffer.String()))
	for scanner.Scan() {
		p.log.Debug(scanner.Text())
	}

	execInspectResp, err := p.client.ContainerExecInspect(ctx, createResp.ID)
	if err != nil {
		return errors.Wrap(err, "inspecting container exec failed")
	}
	if execInspectResp.ExitCode != 0 {
		return errors.New("deployment in container failed, status code: " + strconv.Itoa(execInspectResp.ExitCode))
	}
	return nil
}

func (p *ContainerPlayer) pullImage(
	ctx context.Context,
	imageRef, username, password string,
) (err error) {
	p.log.WithFields(logrus.Fields{
		"image": imageRef,
	}).Debug("Pulling image")

	authConfig := types.AuthConfig{
		Username: username,
		Password: password,
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return errors.Wrap(err, "authConfig JSON marshaling failed")
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	reader, err := p.client.ImagePull(ctx, imageRef, types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		return errors.Wrap(err, "image pulling failed")
	}

	defer func() {
		cerr := reader.Close()
		if err == nil {
			err = errors.Wrap(cerr, "closing reader from image pull response failed")
		}
	}()

	_, err = ioutil.ReadAll(reader)
	return errors.Wrap(err, "reading from image pull response failed")
}

func (p *ContainerPlayer) searchContainer(ctx context.Context, imageRef string) (*types.Container, bool, error) {
	p.log.WithFields(logrus.Fields{
		"image": imageRef,
	}).Debug("Searching for existing container created from specified image")

	containerInstances, err := p.client.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, false, errors.Wrap(err, "listing all containers failed")
	}

	for _, containerInstance := range containerInstances {
		if containerInstance.Image == imageRef {
			return &containerInstance, true, nil
		}
	}

	return nil, false, nil
}

func getContainerName(containerInstance *types.Container) string {
	if containerInstance == nil {
		return ""
	}
	return strings.Join(containerInstance.Names, "")
}

func (p *ContainerPlayer) startExistingContainer(
	ctx context.Context,
	containerInstance *types.Container,
) (string, error) {
	p.log.WithFields(logrus.Fields{
		"container-name": getContainerName(containerInstance),
	}).Debug("Starting existing containery")

	if containerInstance.State != "running" {
		err := p.client.ContainerStart(ctx, containerInstance.ID, types.ContainerStartOptions{})
		if err != nil {
			return "", errors.Wrap(err, "starting container failed")
		}
	}

	return getContainerName(containerInstance), nil
}

func (p *ContainerPlayer) createRunningContainer(ctx context.Context, cp *ContainerParameters) (string, error) {
	containerName := strings.Join([]string{cp.ContainerPrefix, time.Now().Format("20060102150405")}, "_")

	p.log.WithFields(logrus.Fields{
		"image":          cp.ImageRef,
		"container-name": containerName,
	}).Debug("Creating container from specified image")

	mounts := toMounts(cp.HostVolumes)

	if err := validateVolumesSourcePaths(mounts); err != nil {
		return "", errors.Wrap(err, "container volumes are invalid")
	}

	resp, err := p.client.ContainerCreate(
		ctx,
		&container.Config{
			Tty:   true,
			Image: cp.ImageRef,
			Cmd:   getContainerCmd(cp.OverwriteEntrypoint),
		},
		&container.HostConfig{
			Mounts:      mounts,
			NetworkMode: getNetworkMode(cp.HostNetwork),
		},
		nil,
		containerName,
	)
	if err != nil {
		return "", errors.Wrap(err, "creating container failed")
	}

	err = p.client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", errors.Wrap(err, "starting container failed")
	}

	return containerName, nil
}

// Volume represents a configuration for a Volume that will be added to Container.
type Volume struct {
	Source string
	Target string
}

func toMounts(volumes []Volume) []mount.Mount {
	mounts := make([]mount.Mount, len(volumes))
	for i, v := range volumes {
		mounts[i] = mount.Mount{
			Type:   mount.TypeBind,
			Source: v.Source,
			Target: v.Target,
		}
	}
	return mounts
}

func validateVolumesSourcePaths(volumes []mount.Mount) error {
	fileErrors := []string{}
	for _, v := range volumes {
		if _, err := os.Stat(v.Source); err != nil {
			fileErrors = append(fileErrors, fmt.Sprintf("source path %s is invalid: %s", v.Source, err.Error()))
		}
	}
	if len(fileErrors) != 0 {
		return errors.New(strings.Join(fileErrors, ", "))
	}
	return nil
}

func getContainerCmd(overwriteEntrypoint bool) []string {
	if overwriteEntrypoint {
		return []string{"tail", "-f", "/dev/null"}
	}
	return nil
}

func getNetworkMode(hostNetwork bool) container.NetworkMode {
	if hostNetwork {
		return "host"
	}
	return ""
}

// TODO(dji): Notice that the stopping and starting of a container takes much longer time than create and remove it
// In order to improve efficiency in the future, the container should not be removed between ansible calls,
// but it needs to be removed when all ansible calls finish. This is left as future improvement.
func (p *ContainerPlayer) removeContainer(ctx context.Context, containerName string) error {
	p.log.WithFields(logrus.Fields{
		"container-name": containerName,
	}).Debug("Cleaning up container")

	err := p.client.ContainerStop(ctx, containerName, new(time.Duration))
	if err != nil {
		return errors.Wrap(err, "stopping container failed")
	}

	err = p.client.ContainerRemove(ctx, containerName, types.ContainerRemoveOptions{Force: true})
	return errors.Wrap(err, "force removing container failed")
}

// GetContrailVersion returns Contrail version specified by user.
func GetContrailVersion(cluster *models.ContrailCluster, logger *logrus.Entry) string {
	if version := cluster.ContrailConfiguration.GetValue("CONTRAIL_CONTAINER_TAG"); version != "" {
		if cluster.ContrailVersion != "" && cluster.ContrailVersion != version {
			logger.Warnf(
				"Contrail Version is different for CONTRAIL_CONTAINER_TAG: %s, and CONTRAIL_VERSION: %s. Using: %s",
				version, cluster.ContrailVersion, version)
		}
		return version
	}
	return cluster.ContrailVersion
}

// ImageReference returns image reference.
func ImageReference(registry, imageName, imageTag string) (string, error) {
	missingFields := []string{}
	if registry == "" {
		missingFields = append(missingFields, "Registry")
	}
	if imageName == "" {
		missingFields = append(missingFields, "Image Name")
	}
	if imageTag == "" {
		missingFields = append(missingFields, "Image Tag")
	}
	if len(missingFields) != 0 {
		return "", errors.Errorf("cannot resolve image reference. Missing: %s", strings.Join(missingFields, ", "))
	}
	return registry + "/" + imageName + ":" + imageTag, nil
}
