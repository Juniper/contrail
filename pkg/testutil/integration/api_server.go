package integration

import (
	"net/http/httptest"
	"path"
	"testing"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/testutil"

	kscommon "github.com/Juniper/contrail/pkg/keystone"
	integrationetcd "github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

const (
	authEndpointSuffix = "/keystone/v3"
	dbUser             = "root"
	dbPassword         = "contrail123"
	dbName             = "contrail_test"
)

// Keystone credentials
const (
	DefaultDomainID   = "default"
	DefaultDomainName = "DefaultDomain"
	AdminProjectID    = "admin"
	AdminProjectName  = "AdminProject"
	AdminRoleID       = "admin"
	AdminRoleName     = "admin"
	AdminUserID       = "alice"
	AdminUserName     = "Alice"
	AdminUserPassword = "alice_password"
)

// APIServer is embedded API Server for testing purposes.
type APIServer struct {
	APIServer  *apisrv.Server
	TestServer *httptest.Server
	log        *logrus.Entry
}

// APIServerConfig contains parameters for test API Server.
type APIServerConfig struct {
	CacheDB            *cache.DB
	DBDriver           string
	RepoRootPath       string
	LogLevel           string
	EnableEtcdNotifier bool
	DisableLogAPI      bool
	EnableRBAC         bool
}

// NewRunningAPIServer creates new running test API Server for testing purposes.
// Call Close() method to release its resources.
func NewRunningAPIServer(t *testing.T, c *APIServerConfig) *APIServer {
	s, err := NewRunningServer(c)
	require.NoError(t, err)

	return s
}

// NewRunningServer creates new running API server with default testing configuration.
// Call Close() method to release its resources.
func NewRunningServer(c *APIServerConfig) (*APIServer, error) {
	setDefaultViperConfig(c)

	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		return nil, err
	}

	s, err := apisrv.NewServer()
	if err != nil {
		return nil, errors.Wrapf(err, "creating API Server failed")
	}
	s.Cache = c.CacheDB

	ts := testutil.NewTestHTTPServer(s.Echo)
	viper.Set("keystone.authurl", ts.URL+authEndpointSuffix)
	viper.Set("client.endpoint", ts.URL)

	if err = s.Init(); err != nil {
		return nil, errors.Wrapf(err, "initialization of test API Server failed")
	}

	return &APIServer{
		APIServer:  s,
		TestServer: ts,
		log:        logutil.NewLogger("api-server"),
	}, nil
}

func setDefaultViperConfig(c *APIServerConfig) {
	setViperConfig(map[string]interface{}{
		"database.type":               c.DBDriver,
		"database.host":               "localhost",
		"database.user":               dbUser,
		"database.name":               dbName,
		"database.password":           dbPassword,
		"database.dialect":            c.DBDriver,
		"database.max_open_conn":      100,
		"database.connection_retries": 10,
		"database.retry_period":       3,
		"database.debug":              true,
		constants.ETCDPathVK:          integrationetcd.Prefix,
		"keystone.local":              true,
		"keystone.assignment.type":    "static",
		"keystone.assignment.data":    keystoneAssignment(),
		"keystone.store.type":         "memory",
		"keystone.store.expire":       3600,
		"keystone.insecure":           true,
		"log_level":                   c.LogLevel,
		"server.notify_etcd":          c.EnableEtcdNotifier,
		"server.read_timeout":         10,
		"server.write_timeout":        5,
		"server.log_api":              !c.DisableLogAPI,
		"server.log_body":             !c.DisableLogAPI,
		"static_files.public":         path.Join(c.RepoRootPath, "public"),
		"server.enable_vnc_neutron":   true,
		"tls.enabled":                 false,
		"aaa_mode":                    rbacConfig(c.EnableRBAC),
	})
}

func rbacConfig(enableRBAC bool) string {
	if enableRBAC {
		return "rbac"
	}
	return ""
}

func keystoneAssignment() *keystone.StaticAssignment {
	a := keystone.StaticAssignment{
		Domains: map[string]*kscommon.Domain{
			DefaultDomainID: {
				ID:   DefaultDomainID,
				Name: DefaultDomainName,
			},
		},
		Projects: make(map[string]*kscommon.Project),
		Users:    make(map[string]*kscommon.User),
	}
	a.Projects[AdminProjectID] = &kscommon.Project{
		Domain: a.Domains[DefaultDomainID],
		ID:     AdminProjectID,
		Name:   AdminProjectName,
	}
	a.Users[AdminUserID] = &kscommon.User{
		Domain:   a.Domains[DefaultDomainID],
		ID:       AdminUserID,
		Name:     AdminUserName,
		Password: AdminUserPassword,
		Roles: []*kscommon.Role{
			{
				ID:      AdminRoleID,
				Name:    AdminRoleName,
				Project: a.Projects[AdminProjectID],
			},
		},
	}
	return &a
}

func setViperConfig(config map[string]interface{}) {
	for k, v := range config {
		viper.SetDefault(k, v)
	}
}

// URL returns server base URL.
func (s *APIServer) URL() string {
	return s.TestServer.URL
}

// CloseT closes server.
func (s *APIServer) CloseT(t *testing.T) {
	s.log.Debug("Closing test API server")
	err := s.Close()
	assert.NoError(t, err, "closing API Server failed")

}

// Close closes server.
func (s *APIServer) Close() error {
	s.TestServer.Close()
	return s.APIServer.Close()
}

// ForceProxyUpdate requests an immediate update of endpoints and waits for its completion.
func (s *APIServer) ForceProxyUpdate() {
	s.APIServer.Proxy.ForceUpdate()
}
