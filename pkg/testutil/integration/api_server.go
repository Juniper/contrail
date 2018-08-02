package integration

import (
	"fmt"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/testutil"
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
	AdminRoleName     = "AdminRole"
	AdminUserID       = "alice"
	AdminUserName     = "Alice"
	AdminUserPassword = "alice_password"
)

// APIServer is embedded API Server for testing purposes.
type APIServer struct {
	apiServer  *apisrv.Server
	testServer *httptest.Server
	log        *logrus.Entry
}

// NewRunningAPIServer creates new running test API Server.
// Call Close() method to release its resources.
func NewRunningAPIServer(t *testing.T, repoRootPath, dbDriver string, dbDebug bool) *APIServer {
	setViperConfig(map[string]interface{}{
		"database.type":               dbDriver,
		"database.host":               "localhost",
		"database.user":               dbUser,
		"database.name":               dbName,
		"database.password":           dbPassword,
		"database.dialect":            dbDriver,
		"database.max_open_conn":      100,
		"database.connection_retries": 10,
		"database.retry_period":       3,
		"database.debug":              dbDebug,
		"etcd.path":                   EtcdJSONPrefix,
		"keystone.local":              true,
		"keystone.assignment.type":    "static",
		"keystone.assignment.data":    keystoneAssignment(),
		"keystone.store.type":         "memory",
		"keystone.store.expire":       3600,
		"keystone.insecure":           true,
		"log_level":                   "debug",
		"server.read_timeout":         10,
		"server.write_timeout":        5,
		"server.log_api":              true,
		"static_files.public":         path.Join(repoRootPath, "public"),
		"tls.enabled":                 false,
	})
	configureDebugLogging(t)

	log := pkglog.NewLogger("api-server")
	log.WithField("config", fmt.Sprintf("%+v", viper.AllSettings())).Debug("Creating API Server")
	s, err := apisrv.NewServer()
	require.NoError(t, err, "creating API Server failed")

	ts := testutil.NewTestHTTPServer(s.Echo)

	viper.Set("keystone.authurl", ts.URL+authEndpointSuffix)
	err = s.Init()
	require.NoError(t, err, "initialization of test API Server failed")

	return &APIServer{
		apiServer:  s,
		testServer: ts,
		log:        log,
	}
}

func keystoneAssignment() *keystone.StaticAssignment {
	a := keystone.StaticAssignment{
		Domains: map[string]*keystone.Domain{
			DefaultDomainID: {
				ID:   DefaultDomainID,
				Name: DefaultDomainName,
			},
		},
		Projects: make(map[string]*keystone.Project),
		Users:    make(map[string]*keystone.User),
	}
	a.Projects[AdminProjectID] = &keystone.Project{
		Domain: a.Domains[DefaultDomainID],
		ID:     AdminProjectID,
		Name:   AdminProjectName,
	}
	a.Users[AdminUserID] = &keystone.User{
		Domain:   a.Domains[DefaultDomainID],
		ID:       AdminUserID,
		Name:     AdminUserName,
		Password: AdminUserPassword,
		Roles: []*keystone.Role{
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
		viper.Set(k, v)
	}
}

func configureDebugLogging(t *testing.T) {
	err := pkglog.Configure("debug")
	assert.NoError(t, err, "configuring logging failed")
}

// URL returns server base URL.
func (s *APIServer) URL() string {
	return s.testServer.URL
}

// Close closes server.
func (s *APIServer) Close(t *testing.T) {
	s.log.Debug("Closing test API server")
	err := s.apiServer.Close()
	assert.NoError(t, err, "closing API Server failed")

	s.testServer.Close()
}
