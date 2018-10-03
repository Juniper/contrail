package cloud

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	report "github.com/Juniper/contrail/pkg/cluster"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/logging"
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
	// Type of cloud, could be azure, gcp, aws, private
	Type string `yaml:"cloud_type"`
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
	streamServer *pkglog.StreamServer
}

// NewCloudManager creates cloud reading configuration from given file.
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
	s := &client.HTTP{
		Endpoint: c.Endpoint,
		InSecure: c.InSecure,
	}

	// when auth is enabled
	if c.AuthURL != "" {
		s.AuthURL = c.AuthURL
		s.ID = c.ID
		s.Password = c.Password
		s.Scope = client.GetKeystoneScope(c.DomainID, c.DomainName,
			c.ProjectID, c.ProjectName)
	}
	s.Init()

	if c.CloudID != "" && c.Action == "" {
		return nil, fmt.Errorf("action not specified in the config")
	} else if c.CloudID == "" && c.Action != "" {
		return nil, fmt.Errorf("cloudID not specified in the config")
	}

	//create reporter for cloud
	logger := pkglog.NewFileLogger("reporter", c.LogFile)
	pkglog.SetLogLevel(logger, c.LogLevel)

	r := &report.Reporter{
		API:      s,
		Resource: fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.CloudID),
		Log:      logger,
	}

	// create logger for cloud
	logger = pkglog.NewFileLogger("cloud", c.LogFile)
	pkglog.SetLogLevel(logger, c.LogLevel)
	streamServer := pkglog.NewStreamServer(c.LogFile)

	return &Cloud{
		APIServer:    s,
		config:       c,
		log:          logger,
		reporter:     r,
		streamServer: streamServer,
	}, nil
}

// Manage starts managing the cloud.
func (c *Cloud) Manage() error {
	logging.SetLogLevel()
	// start log server
	c.streamServer.Serve()
	defer c.streamServer.Close()

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
	case deleteAction:
		err := c.delete()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Cloud) create() error {

	status := map[string]interface{}{statusField: statusCreateProgress}

	topo, secret, data, err := c.initialize()
	if err != nil {
		status[statusField] = statusCreateFailed
		c.reporter.ReportStatus(status, defaultCloudResource)
		return err
	}

	if data.isCloudCreated() {
		return nil
	}

	c.log.Infof("Starting %s of cloud: %s", c.config.Action, data.cloudInfo.FQName)

	c.reporter.ReportStatus(status, defaultCloudResource)
	status[statusField] = statusCreateFailed

	err = topo.createTopologyFile(topo.getTopoFile())
	if err != nil {
		c.reporter.ReportStatus(status, defaultCloudResource)
		return err
	}

	err = secret.createSecretFile()
	if err != nil {
		c.reporter.ReportStatus(status, defaultCloudResource)
		return err
	}

	err = c.manageTerraform(topo, secret, c.config.Action)
	if err != nil {
		c.reporter.ReportStatus(status, defaultCloudResource)
		return err
	}

	status[statusField] = statusCreated
	c.reporter.ReportStatus(status, defaultCloudResource)

	return nil

}

func (c *Cloud) update() error {

	status := map[string]interface{}{statusField: statusUpdateProgress}

	topo, secret, data, err := c.initialize()
	if err != nil {
		status[statusField] = statusUpdateFailed
		c.reporter.ReportStatus(status, defaultCloudResource)
		return err
	}

	topoUpdated, err := topo.isUpdated()
	if err != nil {
		status[statusField] = statusUpdateFailed
		c.reporter.ReportStatus(status, defaultCloudResource)
		return err
	}

	if topoUpdated {
		return nil
	}

	c.log.Infof("Starting %s of cloud: %s", c.config.Action, data.cloudInfo.FQName)

	c.reporter.ReportStatus(status, defaultCloudResource)
	status[statusField] = statusUpdateFailed

	err = topo.createTopologyFile(topo.getTopoFile())
	if err != nil {
		c.reporter.ReportStatus(status, defaultCloudResource)
		return err
	}

	//TODO(madhukar) handle if key-pair changes or aws-key

	err = c.manageTerraform(topo, secret, c.config.Action)
	if err != nil {
		c.reporter.ReportStatus(status, defaultCloudResource)
		return err
	}

	status[statusField] = statusUpdated
	c.reporter.ReportStatus(status, defaultCloudResource)
	return nil
}

func (c *Cloud) delete() error {

	status := map[string]interface{}{statusField: statusDeleteProgress}

	topo, secret, _, err := c.initialize()
	if err != nil {
		status[statusField] = statusDeleteFailed
		c.reporter.ReportStatus(status, defaultCloudResource)
		return err
	}

	c.reporter.ReportStatus(status, defaultCloudResource)
	status[statusField] = statusDeleteFailed

	err = c.manageTerraform(topo, secret, c.config.Action)
	if err != nil {
		c.reporter.ReportStatus(status, defaultCloudResource)
		return err
	}

	err = c.removeHomeDir()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cloud) initialize() (*topology, *secret, *Data, error) {

	data, err := c.getCloudData()
	if err != nil {
		return nil, nil, nil, err
	}

	topo := c.newTopology(data)

	secret, err := c.newSecret()
	if err != nil {
		return nil, nil, nil, err
	}

	err = secret.updateFileConfig(data)
	if err != nil {
		return nil, nil, nil, err
	}

	// Get CloudType
	cloudType, err := c.getCloudType()
	if err != nil {
		return nil, nil, nil, err
	}

	if cloudType == azure {
		err = c.authenticate(data)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return topo, secret, data, nil
}

func (c *Cloud) removeHomeDir() error {
	return os.RemoveAll(c.getWorkingDir())
}
