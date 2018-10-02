package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

const (
	maxHistory = 100000
)

// RunCacheDB runs DB Cache with etcd event producer.
func RunCacheDB(t *testing.T) (cacheDB *cache.DB, cancelEtcdEventProducer context.CancelFunc) {
	setViperConfig(map[string]interface{}{
		"cache.timeout":  "10s",
		"etcd.endpoints": []string{integrationetcd.Endpoint},
	})

	cacheDB = cache.NewDB(maxHistory)

	processor, err := etcd.NewEventProducer(cacheDB)
	require.NoError(t, err, "creating etcd event producer for Cache DB failed")

	ctx, cancelEtcdEventProducer := context.WithCancel(context.Background())
	errChan := make(chan error)
	go func() {
		errChan <- processor.Start(ctx)
	}()

	return cacheDB, func() {
		cancelEtcdEventProducer()
		assert.NoError(t, <-errChan, "unexpected etcd event producer runtime error")
	}
}
