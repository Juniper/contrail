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

func runEndpointTest(t *testing.T, clusterName string,
	extraTasks bool) (*TestScenario, *httptest.Server, *httptest.Server, func()) {
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
		t, clusterAName, true)
	defer cleanup1()
	// remove tempfile after test
	defer clusterANeutronPrivate.Close()
	defer clusterANeutronPublic.Close()

	// Wait a sec for the dynamic proxy to be created/updated
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
		t, clusterBName, false)
	defer cleanup2()
	// remove tempfile after test
	defer neutronPrivate.Close()
	defer neutronPublic.Close()
	// Wait a sec for the dynamic proxy to be created/updated
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
		url = fmt.Sprintf("/endpoint/endpoint_%s_neutron_uuid", clusterAName)
		_, err := client.Update(url, &data, &response)
		assert.NoError(t, err, "failed to update neutron endpoint port")
		break
	}

	// Wait 2 sec for the dynamic proxy to update endpointstore
	time.Sleep(2 * time.Second)
	// verify proxy (expected to fail as the port is incorrect)
	url = "/proxy/" + clusterAName + "_uuid/neutron/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, publicPortList)
	assert.False(t, ok, "proxy %s expected to fail", url)
	url = "/proxy/" + clusterAName + "_uuid/neutron/private/ports"
	ok = verifyProxy(t, testScenario, url, clusterAName, privatePortList)
	assert.False(t, ok, "proxy %s expected to fail", url)

	// Delete the neutron endpoint
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url = fmt.Sprintf("/endpoint/endpoint_%s_neutron_uuid", clusterAName)
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
		url = fmt.Sprintf("/endpoints")
		_, err := client.Create(url, &data, &response)
		assert.NoError(t, err, "failed to re-create neutron endpoint port")
		break
	}
	// Wait a sec for the dynamic proxy to be created/updated
	time.Sleep(2 * time.Second)
	// verify proxy
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
		"extra_tasks":   true,
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

	// Wait a sec for the dynamic proxy to be created/updated
	time.Sleep(2 * time.Second)

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
	// Wait a sec for the dynamic proxy to be deleted
	time.Sleep(2 * time.Second)
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
		"extra_tasks":   true,
		"cluster_name":  clusterName,
		"endpoint_name": "keystone",
		"public_url":    ksPublic.URL,
	}
	err = LoadTestScenario(&testScenario, testEndpointFile, context)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup = RunDirtyTestScenario(t, &testScenario)
	defer cleanup()
	// Wait a sec for the dynamic proxy to be created
	time.Sleep(2 * time.Second)
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
	// Wait a sec for the dynamic proxy to be deleted
	time.Sleep(2 * time.Second)
}
