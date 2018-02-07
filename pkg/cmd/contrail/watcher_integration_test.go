package contrail

import (
	//"context"
	//"fmt"
	//"time"
	"encoding/json"
	"path"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/models"
	//"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

const (
	etcdJSONPrefix = "json"
	nodeKey        = "nodes"
	//etcdWatchTimeout = 10 * time.Second
)

func TestWatcherSynchronizesPostgresDataToEtcdUsingJSONStorage(t *testing.T) {
	t.Skip()
	s := integration.NewRunningAPIServer(t, "../../..")
	defer s.Close(t)

	//sc := integration.NewAPIServerClient(t, s.URL())
	//ec := integration.NewEtcdClient(t)
	//defer ec.CloseClient(t)

	// Check there are no resources in DB
	//uuid := "random-uuid"
	//gr := ec.GetAllFromPrefix(t, nodeEtcdKey(uuid))
	//assert.Equal(t, int64(0), gr.Count)

	// Start watching resource on etcd
	//ctx, cancel := context.WithTimeout(context.Background(), etcdWatchTimeout)
	//wc := ec.Watch(ctx, nodeEtcdKey(uuid), clientv3.WithPrefix())

	// Create Node resource
	//node := models.Node{
	//	UUID:      uuid,
	//	FQName:    []string{"test-node"},
	//	Hostname:  "test-hostname",
	//	IPAddress: "test-ip-address",
	//}
	//
	//output := make(map[string]interface{})
	//_, err := sc.Create(
	//	nodeKey,
	//	models.CreateNodeRequest{Node: &node},
	//	nil,
	//)
	//assert.NoError(t, err)
	//log.Debug("Created resource: %v", output)
	//

	w := integration.NewRunningWatcher(t)
	time.Sleep(time.Second)
	defer w.Close(t)

	//// Check Node resource in etcd
	//r := <-wc
	//cancel()
	//checkNodeResourceCreated(t, r, node)
}

func nodeEtcdKey(uuid string) string {
	return path.Join(etcdJSONPrefix, nodeKey, uuid)
}

func checkNodeResourceCreated(t *testing.T, r clientv3.WatchResponse, node *models.Node) {
	if err := r.Err(); err != nil {
		t.Fatalf("Cannot watch etcd events: %s", err)
	}
	log.WithField("response", r).Debug("Received etcd Watch response")

	assert.Equal(t, 1, len(r.Events))
	assert.True(t, r.Events[0].IsCreate())
	assert.Equal(t, node, decodeJSON(t, r.Events[0].Kv.Value))
}

func decodeJSON(t *testing.T, bytes []byte) *models.Node {
	var data models.Node
	assert.NoError(t, json.Unmarshal(bytes, &data))
	return &data
}
