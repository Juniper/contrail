package replication

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
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
	targetClient *client.HTTP
	sourceClient *client.HTTP
	ctx          context.Context
	clusterID    string
	endpointID   string
	log          *logrus.Entry
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
	h.APIServer = client.NewHTTP(&client.HTTPConfig{
		ID:       viper.GetString("client.id"),
		Password: viper.GetString("client.password"),
		Endpoint: viper.GetString("client.endpoint"),
		AuthURL:  viper.GetString("keystone.authurl"),
		Scope: keystone.NewScope(
			viper.GetString("client.domain_id"),
			viper.GetString("client.domain_name"),
			viper.GetString("client.project_id"),
			viper.GetString("client.project_name"),
		),
		Insecure: viper.GetBool("insecure"),
	})

	var err error
	if viper.GetString("keystone.authurl") != "" {
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
		apiClient.ID, []string{defaultProjectName}, "", auth.NewObjPerms(nil))
	var authKey interface{} = "auth"
	ctx = context.WithValue(ctx, authKey, varCtx)
	return ctx
}

func (h *vncAPIHandle) createClient(ep *models.Endpoint) {
	if ep.Prefix != configService {
		return
	}

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
			h.log.Warnf("VNC API client not prepared for %s, %v", ep.ParentUUID, err)
		}
		id = authEndpoint.Username
		password = authEndpoint.Password
	}

	c := client.NewHTTP(&client.HTTPConfig{
		ID:       id,
		Password: password,
		Endpoint: endpoint,
		AuthURL:  authURL,
		Scope: keystone.NewScope(
			domainID,
			domainName,
			projectID,
			projectName,
		),
		Insecure: inSecure,
	})

	ctx := auth.NoAuth(context.Background())
	if authURL != "" {
		ctx = h.getAuthContext(ep.ParentUUID, c)
	}

	_, err := c.Login(ctx)
	if err != nil {
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
	h.clients[ep.ParentUUID].targetClient.ID = ep.Username
	h.clients[ep.ParentUUID].targetClient.Password = ep.Password
	// Login to get fetch auth token
	vncAPI := h.clients[ep.ParentUUID]
	_, err := vncAPI.targetClient.Login(vncAPI.ctx)
	if err != nil {
		h.log.Warnf("Login failed for: %s, %v", ep.ParentUUID, err)
	}
	h.log.Debugf("Updated VNC API client for endpoint: %s", ep.UUID)
}

func (h *vncAPIHandle) deleteClient(endpointID string) {
	for clusterID, apiClient := range h.clients {
		if apiClient.endpointID == endpointID {
			delete(h.clients, clusterID)
			h.log.Debugf("Deleted VNC API client for endpoint: %s", endpointID)
			break
		}
	}
}

func (h *vncAPIHandle) replicate(action, sourceURL string, data interface{}, response interface{}) {
	if len(h.clients) == 0 {
		if err := h.initialize(); err != nil {
			h.log.Errorf("clients not initialized: %v", err)
		}
	}
	for _, vncAPI := range h.clients {
		vncAPI.replicate(action, sourceURL, data, response)
	}
}

func (v *vncAPI) replicate(action, sourceURL string, data interface{}, response interface{}) {
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
	refUpdateURL := strings.Join([]string{
		"/proxy", v.clusterID, configService, services.RefUpdatePath}, "/")
	for _, physicalInterface := range response.Port.PhysicalInterfaceBackRefs {
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
