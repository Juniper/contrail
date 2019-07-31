package apisrv_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/websocket"
)

const (
	key              = "name"
	privatePortList  = "private_port_list"
	publicPortList   = "public_port_list"
	testEndpointFile = "./test_data/test_endpoint.tmpl"
	xClusterIDKey    = "X-Cluster-ID"
)

func TestProxyEndpoint(t *testing.T) {
	ctx := context.Background()
	// Create a cluster and its neutron endpoints(multiple)
	clusterAName := t.Name() + "_clusterA"
	testScenario, clusterANeutronPublic, clusterANeutronPrivate, cleanup1 := runEndpointTest(
		t, clusterAName, "neutron1")
	testScenario, clusterANeutron2Public, clusterANeutron2Private, cleanup2 := runEndpointTest(
		t, clusterAName, "neutron2")
	defer cleanup1()
	defer cleanup2()
	defer clusterANeutronPrivate.Close()
	defer clusterANeutronPublic.Close()
	defer clusterANeutron2Private.Close()
	defer clusterANeutron2Public.Close()

	server.ForceProxyUpdate()

	verifyProxies(ctx, t, testScenario, clusterAName, true)

	// create one more cluster/neutron endpoint for new cluster
	clusterBName := t.Name() + "_clusterB"
	testScenario, neutronPublic, neutronPrivate, cleanup3 := runEndpointTest(
		t, clusterBName, "neutron1")
	defer cleanup3()
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
	clusterAName := t.Name() + "_clusterA"
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
	clusterBName := t.Name() + "_clusterB"
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
			clusterID := c.Request().Header.Get(xClusterIDKey)
			if clusterID != clusterName+"_uuid" {
				return c.JSON(http.StatusBadRequest,
					"clusterID not found in header")
			}
			return c.JSON(http.StatusOK,
				&mockPortsResponse{Name: clusterName + publicPortList})
		}),
	}
	neutronPublic := mockServer(routes)

	manageParent := true
	if endpointName == "neutron2" {
		manageParent = false
	}

	ts, err := integration.LoadTest(testEndpointFile, pongo2.Context{
		"cluster_name":    clusterName,
		"endpoint_name":   fmt.Sprintf("%s_%s", clusterName, endpointName),
		"endpoint_prefix": "neutron",
		"private_url":     neutronPrivate.URL,
		"public_url":      neutronPublic.URL,
		"manage_parent":   manageParent,
	})
	require.NoError(t, err, "failed to load test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)

	return ts, neutronPublic, neutronPrivate, cleanup
}

type mockPortsResponse struct {
	Name string `json:"name"`
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

func verifyProxy(
	ctx context.Context,
	t *testing.T,
	testScenario *integration.TestScenario,
	url string,
	clusterName string,
	expected string,
) bool {
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		_, err := client.Read(ctx, url, &response)
		//assert.NoError(t, err, fmt.Sprint("client endpoint:", client.Endpoint), "url:", url) // TODO

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

// TODO(Daniel): taken from previous CRs - verify this is needed
func TestProxyEndpointWithRetry(t *testing.T) {
	t.Skip("Hoge")

	// Create a cluster and its config endpoints(multiple)
	// unavailable endpoint
	clusterName := "dead_cluster"
	routes := map[string]interface{}{
		"/virtual-networks": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusServiceUnavailable,
				&mockPortsResponse{Name: clusterName + "_serviceUnavailable"})
		}),
	}
	api1 := mockServer(routes)
	defer api1.Close()

	pCtx := pongo2.Context{
		"cluster_name":    clusterName,
		"endpoint_name":   fmt.Sprintf("%s_config1", clusterName),
		"endpoint_prefix": "config",
		"private_url":     api1.URL,
		"public_url":      api1.URL,
		"manage_parent":   true,
	}

	ts, err := integration.LoadTest(testEndpointFile, pCtx)
	assert.NoError(t, err, "failed to load test data")
	cleanup1 := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup1()

	// bad gateway endpoint
	routes = map[string]interface{}{
		"/virtual-networks": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusBadGateway,
				&mockPortsResponse{Name: clusterName + "_badGateway"})
		}),
	}
	api2 := mockServer(routes)
	defer api2.Close()
	pCtx = pongo2.Context{
		"cluster_name":    clusterName,
		"endpoint_name":   fmt.Sprintf("%s_config2", clusterName),
		"endpoint_prefix": "config",
		"private_url":     api2.URL,
		"public_url":      api2.URL,
		"manage_parent":   false,
	}

	ts, err = integration.LoadTest(testEndpointFile, pCtx)
	assert.NoError(t, err, "failed to load test data")
	cleanup2 := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup2()

	// endpoint with gateway timeout
	routes = map[string]interface{}{
		"/virtual-networks": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK,
				&mockPortsResponse{Name: clusterName})
		}),
	}
	api3 := mockServer(routes)
	defer api3.Close()
	pCtx = pongo2.Context{
		"cluster_name":    clusterName,
		"endpoint_name":   fmt.Sprintf("%s_config3", clusterName),
		"endpoint_prefix": "config",
		"private_url":     api3.URL,
		"public_url":      api3.URL,
		"manage_parent":   false,
	}

	ts, err = integration.LoadTest(testEndpointFile, pCtx)
	assert.NoError(t, err, "failed to load test data")
	cleanup3 := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup3()

	server.ForceProxyUpdate()

	// verify proxies, expected to fail as all the endpoints are down
	url := "/proxy/" + clusterName + "_uuid/config/virtual-networks"
	for _, client := range ts.Clients {
		var response map[string]interface{}
		_, err := client.Read(context.Background(), url, &response) // FIXME sometimes returns 503 for "name":"dead_cluster_badGateway"
		assert.NoError(t, err, fmt.Sprintf("Reading: %s, Response: %s", url, err))

		testutil.AssertEqual(t,
			map[string]interface{}{key: clusterName},
			response,
		)
	}
}

func mockServer(routes map[string]interface{}) *httptest.Server {
	// Echo instance
	e := echo.New()

	// Routes
	for route, handler := range routes {
		e.GET(route, handler.(echo.HandlerFunc))
	}

	return httptest.NewServer(e)
}

func TestKeystoneEndpoint(t *testing.T) {
	ctx := context.Background()
	keystoneAuthURL := viper.GetString("keystone.authurl")

	clusterCName := t.Name() + "_clusterC"
	clusterCUser := clusterCName + "_admin"
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterCUser, clusterCUser)
	defer ksPrivate.Close()

	ksPublic := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterCUser, clusterCUser)
	defer ksPublic.Close()

	ts, err := integration.LoadTest(
		testEndpointFile,
		pongo2.Context{
			"cluster_name":  clusterCName,
			"endpoint_name": clusterCName + "_keystone",
			"private_url":   ksPrivate.URL,
			"public_url":    ksPublic.URL,
			"manage_parent": true,
			"admin_user":    clusterCUser,
		})
	require.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Login to new remote keystone
	for _, client := range ts.Clients {
		_, err = client.Login(ctx)
		assert.NoError(t, err, "client failed to login remote keystone")
	}
	// verify auth (remote keystone)
	err = verifyKeystoneEndpoint(ctx, ts, false)
	assert.NoError(t, err,
		"failed to validate token with remote keystone")

	// Delete endpoint test
	for _, client := range ts.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterCName)
		_, err = client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete keystone endpoint")
		break
	}
	server.ForceProxyUpdate()

	// Login to new local keystone
	for _, client := range ts.Clients {
		_, err = client.Login(ctx)
		assert.NoError(t, err, "client failed to login local keystone")
	}
	// verify auth (local keystone)
	err = verifyKeystoneEndpoint(ctx, ts, false)
	assert.NoError(t, err,
		"failed to validate token with local keystone after endpoint delete")

	// Recreate endpoint
	ts, err = integration.LoadTest(
		testEndpointFile,
		pongo2.Context{
			"cluster_name":  clusterCName,
			"endpoint_name": clusterCName + "_keystone",
			"public_url":    ksPublic.URL,
			"admin_user":    clusterCName,
		})
	require.NoError(t, err, "failed to load endpoint create test data")
	cleanup = integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Login to new remote keystone
	for _, client := range ts.Clients {
		_, err = client.Login(ctx)
		assert.NoError(t, err, "client failed to login remote keystone")
	}
	// verify auth (remote keystone)
	err = verifyKeystoneEndpoint(ctx, ts, true)
	assert.NoError(t, err,
		"failed to validate token with remote keystone after endpoint re-create")

	// Cleanup endpoint test
	for _, client := range ts.Clients {
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterCName)
		_, err = client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete keystone endpoint")
		break
	}

	server.ForceProxyUpdate()
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

func TestMultipleClusterKeystoneEndpoint(t *testing.T) {
	ctx := context.Background()
	keystoneAuthURL := viper.GetString("keystone.authurl")
	clusterCName := t.Name() + "_clusterC"
	clusterCUser := clusterCName + "_admin"
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterCUser, clusterCUser)
	defer ksPrivate.Close()

	ksPublic := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterCUser, clusterCUser)
	defer ksPublic.Close()

	ts, err := integration.LoadTest(
		testEndpointFile,
		pongo2.Context{
			"cluster_name":    clusterCName,
			"endpoint_name":   clusterCName + "_keystone",
			"endpoint_prefix": "keystone",
			"private_url":     ksPrivate.URL,
			"public_url":      ksPublic.URL,
			"manage_parent":   true,
			"admin_user":      clusterCUser,
		})
	require.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Create one more cluster's keystone endpoint
	keystoneAuthURL = viper.GetString("keystone.authurl")
	clusterDName := t.Name() + "_clusterD"
	clusterDUser := clusterDName + "_admin"
	ksPrivateClusterD := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterDUser, clusterDUser)
	defer ksPrivateClusterD.Close()

	ksPublicClusterD := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterDUser, clusterDUser)
	defer ksPublicClusterD.Close()

	pContext := pongo2.Context{
		"cluster_name":    clusterDName,
		"endpoint_name":   clusterDName + "_keystone",
		"endpoint_prefix": "keystone",
		"private_url":     ksPrivateClusterD.URL,
		"public_url":      ksPublicClusterD.URL,
		"manage_parent":   true,
		"admin_user":      clusterDUser,
	}

	ts, err = integration.LoadTest(testEndpointFile, pContext)
	require.NoError(t, err, "failed to load endpoint create test data")
	cleanupClusterD := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanupClusterD()

	server.ForceProxyUpdate()

	// Login to new remote keystone(success)
	// when multiple cluster endpoints are present
	// auth middleware should find the keystone endpoint
	// with X-Cluster-ID in the header
	clientTS, err := integration.LoadTest(testEndpointFile, pContext)
	require.NoError(t, err, "failed to load endpoint create test data")

	clients := integration.PrepareClients(ctx, t, clientTS, server)

	for _, client := range clients {
		ctx = auth.WithXClusterID(ctx, clusterDName+"_uuid")
		client.ID = clusterDUser
		client.Password = clusterDUser
		client.Scope = nil
		_, err = client.Login(ctx)
		assert.NoError(t, err, "client failed to login remote keystone")
	}

	// Login to new remote keystone(failure)
	// when multiple cluster endpoints are present
	// auth middleware cannot not find keystone endpoint
	// without X-Cluster-ID in the header
	for _, client := range clients {
		ctx = context.Background()
		client.ID = clusterDUser
		client.Password = clusterDUser
		client.Scope = nil
		_, err = client.Login(ctx)
		assert.Error(t, err, "client logged in to remote keystone unexpectedly")
	}

	// Delete the clusterD's keystone endpoint
	for _, client := range ts.Clients {
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterDName)
		var response map[string]interface{}
		_, err := client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete clusterD's keystone endpoint")
		break
	}
	server.ForceProxyUpdate()

	// Delete the clusterC's keystone endpoint
	for _, client := range ts.Clients {
		ctx = context.Background()
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterCName)
		_, err := client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete clusterC's keystone endpoint")
		break
	}
	server.ForceProxyUpdate()
}

func TestWebsocketEndpoint(t *testing.T) {
	clusterName := t.Name() + "_cluster"
	endpointPrefix := "websocket-server"

	target := echoWebsocketServer(t)
	target.Start()
	defer target.Close()
	cleanup := createEndpoint(t, target, clusterName, endpointPrefix)
	defer cleanup()

	wsURLBase := strings.ReplaceAll(server.URL(), "https://", "wss://")
	url := fmt.Sprintf("%s/proxy/%s_uuid/%s", wsURLBase, clusterName, endpointPrefix)

	config, err := websocket.NewConfig(url, "http://localhost/")
	assert.NoError(t, err, "failed to create websocket config from proxy URL")
	config.TlsConfig = &tls.Config{InsecureSkipVerify: true}
	ws, err := websocket.DialConfig(config)
	assert.NoError(t, err, "failed to connect to a websocket endpoint through the proxy")
	defer func() {
		if err = ws.Close(); err != nil {
			t.Error("Failed to close websocket: ", err)
		}
	}()

	sentMsg := []byte("test message")
	_, err = ws.Write(sentMsg)
	assert.NoError(t, err, "failed to send a message through the proxied websocket")

	receivedMsg := make([]byte, 100)
	n, err := ws.Read(receivedMsg)
	assert.NoError(t, err, "failed to receive a message through the proxied websocket")
	assert.Equal(t, sentMsg, receivedMsg[:n])
}

func echoWebsocketServer(t *testing.T) *httptest.Server {
	return integration.NewWellKnownServer("", websocket.Handler(func(ws *websocket.Conn) {
		if _, err := io.Copy(ws, ws); err != nil {
			t.Error("Failed to echo the message back to the client: ", err)
		}
	}))
}

func createEndpoint(t *testing.T, target *httptest.Server, clusterName, endpointPrefix string) (cleanup func()) {
	ts, err := integration.LoadTest(
		testEndpointFile,
		pongo2.Context{
			"cluster_name":    clusterName,
			"endpoint_name":   clusterName + "_" + endpointPrefix,
			"endpoint_prefix": endpointPrefix,
			"private_url":     target.URL,
			"public_url":      target.URL,
			"manage_parent":   true,
			"admin_user":      clusterName + "_admin",
		})
	require.NoError(t, err, "failed to load endpoint create test data")
	endpointCleanup := integration.RunDirtyTestScenario(t, ts, server)

	server.ForceProxyUpdate()

	return func() {
		endpointCleanup()
		server.ForceProxyUpdate()
	}
}
