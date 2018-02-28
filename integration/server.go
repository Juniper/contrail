package integration

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Server is embedded API Server for testing purposes.
type Server struct {
	apiServer  *apisrv.Server
	testServer *httptest.Server
	log        *logrus.Entry
}

// NewServer creates new test API Server.
func NewServer(t *testing.T) *Server {
	err := pkglog.Configure("debug")
	require.NoError(t, err, "cannot configure logger")

	setServerConfig(map[string]interface{}{
		"address":                  ":9091",
		"database.connection":      "root:contrail123@tcp(localhost:3306)/contrail_test",
		"database.max_open_conn":   100,
		"database.dialect":         "mysql",
		"keystone.local":           true,
		"keystone.assignment.type": "static",
		"keystone.assignment.file": "../keystone.yml",
		"keystone.store.type":      "memory",
		"keystone.store.expire":    3600,
		"log_level":                "debug",
		"proxy./contrail":          []string{"http://localhost:8082"},
		"server.read_timeout":      10,
		"server.write_timeout":     5,
		"static_files.public":      "../../public",
		"tls.enabled":              false,
	})

	s, err := apisrv.NewServer()
	require.NoError(t, err, "creating API Server failed")

	ts := httptest.NewServer(s.Echo)

	viper.Set("keystone.authurl", ts.URL+"/v3")
	err = s.Init()
	require.NoError(t, err, "server initialization failed")

	return &Server{
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

// URL returns Server base URL.
func (s *Server) URL() string {
	return s.testServer.URL
}

// Database returns API Server database handle.
func (s *Server) Database() *sql.DB {
	return s.apiServer.DB
}

// Close closes Server.
func (s *Server) Close(t *testing.T) {
	s.log.Debug("Closing test server")
	err := s.apiServer.Close()
	assert.NoError(t, err, "closing API Server failed")

	s.testServer.Close()
}
