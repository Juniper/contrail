package cluster

import (
	"context"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/fileutil/template"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/contrail/pkg/ansible"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/deploy/rhospd/overcloud"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Config represents Command configuration.
type Config struct {
	// http client of api server
	APIServer *client.HTTP
	// UUID of resource to be managed.
	ClusterID string
	// Action to the performed with the cluster (values: create, update, delete).
	Action string
	// Logging level
	LogLevel string
	// Logging  file
	LogFile string
	// Template root directory
	TemplateRoot string
	// Work root directory
	WorkRoot string
	// Service user name for keystone
	ServiceUserID string
	// Service user password for keystone
	ServiceUserPassword string

	// Optional ansible sudo password
	AnsibleSudoPass string
	// Optional ansible deployer cherry pick url
	AnsibleFetchURL string
	// Optional ansible deployer cherry pick revison(commit id)
	AnsibleCherryPickRevision string
	// Optional ansible deployer revision(commit id)
	AnsibleRevision string
	// Optional Test var to run command in test mode
	Test bool
}

// CommandExecutor interface provides methods to execute a command
type CommandExecutor interface {
	ExecCmdAndWait(r *report.Reporter, cmd string, args []string, dir string, envVars ...string) error
	ExecAndWait(r *report.Reporter, cmd *exec.Cmd) error
}

// Cluster represents contrail cluster manager
type Cluster struct {
	config          *Config
	APIServer       *client.HTTP
	log             *logrus.Entry
	commandExecutor CommandExecutor
}

// NewCluster creates Cluster with given configuration.
func NewCluster(c *Config, commandExecutor CommandExecutor) (*Cluster, error) {
	return &Cluster{
		config:          c,
		APIServer:       c.APIServer,
		log:             logutil.NewFileLogger("cluster", c.LogFile),
		commandExecutor: commandExecutor,
	}, nil
}

// GetDeployer creates new deployer based on the type
// TODO(Daniel): this should not be Cluster's method
func (c *Cluster) GetDeployer() (base.Deployer, error) {
	if c.isMultiCloudRequest() {
		cc := c.config
		return NewMultiCloudProvisioner(
			cc.ClusterID, cc.LogFile, filepath.Join(cc.WorkRoot, cc.ClusterID), cc.APIServer, cc.Test, c.commandExecutor,
		)
	}

	cData, err := data(c)
	if err != nil {
		return nil, err
	}

	switch deployerType(cData, c.config.Action) {
	case "rhospd":
		return newOvercloudDeployer(c)
	case "ansible", "tripleo", "juju":
		return newAnsibleDeployer(c, cData), nil
	}
	return nil, errors.New("unsupported deployer type")
}

func (c *Cluster) isMultiCloudRequest() bool {
	if c.config.Action == deleteAction {
		return false
	}
	resp, err := c.APIServer.GetContrailCluster(context.Background(), &services.GetContrailClusterRequest{
		ID:     c.config.ClusterID,
		Fields: []string{models.ContrailClusterFieldIsMulticloud},
	})
	if err != nil {
		return false
	}
	return resp.GetContrailCluster().GetIsMulticloud()
}

func data(c *Cluster) (*base.Data, error) {
	if c.config.Action == deleteAction {
		return &base.Data{Reader: c.APIServer}, nil
	}
	return base.NewResourceManager(c.APIServer, c.config.LogFile).GetClusterDetails(c.config.ClusterID)
}

func deployerType(cData *base.Data, action string) string {
	if cData.ClusterInfo != nil && cData.ClusterInfo.ProvisionerType != "" {
		return cData.ClusterInfo.ProvisionerType
	}

	return defaultDeployer
}

func newOvercloudDeployer(c *Cluster) (base.Deployer, error) {
	o, err := overcloud.NewOverCloud(&overcloud.Config{
		APIServer:    c.APIServer,
		ResourceID:   c.config.ClusterID,
		Action:       c.config.Action,
		TemplateRoot: c.config.TemplateRoot,
		LogLevel:     c.config.LogLevel,
		LogFile:      c.config.LogFile,
	})
	if err != nil {
		return nil, err
	}

	return o.GetDeployer()
}

func newAnsibleDeployer(c *Cluster, cData *base.Data) (*contrailAnsibleDeployer, error) {
	d := newDeployCluster(c, cData, "contrail-ansible-deployer")

	// TODO(dji): move dependency injection to testing code
	containerPlayer, err := getContainerPlayer(c, d)
	if err != nil {
		return nil, errors.Wrap(err, "New container player creation failed")
	}

	return &contrailAnsibleDeployer{
		deployCluster: *d,
		ansibleClient: ansible.NewCLIClient(
			d.Reporter,
			c.config.LogFile,
			d.getWorkingDir(),
			c.config.Test,
		),
		containerPlayer: containerPlayer,
	}, nil
}

func getContainerPlayer(c *Cluster, d *deployCluster) (Player, error) {
	if c.config.Test {
		return newMockContainerPlayer(d.getWorkingDir())
	}
	return ansible.NewContainerPlayer(d.Reporter, c.config.LogFile)
}

// TODO(dji): move to testing code and inject as dependency
type mockContainerPlayer struct {
	workingDirectory string
}

func newMockContainerPlayer(workingDirectory string) (*mockContainerPlayer, error) {
	return &mockContainerPlayer{workingDirectory: workingDirectory}, nil
}

func (m *mockContainerPlayer) Play(
	ctx context.Context,
	imageRef string,
	imageRefUsername string,
	imageRefPassword string,
	workRoot string,
	ansibleBinaryRepo string,
	ansibleArgs []string,
	keepContainerAlive bool,
) error {
	playBookIndex := len(ansibleArgs) - 1
	content, err := template.Apply("./test_data/test_ansible_playbook.tmpl", pongo2.Context{
		"playBook":    ansibleArgs[playBookIndex],
		"ansibleArgs": strings.Join(ansibleArgs[:playBookIndex], " "),
	})
	if err != nil {
		return err
	}

	return fileutil.AppendToFile(
		filepath.Join(m.workingDirectory, "executed_ansible_playbook.yml"),
		content,
		0600,
	)
}
