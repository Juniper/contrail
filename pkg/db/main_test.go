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
	common.InitConfig()
	common.SetLogLevel()
	var err error
	testDB, err = ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()
	db = &DB{
		DB:      testDB,
		Dialect: NewDialect(viper.GetString("database.dialect")),
	}
	db.initQueryBuilders()

	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	os.Exit(code)
}
