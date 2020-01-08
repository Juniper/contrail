package deploy

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	asfclient "github.com/Juniper/asf/pkg/client"
	yaml "gopkg.in/yaml.v2"
)

const oneShotMgr = "oneshot"

// Config represents Deploy configuration.
type Config struct { // nolint: maligned
	// ID of Deploy account.
	ID string `yaml:"id"`
	// Password of Deploy account.
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
	// Resource type to be managed.
	ResourceType string `yaml:"resource_type,omitempty"`
	// UUID of resource to be managed.
	ResourceID string `yaml:"resource_id,omitempty"`
	// Action to the performed with the resource (values: create, update, delete).
	Action string `yaml:"resource_action,omitempty"`
	// Provisioning tool used to provision the resource (values: ansible, helm).
	ProvisionerType string `yaml:"provisioner_type,omitempty"`
	// Logging level
	LogLevel string `yaml:"log_level"`
	// Logging  file
	LogFile string `yaml:"log_file"`
	// Template root directory
	TemplateRoot string `yaml:"template_root"`
	// Service user name for keystone
	ServiceUserID string `yaml:"service_user_id"`
	// Service user password for keystone
	ServiceUserPassword string `yaml:"service_user_password"`

	// Optional ansible sudo password
	AnsibleSudoPass string `yaml:"ansible_sudo_pass"`
	// Optional ansible deployer cherry pick url
	AnsibleFetchURL string `yaml:"ansible_fetch_url"`
	// Optional ansible deployer cherry pick revison(commit id)
	AnsibleCherryPickRevision string `yaml:"ansible_cherry_pick_revision"`
	// Optional ansible deployer revision(commit id)
	AnsibleRevision string `yaml:"ansible_revision"`
	// Optional Test var to run command in test mode
	Test bool `yaml:"test"`
}

// Deploy represents Deploy service.
type Deploy struct {
	managerType  string
	config       *Config
	APIServer    *client.HTTP
	log          *logrus.Entry
	streamServer *logutil.StreamServer
}

// NewDeployManager creates Deploy reading configuration from given file.
func NewDeployManager(configPath string) (*Deploy, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.UnmarshalStrict(data, &c)
	if err != nil {
		return nil, err
	}

	return NewDeploy(&c)
}

// NewDeploy creates Deploy with given configuration.
func NewDeploy(c *Config) (*Deploy, error) {
	if err := logutil.Configure(c.LogLevel); err != nil {
		return nil, err
	}

	s := client.NewHTTP(&asfclient.HTTPConfig{
		ID:       c.ID,
		Password: c.Password,
		Endpoint: c.Endpoint,
		AuthURL:  c.AuthURL,
		Scope: keystone.NewScope(
			c.DomainID,
			c.DomainName,
			c.ProjectID,
			c.ProjectName,
		),
		Insecure: c.InSecure,
	})

	if c.ResourceID != "" {
		return nil, fmt.Errorf("action not specified in the config for oneshot manager")
	}
	if c.Action != "" {
		return nil, fmt.Errorf("resource ID not specified in the config for oneshot manager")
	}

	return &Deploy{
		managerType:  oneShotMgr,
		APIServer:    s,
		config:       c,
		log:          logutil.NewFileLogger("deploy", c.LogFile),
		streamServer: logutil.NewStreamServer(c.LogFile),
	}, nil
}

// Manage starts managing the resource.
func (c *Deploy) Manage() error {
	c.streamServer.Serve()
	defer c.streamServer.Close()

	c.log.Infof("start handling %s", c.config.ResourceType)
	if err := c.APIServer.Login(context.Background()); err != nil {
		return errors.Wrap(err, "login to API Server failed")
	}

	manager, err := newManager(c)
	if err != nil {
		return err
	}
	err = manager.manage()
	if err != nil {
		return err
	}
	c.log.Infof("stop handling %s", c.config.ResourceType)
	return nil
}
