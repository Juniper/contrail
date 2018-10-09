package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

type Runner interface {
	Run() error
}

func RunConcurrently(r Runner) <-chan error {
	runError := make(chan error)
	go func() {
		runError <- r.Run()
	}()

	return runError
}

type Closer interface {
	Close()
}

func CloseNoError(t *testing.T, c Closer, errChan <-chan error) {
	c.Close()
	assert.NoError(t, <-errChan, "unexpected error while closing")
}

type RunCloser interface {
	Runner
	Closer
}

func RunNoError(t *testing.T, rc RunCloser) (close func(*testing.T)) {
	errChan := RunConcurrently(rc)
	return func(*testing.T) { CloseNoError(t, rc, errChan) }
}

func SetDefaultSyncConfig() {
	setViperConfig(map[string]interface{}{
		"etcd.endpoints":              []string{integrationetcd.Endpoint},
		"sync.storage":                models.JSONCodec.Key(),
		"database.type":               basedb.DriverPostgreSQL,
		"database.host":               "localhost",
		"database.user":               dbUser,
		"database.name":               dbName,
		"database.password":           dbPassword,
		"database.dialect":            basedb.DriverPostgreSQL,
		"database.max_open_conn":      100,
		"database.connection_retries": 10,
		"database.retry_period":       3,
		"database.debug":              true,
	})
}
