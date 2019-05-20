package cloud

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/logutil/report"
)

// Config represents cloud configuration needed by cloudManager
// this config is read from the yaml file </var/tmp/cloud/config/<uuid>/contrail-cloud-config.yml
type Config struct { // nolint: maligned
	// ID of cloud
	ID string `yaml:"id"`
	// Password of Cluster account.
	Password string `yaml:"password"`
	// DomainID is ID of keystone domain used for authentication.
	DomainID string `yaml:"domain_id"`
	// ProjectID is ID of keystone project used for authentication.
	ProjectID string `yaml:"project_id"`
	// DomainName is Name of keystone domain used for authentication.
	DomainName string `yaml:"domain_name"`
	// ProjectName is Name of keystone project used for authentication.
	ProjectName string `yaml:"project_name"`
	// AuthURL defines authentication URL.
	AuthURL string `yaml:"auth_url"`
	// Endpoint of API Server.
	Endpoint string `yaml:"endpoint"`
	// InSecure https connection to endpoint
	InSecure bool `yaml:"insecure"`
	// UUID of cloud to be managed.
	CloudID string `yaml:"cloud_id,omitempty"`
	// Action to the performed with the cloud (values: create, update, delete).
	Action string `yaml:"cloud_action,omitempty"`
	// Logging level
	LogLevel string `yaml:"log_level"`
	// Logging  file
	LogFile string `yaml:"log_file"`
	// Template root directory
	TemplateRoot string `yaml:"template_root"`

	// Optional Test var to run cloud in test mode
	Test bool `yaml:"test"`
}

// Cloud represents fields needed to process a cloud request
type Cloud struct {
	config       *Config
	APIServer    *client.HTTP
	log          *logrus.Entry
	reporter     *report.Reporter
	streamServer *logutil.StreamServer
	ctx          context.Context
}

// NewCloudManager returns cloud field by reading the config file
// which is needed to process a cloud request
func NewCloudManager(configPath string) (*Cloud, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.UnmarshalStrict(data, &c)
	if err != nil {
		return nil, err
	}

	return NewCloud(&c)
}

// NewCloud returns a new Cloud instance
func NewCloud(c *Config) (*Cloud, error) {
	if err := logutil.Configure(c.LogLevel); err != nil {
		return nil, err
	}

	// initialize the apiserver client
	s := &client.HTTP{
		Endpoint: c.Endpoint,
		InSecure: c.InSecure,
	}

	// by default create no auth context
	ctx := auth.NoAuth(context.Background())

	// when auth is enabled
	if c.AuthURL != "" {
		s.AuthURL = c.AuthURL
		s.ID = c.ID
		s.Password = c.Password
		s.Scope = keystone.NewScope(c.DomainID, c.DomainName,
			c.ProjectID, c.ProjectName)

		// as auth is enabled, create ctx with auth
		varCtx := auth.NewContext(c.DomainID, c.ProjectID,
			c.ID, []string{c.ProjectName}, "", auth.NewObjPerms(nil))
		var authKey interface{} = "auth"
		ctx = context.WithValue(context.Background(), authKey, varCtx)
	}
	s.Init()

	if c.CloudID != "" && c.Action == "" {
		return nil, fmt.Errorf("action not specified in the config")
	} else if c.CloudID == "" && c.Action != "" {
		return nil, fmt.Errorf("cloudID not specified in the config")
	}

	return &Cloud{
		APIServer: s,
		config:    c,
		log:       logutil.NewFileLogger("cloud", c.LogFile),
		reporter: report.NewReporter(
			s,
			fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.CloudID),
			logutil.NewFileLogger("reporter", c.LogFile),
		),
		streamServer: logutil.NewStreamServer(c.LogFile),
		ctx:          ctx,
	}, nil
}

// Manage acts on every cloud request trigerred by agent
func (c *Cloud) Manage() error {
	c.streamServer.Serve()
	defer c.streamServer.Close()

	// check if request is a delete cloud request
	// and trigger delete cloud workflow if its a delete cloud request
	isDeleteReq, err := c.isCloudDeleteRequest()
	if err != nil {
		c.log.Errorf("cloud %s processing failed with error: %s",
			c.config.CloudID, err)
		return err
	} else if err == nil && isDeleteReq {
		err = c.delete()
		if err != nil {
			c.log.Errorf("delete cloud %s failed with error: %s",
				c.config.CloudID, err)
		}
		return err
	}

	switch c.config.Action {
	case createAction:
		// trigger cloud create workflow
		err = c.create()
		if err != nil {
			c.log.Errorf("create cloud %s failed with error: %s",
				c.config.CloudID, err)
			return err
		}
	case updateAction:
		// trigger cloud update workflow
		err = c.update()
		if err != nil {
			c.log.Errorf("update cloud %s failed with error: %s",
				c.config.CloudID, err)
			return err
		}
	}

	return nil
}

// create executes create cloud workflow
func (c *Cloud) create() error {

	status := map[string]interface{}{statusField: statusCreateProgress}

	// initializes all data
	topo, secret, data, err := c.initialize()
	if err != nil {
		status[statusField] = statusCreateFailed
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}

	// skip is cloud is already created
	if isCloudCreated(c.config.Action, data) {
		c.log.Infof("Cloud %s already created", c.config.CloudID)
		return nil
	}

	c.log.Infof("Starting %s of cloud: %s", c.config.Action, data.info.FQName)

	c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
	status[statusField] = statusCreateFailed

	// create input files and builds public cloud infra using terraform
	err = c.execMCDeployerWorkflow(data.isCloudPublic(), topo, secret)
	if err != nil {
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}

	// update IP details only when cloud is public
	// basically when instances created by terraform
	if data.isCloudPublic() && (!c.config.Test) {
		err = c.updateIPDetails(c.ctx, data)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
	}

	// report the status to be created
	status[statusField] = statusCreated
	c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)

	return nil

}

// handles update of cloud
// nolint: gocyclo
func (c *Cloud) update() error {

	status := map[string]interface{}{statusField: statusUpdateProgress}

	// initializes all data
	topo, secret, data, err := c.initialize()
	if err != nil {
		status[statusField] = statusUpdateFailed
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}

	// check if update request is valid by comparing the topology file
	// skip this request if there is no change in topology file
	if !isCloudUpdateRequest(c.config.Action, data) {
		var topoIsAlreadyUpdated bool
		topoIsAlreadyUpdated, err = topo.isUpdated(defaultCloudResource)
		if err != nil {
			status[statusField] = statusUpdateFailed
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
		if topoIsAlreadyUpdated {
			return nil
		}
	}

	c.log.Infof("Starting %s of cloud: %s", c.config.Action, data.info.FQName)

	c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
	status[statusField] = statusUpdateFailed

	// create input files and builds public cloud infra using terraform
	err = c.execMCDeployerWorkflow(data.isCloudPublic(), topo, secret)
	if err != nil {
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}

	if data.isCloudPublic() && (!c.config.Test) {
		// update IP address for nodes created on public clouds
		err = c.updateIPDetails(c.ctx, data)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
	}
	// report the status to be updated
	status[statusField] = statusUpdated
	c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
	return nil
}

func (c *Cloud) delete() error {

	// get cloud data
	data, err := GetCloudData(c.ctx, c.config.CloudID, c.APIServer, true)
	if err != nil {
		return err
	}

	status := map[string]interface{}{statusField: statusUpdateFailed}

	if data.isCloudPrivate() {
		// wait for the contrail cluster(back ref of cloud) to complete processing
		err = c.verifyContrailClusterStatus(data)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
	}

	if data.isCloudPublic() {
		if tfStateOutputExists(c.config.CloudID) {
			// depending upon the config action, it takes respective terraform action
			err = c.manageTerraform(deleteAction)
			if err != nil {
				c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
				return err
			}
		}
	}

	// delete all the objects referred/in-tree of this cloud object
	err = c.deleteAPIObjects(data)
	if err != nil {
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}

	return os.RemoveAll(GetCloudDir(c.config.CloudID))

}

// verifyContrailClusterStatus verifies and wait for cluster status to be updated
func (c *Cloud) verifyContrailClusterStatus(data *Data) error {

	for _, clusterRef := range data.info.ContrailClusterBackRefs {
		err := waitForClusterStatusToBeUpdated(c.ctx, c.log,
			c.APIServer, clusterRef.UUID)
		if err != nil {
			return err
		}
	}
	return nil
}

// initialize inits data, topology and secret receiver
func (c *Cloud) initialize() (*topology, *secret, *Data, error) {
	// get cloud data
	data, err := GetCloudData(c.ctx, c.config.CloudID, c.APIServer, false)
	if err != nil {
		return nil, nil, nil, err
	}

	// initialize topo receiver
	topo := c.newTopology(data)

	if data.isCloudPrivate() {
		return topo, nil, data, nil
	}

	// initialize secret receiver
	secret, err := c.newSecret()
	if data.isCloudPublic() {
		if err != nil {
			return nil, nil, nil, err
		}
		err = secret.updateFileConfig(data)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return topo, secret, data, nil
}

// isCloudDeleteRequest check if request is delete
func (c *Cloud) isCloudDeleteRequest() (bool, error) {

	cloudObj, err := GetCloud(c.ctx, c.APIServer, c.config.CloudID)
	if err != nil {
		return false, err
	}

	if c.config.Action == updateAction &&
		cloudObj.ProvisioningAction == deleteCloudAction &&
		cloudObj.ProvisioningState == statusNoState {
		return true, nil
	}
	return false, nil
}

// execMCDeployerWorkflow executes deployer workflow which includes of
// 1. create/update topology file
// 2. create/update secret file
// 3. for public cloud type, run generate topology and execute terraform commands
func (c *Cloud) execMCDeployerWorkflow(isPublic bool, topo *topology, secret *secret) error {

	// create topology file(/var/tmp/cloud/<uuid>topology.yml) for cloud
	err := topo.createTopologyFile(GetTopoFile(topo.cloud.config.CloudID))
	if err != nil {
		return err
	}

	//TODO(madhukar) handle if key-pair changes or aws-key

	if isPublic {
		// create secret file(/var/tmp/cloud/<uuid>topology.yml) for public cloud only
		err = secret.createSecretFile()
		if err != nil {
			return err
		}

		// depending upon the config action, it takes respective terraform action
		err = c.manageTerraform(c.config.Action)
		if err != nil {
			return err
		}
	}
	return nil
}
