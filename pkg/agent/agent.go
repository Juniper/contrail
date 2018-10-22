package agent

import (
	"context"
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/config"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/logging"
	"github.com/Juniper/contrail/pkg/schema"
)

// Agent constants.
const (
	FileBackend    = "file"
	PollingWatcher = "polling"

	serverSchemaRoot = "/public/"
	serverSchemaFile = "schema.json"
)

// Config represents Agent configuration.
// nolint
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
	// InSecure https connection to endpoint
	InSecure bool `yaml:"insecure"`
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
	// Enabled
	Enabled bool `yaml:"enabled"`
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

	return NewAgent(&c)
}

// NewAgent creates Agent with given configuration.
func NewAgent(c *Config) (*Agent, error) {
	s := &client.HTTP{
		Endpoint: c.Endpoint,
		InSecure: c.InSecure,
	}
	// auth enabled
	if c.AuthURL != "" {
		s.AuthURL = c.AuthURL
		s.ID = c.ID
		s.Password = c.Password
		s.Scope = client.GetKeystoneScope(c.DomainID, c.DomainName,
			c.ProjectID, c.ProjectName)
	}
	s.Init()
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

	// create logger for agent
	logger := pkglog.NewLogger("agent")
	pkglog.SetLogLevel(logger, c.LogLevel)

	return &Agent{
		APIServer: s,
		config:    c,
		backend:   b,
		serverAPI: api,
		schemas:   buildSchemaMapping(api.Schemas),
		log:       logger,
	}, nil
}

func fetchServerAPI(ctx context.Context, server *client.HTTP, serverSchema string) (*schema.API, error) {
	var api schema.API
	for {
		_, err := server.Read(ctx, serverSchema, &api)
		if err == nil {
			break
		}
		log.Warn("failed to connect server %d. reconnecting...", err)
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
	// configure global log level
	logging.SetLogLevel()

	a.log.Info("Starting watching for events")
	if a.config.AuthURL != "" {
		err := a.APIServer.Login(ctx)
		if err != nil {
			return fmt.Errorf("login to API Server failed: %s", err)
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(a.config.Tasks))

	for _, task := range a.config.Tasks {
		task.init(a)
		schemaIDPattern := task.SchemaID

		for k := range a.schemas {
			matched, err := regexp.MatchString(schemaIDPattern, k)
			if err != nil {
				continue
			}
			if !matched {
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

	wg.Wait()
	return nil
}
