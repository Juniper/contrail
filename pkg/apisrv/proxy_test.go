package apisrv_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
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

func runEndpointTest(t *testing.T, clusterName, endpointName string) (*integration.TestScenario,
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

	manageParent := true
	if endpointName == "neutron2" {
		manageParent = false
	}
	context := pongo2.Context{
		"cluster_name":    clusterName,
		"endpoint_name":   fmt.Sprintf("%s_%s", clusterName, endpointName),
		"endpoint_prefix": "neutron",
		"private_url":     neutronPrivate.URL,
		"public_url":      neutronPublic.URL,
		"manage_parent":   manageParent,
	}

	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testEndpointFile, context)
	assert.NoError(t, err, "failed to load test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)

	return &testScenario, neutronPublic, neutronPrivate, cleanup
}

func verifyProxy(ctx context.Context, t *testing.T, testScenario *integration.TestScenario, url string,
	clusterName string, expected string) bool {
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		_, err := client.Read(ctx, url, &response)
		if err != nil {
			fmt.Printf("Reading: %s, Response: %s", url, err)
			return false
		}
		ok := testutil.AssertEqual(t,
			map[string]interface{}{key: clusterName + expected},
			response,
			fmt.Sprintf("Unexpected Response: %s", response))
		if !ok {
			return ok
		}
	}
	return true
}

func verifyKeystoneEndpoint(ctx context.Context, testScenario *integration.TestScenario, testInvalidUser bool) error {
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		_, err := client.Read(ctx, "/keystone/v3/auth/tokens", &response)
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
	ctx := context.Background()
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

	server.ForceProxyUpdate()

	// verify proxies
	verifyProxies(ctx, t, testScenario, clusterAName, true)

	// create one more cluster/neutron endpoint for new cluster
	clusterBName := "clusterB"
	testScenario, neutronPublic, neutronPrivate, cleanup3 := runEndpointTest(
		t, clusterBName, "neutron1")
	defer cleanup3()
	// remove tempfile after test
	defer neutronPrivate.Close()
	defer neutronPublic.Close()

	server.ForceProxyUpdate()

	// verify new proxies
	verifyProxies(ctx, t, testScenario, clusterBName, true)

	// verify existing proxies, make sure the proxy prefix is updated with cluster id
	verifyProxies(ctx, t, testScenario, clusterAName, true)

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
			url := fmt.Sprintf("/endpoint/endpoint_%s_%s_uuid", clusterAName, neutron)
			_, err := client.Update(ctx, url, &data, &response)
			assert.NoError(t, err, "failed to update neutron endpoint port")
			break
		}
	}

	server.ForceProxyUpdate()

	// verify proxy (expected to fail as the port is incorrect)
	verifyProxies(ctx, t, testScenario, clusterAName, false)

	// Delete the neutron endpoint
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_neutron1_uuid", clusterAName)
		_, err := client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete neutron1 endpoint")
		url = fmt.Sprintf("/endpoint/endpoint_%s_neutron2_uuid", clusterAName)
		_, err = client.Delete(ctx, url, &response)
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
		url := fmt.Sprintf("/endpoints")
		_, err := client.Create(ctx, url, &data, &response)
		assert.NoError(t, err, "failed to re-create neutron1 endpoint port")
		break
	}

	server.ForceProxyUpdate()

	// verify proxy
	verifyProxies(ctx, t, testScenario, clusterAName, true)
}

// TestProxyEndpointWithSleep tests the first part of TestProxyEndpoint,
// but verifies that endpoint updates are triggered every 2 seconds.
// TODO: Remove this test when proxyService switches to using events instead of Ticker.
func TestProxyEndpointWithSleep(t *testing.T) {
	ctx := context.Background()
	// Create a cluster and its neutron endpoint
	clusterAName := "clusterA"
	testScenario, clusterANeutronPublic, clusterANeutronPrivate, cleanup1 := runEndpointTest(
		t, clusterAName, "neutron1")
	defer cleanup1()
	// remove tempfile after test
	defer clusterANeutronPrivate.Close()
	defer clusterANeutronPublic.Close()

	// wait for proxy endpoints to update
	time.Sleep(2 * time.Second)

	verifyProxies(ctx, t, testScenario, clusterAName, true)

	// create one more cluster/neutron endpoint for new cluster
	clusterBName := "clusterB"
	testScenario, neutronPublic, neutronPrivate, cleanup2 := runEndpointTest(
		t, clusterBName, "neutron1")
	defer cleanup2()
	// remove tempfile after test
	defer neutronPrivate.Close()
	defer neutronPublic.Close()

	// wait for proxy endpoints to update
	time.Sleep(2 * time.Second)

	// verify new proxies
	verifyProxies(ctx, t, testScenario, clusterBName, true)

	// verify existing proxies, make sure the proxy prefix is updated with cluster id
	verifyProxies(ctx, t, testScenario, clusterAName, true)
}

func verifyProxies(
	ctx context.Context, t *testing.T, scenario *integration.TestScenario, clusterName string, isSuccessful bool,
) {
	url := "/proxy/" + clusterName + "_uuid/neutron/ports"
	ok := verifyProxy(ctx, t, scenario, url, clusterName, publicPortList)
	assert.Equal(t, ok, isSuccessful, "failed to proxy %s", url)

	url = "/proxy/" + clusterName + "_uuid/neutron/private/ports"
	ok = verifyProxy(ctx, t, scenario, url, clusterName, privatePortList)
	assert.Equal(t, ok, isSuccessful, "failed to proxy %s", url)
}

func TestKeystoneEndpoint(t *testing.T) {
	ctx := context.Background()
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPrivate := integration.MockServerWithKeystone("", keystoneAuthURL)
	defer ksPrivate.Close()

	ksPublic := integration.MockServerWithKeystone("", keystoneAuthURL)
	defer ksPublic.Close()

	clusterName := "clusterC"
	context := pongo2.Context{
		"cluster_name":  clusterName,
		"endpoint_name": clusterName + "_keystone",
		"private_url":   ksPrivate.URL,
		"public_url":    ksPublic.URL,
		"manage_parent": true,
	}

	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testEndpointFile, context)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Login to new remote keystone
	for _, client := range testScenario.Clients {
		_, err = client.Login(ctx)
		assert.NoError(t, err, "client failed to login remote keystone")
	}
	// verify auth (remote keystone)
	err = verifyKeystoneEndpoint(ctx, &testScenario, false)
	assert.NoError(t, err,
		"failed to validate token with remote keystone")

	// Delete endpoint test
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterName)
		_, err = client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete keystone endpoint")
		break
	}
	server.ForceProxyUpdate()

	// Login to new local keystone
	for _, client := range testScenario.Clients {
		_, err = client.Login(ctx)
		assert.NoError(t, err, "client failed to login local keystone")
	}
	// verify auth (local keystone)
	err = verifyKeystoneEndpoint(ctx, &testScenario, false)
	assert.NoError(t, err,
		"failed to validate token with local keystone after endpoint delete")

	// Recreate endpoint
	context = pongo2.Context{
		"cluster_name":  clusterName,
		"endpoint_name": clusterName + "_keystone",
		"public_url":    ksPublic.URL,
	}
	err = integration.LoadTestScenario(&testScenario, testEndpointFile, context)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup = integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Login to new remote keystone
	for _, client := range testScenario.Clients {
		_, err = client.Login(ctx)
		assert.NoError(t, err, "client failed to login remote keystone")
	}
	// verify auth (remote keystone)
	err = verifyKeystoneEndpoint(ctx, &testScenario, true)
	assert.NoError(t, err,
		"failed to validate token with remote keystone after endpoint re-create")

	// Cleanup endpoint test
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterName)
		_, err = client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete keystone endpoint")
		break
	}

	server.ForceProxyUpdate()
}

func TestMultipleClusterKeystoneEndpoint(t *testing.T) {
	ctx := context.Background()
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPrivate := integration.MockServerWithKeystone("", keystoneAuthURL)
	defer ksPrivate.Close()

	ksPublic := integration.MockServerWithKeystone("", keystoneAuthURL)
	defer ksPublic.Close()

	clusterCName := "clusterC"
	pContext := pongo2.Context{
		"cluster_name":    clusterCName,
		"endpoint_name":   clusterCName + "_keystone",
		"endpoint_prefix": "keystone",
		"private_url":     ksPrivate.URL,
		"public_url":      ksPublic.URL,
		"manage_parent":   true,
	}

	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testEndpointFile, pContext)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Login to new remote keystone(success)
	// when one cluster endpoint is present
	// auth middleware should find the keystone endpoint
	// even without X-Cluster-ID in the header
	for _, client := range testScenario.Clients {
		_, err = client.Login(ctx)
		assert.NoError(t, err, "client failed to login remote keystone")
	}
	// verify auth (remote keystone)
	err = verifyKeystoneEndpoint(ctx, &testScenario, false)
	assert.NoError(t, err,
		"failed to validate token with remote keystone")

	// Create one more cluster's keystone endpoint
	keystoneAuthURL = viper.GetString("keystone.authurl")
	ksPrivateClusterD := integration.MockServerWithKeystone("", keystoneAuthURL)
	defer ksPrivateClusterD.Close()

	ksPublicClusterD := integration.MockServerWithKeystone("", keystoneAuthURL)
	defer ksPublicClusterD.Close()

	clusterDName := "clusterD"
	pContext = pongo2.Context{
		"cluster_name":    clusterDName,
		"endpoint_name":   clusterDName + "_keystone",
		"endpoint_prefix": "keystone",
		"private_url":     ksPrivateClusterD.URL,
		"public_url":      ksPublicClusterD.URL,
		"manage_parent":   true,
	}

	err = integration.LoadTestScenario(&testScenario, testEndpointFile, pContext)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanupClusterD := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanupClusterD()

	server.ForceProxyUpdate()

	// Login to new remote keystone(success)
	// when multiple cluster endpoints are present
	// auth middleware should find the keystone endpoint
	// with X-Cluster-ID in the header
	for _, client := range testScenario.Clients {
		ctx = auth.WithXClusterID(ctx, clusterDName+"_uuid")
		_, err = client.Login(ctx)
		assert.NoError(t, err, "client failed to login remote keystone")
	}
	// verify auth (remote keystone)
	err = verifyKeystoneEndpoint(ctx, &testScenario, false)
	assert.NoError(t, err,
		"failed to validate token with remote keystone")

	// Login to new remote keystone(failure)
	// when multiple cluster endpoints are present
	// auth middleware cannot not find keystone endpoint
	// without X-Cluster-ID in the header
	for _, client := range testScenario.Clients {
		ctx = context.Background()
		_, err = client.Login(ctx)
		assert.Error(t, err, "client logged in to remote keystone unexpectedly")
	}

	// Delete the clusterD's keystone endpoint
	for _, client := range testScenario.Clients {
		ctx = auth.WithXClusterID(ctx, clusterDName+"_uuid")
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterDName)
		_, err := client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete clusterD's keystone endpoint")
		break
	}
	server.ForceProxyUpdate()

	// Delete the clusterC's keystone endpoint
	for _, client := range testScenario.Clients {
		ctx = context.Background()
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterCName)
		_, err := client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete clusterC's keystone endpoint")
		break
	}
	server.ForceProxyUpdate()
}
