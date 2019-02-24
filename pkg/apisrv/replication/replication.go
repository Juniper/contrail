package replication

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	//"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/retry"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/sirupsen/logrus"
)

const (
	createPortURL        = "/ports"
	createNodeProfileURL = "/node-profiles"
	createEndSystemURL   = "/end-systems"
	updatePortURL        = "/port/"
	updateNodeProfileURL = "/node-profile/"
	updateEndSystemURL   = "/end-system/"

	deleteAction = "delete"
	createAction = "create"
	updateAction = "update"

	configService     = "config"
	keystoneService   = "keystone"
	scope             = "private"
	proxySyncInterval = 2 * time.Second
	basicAuth         = "basic-auth"

	defaultProjectName = "admin"
	defaultDomainID    = "default"
	defaultDomainName  = "Default"
)

type vncAPI struct {
	client     *client.HTTP
	ctx        context.Context
	endpointID string
}

// Replicator is an implementation to replicate objects to python API
type Replicator struct {
	serviceWaitGroup   *sync.WaitGroup
	serviceContext     context.Context
	stopServiceContext context.CancelFunc
	producer           *cache.DB
	endpointStore      *apicommon.EndpointStore
	vncAPIClients      map[string]*vncAPI
	log                *logrus.Entry
}

// New initializes replication data
func New(cacheDB *cache.DB,
	epStore *apicommon.EndpointStore) (*Replicator, error) {

	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}
	return &Replicator{
		endpointStore:    epStore,
		producer:         cacheDB,
		serviceWaitGroup: &sync.WaitGroup{},
		vncAPIClients:    make(map[string]*vncAPI),
		log:              logutil.NewLogger("vnc_replication"),
	}, nil
}
func (r *Replicator) readAuthEndpoint(clusterID string) (
	authEndpoint *apicommon.Endpoint, err error) {
	// retry 5 times at interval of 2 seconds
	// config endpoints are created before keystone
	// endpoints
	if err := retry.Do(func() (retry bool, err error) {
		endpointKey := strings.Join(
			[]string{"/proxy", clusterID, keystoneService, scope}, "/")
		keystoneTargets := r.endpointStore.Read(endpointKey)
		if keystoneTargets == nil {
			err = fmt.Errorf("keystone targets not found for: %s", endpointKey)
			return true, err
		} else {
			authEndpoint = keystoneTargets.Next(scope)
			if authEndpoint == nil {
				err = fmt.Errorf("unable to get keystone endpoint for: %s", endpointKey)
				return true, err
			}
			return false, nil
		}
	}, retry.WithLog(logrus.StandardLogger()), retry.WithInterval(proxySyncInterval)); err != nil {
		r.log.Error(err)
		return nil, err
	}
	return authEndpoint, nil
}

func (r *Replicator) createAPIClient(ep *models.Endpoint) {
	if ep.Prefix != configService {
		return
	}
	// get all config data
	var id, password string
	var domainID, domainName string
	var projectID, projectName string
	endpoint := viper.GetString("client.endpoint")
	inSecure := viper.GetBool("insecure")
	authURL := viper.GetString("keystone.authurl")
	if viper.GetString("auth_type") == basicAuth {
		id = viper.GetString("client.id")
		password = viper.GetString("client.password")
		domainID = viper.GetString("client.domain_id")
		projectID = viper.GetString("client.project_id")
		domainName = viper.GetString("client.domain_name")
		projectName = viper.GetString("client.project_name")
	} else {
		// get keystone endpoint
		authEndpoint, err := r.readAuthEndpoint(ep.ParentUUID)
		if err != nil {
			r.log.Warnf("vncAPI client not prepared for %s, %v", ep.ParentUUID, err)
		}
		id = authEndpoint.Username
		password = authEndpoint.Password
	}
	// intialize client data
	s := &client.HTTP{
		Endpoint: endpoint,
		InSecure: inSecure,
	}
	// default: create no auth context
	ctx := auth.NoAuth(context.Background())
	if authURL != "" {
		ctx = auth.WithXClusterID(context.Background(), ep.ParentUUID)
		s.AuthURL = authURL
		s.ID = id
		s.Password = password
		if projectID == "" && projectName == "" &&
			domainID == "" && domainName == "" {
			s.Init()
			projectID, err := s.Keystone.GetProjectIDByName(ctx, id, password, defaultProjectName)
			if err != nil {
				domainID = defaultDomainID
				domainName = defaultDomainName
			} else {
				projectName = defaultProjectName
				projectID = projectID
			}
		}
		s.Scope = keystone.GetScope(domainID, domainName,
			projectID, projectName)
		// as auth is enabled, create ctx with auth
		varCtx := auth.NewContext(domainID, projectID,
			id, []string{projectName}, "")
		var authKey interface{} = "auth"
		ctx = context.WithValue(ctx, authKey, varCtx)
	}
	s.Init()
	_, err := s.Login(ctx)
	if err != nil {
		r.log.Warnf("Login failed for: %s, %v", ep.ParentUUID, err)
	}

	r.vncAPIClients[ep.ParentUUID] = &vncAPI{
		client:     s,
		ctx:        ctx,
		endpointID: ep.UUID,
	}
}

func (r *Replicator) updateAPIClient(ep *models.Endpoint) {
	if ep.Prefix != keystoneService {
		// no need to update the auth credential in the client
		return
	}
	r.vncAPIClients[ep.ParentUUID].client.ID = ep.Username
	r.vncAPIClients[ep.ParentUUID].client.Password = ep.Password
	// Login to get fetch auth token
	vncAPI := r.vncAPIClients[ep.ParentUUID]
	vncAPI.client.Init()
	_, err := vncAPI.client.Login(vncAPI.ctx)
	if err != nil {
		r.log.Warnf("Login failed for: %s, %v", ep.ParentUUID, err)
	}

}

func (r *Replicator) deleteAPIClient(endpointID string) {
	for clusterID, apiClient := range r.vncAPIClients {
		if apiClient.endpointID == endpointID {
			delete(r.vncAPIClients, clusterID)
			break
		}
	}
}

// Start starts replication service
func (r *Replicator) Start() error {

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
				r.log.Info("Stopping VNC API replication service")
				return
			case e := <-watcher.Chan():
				r.process(e)
			}
		}
	}()
	return nil
}

func (r *Replicator) process(e *services.Event) {
	switch e.Request.(type) {
	// watch endpoint event and prepare clients
	case *services.Event_CreateEndpointRequest:
		ep := e.GetCreateEndpointRequest().Endpoint
		r.createAPIClient(ep)
	case *services.Event_UpdateEndpointRequest:
		ep := e.GetUpdateEndpointRequest().Endpoint
		r.updateAPIClient(ep)
	case *services.Event_DeleteEndpointRequest:
		id := e.GetDeleteEndpointRequest().ID
		r.deleteAPIClient(id)
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
func (r *Replicator) replicateToVNCAPI(action string,
	url string, data interface{}, output interface{}) {

	for clusterID, vncAPI := range r.vncAPIClients {
		proxyURL := strings.Join([]string{"/proxy", clusterID, configService, url}, "/")
		r.log.Debugf("replicating %v to cluster: %s", data, clusterID)
		switch action {
		case createAction:
			_, err := vncAPI.client.Create(vncAPI.ctx, proxyURL, data, output)
			if err != nil {
				r.log.Errorf("while creating %s on vncAPI: %v", proxyURL, err)
			}
		case updateAction:
			_, err := vncAPI.client.Update(vncAPI.ctx, proxyURL, data, output)
			if err != nil {
				r.log.Errorf("while updating %s on vncAPI: %v", proxyURL, err)
			}
		case deleteAction:
			_, err := vncAPI.client.Delete(vncAPI.ctx, proxyURL, output)
			if err != nil {
				r.log.Errorf("while deleting %s on vncAPI: %v", proxyURL, err)
			}
		}
	}
	return
}

// Stop replication routine
func (r *Replicator) Stop() {
	r.stopServiceContext()
	r.serviceWaitGroup.Wait()
}
