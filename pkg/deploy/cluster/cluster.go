package cluster

import (
	"os/exec"
	"context"
	"path/filepath"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/contrail/pkg/ansible"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/deploy/rhospd/overcloud"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
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
		return c.newMCProvisioner()
	}

	cData, err := data(c)
	if err != nil {
		return nil, err
	}

	switch deployerType(cData, c.config.Action) {
	case "rhospd":
		return newOvercloudDeployer(c)
	case "ansible", "tripleo":
		return newAnsibleDeployer(c, cData), nil
	}
	return nil, errors.New("unsupported deployer type")
}

func (c *Cluster) isMultiCloudRequest() bool {
	if c.config.Action == deleteAction {
		return false
	}
	resp, err := c.APIServer.GetContrailCluster(context.Background(), &services.GetContrailClusterRequest{
		ID: c.config.ClusterID,
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

func newAnsibleDeployer(c *Cluster, cData *base.Data) *contrailAnsibleDeployer {
	d := newDeployCluster(c, cData, "contrail-ansible-deployer")
	return &contrailAnsibleDeployer{
		deployCluster: *d,
		ansibleClient: ansible.NewCLIClient(
			d.Reporter,
			c.config.LogFile,
			d.getWorkingDir(),
			c.config.Test,
		),
	}
}

func (c *Cluster) newMCProvisioner() (*multiCloudProvisioner, error) {
	resp, err := c.APIServer.GetContrailCluster(context.Background(), &services.GetContrailClusterRequest{
		ID: c.config.ClusterID,
		Fields: []string{
			models.ContrailClusterFieldProvisioningState,
			models.ContrailClusterFieldProvisioningAction,
			models.ContrailClusterFieldContainerRegistry,
		},
	})
	if err != nil {
		return nil, err
	}
	// registry, err := cloud.NewAuthorizedRegistry(c.config.)
	// if err != nil {

	// }
	return &multiCloudProvisioner{
		workDir: filepath.Join(c.getWorkRoot(), mcWorkDir),
		clusterUUID: c.config.ClusterID,
		clusterProvisioningState: resp.GetContrailCluster().GetProvisioningState(),
		mcAction: resp.GetContrailCluster().GetProvisioningAction(),
		reporter: newReporter(c),
		log:      logutil.NewFileLogger("multi-cloud-provisioner", c.config.LogFile),
		test: c.config.Test,
		apiServer: c.APIServer,
		clouds: c.getCloudRefs(),
		// contrailRegistries: resp.GetContrailCluster().GetContainerRegistry(),
	}, nil
}

func (c *Cluster) getCloudRefs() []cloudRef {

}

func (c *Cluster) getWorkRoot() string {
	workRoot := c.config.WorkRoot
	if workRoot == "" {
		workRoot = defaultWorkRoot
	}
	return workRoot
}
