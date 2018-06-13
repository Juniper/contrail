package integration

import (
	"database/sql"
	"fmt"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/db"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	dbUser     = "root"
	dbPassword = "contrail123"
	dbName     = "contrail_test"

	authEndpointSuffix = "/keystone/v3"
	defaultDomainID    = "default"
	defaultDomainName  = "DefaultDomain"
	adminProjectID     = "admin"
	adminProjectName   = "AdminProject"
	adminRoleID        = "admin"
	adminRoleName      = "AdminRole"
	adminUserID        = "admin"
	adminUserName      = "adminUser"
	adminUserPassword  = "adminPassword"
)

// APIServer is embedded API Server for testing purposes.
type APIServer struct {
	apiServer  *apisrv.Server
	testServer *httptest.Server
	log        *logrus.Entry
}

// NewRunningAPIServer creates new running test API Server.
// Call Close() method to release its resources.
func NewRunningAPIServer(t *testing.T, repoRootPath string) *APIServer {
	setViperConfig(map[string]interface{}{
		"database.type":               db.DriverPostgreSQL,
		"database.host":               "localhost",
		"database.user":               dbUser,
		"database.name":               dbName,
		"database.password":           dbPassword,
		"database.dialect":            db.DriverPostgreSQL,
		"database.max_open_conn":      100,
		"database.connection_retries": 10,
		"database.retry_period":       3,
		"database.debug":              true,
		"keystone.local":              true,
		"keystone.assignment.type":    "static",
		"keystone.assignment.data":    keystoneAssignment(),
		"keystone.store.type":         "memory",
		"keystone.store.expire":       3600,
		"keystone.insecure":           true,
		"server.read_timeout":         10,
		"server.write_timeout":        5,
		"server.log_api":              true,
		"static_files.public":         path.Join(repoRootPath, "public"),
		"tls.enabled":                 false,
		"log_level":                   "debug",
	})
	configureDebugLogging(t)

	log := pkglog.NewLogger("api-server")
	log.WithField("config", fmt.Sprintf("%+v", viper.AllSettings())).Debug("Creating API Server")
	s, err := apisrv.NewServer()
	require.NoError(t, err, "creating API Server failed")

	ts := httptest.NewServer(s.Echo)

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
			defaultDomainID: {
				ID:   defaultDomainID,
				Name: defaultDomainName,
			},
		},
		Projects: make(map[string]*keystone.Project),
		Users:    make(map[string]*keystone.User),
	}
	a.Projects[adminProjectID] = &keystone.Project{
		Domain: a.Domains[defaultDomainID],
		ID:     adminProjectID,
		Name:   adminProjectName,
	}
	a.Users[adminUserID] = &keystone.User{
		Domain:   a.Domains[defaultDomainID],
		ID:       adminUserID,
		Name:     adminUserName,
		Password: adminUserPassword,
		Roles: []*keystone.Role{
			{
				ID:      adminRoleID,
				Name:    adminRoleName,
				Project: a.Projects[adminProjectID],
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

// Database returns database handle.
func (s *APIServer) Database() *sql.DB {
	return s.apiServer.DB()
}

// Close closes server.
func (s *APIServer) Close(t *testing.T) {
	s.log.Debug("Closing test API server")
	err := s.apiServer.Close()
	assert.NoError(t, err, "closing API Server failed")

	s.testServer.Close()
}
