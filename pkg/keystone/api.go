package keystone

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/proxy"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/config"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Keystone constants.
const (
	LocalAuthPath = "/keystone/v3"

	authAPIVersion   = "v3"
	basicAuth        = "basic-auth"
	configService    = "config"
	keystoneAuth     = "keystone"
	keystoneService  = "keystone"
	openstack        = "openstack"
	xAuthTokenKey    = "X-Auth-Token"
	xSubjectTokenKey = "X-Subject-Token"
	xClusterIDKey    = "X-Cluster-ID"
)

const (
	PublicURLScope  = "public"
	PrivateURLScope = "private"
)

type EndpointStore interface {
	Range(f func(scope, value interface{}) bool)
	InStore(scope string) bool
	ReadEndpointUrls(scope string, endpointKey string) []string
	Remove(scope string)
	UpdateEndpoint(scope string, endpoint *models.Endpoint) error
	InitScope(scope string)
	RemoveDeleted(scope, value interface{}, endpoints map[string]*models.Endpoint, log *logrus.Entry) bool
	GetEndpoints(scope string, endpointKey string) bool
	ReadAuthURLs(scope string, endpointKey string, authApiVersion string) []string
	ReadUsername(scope string, endpointKey string) string
	ReadPassword(scope string, endpointKey string) string
	GetEndpointUrl(clusterID, prefix string) (string, bool)
}

//Keystone is used to represents Keystone Controller.
type Keystone struct {
	store            Store
	Assignment       Assignment
	staticAssignment *StaticAssignment
	endpointStore    EndpointStore
	apiClient        *client.HTTP

	log *logrus.Entry
}

//Init is used to initialize echo with Keystone capability.
//This function reads config from viper.
func (keystone *Keystone) Init(e *echo.Echo, es EndpointStore) error {

	keystone.endpointStore = es
	keystone.apiClient = client.NewHTTPFromConfig()
	keystone.log = logutil.NewLogger("keystone-api")
	assignmentType := viper.GetString("keystone.assignment.type")
	if assignmentType == "static" {
		var staticAssignment StaticAssignment
		err := config.LoadConfig("keystone.assignment.data", &staticAssignment)
		if err != nil {
			return err
		}
		keystone.staticAssignment = &staticAssignment
		keystone.Assignment = &staticAssignment
	}
	storeType := viper.GetString("keystone.store.type")
	if storeType == "memory" {
		expire := viper.GetInt64("keystone.store.expire")
		keystone.store = MakeInMemoryStore(time.Duration(expire) * time.Second)
	}
	e.POST("/keystone/v3/auth/tokens", keystone.CreateTokenAPI)
	e.GET("/keystone/v3/auth/tokens", keystone.ValidateTokenAPI)

	// TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
	e.GET("/keystone/v3/auth/projects", keystone.ListProjectsAPI)
	e.GET("/keystone/v3/auth/domains", keystone.listDomainsAPI)

	e.GET("/keystone/v3/projects", keystone.ListProjectsAPI)
	e.GET("/keystone/v3/projects/:id", keystone.GetProjectAPI)
	e.GET("/keystone/v3/domains", keystone.listDomainsAPI)

	return nil
}

//GetProjectAPI is an API handler to list projects.
func (k *Keystone) GetProjectAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	endpointKey := path.Join("/proxy", clusterID, keystoneService, PrivateURLScope)
	if k.endpointStore.GetEndpoints(PrivateURLScope, endpointKey) {
		c.Request().URL.Path = path.Join(c.Request().URL.Path, c.Param("id"))
		return k.proxyRequest(c, PrivateURLScope, endpointKey)
	}

	token, err := k.validateToken(c.Request())
	if err != nil {
		return err
	}

	// TODO(dfurman): prevent panic: use fields without pointers in models and/or provide getters with nil checks
	for _, role := range token.User.Roles {
		if role.Project.ID == c.Param("id") {
			return c.JSON(http.StatusOK, &ProjectResponse{
				Project: role.Project,
			})
		}
	}

	return c.JSON(http.StatusNotFound, nil)
}

func (k *Keystone) proxyRequest(c echo.Context, scope string, endpointKey string) error {
	r := c.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, LocalAuthPath)
	return proxy.HandleRequest(c, k.endpointStore.ReadAuthURLs(scope, endpointKey, authAPIVersion), k.log)
}

//listDomainsAPI is an API handler to list domains.
func (k *Keystone) listDomainsAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	endpointKey := path.Join("/proxy", clusterID, keystoneService, PrivateURLScope)
	if k.endpointStore.GetEndpoints(PrivateURLScope, endpointKey) {
		return k.proxyRequest(c, PrivateURLScope, endpointKey)
	}

	_, err := k.validateToken(c.Request())
	if err != nil {
		return err
	}
	_, err = k.setAssignment(clusterID)
	if err != nil {
		return err
	}
	domains := k.Assignment.ListDomains()
	domainsResponse := &DomainListResponse{
		Domains: domains,
	}
	return c.JSON(http.StatusOK, domainsResponse)
}

//ListProjectsAPI is an API handler to list projects.
func (k *Keystone) ListProjectsAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	endpointKey := path.Join("/proxy", clusterID, keystoneService, PrivateURLScope)
	if k.endpointStore.GetEndpoints(PrivateURLScope, endpointKey) {
		return k.proxyRequest(c, PrivateURLScope, endpointKey)
	}

	token, err := k.validateToken(c.Request())
	if err != nil {
		return err
	}
	configEndpoint, err := k.setAssignment(clusterID)
	if err != nil {
		return err
	}
	userProjects := []*keystone.Project{}
	user := token.User
	projects := k.Assignment.ListProjects()
	if configEndpoint == "" {
		for _, project := range projects {
			for _, role := range user.Roles {
				if role.Project.Name == project.Name {
					userProjects = append(userProjects, role.Project)
				}
			}
		}
	} else {
		userProjects = append(userProjects, projects...)
	}
	projectsResponse := &ProjectListResponse{
		Projects: userProjects,
	}
	return c.JSON(http.StatusOK, projectsResponse)
}

func (k *Keystone) validateToken(r *http.Request) (*keystone.Token, error) {
	tokenID := r.Header.Get(xAuthTokenKey)
	if tokenID == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	token, ok := k.store.ValidateToken(tokenID)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}

	return token, nil
}

func (k *Keystone) setAssignment(clusterID string) (configEndpoint string, err error) {
	if clusterID == "" {
		return "", nil
	}
	authType, err := k.GetAuthType(clusterID)
	if err != nil {
		logrus.Errorf("Not able to find auth type for cluster %s, %v", clusterID, err)
		return "", err
	}
	if authType != basicAuth {
		return "", nil
	}
	e, ok := k.endpointStore.GetEndpointUrl(clusterID, configService)
	if !ok {
		k.Assignment = k.staticAssignment
		return "", nil
	}
	configEndpoint = e
	apiAssignment := &VNCAPIAssignment{}
	err = apiAssignment.Init(
		configEndpoint, k.staticAssignment.ListUsers())
	if err != nil {
		logrus.Error(err)
		return configEndpoint, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	k.Assignment = apiAssignment
	return configEndpoint, nil
}

// GetAuthType from the cluster configuration
func (k *Keystone) GetAuthType(clusterID string) (authType string, err error) {
	resp, err := k.apiClient.GetContrailCluster(
		context.Background(),
		&services.GetContrailClusterRequest{
			ID: clusterID,
		})
	if err != nil {
		return "", err
	}

	authType = keystoneAuth
	if resp.GetContrailCluster().GetOrchestrator() != openstack {
		authType = basicAuth
	}
	return authType, nil
}

// CreateTokenAPI is an API handler for issuing new authentication token.
func (k *Keystone) CreateTokenAPI(c echo.Context) error {
	sar := keystone.ScopedAuthRequest{}
	if err := c.Bind(&sar); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	clusterID := c.Request().Header.Get(xClusterIDKey)
	endpointKey := path.Join("/proxy", clusterID, keystoneService, PrivateURLScope)
	endpoints := k.endpointStore.GetEndpoints(PrivateURLScope, endpointKey)
	if sar.GetIdentity().Cluster != nil {
		if !endpoints {
			return k.createLocalToken(c, k.newLocalAuthRequest())
		}
		return k.fetchServerTokenWithClusterToken(c, sar.GetIdentity(), PrivateURLScope, endpointKey)
	}

	if endpoints {
		return k.fetchClusterToken(c, sar, PrivateURLScope, endpointKey)
	}

	return k.createLocalToken(c, sar)
}

func (k *Keystone) fetchServerTokenWithClusterToken(
	c echo.Context, i *keystone.Identity, scope string, endpointKey string) error {
	if i.Cluster.Token != nil {
		c.Request().Header.Set(xAuthTokenKey, i.Cluster.Token.ID)
		c.Request().Header.Set(xSubjectTokenKey, i.Cluster.Token.ID)
	}

	return k.proxyRequest(c, scope, endpointKey)
}

func (k *Keystone) newLocalAuthRequest() keystone.AuthRequest {
	scope := keystone.NewScope(
		viper.GetString("client.domain_id"),
		viper.GetString("client.domain_name"),
		viper.GetString("client.project_id"),
		viper.GetString("client.project_name"),
	)
	authRequest := keystone.ScopedAuthRequest{
		Auth: &keystone.ScopedAuth{
			Identity: &keystone.Identity{
				Methods: []string{"password"},
				Password: &keystone.Password{
					User: &keystone.User{
						Name:     viper.GetString("client.id"),
						Password: viper.GetString("client.password"),
						Domain:   scope.GetDomain(),
					},
				},
			},
			Scope: scope,
		},
	}
	return authRequest
}

func (k *Keystone) fetchClusterToken(c echo.Context, sar keystone.AuthRequest, scope string, endpointKey string) error {
	// TODO(dfurman): prevent panic: use fields without pointers in models and/or provide getters with nil checks
	if sar.GetIdentity().Password != nil {
		if sar.GetScope() == nil {
			sar = keystone.UnScopedAuthRequest{
				Auth: &keystone.UnScopedAuth{
					Identity: sar.GetIdentity(),
				},
			}
		}
		// TODO(dfurman): prevent panic: use fields without pointers in models and/or provide getters with nil checks
		if !sar.GetIdentity().Password.User.HasCredential() {
			tokenID := c.Request().Header.Get(xAuthTokenKey)
			if tokenID == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
			}
			_, ok := k.store.ValidateToken(tokenID)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
			}
			k.setCredential(sar, scope, endpointKey)
		}
	}

	c = withRequestBody(c, sar)
	return k.proxyRequest(c, scope, endpointKey)
}

func (k *Keystone) setCredential(sar keystone.AuthRequest, scope string, endpointKey string) {
	username := k.endpointStore.ReadUsername(scope, endpointKey)
	password := k.endpointStore.ReadPassword(scope, endpointKey)
	sar.SetCredential(username, password)
}

func withRequestBody(c echo.Context, ar keystone.AuthRequest) echo.Context {
	b, _ := json.Marshal(ar) // nolint: errcheck
	c.Request().Body = ioutil.NopCloser(bytes.NewReader(b))
	c.Request().ContentLength = int64(len(b))
	return c
}

func (k *Keystone) createLocalToken(c echo.Context, authRequest keystone.AuthRequest) error {
	var err error
	var user *keystone.User
	var token *keystone.Token
	tokenID := ""
	// TODO(dfurman): prevent panic: use fields without pointers in models and/or provide getters with nil checks
	identity := authRequest.GetIdentity()
	if identity.Token != nil {
		tokenID = identity.Token.ID
	}
	if tokenID != "" { // user trying to get a token from token
		token, err = k.store.RetrieveToken(tokenID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}
		user = token.User
	} else {
		// TODO(dfurman): prevent panic: use fields without pointers in models and/or provide getters with nil checks
		user, err = k.Assignment.FetchUser(
			identity.Password.User.Name,
			identity.Password.User.Password,
		)
		if err != nil {
			logrus.WithField("err", err).Debug("User not found")
			return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
		}
		if user == nil {
			logrus.Debug("User not found")
			return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
		}
	}
	var project *keystone.Project
	project, err = filterProject(user, authRequest.GetScope())
	if err != nil {
		logrus.WithField("err", err).Debug("filter project error")
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	tokenID, token = k.store.CreateToken(user, project)
	c.Response().Header().Set(xSubjectTokenKey, tokenID)
	authResponse := &keystone.AuthResponse{
		Token: token,
	}
	return c.JSON(http.StatusOK, authResponse)
}

func filterProject(user *keystone.User, scope *keystone.Scope) (*keystone.Project, error) {
	if scope == nil {
		return nil, nil
	}
	domain := scope.Domain
	if domain != nil {
		if domain.ID != user.Domain.ID {
			return nil, fmt.Errorf("domain unmatched for user %s", user.ID)
		}
	}
	project := scope.Project
	if project == nil {
		return nil, nil
	}
	for _, role := range user.Roles {
		if project.Name != "" {
			if role.Project.Name == project.Name {
				return role.Project, nil
			}
		} else if project.ID != "" {
			if role.Project.ID == project.ID {
				return role.Project, nil
			}
		}
	}
	return nil, nil
}

//ValidateTokenAPI is an API token for validating Token.
func (k *Keystone) ValidateTokenAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	endpointKey := path.Join("/proxy", clusterID, keystoneService, PrivateURLScope)
	if k.endpointStore.GetEndpoints(PrivateURLScope, endpointKey) {
		return k.proxyRequest(c, PrivateURLScope, endpointKey)
	}

	tokenID := c.Request().Header.Get(xAuthTokenKey)
	if tokenID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	token, ok := k.store.ValidateToken(tokenID)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	validateTokenResponse := &keystone.ValidateTokenResponse{
		Token: token,
	}
	return c.JSON(http.StatusOK, validateTokenResponse)
}
