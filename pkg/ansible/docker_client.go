package ansible

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/fileutil/template"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/logutil/report"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"
)

// DockerClient allows to play Ansible playbooks via docker exec.
type DockerClient struct {
	reporter *report.Reporter
	log      *logrus.Entry

	imageRegistry           string
	registryPrivateInsecure bool
	imageRegistryUsername   string
	imageRegistryPassword   string

	keepContainerAlive bool

	// Test setting
	workingDirectory string
	isTestInstance   bool
}

// NewDockerClient returns DockerClient.
func NewDockerClient(
	reporter *report.Reporter,
	logFilePath string,
	imageRegistry string,
	registryPrivateInsecure bool,
	imageRegistryUsername string,
	imageRegistryPassword string,
	keepContainerAlive bool,
	workingDirectory string,
	isTestInstance bool,
) *DockerClient {
	return &DockerClient{
		reporter:                reporter,
		log:                     logutil.NewFileLogger("ansible-client-container", logFilePath),
		imageRegistry:           imageRegistry,
		registryPrivateInsecure: registryPrivateInsecure,
		imageRegistryUsername:   imageRegistryUsername,
		imageRegistryPassword:   imageRegistryPassword,
		keepContainerAlive:      keepContainerAlive,
		workingDirectory:        workingDirectory,
		isTestInstance:          isTestInstance,
	}
}

// Play plays Ansible playbook via ansible-playbook CLI.
// It runs playbook from in docker container.
func (c *DockerClient) Play(repositoryPath string, ansibleArgs []string) error {
	c.log.WithFields(logrus.Fields{
		"ansible-args": ansibleArgs,
	}).Info("Running playbook in container")

	if c.isTestInstance {
		return c.mockPlay(ansibleArgs)
	}

	// Prepare command as exec.Cmd type
	cmd := prepareDockerCmd(repositoryPath, ansibleArgs)

	// Prepare context and docker client
	ctx := context.Background()

	c.log.Info("Creating http client fot accessing docker API")
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// Prepare container for command execution
	// Notice that imageRepo should already be added to docker daemon's insecure registry list by contrail-command-deployer,
	// so private insecure registries don't need extra care when getting pulled.
	c.log.WithFields(logrus.Fields{
		"image-registry": c.imageRegistry,
		"username":       c.imageRegistryUsername,
		"password":       c.imageRegistryPassword,
	}).Info("Pulling images")
	err = pullImage(ctx, cli, c.imageRegistry, c.imageRegistryUsername, c.imageRegistryPassword)
	if err != nil {
		return err
	}

	// Because we don't use fixed name ansible deployer container, and the user might choose to not remove the
	// container after finishing ansible deployment, we need to search for existing container created from specified
	// image registry first to use it.
	c.log.WithFields(logrus.Fields{
		"image-registry": c.imageRegistry,
	}).Info("Searching for exsiting container created from specified image registry")
	containerInstance, err := searchContainer(ctx, cli, c.imageRegistry)
	if err != nil {
		return err
	}

	// If we find a container in the previous step, we start it, if not, we create a new one.
	c.log.WithFields(logrus.Fields{
		"image-registry": c.imageRegistry,
	}).Info("Starting container created from specified image registry")
	containerID, containerName, err := startContainer(ctx, cli, containerInstance, c.imageRegistry)
	if err != nil {
		panic(err)
	}

	c.log.WithFields(logrus.Fields{
		"container-name": containerName,
		"container-ID":   containerID,
		"command":        cmd,
	}).Info("Executing command")
	err = execCmd(ctx, cli, containerID, cmd, true)
	if err != nil {
		return err
	}

	c.log.WithFields(logrus.Fields{
		"container-name": containerName,
		"container-ID":   containerID,
		"remove":         c.keepContainerAlive,
	}).Info("Cleaning up container")
	err = cleanContainer(ctx, cli, containerID, c.keepContainerAlive)
	if err != nil {
		return err
	}

	c.log.WithFields(logrus.Fields{
		"ansible-args": ansibleArgs,
	}).Info("Finished running playbook in container")

	return nil
}

func prepareDockerCmd(repositoryPath string, ansibleArgs []string) *exec.Cmd {
	var cmd *exec.Cmd
	cmd = exec.Command(playbookCmd, ansibleArgs...)
	cmd.Dir = repositoryPath
	return cmd
}

func pullImage(ctx context.Context, cli apiClient, imageRegistry string, username string, password string) error {
	authConfig := types.AuthConfig{
		Username: username,
		Password: password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	reader, err := cli.ImagePull(ctx, imageRegistry, types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		return err
	}
	defer reader.Close()
	_, err = ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return nil
}

func searchContainer(ctx context.Context, cli apiClient, imageRegistry string) (types.Container, error) {
	containerListOpts := types.ContainerListOptions{All: true}
	containerInstances, err := cli.ContainerList(ctx, containerListOpts)
	if err != nil {
		return types.Container{}, err
	}
	for _, containerInstance := range containerInstances {
		if containerInstance.Image == imageRegistry {
			return containerInstance, nil
		}
	}
	return types.Container{}, nil
}

func startContainer(ctx context.Context, cli apiClient, containerInstance types.Container, imageRegistry string) (string, string, error) {
	var containerID string
	var containerName string
	if containerInstance.ID == "" {
		containerName = "contrail-ansible-deployer_" + time.Now().Format("20060102150405")
		resp, err := cli.ContainerCreate(ctx, &container.Config{Tty: true, Image: imageRegistry}, nil, nil, containerName)
		if err != nil {
			return "", "", err
		}
		err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
		if err != nil {
			return "", "", err
		}
		containerID = resp.ID
	} else {
		if containerInstance.State != "running" {
			err := cli.ContainerStart(ctx, containerInstance.ID, types.ContainerStartOptions{})
			if err != nil {
				return "", "", err
			}
		}
		containerName = strings.Join(containerInstance.Names, "")
		containerID = containerInstance.ID
	}

	return containerID, containerName, nil
}

func execCmd(ctx context.Context, cli apiClient, containerID string, cmd *exec.Cmd, writeLog bool) error {
	resp, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{WorkingDir: cmd.Dir, Cmd: cmd.Args})
	if err != nil {
		return err
	}

	hijack, err := cli.ContainerExecAttach(ctx, resp.ID, types.ExecStartCheck{})
	if err != nil {
		return err
	}
	defer hijack.Close()
	if writeLog {
		file, err := os.OpenFile("/var/log/contrail/deploy.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
		if err != nil {
			return err
		}
		defer file.Close()
		io.Copy(file, hijack.Reader)
	} else {
		_, err = ioutil.ReadAll(hijack.Reader)
		if err != nil {
			return err
		}
	}

	return nil
}

func cleanContainer(ctx context.Context, cli apiClient, containerID string, keepContainerAlive bool) error {
	if keepContainerAlive {
		err := cli.ContainerStop(ctx, containerID, new(time.Duration))
		if err != nil {
			return err
		}
		err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

// WorkingDirectory returns working directory.
func (c *DockerClient) WorkingDirectory() string {
	return c.workingDirectory
}

// IsTest returns true for testing instance.
func (c *DockerClient) IsTest() bool {
	return c.isTestInstance
}

// TODO: mockPlay docker client - expected yml files needed
func (c *DockerClient) mockPlay(ansibleArgs []string) error {
	playBookIndex := len(ansibleArgs) - 1
	content, err := template.Apply("./test_data/test_ansible_playbook_container.tmpl", pongo2.Context{
		"containerName": "ansible-deployer-container",
		"playBook":      ansibleArgs[playBookIndex],
		"ansibleArgs":   strings.Join(ansibleArgs[:playBookIndex], " "),
	})
	if err != nil {
		return err
	}

	return fileutil.AppendToFile(
		filepath.Join(c.workingDirectory, "executed_ansible_playbook.yml"),
		content,
		filePermRWOnly,
	)
}
