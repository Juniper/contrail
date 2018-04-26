package keystone

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/pkg/errors"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Keystone is used to represents Keystone Controller.
type Keystone struct {
	Store      Store
	Assignment Assignment
}

//Init is used to initialize echo with Kesytone capability.
//This function reads config from viper.
func Init(e *echo.Echo) (*Keystone, error) {
	keystone := &Keystone{}
	assignmentType := viper.GetString("keystone.assignment.type")
	if assignmentType == "static" {
		filepath := viper.GetString("keystone.assignment.file")
		var staticAssignment StaticAssignment
		err := common.LoadFile(filepath, &staticAssignment)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to load static assignment")
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

func filterProject(user *User, scope *Scope) (*Project, error) {
	if scope == nil {
		return nil, nil
	}
	domain := scope.Domain
	if domain != nil {
		if domain.ID != user.Domain.ID {
			return nil, fmt.Errorf("Domain unmatched for user %s", user.ID)
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

//GetProjectAPI is an API handler to list projects.
func (keystone *Keystone) GetProjectAPI(c echo.Context) error {
	tokenID := c.Request().Header.Get("X-Auth-Token")
	if tokenID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	token, ok := keystone.Store.ValidateToken(tokenID)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	user := token.User
	projects := keystone.Assignment.ListProjects()
	userProjects := []*Project{}
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
	var authRequest AuthRequest
	if err := c.Bind(&authRequest); err != nil {
		log.WithField("error", err).Debug("Validation failed")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	var user *User
	tokenID := ""
	if authRequest.Auth.Identity.Token != nil {
		tokenID = authRequest.Auth.Identity.Token.ID
	}
	if tokenID != "" { // user trying to get a token from token
		token, err := keystone.Store.RetrieveToken(tokenID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}
		user = token.User
	} else {
		var err error
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
	project, err := filterProject(user, authRequest.Auth.Scope)
	if err != nil {
		log.WithField("err", err).Debug("filter project error")
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	tokenID, token := keystone.Store.CreateToken(user, project)
	c.Response().Header().Set("X-Subject-Token", tokenID)
	authResponse := &AuthResponse{
		Token: token,
	}
	return c.JSON(http.StatusCreated, authResponse)
}

//ValidateTokenAPI is an API token for validating Token.
func (keystone *Keystone) ValidateTokenAPI(c echo.Context) error {
	tokenID := c.Request().Header.Get("X-Auth-Token")
	if tokenID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	token, ok := keystone.Store.ValidateToken(tokenID)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate")
	}
	validateTokenResponse := &ValidateTokenResponse{
		Token: token,
	}
	return c.JSON(http.StatusOK, validateTokenResponse)
}
