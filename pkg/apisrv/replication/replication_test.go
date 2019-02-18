package replication

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/deploy/cluster"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	testClusterTmpl       = "./test_data/test_cluster.tmpl"
	createReplicationTmpl = "./test_data/create_replication.tmpl"
	updateReplicationTmpl = "./test_data/update_replication.tmpl"
	postReq               = "POST"
	putReq                = "PUT"
	deleteReq             = "DELETE"
)

var server *integration.APIServer

type httpBodyStore map[string]interface{}

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func TestReplication(t *testing.T) {
	runReplicationTest(t)
}

func runReplicationTest(t *testing.T) {

	var vncReqStore map[string][]*httpBodyStore
	testServer := createMockVNCServer(t, vncReqStore)
	defer testServer.Close()

	parsedURL, err := url.Parse(testServer.URL)
	assert.NoError(t, err, "while parsing mock vnc server URL")

	rContext := pongo2.Context{
		"config_port": parsedURL.Port(),
	}

	expectedEndpoints := map[string]string{
		"config":    testServer.URL,
		"nodejs":    "https://127.0.0.1:8143",
		"telemetry": "http://127.0.0.1:8081",
	}

	wg := &sync.WaitGroup{}
	runTestAndVerify(t, wg, rContext, expectedEndpoints, vncReqStore)
	wg.Wait()

}

func runTestAndVerify(t *testing.T, wg *sync.WaitGroup,
	rContext map[string]interface{}, expectedEndpoints map[string]string,
	vncReqStore map[string][]*httpBodyStore) {

	// Need to start replication service

	cleanupContrailCluster, initTestCleanUp := initTestSetup(t,
		rContext, expectedEndpoints)
	defer cleanupContrailCluster()
	defer initTestCleanUp()

	//create node-profile, node, port object
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario,
		createReplicationTmpl, rContext)
	assert.NoError(t, err, "failed to load create replication template")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()
	fmt.Printf("%v", vncReqStore)

	//verify update objects
}

// create cluster objects needed to create endpoints
func initTestSetup(t *testing.T, rContext map[string]interface{},
	expectedEndpoints map[string]string) (func(), func()) {

	// mock keystone to let access server after cluster create
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPublic := integration.MockServerWithKeystone("127.0.0.1:35357", keystoneAuthURL)
	defer ksPublic.Close()
	ksPrivate := integration.MockServerWithKeystone("127.0.0.1:5000", keystoneAuthURL)
	defer ksPrivate.Close()

	// create cluster config
	s := &client.HTTP{
		Endpoint: server.URL(),
		InSecure: true,
		AuthURL:  server.URL() + "/keystone/v3",
		ID:       "alice",
		Password: "alice_password",
		Scope: client.GetKeystoneScope(
			"default", "default", "admin", "admin"),
	}
	s.Init()
	_, err := s.Login(context.Background())
	assert.NoError(t, err, "failed to login")
	cleanupContrailCluster := func() {
		deleteCluster(s)
	}
	return cleanupContrailCluster,
		createCluster(t, rContext, s, expectedEndpoints)

}

func createCluster(t *testing.T, rContext map[string]interface{},
	s *client.HTTP, expectedEndpoints map[string]string) func() {

	// create contrail cluster and its depenendent objects
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testClusterTmpl, rContext)
	assert.NoError(t, err, "failed to load cluster test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)

	config := &cluster.Config{
		APIServer:        s,
		ClusterID:        "test_cluster_uuid",
		Action:           "create",
		LogLevel:         "debug",
		TemplateRoot:     "../../deploy/cluster/templates/",
		TestTemplateRoot: "../../deploy/cluster/test_data/",
		WorkRoot:         "/tmp/contrail_cluster",
		Test:             true,
		LogFile:          "/tmp/contrail_cluster" + "/deploy.log",
	}
	// create cluster
	clusterDeployer, err := cluster.NewCluster(config)
	assert.NoError(t, err, "failed to create cluster manager to create cluster")
	deployer, err := clusterDeployer.GetDeployer()
	assert.NoError(t, err, "failed to create deployer")
	err = deployer.Deploy()
	assert.NoError(t, err, "failed to manage(create) cluster")
	// make sure all endpoints are created
	err = integration.VerifyEndpoints(t, "test_cluster_uuid", &testScenario, expectedEndpoints)
	if err != nil {
		assert.NoError(t, err, err.Error())
	}

	return cleanup
}

func deleteCluster(s *client.HTTP) {
	config := &cluster.Config{
		APIServer:        s,
		ClusterID:        "test_cluster_uuid",
		Action:           "delete",
		LogLevel:         "debug",
		TemplateRoot:     "../../deploy/cluster/templates/",
		TestTemplateRoot: "../../deploy/cluster/test_data/",
		WorkRoot:         "/tmp/contrail_cluster",
		Test:             true,
		LogFile:          "/tmp/contrail_cluster" + "/deploy.log",
	}
	// create cluster
	clusterDeployer, err := cluster.NewCluster(config)
	if err != nil {
		logrus.WithError(err).Error("deleting contrail cluster failed")
	}
	deployer, err := clusterDeployer.GetDeployer()
	if err != nil {
		logrus.WithError(err).Error("deleting contrail cluster failed")
	}
	err = deployer.Deploy()
	if err != nil {
		logrus.WithError(err).Error("deleting contrail cluster failed")
	}
}

func createMockVNCServer(t *testing.T,
	vncReqStore map[string][]*httpBodyStore) *httptest.Server {

	testServer := httptest.NewServer(

		// NewServer takes a handler.
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			switch r.Method {
			// handle create
			case postReq:
				switch r.URL.Path {
				case "/ports":
					processReqBody(t, w, r, postReq, vncReqStore)
				case "/node-profiles/":
					processReqBody(t, w, r, postReq, vncReqStore)
				case "/end-systems/":
					processReqBody(t, w, r, postReq, vncReqStore)
				default:
					writeJSONResponse(w, 404, nil)
				}
			// handle update
			case putReq:
				switch r.URL.Path {
				case "/port/:id":
					processReqBody(t, w, r, putReq, vncReqStore)
				case "/node-profile/:id":
					processReqBody(t, w, r, putReq, vncReqStore)
				case "/end-system/:id":
					processReqBody(t, w, r, putReq, vncReqStore)
				default:
					writeJSONResponse(w, 404, nil)
				}
			// handle delete
			case deleteReq:
				switch r.URL.Path {
				case "/port/:id":
					processReqBody(t, w, r, deleteReq, vncReqStore)
				case "/node-profile/:id":
					processReqBody(t, w, r, deleteReq, vncReqStore)
				case "/end-system/:id":
					processReqBody(t, w, r, deleteReq, vncReqStore)
				default:
					writeJSONResponse(w, 404, nil)
				}
			default:
				writeJSONResponse(w, 404, nil)
			}
		}),
	)
	return testServer
}

func writeJSONResponse(w http.ResponseWriter,
	status int, resp interface{}) error {
	bytes, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(bytes)
	w.Write([]byte("\n"))
	return nil
}

func processReqBody(t *testing.T, w http.ResponseWriter,
	r *http.Request, reqType string, vncReqStore map[string][]*httpBodyStore) {

	// read the body and save it to a var
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		assert.NoError(t, err, "while reading REST req body")
	}

	reqBody := &httpBodyStore{}
	err = json.Unmarshal(body, reqBody)
	if err != nil {
		assert.NoError(t, err, "while converting vnc req body to json")
	}

	vncReqStore[reqType] = append(vncReqStore[reqType], reqBody)
	writeJSONResponse(w, 200, httpBodyStore{})
}
