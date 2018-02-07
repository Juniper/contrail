package integration

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/sync"
)

// RunSync runs Sync process and returns function closing Sync.
func RunSync(t *testing.T) (closeSync func()) {
	setViperConfig(map[string]interface{}{
		"database.debug": false,
		"etcd.endpoints": []string{etcdEndpoint},
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
		// TODO(Daniel): ignore premature Sync close errors in a better way
		if !isContextCancellationError(err) {
			assert.NoError(t, err, "unexpected Sync runtime error")
		}
	}
}

func isContextCancellationError(err error) bool {
	return strings.Contains(err.Error(), "context canceled") ||
		strings.Contains(err.Error(), "canceling statement due to user request")
}
