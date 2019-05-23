package contrailcli

import (
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	// TODO(Daniel): remove that in order not to depend on Viper and use constructors' parameters instead
	viper.Set("server.static_files.public", "../../../public")

	var err error
	if server, err = integration.NewRunningServer(&integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: true,
	}); err != nil {
		logutil.FatalWithStackTrace(errors.Wrap(err, "initializing integration APIServer failed"))
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
