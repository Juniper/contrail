package contrail

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	etcdJSONPrefix             = "json"
	virtualNetworkKey          = "virtual_network"
	virtualNetworkSingularPath = "/virtual-network"
	virtualNetworkPluralPath   = "/virtual-networks"
	etcdWatchTimeout           = 1 * time.Second
)

func TestWatcherSynchronizesPostgresDataToEtcdUsingJSONStorage(t *testing.T) {
	//t.Skip("Not implemented") // TODO(Daniel): resolve
	s := integration.NewRunningAPIServer(t, "../../..")
	defer s.Close(t)
	sc := integration.NewAPIServerClient(t, s.URL())
	ec := integration.NewEtcdClient(t)
	defer ec.Close(t)

	uuid := withTestID(t, "uuid")

	// Check there are no resources in API Server with test UUID
	_, err := sc.Do(
		echo.GET,
		path.Join(virtualNetworkSingularPath, uuid),
		nil,
		nil,
		[]int{http.StatusNotFound},
	)
	assert.NoError(t, err, "getting API Server resource failed")

	// Check there are no resources in etcd with test UUID
	gr := ec.GetAllWithPrefix(t, virtualNetworkEtcdKey(uuid))
	assert.Equal(t, int64(0), gr.Count)

	// Start watching resource on etcd
	ctx, _ := context.WithTimeout(context.Background(), etcdWatchTimeout)
	vnWatch := ec.Watch(ctx, virtualNetworkEtcdKey(uuid), clientv3.WithPrefix())

	// Create Virtual Network resource
	vn := &models.VirtualNetwork{
		UUID:   uuid,
		FQName: []string{withTestID(t, "vn-blue")},
	}

	createOutput := &models.CreateVirtualNetworkResponse{}
	_, err = sc.Create(
		virtualNetworkPluralPath,
		models.CreateVirtualNetworkRequest{VirtualNetwork: vn},
		createOutput,
	)
	assert.NoError(t, err, "creating resource in API Server failed")
	//assert.Equal(t, vn, createOutput.VirtualNetwork)

	// Create Watcher service
	//w := integration.NewRunningWatcher(t)
	//defer w.Close(t)

	// Check Virtual Network created in etcd
	wr := <-vnWatch
	require.NoError(t, wr.Err(), "watching virtual network failed")
	if ctx.Err() == context.DeadlineExceeded {
		assert.FailNow(t, "watching virtual network timed out")
	}
	log.WithField("response", wr).Debug("Received etcd Watch response")

	assert.Equal(t, 1, len(wr.Events))
	assert.True(t, wr.Events[0].IsCreate())
	assert.Equal(t, vn, decodeVirtualNetworkJSON(t, wr.Events[0].Kv.Value))
}

func withTestID(t *testing.T, s string) string {
	return fmt.Sprintf("%s-%s", t.Name(), s)
}

func virtualNetworkEtcdKey(uuid string) string {
	return path.Join(etcdJSONPrefix, virtualNetworkKey, uuid)
}

func decodeVirtualNetworkJSON(t *testing.T, bytes []byte) *models.VirtualNetwork {
	var data models.VirtualNetwork
	assert.NoError(t, json.Unmarshal(bytes, &data))
	return &data
}
