package undercloud

import (
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/deploy/base"
	"github.com/Juniper/contrail/pkg/logutil"
)

// Config represents Command configuration.
type Config struct {
	// http client of api server
	APIServer *client.HTTP
	// UUID of resource to be managed.
	ResourceID string
	// Action to the performed with the resource (values: create, update, delete).
	Action string
	// Logging level
	LogLevel string
	// Logging  file
	LogFile string
	// Template root directory
	TemplateRoot string

	// Optional Test var to run command in test mode
	Test bool
	// Work root directory
	WorkRoot string
}

// UnderCloud represents contrail undercloud manager
type UnderCloud struct {
	config    *Config
	APIServer *client.HTTP
	log       *logrus.Entry
}

// NewUnderCloud creates UnderCloud with given configuration.
func NewUnderCloud(c *Config) (*UnderCloud, error) {
	return &UnderCloud{
		config:    c,
		APIServer: c.APIServer,
		log:       logutil.NewFileLogger("undercloud", c.LogFile),
	}, nil
}

// GetDeployer creates new deployer based on the type
func (u *UnderCloud) GetDeployer() (base.Deployer, error) {
	return newDeployerByID(u)
}
