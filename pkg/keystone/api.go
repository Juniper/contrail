package keystone

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/config"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	asfkeystone "github.com/Juniper/asf/pkg/keystone"
)

// Keystone constants.
const (
	LocalAuthPath = "/keystone/v3"
	authAPIVersion   = "v3"
	xAuthTokenKey    = "X-Auth-Token"
	xSubjectTokenKey = "X-Subject-Token"
)

//Keystone is used to represents Keystone Controller.
type Keystone struct {
	store      Store
	Assignment *asfkeystone.StaticAssignment
	apiClient  *client.HTTP

	log *logrus.Entry
}

//Init is used to initialize echo with Keystone capability.
//This function reads config from viper.
func Init() (*Keystone, error) {
	keystone := &Keystone{
		apiClient:     client.NewHTTPFromConfig(),
		log:           logutil.NewLogger("keystone-api"),
	}
	assignmentType := viper.GetString("keystone.assignment.type")
	if assignmentType == "static" {
		var staticAssignment asfkeystone.StaticAssignment
		err := config.LoadConfig("keystone.assignment.data", &staticAssignment)
		if err != nil {
			return nil, errors.Wrap(err, "creating local keystone server: failed to parse keystone assignment configuration")
		}
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
func (k *Keystone) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.POST("/keystone/v3/auth/tokens", k.CreateTokenAPI, apiserver.WithNoAuth())
	r.GET("/keystone/v3/auth/tokens", k.ValidateTokenAPI, apiserver.WithNoAuth())

	// TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
	r.GET(
		"/keystone/v3/auth/projects",
		k.ListAuthProjectsAPI,
		apiserver.WithNoAuth(),
		apiserver.WithHomepageType(apiserver.CollectionEndpoint),
	)
	r.GET("/keystone/v3/auth/domains", k.listDomainsAPI, apiserver.WithHomepageType(apiserver.CollectionEndpoint))

	r.GET(
		"/keystone/v3/projects",
		k.ListProjectsAPI,
		apiserver.WithNoAuth(),
		apiserver.WithHomepageType(apiserver.CollectionEndpoint),
	)
	r.GET("/keystone/v3/projects/:id", k.GetProjectAPI)
	r.GET("/keystone/v3/domains", k.listDomainsAPI, apiserver.WithHomepageType(apiserver.CollectionEndpoint))

	r.GET("/keystone/v3/users", k.ListUsersAPI, apiserver.WithHomepageType(apiserver.CollectionEndpoint))
}

// RegisterGRPCAPI does nothing, as Keystone has no GRPC API.
func (*Keystone) RegisterGRPCAPI(r apiserver.GRPCRouter) {
}

//GetProjectAPI is an API handler to list projects.
func (k *Keystone) GetProjectAPI(c echo.Context) error {
	token, err := k.validateToken(c.Request())
	if err != nil {
		return err
	}

	// TODO(dfurman): prevent panic: use fields without pointers in models and/or provide getters with nil checks
	for _, role := range token.User.Roles {
		if role.Project.ID == c.Param("id") {
			return c.JSON(http.StatusOK, &asfkeystone.ProjectResponse{
				Project: role.Project,
			})
		}
	}

	return c.JSON(http.StatusNotFound, nil)
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
	_, err := k.validateToken(c.Request())
	if err != nil {
		return err
	}
	domains := k.Assignment.ListDomains()
	domainsResponse := &asfkeystone.DomainListResponse{
		Domains: domains,
	}
	return c.JSON(http.StatusOK, domainsResponse)
}

//ListAuthProjectsAPI is an API handler to list projects available to be scoped to based on the token.
func (k *Keystone) ListAuthProjectsAPI(c echo.Context) error {
	_, err := k.validateToken(c.Request())
	if err != nil {
		return err
	}
	userProjects := []*keystone.Project{}
	projects := k.Assignment.ListProjects()
	userProjects = append(userProjects, projects...)
	projectsResponse := &asfkeystone.ProjectListResponse{
		Projects: userProjects,
	}
	return c.JSON(http.StatusOK, projectsResponse)
}

//ListProjectsAPI is an API handler to list projects.
func (k *Keystone) ListProjectsAPI(c echo.Context) error {
	_, err := k.validateToken(c.Request())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &asfkeystone.ProjectListResponse{
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

// CreateTokenAPI is an API handler for issuing new authentication token.
func (k *Keystone) CreateTokenAPI(c echo.Context) error {
	sar := keystone.ScopedAuthRequest{}
	if err := c.Bind(&sar); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	return k.createLocalToken(c, sar)
}

//ListUsersAPI is an API handler to list users.
func (k *Keystone) ListUsersAPI(c echo.Context) error {
	_, err := k.validateToken(c.Request())
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
