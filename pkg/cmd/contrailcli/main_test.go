package contrailcli

import (
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../apisrv")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.Set("server.static_files.public", "../../../public")

	if server, err = integration.NewRunningServer(&integration.APIServerConfig{
		DBDriver:           viper.GetString("type"),
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: true,
	}); err != nil {
		log.Fatalf("Error initializing integration APIServer: %+v", err)
	}
	defer testutil.LogFatalIfErr(server.Close)

	os.Exit(m.Run())
}

func setupClient(testID string) {
	integration.AddKeystoneProjectAndUser(server.APIServer, testID)
	viper.SetDefault("client.id", testID)
	viper.SetDefault("client.password", testID)
	viper.SetDefault("client.project_id", testID)
	viper.SetDefault("client.domain_id", "default")
	viper.SetDefault("client.endpoint", server.TestServer.URL)
	viper.SetDefault("client.schema_root", "/public")
	viper.SetDefault("insecure", true)
}
