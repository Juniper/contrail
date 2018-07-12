package integration

import (
	"context"
	"testing"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/sync"
)

const pgQueryCanceledErrorCode = "57014"

// RunSyncService runs Sync process and returns function closing it.
func RunSyncService(t *testing.T) (closeSync func()) {
	setViperConfig(map[string]interface{}{
		"database.debug": false,
		"etcd.endpoints": []string{EtcdEndpoint},
	})

	s, err := sync.NewService()
	require.NoError(t, err, "creating Sync service failed")

	runError := make(chan error)
	go func() {
		runError <- s.Run()
	}()

	return func() {
		s.Close()
		err := <-runError
		if !isContextCancellationError(err) {
			assert.NoError(t, err, "unexpected Sync runtime error")
		}
	}
}

func isContextCancellationError(err error) bool {
	if pqErr, ok := errors.Cause(err).(*pq.Error); ok {
		if pqErr.Code == pgQueryCanceledErrorCode {
			return true
		}
	}
	if errors.Cause(err) == context.Canceled {
		return true
	}
	return false
}
