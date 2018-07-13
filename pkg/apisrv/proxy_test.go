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

func runEndpointTest(t *testing.T, clusterName, endpointName string) (*TestScenario,
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

	var manage_parent = true
	if endpointName == "neutron2" {
		manage_parent = false
	}
	context := pongo2.Context{
		"cluster_name":    clusterName,
		"endpoint_name":   fmt.Sprintf("%s_%s", clusterName, endpointName),
		"endpoint_prefix": "neutron",
		"private_url":     neutronPrivate.URL,
		"public_url":      neutronPublic.URL,
		"manage_parent":   manage_parent,
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
	// Create a cluster and its neutron endpoints(multiple)
	clusterAName := "clusterA"
	testScenario, clusterANeutronPublic, clusterANeutronPrivate, cleanup1 := runEndpointTest(
		t, clusterAName, "neutron1")
	testScenario, clusterANeutron2Public, clusterANeutron2Private, cleanup2 := runEndpointTest(
		t, clusterAName, "neutron2")
	defer cleanup1()
	defer cleanup2()
	// remove tempfile after test
	defer clusterANeutronPrivate.Close()
	defer clusterANeutronPublic.Close()
	defer clusterANeutron2Private.Close()
	defer clusterANeutron2Public.Close()

	APIServer.ForceProxyUpdate()

	// verify proxies
	url := "/proxy/" + clusterAName + "_uuid/neutron/ports"
	ok := verifyProxy(t, testScenario, url, clusterAName, publicPortList)
	assert.True(t, ok, "failed to proxy %s", url)
	url = "/proxy/" + clusterAName + "_uuid/neutron/private/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, privatePortList)
	assert.True(t, ok, "failed to proxy %s", url)

	// create one more cluster/neutron endpoint for new cluster
	clusterBName := "clusterB"
	testScenario, neutronPublic, neutronPrivate, cleanup3 := runEndpointTest(
		t, clusterBName, "neutron1")
	defer cleanup3()
	// remove tempfile after test
	defer neutronPrivate.Close()
	defer neutronPublic.Close()

	APIServer.ForceProxyUpdate()

	// verify new proxies
	url = "/proxy/" + clusterBName + "_uuid/neutron/ports"
	ok = verifyProxy(t, testScenario, url, clusterBName, publicPortList)
	assert.True(t, ok, "failed to proxy %s", url)
	url = "/proxy/" + clusterBName + "_uuid/neutron/private/ports"
	ok = verifyProxy(t, testScenario, url, clusterBName, privatePortList)
	assert.True(t, ok, "failed to proxy %s", url)

	// verify existing proxies, make sure the proxy prefix is updated with cluster id
	url = "/proxy/" + clusterAName + "_uuid/neutron/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, publicPortList)
	assert.True(t, ok, "failed to proxy %s", url)
	url = "/proxy/" + clusterAName + "_uuid/neutron/private/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, privatePortList)
	assert.True(t, ok, "failed to proxy %s", url)

	// Update neutron endpoints with incorrect port
	for _, neutron := range []string{"neutron1", "neutron2"} {

		var data interface{}
		endpointUUID := fmt.Sprintf("endpoint_%s_%s_uuid", clusterAName, neutron)
		endpoint := map[string]interface{}{"uuid": endpointUUID,
			"public_url":  "http://127.0.0.1",
			"private_url": "http://127.0.0.1",
		}
		data = map[string]interface{}{"endpoint": endpoint}
		for _, client := range testScenario.Clients {
			var response map[string]interface{}
			url = fmt.Sprintf("/endpoint/endpoint_%s_%s_uuid", clusterAName, neutron)
			_, err := client.Update(url, &data, &response)
			assert.NoError(t, err, "failed to update neutron endpoint port")
			break
		}
	}

	APIServer.ForceProxyUpdate()

	// verify proxy (expected to fail)
	url = "/proxy/" + clusterAName + "_uuid/neutron/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, publicPortList)
	assert.False(t, ok, "proxy %s expected to fail", url)
	url = "/proxy/" + clusterAName + "_uuid/neutron/private/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, privatePortList)
	assert.False(t, ok, "proxy %s expected to fail", url)

	// Delete the neutron endpoint
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url = fmt.Sprintf("/endpoint/endpoint_%s_neutron1_uuid", clusterAName)
		_, err := client.Delete(url, &response)
		assert.NoError(t, err, "failed to delete neutron1 endpoint")
		url = fmt.Sprintf("/endpoint/endpoint_%s_neutron2_uuid", clusterAName)
		_, err = client.Delete(url, &response)
		assert.NoError(t, err, "failed to delete neutron2 endpoint")
		break
	}

	// Re create the neutron endpoint
	endpointUUID := fmt.Sprintf("endpoint_%s_neutron1_uuid", clusterAName)
	endpoint := map[string]interface{}{"uuid": endpointUUID,
		"public_url":  clusterANeutronPublic.URL,
		"private_url": clusterANeutronPrivate.URL,
		"parent_type": "contrail-cluster",
		"parent_uuid": clusterAName + "_uuid",
		"name":        clusterAName + "_neutron1",
		"prefix":      "neutron",
	}
	data := map[string]interface{}{"endpoint": endpoint}
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url = fmt.Sprintf("/endpoints")
		_, err := client.Create(url, &data, &response)
		assert.NoError(t, err, "failed to re-create neutron1 endpoint port")
		break
	}

	APIServer.ForceProxyUpdate()

	// verify proxy
	url = "/proxy/" + clusterAName + "_uuid/neutron/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, publicPortList)
	assert.True(t, ok, "failed to proxy %s", url)
	url = "/proxy/" + clusterAName + "_uuid/neutron/private/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, privatePortList)
	assert.True(t, ok, "failed to proxy %s", url)
}

// TestProxyEndpointWithSleep tests the first part of TestProxyEndpoint,
// but verifies that endpoint updates are triggered every 2 seconds.
// TODO: Remove this test when proxyService switches to using events instead of Ticker.
func TestProxyEndpointWithSleep(t *testing.T) {
	// Create a cluster and its neutron endpoint
	clusterAName := "clusterA"
	testScenario, clusterANeutronPublic, clusterANeutronPrivate, cleanup1 := runEndpointTest(
		t, clusterAName, "neutron1")
	defer cleanup1()
	// remove tempfile after test
	defer clusterANeutronPrivate.Close()
	defer clusterANeutronPublic.Close()

	time.Sleep(2 * time.Second)

	// verify proxies
	url := "/proxy/" + clusterAName + "_uuid/neutron/ports"
	ok := verifyProxy(t, testScenario, url, clusterAName, publicPortList)
	assert.True(t, ok, "failed to proxy %s", url)
	url = "/proxy/" + clusterAName + "_uuid/neutron/private/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, privatePortList)
	assert.True(t, ok, "failed to proxy %s", url)

	// create one more cluster/neutron endpoint for new cluster
	clusterBName := "clusterB"
	testScenario, neutronPublic, neutronPrivate, cleanup2 := runEndpointTest(
		t, clusterBName, "neutron1")
	defer cleanup2()
	// remove tempfile after test
	defer neutronPrivate.Close()
	defer neutronPublic.Close()

	time.Sleep(2 * time.Second)

	// verify new proxies
	url = "/proxy/" + clusterBName + "_uuid/neutron/ports"
	ok = verifyProxy(t, testScenario, url, clusterBName, publicPortList)
	assert.True(t, ok, "failed to proxy %s", url)
	url = "/proxy/" + clusterBName + "_uuid/neutron/private/ports"
	ok = verifyProxy(t, testScenario, url, clusterBName, privatePortList)
	assert.True(t, ok, "failed to proxy %s", url)

	// verify existing proxies, make sure the proxy prefix is updated with cluster id
	url = "/proxy/" + clusterAName + "_uuid/neutron/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, publicPortList)
	assert.True(t, ok, "failed to proxy %s", url)
	url = "/proxy/" + clusterAName + "_uuid/neutron/private/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, privatePortList)
	assert.True(t, ok, "failed to proxy %s", url)
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
		"endpoint_name": clusterName + "_keystone",
		"private_url":   ksPrivate.URL,
		"public_url":    ksPublic.URL,
		"manage_parent": true,
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
		"endpoint_name": clusterName + "_keystone",
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
