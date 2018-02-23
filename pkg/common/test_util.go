package common

import (
	"database/sql"
	"sync"

	log "github.com/sirupsen/logrus"
)

var dbMutex *sync.Mutex
var tableMutex = map[string]*sync.Mutex{}

func init() {
	dbMutex = new(sync.Mutex)
}

//UseTable lock and initialize a table for testing.
func UseTable(db *sql.DB, table string) *sync.Mutex {
	dbMutex.Lock()
	mutex, ok := tableMutex[table]
	if !ok {
		mutex = &sync.Mutex{}
		tableMutex[table] = new(sync.Mutex)
	}
	dbMutex.Unlock()
	mutex.Lock()
	_, err := db.Exec("delete from " + table)
	if err != nil {
		log.Fatal(err)
	}
	return mutex
}
