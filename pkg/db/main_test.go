package db

import (
	"database/sql"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/common"
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
	for _, iConfig := range viper.GetStringMap("test_database") {
		config := common.InterfaceToInterfaceMap(iConfig)
		testDB, cErr := OpenConnection(ConnectionConfig{
			Driver:   config["type"].(string),
			User:     config["user"].(string),
			Password: config["password"].(string),
			Host:     config["host"].(string),
			Name:     config["name"].(string),
		})
		if cErr != nil {
			log.Fatal(cErr)
		}
		defer closeDB(testDB)

		db = &Service{
			db:      testDB,
			Dialect: NewDialect(config["dialect"].(string)),
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
