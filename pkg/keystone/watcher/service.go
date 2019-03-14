package watcher

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
)

const (
	poolingInterval = time.Duration(30) * time.Second
	httpTimeout     = time.Duration(15) * time.Second
)

type projectsSet map[string]struct{}

// Service is a servive type for keystone project watcher
type Service struct {
	apiServer     *client.HTTP
	knownProjects projectsSet
}

// NewKeystoneWatcherByConfig creates a service watcher that listen kieystone for project changes
func NewKeystoneWatcherByConfig() (*Service, error) {
	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}
	authURL := viper.GetString("keystone.authurl")
	if authURL == "" {
		return nil, errors.New("missing config option keystone.authurl needed by keystone watcher")
	}

	apiClient := getClient()
	return &Service{
		apiServer: apiClient,
	}, nil
}

// Watch starts listenting for project cheges in keystone
func (sv *Service) Watch() error {
	var err error
	if sv.apiServer.Keystone.URL != "" {
		if _, err = sv.apiServer.Login(context.Background()); err != nil {
			return err
		}
	}
	if err := sv.gatherInitialProjects(); err != nil {
		return err
	}
	for {
		time.Sleep(poolingInterval)
		token, err := sv.getKeystoneToken()
		if err != nil {
			log.Errorf("KeystoneProjectWatcher: Failed to get keystone token: %v", err)
			continue
		}
		projects, err := sv.apiServer.Keystone.GetProjects(context.Background(), token)
		if err != nil {
			log.Errorf("KeystoneProjectWatcher: Failed to get projects from keystone: %v", err)
			continue
		}
		sv.syncProjects(projects)
	}
}

func getClient() *client.HTTP {
	authURL := viper.GetString("keystone.authurl")
	// For getting all the projects we need a domain scope
	// unscoped reques won't work: https://bugs.launchpad.net/keystone/+bug/968696
	scope := &keystone.Scope{
		Domain: &keystone.Domain{
			ID: viper.GetString("client.domain_id"),
		},
	}
	client := client.NewHTTP(
		viper.GetString("client.endpoint"),
		authURL,
		viper.GetString("client.id"),
		viper.GetString("client.password"),
		viper.GetBool("insecure"),
		scope,
	)
	client.InitTimeout(httpTimeout)
	log.Infof("Making keystone watcher: %+v with Scope %+v (%+v; %+v)", client, client.Scope, client.Scope.Domain, client.Scope.Project)
	return client
}

func (sv *Service) gatherInitialProjects() error {
	// TODO For better startup performance projects should be gathered on startup
	return nil
}

func (sv *Service) getKeystoneToken() (string, error) {
	resp, err := sv.apiServer.Keystone.ObtainToken(context.Background(), sv.apiServer.ID, sv.apiServer.Password, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed getting keystone token")
	}
	return resp.Header.Get("X-Subject-Token"), nil
}

/* syncProjects compare list of given projects from keystone and create/delete missing/excess projects
Errors are logged only - it is not point to report those anywhere
It is desired to do all the work even if some sync operations fails */
func (sv *Service) syncProjects(projects []*keystone.Project) {
	currentProjects := sv.getKnownProjects()
	for _, prj := range projects {
		if _, have := currentProjects[prj.ID]; !have {
			if sv.addKeystoneProject(prj) {
				sv.knownProjects[prj.ID] = struct{}{}
			}
		}
		delete(currentProjects, prj.ID)
	}
	// All the ids in currentProjects set are those which keystone don't know - delete them
	for prj := range currentProjects {
		sv.deleteProject(prj)
	}
}

func (sv *Service) addKeystoneProject(prj *keystone.Project) bool {
	var output interface{}
	project := models.Project{
		ParentUUID: prj.Domain.ID,
		UUID:       prj.ID,
		Name:       prj.Name,
	}
	var data []byte
	data, err := json.Marshal(project)
	if err != nil {
		log.Errorf("Failed to marshal project %+v to json: %v", project, err)
		return false
	}
	_, err = sv.apiServer.Create(context.Background(), "/projects", data, &output)
	if err != nil {
		log.Errorf("KeystoneProjectWatcher: Failed to create project from keystone: %v", err)
		return false
	}

	return true
}

func (sv *Service) deleteProject(id string) {
	var output interface{}
	_, err := sv.apiServer.Delete(context.Background(), "/projects/"+id, &output)
	if err != nil {
		log.Errorf("Failed to delete project uuid=%v msg: %v", id, err)
		return
	}
}

func (sv *Service) getKnownProjects() projectsSet {
	current := projectsSet{}
	for p := range sv.knownProjects {
		current[p] = struct{}{}
	}
	return current
}
