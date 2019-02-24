package replication_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	testReplicationTmpl       = "./test_data/test_replication.tmpl"
	createReplicationTestFile = "./test_data/create_replication.yml"
	//updateReplicationTestFile = "./test_data/update_replication.yml"
	postReq   = "POST"
	putReq    = "PUT"
	deleteReq = "DELETE"
)

var server *integration.APIServer

type httpBodyStore map[string]interface{}

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func handleCreate(t *testing.T, w http.ResponseWriter,
	r *http.Request, vncReqStore map[string][]*httpBodyStore) {
	switch r.URL.Path {
	case "/ports":
		processReqBody(t, w, r, postReq, vncReqStore)
	case "/node-profiles":
		processReqBody(t, w, r, postReq, vncReqStore)
	case "/end-systems":
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
	case "/end-system/:id":
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
	case "/end-system/:id":
		processReqBody(t, w, r, deleteReq, vncReqStore)
	default:
		writeJSONResponse(t, w, 404, nil)
	}
}

func createMockVNCServer(t *testing.T) (*httptest.Server, map[string][]*httpBodyStore) {
	vncReqStore := map[string][]*httpBodyStore{}
	vncServer := httptest.NewServer(
		// NewServer takes a handler.
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case postReq:
				handleCreate(t, w, r, vncReqStore)
			case putReq:
				handleUpdate(t, w, r, vncReqStore)
			case deleteReq:
				handleDelete(t, w, r, vncReqStore)
			default:
				writeJSONResponse(t, w, 404, nil)
			}
		}),
	)
	return vncServer, vncReqStore
}

func writeJSONResponse(t *testing.T, w http.ResponseWriter,
	status int, resp interface{}) {
	bytes, err := json.MarshalIndent(resp, "", "  ")
	assert.NoError(t, err, "failed to marshal response")
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(bytes)
	//_, err = w.Write([]byte("\n"))
	//assert.NoError(t, err, "failed to write response")
}

func processReqBody(t *testing.T, w http.ResponseWriter,
	r *http.Request, reqType string, vncReqStore map[string][]*httpBodyStore) {

	// read the body and save it to a var
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err, "while reading REST req body")

	reqBody := &httpBodyStore{}
	err = json.Unmarshal(body, reqBody)
	assert.NoError(t, err, "while converting vnc req body to json")

	vncReqStore[reqType] = append(vncReqStore[reqType], reqBody)
	writeJSONResponse(t, w, 200, httpBodyStore{})
}

// create cluster and its endpoints
func initTestCluster(t *testing.T, clusterName string) (
	func(), map[string][]*httpBodyStore) {

	keystoneAuthURL := viper.GetString("keystone.authurl")
	clusterUser := clusterName + "_admin"

	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterUser, clusterUser)

	ksPublic := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterUser, clusterUser)

	vncServer, vncReqStore := createMockVNCServer(t)

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

	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testReplicationTmpl, pContext)
	assert.NoError(t, err, "failed to load replication test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	server.ForceProxyUpdate()

	cleanupTestCluster := func() {
		defer cleanup()
		defer ksPrivate.Close()
		defer ksPublic.Close()
		defer vncServer.Close()
	}
	return cleanupTestCluster, vncReqStore
}

func runReplicationTest(t *testing.T) {
	//create test clusters with keystone/config endpoint.
	cleanupTestClusterA, vncReqStoreA := initTestCluster(t, "clusterA")
	defer cleanupTestClusterA()
	cleanupTestClusterB, vncReqStoreB := initTestCluster(t, "clusterB")
	defer cleanupTestClusterB()

	time.Sleep(5 * time.Second)
	//create node-profile, node, port object
	testScenario, err := integration.LoadTest(createReplicationTestFile, nil)
	assert.NoError(t, err, "failed to load test data")
	integration.RunCleanTestScenario(t, testScenario, server)

	//verify update objects
	fmt.Printf("%v", vncReqStoreA)
	fmt.Printf("%v", vncReqStoreB)
}

func TestReplication(t *testing.T) {
	runReplicationTest(t)
}
