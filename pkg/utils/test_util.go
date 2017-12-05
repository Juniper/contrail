package utils

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
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

func ensureDB(db *sql.DB, dbName string) error {
	_, err := db.Exec("drop database if exists " + dbName)
	if err != nil {
		return errors.Wrap(err, "drop db if exists failed")
	}
	_, err = db.Exec("create database " + dbName)
	if err != nil {
		return errors.Wrap(err, "create database failed")
	}
	return nil
}

func initDB(db *sql.DB, initSQL []string) error {
	for _, query := range initSQL {
		_, err := db.Exec(query)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("init DB failed for %s", query))
		}
	}

	return nil
}

func dropDB(db *sql.DB, dbName string) error {
	_, err := db.Exec("drop database " + dbName)
	if err != nil {
		return errors.Wrap(err, "drop database failed")
	}
	return nil
}

//NewTestServer makes new test server based on test id.
//In a test process, we will create database and initialize it.
//You should call close() method to destroy test environment.
func NewTestServer(testID string, initSQL []string) (*TestServer, error) {
	err := InitConfig()
	if err != nil {
		return nil, err
	}
	databaseConnection := viper.GetString("database.connection")
	connectionParts := strings.Split(databaseConnection, "/")
	db, err := connectDB(databaseConnection)
	if err != nil {
		return nil, err
	}
	err = ensureDB(db, testID)
	if err != nil {
		return nil, err
	}
	connectionString := connectionParts[0] + "/" + testID
	db, err = connectDB(connectionString)
	if err != nil {
		return nil, err
	}
	err = initDB(db, initSQL)
	if err != nil {
		return nil, err
	}
	return &TestServer{
		DB:     db,
		DBName: testID,
	}, err
}

//Close stops test server and clean env.
func (s *TestServer) Close() {
	defer s.DB.Close()
	dropDB(s.DB, s.DBName)
}
