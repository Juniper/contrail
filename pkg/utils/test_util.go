package utils

import (
	"database/sql"

	//loading mysql for testing
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//TestServer for test mode
type TestServer struct {
	DB     *sql.DB
	DBName string
}

func connectDB(databaseConnection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", databaseConnection)
	if err != nil {
		return nil, errors.Wrap(err, "connect DB failed")
	}
	return db, nil
}

//NewTestServer makes new test server based on test id.
//You should call close() method to destroy test environment.
func NewTestServer() *TestServer {
	err := InitConfig()
	log.SetLevel(log.DebugLevel)
	log.Debug("Test server started")
	databaseConnection := viper.GetString("database.connection")
	db, err := connectDB(databaseConnection)
	if err != nil {
		log.Fatal(err)
	}
	return &TestServer{
		DB: db,
	}
}

//Close stops test server and clean env.
func (s *TestServer) Close() {
	err := s.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Closing test server")
}
