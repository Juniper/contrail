package watcher_test

import (
	"fmt"
	"path"
	"testing"

	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/watcher"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	jsonPrefix = "json"
	nodeKey    = "nodes"
)

func TestWatcherSynchronizesDataToEtcdUsingJSONStorage(t *testing.T) {
	t.Skip("Not implemented") //TODO(daniel): finish test

	s := givenAPIServer(t)
	defer s.Close(t)
	w := givenRunningWatcher(t, watcher.StorageJSON)
	defer w.Close()
	sc := testutil.NewAPIServerClient(t, s.URL())
	ec := testutil.NewEtcdClient(t)
	defer ec.CloseClient(t)

	uuid := "random-uuid"
	gr := ec.GetAllFromPrefix(t, path.Join(jsonPrefix, nodeKey, uuid))
	assert.Equal(t, int64(0), gr.Count)

	// Create Node resource
	output := make(map[string]interface{})
	_, err := sc.Create(
		nodeKey,
		models.CreateNodeRequest{Node: &models.Node{
			UUID:      uuid,
			FQName:    []string{"test-node"},
			Hostname:  "test-hostname",
			IPAddress: "test-ip-address",
		}},
		nil,
	)
	assert.NoError(t, err)
	fmt.Printf("Created resource: %v", output)

	// Check Node resource in etcd

	//account := e.GetAllFromPrefix(t, path.Join(accountKey, "json", "d9f1c415-f53c-4f78-8dba-2cf62c7165f4"))
	//assert.Len(t, account.Kvs, 1)
	//assert.Equal(t, path.Join(accountKey, "json", "d9f1c415-f53c-4f78-8dba-2cf62c7165f4"), string(account.Kvs[0].Key))
	//assert.Equal(t,
	//	map[string]interface{}{
	//		"id":          "d9f1c415-f53c-4f78-8dba-2cf62c7165f4",
	//		"name":        "default",
	//		"description": "",
	//		"status":      "disabled",
	//	},
	//	decodeJSON(t, account.Kvs[0].Value),
	//)
	//
	//userRole := e.GetAllFromPrefix(t, path.Join(userRoleKey, "json", "00b401f3-b571-4210-b35e-0c4fbcef8dda"))
	//assert.Len(t, userRole.Kvs, 1)
	//assert.Equal(t, path.Join(userRoleKey, "json", "00b401f3-b571-4210-b35e-0c4fbcef8dda"), string(userRole.Kvs[0].Key))
	//assert.Equal(t,
	//	map[string]interface{}{
	//		"id":          "00b401f3-b571-4210-b35e-0c4fbcef8dda",
	//		"name":        "billing",
	//		"description": "",
	//		"public":      float64(1),
	//		"account_id":  "d9f1c415-f53c-4f78-8dba-2cf62c7165f4",
	//	},
	//	decodeJSON(t, userRole.Kvs[0].Value),
	//)
	//
	//rule := e.GetAllFromPrefix(t, path.Join(ruleKey, "json", "ff0b21b1-63d3-41a4-adac-fcdfde136692"))
	//assert.Len(t, rule.Kvs, 1)
	//assert.Equal(t, path.Join(ruleKey, "json", "ff0b21b1-63d3-41a4-adac-fcdfde136692"), string(rule.Kvs[0].Key))
	//assert.Equal(t,
	//	map[string]interface{}{
	//		"id":             "ff0b21b1-63d3-41a4-adac-fcdfde136692",
	//		"name":           "billing_readonly",
	//		"description":    "",
	//		"allow_create":   float64(0),
	//		"allow_delete":   float64(0),
	//		"allow_read":     float64(1),
	//		"allow_update":   float64(0),
	//		"is_owner_group": float64(0),
	//		"is_owner_user":  float64(0),
	//		"resource":       "user,group,member",
	//		"user_role_id":   "00b401f3-b571-4210-b35e-0c4fbcef8dda",
	//	},
	//	decodeJSON(t, rule.Kvs[0].Value),
	//)
}

func TestWatcherSynchronizesDataToEtcdUsingNestedStorage(t *testing.T) {
	t.Skip("Not implemented") //TODO(daniel): implement test

	givenRunningWatcher(t, watcher.StorageNested)

	e := testutil.NewEtcdClient(t)
	defer e.CloseClient(t)

	//account := e.GetAllFromPrefix(t, path.Join(accountKey, "d9f1c415-f53c-4f78-8dba-2cf62c7165f4"))
	//checkNestedKeyValues(t,
	//	path.Join(accountKey, "d9f1c415-f53c-4f78-8dba-2cf62c7165f4"),
	//	account.Kvs,
	//	map[string]interface{}{
	//		"id":          "d9f1c415-f53c-4f78-8dba-2cf62c7165f4",
	//		"name":        "default",
	//		"description": "",
	//		"status":      "disabled",
	//	},
	//)
	//
	//userRole := e.GetAllFromPrefix(t, path.Join(userRoleKey, "00b401f3-b571-4210-b35e-0c4fbcef8dda"))
	//checkNestedKeyValues(t,
	//	path.Join(userRoleKey, "00b401f3-b571-4210-b35e-0c4fbcef8dda"),
	//	userRole.Kvs,
	//	map[string]interface{}{
	//		"id":          "00b401f3-b571-4210-b35e-0c4fbcef8dda",
	//		"name":        "billing",
	//		"description": "",
	//		"public":      "1",
	//		"account_id":  "d9f1c415-f53c-4f78-8dba-2cf62c7165f4",
	//	},
	//)
	//
	//rule := e.GetAllFromPrefix(t, path.Join(ruleKey, "ff0b21b1-63d3-41a4-adac-fcdfde136692"))
	//checkNestedKeyValues(t,
	//	path.Join(ruleKey, "ff0b21b1-63d3-41a4-adac-fcdfde136692"),
	//	rule.Kvs,
	//	map[string]interface{}{
	//		"id":             "ff0b21b1-63d3-41a4-adac-fcdfde136692",
	//		"name":           "billing_readonly",
	//		"description":    "",
	//		"allow_create":   "0",
	//		"allow_delete":   "0",
	//		"allow_read":     "1",
	//		"allow_update":   "0",
	//		"is_owner_group": "0",
	//		"is_owner_user":  "0",
	//		"resource":       "user,group,member",
	//		"user_role_id":   "00b401f3-b571-4210-b35e-0c4fbcef8dda",
	//	},
	//)
}

//func TestWatcherCreatesResourceInEtcd(t *testing.T) {
//	t.Skip("Not implemented") //TODO(daniel): implement test
//	e := clients.NewEtcdClient(t)
//	s := clients.NewAPIServer(t)
//
//	accountID := "test-account-id"
//	account := map[string]interface{}{
//		"id":          accountID,
//		"name":        "test-name",
//		"description": "test-description",
//		"status":      "disabled",
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), watchTimeout)
//	w := e.Watch(ctx, path.Join(accountKey, accountID), clientv3.WithPrefix())
//
//	s.Create(account)
//
//	r := <-w
//	cancel()
//	checkWatchedKeyValues(t, r, account)
//
//	e.CloseClient(t)
//}
//
//func checkWatchedKeyValues(t *testing.T, r clientv3.WatchResponse, resource map[string]interface{}) {
//	if err := r.Err(); err != nil {
//		t.Fatalf("Cannot watch etcd events: %s", err)
//	}
//	log.WithField("response", r).Debug("Received etcd Watch response")
//
//	etcdResource := make(map[string]interface{})
//	for _, event := range r.Events {
//		if !event.IsCreate() {
//			t.Errorf("Event %s should be create event", event)
//		}
//		etcdResource[string(event.Kv.Key)] = event.Kv.Value
//	}
//	assert.Equal(t, resource, etcdResource)
//}

//func decodeJSON(t *testing.T, bytes []byte) map[string]interface{} {
//	var data map[string]interface{}
//	if err := json.Unmarshal(bytes, &data); err != nil {
//		t.FailNow()
//	}
//	return data
//}

func TestWatcherDumpsExistingDataToEtcdUsingJSONStorage(t *testing.T) {
	t.Skip("Not implemented") //TODO(daniel): implement test
}

func TestWatcherDumpsExistingDataToEtcdUsingNestedStorage(t *testing.T) {
	t.Skip("Not implemented") //TODO(daniel): implement test
}

func givenAPIServer(t *testing.T) *testutil.APIServer {
	return testutil.NewAPIServer(t, "../..")
}

func givenRunningWatcher(t *testing.T, storage string) *watcher.Service {
	s, err := watcher.NewService(&watcher.Config{
		Database: watcher.DBConfig{
			Host:     fmt.Sprintf("%s:%d", testutil.DBHostname, testutil.DBPort),
			User:     testutil.DBUser,
			Password: testutil.DBPassword,
			Name:     testutil.DBName,
		},
		Etcd: watcher.EtcdConfig{
			Endpoints: []string{testutil.EtcdEndpoint},
		},
		Storage: storage,
	})
	require.NoError(t, err)

	go func() {
		err := s.Run()
		assert.NoError(t, err)
	}()

	return s
}

//func checkNestedKeyValues(t *testing.T, prefix string, actualKvs []*mvccpb.KeyValue,
//	expectedResource map[string]interface{}) {
//	actualResource := make(map[string]interface{})
//	for _, kv := range actualKvs {
//		actualResource[string(kv.Key)] = string(kv.Value)
//	}
//	assert.Equal(t, withPrefix(expectedResource, prefix), actualResource)
//}

//func withPrefix(resource map[string]interface{}, prefix string) map[string]interface{} {
//	r := make(map[string]interface{})
//	for k, v := range resource {
//		r[path.Join(prefix, k)] = v
//	}
//	return r
//}
