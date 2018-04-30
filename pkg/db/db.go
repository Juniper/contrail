package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ExpansiveWorlds/instrumentedsql"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/go-sql-driver/mysql"
	"github.com/gogo/protobuf/proto"
	"github.com/lib/pq"
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
func NewService(db *sql.DB, dialect string) *DB {
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

// Object is generic database model instance.
type Object = proto.Message

// ObjectWriter processes rows
type ObjectWriter interface {
	WriteObject(schemaID, objUUID string, obj Object) error
}

// Dump selects all data from every table and writes each row to ObjectWriter.
//
// Note that dumping the whole database using SELECT statements may take a lot
// of time and memory, increasing both server and database load thus it should
// be used as a first shot operation only.
//
// An example application of that function is loading initial database snapshot
// in Watcher.
func (db *DB) Dump(ctx context.Context, ow ObjectWriter) error {
	return DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.dump(ctx, ow)
		},
	)
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
		err = handleError(err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		err = handleError(err)
		return err
	}
	return nil
}

//LogQuery log sql query
// nolint
func LogQuery(ctx context.Context, command string, args ...interface{}) {
	log.Debug(command, args)
}

func makeConnection(dbType, databaseConnection string) (*sql.DB, error) {
	if viper.GetBool("database.debug") {
		logger := instrumentedsql.LoggerFunc(LogQuery)
		switch dbType {
		case MYSQL:
			dbType = "instrumented-" + dbType
			sql.Register(dbType, instrumentedsql.WrapDriver(&mysql.MySQLDriver{}, instrumentedsql.WithLogger(logger)))
		case POSTGRES:
			dbType = "instrumented-" + dbType
			sql.Register(dbType, instrumentedsql.WrapDriver(&pq.Driver{}, instrumentedsql.WithLogger(logger)))
		}
	}

	return sql.Open(dbType, databaseConnection)
}

//ConnectDB connect to the db based on viper configuration.
func ConnectDB() (*sql.DB, error) {
	dbType := viper.GetString("database.type")
	databaseConnection := viper.GetString("database.connection")
	maxConn := viper.GetInt("database.max_open_conn")
	db, err := makeConnection(dbType, databaseConnection)
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
