package integration

import (
	"context"

	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

const (
	maxHistory = 100000
)

// RunCacheDB runs DB Cache with etcd event producer.
func RunCacheDB() (*cache.DB, error, func() error) {
	setViperConfig(map[string]interface{}{
		"cache.timeout":  "10s",
		"etcd.endpoints": []string{integrationetcd.Endpoint},
	})

	cacheDB := cache.NewDB(maxHistory)

	processor, err := etcd.NewEventProducer(cacheDB)
	if err != nil {
		return nil, err, func() error { return nil }
	}

	ctx, cancelEtcdEventProducer := context.WithCancel(context.Background())
	errChan := make(chan error)
	go func() {
		errChan <- processor.Start(ctx)
	}()

	return cacheDB, nil, func() error {
		cancelEtcdEventProducer()
		return <-errChan
	}
}
