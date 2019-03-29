package integration

import (
	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/models"
	integrationetcd "github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

// SetDefaultSyncConfig sets config options required by sync.
func SetDefaultSyncConfig(shouldDump bool) {
	setViperConfig(map[string]interface{}{
		constants.ETCDEndpointsVK:     []string{integrationetcd.Endpoint},
		"sync.storage":                models.JSONCodec.Key(),
		"sync.dump":                   shouldDump,
		"database.type":               basedb.DriverPostgreSQL,
		"database.host":               "localhost",
		"database.user":               dbUser,
		"database.name":               dbName,
		"database.password":           dbPassword,
		"database.dialect":            basedb.DriverPostgreSQL,
		"database.max_open_conn":      100,
		"database.connection_retries": 10,
		"database.retry_period":       3,
		"database.debug":              false,
	})
}
