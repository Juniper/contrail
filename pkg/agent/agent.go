package agent

import (
	"context"
	"sync"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/config"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	syncp "github.com/Juniper/contrail/pkg/sync"
)

type handler interface {
	processCluster(e *services.Event, c *Config) error
	processCloud(e *services.Event, c *Config) error
}

// Config holds info
type Config struct {
	// ID of Agent account.
	ID string `yaml:"id"`
	// Password of Agent account.
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
	// Server schema path
	SchemaRoot string `yaml:"schema_root"`
	// Logging level
	LogLevel string `yaml:"log_level"`
	// Backend specifies backend to be used (values: "file").
	Backend string `yaml:"backend"`
	// Watcher specifies resource event watching strategy to be used (values: "polling").
	Watcher string `yaml:"watcher"`
	// InSecure https connection to endpoint
	InSecure bool `yaml:"insecure"`
	// Enabled
	Enabled bool `yaml:"enabled"`
	// Service user name for keystone
	ServiceUserID string `yaml:"service_user_id"`
	// Service user password for keystone
	ServiceUserPassword string `yaml:"service_user_password"`
}

// Agent is here
type Agent struct {
	config             *Config
	handler            handler
	serviceWaitGroup   *sync.WaitGroup
	stopServiceContext context.CancelFunc
	log                *logrus.Entry
}

// NewAgentByConfig creates Agent reading configuration from viper config.
func NewAgentByConfig() (*Agent, error) {
	var c Config
	err := config.LoadConfig("agent", &c)
	if err != nil {
		return nil, err
	}
	c.ID = viper.GetString("client.id")
	c.Password = viper.GetString("client.password")
	c.DomainID = viper.GetString("client.domain_id")
	c.ProjectID = viper.GetString("client.project_id")
	c.DomainName = viper.GetString("client.domain_name")
	c.ProjectName = viper.GetString("client.project_name")
	c.AuthURL = viper.GetString("keystone.authurl")
	c.InSecure = viper.GetBool("insecure")
	c.SchemaRoot = viper.GetString("client.schema_root")
	c.Endpoint = viper.GetString("client.endpoint")
	c.ServiceUserID = viper.GetString("keystone.service_user.id")
	c.ServiceUserPassword = viper.GetString("keystone.service_user.password")

	return NewAgent(&c)
}

// NewAgent creates Agent with given configuration.
func NewAgent(c *Config) (*Agent, error) {
	if err := logutil.Configure(c.LogLevel); err != nil {
		return nil, err
	}

	return &Agent{
		config:  c,
		handler: newEventHandler(),
		log:     logutil.NewLogger("agent"),
	}, nil
}

// Start replication service
func (a *Agent) Start() error {
	processor := &services.EventListProcessor{
		EventProcessor:    a,
		InTransactionDoer: services.NoTransaction,
	}
	producer, err := syncp.NewEventProducer("agent", processor)
	if err != nil {
		return err
	}
	defer producer.Close()

	var ctx context.Context
	ctx, a.stopServiceContext = context.WithCancel(context.Background())

	return producer.Start(ctx)
}

// Process processes event by sending requests to all registered clusters.
func (a *Agent) Process(ctx context.Context, e *services.Event) (*services.Event, error) { //nolint: gocyclo
	a.log.Infof("Received event: %v", e)
	if e == nil {
		return nil, nil
	}

	switch e.Kind() {
	// determine contrail cluster event kind
	// TODO(dji): shift command execution to function call
	case "contrail-cluster", "rhospd-cloud-manager":
		return e, a.handler.processCluster(e, a.config)
	case "cloud":
		return e, a.handler.processCloud(e, a.config)
	}

	return e, nil
}

// Stop replication routine
func (a *Agent) Stop() {
	a.stopServiceContext()
}
