package integration

import (
	"net/http/httptest"
	"testing"

	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/labstack/echo"

	asfkeystone "github.com/Juniper/asf/pkg/keystone"
)

// NewKeystoneServerFake creates started Keystone server fake.
// It registers given user in default domain with admin role.
func NewKeystoneServerFake(t *testing.T, keystoneAuthURL, user, password string) *httptest.Server {
	e := echo.New()

	k, err := keystone.Init()
	if err != nil {
		return nil
	}
	if user != "" {
		k = withKeystoneUser(k, user, password)
	}

	e = withRegisteredRoutes(e, k)
	return httptest.NewServer(e)
}

func withKeystoneUser(k *keystone.Keystone, user, password string) *keystone.Keystone {
	sa := asfkeystone.StaticAssignment{
		Users: make(map[string]*asfkeystone.User),
	}
	sa.Users[user] = &asfkeystone.User{
		Domain:   sa.Domains[DefaultDomainID],
		ID:       user,
		Name:     user,
		Password: password,
		Roles: []*asfkeystone.Role{
			{
				ID:      AdminRoleID,
				Name:    AdminRoleName,
				Project: sa.Projects[AdminProjectID],
			},
		},
	}
	return k
}

func withRegisteredRoutes(e *echo.Echo, k *keystone.Keystone) *echo.Echo {
	e.POST("/v3/auth/tokens", k.CreateTokenAPI)
	e.GET("/v3/auth/tokens", k.ValidateTokenAPI)

	// TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
	e.GET("/v3/auth/projects", k.ListProjectsAPI)

	e.GET("/v3/projects", k.ListProjectsAPI)
	e.GET("/v3/project/:id", k.GetProjectAPI)

	return e
}
