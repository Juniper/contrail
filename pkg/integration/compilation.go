package integration

import (
	"context"
	"log"

	"github.com/Juniper/contrail/pkg/compilation"
	"github.com/Juniper/contrail/pkg/integration/etcd"
)

const (
	messageIndex    = "MsgIndex"
	pluginDirectory = "etc/plugins/"
	readLock        = "MsgReadLock"
)

// RunIntentCompilationService runs Intent Compilation process and returns function closing it.
func RunIntentCompilationService(apiURL string) context.CancelFunc {
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
		"etcd.endpoints":     integrationetcd.Endpoint,
		"etcd.grpc_insecure": true,
		"etcd.path":          integrationetcd.JSONPrefix,
		"client.id":          AdminUserID,
		"client.password":    AdminUserPassword,
		"client.project_id":  AdminProjectID,
		"client.domain_id":   DefaultDomainID,
		"client.schema_root": "/public",
		"client.endpoint":    apiURL,
		"insecure":           true,
	})

	ics, err := compilation.NewIntentCompilationService()
	if err != nil {
		log.Fatalf("creating Intent Compilation service failed: %+v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	errChan := make(chan error)
	go func() {
		errChan <- ics.Run(ctx)
	}()

	return func() {
		cancel()
		if err != nil {
			log.Fatalf("unexpected Intent Compilation runtime error: %+v", err)
		}
	}
}
