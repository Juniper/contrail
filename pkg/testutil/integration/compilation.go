package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/compilation"
)

const (
	messageIndex    = "MsgIndex"
	pluginDirectory = "etc/plugins/"
	readLock        = "MsgReadLock"
)

// RunIntentCompilationService runs Intent Compilation process and returns function closing it.
func RunIntentCompilationService(t *testing.T) (cancel context.CancelFunc) {
	setViperConfig(map[string]interface{}{
		"compilation.plugin_directory":    pluginDirectory,
		"compilation.number_of_workers":   4,
		"compilation.max_job_queue_len":   5,
		"compilation.msg_queue_lock_time": 30,
		"compilation.msg_index_string":    messageIndex,
		"compilation.read_lock_string":    readLock,
		"compilation.master_election":     true,
		"compilation.plugin": map[string]map[string]string{
			"handlers": {
				"create_handler": "HandleCreate",
				"update_handler": "HandleUpdate",
				"delete_handler": "HandleDelete",
			},
		},
		"etcd.endpoints":     EtcdEndpoint,
		"etcd.grpc_insecure": true,
		"etcd.path":          EtcdJSONPrefix,
	})

	ics, err := compilation.NewIntentCompilationService()
	require.NoError(t, err, "creating Intent Compilation service failed")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err = ics.Run(ctx)
		assert.NoError(t, err, "unexpected Intent Compilation runtime error")
	}()

	return cancel
}
