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

// Cloud represents cloud service.
type Cloud struct {
	config       *Config
	APIServer    *client.HTTP
	log          *logrus.Entry
	reporter     *report.Reporter
	streamServer *logutil.StreamServer
	ctx          context.Context
}

// NewCloudManager creates cloud fields by reading config from given configPath
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
			c.ID, []string{c.ProjectName}, "")
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

// Manage starts managing the cloud.
func (c *Cloud) Manage() error {
	c.streamServer.Serve()
	defer c.streamServer.Close()

	if isDeleteReq, err := c.isCloudDeleteRequest(); err == nil && isDeleteReq {
		return c.delete()
	}

	switch c.config.Action {
	case createAction:
		err := c.create()
		if err != nil {
			return err
		}
	case updateAction:
		err := c.update()
		if err != nil {
			return err
		}
	}

	return nil
}

// handles creation of cloud
func (c *Cloud) create() error {

	status := map[string]interface{}{statusField: statusCreateProgress}

	// Run pre-install steps
	topo, secret, data, err := c.initialize()
	if err != nil {
		status[statusField] = statusCreateFailed
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}

	if data.isCloudCreated() {
		return nil
	}

	c.log.Infof("Starting %s of cloud: %s", c.config.Action, data.info.FQName)

	c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
	status[statusField] = statusCreateFailed

	err = topo.createTopologyFile(GetTopoFile(c.config.CloudID))
	if err != nil {
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}

	if data.isCloudPublic() {
		err = secret.createSecretFile()
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
		// depending upon the config action, it takes respective terraform action
		err = c.manageTerraform(c.config.Action)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
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

	status[statusField] = statusCreated
	c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)

	return nil

}

// handles update of cloud
// nolint: gocyclo
func (c *Cloud) update() error {

	status := map[string]interface{}{statusField: statusUpdateProgress}

	// Run pre-install steps
	topo, secret, data, err := c.initialize()
	if err != nil {
		status[statusField] = statusUpdateFailed
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}

	if !data.isCloudUpdateRequest() {
		topoIsAlreadyUpdated, err := topo.isUpdated(defaultCloudResource)
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

	err = topo.createTopologyFile(GetTopoFile(topo.cloud.config.CloudID))
	if err != nil {
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}

	//TODO(madhukar) handle if key-pair changes or aws-key

	if data.isCloudPublic() {

		err = secret.createSecretFile()
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}

		// depending upon the config action, it takes respective terraform action
		err = c.manageTerraform(c.config.Action)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
	}

	//update IP address
	if data.isCloudPublic() && (!c.config.Test) {
		err = c.updateIPDetails(c.ctx, data)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
	}

	status[statusField] = statusUpdated
	c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
	return nil
}

func (c *Cloud) delete() error {

	// get cloud data
	data, err := c.getCloudData()
	if err != nil {
		return err
	}

	if data.isCloudPublic() {
		if tfStateOutputExists(c.config.CloudID) {
			err = c.manageTerraform(deleteAction)
			if err != nil {
				return err
			}
		}
	}

	// delete all the objects referred/in-tree of this cloud object
	err = c.deleteAPIObjects(data)
	if err != nil {
		return err
	}

	return os.RemoveAll(GetCloudDir(c.config.CloudID))

}

func (c *Cloud) initialize() (*topology, *secret, *Data, error) {

	data, err := c.getCloudData()
	if err != nil {
		return nil, nil, nil, err
	}
	topo := c.newTopology(data)

	if data.isCloudPrivate() {
		return topo, nil, data, nil
	}

	// initialize secret struct
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

func (c *Cloud) isCloudDeleteRequest() (bool, error) {

	cloudObj, err := GetCloud(c.ctx, c.APIServer, c.config.CloudID)
	if err != nil {
		return false, err
	}

	if c.config.Action == updateAction &&
		(cloudObj.ProvisioningAction == deleteCloudAction) {
		return true, nil
	}
	return false, nil
}
