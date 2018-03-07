package testutil

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
	apiServerEndpoint = "localhost:9091"
	staticFilesPath   = "../../public"
)

// APIServer is embedded API Server for testing purposes.
type APIServer struct {
	apiServer  *apisrv.Server
	testServer *httptest.Server
	log        *logrus.Entry
}

// NewAPIServer creates new test API Server.
func NewAPIServer(t *testing.T, repoRootPath string) *APIServer {
	err := pkglog.Configure("debug")
	require.NoError(t, err, "cannot configure logger")

	setServerConfig(map[string]interface{}{
		"address":                  apiServerEndpoint,
		"database.connection":      fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DBUser, DBPassword, DBHostname, DBPort, DBName),
		"database.max_open_conn":   100,
		"keystone.local":           true,
		"keystone.assignment.type": "static",
		"keystone.assignment.file": path.Join(repoRootPath, "tools/keystone.yml"),
		"keystone.store.type":      "memory",
		"keystone.store.expire":    3600,
		"log_level":                "debug",
		"server.read_timeout":      10,
		"server.write_timeout":     5,
		"static_files.public":      staticFilesPath,
		"tls.enabled":              false,
	})

	s, err := apisrv.NewServer()
	require.NoError(t, err, "creating API Server failed")

	ts := httptest.NewServer(s.Echo)

	viper.Set("keystone.authurl", ts.URL+"/v3")
	err = s.Init()
	require.NoError(t, err, "server initialization failed")

	return &APIServer{
		apiServer:  s,
		testServer: ts,
		log:        pkglog.NewLogger("server"),
	}
}

func setServerConfig(config map[string]interface{}) {
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
	s.log.Debug("Closing test server")
	err := s.apiServer.Close()
	assert.NoError(t, err, "closing API Server failed")

	s.testServer.Close()
}
