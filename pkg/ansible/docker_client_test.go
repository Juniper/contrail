package ansible

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"os/exec"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

type apiClient interface {
	ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)
	ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
	ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)
	ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error
	ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
	ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error)
	ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
	ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error
}

type mockConnection struct {
	net.Conn
}

type mockAPIClient struct {
	expectedUsername      string
	expectedPassword      string
	containerInstances    []types.Container
	containerInstance     types.Container
	expectedContainerID   string
	expectedContainerName string
}

func (mockConn *mockConnection) Close() error {
	return nil
}

func (mockClient *mockAPIClient) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {
	authConfig := types.AuthConfig{
		Username: mockClient.expectedUsername,
		Password: mockClient.expectedPassword,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return nil, err
	}
	expectedAuth := base64.URLEncoding.EncodeToString(encodedJSON)
	if options.RegistryAuth != expectedAuth {
		return nil, errors.New("Authentication failed")
	}
	return ioutil.NopCloser(strings.NewReader("Success")), nil
}

func (mockClient *mockAPIClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return mockClient.containerInstances, nil
}

func (mockClient *mockAPIClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	if mockClient.containerInstance.ID != "" {
		return container.ContainerCreateCreatedBody{
			ID: mockClient.containerInstance.ID,
		}, nil
	}
	return container.ContainerCreateCreatedBody{
		ID: mockClient.expectedContainerID,
	}, nil
}

func (mockClient *mockAPIClient) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	return nil
}

func (mockClient *mockAPIClient) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error) {
	return types.IDResponse{
		ID: "1",
	}, nil
}

func (mockClient *mockAPIClient) ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error) {
	return types.HijackedResponse{
		Conn:   &mockConnection{},
		Reader: bufio.NewReader(strings.NewReader("Output")),
	}, nil
}

func (mockClient *mockAPIClient) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error {
	return nil
}

func (mockClient *mockAPIClient) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	return nil
}

func TestPullImageAuthentication(t *testing.T) {
	pullCases := map[string]struct {
		imageRegistry    string
		username         string
		password         string
		expectedUsername string
		expectedPassword string
		isSuccessful     bool
	}{
		"Correct credentials": {
			imageRegistry:    "busybox",
			username:         "admin",
			password:         "123456",
			expectedUsername: "admin",
			expectedPassword: "123456",
			isSuccessful:     true,
		},
		"Wrong password": {
			imageRegistry:    "busybox",
			username:         "admin",
			password:         "654321",
			expectedUsername: "admin",
			expectedPassword: "123456",
			isSuccessful:     false,
		},
		"Wrong username": {
			imageRegistry:    "busybox",
			username:         "nimda",
			password:         "123456",
			expectedUsername: "admin",
			expectedPassword: "123456",
			isSuccessful:     false,
		},
		"Wrong usernam and password": {
			imageRegistry:    "busybox",
			username:         "nimda",
			password:         "654321",
			expectedUsername: "admin",
			expectedPassword: "123456",
			isSuccessful:     false,
		},
	}
	for _, pullCase := range pullCases {
		ctx := context.Background()
		mockClient := &mockAPIClient{
			expectedUsername: pullCase.expectedUsername,
			expectedPassword: pullCase.expectedPassword,
		}
		err := pullImage(ctx, mockClient, pullCase.imageRegistry, pullCase.username, pullCase.password)
		if pullCase.isSuccessful {
			if err != nil {
				t.Errorf("Authentication failed with username %s and password %s, expected username %s and password %s", pullCase.username, pullCase.password, pullCase.expectedUsername, pullCase.expectedPassword)
			}
		} else {
			if err == nil {
				t.Errorf("Authentication succeeded with username %s and password %s, expected username %s and password %s", pullCase.username, pullCase.password, pullCase.expectedUsername, pullCase.expectedPassword)
			}
		}
	}
}

func TestSearchContainer(t *testing.T) {
	searchCases := map[string]struct {
		imageName                 string
		containerInstances        []types.Container
		expectedContainerInstance types.Container
		isFound                   bool
	}{
		"Container exists": {
			imageName: "busybox",
			containerInstances: []types.Container{
				{
					ID:    "1",
					Names: []string{"1"},
					Image: "hello-world",
				},
				{
					ID:    "2",
					Names: []string{"2"},
					Image: "ubuntu",
				},
				{
					ID:    "3",
					Names: []string{"3"},
					Image: "busybox",
				},
			},
			expectedContainerInstance: types.Container{
				ID:    "3",
				Names: []string{"3"},
				Image: "busybox",
			},
			isFound: true,
		},
		"Container doesn't exist": {
			imageName: "busybox",
			containerInstances: []types.Container{
				{
					ID:    "1",
					Names: []string{"1"},
					Image: "hello-world",
				},
				{
					ID:    "2",
					Names: []string{"2"},
					Image: "ubuntu",
				},
			},
			expectedContainerInstance: types.Container{},
			isFound:                   false,
		},
	}
	for _, searchCase := range searchCases {
		ctx := context.Background()
		mock := &mockAPIClient{
			containerInstances: searchCase.containerInstances,
		}
		containerInstance, err := searchContainer(ctx, mock, searchCase.imageName)
		if err != nil {
			t.Errorf("%v when searching for container", err)
		}
		if !reflect.DeepEqual(containerInstance, searchCase.expectedContainerInstance) {
			if searchCase.isFound {
				t.Errorf("Expected %#+v, got %#+v", searchCase.expectedContainerInstance, containerInstance)
			} else {
				t.Errorf("Expected empty container instance, got %#+v", containerInstance)
			}
		}
	}
}

func TestStartContainer(t *testing.T) {
	startCases := map[string]struct {
		containerInstance     types.Container
		imageName             string
		expectedContainerID   string
		expectedContainerName string
	}{
		"Existing container": {
			containerInstance: types.Container{
				ID:    "1",
				Names: []string{"1"},
			},
			imageName:             "busybox",
			expectedContainerID:   "1",
			expectedContainerName: "1",
		},
		"No container": {
			containerInstance:     types.Container{},
			imageName:             "busybox",
			expectedContainerID:   "1",
			expectedContainerName: "contrail-ansible-deployer",
		},
	}
	for _, startCase := range startCases {
		ctx := context.Background()
		mockClient := &mockAPIClient{
			containerInstance:     startCase.containerInstance,
			expectedContainerID:   startCase.expectedContainerID,
			expectedContainerName: startCase.expectedContainerName,
		}
		containerID, containerName, err := startContainer(ctx, mockClient, startCase.containerInstance, startCase.imageName)
		if err != nil {
			t.Errorf("%v when starting container", err)
		}
		if !((containerID == startCase.expectedContainerID) && (containerName == startCase.expectedContainerName) || (strings.Contains(containerName, startCase.expectedContainerName))) {
			t.Errorf("Expected container ID %s and container name or prefix %s, got container ID %s and container name %s", startCase.expectedContainerID, startCase.expectedContainerName, containerID, containerName)
		}
	}
}

func TestExecCmd(t *testing.T) {
	execCases := map[string]struct {
		containerID string
		command     *exec.Cmd
	}{
		"Ansible-playbook command": {
			containerID: "1",
			command: &exec.Cmd{
				Path: "ansible-playbook",
				Args: []string{"ansible-playbook", "provision_instances.yml"},
				Dir:  "/usr/share/contrail/contrail-ansible-deployer",
			},
		},
	}
	for _, execCase := range execCases {
		ctx := context.Background()
		mock := &mockAPIClient{}
		err := execCmd(ctx, mock, execCase.containerID, execCase.command, false)
		if err != nil {
			t.Errorf("%v when executing command in container", err)
		}
	}
}

func TestCleanContainer(t *testing.T) {
	cleanCases := map[string]struct {
		containerID     string
		removeContainer bool
	}{
		"Stop and keep container": {
			containerID:     "1",
			removeContainer: false,
		},
		"Stop and destroy container": {
			containerID:     "2",
			removeContainer: true,
		},
	}
	for _, cleanCase := range cleanCases {
		ctx := context.Background()
		mock := &mockAPIClient{}
		err := cleanContainer(ctx, mock, cleanCase.containerID, cleanCase.removeContainer)
		if err != nil {
			t.Errorf("%v when cleaning container", err)
		}
	}
}
