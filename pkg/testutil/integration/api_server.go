package integration

import (
	"database/sql"
	"fmt"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
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

	aliceUserID        = "alice"
	alicePassword      = "alice_password"
	adminProjectID     = "admin"
	adminProjectName   = "admin"
	defaultDomainID    = "default"
	authEndpointSuffix = "/keystone/v3"
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
	configureDebugLogging(t)
	setViperConfig(map[string]interface{}{
		"database.connection":      fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName),
		"database.type":            "postgres",
		"database.dialect":         "postgres",
		"keystone.local":           true,
		"keystone.assignment.type": "static",
		"keystone.assignment.file": path.Join(repoRootPath, "sample/keystone.yml"),
		"keystone.store.type":      "memory",
		"keystone.store.expire":    3600,
		"keystone.insecure":        true,
		"log_level":                "debug",
		"server.read_timeout":      10,
		"server.write_timeout":     5,
		"static_files.public":      path.Join(repoRootPath, "public"),
		"tls.enabled":              false,
	})

	s, err := apisrv.NewServer()
	require.NoError(t, err, "creating API Server failed")

	ts := httptest.NewServer(s.Echo)

	viper.Set("keystone.authurl", ts.URL+"/v3")
	err = s.Init()
	require.NoError(t, err, "initialization of test API Server failed")

	return &APIServer{
		apiServer:  s,
		testServer: ts,
		log:        pkglog.NewLogger("server"),
	}
}

func configureDebugLogging(t *testing.T) {
	err := pkglog.Configure("debug")
	require.NoError(t, err, "configuring logging failed")
}

func setViperConfig(config map[string]interface{}) {
	for k, v := range config {
		viper.Set(k, v)
	}
}

// URL returns server base URL.
func (s *APIServer) URL() string {
	return s.testServer.URL
}

// Database returns server database handle.
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
