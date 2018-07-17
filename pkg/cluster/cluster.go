package cluster

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/common"
	pkglog "github.com/Juniper/contrail/pkg/log"
)

// Config represents Cluster configuration.
type Config struct { // nolint: maligned
	// ID of Cluster account.
	ID string `yaml:"id"`
	// Password of Cluster account.
	Password string `yaml:"password"`
	// ProjectID is ID of keystone project used for authentication.
	ProjectID string `yaml:"project_id"`
	// AuthURL defines authentication URL.
	AuthURL string `yaml:"auth_url"`
	// Endpoint of API Server.
	Endpoint string `yaml:"endpoint"`
	// InSecure https connection to endpoint
	InSecure bool `yaml:"insecure"`
	// UUID of cluster to be managed.
	ClusterID string `yaml:"cluster_id,omitempty"`
	// Action to the performed with the cluster (values: create, update, delete).
	Action string `yaml:"cluster_action,omitempty"`
	// Provisioning tool used to provision the cluster (values: ansible, helm).
	ProvisionerType string `yaml:"provisioner_type,omitempty"`
	// Logging level
	LogLevel string `yaml:"log_level"`
	// Template root directory
	TemplateRoot string `yaml:"template_root"`

	// Optional ansible sudo password
	AnsibleSudoPass string `yaml:"ansible_sudo_pass"`
	// Optional ansible deployer cherry pick url
	AnsibleFetchURL string `yaml:"ansible_fetch_url"`
	// Optional ansible deployer cherry pick revison(commit id)
	AnsibleCherryPickRevision string `yaml:"ansible_cherry_pick_revision"`
	// Optional ansible deployer revision(commit id)
	AnsibleRevision string `yaml:"ansible_revision"`
	// Optional Test var to run cluster in test mode
	Test bool `yaml:"test"`
}

// Cluster represents Cluster service.
type Cluster struct {
	managerType string
	config      *Config
	APIServer   *client.HTTP
	log         *logrus.Entry
}

// NewClusterManager creates Cluster reading configuration from given file.
func NewClusterManager(configPath string) (*Cluster, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.UnmarshalStrict(data, &c)
	if err != nil {
		return nil, err
	}

	return NewCluster(&c)
}

// NewCluster creates Cluster with given configuration.
func NewCluster(c *Config) (*Cluster, error) {
	s := &client.HTTP{
		Endpoint: c.Endpoint,
		InSecure: c.InSecure,
	}
	// auth enabled
	if c.AuthURL != "" {
		s.AuthURL = c.AuthURL
		s.ID = c.ID
		s.Password = c.Password
		s.Scope = &keystone.Scope{
			Project: &keystone.Project{
				ID: c.ProjectID,
			},
		}
	}
	s.Init()

	t := "daemon"
	if c.ClusterID != "" && c.Action != "" {
		t = "oneshot"
	} else if c.ClusterID != "" && c.Action == "" {
		return nil, fmt.Errorf("Action not specified in the config for oneshot manager")
	} else if c.Action != "" && c.ClusterID == "" {
		return nil, fmt.Errorf("Cluster ID not specified in the config for oneshot manager")
	}

	// create logger for cluster
	logger := pkglog.NewLogger("cluster")
	pkglog.SetLogLevel(logger, c.LogLevel)

	return &Cluster{
		managerType: t,
		APIServer:   s,
		config:      c,
		log:         logger,
	}, nil
}

// Manage starts managing the clusters.
func (c *Cluster) Manage() error {
	common.SetLogLevel()
	c.log.Info("Start managing contrail clusters")
	if c.config.AuthURL != "" {
		err := c.APIServer.Login()
		if err != nil {
			return fmt.Errorf("login to API Server failed: %s", err)
		}
	}

	manager, err := newManager(c)
	if err != nil {
		return err
	}
	err = manager.manage()
	if err != nil {
		return err
	}
	c.log.Info("Stop managing contrail clusters")
	return nil
}
