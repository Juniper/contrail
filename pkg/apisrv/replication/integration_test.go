package replication_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testReplicationTmpl       = "./test_data/test_replication.tmpl"
	createReplicationTestFile = "./test_data/create_replication.yml"
	updateReplicationTestFile = "./test_data/update_replication.yml"
	deleteReplicationTestFile = "./test_data/delete_replication.yml"
	postReq                   = "POST"
	putReq                    = "PUT"
	deleteReq                 = "DELETE"
	readReq                   = "GET"
)

var server *integration.APIServer

type httpBodyStore map[string]interface{}

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func handleRead(t *testing.T, w http.ResponseWriter,
	r *http.Request, vncReqStore map[string][]*httpBodyStore) {
	processReqBody(t, w, r, readReq, vncReqStore)
}

func handleCreate(t *testing.T, w http.ResponseWriter,
	r *http.Request, vncReqStore map[string][]*httpBodyStore) {
	switch r.URL.Path {
	case "/ports":
		processReqBody(t, w, r, postReq, vncReqStore)
	case "/node-profiles":
		processReqBody(t, w, r, postReq, vncReqStore)
	case "/nodes":
		processReqBody(t, w, r, postReq, vncReqStore)
	case "/ref-update":
		processReqBody(t, w, r, postReq, vncReqStore)
	default:
		writeJSONResponse(t, w, 404, nil)
	}
}

func handleUpdate(t *testing.T, w http.ResponseWriter,
	r *http.Request, vncReqStore map[string][]*httpBodyStore) {
	switch r.URL.Path {
	case "/port/:id":
		processReqBody(t, w, r, putReq, vncReqStore)
	case "/node-profile/:id":
		processReqBody(t, w, r, putReq, vncReqStore)
	case "/node/:id":
		processReqBody(t, w, r, putReq, vncReqStore)
	default:
		writeJSONResponse(t, w, 404, nil)
	}
}

func handleDelete(t *testing.T, w http.ResponseWriter,
	r *http.Request, vncReqStore map[string][]*httpBodyStore) {
	switch r.URL.Path {
	case "/port/:id":
		processReqBody(t, w, r, deleteReq, vncReqStore)
	case "/node-profile/:id":
		processReqBody(t, w, r, deleteReq, vncReqStore)
	case "/node/:id":
		processReqBody(t, w, r, deleteReq, vncReqStore)
	default:
		writeJSONResponse(t, w, 404, nil)
	}
}

func createMockVNCServer(t *testing.T, expectedCount int) (
	*httptest.Server, map[string][]*httpBodyStore, chan struct{},
) {
	vncReqStore := map[string][]*httpBodyStore{}
	done := make(chan struct{})
	vncServer := httptest.NewServer(
		// NewServer takes a handler.
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if expectedCount == 0 {
				close(done)
			}
			expectedCount--
			switch r.Method {
			case postReq:
				handleCreate(t, w, r, vncReqStore)
			case putReq:
				handleUpdate(t, w, r, vncReqStore)
			case deleteReq:
				handleDelete(t, w, r, vncReqStore)
			case readReq:
				handleRead(t, w, r, vncReqStore)
			default:
				writeJSONResponse(t, w, 404, nil)
			}
		}),
	)
	return vncServer, vncReqStore, done
}

//nolint: govet
func writeJSONResponse(t *testing.T, w http.ResponseWriter,
	status int, resp interface{}) {
	bytes, err := json.MarshalIndent(resp, "", "  ")
	assert.NoError(t, err, "failed to marshal response")
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	//nolint: errcheck
	w.Write(bytes)
	//nolint: errcheck
	w.Write([]byte("\n"))
}

func processReqBody(t *testing.T, w http.ResponseWriter,
	r *http.Request, reqType string, vncReqStore map[string][]*httpBodyStore) {

	// read the body and save it to a var
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err, "while reading REST req body")

	reqBody := &httpBodyStore{}
	if len(body) > 0 {
		err = json.Unmarshal(body, reqBody)
		assert.NoError(t, err, "while converting vnc req body to json")
	}

	vncReqStore[reqType] = append(vncReqStore[reqType], reqBody)
	var resp httpBodyStore
	if reqType == readReq {
		urlParts := strings.Split(r.URL.Path, "/")
		switch urlParts[1] {
		case "port":
			var pi []interface{}
			pi = append(pi, map[string]string{
				"uuid": "test_pi_uuid",
			})
			resp = httpBodyStore{
				"port": map[string]interface{}{
					"physical_interface_back_refs": pi,
				},
			}
		}
	} else {
		resp = httpBodyStore{}
	}
	writeJSONResponse(t, w, 200, resp)
}

// create cluster and its endpoint
func initTestCluster(
	t *testing.T, clusterName string, expectedCount int,
) (func(), map[string][]*httpBodyStore, chan struct{}) {

	keystoneAuthURL := viper.GetString("keystone.authurl")
	clusterUser := clusterName + "_admin"

	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterUser, clusterUser)

	ksPublic := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterUser, clusterUser)

	vncServer, vncReqStore, done := createMockVNCServer(t, expectedCount)

	pContext := pongo2.Context{
		"cluster_name":    clusterName,
		"endpoint_name":   clusterName + "_keystone",
		"endpoint_prefix": "keystone",
		"private_url":     ksPrivate.URL,
		"public_url":      ksPublic.URL,
		"manage_parent":   true,
		"admin_user":      clusterUser,
		"config_url":      vncServer.URL,
	}

	ts, err := integration.LoadTest(testReplicationTmpl, pContext)
	require.NoError(t, err, "failed to load replication test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	server.ForceProxyUpdate()

	cleanupTestCluster := func() {
		defer cleanup()
		defer ksPrivate.Close()
		defer ksPublic.Close()
		defer vncServer.Close()
	}
	return cleanupTestCluster, vncReqStore, done
}

func TestReplication(t *testing.T) {
	//create test clusters with keystone/config endpoint.
	cleanupTestClusterA, vncReqStoreA, doneA := initTestCluster(t, t.Name()+"_clusterA", 2)
	defer cleanupTestClusterA()
	cleanupTestClusterB, vncReqStoreB, doneB := initTestCluster(t, t.Name()+"_clusterB", 2)
	defer cleanupTestClusterB()

	//create node-profile, node, port object
	testScenario, err := integration.LoadTest(createReplicationTestFile, nil)
	require.NoError(t, err, "failed to load test data")
	cleanup := integration.RunDirtyTestScenario(t, testScenario, server)
	defer cleanup()

	assertCloses(t, doneA)
	assertCloses(t, doneB)
	//verify created objects
	verifyVNCReqStore(t, postReq, vncReqStoreA, testScenario)
	verifyVNCReqStore(t, postReq, vncReqStoreB, testScenario)
}

func TestReplicateAlreadyCreatedResources(t *testing.T) {
	//create node-profile, node, port object
	ts, err := integration.LoadTest(createReplicationTestFile, nil)
	require.NoError(t, err, "failed to load test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	//create test clusters with keystone/config endpoint.
	cleanupTestClusterA, vncReqStoreA, doneA := initTestCluster(t, "clusterA", 2)
	defer cleanupTestClusterA()
	cleanupTestClusterB, vncReqStoreB, doneB := initTestCluster(t, "clusterB", 2)
	defer cleanupTestClusterB()

	// Update node* objects to force replication to the new endpoints.
	ts, err = integration.LoadTest(updateReplicationTestFile, nil)
	require.NoError(t, err, "failed to load update data")
	integration.RunDirtyTestScenario(t, ts, server)

	server.ForceProxyUpdate()

	assertCloses(t, doneA)
	assertCloses(t, doneB)
	//verify created objects
	verifyVNCReqStore(t, postReq, vncReqStoreA, ts)
	verifyVNCReqStore(t, postReq, vncReqStoreB, ts)
}

func TestDeleteNonReplicatedResources(t *testing.T) {
	//create node-profile, node, port object
	ts, err := integration.LoadTest(createReplicationTestFile, nil)
	require.NoError(t, err, "failed to load test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	//create test clusters with keystone/config endpoint.
	cleanupTestClusterA, _, doneA := initTestCluster(t, "clusterA", 2)
	defer cleanupTestClusterA()
	cleanupTestClusterB, _, doneB := initTestCluster(t, "clusterB", 2)
	defer cleanupTestClusterB()

	// Delete node* objects.
	ts, err = integration.LoadTest(deleteReplicationTestFile, nil)
	require.NoError(t, err, "failed to load delete data")
	integration.RunDirtyTestScenario(t, ts, server)

	server.ForceProxyUpdate()

	assertCloses(t, doneA)
	assertCloses(t, doneB)
}

func assertCloses(t *testing.T, c chan struct{}) {
	select {
	case <-time.After(5 * time.Second):
		assert.Fail(t, "timeout passed waiting for channel to close")
	case <-c:
	}
}

// nolint: gocyclo
func verifyVNCReqStore(t *testing.T, req string,
	vncReqStore map[string][]*httpBodyStore,
	testScenario *integration.TestScenario) {

	postReqs, ok := vncReqStore[req]
	if ok {
		for _, eachReq := range postReqs {
			for _, task := range testScenario.Workflow {
				//nolint: errcheck
				expectMap, _ := task.Expect.(map[string]interface{})
				reqData, schemaPresent := (*eachReq)["node-profile"]
				if schemaPresent && task.Request.Path == "/node-profiles" {

					_ = testutil.AssertEqual(t, expectMap["node-profile"], reqData,
						fmt.Sprintf("node-profile req not replicated to vnc server"))
				}

				reqData, schemaPresent = (*eachReq)["node"]
				if schemaPresent && task.Request.Path == "/nodes" {
					_ = testutil.AssertEqual(t, expectMap["node"], reqData,
						fmt.Sprintf("node req not replicated to vnc server"))
				}

				reqData, schemaPresent = (*eachReq)["port"]
				if schemaPresent && task.Request.Path == "/ports" {
					//nolint: errcheck
					expectedPort, _ := expectMap["port"].(map[string]interface{})
					expectedPort["parent_type"] = "node"
					_ = testutil.AssertEqual(t, expectedPort, reqData,
						fmt.Sprintf("port req not replicated to vnc server"))
				}
			}
		}
	}
	assert.True(t, ok, "post req not found in test replication store")
}
