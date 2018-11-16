package cluster

import (
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/deploy/base"
	pkglog "github.com/Juniper/contrail/pkg/log"
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

// Cluster represents contrail cluster manager
type Cluster struct {
	config    *Config
	APIServer *client.HTTP
	log       *logrus.Entry
}

// NewCluster creates Cluster with given configuration.
func NewCluster(c *Config) (*Cluster, error) {
	// create logger for cluster
	logger := pkglog.NewFileLogger("cluster", c.LogFile)
	pkglog.SetLogLevel(logger, c.LogLevel)

	return &Cluster{
		config:    c,
		APIServer: c.APIServer,
		log:       logger,
	}, nil
}

// GetDeployer creates new deployer based on the type
func (c *Cluster) GetDeployer() (base.Deployer, error) {
	return newDeployerByID(c)
}
