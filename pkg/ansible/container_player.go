package ansible

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
	ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error)
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

// Play plays Ansible playbook via ansible-playbook CLI.
// It runs playbook from in docker container.
// Notice that imageRepo should already be added to docker daemon's insecure registry list by contrail-command-deployer,
// so private insecure registries don't need extra care when getting pulled.
func (p *ContainerPlayer) Play(
	ctx context.Context,
	imageRef string,
	imageRefUsername string,
	imageRefPassword string,
	workRoot string,
	workingDirectory string,
	ansibleArgs []string,
	removeContainer bool,
) error {
	p.log.WithFields(logrus.Fields{
		"ansible-args": ansibleArgs,
	}).Info("Running playbook in container")

	err := p.pullImage(ctx, imageRef, imageRefUsername, imageRefPassword)
	if err != nil {
		return errors.Wrap(err, "image pulling failed")
	}

	// Because we don't use fixed name ansible deployer container, and the user might choose to not remove the
	// container after finishing ansible deployment, we need to search for existing container created from specified
	// image registry first to use it.
	containerInstance, isFound, err := p.searchContainer(ctx, imageRef)
	if err != nil {
		return errors.Wrap(err, "container search failed")
	}

	var containerName string
	if isFound {
		containerName, err = p.startExistingContainer(ctx, containerInstance)
	} else {
		containerName, err = p.createRunningContainer(ctx, imageRef, workRoot)
	}
	if err != nil {
		return errors.Wrap(err, "container Initialization failed")
	}

	err = p.execCmd(ctx, containerName, workingDirectory, ansibleArgs)
	if err != nil {
		return errors.Wrap(err, "container command execution failed")
	}

	if removeContainer {
		err = p.removeContainer(ctx, containerName)
		if err != nil {
			return errors.Wrap(err, "container cleaning failed")
		}
	}

	p.log.WithFields(logrus.Fields{
		"ansible-args": ansibleArgs,
	}).Info("Finished running playbook in container")

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
		return err
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	reader, err := p.client.ImagePull(ctx, imageRef, types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		return err
	}

	defer func() {
		cerr := reader.Close()
		if err == nil {
			err = cerr
		}
	}()

	_, err = ioutil.ReadAll(reader)

	return err
}

func (p *ContainerPlayer) searchContainer(ctx context.Context, imageRef string) (*types.Container, bool, error) {
	p.log.WithFields(logrus.Fields{
		"image": imageRef,
	}).Debug("Searching for existing container created from specified image")

	containerInstances, err := p.client.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, false, err
	}

	for _, containerInstance := range containerInstances {
		if containerInstance.Image == imageRef {
			return &containerInstance, true, nil
		}
	}

	return nil, false, nil
}

func (p *ContainerPlayer) startExistingContainer(
	ctx context.Context,
	containerInstance *types.Container,
) (string, error) {
	p.log.WithFields(logrus.Fields{
		"container-name": strings.Join(containerInstance.Names, ""),
	}).Debug("Starting existing containery")

	if containerInstance.State != "running" {
		err := p.client.ContainerStart(ctx, containerInstance.ID, types.ContainerStartOptions{})
		if err != nil {
			return "", err
		}
	}

	return strings.Join(containerInstance.Names, ""), nil
}

func (p *ContainerPlayer) createRunningContainer(ctx context.Context, imageRef, workRoot string) (string, error) {
	containerName := "ansible-deployer_" + time.Now().Format("20060102150405")

	p.log.WithFields(logrus.Fields{
		"image":          imageRef,
		"container-name": containerName,
	}).Debug("Creating container from specified image")

	resp, err := p.client.ContainerCreate(
		ctx,
		&container.Config{Tty: true, Image: imageRef},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: workRoot,
					Target: workRoot,
				},
			},
		},
		nil,
		containerName,
	)
	if err != nil {
		return "", err
	}

	err = p.client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	return containerName, nil
}

// playbooks: in /root/contrail-ansible-deployer
// instances: /var/tmp/contrail-cluster, but you wanna mount /var/tmp
func (p *ContainerPlayer) execCmd(
	ctx context.Context, containerName, workingDirectory string,
	ansibleArgs []string,
) error {
	p.log.WithFields(logrus.Fields{
		"container-name":    containerName,
		"workingDirectory":  workingDirectory,
		"ansible-arguments": ansibleArgs,
	}).Debug("Executing command")

	createResp, err := p.client.ContainerExecCreate(
		ctx,
		containerName,
		types.ExecConfig{
			AttachStdout: true,
			AttachStderr: true,
			WorkingDir:   workingDirectory,
			Cmd:          append([]string{playbookCmd}, ansibleArgs...),
		},
	)
	if err != nil {
		return err
	}

	hijack, err := p.client.ContainerExecAttach(ctx, createResp.ID, types.ExecStartCheck{})
	if err != nil {
		return err
	}
	defer hijack.Close()
	scanner := bufio.NewScanner(hijack.Reader)
	for scanner.Scan() {
		p.log.Debug(scanner.Text())
	}

	execInspectResp, err := p.client.ContainerExecInspect(ctx, createResp.ID)
	if err != nil {
		return err
	}
	if execInspectResp.ExitCode != 0 {
		return errors.New("deployment in container failed, status code: " + strconv.Itoa(execInspectResp.ExitCode))
	}

	return nil
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
		return err
	}

	return p.client.ContainerRemove(ctx, containerName, types.ContainerRemoveOptions{Force: true})
}
