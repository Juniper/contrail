package agent

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
	"sync"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/common"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const serverSchemaPath = "/public/schema.json"

// Config represents Agent configuration.
type Config struct {
	// ID of Agent account.
	ID string `yaml:"id"`
	// Password of Agent account.
	Password string `yaml:"password"`
	// ProjectID is ID of keystone project used for authentication.
	ProjectID string `yaml:"project_id"`
	// AuthURL defines authentication URL.
	AuthURL string `yaml:"auth_url"`
	// Endpoint of API Server.
	Endpoint string `yaml:"endpoint"`
	// Backend specifies backend to be used (values: "file").
	Backend string `yaml:"backend"`
	// Watcher specifies resource event watching strategy to be used (values: "polling").
	Watcher string `yaml:"watcher"`
	// List of tasks for Agent to perform on events that involve specified resources.
	Tasks []*task `yaml:"tasks"`
}

// Agent represents Agent service.
type Agent struct {
	config    *Config
	backend   backend
	APIServer *apisrv.Client
	serverAPI *common.API
	// schemas map schema IDs to API Server schemas.
	schemas map[string]*common.Schema
	log     *logrus.Entry
}

// NewAgentByFile creates Agent reading configuration from given file.
func NewAgentByFile(configPath string) (*Agent, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.UnmarshalStrict(data, &c)
	if err != nil {
		return nil, err
	}

	return NewAgent(&c)
}

// NewAgent creates Agent with given configuration.
func NewAgent(c *Config) (*Agent, error) {
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

	api, err := fetchServerAPI(s)
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
		log:       pkglog.NewLogger("agent"),
	}, nil
}

func fetchServerAPI(server *apisrv.Client) (*common.API, error) {
	var api common.API
	_, err := server.Read(serverSchemaPath, &api)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch API Server schemas")
	}
	return &api, nil
}

func buildSchemaMapping(schemas []*common.Schema) map[string]*common.Schema {
	s := make(map[string]*common.Schema)
	for _, schema := range schemas {
		// Compensate for empty Path and PluralPath fields in schema
		schema.Path = path.Join(schema.Prefix, strings.Replace(schema.ID, "_", "-", -1))
		schema.PluralPath = path.Join(schema.Prefix, strings.Replace(schema.Plural, "_", "-", -1))
		s[schema.ID] = schema
	}
	return s
}

// Watch starts watching for events on API Server resources.
func (a *Agent) Watch() error {
	a.log.Info("Starting watching for events")
	err := a.APIServer.Login()
	if err != nil {
		return fmt.Errorf("login to API Server failed: %s", err)
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
				watcher.watch()
			}()
		}
	}

	wg.Wait()
	return nil
}
