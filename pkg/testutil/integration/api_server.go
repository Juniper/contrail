package integration

import (
	"database/sql"
	"fmt"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/db"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

const (
	dbUser     = "root"
	dbPassword = "contrail123"
	dbName     = "contrail_test"

	aliceUserID        = "alice"
	alicePassword      = "alice_password"
	adminProjectID     = "admin"
	adminProjectName   = "admin"
	defaultDomainID    = "default"
	authEndpointSuffix = "/keystone/v3"
)

const keystoneAssignmentYAML = `
domains:
  default: &default
    id: default
    name: default
projects:
  admin: &admin
    id: admin
    name: admin
    domain: *default
  demo: &demo
    id: demo
    name: demo
    domain: *default
users:
  alice:
    id: alice
    name: Alice
    domain: *default
    password: alice_password
    email: alice@juniper.nets
    roles:
    - id: admin
      name: Admin
      project: *admin
  bob:
    id: bob
    name: Bob
    domain: *default
    password: bob_password
    email: bob@juniper.net
    roles:
    - id: Member
      name: Member
      project: *demo
`

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
		"keystone.assignment.data":    keystoneAssignment(t),
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

func keystoneAssignment(t *testing.T) interface{} {
	var keystoneAssignment interface{}
	err := yaml.Unmarshal([]byte(keystoneAssignmentYAML), &keystoneAssignment)
	require.NoError(t, err)

	return keystoneAssignment
}

//func keystoneAssignment2() *keystone.StaticAssignment {
//	assignment := s.Keystone.Assignment.(*keystone.StaticAssignment)
//	assignment.Projects[testID] = &keystone.Project{
//		Domain: assignment.Domains["default"],
//		ID:     testID,
//		Name:   testID,
//	}
//
//	assignment.Users[testID] = &keystone.User{
//		Domain:   assignment.Domains["default"],
//		ID:       testID,
//		Name:     testID,
//		Password: testID,
//		Roles: []*keystone.Role{
//			{
//				ID:      "member",
//				Name:    "Member",
//				Project: assignment.Projects[testID],
//			},
//		},
//	}
//}

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
	return s.apiServer.DB
}

// Close closes server.
func (s *APIServer) Close(t *testing.T) {
	s.log.Debug("Closing test API server")
	err := s.apiServer.Close()
	assert.NoError(t, err, "closing API Server failed")

	s.testServer.Close()
}
