package integration

import (
	"context"

	"github.com/Juniper/contrail/pkg/cache"
	"github.com/Juniper/contrail/pkg/etcd"

	asfetcd "github.com/Juniper/asf/pkg/etcd"
	integrationetcd "github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

const (
	maxHistory = 100000
)

// RunCacheDB runs DB Cache with etcd event producer.
func RunCacheDB() (*cache.DB, func() error, error) {
	setViper(map[string]interface{}{
		"cache.timeout":         "10s",
		asfetcd.ETCDEndpointsVK: []string{integrationetcd.Endpoint},
	})

	cacheDB := cache.NewDB(maxHistory)

	processor, err := etcd.NewEventProducer(cacheDB, "integration-cache-db")
	if err != nil {
		return nil, func() error { return nil }, err
	}

	ctx, cancelEtcdEventProducer := context.WithCancel(context.Background())
	errChan := make(chan error)
	go func() {
		errChan <- processor.Start(ctx)
	}()

	return cacheDB, func() error {
		cancelEtcdEventProducer()
		return <-errChan
	}, nil
}
