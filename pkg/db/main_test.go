package db

import (
	"database/sql"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var testDB *sql.DB
var db *Service

func TestMain(m *testing.M) {
	viper.SetConfigName("server")
	viper.AddConfigPath("../apisrv")
	viper.ReadInConfig()
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	common.SetLogLevel()
	var err error
	dbConfig := viper.GetStringMap("test_database")
	for _, iConfig := range dbConfig {
		config := common.InterfaceToInterfaceMap(iConfig)
		dbType := config["type"].(string)
		testDB, err = makeConnection(dbType, config["connection"].(string))
		if err != nil {
			log.Fatal(err)
		}
		defer testDB.Close()
		db = &Service{
			db:      testDB,
			Dialect: NewDialect(config["dialect"].(string)),
		}
		db.initQueryBuilders()

		log.Info("Running test for " + dbType)
		code := m.Run()
		log.Info("finished")
		if code != 0 {
			os.Exit(code)
		}
	}
}
