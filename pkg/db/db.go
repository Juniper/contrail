package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ExpansiveWorlds/instrumentedsql"
	"github.com/go-sql-driver/mysql"
	"github.com/gogo/protobuf/proto"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/services"
)

// Database drivers
const (
	DriverMySQL      = "mysql"
	DriverPostgreSQL = "postgres"
)

const (
	dbDSNFormatMySQL      = "%s:%s@tcp(%s:3306)/%s"
	dbDSNFormatPostgreSQL = "sslmode=disable user=%s password=%s host=%s dbname=%s"
)

//Service struct
type Service struct {
	services.BaseService
	db            *sql.DB
	Dialect       Dialect
	queryBuilders map[string]*QueryBuilder
}

//NewServiceFromConfig makes db service from viper config.
func NewServiceFromConfig() (*Service, error) {
	sqlDB, err := ConnectDB()
	if err != nil {
		return nil, errors.Wrap(err, "Init DB failed")
	}
	return NewService(sqlDB, viper.GetString("database.dialect")), nil
}

//NewService makes a DB service.
func NewService(db *sql.DB, dialect string) *Service {
	dbService := &Service{
		BaseService: services.BaseService{},
		db:          db,
		Dialect:     NewDialect(dialect),
	}
	dbService.initQueryBuilders()
	return dbService
}

//DB gets db object.
func (db *Service) DB() *sql.DB {
	return db.db
}

//Close closes db.
func (db *Service) Close() error {
	return db.db.Close()
}

//SetDB sets db object.
func (db *Service) SetDB(sqlDB *sql.DB) {
	db.db = sqlDB
}

// Object is generic database model instance.
type Object interface {
	proto.Message
	ToMap() map[string]interface{}
}

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
func (db *Service) Dump(ctx context.Context, ow ObjectWriter) error {
	return db.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			return db.dump(ctx, ow)
		},
	)
}

//Transaction is a context key for tx object.
var Transaction interface{} = "transaction"

//GetTransaction get a transaction from context.
func GetTransaction(ctx context.Context) *sql.Tx {
	iTx := ctx.Value(Transaction)
	tx, _ := iTx.(*sql.Tx)
	return tx
}

//DoInTransaction runs a function inside of DB transaction.
func (db *Service) DoInTransaction(ctx context.Context, do func(context.Context) error) error {
	tx := GetTransaction(ctx)
	if tx != nil {
		return do(ctx)
	}

	conn, err := db.DB().Conn(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve DB connection")
	}
	defer conn.Close() // nolint: errcheck

	tx, err = conn.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "failed to start DB transaction")
	}
	defer rollbackOnPanic(tx)

	err = do(context.WithValue(ctx, Transaction, tx))
	if err != nil {
		tx.Rollback() // nolint: errcheck
		return handleError(err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback() // nolint: errcheck
		return handleError(err)
	}
	return nil
}

func rollbackOnPanic(tx *sql.Tx) {
	if p := recover(); p != nil {
		err := tx.Rollback()
		if err != nil {
			panic(fmt.Sprintf("%v; also transaction rollback failed: %v", p, err))
		}
		panic(p)
	}
}

//ConnectDB connect to the db based on viper configuration.
func ConnectDB() (*sql.DB, error) {
	db, err := OpenConnection(ConnectionConfig{
		Driver:   viper.GetString("database.type"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Name:     viper.GetString("database.name"),
		Debug:    viper.GetBool("database.debug"),
	})
	if err != nil {
		return nil, err
	}

	maxConn := viper.GetInt("database.max_open_conn")
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

// ConnectionConfig holds DB connection configuration.
type ConnectionConfig struct {
	Driver   string
	User     string
	Password string
	Host     string
	Name     string
	Debug    bool
}

// OpenConnection opens DB connection.
func OpenConnection(c ConnectionConfig) (*sql.DB, error) {
	if c.Debug {
		c.Driver = wrapDriver(c.Driver)
	}

	dsn, err := dataSourceName(&c)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(c.Driver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open DB connection")
	}
	return db, nil
}

func logQuery(_ context.Context, command string, args ...interface{}) {
	log.Debug(command, args)
}

func wrapDriver(driver string) string {
	switch driver {
	case MYSQL:
		driver = "instrumented-" + driver
		sql.Register(driver, instrumentedsql.WrapDriver(
			&mysql.MySQLDriver{},
			instrumentedsql.WithLogger(instrumentedsql.LoggerFunc(logQuery))),
		)
		return driver
	case POSTGRES:
		driver = "instrumented-" + driver
		sql.Register(driver, instrumentedsql.WrapDriver(
			&pq.Driver{},
			instrumentedsql.WithLogger(instrumentedsql.LoggerFunc(logQuery))),
		)
		return driver
	}
	return driver
}

func dataSourceName(c *ConnectionConfig) (string, error) {
	f, err := getDSNFormat(c.Driver)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(f, c.User, c.Password, c.Host, c.Name), nil
}

func getDSNFormat(driver string) (string, error) {
	switch driver {
	case DriverPostgreSQL:
		return dbDSNFormatPostgreSQL, nil
	case DriverMySQL:
		return dbDSNFormatMySQL, nil
	default:
		return "", errors.Errorf("undefined database driver: %v", driver)
	}
}
