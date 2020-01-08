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
	"github.com/Juniper/contrail/pkg/apisrv/baseapisrv"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/config"
	"github.com/Juniper/contrail/pkg/endpoint"
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

//Keystone is used to represents Keystone Controller.
type Keystone struct {
	store            Store
	Assignment       Assignment
	staticAssignment *StaticAssignment
	endpointStore    *endpoint.Store
	apiClient        *client.HTTP

	log *logrus.Entry
}

//Init is used to initialize echo with Keystone capability.
//This function reads config from viper.
func Init(es *endpoint.Store) (*Keystone, error) {
	keystone := &Keystone{
		endpointStore: es,
		apiClient:     client.NewHTTPFromConfig(),
		log:           logutil.NewLogger("keystone-api"),
	}
	assignmentType := viper.GetString("keystone.assignment.type")
	if assignmentType == "static" {
		var staticAssignment StaticAssignment
		err := config.LoadConfig("keystone.assignment.data", &staticAssignment)
		if err != nil {
			return nil, err
		}
		keystone.staticAssignment = &staticAssignment
		keystone.Assignment = &staticAssignment
	}
	storeType := viper.GetString("keystone.store.type")
	if storeType == "memory" {
		expire := viper.GetInt64("keystone.store.expire")
		keystone.store = MakeInMemoryStore(time.Duration(expire) * time.Second)
	}

	return keystone, nil
}

// RegisterHTTPAPI registers local Keystone endpoints.
func (k *Keystone) RegisterHTTPAPI(r baseapisrv.HTTPRouter) error {
	r.POST("/keystone/v3/auth/tokens", k.CreateTokenAPI)
	r.GET("/keystone/v3/auth/tokens", k.ValidateTokenAPI)

	// TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
	r.GET("/keystone/v3/auth/projects", k.ListProjectsAPI)
	r.GET("/keystone/v3/auth/domains", k.listDomainsAPI)

	r.GET("/keystone/v3/projects", k.ListProjectsAPI)
	r.GET("/keystone/v3/projects/:id", k.GetProjectAPI)
	r.GET("/keystone/v3/domains", k.listDomainsAPI)

	return nil
}

// RegisterGRPCAPI does nothing, as Keystone has no GRPC API.
func (*Keystone) RegisterGRPCAPI(r baseapisrv.GRPCRouter) error {
	return nil
}

//GetProjectAPI is an API handler to list projects.
func (k *Keystone) GetProjectAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	if ke := getKeystoneEndpoints(clusterID, k.endpointStore); len(ke) > 0 {
		c.Request().URL.Path = path.Join(c.Request().URL.Path, c.Param("id"))
		return k.proxyRequest(c, ke)
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

func (k *Keystone) proxyRequest(c echo.Context, endpoints []*endpoint.Endpoint) error {
	r := c.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, LocalAuthPath)
	return proxy.HandleRequest(c, rawAuthURLs(endpoints), k.log)
}

//listDomainsAPI is an API handler to list domains.
func (k *Keystone) listDomainsAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	if ke := getKeystoneEndpoints(clusterID, k.endpointStore); len(ke) > 0 {
		return k.proxyRequest(c, ke)
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
	if ke := getKeystoneEndpoints(clusterID, k.endpointStore); len(ke) > 0 {
		return k.proxyRequest(c, ke)
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
	e := k.endpointStore.GetEndpoint(clusterID, configService)
	if e == nil {
		k.Assignment = k.staticAssignment
		return "", nil
	}
	configEndpoint = e.URL
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

	ke := getKeystoneEndpoints(c.Request().Header.Get(xClusterIDKey), k.endpointStore)
	if sar.GetIdentity().Cluster != nil {
		if len(ke) == 0 {
			return k.createLocalToken(c, k.newLocalAuthRequest())
		}
		return k.fetchServerTokenWithClusterToken(c, sar.GetIdentity(), ke)
	}

	if len(ke) > 0 {
		return k.fetchClusterToken(c, sar, ke)
	}

	return k.createLocalToken(c, sar)
}

func (k *Keystone) fetchServerTokenWithClusterToken(
	c echo.Context, i *keystone.Identity, ke []*endpoint.Endpoint,
) error {
	if i.Cluster.Token != nil {
		c.Request().Header.Set(xAuthTokenKey, i.Cluster.Token.ID)
		c.Request().Header.Set(xSubjectTokenKey, i.Cluster.Token.ID)
	}

	return k.proxyRequest(c, ke)
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

func (k *Keystone) fetchClusterToken(c echo.Context, sar keystone.AuthRequest, ke []*endpoint.Endpoint) error {
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

			sar.SetCredential(ke[0].Username, ke[0].Password)
		}
	}

	c = withRequestBody(c, sar)
	return k.proxyRequest(c, ke)
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
	if ke := getKeystoneEndpoints(clusterID, k.endpointStore); len(ke) > 0 {
		return k.proxyRequest(c, ke)
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

func getKeystoneEndpoints(clusterID string, es *endpoint.Store) []*endpoint.Endpoint {
	if es == nil {
		// getKeystoneEndpoints called from CreateTokenAPI, ValidateTokenAPI or GetProjectAPI of the mock keystone
		return nil
	}

	// TODO(dfurman): "server.dynamic_proxy_path" or DefaultDynamicProxyPath should be used
	ts := es.Read(path.Join("/proxy", clusterID, keystoneService, endpoint.PrivateURLScope))
	if ts == nil {
		return nil
	}

	return ts.ReadAll(endpoint.PrivateURLScope)
}

func rawAuthURLs(targets []*endpoint.Endpoint) []string {
	var u []string
	for _, target := range targets {
		u = append(u, withAPIVersionSuffix(target.URL))
	}
	return u
}

func withAPIVersionSuffix(url string) string {
	return fmt.Sprintf("%s/%s", url, authAPIVersion)
}
