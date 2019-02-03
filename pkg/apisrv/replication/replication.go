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
	node        = "node"
	port        = "port"
	nodeProfile = "node-profile"
	endSystem   = "end-system"

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
	case *services.Event_CreateNodeProfileRequest:
		r.replicateToVNCAPI(createAction, nodeProfile,
			e.Request, &services.CreateNodeProfileResponse{})
	case *services.Event_UpdateNodeProfileRequest:
		r.replicateToVNCAPI(updateAction, nodeProfile,
			e.Request, &services.UpdateNodeProfileResponse{})
	case *services.Event_DeleteNodeProfileRequest:
		r.replicateToVNCAPI(deleteAction, nodeProfile,
			e.Request, &services.DeleteNodeProfileResponse{})
	case *services.Event_CreateNodeRequest:
		r.replicateToVNCAPI(createAction, node,
			e.Request, &services.CreateNodeResponse{})
	case *services.Event_UpdateNodeRequest:
		r.replicateToVNCAPI(updateAction, node,
			e.Request, &services.UpdateNodeResponse{})
	case *services.Event_DeleteNodeRequest:
		r.replicateToVNCAPI(deleteAction, node,
			e.Request, &services.DeleteNodeResponse{})
	case *services.Event_CreatePortRequest:
		r.replicateToVNCAPI(createAction, port,
			e.Request, &services.CreatePortResponse{})
	case *services.Event_UpdatePortRequest:
		r.replicateToVNCAPI(updateAction, port,
			e.Request, &services.UpdatePortResponse{})
	case *services.Event_DeletePortRequest:
		r.replicateToVNCAPI(deleteAction, port,
			e.Request, &services.DeletePortResponse{})
	}
}

func (r *Replication) replicateToVNCAPI(action string,
	url string, data services.isEvent_Request, output interface{}) {

	configEndpoint, err := r.endpointStore.GetEndpoint(configService)
	if err != nil {
		r.log.Errorf("while getting config endpoint url")
		return
	}

	vncUrl := configEndpoint + "/" + url

	switch action {
	case createAction:
		_, err = r.apiClient.Create(r.apiClientCtx, vncUrl, data, output)
		if err != nil {
			r.log.Errorf("while creating %s on vncAPI", vncUrl)
		}
	case updateAction:
		_, err = r.apiClient.Update(r.apiClientCtx, vncUrl, data, output)
		if err != nil {
			r.log.Errorf("while updating %s on vncAPI", vncUrl)
		}
	case deleteAction:
		_, err = r.apiClient.Delete(r.apiClientCtx, vncUrl, data, output)
		if err != nil {
			r.log.Errorf("while deleting %s on vncAPI", vncUrl)
		}
	}
	return
}

// Stop replication routine
func (r *Replication) Stop() {
	r.stopServiceContext()
	r.serviceWaitGroup.Wait()
}
