package compilation

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/compilation/watch"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	routinesCount          = 1
	testMessageIndexString = "test-message-index-string"
)

func setTestConfig() {
	viper.Set("etcd.endpoints", "localhost:2379")
	viper.Set("etcd.path", "contrail")
	viper.Set("compilation.msg_index_string", testMessageIndexString)
}

func TestIntentCompilationSeviceConcurrency(t *testing.T) {
	etcdClient := integration.NewEtcdClient(t)
	_, err := etcdClient.Delete(context.Background(), testMessageIndexString)
	assert.NoError(t, err)

	setTestConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(routinesCount)
	for i := 0; i < routinesCount; i++ {
		go func() {
			ics, err := NewIntentCompilationService()
			assert.NoError(t, err)
			err = ics.Run(ctx)
			assert.NoError(t, err)
			wg.Done()
		}()
	}

	for {
		if watch.JobQueue != nil {
			break
		}
		time.Sleep(1 * time.Millisecond)
	}
	close(watch.JobQueue)

	_, err = etcdClient.Put(context.Background(), "/contrail/test", "value")
	_, err = etcdClient.Put(context.Background(), "/contrail/test", "valueu")

	wg.Wait()
}
