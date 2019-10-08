package integration

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"

	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/labstack/echo"

	pkgkeystone "github.com/Juniper/contrail/pkg/keystone"
)

// ServeKeystoneMock serves mock Keystone server with given Keystone user.
func ServeKeystoneMock(address, keystoneAuthURL, testUser, testPassword string) *httptest.Server {
	e := echo.New()
	k, err := keystone.Init(e, nil, keystone.NewClient(keystoneAuthURL, true))
	if err != nil {
		return nil
	}

	if len(testUser) > 0 {
		addKeystoneUser(k, testUser, testPassword)
	}

	registerRoutes(e, k)

	s := NewWellKnownServer(address, e)
	s.Start()
	return s
}

func addKeystoneUser(k *keystone.Keystone, user, password string) {
	sa, _ := k.Assignment.(*keystone.StaticAssignment) // nolint: errcheck
	sa.Users = map[string]*pkgkeystone.User{}
	sa.Users[user] = &pkgkeystone.User{
		Domain:   sa.Domains[DefaultDomainID],
		ID:       user,
		Name:     user,
		Password: password,
		Roles: []*pkgkeystone.Role{
			{
				ID:      AdminRoleID,
				Name:    AdminRoleName,
				Project: sa.Projects[AdminProjectID],
			},
		},
	}
}

func registerRoutes(e *echo.Echo, k *keystone.Keystone) {
	e.POST("/v3/auth/tokens", k.CreateTokenAPI)
	e.GET("/v3/auth/tokens", k.ValidateTokenAPI)

	// TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
	e.GET("/v3/auth/projects", k.ListProjectsAPI)

	e.GET("/v3/projects", k.ListProjectsAPI)
	e.GET("/v3/project/:id", k.GetProjectAPI)
}

// NewWellKnownServer returns a new server with given port
func NewWellKnownServer(address string, handler http.Handler) *httptest.Server {
	return &httptest.Server{
		Listener: newWellKnownListener(address),
		Config:   &http.Server{Handler: handler},
	}
}

func newWellKnownListener(address string) net.Listener {
	if address != "" {
		l, err := net.Listen("tcp", address)
		if err != nil {
			panic(fmt.Sprintf("httptest: failed to listen on %v: %v", address, err))
		}
		return l
	}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			panic(fmt.Sprintf("httptest: failed to listen on a port: %v", err))
		}
	}
	return l
}
