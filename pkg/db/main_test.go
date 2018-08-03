package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db/basedb"
)

var db *Service

func TestMain(m *testing.M) {
	viper.SetConfigName("contrail")
	viper.AddConfigPath("../apisrv")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	common.SetLogLevel()
	dbConfig := viper.GetStringMap("test_database")
	for _, iConfig := range dbConfig {
		config := common.InterfaceToInterfaceMap(iConfig)
		driver := config["type"].(string)

		var dbDSNFormat string
		switch driver {
		case DriverPostgreSQL:
			dbDSNFormat = dbDSNFormatPostgreSQL
		case DriverMySQL:
			dbDSNFormat = dbDSNFormatMySQL
		}
		dsn := fmt.Sprintf(
			dbDSNFormat,
			config["user"].(string),
			config["password"].(string),
			config["host"].(string),
			config["name"].(string),
		)

		testDB, err := makeConnection(driver, dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer closeDB(testDB)
		db = &Service{
			BaseDB: basedb.NewBaseDB(testDB, config["dialect"].(string)),
		}
		db.initQueryBuilders()

		log.Info("Running test for " + driver)
		code := m.Run()
		log.Info("finished")
		if code != 0 {
			os.Exit(code)
		}
	}
}

func closeDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.WithError(err).Fatal("Closing test DB failed")
	}
}
