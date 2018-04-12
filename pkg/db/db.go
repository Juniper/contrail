package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ExpansiveWorlds/instrumentedsql"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Replication drivers
const (
	DriverMySQL      = "mysql"
	DriverPostgreSQL = "postgres"
)

const (
	dbDSNFormatMySQL      = "%s:%s@tcp(%s:3306)/%s"
	dbDSNFormatPostgreSQL = "sslmode=disable user=%s password=%s host=%s dbname=%s"
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
	driver := viper.GetString("database.type")
	maxConn := viper.GetInt("database.max_open_conn")

	var dbDSNFormat string
	switch driver {
	case DriverPostgreSQL:
		dbDSNFormat = dbDSNFormatPostgreSQL
	case DriverMySQL:
		dbDSNFormat = dbDSNFormatMySQL
	default:
		return nil, errors.New("undefined database type")
	}

	dsn := fmt.Sprintf(
		dbDSNFormat,
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.name"),
	)
	db, err := makeConnection(driver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db connection")
	}
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxConn)

	retries, period := viper.GetInt("database.connection_retries"), viper.GetDuration("database.retry_period")
	for i := 0; i < retries; i++ {
		err = db.Ping()
		if err == nil {
			log.Info("connected to the database")
			return db, nil
		}
		time.Sleep(period)
		log.Printf("Retrying db connection... (%s)", err)
	}
	return nil, fmt.Errorf("failed to open db connection")
}
