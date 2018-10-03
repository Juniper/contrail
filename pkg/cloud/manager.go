package cloud

import (
	"io/ioutil"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/common"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// Config represents cloud configuration needed by cloudManager
type Config struct { // nolint: maligned
	// ID of cloud
	ID string `yaml:"id"`
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
	s.Init()

	//TODO(madhukar) - Handle extension of existing contrail cluster to public
	// cloud, which involves taking care of keystone auth

	// create logger for cloud
	logger := pkglog.NewFileLogger("cloud", c.LogFile)
	pkglog.SetLogLevel(logger, c.LogLevel)
	streamServer := pkglog.NewStreamServer(c.LogFile)

	return &Cloud{
		APIServer:    s,
		config:       c,
		log:          logger,
		streamServer: streamServer,
	}, nil
}

// Manage starts managing the cloud.
func (c *Cloud) Manage() error {
	common.SetLogLevel()
	// start log server
	c.streamServer.Serve()
	defer c.streamServer.Close()
	c.log.Info("Start managing public clouds")

	//TODO(madhukar) - login to API server if InSecure is false

	cloudProvisioner, err := newProvisioner(c)
	if err != nil {
		return err
	}
	err = cloudProvisioner.modifyTopology()
	if err != nil {
		return err
	}
	err = cloudProvisioner.selectCloudCredential()
	if err != nil {
		return err
	}
	err = cloudProvisioner.invokeTerraform()
	if err != nil {
		return err
	}
	c.log.Info("Stop managing public clouds")
	return nil
}
