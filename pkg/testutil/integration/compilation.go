package integration

import (
	"context"
	"testing"

	"github.com/Juniper/asf/pkg/etcd"
	"github.com/Juniper/contrail/pkg/compilation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	integrationetcd "github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

const (
	messageIndex    = "MsgIndex"
	pluginDirectory = "etc/plugins/"
	readLock        = "MsgReadLock"
)

// RunIntentCompilationService runs Intent Compilation process and returns function closing it.
func RunIntentCompilationService(t *testing.T, apiURL string) context.CancelFunc {
	setViper(map[string]interface{}{
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
		etcd.ETCDEndpointsVK:    integrationetcd.Endpoint,
		etcd.ETCDGRPCInsecureVK: true,
		etcd.ETCDPathVK:         integrationetcd.Prefix,
		"client.id":             AdminUserID,
		"client.password":       AdminUserPassword,
		"client.project_id":     AdminProjectID,
		"client.domain_id":      DefaultDomainID,
		"client.schema_root":    "/public",
		"client.endpoint":       apiURL,
		"insecure":              true,
	})

	ics, err := compilation.NewIntentCompilationService()
	require.NoError(t, err, "creating Intent Compilation service failed")

	ctx, cancel := context.WithCancel(context.Background())

	errChan := make(chan error)
	go func() {
		errChan <- ics.Run(ctx)
	}()

	return func() {
		cancel()
		assert.NoError(t, <-errChan, "unexpected Intent Compilation runtime error")
	}
}
