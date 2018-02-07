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
	s := integration.NewRunningAPIServer(t, "../../..")
	defer s.Close(t)
	sc := integration.NewAPIServerClient(t, s.URL())
	ec := integration.NewEtcdClient(t)
	defer ec.Close(t)

	testID := getTestID(t)
	uuid := testID + "-uuid"

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
	//vn := &models.VirtualNetwork{
	//	UUID: uuid,
	//	Perms2: &models.PermType2{
	//		Owner:        "",
	//		OwnerAccess:  0,
	//		GlobalAccess: 0,
	//
	//		//Share: MakeShareTypeSlice(),
	//	},
	//	FQName: []string{withTestID(testID, "vn")},
	//}
	vn := models.MakeVirtualNetwork()
	vn.UUID = uuid
	vn.FQName = []string{"default", "admin", testID + "-vn"}
	vn.Perms2 = &models.PermType2{Owner: "admin"}

	log.Debug("Creating virtual network resource")
	createOutput := &models.CreateVirtualNetworkResponse{}
	resp, err := sc.Create(
		virtualNetworkPluralPath,
		models.CreateVirtualNetworkRequest{VirtualNetwork: vn},
		createOutput,
	)
	log.WithField("resp", resp).Debug("Got virtual network create response")
	assert.NoError(t, err, "creating resource in API Server failed")

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

	// Some debugs
	//assert.Equal(t, vn, createOutput.VirtualNetwork)
	//spew.Dump("vn: ", vn)
	//spew.Dump("createOutput.VirtualNetwork: ", createOutput.VirtualNetwork)
	//vnJSON, _ := json.MarshalIndent(vn, "", "    ")
	//vn2JSON, _ := json.MarshalIndent(createOutput.VirtualNetwork, "", "    ")
	//fmt.Print("vnJSON: ", string(vnJSON))
	//fmt.Print("vn2JSON: ", string(vn2JSON))

	// Delete resource
	_, err = sc.Delete(path.Join(virtualNetworkSingularPath, uuid), nil)
	assert.NoError(t, err, "deleting API Server resource failed")
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
