package contrailcli

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/common"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var testServer *httptest.Server
var server *apisrv.Server

func TestMain(m *testing.M) {
	viper.SetConfigName("contrail")
	viper.AddConfigPath("../../apisrv")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.Set("server.static_files.public", "../../../public")
	common.SetLogLevel()
	server, err = apisrv.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	testServer = testutil.NewTestHTTPServer(server.Echo)
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
	apisrv.AddKeystoneProjectAndUser(server, testID)
	viper.SetDefault("client.id", testID)
	viper.SetDefault("client.password", testID)
	viper.SetDefault("client.project_id", testID)
	viper.SetDefault("client.domain_id", "default")
	viper.SetDefault("client.endpoint", testServer.URL)
	viper.SetDefault("client.schema_root", "/public")
	viper.SetDefault("insecure", true)
}
