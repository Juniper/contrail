package keystone

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/config"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	kscommon "github.com/Juniper/contrail/pkg/keystone"
)

// Keystone constants.
const (
	AuthEndpointSuffix = "/keystone/v3" // TODO(Daniel) use this constant where possible

	configService = "config"
	xClusterIDKey = "X-Cluster-ID"
)

//Keystone is used to represents Keystone Controller.
type Keystone struct {
	Store      Store
	Assignment Assignment
	Endpoints  *apicommon.EndpointStore
	Client     *Client

	staticAssignment *StaticAssignment
}

//Init is used to initialize echo with Keystone capability.
//This function reads config from viper.
func Init(e *echo.Echo, endpoints *apicommon.EndpointStore,
	keystoneClient *Client) (*Keystone, error) {
	keystone := &Keystone{
		Endpoints: endpoints,
		Client:    keystoneClient,
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
		keystone.Store = MakeInMemoryStore(time.Duration(expire) * time.Second)
	}
	e.POST("/keystone/v3/auth/tokens", keystone.CreateTokenAPI)
	e.GET("/keystone/v3/auth/tokens", keystone.ValidateTokenAPI)

	// TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
	e.GET("/keystone/v3/auth/projects", keystone.ListProjectsAPI)
	e.GET("/keystone/v3/auth/domains", keystone.ListDomainsAPI)

	e.GET("/keystone/v3/projects", keystone.ListProjectsAPI)
	e.GET("/keystone/v3/projects/:id", keystone.GetProjectAPI)
	e.GET("/keystone/v3/domains", keystone.ListDomainsAPI)

	return keystone, nil
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

func (keystone *Keystone) setAssignment(clusterID string) (configEndpoint string, err error) {
	authType := viper.GetString("auth_type")
	if authType != "basic-auth" {
		return "", nil
	}
	e := keystone.Endpoints.GetEndpoint(clusterID, configService)
	if e != nil {
		configEndpoint = e.URL
		apiAssignment := &VNCAPIAssignment{}
		err := apiAssignment.Init(
			configEndpoint, keystone.staticAssignment.ListUsers())
		if err != nil {
			logrus.Error(err)
			return configEndpoint, echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		keystone.Assignment = apiAssignment
	} else {
		keystone.Assignment = keystone.staticAssignment
	}
	return configEndpoint, nil
}

func (keystone *Keystone) validateToken(r *http.Request) (*kscommon.Token, error) {
	tokenID := r.Header.Get("X-Auth-Token")
	if tokenID == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	token, ok := keystone.Store.ValidateToken(tokenID)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}

	return token, nil
}

//GetProjectAPI is an API handler to list projects.
func (keystone *Keystone) GetProjectAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	keystoneEndpoint := getKeystoneEndpoint(clusterID, keystone.Endpoints)

	id := c.Param("id")
	if keystoneEndpoint != nil {
		keystone.Client.SetAuthURL(keystoneEndpoint.URL)
		return keystone.Client.GetProject(c, id)
	}

	token, err := keystone.validateToken(c.Request())
	if err != nil {
		return err
	}

	user := token.User
	for _, role := range user.Roles {
		if role.Project.ID == id {
			return c.JSON(http.StatusOK, &ProjectResponse{
				Project: role.Project,
			})
		}
	}

	return c.JSON(http.StatusNotFound, nil)
}

//ListDomainsAPI is an API handler to list domains.
func (keystone *Keystone) ListDomainsAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	keystoneEndpoint := getKeystoneEndpoint(clusterID, keystone.Endpoints)
	if keystoneEndpoint != nil {
		keystone.Client.SetAuthURL(keystoneEndpoint.URL)
		return keystone.Client.GetDomains(c)
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
	keystoneEndpoint := getKeystoneEndpoint(clusterID, keystone.Endpoints)
	if keystoneEndpoint != nil {
		keystone.Client.SetAuthURL(keystoneEndpoint.URL)
		return keystone.Client.GetProjects(c)
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

func (keystone *Keystone) fetchServerTokenWithClusterToken(
	c echo.Context, identity *kscommon.Identity) error {
	clusterID := identity.Cluster.ID
	keystoneEndpoint := getKeystoneEndpoint(clusterID, keystone.Endpoints)
	if keystoneEndpoint != nil {
		tokenURL := keystoneEndpoint.URL + "/v3/auth/tokens"
		request, err := http.NewRequest(echo.GET, tokenURL, nil)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		request.Header.Set("X-Auth-Token", identity.Cluster.Token.ID)
		request.Header.Set("X-Subject-Token", identity.Cluster.Token.ID)
		resp, err := keystone.Client.httpClient.Do(request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		defer resp.Body.Close() // nolint: errcheck
		validateTokenResponse := &kscommon.ValidateTokenResponse{}
		if err = json.NewDecoder(resp.Body).Decode(validateTokenResponse); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		if resp.StatusCode != 200 {
			return c.JSON(resp.StatusCode, validateTokenResponse)
		}
	}
	// Get token from local keystone
	return keystone.createToken(c, keystone.newLocalAuthRequest())
}

func (keystone *Keystone) fetchClusterToken(c echo.Context,
	identity *kscommon.Identity, authRequest kscommon.AuthRequest,
	keystoneEndpoint *apicommon.Endpoint) error {
	keystone.Client.SetAuthURL(keystoneEndpoint.URL)
	if identity.Password != nil {
		if authRequest.GetScope() == nil {
			unScopedRequest := kscommon.UnScopedAuthRequest{
				Auth: &kscommon.UnScopedAuth{
					Identity: identity,
				},
			}
			authRequest = unScopedRequest
		}
		if !identity.Password.User.HasCredential() {
			tokenID := c.Request().Header.Get("X-Auth-Token")
			if tokenID == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
			}
			_, ok := keystone.Store.ValidateToken(tokenID)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
			}
			authRequest.SetCredential(
				keystoneEndpoint.Username, keystoneEndpoint.Password)
		}
	}
	c = keystone.Client.SetAuthIdentity(c, authRequest)
	return keystone.Client.CreateToken(c)
}

//CreateTokenAPI is an API handler for issuing new Token.
func (keystone *Keystone) CreateTokenAPI(c echo.Context) error {
	var authRequest kscommon.AuthRequest
	scopedRequest := kscommon.ScopedAuthRequest{}
	if err := c.Bind(&scopedRequest); err != nil {
		logrus.WithField("error", err).Debug("Validation failed")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	authRequest = scopedRequest
	identity := authRequest.GetIdentity()
	if identity.Cluster != nil {
		return keystone.fetchServerTokenWithClusterToken(c, identity)
	}
	clusterID := c.Request().Header.Get(xClusterIDKey)
	keystoneEndpoint := getKeystoneEndpoint(clusterID, keystone.Endpoints)
	if keystoneEndpoint != nil {
		return keystone.fetchClusterToken(c, identity, authRequest, keystoneEndpoint)
	}
	return keystone.createToken(c, authRequest)
}

func (keystone *Keystone) createToken(c echo.Context, authRequest kscommon.AuthRequest) error {
	var err error
	var user *kscommon.User
	var token *kscommon.Token
	tokenID := ""
	identity := authRequest.GetIdentity()
	if identity.Token != nil {
		tokenID = identity.Token.ID
	}
	if tokenID != "" { // user trying to get a token from token
		token, err = keystone.Store.RetrieveToken(tokenID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}
		user = token.User
	} else {
		clusterID := c.Request().Header.Get(xClusterIDKey)
		_, err = keystone.setAssignment(clusterID)
		if err != nil {
			return err
		}
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
	tokenID, token = keystone.Store.CreateToken(user, project)
	c.Response().Header().Set("X-Subject-Token", tokenID)
	authResponse := &kscommon.AuthResponse{
		Token: token,
	}
	return c.JSON(http.StatusOK, authResponse)
}

//ValidateTokenAPI is an API token for validating Token.
func (keystone *Keystone) ValidateTokenAPI(c echo.Context) error {
	clusterID := c.Request().Header.Get(xClusterIDKey)
	keystoneEndpoint := getKeystoneEndpoint(clusterID, keystone.Endpoints)
	if keystoneEndpoint != nil {
		keystone.Client.SetAuthURL(keystoneEndpoint.URL)
		return keystone.Client.ValidateToken(c)
	}

	tokenID := c.Request().Header.Get("X-Auth-Token")
	if tokenID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	token, ok := keystone.Store.ValidateToken(tokenID)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	validateTokenResponse := &kscommon.ValidateTokenResponse{
		Token: token,
	}
	return c.JSON(http.StatusOK, validateTokenResponse)
}
