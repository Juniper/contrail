package testutils

import (
	"database/sql"
	"fmt"
	"github.com/Juniper/contrail/pkg/db"
	"os"
	"strings"
	"testing"

	"github.com/ExpansiveWorlds/instrumentedsql"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	dbDSNFormatMySQL      = "%s:%s@tcp(%s:3306)/%s"
	dbDSNFormatPostgreSQL = "sslmode=disable user=%s password=%s host=%s dbname=%s"
)

var testDB *sql.DB
var TestDbService *db.Service

func makeConnection(dbType, databaseConnection string) (*sql.DB, error) {
	if viper.GetBool("database.debug") {
		logger := instrumentedsql.LoggerFunc(db.LogQuery)
		switch dbType {
		case db.MYSQL:
			dbType = "instrumented-" + dbType
			sql.Register(dbType, instrumentedsql.WrapDriver(&mysql.MySQLDriver{}, instrumentedsql.WithLogger(logger)))
		case db.POSTGRES:
			dbType = "instrumented-" + dbType
			sql.Register(dbType, instrumentedsql.WrapDriver(&pq.Driver{}, instrumentedsql.WithLogger(logger)))
		}
	}

	return sql.Open(dbType, databaseConnection)
}

//CreateTestDbService Create TestDB and DBService for Unit tests
func CreateTestDbService(m *testing.M) {
	viper.SetConfigName("contrail")
	viper.AddConfigPath("../apisrv")
	viper.ReadInConfig()
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	common.SetLogLevel()
	var err error
	dbConfig := viper.GetStringMap("test_database")

	var code int
	for _, iConfig := range dbConfig {
		config := common.InterfaceToInterfaceMap(iConfig)
		driver := config["type"].(string)

		var dbDSNFormat string
		switch driver {
		case db.DriverPostgreSQL:
			dbDSNFormat = dbDSNFormatPostgreSQL
		case db.DriverMySQL:
			dbDSNFormat = dbDSNFormatMySQL
		}
		dsn := fmt.Sprintf(
			dbDSNFormat,
			config["user"].(string),
			config["password"].(string),
			config["host"].(string),
			config["name"].(string),
		)

		testDB, err = makeConnection(driver, dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer testDB.Close()
		TestDbService = db.NewService(testDB, dbDSNFormat)

		log.Info("Running test for " + driver)
		code = m.Run()
		log.Info("finished")
		//if code != 0 {
		//	os.Exit(code)
		//}
	}
	os.Exit(code)
}
