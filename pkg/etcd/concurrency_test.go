package etcd_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/Juniper/asf/pkg/db/etcd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	integrationetcd "github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

// TODO: move to ASF

const (
	testResourceKey          = "concurrent-resource-test"
	testResourceInitialValue = "test-initial-value"
)

func TestDoInTransactionEtcdDataRace(t *testing.T) {
	e := integrationetcd.NewEtcdClient(t)
	defer e.Close(t)

	client, err := etcd.NewClient(&etcd.Config{
		Client:      e.Client,
		ServiceName: t.Name(),
	})
	require.NoError(t, err)

	collectKeyHistory := e.WatchKey(testResourceKey)

	// Set initial state
	err = client.Put(context.Background(), testResourceKey, []byte(testResourceInitialValue))
	assert.NoError(t, err)

	firstValue, interrupterValue := "first-value", "interrupter-value"
	var firstProcessedKeys, interrupterProcessedKeys []string

	// Interrupter is etcd user who updated key during first process operations.
	interrupterProcess := func() error {
		var value []byte

		err = client.DoInTransaction(context.Background(), func(ctx context.Context) error {
			txn := etcd.GetTxn(ctx)
			value = txn.Get(testResourceKey)

			txn.Put(testResourceKey, []byte(interrupterValue))

			return nil
		})

		// Collect all keys that interrupter "processed".
		interrupterProcessedKeys = append(interrupterProcessedKeys, string(value))
		return err
	}

	// Allow running interrupter once
	var once sync.Once
	runInterrupterOnce := func() {
		once.Do(func() {
			err = interrupterProcess()
			assert.NoError(t, err)
		})
	}

	// First process is etcd user who got interrupted by interrupter before updating value.
	firstProcess := func() error {
		var value []byte
		err = client.DoInTransaction(context.Background(), func(ctx context.Context) error {
			txn := etcd.GetTxn(ctx)

			value = txn.Get(testResourceKey)
			fmt.Println(string(value))

			runInterrupterOnce() // Run interrupter before saving

			txn.Put(testResourceKey, []byte(firstValue))

			return nil
		})

		// Collect all keys that first "processed".
		firstProcessedKeys = append(firstProcessedKeys, string(value))

		return err
	}

	err = firstProcess()
	assert.NoError(t, err)

	assert.Equal(t, []string{testResourceInitialValue}, interrupterProcessedKeys,
		"interrupter should process only initial value")
	assert.Equal(t, []string{interrupterValue}, firstProcessedKeys, "first should only process interrupter value")

	err = e.Client.Sync(context.Background())
	assert.NoError(t, err)

	keyHistory := collectKeyHistory()
	assert.Equal(t, []string{testResourceInitialValue, interrupterValue, firstValue}, keyHistory)

	// cleanup
	err = client.Delete(context.Background(), testResourceKey)
	assert.NoError(t, err)
}
