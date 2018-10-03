package cloud

import (
	"fmt"
	"io/ioutil"

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
		return nil, fmt.Errorf("CloudID not specified in the config")
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
	status := map[string]interface{}{statusField: statusCreateProgress}
	c.reporter.ReportStatus(status)

	c.log.Info("Start managing public clouds")

	// Manage topology
	status[statusField] = statusCreateFailed
	topo, err := c.getTopology()
	if err != nil {
		c.reporter.ReportStatus(status)
		return err
	}

	// Manage secret file
	sec, err := c.getSecret(topo.cloudData)
	if err != nil {
		c.reporter.ReportStatus(status)
		return err
	}

	cloudType, err := c.getCloudType()
	if err != nil {
		c.reporter.ReportStatus(status)
		return err
	}

	//Authenticate for azure cloud type
	if cloudType == "azure" {
		err = c.authenticate(topo.cloudData)
		if err != nil {
			c.reporter.ReportStatus(status)
			return err
		}
	}

	err = c.manageTerraform(topo, sec)
	if err != nil {
		c.reporter.ReportStatus(status)
		return err
	}

	c.log.Info("Done managing public clouds")
	return nil
}
