package contrailcli

import (
	"crypto/tls"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/common"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var testServer *httptest.Server
var server *apisrv.Server

func TestMain(m *testing.M) {
	viper.SetConfigName("server")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	common.SetLogLevel()
	server, err = apisrv.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	testServer = httptest.NewUnstartedServer(server.Echo)
	testServer.TLS = new(tls.Config)
	testServer.TLS.NextProtos = append(testServer.TLS.NextProtos, "h2")
	testServer.StartTLS()
	defer testServer.Close()

	viper.Set("keystone.authurl", testServer.URL+"/keystone/v3")
	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close() // nolint: errcheck

	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	os.Exit(code)
}

func setupClient(testID string) {
	apisrv.CreateTestProject(server, testID)
	viper.SetDefault("id", testID)
	viper.SetDefault("password", testID)
	viper.SetDefault("project_id", testID)
	viper.SetDefault("endpoint", testServer.URL)
	viper.SetDefault("auth_url", testServer.URL+"/keystone/v3")
	viper.SetDefault("insecure", true)
}
