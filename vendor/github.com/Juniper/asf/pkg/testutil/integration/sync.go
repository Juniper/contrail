package integration

import (
	"github.com/Juniper/asf/pkg/constants"
	"github.com/Juniper/asf/pkg/models"
	integrationetcd "github.com/Juniper/asf/pkg/testutil/integration/etcd"
)

// SetDefaultSyncConfig sets config options required by sync.
func SetDefaultSyncConfig(shouldDump bool) {
	setViper(map[string]interface{}{
		constants.ETCDEndpointsVK:     []string{integrationetcd.Endpoint},
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
