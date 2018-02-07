package contrail

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
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
	etcdWatchTimeout           = 10 * time.Second
)

func TestSyncSynchronizesPostgresDataToEtcdUsingJSONStorage(t *testing.T) {
	t.Skip("Skipping because of not sufficient CI test env") // TODO(Daniel): resolve

	s := integration.NewRunningAPIServer(t, "../../..")
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())
	ec := integration.NewEtcdClient(t)
	defer ec.Close(t)

	testID := getTestID(t)
	uuid := testID + "-uuid"

	// Check there are no resources in API Server with test UUID
	_, err := hc.Do(
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
	vn := &models.VirtualNetwork{}
	//vn := models.MakeVirtualNetwork()
	vn.UUID = uuid
	vn.FQName = []string{"default", "admin", testID + "-vn"}
	vn.Perms2 = &models.PermType2{Owner: "admin"}
	vn.ExternalIpam = true
	vn.IsShared = true

	log.Debug("Creating virtual network resource")
	createOutput := &models.CreateVirtualNetworkResponse{}
	resp, err := hc.Create(
		virtualNetworkPluralPath,
		models.CreateVirtualNetworkRequest{VirtualNetwork: vn},
		createOutput,
	)
	log.WithField("resp", resp).Debug("Got virtual network create response")
	assert.NoError(t, err, "creating resource in API Server failed")

	var readOutput models.GetVirtualNetworkResponse
	_, err = hc.Read(virtualNetworkSingularPath+"/"+uuid, &readOutput)
	assert.NoError(t, err)

	// Create Sync service
	sync := integration.NewRunningSync(t)
	defer sync.Close(t)

	// Check Virtual Network created in etcd
	wr := <-vnWatch
	require.NoError(t, wr.Err(), "watching virtual network failed")
	require.NotEqual(t, ctx.Err(), context.DeadlineExceeded, "watching virtual network timed out")
	log.WithField("response", wr).Debug("Received etcd Watch response")

	assert.Equal(t, 1, len(wr.Events))
	assert.True(t, wr.Events[0].IsCreate())
	assert.Equal(t, createOutput.VirtualNetwork, decodeVirtualNetworkJSON(t, wr.Events[0].Kv.Value))

	createdVN, _ := json.MarshalIndent(createOutput.VirtualNetwork, "", "    ")
	readVN, _ := json.MarshalIndent(readOutput.VirtualNetwork, "", "    ")
	syncedVN, _ := json.MarshalIndent(decodeVirtualNetworkJSON(t, wr.Events[0].Kv.Value), "", "    ")
	fmt.Print("createdVN: ", string(createdVN))
	fmt.Print("readVN: ", string(readVN))
	fmt.Print("syncedVN: ", string(syncedVN))
	assert.Equal(t, readOutput.VirtualNetwork, decodeVirtualNetworkJSON(t, wr.Events[0].Kv.Value), "readVN is not equal")

	// Delete resource
	_, err = hc.Delete(path.Join(virtualNetworkSingularPath, uuid), nil)
	assert.NoError(t, err, "deleting API Server resource failed")

	time.Sleep(1 * time.Second) // TODO(Daniel): change Sync not to throw error on Dump context cancellation
}

func getTestID(t *testing.T) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v-%v", t.Name(), rand.Uint64())
}

func virtualNetworkEtcdKey(uuid string) string {
	return path.Join(etcdJSONPrefix, virtualNetworkKey, uuid)
}

func decodeVirtualNetworkJSON(t *testing.T, vnBytes []byte) *models.VirtualNetwork {
	var vn models.VirtualNetwork
	assert.NoError(t, json.Unmarshal(vnBytes, &vn))
	return &vn
}
