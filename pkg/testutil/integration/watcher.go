package integration

import (
	"testing"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/watcher"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// Watcher is embedded Watcher service for testing purposes.
type Watcher struct {
	watcherService *watcher.Service
	runError       chan error
	log            *logrus.Entry
}

// NewRunningWatcher creates new running test Watcher service.
// Call Close() method to release its resources.
func NewRunningWatcher(t *testing.T) *Watcher {
	setViperConfig(map[string]interface{}{
		"database.user":                       dbUser,
		"database.password":                   dbPassword,
		"database.name":                       dbName,
		"database.connection_retries":         10,
		"database.driver":                     watcher.DriverPostgreSQL,
		"database.replication_status_timeout": "10s",
		"database.retry_period":               "1s",
		"etcd.dial_timeout":                   "60s",
		"etcd.endpoints":                      []string{etcdEndpoint},
	})

	s, err := watcher.NewService()
	if err != nil {
		t.Fatalf("creating test Watcher service failed: %s", err)
	}

	runError := make(chan error)
	go func() {
		runError <- s.Run()
	}()

	return &Watcher{
		watcherService: s,
		runError:       runError,
		log:            pkglog.NewLogger("server"),
	}
}

func (w *Watcher) Close(t *testing.T) {
	w.log.Debug("Closing test Watcher service")
	w.watcherService.Close()

	assert.NoError(t, <-w.runError, "Unexpected Watcher Run error")
}
