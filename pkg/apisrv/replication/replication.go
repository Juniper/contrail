package replication

import (
	"context"
	"sync"

	"github.com/siddontang/go/log"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/db/cache"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/sirupsen/logrus"
)

const (
	createPortURL        	= "/ports/"
	createNodeProfileURL 	= "/node-profiles/"
	createEndSystemURL   	= "/end-systems/"
	updatePortURL				 	= "/port/"
	updateNodeProfileURL	= "/node-profile/"
	updateEndSystemURL		= "/end-system/"

	deleteAction = "delete"
	createAction = "create"
	updateAction = "update"

	configService = "config"
)

// Replication is an implementation to replicate objects to python API
type Replication struct {
	serviceWaitGroup   *sync.WaitGroup
	serviceContext     context.Context
	stopServiceContext context.CancelFunc
	producer           *cache.DB
	endpointStore      *apicommon.EndpointStore
	apiClient          *client.HTTP
	log                *logrus.Entry
	apiClientCtx       context.Context
}

// New initializes replication data
func New(cacheDB *cache.DB,
	epStore *apicommon.EndpointStore) *Replication {

	apiClient, ctx := newAPIClient()

	// create logger for replicate
	logger := pkglog.NewLogger("replicate")
	pkglog.SetLogLevel(logger, viper.GetString("log_level"))

	return &Replication{
		endpointStore:    epStore,
		producer:         cacheDB,
		serviceWaitGroup: &sync.WaitGroup{},
		apiClient:        apiClient,
		apiClientCtx:     ctx,
		log:              logger,
	}
}

func newAPIClient() (*client.HTTP, context.Context) {

	// get all config data
	id := viper.GetString("client.id")
	password := viper.GetString("client.password")
	domainID := viper.GetString("client.domain_id")
	projectID := viper.GetString("client.project_id")
	domainName := viper.GetString("client.domain_name")
	projectName := viper.GetString("client.project_name")
	authURL := viper.GetString("keystone.authurl")
	inSecure := viper.GetBool("insecure")
	endpoint := viper.GetString("client.endpoint")

	// intialize client data
	s := &client.HTTP{
		Endpoint: endpoint,
		InSecure: inSecure,
	}

	// default: create no auth context
	ctx := auth.NoAuth(context.Background())
	if authURL != "" {
		s.AuthURL = authURL
		s.ID = id
		s.Password = password
		s.Scope = client.GetKeystoneScope(domainID, domainName,
			projectID, projectName)

		// as auth is enabled, create ctx with auth
		varCtx := auth.NewContext(domainID, projectID,
			id, []string{projectName})
		var authKey interface{} = "auth"
		ctx = context.WithValue(context.Background(), authKey, varCtx)
	}
	s.Init()

	return s, ctx
}

// Start starts replication service
func (r *Replication) Start() error {

	r.serviceContext, r.stopServiceContext = context.WithCancel(context.Background())
	watcher, err := r.producer.AddWatcher(r.serviceContext, 0)
	if err != nil {
		return err
	}

	r.serviceWaitGroup.Add(1)
	go func() {
		defer r.serviceWaitGroup.Done()
		for {
			select {
			case <-r.serviceContext.Done():
				log.Info("Stopping VNC API replication service")
				return
			case e := <-watcher.Chan():
				r.process(e)
			}
		}
	}()

	return nil
}

func (r *Replication) process(e *services.Event) {
	// do certain action on received event

	switch e.Request.(type) {
	// handle request related to node-profile
	case *services.Event_CreateNodeProfileRequest:
		r.replicateToVNCAPI(createAction, createNodeProfileURL,
			e.GetCreateNodeProfileRequest(), &services.CreateNodeProfileResponse{})
	case *services.Event_UpdateNodeProfileRequest:
		r.replicateToVNCAPI(updateAction, updateNodeProfileURL,
			e.GetUpdateNodeProfileRequest(), &services.UpdateNodeProfileResponse{})
	case *services.Event_DeleteNodeProfileRequest:
		r.replicateToVNCAPI(deleteAction, updateNodeProfileURL,
			e.GetDeleteNodeProfileRequest(), &services.DeleteNodeProfileResponse{})
	// handle request related to nodes/end-systems
	case *services.Event_CreateNodeRequest:
		r.replicateToVNCAPI(createAction, createEndSystemURL,
			e.GetCreateNodeRequest(), &services.CreateNodeResponse{})
	case *services.Event_UpdateNodeRequest:
		r.replicateToVNCAPI(updateAction, updateEndSystemURL,
			e.GetUpdateNodeRequest(), &services.UpdateNodeResponse{})
	case *services.Event_DeleteNodeRequest:
		r.replicateToVNCAPI(deleteAction, updateEndSystemURL,
			e.GetDeleteNodeRequest(), &services.DeleteNodeResponse{})
	// handle request related to ports
	case *services.Event_CreatePortRequest:
		r.replicateToVNCAPI(createAction, createPortURL,
			e.GetCreatePortRequest(), &services.CreatePortResponse{})
	case *services.Event_UpdatePortRequest:
		r.replicateToVNCAPI(updateAction, updatePortURL,
			e.GetUpdatePortRequest(), &services.UpdatePortResponse{})
	case *services.Event_DeletePortRequest:
		r.replicateToVNCAPI(deleteAction, updatePortURL,
			e.GetDeletePortRequest(), &services.DeletePortResponse{})
	}
}

//nolint: vetshadow,vet
func (r *Replication) replicateToVNCAPI(action string,
	url string, data interface{}, output interface{}) {

	configEndpoint, err := r.endpointStore.GetEndpoint(configService)
	if err != nil {
		r.log.Errorf("while getting config endpoint url")
		return
	}

	vncURL := configEndpoint + url

	switch action {
	case createAction:

		_, err = r.apiClient.Create(r.apiClientCtx, vncURL, data, output)
		if err != nil {
			r.log.Errorf("while creating %s on vncAPI", vncURL)
		}
	case updateAction:
		_, err = r.apiClient.Update(r.apiClientCtx, vncURL, data, output)
		if err != nil {
			r.log.Errorf("while updating %s on vncAPI", vncURL)
		}
	case deleteAction:
		_, err = r.apiClient.Delete(r.apiClientCtx, vncURL, output)
		if err != nil {
			r.log.Errorf("while deleting %s on vncAPI", vncURL)
		}
	}
	return
}

// Stop replication routine
func (r *Replication) Stop() {
	r.stopServiceContext()
	r.serviceWaitGroup.Wait()
}
