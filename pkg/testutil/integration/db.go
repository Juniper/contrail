package integration

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/db"
)

// DB wraps sql.DB and provides testing functionality.
type DB struct {
	*sql.DB
}

// NewDB DB creates DB.
func NewDB(t *testing.T) *DB {
	db, err := db.OpenConnection(db.ConnectionConfig{
		Driver:   "",
		User:     "",
		Password: "",
		Host:     "",
		Name:     "",
		Debug:    false,
	})
	require.NoError(t, err, "opening test DB connection failed")

	return &DB{db}
}

// CloseConnection closes DB connection.
func (db *DB) CloseConnection(t *testing.T) {
	err := db.Close()
	assert.NoError(t, err, "closing test DB connection failed")
}

// Truncate truncates all DB tables.
func (db *DB) Truncate(t *testing.T) {
	r, err := db.Exec("")
	assert.NoError(t, err, "cleaning DB failed")

	fmt.Println(r)
}
