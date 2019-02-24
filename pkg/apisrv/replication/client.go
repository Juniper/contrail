package replication

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/retry"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

const (
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
	clusterID  string
	endpointID string
	log        *logrus.Entry
}

type vncAPIHandle struct {
	APIServer     *client.HTTP
	clients       map[string]*vncAPI
	endpointStore *apicommon.EndpointStore
	log           *logrus.Entry
}

func newVncAPIHandle(epStore *apicommon.EndpointStore) *vncAPIHandle {
	handle := &vncAPIHandle{
		clients:       make(map[string]*vncAPI),
		endpointStore: epStore,
		log:           logutil.NewLogger("vnc_replication_client"),
	}
	return handle
}

func (h *vncAPIHandle) initialize() (err error) {
	if err = h.initClient(); err != nil {
		return err
	}
	var endpoints []*models.Endpoint
	if endpoints, err = h.ListConfigEndpoints(); err != nil {
		return err
	}
	for _, e := range endpoints {
		h.createClient(e)
	}
	return nil
}

func (h *vncAPIHandle) initClient() error {
	authURL := viper.GetString("keystone.authurl")
	scope := keystone.NewScope(
		viper.GetString("client.domain_id"),
		viper.GetString("client.domain_name"),
		viper.GetString("client.project_id"),
		viper.GetString("client.project_name"),
	)
	h.APIServer = client.NewHTTP(
		viper.GetString("client.endpoint"),
		authURL,
		viper.GetString("client.id"),
		viper.GetString("client.password"),
		viper.GetBool("insecure"),
		scope,
	)
	var err error
	if authURL != "" {
		_, err = h.APIServer.Login(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *vncAPIHandle) ListConfigEndpoints() (endpoints []*models.Endpoint, err error) {

	request := &services.ListEndpointRequest{
		Spec: &baseservices.ListSpec{
			Fields: []string{"uuid", "parent_uuid", "prefix"},
			Filters: []*baseservices.Filter{
				{
					Key:    "prefix",
					Values: []string{configService},
				},
			},
		},
	}
	var resp *services.ListEndpointResponse
	resp, err = h.APIServer.ListEndpoint(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp.GetEndpoints(), nil
}

func (h *vncAPIHandle) readAuthEndpoint(clusterID string) (authEndpoint *apicommon.Endpoint, err error) {
	// retry 5 times at interval of 2 seconds
	// config endpoints are created before keystone
	// endpoints
	if err := retry.Do(func() (retry bool, err error) {
		endpointKey := strings.Join(
			[]string{"/proxy", clusterID, keystoneService, scope}, "/")
		keystoneTargets := h.endpointStore.Read(endpointKey)
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
	}, retry.WithLog(logrus.StandardLogger()),
		retry.WithInterval(proxySyncInterval)); err != nil {
		h.log.Error(err)
		return nil, err
	}
	return authEndpoint, nil
}

func (h *vncAPIHandle) getAuthContext(clusterID string, apiClient *client.HTTP) context.Context {
	var err error
	var projectID string
	ctx := auth.WithXClusterID(context.Background(), clusterID)
	if apiClient.Scope.Project.Name == "" && apiClient.Scope.Project.ID == "" {
		apiClient.Init()
		projectID, err = apiClient.Keystone.GetProjectIDByName(
			ctx, apiClient.ID, apiClient.Password, defaultProjectName,
			apiClient.Scope.Project.Domain)
		if err == nil {
			apiClient.Scope = keystone.NewScope(
				defaultDomainID, defaultDomainName,
				projectID, defaultProjectName)
		}
	}
	// as auth is enabled, create ctx with auth
	varCtx := auth.NewContext(defaultDomainID, projectID,
		apiClient.ID, []string{defaultProjectName}, "")
	var authKey interface{} = "auth"
	ctx = context.WithValue(ctx, authKey, varCtx)
	return ctx
}

func (h *vncAPIHandle) createClient(ep *models.Endpoint) {
	if ep.Prefix != configService {
		return
	}
	// get all config data
	var id, password string
	var projectID, projectName string
	var domainID, domainName string
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
		domainID = defaultDomainID
		domainName = defaultDomainName
		// get keystone endpoint
		authEndpoint, err := h.readAuthEndpoint(ep.ParentUUID)
		if err != nil {
			h.log.Warnf("vncAPI client not prepared for %s, %v", ep.ParentUUID, err)
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
		ctx = h.getAuthContext(ep.ParentUUID, apiClient)
	}
	apiClient.Init()
	_, err := apiClient.Login(ctx)
	if err != nil {
		h.log.Warnf("Login failed for: %s, %v", ep.ParentUUID, err)
	}

	h.clients[ep.ParentUUID] = &vncAPI{
		client:     apiClient,
		ctx:        ctx,
		clusterID:  ep.ParentUUID,
		endpointID: ep.UUID,
		log:        h.log,
	}
	h.log.Debugf("created vnc client for endpoint: %s", ep.UUID)
}

func (h *vncAPIHandle) updateClient(ep *models.Endpoint) {
	if ep.Prefix == configService {
		if _, ok := h.clients[ep.ParentUUID]; !ok {
			h.createClient(ep)
		}
	}
	if ep.Prefix != keystoneService {
		// no need to update the auth credential in the client
		return
	}
	h.clients[ep.ParentUUID].client.ID = ep.Username
	h.clients[ep.ParentUUID].client.Password = ep.Password
	// Login to get fetch auth token
	vncAPI := h.clients[ep.ParentUUID]
	vncAPI.client.Init()
	_, err := vncAPI.client.Login(vncAPI.ctx)
	if err != nil {
		h.log.Warnf("Login failed for: %s, %v", ep.ParentUUID, err)
	}
	h.log.Debugf("updated vnc client for endpoint: %s", ep.UUID)
}

func (h *vncAPIHandle) deleteClient(endpointID string) {
	for clusterID, apiClient := range h.clients {
		if apiClient.endpointID == endpointID {
			delete(h.clients, clusterID)
			h.log.Debugf("deleted vnc client for endpoint: %s", endpointID)
			break
		}
	}
}

func (h *vncAPIHandle) replicate(action, url string, data interface{}, response interface{}) {
	if len(h.clients) == 0 {
		if err := h.initialize(); err != nil {
			h.log.Errorf("clients not initialized: %v", err)
		}
	}
	for _, vncAPI := range h.clients {
		vncAPI.replicate(action, url, data, response)
	}
}

func (v *vncAPI) replicate(action, url string, data interface{}, response interface{}) {
	proxyURL := strings.Join([]string{"/proxy", v.clusterID, configService, url}, "/")
	v.log.Debugf("replicating %v to cluster: %s", data, v.clusterID)
	switch action {
	case createAction:
		_, err := v.client.Create(v.ctx, proxyURL, data, response)
		if err != nil {
			v.log.Errorf("while creating %s on vncAPI: %v", proxyURL, err)
		}
	case updateAction:
		_, err := v.client.Update(v.ctx, proxyURL, data, response)
		if err != nil {
			v.log.Errorf("while updating %s on vncAPI: %v", proxyURL, err)
		}
	case deleteAction:
		_, err := v.client.Delete(v.ctx, proxyURL, response)
		if err != nil {
			v.log.Errorf("while deleting %s on vncAPI: %v", proxyURL, err)
		}
	}
}
