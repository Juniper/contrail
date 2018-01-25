package integration

import (
	"database/sql"

	"github.com/Juniper/contrail/pkg/common"
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
