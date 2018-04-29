package apisrv

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/common"
)

const (
	key             = "name"
	privatePortList = "private_port_list"
	publicPortList  = "public_port_list"
)

type mockPortsResponse struct {
	Name string `json:"name"`
}

func mockServer(routes map[string]interface{}) *httptest.Server {
	// Echo instance
	e := echo.New()

	// Routes
	for route, handler := range routes {
		e.GET(route, handler.(echo.HandlerFunc))
	}
	mockServer := httptest.NewServer(e)
	return mockServer
}

func mockServerWithKeystone() *httptest.Server {
	// Echo instance
	e := echo.New()
	keystoneAuthURL := viper.GetString("keystone.authurl")
	keystoneClient := keystone.NewKeystoneClient(keystoneAuthURL, true)
	endpointStore := apicommon.MakeEndpointStore()
	k, err := keystone.Init(e, endpointStore, keystoneClient)
	if err != nil {
		return nil
	}

	// Routes
	e.POST("/v3/auth/tokens", k.CreateTokenAPI)
	e.GET("/v3/auth/tokens", k.ValidateTokenAPI)
	e.GET("/v3/auth/projects", k.GetProjectAPI)
	mockServer := httptest.NewServer(e)
	return mockServer
}

func runEndpointTest(t *testing.T, clusterName string,
	extraTasks bool) (string, *TestScenario, *httptest.Server, *httptest.Server) {
	routes := map[string]interface{}{
		"/ports": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK,
				&mockPortsResponse{Name: clusterName + privatePortList})
		}),
	}
	neutronPrivate := mockServer(routes)
	routes = map[string]interface{}{
		"/ports": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK,
				&mockPortsResponse{Name: clusterName + publicPortList})
		}),
	}
	neutronPublic := mockServer(routes)

	context := pongo2.Context{
		"extra_tasks":   extraTasks,
		"cluster_name":  clusterName,
		"endpoint_name": "neutron",
		"private_url":   neutronPrivate.URL,
		"public_url":    neutronPublic.URL,
	}

	testFile := GetTestFromTemplate(t, "./test_data/test_endpoint.tmpl", context)

	var testScenario TestScenario
	err := LoadTestScenario(&testScenario, testFile)
	assert.NoError(t, err, "failed to load test data")
	RunTestScenario(t, &testScenario)

	return testFile, &testScenario, neutronPublic, neutronPrivate
}

func verifyProxy(t *testing.T, testScenario *TestScenario, url string,
	clusterName string, expected string) bool {
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		_, err := client.Read(url, &response)
		ok := assert.NoError(t, err, "failed to proxy %s", url)
		if !ok {
			return ok
		}
		ok = common.AssertEqual(t,
			map[string]interface{}{key: clusterName + expected},
			response,
			fmt.Sprintf("Unexpected Response: %s", response))
		if !ok {
			return ok
		}
	}
	return true
}

func verifyKeystoneEndpoint(testScenario *TestScenario) error {
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		_, err := client.Read("/keystone/v3/auth/tokens", &response)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestProxyEndpoint(t *testing.T) {
	// Create a cluster and its neutron endpoint
	clusterAName := "clusterA"
	testFile, testScenario, neutronPublic, neutronPrivate := runEndpointTest(
		t, clusterAName, true)
	// remove tempfile after test
	defer os.Remove(testFile) // nolint: errcheck
	defer neutronPrivate.Close()
	defer neutronPublic.Close()

	// Wait a sec for the dynamic proxy to be created/updated
	time.Sleep(2 * time.Second)

	// verify proxies
	ok := verifyProxy(t, testScenario,
		"/proxy/"+clusterAName+"_uuid/neutron/ports", clusterAName, publicPortList)
	if !ok {
		return
	}
	ok = verifyProxy(t, testScenario,
		"/proxy/"+clusterAName+"_uuid/neutron/private/ports", clusterAName, privatePortList)
	if !ok {
		return
	}

	// create one more cluster/neutron endpoint for new cluster
	clusterBName := "clusterB"
	testFile, testScenario, neutronPublic, neutronPrivate = runEndpointTest(
		t, clusterBName, false)
	// remove tempfile after test
	defer os.Remove(testFile) // nolint: errcheck
	defer neutronPrivate.Close()
	defer neutronPublic.Close()
	// Wait a sec for the dynamic proxy to be created/updated
	time.Sleep(2 * time.Second)

	// verify new proxies
	ok = verifyProxy(t, testScenario,
		"/proxy/"+clusterBName+"_uuid/neutron/ports",
		clusterBName, publicPortList)
	if !ok {
		return
	}
	ok = verifyProxy(t, testScenario,
		"/proxy/"+clusterBName+"_uuid/neutron/private/ports",
		clusterBName, privatePortList)
	if !ok {
		return
	}

	// verify existing proxies, make sure the proxy prefix is updated with cluster id
	ok = verifyProxy(t, testScenario,
		"/proxy/"+clusterAName+"_uuid/neutron/ports",
		clusterAName, publicPortList)
	if !ok {
		return
	}
	ok = verifyProxy(t, testScenario,
		"/proxy/"+clusterAName+"_uuid/neutron/private/ports",
		clusterAName, privatePortList)
	if !ok {
		return
	}
}

func TestKeystoneEndpoint(t *testing.T) {
	ksPrivate := mockServerWithKeystone()
	defer ksPrivate.Close()

	ksPublic := mockServerWithKeystone()
	defer ksPublic.Close()

	clusterName := "clusterA"
	context := pongo2.Context{
		"extra_tasks":   true,
		"cluster_name":  clusterName,
		"endpoint_name": "keystone",
		"private_url":   ksPrivate.URL,
		"public_url":    ksPublic.URL,
	}

	testFile := GetTestFromTemplate(t, "./test_data/test_endpoint.tmpl", context)
	// remove tempfile after test
	defer os.Remove(testFile) // nolint: errcheck

	var testScenario TestScenario
	err := LoadTestScenario(&testScenario, testFile)
	assert.NoError(t, err, "failed to load endpoint create test data")
	RunTestScenario(t, &testScenario)

	// Wait a sec for the dynamic proxy to be created/updated
	time.Sleep(2 * time.Second)

	// Login to new remote keystone
	for _, client := range testScenario.Clients {
		err := client.Login()
		assert.NoError(t, err, "client failed to login remote keystone")
	}
	// verify auth (remote keystone)
	err = verifyKeystoneEndpoint(&testScenario)
	assert.NoError(t, err,
		"failed to validate token with remote keystone")

	// Delete endpoint test
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterName)
		_, err = client.Delete(url, &response)
		assert.NoError(t, err, "failed to delete keystone endpoint")
		break
	}
	// Wait a sec for the dynamic proxy to be deleted
	time.Sleep(2 * time.Second)
	// Login to new local keystone
	for _, client := range testScenario.Clients {
		err = client.Login()
		assert.NoError(t, err, "client failed to login local keystone")
	}
	// verify auth (local keystone)
	err = verifyKeystoneEndpoint(&testScenario)
	assert.NoError(t, err,
		"failed to validate token with local keystone after endpoint delete")

	// Recreate endpoint
	err = LoadTestScenario(&testScenario, testFile)
	assert.NoError(t, err, "failed to load endpoint create test data")
	RunTestScenario(t, &testScenario)
	// Wait a sec for the dynamic proxy to be created
	time.Sleep(2 * time.Second)
	// Login to new remote keystone
	for _, client := range testScenario.Clients {
		err = client.Login()
		assert.NoError(t, err, "client failed to login remote keystone")
	}
	// verify auth (remote keystone)
	err = verifyKeystoneEndpoint(&testScenario)
	assert.NoError(t, err,
		"failed to validate token with remote keystone after endpoint re-create")

	// Cleanup endpoint test
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterName)
		_, err = client.Delete(url, &response)
		assert.NoError(t, err, "failed to delete keystone endpoint")
		break
	}
	// Wait a sec for the dynamic proxy to be deleted
	time.Sleep(2 * time.Second)
}
