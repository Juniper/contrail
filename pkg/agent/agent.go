package agent

import (
	"context"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/schema"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	asfclient "github.com/Juniper/asf/pkg/client"
)

// Agent constants.
const (
	FileBackend    = "file"
	PollingWatcher = "polling"

	serverSchemaRoot = "/public/"
	serverSchemaFile = "schema.json"
)

// Config represents Agent configuration.
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
	// List of tasks for Agent to perform on events that involve specified resources.
	Tasks []*task `yaml:"tasks"`
	// InSecure https connection to endpoint
	InSecure bool `yaml:"insecure"`
	// Enabled
	Enabled bool `yaml:"enabled"`
	// Service user name for keystone
	ServiceUserID string `yaml:"service_user_id"`
	// Service user password for keystone
	ServiceUserPassword string `yaml:"service_user_password"`
}

// Agent represents Agent service.
type Agent struct {
	config    *Config
	backend   backend
	APIServer *client.HTTP
	serverAPI *schema.API
	// schemas map schema IDs to API Server schemas.
	schemas map[string]*schema.Schema
	log     *logrus.Entry
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

	serverSchema := filepath.Join(serverSchemaRoot, serverSchemaFile)
	if c.SchemaRoot != "" {
		serverSchema = filepath.Join(c.SchemaRoot, serverSchemaFile)
	}
	api, err := fetchServerAPI(context.Background(), s, serverSchema)
	if err != nil {
		return nil, err
	}

	b, err := newBackend(c.Backend)
	if err != nil {
		return nil, err
	}

	return &Agent{
		APIServer: s,
		config:    c,
		backend:   b,
		serverAPI: api,
		schemas:   buildSchemaMapping(api.Schemas),
		log:       logutil.NewLogger("agent"),
	}, nil
}

func fetchServerAPI(ctx context.Context, server *client.HTTP, serverSchema string) (*schema.API, error) {
	var api schema.API
	for {
		_, err := server.Read(ctx, serverSchema, &api)
		if err == nil {
			break
		}
		logrus.Warnf("failed to connect server %v. reconnecting...", err)
		time.Sleep(time.Second)
	}
	return &api, nil
}

func buildSchemaMapping(schemas []*schema.Schema) map[string]*schema.Schema {
	s := make(map[string]*schema.Schema)
	for _, schema := range schemas {
		// Compensate for empty Path and PluralPath fields in schema
		// TODO(daniel): remove this after following issue is fixed: https://github.com/Juniper/contrail/issues/72
		schema.Path = path.Join(schema.Prefix, strings.Replace(schema.ID, "_", "-", -1))
		schema.PluralPath = path.Join(schema.Prefix, strings.Replace(schema.Plural, "_", "-", -1))
		s[schema.ID] = schema
	}
	return s
}

// Watch starts watching for events on API Server resources.
func (a *Agent) Watch(ctx context.Context) error {
	a.log.Info("Starting watching for events")
	if err := a.APIServer.Login(ctx); err != nil {
		return errors.Wrap(err, "login to API Server failed")
	}

	var wg sync.WaitGroup
	wg.Add(len(a.config.Tasks))

	for _, task := range a.config.Tasks {
		task.init(a)

		for k := range a.schemas {
			for _, schemaID := range task.SchemaIDs {
				if schemaID != k {
					continue
				}

				watcher, err := newWatcher(a, task, k)
				if err != nil {
					return err
				}

				go func() {
					defer wg.Done()
					watcher.watch(ctx)
				}()
			}
		}
	}

	wg.Wait()
	return nil
}
