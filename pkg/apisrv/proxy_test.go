package apisrv

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/common"
)

const (
	key              = "name"
	privatePortList  = "private_port_list"
	publicPortList   = "public_port_list"
	testEndpointFile = "./test_data/test_endpoint.tmpl"
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

func runEndpointTest(t *testing.T, clusterName string) (*TestScenario,
	*httptest.Server, *httptest.Server, func()) {
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
		"cluster_name":  clusterName,
		"endpoint_name": "neutron",
		"private_url":   neutronPrivate.URL,
		"public_url":    neutronPublic.URL,
	}

	var testScenario TestScenario
	err := LoadTestScenario(&testScenario, testEndpointFile, context)
	assert.NoError(t, err, "failed to load test data")
	cleanup := RunDirtyTestScenario(t, &testScenario)

	return &testScenario, neutronPublic, neutronPrivate, cleanup
}

func verifyProxy(t *testing.T, testScenario *TestScenario, url string,
	clusterName string, expected string) bool {
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		_, err := client.Read(url, &response)
		if err != nil {
			fmt.Printf("Reading: %s, Response: %s", url, err)
			return false
		}
		ok := common.AssertEqual(t,
			map[string]interface{}{key: clusterName + expected},
			response,
			fmt.Sprintf("Unexpected Response: %s", response))
		if !ok {
			return ok
		}
	}
	return true
}

func verifyKeystoneEndpoint(testScenario *TestScenario, testInvalidUser bool) error {
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		_, err := client.Read("/keystone/v3/auth/tokens", &response)
		if err != nil {
			return err
		}
		if !testInvalidUser {
			break
		}
	}
	return nil
}

func TestProxyEndpoint(t *testing.T) {
	// Create a cluster and its neutron endpoint
	clusterAName := "clusterA"
	testScenario, clusterANeutronPublic, clusterANeutronPrivate, cleanup1 := runEndpointTest(
		t, clusterAName)
	defer cleanup1()
	// remove tempfile after test
	defer clusterANeutronPrivate.Close()
	defer clusterANeutronPublic.Close()

	APIServer.ForceProxyUpdate()

	// verify proxies
	verifyProxyAndTestIt(t, testScenario, clusterAName, true)

	// create one more cluster/neutron endpoint for new cluster
	clusterBName := "clusterB"
	testScenario, neutronPublic, neutronPrivate, cleanup2 := runEndpointTest(
		t, clusterBName)
	defer cleanup2()
	// remove tempfile after test
	defer neutronPrivate.Close()
	defer neutronPublic.Close()

	APIServer.ForceProxyUpdate()

	// verify new proxies
	verifyProxyAndTestIt(t, testScenario, clusterBName, true)

	// verify existing proxies, make sure the proxy prefix is updated with cluster id
	verifyProxyAndTestIt(t, testScenario, clusterAName, true)

	// Update endpoint with incorrect port
	var data interface{}
	endpointUUID := fmt.Sprintf("endpoint_%s_neutron_uuid", clusterAName)
	endpoint := map[string]interface{}{"uuid": endpointUUID,
		"public_url":  "http://127.0.0.1",
		"private_url": "http://127.0.0.1",
	}
	data = map[string]interface{}{"endpoint": endpoint}
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_neutron_uuid", clusterAName)
		_, err := client.Update(url, &data, &response)
		assert.NoError(t, err, "failed to update neutron endpoint port")
		break
	}

	APIServer.ForceProxyUpdate()

	// verify proxy (expected to fail as the port is incorrect)
	verifyProxyAndTestIt(t, testScenario, clusterAName, false)

	// Delete the neutron endpoint
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_neutron_uuid", clusterAName)
		_, err := client.Delete(url, &response)
		assert.NoError(t, err, "failed to delete neutron endpoint")
		break
	}

	// Re create the neutron endpoint
	endpoint = map[string]interface{}{"uuid": endpointUUID,
		"public_url":  clusterANeutronPublic.URL,
		"private_url": clusterANeutronPrivate.URL,
		"parent_type": "contrail-cluster",
		"parent_uuid": clusterAName + "_uuid",
		"name":        "neutron",
	}
	data = map[string]interface{}{"endpoint": endpoint}
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoints")
		_, err := client.Create(url, &data, &response)
		assert.NoError(t, err, "failed to re-create neutron endpoint port")
		break
	}

	APIServer.ForceProxyUpdate()

	// verify proxy
	verifyProxyAndTestIt(t, testScenario, clusterAName, true)
}

// TestProxyEndpointWithSleep tests the first part of TestProxyEndpoint,
// but verifies that endpoint updates are triggered every 2 seconds.
// TODO: Remove this test when proxyService switches to using events instead of Ticker.
func TestProxyEndpointWithSleep(t *testing.T) {
	// Create a cluster and its neutron endpoint
	clusterAName := "clusterA"
	testScenario, clusterANeutronPublic, clusterANeutronPrivate, cleanup1 := runEndpointTest(
		t, clusterAName)
	defer cleanup1()
	// remove tempfile after test
	defer clusterANeutronPrivate.Close()
	defer clusterANeutronPublic.Close()

	// wait for proxy endpoints to update
	time.Sleep(2 * time.Second)

	verifyProxyAndTestIt(t, testScenario, clusterAName, true)

	// create one more cluster/neutron endpoint for new cluster
	clusterBName := "clusterB"
	testScenario, neutronPublic, neutronPrivate, cleanup2 := runEndpointTest(
		t, clusterBName)
	defer cleanup2()
	// remove tempfile after test
	defer neutronPrivate.Close()
	defer neutronPublic.Close()

	// wait for proxy endpoints to update
	time.Sleep(2 * time.Second)

	// verify new proxies
	verifyProxyAndTestIt(t, testScenario, clusterBName, true)

	// verify existing proxies, make sure the proxy prefix is updated with cluster id
	verifyProxyAndTestIt(t, testScenario, clusterAName, true)
}

func verifyProxyAndTestIt(t *testing.T, scenario *TestScenario, clusterName string, expectTrueAssertion bool) {
	urlPub := "/proxy/" + clusterName + "_uuid/neutron/ports"
	okPub := verifyProxy(t, scenario, urlPub, clusterName, publicPortList)

	urlPriv := "/proxy/" + clusterName + "_uuid/neutron/private/ports"
	okPriv := verifyProxy(t, scenario, urlPriv, clusterName, privatePortList)

	if expectTrueAssertion {
		assert.True(t, okPub, "failed to proxy %s", urlPub)
		assert.True(t, okPriv, "failed to proxy %s", urlPriv)
	} else {
		assert.False(t, okPub, "proxy %s expected to fail", urlPub)
		assert.False(t, okPriv, "proxy %s expected to fail", urlPriv)
	}
}

func TestKeystoneEndpoint(t *testing.T) {
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPrivate := MockServerWithKeystone("", keystoneAuthURL)
	defer ksPrivate.Close()

	ksPublic := MockServerWithKeystone("", keystoneAuthURL)
	defer ksPublic.Close()

	clusterName := "clusterC"
	context := pongo2.Context{
		"cluster_name":  clusterName,
		"endpoint_name": "keystone",
		"private_url":   ksPrivate.URL,
		"public_url":    ksPublic.URL,
	}

	var testScenario TestScenario
	err := LoadTestScenario(&testScenario, testEndpointFile, context)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup := RunDirtyTestScenario(t, &testScenario)
	defer cleanup()

	APIServer.ForceProxyUpdate()

	// Login to new remote keystone
	for _, client := range testScenario.Clients {
		err = client.Login()
		assert.NoError(t, err, "client failed to login remote keystone")
	}
	// verify auth (remote keystone)
	err = verifyKeystoneEndpoint(&testScenario, false)
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
	APIServer.ForceProxyUpdate()

	// Login to new local keystone
	for _, client := range testScenario.Clients {
		err = client.Login()
		assert.NoError(t, err, "client failed to login local keystone")
	}
	// verify auth (local keystone)
	err = verifyKeystoneEndpoint(&testScenario, false)
	assert.NoError(t, err,
		"failed to validate token with local keystone after endpoint delete")

	// Recreate endpoint
	context = pongo2.Context{
		"cluster_name":  clusterName,
		"endpoint_name": "keystone",
		"public_url":    ksPublic.URL,
	}
	err = LoadTestScenario(&testScenario, testEndpointFile, context)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup = RunDirtyTestScenario(t, &testScenario)
	defer cleanup()

	APIServer.ForceProxyUpdate()

	// Login to new remote keystone
	for _, client := range testScenario.Clients {
		err = client.Login()
		assert.NoError(t, err, "client failed to login remote keystone")
	}
	// verify auth (remote keystone)
	err = verifyKeystoneEndpoint(&testScenario, true)
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

	APIServer.ForceProxyUpdate()
}
