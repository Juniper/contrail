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
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/retry"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/sirupsen/logrus"
)

const (
	createTagURL            = "/tags"
	createVirtualNetworkURL = "/virtual-networks"
	createNetworkIpamURL    = "/network-ipams"
	createPortURL           = "/ports"
	createNodeProfileURL    = "/node-profiles"
	createEndSystemURL      = "/end-systems"
	updateTagURL            = "/tag"
	updateVirtualNetworkURL = "/virtual-network"
	updateNetworkIpamURL    = "/network-ipam"
	updatePortURL           = "/port/"
	updateNodeProfileURL    = "/node-profile/"
	updateEndSystemURL      = "/end-system/"

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
		}
		authEndpoint = keystoneTargets.Next(scope)
		if authEndpoint == nil {
			err = fmt.Errorf("unable to get keystone endpoint for: %s", endpointKey)
			return true, err
		}
		return false, nil
	}, retry.WithLog(logrus.StandardLogger()), retry.WithInterval(proxySyncInterval)); err != nil {
		r.log.Error(err)
		return nil, err
	}
	return authEndpoint, nil
}

func (r *Replicator) getAuthContext(clusterID string, apiClient *client.HTTP) context.Context {
	var err error
	var domainID, domainName string
	var projectID, projectName string
	ctx := auth.WithXClusterID(context.Background(), clusterID)
	if apiClient.Scope.Domain == nil &&
		apiClient.Scope.Project.Name == "" && apiClient.Scope.Project.ID == "" {
		apiClient.Init()
		projectID, err = apiClient.Keystone.GetProjectIDByName(
			ctx, apiClient.ID, apiClient.Password, defaultProjectName)
		if err != nil {
			domainID = defaultDomainID
			domainName = defaultDomainName
		} else {
			projectName = defaultProjectName
		}
	}
	apiClient.Scope = keystone.NewScope(domainID, domainName,
		projectID, projectName)
	// as auth is enabled, create ctx with auth
	varCtx := auth.NewContext(domainID, projectID,
		apiClient.ID, []string{projectName}, "")
	var authKey interface{} = "auth"
	ctx = context.WithValue(ctx, authKey, varCtx)
	return ctx
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
	// initialize client data
	apiClient := &client.HTTP{
		Endpoint: endpoint,
		InSecure: inSecure,
		Scope: keystone.NewScope(domainID, domainName,
			projectID, projectName),
	}
	// default: create no auth context
	ctx := auth.NoAuth(context.Background())
	if authURL != "" {
		apiClient.AuthURL = authURL
		apiClient.ID = id
		apiClient.Password = password
		ctx = r.getAuthContext(ep.ParentUUID, apiClient)
	}
	apiClient.Init()
	_, err := apiClient.Login(ctx)
	if err != nil {
		r.log.Warnf("Login failed for: %s, %v", ep.ParentUUID, err)
	}

	r.vncAPIClients[ep.ParentUUID] = &vncAPI{
		client:     apiClient,
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

func (r *Replicator) process(e *services.Event) { //nolint: gocyclo
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
		// handle tag
	case *services.Event_CreateTagRequest:
		r.replicateToVNCAPI(createAction, createTagURL,
			e.GetCreateTagRequest(), &services.CreateTagResponse{})
	case *services.Event_UpdateTagRequest:
		r.replicateToVNCAPI(updateAction, updateTagURL,
			e.GetUpdateTagRequest(), &services.UpdateTagResponse{})
	case *services.Event_DeleteTagRequest:
		r.replicateToVNCAPI(deleteAction, updateTagURL,
			e.GetDeleteTagRequest(), &services.DeleteTagResponse{})
	// handle virtual-network
	case *services.Event_CreateVirtualNetworkRequest:
		r.replicateToVNCAPI(createAction, createVirtualNetworkURL,
			e.GetCreateVirtualNetworkRequest(), &services.CreateVirtualNetworkResponse{})
	case *services.Event_UpdateVirtualNetworkRequest:
		r.replicateToVNCAPI(updateAction, updateVirtualNetworkURL,
			e.GetUpdateVirtualNetworkRequest(), &services.UpdateVirtualNetworkResponse{})
	case *services.Event_DeleteVirtualNetworkRequest:
		r.replicateToVNCAPI(deleteAction, updateVirtualNetworkURL,
			e.GetDeleteVirtualNetworkRequest(), &services.DeleteVirtualNetworkResponse{})
		// handle network-ipam
	case *services.Event_CreateNetworkIpamRequest:
		r.replicateToVNCAPI(createAction, createNetworkIpamURL,
			e.GetCreateNetworkIpamRequest(), &services.CreateNetworkIpamResponse{})
	case *services.Event_UpdateNetworkIpamRequest:
		r.replicateToVNCAPI(updateAction, updateNetworkIpamURL,
			e.GetUpdateNetworkIpamRequest(), &services.UpdateNetworkIpamResponse{})
	case *services.Event_DeleteNetworkIpamRequest:
		r.replicateToVNCAPI(deleteAction, updateNetworkIpamURL,
			e.GetDeleteNetworkIpamRequest(), &services.DeleteNetworkIpamResponse{})
	// handle node-profile
	case *services.Event_CreateNodeProfileRequest:
		r.replicateToVNCAPI(createAction, createNodeProfileURL,
			e.GetCreateNodeProfileRequest(), &services.CreateNodeProfileResponse{})
	case *services.Event_UpdateNodeProfileRequest:
		r.replicateToVNCAPI(updateAction, updateNodeProfileURL,
			e.GetUpdateNodeProfileRequest(), &services.UpdateNodeProfileResponse{})
	case *services.Event_DeleteNodeProfileRequest:
		r.replicateToVNCAPI(deleteAction, updateNodeProfileURL,
			e.GetDeleteNodeProfileRequest(), &services.DeleteNodeProfileResponse{})
	// handle nodes(end-systems)
	case *services.Event_CreateNodeRequest:
		r.replicateToVNCAPI(createAction, createEndSystemURL,
			e.GetCreateNodeRequest(), &services.CreateNodeResponse{})
	case *services.Event_UpdateNodeRequest:
		r.replicateToVNCAPI(updateAction, updateEndSystemURL,
			e.GetUpdateNodeRequest(), &services.UpdateNodeResponse{})
	case *services.Event_DeleteNodeRequest:
		r.replicateToVNCAPI(deleteAction, updateEndSystemURL,
			e.GetDeleteNodeRequest(), &services.DeleteNodeResponse{})
	// handle ports
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
