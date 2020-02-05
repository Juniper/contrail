package integration

import (
	"github.com/Juniper/asf/pkg/etcd"
	"github.com/Juniper/contrail/pkg/models"
	integrationetcd "github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

// SetDefaultSyncConfig sets config options required by sync.
func SetDefaultSyncConfig(shouldDump bool) {
	setViper(map[string]interface{}{
		etcd.ETCDEndpointsVK:          []string{integrationetcd.Endpoint},
		"sync.storage":                models.JSONCodec.Key(),
		"sync.dump":                   shouldDump,
		"database.host":               "localhost",
		"database.user":               dbUser,
		"database.name":               dbName,
		"database.password":           dbPassword,
		"database.max_open_conn":      100,
		"database.connection_retries": 10,
		"database.retry_period":       3,
		"database.debug":              true,
	})
}
