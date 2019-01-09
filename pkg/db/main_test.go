package db

import (
	"database/sql"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/logging"
)

var db *Service

func TestMain(m *testing.M) {
	viper.SetConfigType("yml")
	viper.SetConfigName("test_config")
	viper.AddConfigPath("../../sample")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	logging.SetLogLevel()
	for _, iConfig := range viper.GetStringMap("test_database") {
		config := format.InterfaceToInterfaceMap(iConfig)
		driver := config["type"].(string) //nolint: errcheck
		testDB, err := basedb.OpenConnection(basedb.ConnectionConfig{
			Driver:   driver,
			User:     config["user"].(string),
			Password: config["password"].(string),
			Host:     config["host"].(string),
			Port:     format.InterfaceToString(config["port"]),
			Name:     config["name"].(string),
		})
		if err != nil {
			log.Fatal(err)
		}
		defer closeDB(testDB)

		db = &Service{
			BaseDB: basedb.NewBaseDB(testDB, config["dialect"].(string)),
		}
		db.initQueryBuilders()

		log.WithField("dbType", config["type"]).Info("Starting tests for DB")
		code := m.Run()
		log.WithField("dbType", config["type"]).Info("Finished tests for DB")
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
