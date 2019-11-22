package ansible

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
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
	ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
	ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error
}

// DockerClient allows to play Ansible playbooks via docker exec.
type DockerClient struct {
	reporter *report.Reporter
	log      *logrus.Entry
	client   apiClient
}

// NewDockerClient returns DockerClient.
func NewDockerClient(reporter *report.Reporter, logFilePath string) (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &DockerClient{
		reporter: reporter,
		log:      logutil.NewFileLogger("ansible-deployer-container", logFilePath),
		client:   cli,
	}, nil
}

// Play plays Ansible playbook via ansible-playbook CLI.
// It runs playbook from in docker container.
// Notice that imageRepo should already be added to docker daemon's insecure registry list by contrail-command-deployer,
// so private insecure registries don't need extra care when getting pulled.
func (c *DockerClient) Play(
	ctx context.Context,
	imageRef string,
	imageRefUsername string,
	imageRefPassword string,
	repositoryPath string,
	ansibleArgs []string,
	keepContainerAlive bool,
) error {
	c.log.WithFields(logrus.Fields{
		"ansible-args": ansibleArgs,
	}).Info("Running playbook in container")

	c.log.WithFields(logrus.Fields{
		"image-registry": imageRef,
	}).Debug("Pulling images")

	err := c.pullImage(ctx, imageRef, imageRefUsername, imageRefPassword)
	if err != nil {
		return err
	}

	// Because we don't use fixed name ansible deployer container, and the user might choose to not remove the
	// container after finishing ansible deployment, we need to search for existing container created from specified
	// image registry first to use it.
	c.log.WithFields(logrus.Fields{
		"image-registry": imageRef,
	}).Debug("Searching for existing container created from specified image registry")

	containerInstance, isFound, err := c.searchContainer(ctx, imageRef)
	if err != nil {
		return err
	}

	c.log.WithFields(logrus.Fields{
		"image-registry": imageRef,
	}).Debug("Starting container created from specified image registry")

	var containerName string
	if isFound {
		containerName, err = c.startExistingContainer(ctx, containerInstance)
	} else {
		containerName, err = c.createRunningContainer(ctx, imageRef)
	}
	if err != nil {
		return err
	}

	cmd := c.prepareDockerCmd(repositoryPath, ansibleArgs)

	c.log.WithFields(logrus.Fields{
		"container-name": containerName,
		"command":        cmd,
	}).Debug("Executing command")

	err = c.execCmd(ctx, containerName, cmd)
	if err != nil {
		return err
	}

	c.log.WithFields(logrus.Fields{
		"container-name": containerName,
		"remove":         keepContainerAlive,
	}).Debug("Cleaning up container")

	err = c.cleanContainer(ctx, containerName, keepContainerAlive)
	if err != nil {
		return err
	}

	c.log.WithFields(logrus.Fields{
		"ansible-args": ansibleArgs,
	}).Info("Finished running playbook in container")

	return nil
}

func (c *DockerClient) pullImage(ctx context.Context, imageRef string, username string, password string) error {
	authConfig := types.AuthConfig{
		Username: username,
		Password: password,
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	reader, err := c.client.ImagePull(ctx, imageRef, types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		return err
	}

	defer func() {
		tempErr := reader.Close()
		if err == nil {
			err = tempErr
		}
	}()

	_, err = ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	return err
}

func (c *DockerClient) searchContainer(ctx context.Context, imageRef string) (*types.Container, bool, error) {
	containerInstances, err := c.client.ContainerList(ctx, types.ContainerListOptions{All: true})
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

func (c *DockerClient) startExistingContainer(ctx context.Context, containerInstance *types.Container) (string, error) {
	if containerInstance.State != "running" {
		err := c.client.ContainerStart(ctx, containerInstance.ID, types.ContainerStartOptions{})
		if err != nil {
			return "", err
		}
	}

	return strings.Join(containerInstance.Names, ""), nil
}

func (c *DockerClient) createRunningContainer(ctx context.Context, imageRef string) (string, error) {
	containerName := "ansible-deployer_" + time.Now().Format("20060102150405")

	resp, err := c.client.ContainerCreate(
		ctx,
		&container.Config{Tty: true, Image: imageRef},
		nil,
		nil,
		containerName,
	)
	if err != nil {
		return "", err
	}

	err = c.client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	return containerName, nil
}

func (c *DockerClient) prepareDockerCmd(repositoryPath string, ansibleArgs []string) *exec.Cmd {
	cmd := exec.Command(playbookCmd, ansibleArgs...)
	cmd.Dir = repositoryPath

	return cmd
}

func (c *DockerClient) execCmd(ctx context.Context, containerName string, cmd *exec.Cmd) error {
	resp, err := c.client.ContainerExecCreate(
		ctx,
		containerName,
		types.ExecConfig{
			AttachStdout: true,
			AttachStderr: true,
			WorkingDir:   cmd.Dir,
			Cmd:          cmd.Args,
		},
	)
	if err != nil {
		return err
	}

	hijack, err := c.client.ContainerExecAttach(ctx, resp.ID, types.ExecStartCheck{})
	if err != nil {
		return err
	}
	defer hijack.Close()
	// TODO: use logrus formatter for ansible deployment log
	bytes, err := ioutil.ReadAll(hijack.Reader)
	if err != nil {
		return err
	}
	c.log.WithFields(logrus.Fields{
		"content": bytes,
	}).Debug("ansible deployment log")

	return nil
}

func (c *DockerClient) cleanContainer(ctx context.Context, containerName string, keepContainerAlive bool) error {
	if !keepContainerAlive {
		err := c.client.ContainerStop(ctx, containerName, new(time.Duration))
		if err != nil {
			return err
		}

		err = c.client.ContainerRemove(ctx, containerName, types.ContainerRemoveOptions{})

		return err
	}

	return nil
}
