package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/compilation"
)

const (
	messageIndex = "MsgIndex"
	readLock     = "MsgReadLock"
)

// RunIntentCompilationService runs Intent Compilation process and returns function closing it.
func RunIntentCompilationService(t *testing.T) (closeIntentCompilation func()) {
	setViperConfig(map[string]interface{}{
		"compilation.number_of_workers":   4,
		"compilation.max_job_queue_len":   5,
		"compilation.msg_queue_lock_time": 30,
		"compilation.msg_index_string":    messageIndex,
		"compilation.read_lock_string":    readLock,
		"compilation.master_election":     false,
		"etcd.endpoints":                  EtcdEndpoint,
		"etcd.grpc_insecure":              true,
		"etcd.path":                       EtcdJSONPrefix,
	})

	ics, err := compilation.NewIntentCompilationService()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err = ics.Run(ctx)
		assert.NoError(t, err, "unexpected Intent Compilation runtime error")
	}()

	return func() {
		cancel()
	}
}
