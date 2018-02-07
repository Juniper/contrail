package integration

import (
	"testing"

	"github.com/Juniper/contrail/pkg/sync"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RunSync runs Sync process and returns function closing Sync.
func RunSync(t *testing.T) (closeSync func(t *testing.T)) {
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

	return func(t *testing.T) {
		s.Close()
		assert.NoError(t, <-runError, "unexpected Sync runtime error")
	}
}
