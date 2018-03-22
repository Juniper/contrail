package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//DB service struct
type DB struct {
	serviceif.BaseService
	DB            *sql.DB
	Dialect       Dialect
	queryBuilders map[string]*QueryBuilder
}

//NewService makes a DB service.
func NewService(db *sql.DB, dialect string) serviceif.Service {
	dbService := &DB{
		BaseService: serviceif.BaseService{},
		DB:          db,
		Dialect:     NewDialect(dialect),
	}
	dbService.initQueryBuilders()
	return dbService
}

//SetDB sets db object.
func (db *DB) SetDB(sqlDB *sql.DB) {
	db.DB = sqlDB
}

//Transaction is a context key for tx object.
var Transaction interface{} = "transaction"

const (
	retryDB     = 10
	retryDBWait = 10
)

//GetTransaction get a transaction from context.
func GetTransaction(ctx context.Context) *sql.Tx {
	iTx := ctx.Value(Transaction)
	tx, _ := iTx.(*sql.Tx)
	return tx
}

//DoInTransaction run a function inside of DB transaction
func DoInTransaction(ctx context.Context, db *sql.DB, do func(context.Context) error) error {
	var err error
	tx := GetTransaction(ctx)
	if tx != nil {
		return do(ctx)
	}
	conn, err := db.Conn(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}
	tx, err = conn.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
		conn.Close()
	}()
	newCTX := context.WithValue(ctx, Transaction, tx)
	err = do(newCTX)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}
	return nil
}

//ConnectDB connect to the db based on viper configuration.
func ConnectDB() (*sql.DB, error) {
	databaseConnection := viper.GetString("database.connection")
	maxConn := viper.GetInt("database.max_open_conn")
	db, err := sql.Open(viper.GetString("database.type"), databaseConnection)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db connection")
	}
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxConn)
	for i := 0; i < retryDB; i++ {
		err = db.Ping()
		if err == nil {
			log.Info("connected to the database")
			return db, nil
		}
		time.Sleep(retryDBWait * time.Second)
		log.Printf("Retrying db connection... (%s)", err)
	}
	return nil, fmt.Errorf("failed to open db connection")
}
