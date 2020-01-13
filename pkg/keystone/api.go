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
	privateURLScope  = "private"
	xAuthTokenKey    = "X-Auth-Token"
	xSubjectTokenKey = "X-Subject-Token"
	xClusterIDKey    = "X-Cluster-ID"
)

type endpointStore interface {
	GetUsername(prefix string, endpointKey string) string
	GetPassword(prefix string, endpointKey string) string
	GetEndpointURL(clusterID, prefix string) (string, bool)
	GetEndpointURLs(prefix string, endpointKey string) []string
}

//Keystone is used to represents Keystone Controller.
type Keystone struct {
	store            Store
	Assignment       Assignment
	staticAssignment *StaticAssignment
	endpointStore    endpointStore
	apiClient        *client.HTTP

	log *logrus.Entry
}

//Init is used to initialize echo with Keystone capability.
//This function reads config from viper.
func Init(e *echo.Echo, es endpointStore) (*Keystone, error) {
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
	e.POST("/keystone/v3/auth/tokens", keystone.CreateTokenAPI)
	e.GET("/keystone/v3/auth/tokens", keystone.ValidateTokenAPI)

	// TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
	e.GET("/keystone/v3/auth/projects", keystone.ListAuthProjectsAPI)
	e.GET("/keystone/v3/auth/domains", keystone.listDomainsAPI)

	e.GET("/keystone/v3/projects", keystone.ListProjectsAPI)
	e.GET("/keystone/v3/projects/:id", keystone.GetProjectAPI)
	e.GET("/keystone/v3/domains", keystone.listDomainsAPI)

	e.GET("/keystone/v3/users", keystone.ListUsersAPI)

	return keystone, nil
}

// NoAuthPaths returns paths that require no authentication.
func (k *Keystone) NoAuthPaths() []string {
	return []string{
		"/keystone/v3/auth/tokens",
		"/keystone/v3/projects",
		"/keystone/v3/auth/projects", // TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
	}
}

//GetProjectAPI is an API handler to list projects.
func (k *Keystone) GetProjectAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	if ke := getKeystoneEndpoints(clusterID, k.endpointStore); len(ke) > 0 {
		c.Request().URL.Path = path.Join(c.Request().URL.Path, c.Param("id"))
		return k.proxyRequest(c)
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

func (k *Keystone) proxyRequest(c echo.Context) error {
	r := c.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, LocalAuthPath)
	clusterID := c.Request().Header.Get(xClusterIDKey)
	ke := getKeystoneEndpoints(clusterID, k.endpointStore)
	return proxy.HandleRequest(c, rawAuthURLs(ke), k.log)
}

func rawAuthURLs(urls []string) (u []string) {
	for _, url := range urls {
		u = append(u, withAPIVersionSuffix(url))
	}
	return
}

func withAPIVersionSuffix(url string) string {
	return fmt.Sprintf("%s/%s", url, authAPIVersion)
}

//listDomainsAPI is an API handler to list domains.
func (k *Keystone) listDomainsAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	if ke := getKeystoneEndpoints(clusterID, k.endpointStore); len(ke) > 0 {
		return k.proxyRequest(c)
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

//ListAuthProjectsAPI is an API handler to list projects available to be scoped to based on the token.
func (k *Keystone) ListAuthProjectsAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	if ke := getKeystoneEndpoints(clusterID, k.endpointStore); len(ke) > 0 {
		return k.proxyRequest(c)
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

//ListProjectsAPI is an API handler to list projects.
func (k *Keystone) ListProjectsAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	if ke := getKeystoneEndpoints(clusterID, k.endpointStore); len(ke) > 0 {
		return k.proxyRequest(c)
	}
	_, err := k.validateToken(c.Request())
	if err != nil {
		return err
	}
	_, err = k.setAssignment(clusterID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &ProjectListResponse{
		Projects: k.Assignment.ListProjects(),
	})
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
	e, ok := k.endpointStore.GetEndpointURL(clusterID, configService)
	if !ok {
		k.Assignment = k.staticAssignment
		return "", nil
	}
	apiAssignment := &VNCAPIAssignment{}
	err = apiAssignment.Init(
		e, k.staticAssignment.Users)
	if err != nil {
		logrus.Error(err)
		return e, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	k.Assignment = apiAssignment
	return e, nil
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
	endpointKey := path.Join("/proxy", clusterID, keystoneService, privateURLScope)
	ke := getKeystoneEndpoints(clusterID, k.endpointStore)
	if sar.GetIdentity().Cluster != nil {
		if len(ke) == 0 {
			return k.createLocalToken(c, k.newLocalAuthRequest())
		}
		return k.fetchServerTokenWithClusterToken(c, sar.GetIdentity())
	}

	if len(ke) > 0 {
		return k.fetchClusterToken(c, sar, privateURLScope, endpointKey)
	}

	return k.createLocalToken(c, sar)
}

//ListUsersAPI is an API handler to list users.
func (k *Keystone) ListUsersAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	if ke := getKeystoneEndpoints(clusterID, k.endpointStore); len(ke) > 0 {
		return k.proxyRequest(c)
	}
	_, err := k.validateToken(c.Request())
	if err != nil {
		return err
	}
	_, err = k.setAssignment(clusterID)
	if err != nil {
		return err
	}

	n := c.QueryParam("name")
	if n != "" {
		for _, u := range k.Assignment.ListUsers() {
			if u.Name == n {
				return c.JSON(http.StatusOK, &UserListResponse{
					Users: []*keystone.User{u},
				})
			}
		}
		return c.JSON(http.StatusOK, &UserListResponse{
			Users: nil,
		})
	}
	return c.JSON(http.StatusOK, &UserListResponse{
		Users: k.Assignment.ListUsers(),
	})
}

func (k *Keystone) fetchServerTokenWithClusterToken(
	c echo.Context, i *keystone.Identity) error {
	if i.Cluster.Token != nil {
		c.Request().Header.Set(xAuthTokenKey, i.Cluster.Token.ID)
		c.Request().Header.Set(xSubjectTokenKey, i.Cluster.Token.ID)
	}

	return k.proxyRequest(c)
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
			sar.SetCredential(k.endpointStore.GetUsername(scope, endpointKey), k.endpointStore.GetPassword(scope, endpointKey))
		}
	}

	c = withRequestBody(c, sar)
	return k.proxyRequest(c)
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
		return k.proxyRequest(c)
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

func getKeystoneEndpoints(clusterID string, es endpointStore) []string {
	if es == nil {
		// getKeystoneEndpoints called from CreateTokenAPI, ValidateTokenAPI or GetProjectAPI of the mock keystone
		return nil
	}

	// TODO(dfurman): "server.dynamic_proxy_path" or DefaultDynamicProxyPath should be used
	return es.GetEndpointURLs(privateURLScope, path.Join("/proxy", clusterID, keystoneService, privateURLScope))
}
