package contrailcli

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	viper.SetConfigType("yml")
	viper.SetConfigName("test_config")
	viper.AddConfigPath("../../../sample")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	// TODO(Daniel): remove that in order not to depend on Viper and use constructors' parameters instead
	viper.Set("server.static_files.public", "../../../public")

	if server, err = integration.NewRunningServer(&integration.APIServerConfig{
		DBDriver:           viper.GetString("type"),
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: true,
	}); err != nil {
		logrus.Fatalf("Error initializing integration APIServer: %+v", err)
	}
	defer testutil.LogFatalIfError(server.Close)

	if code := m.Run(); code != 0 {
		os.Exit(code)
	}
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
