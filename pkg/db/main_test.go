package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
	//Import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB
var db *DB

func TestMain(m *testing.M) {
	common.InitConfig()
	common.SetLogLevel()
	var err error
	fmt.Println("connected db")
	testDB, err = ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()
	db = &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()

	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	os.Exit(code)
}
