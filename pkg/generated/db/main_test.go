package db

import (
	"database/sql"
	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
	//Import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	common.InitConfig()
	common.SetLogLevel()
	var err error
	testDB, err = common.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()
	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	os.Exit(code)
}
