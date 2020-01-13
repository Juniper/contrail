package replication

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/retry"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/endpoint"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	asfauth "github.com/Juniper/asf/pkg/auth"
	asfclient "github.com/Juniper/asf/pkg/client"
	kstypes "github.com/Juniper/asf/pkg/keystone"
)

const (
	configService     = "config"
	keystoneService   = "keystone"
	scope             = "private"
	proxySyncInterval = 2 * time.Second
	basicAuth         = "basic-auth"

	defaultProjectName = "admin"
)

type vncAPIHandle struct {
	APIServer     *client.HTTP
	clients       map[string]*vncAPI
	endpointStore *endpoint.Store
	log           *logrus.Entry
	auth          *keystone.Keystone
}

func newVncAPIHandle(epStore *endpoint.Store, auth *keystone.Keystone) *vncAPIHandle {
	return &vncAPIHandle{
		clients:       make(map[string]*vncAPI),
		endpointStore: epStore,
		log:           logutil.NewLogger("vnc_replication_client"),
		auth:          auth,
		APIServer:     client.NewHTTPFromConfig(),
	}
}

// CreateClient creates client for given endpoint.
func (h *vncAPIHandle) CreateClient(ep *models.Endpoint) {
	if ep.Prefix != configService {
		return
	}

	config := asfclient.LoadHTTPConfig()
	authType, err := h.auth.GetAuthType(ep.ParentUUID)
	if err != nil {
		h.log.Errorf("Not able to find auth type for cluster %s, %v", ep.ParentUUID, err)
	}
	if authType != basicAuth {
		config.Scope = &kstypes.Scope{Project: &kstypes.Project{Domain: kstypes.DefaultDomain()}}
		// get keystone endpoint
		var e *endpoint.Endpoint
		e, err = h.readAuthEndpoint(ep.ParentUUID)
		if err != nil {
			h.log.Warnf("VNC API client not prepared for %s, %v", ep.ParentUUID, err)
		}
		config.SetCredentials(e.Username, e.Password)
	}

	c := client.NewHTTP(config)

	ctx := asfauth.NoAuth(context.Background())
	if viper.GetString("keystone.authurl") != "" {
		ctx = h.getAuthContext(ep.ParentUUID, c)
	}

	if err := c.Login(ctx); err != nil {
		h.log.Warnf("Login failed for: %s, %v", ep.ParentUUID, err)
	}

	h.clients[ep.ParentUUID] = &vncAPI{
		targetClient: c,
		sourceClient: h.APIServer,
		ctx:          ctx,
		clusterID:    ep.ParentUUID,
		endpointID:   ep.UUID,
		log:          h.log,
	}
	h.log.Debugf("Created VNC API client for endpoint: %s", ep.UUID)
}

func (h *vncAPIHandle) readAuthEndpoint(clusterID string) (authEndpoint *endpoint.Endpoint, err error) {
	// retry 5 times at interval of 2 seconds
	// config endpoints are created before keystone
	// endpoints
	if err := retry.Do(func() (retry bool, err error) {
		// TODO(dfurman): "server.dynamic_proxy_path" or DefaultDynamicProxyPath should be used
		endpointKey := strings.Join(
			[]string{"/proxy", clusterID, keystoneService, scope}, "/")
		keystoneTargets := h.endpointStore.Get(endpointKey)
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
		projectID, err = h.getProjectIDByName(ctx, apiClient)
		if err == nil {
			apiClient.Scope = kstypes.NewScope(
				kstypes.DefaultDomainID, kstypes.DefaultDomainName,
				projectID, defaultProjectName)
		}
	}
	// as auth is enabled, create ctx with auth
	varCtx := kstypes.NewAuthIdentity(kstypes.DefaultDomainID, projectID, apiClient.ID, []string{defaultProjectName})
	return asfauth.WithIdentity(ctx, varCtx)
}

func (h *vncAPIHandle) getProjectIDByName(ctx context.Context, apiClient *client.HTTP) (string, error) {
	token, err := apiClient.Keystone.ObtainUnscopedToken(
		ctx, apiClient.ID, apiClient.Password, apiClient.Scope.Project.Domain,
	)
	if err != nil {
		return "", err
	}
	ctx = kstypes.WithXAuthToken(ctx, token)
	return apiClient.Keystone.GetProjectIDByName(
		ctx, defaultProjectName, apiClient.Scope.Project.Domain,
	)
}

// UpdateClient updates client for given endpoint.
func (h *vncAPIHandle) UpdateClient(ep *models.Endpoint) {
	if ep.Prefix == configService {
		if _, ok := h.clients[ep.ParentUUID]; !ok {
			h.CreateClient(ep)
		}
	}
	if ep.Prefix != keystoneService {
		// no need to update the auth credential in the client
		return
	}
	h.clients[ep.ParentUUID].targetClient.ID = ep.Username
	h.clients[ep.ParentUUID].targetClient.Password = ep.Password
	// Login to get fetch auth token
	vncAPI := h.clients[ep.ParentUUID]
	if err := vncAPI.targetClient.Login(vncAPI.ctx); err != nil {
		h.log.Warnf("Login failed for: %s, %v", ep.ParentUUID, err)
	}
	h.log.Debugf("Updated VNC API client for endpoint: %s", ep.UUID)
}

// DeleteClient deletes client for given endpoint.
func (h *vncAPIHandle) DeleteClient(endpointID string) {
	for clusterID, apiClient := range h.clients {
		if apiClient.endpointID == endpointID {
			delete(h.clients, clusterID)
			h.log.Debugf("Deleted VNC API client for endpoint: %s", endpointID)
			break
		}
	}
}

// Replicate propagates given event to VNC API.
func (h *vncAPIHandle) Replicate(action, sourceURL string, data interface{}, response interface{}) {
	if len(h.clients) == 0 {
		if err := h.initClients(); err != nil {
			h.log.Errorf("Clients not initialized: %v", err)
		}
	}
	for _, vncAPI := range h.clients {
		vncAPI.Replicate(action, sourceURL, data, response)
	}
}

func (h *vncAPIHandle) initClients() error {
	endpoints, err := h.listConfigEndpoints()
	if err != nil {
		return err
	}
	for _, e := range endpoints {
		h.CreateClient(e)
	}
	return nil
}

func (h *vncAPIHandle) listConfigEndpoints() ([]*models.Endpoint, error) {
	request := &services.ListEndpointRequest{
		Spec: &baseservices.ListSpec{
			Fields:  []string{"uuid", "parent_uuid", "prefix"},
			Filters: []*baseservices.Filter{{Key: "prefix", Values: []string{configService}}},
		},
	}
	resp, err := h.APIServer.ListEndpoint(context.Background(), request)
	return resp.GetEndpoints(), err
}

type vncAPI struct {
	targetClient *client.HTTP
	sourceClient *client.HTTP
	ctx          context.Context
	clusterID    string
	endpointID   string
	log          *logrus.Entry
}

// Replicate propagates given event to VNC API.
func (v *vncAPI) Replicate(action, sourceURL string, data interface{}, response interface{}) {
	// TODO(dfurman): "server.dynamic_proxy_path" or DefaultDynamicProxyPath should be used
	targetURL := strings.Join([]string{"/proxy", v.clusterID, configService, sourceURL}, "/")
	v.log.WithFields(logrus.Fields{
		"data":      data,
		"clusterID": v.clusterID,
	}).Debug("Replicating to cluster")
	switch action {
	case createAction:
		v.replicateCreate(targetURL, data, response)
	case updateAction:
		v.replicateUpdate(sourceURL, targetURL, data, response)
	case deleteAction:
		v.replicateDelete(sourceURL, targetURL, data, response)
	case refUpdateAction:
		v.replicateRefUpdate(targetURL, data, response)
	}
}

func (v *vncAPI) replicateCreate(targetURL string, data, response interface{}) {
	_, err := v.targetClient.Create(v.ctx, targetURL, data, response)
	if err != nil {
		v.log.WithError(err).WithField(
			"targetURL", targetURL,
		).Error("Failed to create resource in VNC API")
	}
}

func (v *vncAPI) replicateUpdate(sourceURL, targetURL string, data, response interface{}) {
	resp, err := v.targetClient.Update(v.ctx, targetURL, data, response)
	if err != nil {
		if isNotFound(resp) {
			v.log.WithError(err).WithField(
				"targetURL", targetURL,
			).Debug("Failed to update resource in VNC API; fetching it from API to create it instead")
			v.fetchAndCreate(sourceURL, targetURL)
		} else {
			v.log.WithError(err).WithField(
				"targetURL", targetURL,
			).Error("Failed to update resource in VNC API")
		}
	}
}

func (v *vncAPI) fetchAndCreate(sourceURL, targetURL string) {
	sourceURL = "/" + sourceURL
	var resource map[string]interface{}
	_, err := v.sourceClient.Read(v.ctx, sourceURL, &resource)
	if err != nil {
		v.log.WithError(err).WithField("sourceURL", sourceURL).
			Error("Failed to fetch from replication source")
		return
	}

	createURL := singularToPluralURL(targetURL)
	var response map[string]interface{}
	_, err = v.targetClient.Create(v.ctx, createURL, resource, response)
	if err != nil {
		v.log.WithError(err).WithField(
			"targetURL", createURL,
		).Error("Failed to create resource in VNC API to work around updating")
	}
}

func singularToPluralURL(singularURL string) string {
	parts := strings.Split(singularURL, "/")
	if len(parts) == 0 {
		return ""
	}
	singularURLWithoutUUID := strings.Join(parts[:len(parts)-1], "/")
	return singularURLWithoutUUID + "s"
}

func (v *vncAPI) replicateDelete(sourceURL, targetURL string, data, response interface{}) {
	urlParts := strings.Split(sourceURL, "/")
	if urlParts[0] == "port" {
		v.deletePhysicaInterfaceToPortRefs(targetURL, urlParts[1])
	}
	resp, err := v.targetClient.Delete(v.ctx, targetURL, response)
	if err != nil {
		if isNotFound(resp) {
			v.log.WithError(err).WithField(
				"targetURL", targetURL,
			).Debug("Failed to delete resource in VNC API: not found")
		} else {
			v.log.WithError(err).WithField(
				"targetURL", targetURL,
			).Error("Failed to delete resource in VNC API")
		}
	}
}

func (v *vncAPI) deletePhysicaInterfaceToPortRefs(portURL, portID string) {
	// Read physical-interface back_refs from vnc_api
	response := &services.GetPortResponse{}
	_, err := v.targetClient.Read(v.ctx, portURL, response)
	if err != nil {
		v.log.WithError(err).WithField(
			"targetURL", portURL,
		).Error("Failed to read port from VNC API")
	}
	// Delete physical-interface to this port ref
	// TODO(dfurman): "server.dynamic_proxy_path" or DefaultDynamicProxyPath should be used
	refUpdateURL := strings.Join([]string{
		"/proxy", v.clusterID, configService, services.RefUpdatePath}, "/")
	for _, physicalInterface := range response.GetPort().GetPhysicalInterfaceBackRefs() {
		data := services.RefUpdate{
			Operation: "DELETE",
			Type:      "physical-interface",
			UUID:      physicalInterface.UUID,
			RefType:   "port",
			RefUUID:   portID,
		}
		expected := []int{http.StatusOK}
		_, err := v.targetClient.Do(
			v.ctx, echo.POST, refUpdateURL, nil, data, map[string]interface{}{}, expected,
		)
		if err != nil {
			v.log.WithError(err).WithField(
				"data", data,
			).Error("Failed to update ref in VNC API")
		}
	}
}

func isNotFound(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusNotFound
}

func (v *vncAPI) replicateRefUpdate(targetURL string, data, response interface{}) {
	expected := []int{http.StatusOK}
	_, err := v.targetClient.Do(v.ctx, echo.POST, targetURL, nil, data, response, expected)
	if err != nil {
		v.log.WithError(err).WithField(
			"data", data,
		).Error("Failed to update ref in VNC API")
	}
}
