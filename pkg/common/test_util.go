package common

import (
	"database/sql"
	"sync"

	log "github.com/sirupsen/logrus"
)

var dbMutex = &sync.Mutex{}
var tableMutex = map[string]*sync.Mutex{}

//UseTable lock and initialize a table for testing.
func UseTable(db *sql.DB, table string) {
	dbMutex.Lock()
	mutex, ok := tableMutex[table]
	if !ok {
		mutex = &sync.Mutex{}
		tableMutex[table] = mutex
	}
	dbMutex.Unlock()
	mutex.Lock()
	_, err := db.Exec("delete from " + table)
	if err != nil {
		log.Fatal(err)
	}
}

//ClearTable clean and unlock the table.
func ClearTable(db *sql.DB, table string) {
	dbMutex.Lock()
	mutex, ok := tableMutex[table]
	dbMutex.Unlock()
	if !ok {
		return
	}
	_, err := db.Exec("delete from " + table)
	if err != nil {
		log.Fatal(err)
	}
	mutex.Unlock()
}
