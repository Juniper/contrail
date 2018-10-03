package cloud

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/auth"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/log/report"
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
	ctx          context.Context
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

	ctx := auth.NoAuth(context.Background())

	// when auth is enabled
	if c.AuthURL != "" {
		s.AuthURL = c.AuthURL
		s.ID = c.ID
		s.Password = c.Password
		s.Scope = client.GetKeystoneScope(c.DomainID, c.DomainName,
			c.ProjectID, c.ProjectName)

		varCtx := auth.NewContext(c.DomainID, c.ProjectID, c.ID, []string{c.ProjectName})
		var authKey interface{} = "auth"
		ctx = context.WithValue(context.Background(), authKey, varCtx)
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

	r := report.NewReporter(s,
		fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.CloudID), logger)

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
		ctx:          ctx,
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
	}

	return nil
}

func (c *Cloud) create() error {

	status := map[string]interface{}{statusField: statusCreateProgress}

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

		err = c.manageTerraform(c.config.Action)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
	}

	if data.isCloudPublic() && (!c.config.Test) {
		err = c.updateNodeIP(c.ctx, data)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
	}

	status[statusField] = statusCreated
	c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)

	return nil

}

func (c *Cloud) update() error {

	status := map[string]interface{}{statusField: statusUpdateProgress}

	topo, secret, data, err := c.initialize()
	if err != nil {
		status[statusField] = statusUpdateFailed
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}

	if data.isCloudDeleteRequest() {
		return c.delete(data)
	}

	if data.isCloudUpdated() {
		return nil
	}

	topoUpdated, err := topo.isUpdated(defaultCloudResource)
	if err != nil {
		status[statusField] = statusUpdateFailed
		c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
		return err
	}
	if topoUpdated {
		return nil
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

		err = c.manageTerraform(c.config.Action)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
	}

	//update IP address
	if data.isCloudPublic() && (!c.config.Test) {
		err = c.updateNodeIP(c.ctx, data)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
			return err
		}
	}

	status[statusField] = statusUpdated
	c.reporter.ReportStatus(c.ctx, status, defaultCloudResource)
	return nil
}

func (c *Cloud) delete(cData *Data) error {

	var errList []string
	if c.config.Type != onPrem {
		err := c.manageTerraform(c.config.Action)
		// log the error and continue
		if err != nil {
			errList = append(errList, err.Error())
		}
	}

	err := c.deleteAPIObjects(cData)
	if err != nil {
		errList = append(errList, err.Error())
	}

	err = c.removeHomeDir()
	if err != nil {
		errList = append(errList, err.Error())
	}

	if len(errList) > 0 {
		errors.New(strings.Join(errList, "\n"))
	}
	return nil

}

func (c *Cloud) initialize() (*topology, *secret, *Data, error) {

	data, err := c.getCloudData()
	if err != nil {
		return nil, nil, nil, err
	}
	topo := c.newTopology(data)

	if c.config.Type == onPrem {
		return topo, nil, data, nil
	}

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

	if c.config.Type == azure {
		err = c.authenticate(data)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return topo, secret, data, nil
}

func (c *Cloud) removeHomeDir() error {
	return os.RemoveAll(GetCloudDir(c.config.CloudID))
}
