package cluster

import (
	"fmt"
	"io/ioutil"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/generated/models"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Config represents Cluster configuration.
type Config struct {
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
	// UUID of cluster to be managed.
	ClusterID string `yaml:"cluster_id,omitempty"`
	// Action to the performed with the cluster (values: create, update, delete).
	Action string `yaml:"cluster_action,omitempty"`
}

// Cluster represents Cluster service.
type Cluster struct {
	managerType string
	config      *Config
	APIServer   *apisrv.Client
	log         *logrus.Entry
}

// Data is the representation of cluster details.
type Data struct {
	clusterInfo *models.ContrailCluster
	nodesInfo   []*models.Node
	// TODO (ijohnson): Add gce/aws/kvm info
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
	s := apisrv.NewClient(
		c.Endpoint,
		c.AuthURL,
		c.ID,
		c.Password,
		&keystone.Scope{
			Project: &keystone.Project{
				ID: c.ProjectID,
			},
		},
	)

	t := "daemon"
	if c.ClusterID != "" && c.Action != "" {
		t = "oneshot"
	} else if c.ClusterID != "" && c.Action == "" {
		return nil, fmt.Errorf("Action not specified in the config for oneshot manager")
	} else if c.Action != "" && c.ClusterID == "" {
		return nil, fmt.Errorf("Cluster ID not specified in the config for oneshot manager")
	}

	return &Cluster{
		managerType: t,
		APIServer:   s,
		config:      c,
		log:         pkglog.NewLogger("cluster"),
	}, nil
}

// Manage starts managing the clusters.
func (c *Cluster) Manage() error {
	c.log.Info("Start managing contrail clusters")
	err := c.APIServer.Login()
	if err != nil {
		return fmt.Errorf("login to API Server failed: %s", err)
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
