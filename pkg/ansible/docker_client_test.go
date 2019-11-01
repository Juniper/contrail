package ansible

import (
	"bufio"
	"context"
	"errors"
	"io"
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

func TestPullImageAuthentication(t *testing.T) {
	pullCases := map[string]struct {
		imageRegistry    string
		username         string
		password         string
		expectedUsername string
		expectedPassword string
		expectedAuth     string
		fail             bool
	}{
		"Correct credentials": {
			imageRegistry:    "busybox",
			username:         "admin",
			password:         "123456",
			expectedUsername: "admin",
			expectedPassword: "123456",
			expectedAuth:     "eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIxMjM0NTYifQ==", // for expected username admin and password 123456
			fail:             false,
		},
		"Wrong password": {
			imageRegistry:    "busybox",
			username:         "admin",
			password:         "654321",
			expectedUsername: "admin",
			expectedPassword: "123456",
			expectedAuth:     "eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIxMjM0NTYifQ==",
			fail:             true,
		},
		"Wrong username": {
			imageRegistry:    "busybox",
			username:         "nimda",
			password:         "123456",
			expectedUsername: "admin",
			expectedPassword: "123456",
			expectedAuth:     "eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIxMjM0NTYifQ==",
			fail:             true,
		},
		"Wrong usernam and password": {
			imageRegistry:    "busybox",
			username:         "nimda",
			password:         "654321",
			expectedUsername: "admin",
			expectedPassword: "123456",
			expectedAuth:     "eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIxMjM0NTYifQ==",
			fail:             true,
		},
	}
	for name, pullCase := range pullCases {
		t.Run(name, func(t *testing.T) {
			mockRC := &mockReadCloser{Reader: strings.NewReader("Success")}
			err := pullImage(
				context.Background(),
				&mockAPIClient{
					expectedAuth: pullCase.expectedAuth,
					imageReader:  mockRC,
				},
				pullCase.imageRegistry,
				pullCase.username,
				pullCase.password,
			)
			if pullCase.fail {
				if err == nil {
					t.Errorf("Pulling image succeeded with username %s and password %s, expected username %s and password %s", pullCase.username, pullCase.password, pullCase.expectedUsername, pullCase.expectedPassword)
				}
			} else {
				if err != nil {
					t.Errorf("Pulling image failed with username %s and password %s, expected username %s and password %s", pullCase.username, pullCase.password, pullCase.expectedUsername, pullCase.expectedPassword)
				}
				if !mockRC.isClosed {
					t.Errorf("ReadCloser form cli.imagePull() didn't close successfully")
				}
			}
		})
	}
}

func TestSearchContainer(t *testing.T) {
	searchCases := map[string]struct {
		imageName                 string
		containerInstances        []types.Container
		expectedContainerInstance *types.Container
		found                     bool
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
			expectedContainerInstance: &types.Container{
				ID:    "3",
				Names: []string{"3"},
				Image: "busybox",
			},
			found: true,
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
			expectedContainerInstance: nil,
			found:                     false,
		},
	}
	for name, searchCase := range searchCases {
		t.Run(name, func(t *testing.T) {
			containerInstance, found, err := searchContainer(
				context.Background(),
				&mockAPIClient{
					containerInstances: searchCase.containerInstances,
				},
				searchCase.imageName,
			)
			if err != nil {
				t.Errorf("%v when searching for container", err)
			}
			if searchCase.found {
				if !found {
					t.Errorf("Expected to find existing container, got nothing")
				}
				if !reflect.DeepEqual(containerInstance, searchCase.expectedContainerInstance) {
					t.Errorf("Expected %#+v, got %#+v", searchCase.expectedContainerInstance, containerInstance)
				}
			} else {
				if found {
					t.Errorf("Expected nothing, got an existing container")
				}
			}
		})
	}
}

func TestStartExistingContainer(t *testing.T) {
	startCases := map[string]struct {
		containerInstance     *types.Container
		imageName             string
		expectedContainerName string
	}{
		"Existing container": {
			containerInstance: &types.Container{
				ID:    "1",
				Names: []string{"1"},
			},
			imageName:             "busybox",
			expectedContainerName: "1",
		},
	}
	for name, startCase := range startCases {
		t.Run(name, func(t *testing.T) {
			containerName, err := startExistingContainer(
				context.Background(),
				&mockAPIClient{
					containerInstance:     startCase.containerInstance,
					expectedContainerName: startCase.expectedContainerName,
				},
				startCase.containerInstance,
			)
			if err != nil {
				t.Errorf("%v when starting container", err)
			}
			if containerName != startCase.expectedContainerName {
				t.Errorf("Expected container name %s, got container name %s", startCase.expectedContainerName, containerName)
			}
		})
	}
}

func TestCreatRunningContainer(t *testing.T) {
	createCases := map[string]struct {
		imageName                   string
		expectedContainerNamePrefix string
	}{
		"No container": {
			imageName:                   "busybox",
			expectedContainerNamePrefix: "ansible-deployer",
		},
	}
	for name, createCase := range createCases {
		t.Run(name, func(t *testing.T) {
			containerName, err := createRunningContainer(
				context.Background(),
				&mockAPIClient{
					expectedContainerNamePrefix: createCase.expectedContainerNamePrefix,
				},
				createCase.imageName,
			)
			if err != nil {
				t.Errorf("%v when starting container", err)
			}
			if !strings.Contains(containerName, createCase.expectedContainerNamePrefix) {
				t.Errorf("Expected container name prefix %s, container name %s", createCase.expectedContainerNamePrefix, containerName)
			}
		})
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
	for name, execCase := range execCases {
		t.Run(name, func(t *testing.T) {
			err := execCmd(context.Background(), &mockAPIClient{}, execCase.containerID, execCase.command, false)
			if err != nil {
				t.Errorf("%v when executing command in container", err)
			}
		})
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
	for name, cleanCase := range cleanCases {
		t.Run(name, func(t *testing.T) {
			err := cleanContainer(context.Background(), &mockAPIClient{}, cleanCase.containerID, cleanCase.removeContainer)
			if err != nil {
				t.Errorf("%v when cleaning container", err)
			}
		})
	}
}

type mockConnection struct {
	net.Conn
}

func (m mockConnection) Close() error {
	return nil
}

type mockReadCloser struct {
	io.Reader
	isClosed bool
}

func (m *mockReadCloser) Close() error {
	m.isClosed = true
	return nil
}

type mockAPIClient struct {
	expectedAuth                string
	imageReader                 io.ReadCloser
	containerInstances          []types.Container
	containerInstance           *types.Container
	expectedContainerName       string
	expectedContainerNamePrefix string
}

func (m *mockAPIClient) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {
	if options.RegistryAuth != m.expectedAuth {
		return nil, errors.New("Authentication failed")
	}
	return m.imageReader, nil
}

func (m *mockAPIClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return m.containerInstances, nil
}

func (m *mockAPIClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	return container.ContainerCreateCreatedBody{}, nil
}

func (m *mockAPIClient) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	return nil
}

func (m *mockAPIClient) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error) {
	return types.IDResponse{
		ID: "1",
	}, nil
}

func (m *mockAPIClient) ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error) {
	return types.HijackedResponse{
		Conn:   &mockConnection{},
		Reader: bufio.NewReader(strings.NewReader("Output")),
	}, nil
}

func (m *mockAPIClient) ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error {
	return nil
}

func (m *mockAPIClient) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	return nil
}
