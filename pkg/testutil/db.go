package testutil

import (
	"database/sql"

	"github.com/Juniper/contrail/pkg/common"
)

// Database connection constants
const (
	DBHostname = "localhost"
	DBPort     = 3306
	DBUser     = "root"
	DBPassword = "contrail123"
	DBName     = "contrail_test"
)

// LockAndClearTables locks and initializes given database tables.
func LockAndClearTables(db *sql.DB, tables ...string) {
	for _, t := range tables {
		common.UseTable(db, t)
	}
}

// ClearAndUnlockTables clears and unlocks given database tables.
func ClearAndUnlockTables(db *sql.DB, tables ...string) {
	for _, t := range tables {
		common.ClearTable(db, t)
	}
}
