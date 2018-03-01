package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
	//Import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	common.InitConfig()
	common.SetLogLevel()
	var err error
	testDB, err = ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()
	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	os.Exit(code)
}
