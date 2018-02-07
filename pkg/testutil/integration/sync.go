package integration

import (
	"testing"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/sync"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Sync is embedded Sync service for testing purposes.
type Sync struct {
	service  *sync.Service
	runError chan error
	log      *logrus.Entry
}

// NewRunningSync creates new running test Sync service.
// Call Close() method to release its resources.
func NewRunningSync(t *testing.T) *Sync {
	setViperConfig(map[string]interface{}{
		"database.debug": false,
		"etcd.endpoints": []string{etcdEndpoint},
	})

	s, err := sync.NewService()
	require.NoError(t, err, "creating test Sync service failed")

	runError := make(chan error)
	go func() {
		runError <- s.Run()
	}()

	return &Sync{
		service:  s,
		runError: runError,
		log:      pkglog.NewLogger("server"),
	}
}

func (w *Sync) Close(t *testing.T) {
	w.log.Debug("Closing test Sync service")
	w.service.Close()

	assert.NoError(t, <-w.runError, "Unexpected Sync runtime error")
}
