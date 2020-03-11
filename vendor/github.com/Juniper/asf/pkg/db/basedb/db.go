package basedb

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/ExpansiveWorlds/instrumentedsql"
	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/models"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Database driver
const (
	DriverPostgreSQL = "postgres"

	dbDSNFormatPostgreSQL = "sslmode=disable user=%s password=%s host=%s dbname=%s"
)

//BaseDB struct for base function.
type BaseDB struct {
	db            *sql.DB
	Dialect       Dialect
	QueryBuilders map[string]*QueryBuilder
}

//NewBaseDB makes new base db instance.
func NewBaseDB(db *sql.DB) BaseDB {
	return BaseDB{
		db:      db,
		Dialect: NewDialect(),
	}
}

//DB gets db object.
func (db *BaseDB) DB() *sql.DB {
	return db.db
}

//Close closes db.
func (db *BaseDB) Close() error {
	if err := db.db.Close(); err != nil {
		return errors.Wrap(err, "close DB handle")
	}
	return nil
}

//DoInTransaction runs a function inside of DB transaction.
func (db *BaseDB) DoInTransaction(ctx context.Context, do func(context.Context) error) error {
	return db.DoInTransactionWithOpts(ctx, do, nil)
}

//DoInTransactionWithOpts runs a function inside of DB transaction with extra options.
func (db *BaseDB) DoInTransactionWithOpts(
	ctx context.Context, do func(context.Context) error, opts *sql.TxOptions,
) error {
	tx := GetTransaction(ctx)
	if tx != nil {
		return do(ctx)
	}

	conn, err := db.DB().Conn(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve DB connection")
	}
	defer conn.Close() // nolint: errcheck

	tx, err = conn.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(FormatDBError(err), "failed to start DB transaction")
	}
	defer rollbackOnPanic(tx)

	if err = do(context.WithValue(ctx, Transaction, tx)); err != nil {
		tx.Rollback() // nolint: errcheck
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback() // nolint: errcheck
		return FormatDBError(err)
	}

	return nil
}

//Transaction is a context key for tx object.
var Transaction interface{} = "transaction"

//GetTransaction get a transaction from context.
func GetTransaction(ctx context.Context) *sql.Tx {
	tx, _ := ctx.Value(Transaction).(*sql.Tx) //nolint: errcheck
	return tx
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

// DoWithoutConstraints executes function without checking DB constraints
func (db *BaseDB) DoWithoutConstraints(ctx context.Context, do func(context.Context) error) (err error) {
	if err = db.disableConstraints(); err != nil {
		return err
	}
	defer func() {
		if enerr := db.enableConstraints(); enerr != nil {
			if err != nil {
				err = errutil.MultiError{err, enerr}
				return
			}
			err = enerr
		}
	}()
	err = do(ctx)
	return err
}

// disableConstraints globally disables constraints checking in DB - USE WITH CAUTION!
func (db *BaseDB) disableConstraints() error {
	_, err := db.DB().Exec(db.Dialect.DisableConstraints())
	return errors.Wrapf(err, "Disabling constraints checking (%s): ", db.Dialect.DisableConstraints())
}

// enableConstraints globally enables constraints checking - reverts behavior of DisableConstraints()
func (db *BaseDB) enableConstraints() error {
	_, err := db.DB().Exec(db.Dialect.EnableConstraints())
	return errors.Wrapf(err, "Enabling constraints checking (%s): ", db.Dialect.EnableConstraints())
}

// Delete deletes a resource
func (db *BaseDB) Delete(
	ctx context.Context,
	qb *QueryBuilder,
	uuid string,
	backrefFields map[string][]string,
) error {
	if err := db.CheckPolicy(ctx, qb, uuid); err != nil {
		return err
	}

	tx := GetTransaction(ctx)

	for backref := range backrefFields {
		_, err := tx.ExecContext(ctx, qb.DeleteRelaxedBackrefsQuery(backref), uuid)
		if err != nil {
			return errors.Wrapf(
				FormatDBError(err),
				"deleting all relaxed references from %s to resource with UUID '%v' from DB failed",
				backref, uuid,
			)
		}
	}

	_, err := tx.ExecContext(ctx, qb.DeleteQuery(), uuid)
	if err != nil {
		err = FormatDBError(err)
		return errors.Wrapf(err, "deleting resource with UUID '%v' from DB failed", uuid)
	}

	return db.DeleteMetadata(ctx, uuid)
}

// Count counts rows for given ListSpec.
func (db *BaseDB) Count(
	ctx context.Context, qb *QueryBuilder, spec *baseservices.ListSpec,
) (count int64, err error) {
	query, values := qb.CountQuery(auth.GetIdentity(ctx), spec)

	tx := GetTransaction(ctx)
	row := tx.QueryRowContext(ctx, query, values...)

	if err = row.Scan(&count); err != nil {
		return 0, errors.Wrap(FormatDBError(err), "count query failed")
	}

	return count, nil
}

// checkPolicy check ownership of resources.
func (db *BaseDB) CheckPolicy(ctx context.Context, qb *QueryBuilder, uuid string) (err error) {
	tx := GetTransaction(ctx)
	auth := auth.GetIdentity(ctx)

	selectQuery := qb.SelectAuthQuery(auth.IsAdmin())

	var row *sql.Row
	if auth.IsAdmin() {
		row = tx.QueryRowContext(ctx, selectQuery, uuid)
	} else {
		row = tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
	}

	var count int
	err = row.Scan(&count)
	if err != nil {
		return FormatDBError(err)
	}
	if count == 0 {
		return errutil.ErrorNotFound
	}

	return nil
}

type childObject interface {
	GetUUID() string
	GetParentUUID() string
	GetParentType() string
}

func (db *BaseDB) CreateParentReference(
	ctx context.Context,
	obj childObject,
	qb *QueryBuilder,
	possibleParents []string,
	optional bool,
) (err error) {
	parentSchemaID := models.KindToSchemaID(obj.GetParentType())

	if !format.ContainsString(possibleParents, parentSchemaID) {
		if optional {
			return nil
		}
		return errutil.ErrorBadRequest("invalid parent type")
	}

	tx := GetTransaction(ctx)
	_, err = tx.ExecContext(ctx, qb.CreateParentRefQuery(parentSchemaID), obj.GetUUID(), obj.GetParentUUID())

	return errors.Wrapf(
		FormatDBError(err), "creating resource %T with UUID '%v' in DB failed", obj, obj.GetUUID(),
	)
}

func (db *BaseDB) CreateRef(
	ctx context.Context,
	fromID, toID string,
	fromSchemaID, toSchemaID string,
	attrs ...interface{},
) error {
	qb := db.QueryBuilders[fromSchemaID]
	tx := GetTransaction(ctx)

	_, err := tx.ExecContext(ctx, qb.CreateRefQuery(toSchemaID), append([]interface{}{fromID, toID}, attrs...)...)
	if err != nil {
		return errors.Wrapf(
			FormatDBError(err),
			"%s_ref create failed for object %s with UUID: '%v' and ref UUID '%v'",
			toSchemaID, fromSchemaID, fromID, toID,
		)
	}
	return nil
}

func (db *BaseDB) DeleteRef(
	ctx context.Context,
	fromID, toID string,
	fromSchemaID, toSchemaID string,
) error {
	query := db.QueryBuilders[fromSchemaID].DeleteRefQuery(toSchemaID)
	tx := GetTransaction(ctx)

	_, err := tx.ExecContext(ctx, query, fromID, toID)

	return errors.Wrapf(
		FormatDBError(err),
		"%s_ref delete failed for object %s with UUID: '%v' and ref UUID '%v'",
		toSchemaID, fromSchemaID, fromID, toID,
	)
}

// ListRows gets rows for given schema ID.
func (db *BaseDB) ListRows(ctx context.Context, schemaID string, spec *baseservices.ListSpec) (*MapRows, error) {
	qb := db.QueryBuilders[schemaID]
	query, columns, values := qb.ListQuery(auth.GetIdentity(ctx), spec)

	tx := GetTransaction(ctx)
	rows, err := tx.QueryContext(ctx, query, values...)
	if err != nil {
		err = FormatDBError(err)
		return nil, errors.Wrap(err, "select query failed")
	}

	return &MapRows{Rows: rows, columns: columns}, nil
}

// MapRows is a sql.Rows wrapper that allows reading rows as maps.
type MapRows struct {
	*sql.Rows
	columns Columns
}

// ReadMap scans a row of values from sql.Rows and returns it in a form of RowData.
func (r *MapRows) ReadMap() (RowData, error) {
	return scanRow(r, r.columns)
}

type sqlScanner interface {
	Scan(...interface{}) error
}

func scanRow(s sqlScanner, c Columns) (RowData, error) {
	if s == nil {
		return nil, errors.New("scanner is nil")
	}
	values := makeInterfacePointerArray(len(c))
	if err := s.Scan(values...); err != nil {
		return nil, errors.Wrap(err, "scan failed")
	}

	valuesMap := make(map[string]interface{}, len(c))
	for column, index := range c {
		if index >= len(values) {
			return nil, errors.New("column index is greater than scanned values")
		}
		val := values[index].(*interface{})
		valuesMap[column] = *val
	}

	return valuesMap, nil
}

// Dump selects all data from every table and returns objects as RowData maps.
//
// Note that dumping the whole database using SELECT statements may take a lot
// of time and memory, increasing both server and database load, thus it should
// be used as a first shot operation only.
//
// An example application of that function is loading the initial database snapshot
// in Watcher.
func (db *BaseDB) Dump(ctx context.Context) (DatabaseData, error) {
	result := DatabaseData{}

	for schemaID := range db.QueryBuilders {
		table, err := db.dumpTable(ctx, schemaID)
		if err != nil {
			return nil, errors.Wrap(err, "select query failed")
		}
		result[schemaID] = table
	}

	return result, nil
}

func (db *BaseDB) dumpTable(ctx context.Context, schemaID string) (TableData, error) {
	var result TableData
	rows, err := db.ListRows(ctx, schemaID, &baseservices.ListSpec{Detail: true})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		m, err := rows.ReadMap()
		if err != nil {
			return nil, err
		}
		result = append(result, m)
	}
	return result, nil
}

// DriverWrapper is a function that wraps driver.Driver adding some functionalities.
type DriverWrapper func(driver.Driver) driver.Driver

// WithInstrumentedSQL creates DriverWrapper that add instrumentedsql logger.
func WithInstrumentedSQL() func(driver.Driver) driver.Driver {
	return func(d driver.Driver) driver.Driver {
		return instrumentedsql.WrapDriver(d, instrumentedsql.WithLogger(instrumentedsql.LoggerFunc(logQuery)))
	}
}

func logQuery(_ context.Context, command string, args ...interface{}) {
	logrus.Debug(command, args)
}

//ConnectDB connect to the db based on viper configuration.
func ConnectDB(wrappers ...DriverWrapper) (*sql.DB, error) {
	if debug := viper.GetBool("database.debug"); debug {
		wrappers = append(wrappers, WithInstrumentedSQL())
	}
	config := ConnectionConfigFromViper()
	config.DriverWrappers = wrappers

	db, err := OpenConnection(config)
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
			logrus.Debug("Connected to the database")
			return db, nil
		}
		time.Sleep(period)
		logrus.WithError(err).Debug("DB connection error. Retrying...")
	}
	return nil, fmt.Errorf("failed to open DB connection")
}

// ConnectionConfig holds DB connection configuration.
type ConnectionConfig struct {
	DriverWrappers []DriverWrapper
	User           string
	Password       string
	Host           string
	Name           string
}

// ConnectionConfigFromViper loads ConnectionConfig from viper.
func ConnectionConfigFromViper() ConnectionConfig {
	return ConnectionConfig{
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Name:     viper.GetString("database.name"),
	}
}

// OpenConnection opens DB connection.
func OpenConnection(c ConnectionConfig) (*sql.DB, error) {
	dsn, err := dataSourceName(&c)
	if err != nil {
		return nil, err
	}

	driverName := registerDriver(c.DriverWrappers)

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open DB connection")
	}
	return db, nil
}

func dataSourceName(c *ConnectionConfig) (string, error) {
	return fmt.Sprintf(dbDSNFormatPostgreSQL, c.User, c.Password, c.Host, c.Name), nil
}

func registerDriver(wrappers []DriverWrapper) string {
	driverName := "wrapped-" + DriverPostgreSQL

	if !isDriverRegistered(driverName) {
		var d driver.Driver = &pq.Driver{}
		for _, w := range wrappers {
			d = w(d)
		}
		sql.Register(driverName, d)
	}

	return driverName
}

func isDriverRegistered(driver string) bool {
	for _, d := range sql.Drivers() {
		if d == driver {
			return true
		}
	}
	return false
}
