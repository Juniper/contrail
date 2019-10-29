package keystone

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"time"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/apisrv/endpoint"
	"github.com/Juniper/contrail/pkg/config"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	kscommon "github.com/Juniper/contrail/pkg/keystone"
)

// Keystone constants.
const (
	LocalAuthPath = "/keystone/v3"

	authAPIVersion  = "v3"
	basicAuth       = "basic-auth"
	configService   = "config"
	keystoneAuth    = "keystone"
	keystoneService = "keystone"
	openstack       = "openstack"
	xClusterIDKey   = "X-Cluster-ID"
)

//Keystone is used to represents Keystone Controller.
type Keystone struct {
	store            Store
	Assignment       Assignment
	staticAssignment *StaticAssignment
	endpointStore    *endpoint.Store
	client           *Client
	apiClient        *client.HTTP
	httpClient       *http.Client
}

//Init is used to initialize echo with Keystone capability.
//This function reads config from viper.
func Init(e *echo.Echo, es *endpoint.Store, httpInsecure bool) (*Keystone, error) {
	keystone := &Keystone{
		endpointStore: es,
		client:        &Client{},
		apiClient:     client.NewHTTPFromConfig(),
		httpClient: &http.Client{
			Transport: httpTransport(httpInsecure),
		},
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
	e.POST("/keystone/v3/auth/tokens", keystone.CreateTokenAPI)
	e.GET("/keystone/v3/auth/tokens", keystone.ValidateTokenAPI)

	// TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
	e.GET("/keystone/v3/auth/projects", keystone.ListProjectsAPI)
	e.GET("/keystone/v3/auth/domains", keystone.listDomainsAPI)

	e.GET("/keystone/v3/projects", keystone.ListProjectsAPI)
	e.GET("/keystone/v3/projects/:id", keystone.GetProjectAPI)
	e.GET("/keystone/v3/domains", keystone.listDomainsAPI)

	return keystone, nil
}

//GetProjectAPI is an API handler to list projects.
func (keystone *Keystone) GetProjectAPI(c echo.Context) error {
	ke := getKeystoneEndpoints(c.Request().Header.Get(xClusterIDKey), keystone.endpointStore)
	if len(ke) > 0 {
		return keystone.client.GetProject(c, rawAuthURLs(ke), c.Param("id"))
	}

	token, err := keystone.validateToken(c.Request())
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

//listDomainsAPI is an API handler to list domains.
func (keystone *Keystone) listDomainsAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	ke := getKeystoneEndpoints(clusterID, keystone.endpointStore)
	if len(ke) > 0 {
		return keystone.client.GetDomains(c, rawAuthURLs(ke))
	}

	_, err := keystone.validateToken(c.Request())
	if err != nil {
		return err
	}
	_, err = keystone.setAssignment(clusterID)
	if err != nil {
		return err
	}
	domains := keystone.Assignment.ListDomains()
	domainsResponse := &DomainListResponse{
		Domains: domains,
	}
	return c.JSON(http.StatusOK, domainsResponse)
}

//ListProjectsAPI is an API handler to list projects.
func (keystone *Keystone) ListProjectsAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	ke := getKeystoneEndpoints(clusterID, keystone.endpointStore)
	if len(ke) > 0 {
		return keystone.client.GetProjects(c, rawAuthURLs(ke))
	}

	token, err := keystone.validateToken(c.Request())
	if err != nil {
		return err
	}
	configEndpoint, err := keystone.setAssignment(clusterID)
	if err != nil {
		return err
	}
	userProjects := []*kscommon.Project{}
	user := token.User
	projects := keystone.Assignment.ListProjects()
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

func (keystone *Keystone) validateToken(r *http.Request) (*kscommon.Token, error) {
	tokenID := r.Header.Get("X-Auth-Token")
	if tokenID == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	token, ok := keystone.store.ValidateToken(tokenID)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}

	return token, nil
}

func (keystone *Keystone) setAssignment(clusterID string) (configEndpoint string, err error) {
	if clusterID == "" {
		return "", nil
	}
	authType, err := keystone.GetAuthType(clusterID)
	if err != nil {
		logrus.Errorf("Not able to find auth type for cluster %s, %v", clusterID, err)
		return "", err
	}
	if authType != basicAuth {
		return "", nil
	}
	e := keystone.endpointStore.GetEndpoint(clusterID, configService)
	if e == nil {
		keystone.Assignment = keystone.staticAssignment
		return "", nil
	}
	configEndpoint = e.URL
	apiAssignment := &VNCAPIAssignment{}
	err = apiAssignment.Init(
		configEndpoint, keystone.staticAssignment.ListUsers())
	if err != nil {
		logrus.Error(err)
		return configEndpoint, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	keystone.Assignment = apiAssignment
	return configEndpoint, nil
}

// GetAuthType from the cluster configuration
func (keystone *Keystone) GetAuthType(clusterID string) (authType string, err error) {
	resp, err := keystone.apiClient.GetContrailCluster(
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
func (keystone *Keystone) CreateTokenAPI(c echo.Context) error {
	sar := kscommon.ScopedAuthRequest{}
	if err := c.Bind(&sar); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	ke := getKeystoneEndpoints(c.Request().Header.Get(xClusterIDKey), keystone.endpointStore)
	if sar.GetIdentity().Cluster != nil {
		if len(ke) == 0 {
			return keystone.createLocalToken(c, keystone.newLocalAuthRequest())
		}
		return keystone.fetchServerTokenWithClusterToken(c, sar.GetIdentity(), ke)
	}

	if len(ke) > 0 {
		return keystone.fetchClusterToken(c, sar, ke)
	}

	return keystone.createLocalToken(c, sar)
}

func (keystone *Keystone) fetchServerTokenWithClusterToken(
	c echo.Context, identity *kscommon.Identity, ke []*endpoint.Endpoint,
) error {
	// TODO(dfurman): use keystone.client.CreateToken() instead
	r, err := http.NewRequest(echo.GET, ke[0].URL+"/v3/auth/tokens", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	r.Header.Set("X-Auth-Token", identity.Cluster.Token.ID)
	r.Header.Set("X-Subject-Token", identity.Cluster.Token.ID)

	resp, err := keystone.httpClient.Do(r)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close() // nolint: errcheck

	return c.JSON(resp.StatusCode, resp.Body)
}

func (keystone *Keystone) newLocalAuthRequest() kscommon.AuthRequest {
	scope := kscommon.NewScope(
		viper.GetString("client.domain_id"),
		viper.GetString("client.domain_name"),
		viper.GetString("client.project_id"),
		viper.GetString("client.project_name"),
	)
	authRequest := kscommon.ScopedAuthRequest{
		Auth: &kscommon.ScopedAuth{
			Identity: &kscommon.Identity{
				Methods: []string{"password"},
				Password: &kscommon.Password{
					User: &kscommon.User{
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

func (keystone *Keystone) fetchClusterToken(ctx echo.Context, sar kscommon.AuthRequest, ke []*endpoint.Endpoint) error {
	// TODO(dfurman): prevent panic: use fields without pointers in models and/or provide getters with nil checks
	if sar.GetIdentity().Password != nil {
		if sar.GetScope() == nil {
			sar = kscommon.UnScopedAuthRequest{
				Auth: &kscommon.UnScopedAuth{
					Identity: sar.GetIdentity(),
				},
			}
		}
		// TODO(dfurman): prevent panic: use fields without pointers in models and/or provide getters with nil checks
		if !sar.GetIdentity().Password.User.HasCredential() {
			tokenID := ctx.Request().Header.Get("X-Auth-Token")
			if tokenID == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
			}
			_, ok := keystone.store.ValidateToken(tokenID)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
			}

			// TODO(dfurman): verify if it is acceptable
			sar.SetCredential(ke[0].Username, ke[0].Password)
		}
	}

	ctx = withRequestBody(ctx, sar)
	return keystone.client.CreateToken(ctx, rawAuthURLs(ke))
}

func withRequestBody(c echo.Context, ar kscommon.AuthRequest) echo.Context {
	b, _ := json.Marshal(ar) // nolint: errcheck
	c.Request().Body = ioutil.NopCloser(bytes.NewReader(b))
	c.Request().ContentLength = int64(len(b))
	return c
}

func (keystone *Keystone) createLocalToken(c echo.Context, authRequest kscommon.AuthRequest) error {
	var err error
	var user *kscommon.User
	var token *kscommon.Token
	tokenID := ""
	// TODO(dfurman): prevent panic: use fields without pointers in models and/or provide getters with nil checks
	identity := authRequest.GetIdentity()
	if identity.Token != nil {
		tokenID = identity.Token.ID
	}
	if tokenID != "" { // user trying to get a token from token
		token, err = keystone.store.RetrieveToken(tokenID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}
		user = token.User
	} else {
		// TODO(dfurman): prevent panic: use fields without pointers in models and/or provide getters with nil checks
		user, err = keystone.Assignment.FetchUser(
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
	var project *kscommon.Project
	project, err = filterProject(user, authRequest.GetScope())
	if err != nil {
		logrus.WithField("err", err).Debug("filter project error")
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	tokenID, token = keystone.store.CreateToken(user, project)
	c.Response().Header().Set("X-Subject-Token", tokenID)
	authResponse := &kscommon.AuthResponse{
		Token: token,
	}
	return c.JSON(http.StatusOK, authResponse)
}

func filterProject(user *kscommon.User, scope *kscommon.Scope) (*kscommon.Project, error) {
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
func (keystone *Keystone) ValidateTokenAPI(c echo.Context) error {
	ke := getKeystoneEndpoints(c.Request().Header.Get(xClusterIDKey), keystone.endpointStore)
	if ke != nil {
		return keystone.client.ValidateToken(c, rawAuthURLs(ke))
	}

	tokenID := c.Request().Header.Get("X-Auth-Token")
	if tokenID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	token, ok := keystone.store.ValidateToken(tokenID)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	validateTokenResponse := &kscommon.ValidateTokenResponse{
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
