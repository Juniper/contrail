package keystone

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/config"
	kscommon "github.com/Juniper/contrail/pkg/keystone"
)

const (
	configService = "config"
)

//Keystone is used to represents Keystone Controller.
type Keystone struct {
	Store      Store
	Assignment Assignment
	Endpoints  *apicommon.EndpointStore
	Client     *KeystoneClient
	vncClient  *client.HTTP
}

//Init is used to initialize echo with Kesytone capability.
//This function reads config from viper.
func Init(e *echo.Echo, endpoints *apicommon.EndpointStore,
	keystoneClient *KeystoneClient) (*Keystone, error) {
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
		keystone.Assignment = &staticAssignment
	}
	storeType := viper.GetString("keystone.store.type")
	if storeType == "memory" {
		expire := viper.GetInt64("keystone.store.expire")
		keystone.Store = MakeInMemoryStore(time.Duration(expire) * time.Second)
	}
	e.POST("/keystone/v3/auth/tokens", keystone.CreateTokenAPI)
	e.GET("/keystone/v3/auth/tokens", keystone.ValidateTokenAPI)
	e.GET("/keystone/v3/auth/projects", keystone.GetProjectAPI)

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

func getVncConfigEndpoint(endpoints *apicommon.EndpointStore) (configEndpoint string, err error) {
	configEndpoint, err = endpoints.GetEndpoint(configService)
	return configEndpoint, err
}

func (keystone *Keystone) getVncProjects(c echo.Context, configEndpoint string) ([]*kscommon.Project, error) {
	tokenID := c.Request().Header.Get("X-Auth-Token")
	if tokenID == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	_, ok := keystone.Store.ValidateToken(tokenID)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	if keystone.vncClient == nil {
		keystone.vncClient = client.NewHTTP("", "", "", "", true, nil)
	}
	keystone.vncClient.Endpoint = configEndpoint
	keystone.vncClient.Init()
	projectURI := "/projects"
	query := url.Values{"detail": []string{"True"}}
	vncProjectsResponse := &VncProjectListResponse{}
	_, err := keystone.vncClient.ReadWithQuery(
		context.Background(), projectURI, query, vncProjectsResponse)
	if err != nil {
		return nil, err
	}
	projects := []*kscommon.Project{}
	for _, vncProject := range vncProjectsResponse.Projects {
		projects = append(projects, &kscommon.Project{
			Name: vncProject.Project.Name,
			ID:   vncProject.Project.UUID,
		})
	}
	return projects, nil
}

//GetProjectAPI is an API handler to list projects.
func (keystone *Keystone) GetProjectAPI(c echo.Context) error { // nolint: gocyclo
	keystoneEndpoint, err := getKeystoneEndpoint(keystone.Endpoints)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if keystoneEndpoint != "" {
		keystone.Client.SetAuthURL(keystoneEndpoint)
		return keystone.Client.GetProjects(c)
	}

	tokenID := c.Request().Header.Get("X-Auth-Token")
	if tokenID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	token, ok := keystone.Store.ValidateToken(tokenID)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	userProjects := []*kscommon.Project{}
	authType := viper.GetString("auth_type")
	if authType == "basic-auth" {
		configEndpoint, err := getVncConfigEndpoint(keystone.Endpoints)
		if err != nil {
			log.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		if configEndpoint != "" {
			userProjects, err = keystone.getVncProjects(c, configEndpoint)
			if err != nil {
				return err
			}
		}
	}
	user := token.User
	projects := keystone.Assignment.ListProjects()
	for _, project := range projects {
		for _, role := range user.Roles {
			if role.Project.Name == project.Name {
				userProjects = append(userProjects, role.Project)
			}
		}
	}
	projectsResponse := &ProjectListResponse{
		Projects: userProjects,
	}
	return c.JSON(http.StatusOK, projectsResponse)
}

//CreateTokenAPI is an API handler for issuing new Token.
func (keystone *Keystone) CreateTokenAPI(c echo.Context) error {
	keystoneEndpoint, err := getKeystoneEndpoint(keystone.Endpoints)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	if keystoneEndpoint != "" {
		keystone.Client.SetAuthURL(keystoneEndpoint)
		return keystone.Client.CreateToken(c)
	}
	var authRequest kscommon.AuthRequest
	if err = c.Bind(&authRequest); err != nil {
		log.WithField("error", err).Debug("Validation failed")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	var user *kscommon.User
	var token *kscommon.Token
	tokenID := ""
	if authRequest.Auth.Identity.Token != nil {
		tokenID = authRequest.Auth.Identity.Token.ID
	}
	if tokenID != "" { // user trying to get a token from token
		token, err = keystone.Store.RetrieveToken(tokenID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}
		user = token.User
	} else {
		user, err = keystone.Assignment.FetchUser(
			authRequest.Auth.Identity.Password.User.Name,
			authRequest.Auth.Identity.Password.User.Password,
		)
		if err != nil {
			log.WithField("err", err).Debug("User not found")
			return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
		}
		if user == nil {
			log.Debug("User not found")
			return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
		}
	}
	var project *kscommon.Project
	project, err = filterProject(user, authRequest.Auth.Scope)
	if err != nil {
		log.WithField("err", err).Debug("filter project error")
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
	keystoneEndpoint, err := getKeystoneEndpoint(keystone.Endpoints)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	if keystoneEndpoint != "" {
		keystone.Client.SetAuthURL(keystoneEndpoint)
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
