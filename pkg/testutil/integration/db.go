package integration

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pkgdb "github.com/Juniper/contrail/pkg/db"
)

// DB wraps sql.DB and provides testing functionality.
type DB struct {
	*sql.DB
	driver string
}

// NewDB DB creates DB.
func NewDB(t *testing.T, driver string) *DB {
	db, err := pkgdb.OpenConnection(pkgdb.ConnectionConfig{
		Driver:   driver,
		User:     dbUser,
		Password: dbPassword,
		Host:     dbHost,
		Name:     dbName,
		Debug:    true,
	})
	require.NoError(t, err, "opening test DB connection failed")

	return &DB{
		DB:     db,
		driver: driver,
	}
}

// CloseConnection closes DB connection.
func (db *DB) CloseConnection(t *testing.T) {
	err := db.Close()
	assert.NoError(t, err, "closing test DB connection failed")
}

// Truncate truncates all DB tables.
func (db *DB) Truncate(t *testing.T) {
	switch db.driver {
	case pkgdb.DriverPostgreSQL:
		// TODO: list all tables
		tables := []string{"metadata", "int_pool", "ipaddress_pool"}

		r, err := db.Exec(fmt.Sprintf("TRUNCATE table %v CASCADE;", strings.Join(tables, ", ")))
		assert.NoError(t, err, "truncating DB tables failed")

		fmt.Printf("hogehoge: %+v", r)

	case pkgdb.DriverMySQL:
		// TODO: implement
		require.FailNow(t, fmt.Sprintf("unsupported DB driver: %v", db.driver))
	default:
		require.FailNow(t, fmt.Sprintf("unsupported DB driver: %v", db.driver))
	}
}
