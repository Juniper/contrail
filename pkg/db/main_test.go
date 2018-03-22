package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	//Import mysql driver
	_ "github.com/go-sql-driver/mysql"
	//Import psql driver
	_ "github.com/lib/pq"
)

var testDB *sql.DB
var db *DB

func TestMain(m *testing.M) {

	viper.SetConfigName("server")
	viper.AddConfigPath("../apisrv")
	viper.ReadInConfig()

	common.SetLogLevel()
	var err error
	dbConfig := viper.GetStringMap("test_database")
	for _, iConfig := range dbConfig {
		config := common.InterfaceToInterfaceMap(iConfig)
		testDB, err = sql.Open(config["type"].(string), config["connection"].(string))
		if err != nil {
			log.Fatal(err)
		}
		defer testDB.Close()
		db = &DB{
			DB:      testDB,
			Dialect: NewDialect(config["dialect"].(string)),
		}
		db.initQueryBuilders()

		log.Info("starting test")
		code := m.Run()
		log.Info("finished test")
		if code != 0 {
			os.Exit(code)
		}
	}
}
