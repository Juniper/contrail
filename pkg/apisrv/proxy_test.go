package apisrv_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/endpoint"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/websocket"
)

const (
	neutronEndpointPrefix = "neutron"
	portsPath             = "/ports"
	testEndpointFile      = "./test_data/test_endpoint.tmpl"
)

//////////////////////////////
// Dynamic proxy HTTP tests //
//////////////////////////////

func TestDynamicProxyServiceHTTPSupport(t *testing.T) {
	for _, tt := range []struct {
		name                      string
		synchronizeProxyEndpoints func(s *integration.APIServer)
	}{
		{
			name: "synchronizing proxy endpoints with ForceProxyUpdate",
			synchronizeProxyEndpoints: func(s *integration.APIServer) {
				s.ForceProxyUpdate()
			},
		},
		{
			// TODO: Remove this test when proxyService switches to using events instead of Ticker.
			name: "synchronizing proxy endpoints with sleep",
			synchronizeProxyEndpoints: func(_ *integration.APIServer) {
				time.Sleep(apisrv.ProxySyncInterval)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			testDynamicProxyServiceHTTPSupport(t, tt.synchronizeProxyEndpoints)
		})
	}
}

func testDynamicProxyServiceHTTPSupport(t *testing.T, synchronizeProxyEndpoints func(s *integration.APIServer)) {
	// arrange
	clusterAName, clusterBName := contrailClusterName(t, "A"), contrailClusterName(t, "B")
	const neutron1EndpointName, neutron2EndpointName = "neutron1", "neutron2"
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	cleanupCCA := createContrailCluster(t, hc, clusterAName)
	defer cleanupCCA()

	neutronA1Public, neutronA1Private, closeNeutronA1 := createNeutronServers(clusterAName)
	defer closeNeutronA1()
	cleanupEA1 := createNeutronEndpoint(t, hc, endpointParameters{
		clusterName:  clusterAName,
		endpointName: neutron1EndpointName,
		privateURL:   neutronA1Private.URL,
		publicURL:    neutronA1Public.URL,
	})
	defer cleanupEA1()

	neutronA2Public, neutronA2Private, closeNeutronA2 := createNeutronServers(clusterAName)
	defer closeNeutronA2()
	cleanupEA2 := createNeutronEndpoint(t, hc, endpointParameters{
		clusterName:  clusterAName,
		endpointName: neutron2EndpointName,
		privateURL:   neutronA2Private.URL,
		publicURL:    neutronA2Public.URL,
	})
	defer cleanupEA2()

	// act/assert
	synchronizeProxyEndpoints(server)
	verifyNeutronReadRequests(t, hc, clusterAName)

	cleanupCCB := createContrailCluster(t, hc, clusterBName)
	defer cleanupCCB()

	neutronB1Public, neutronB1Private, closeNeutronB1 := createNeutronServers(clusterBName)
	defer closeNeutronB1()
	cleanupEB1 := createNeutronEndpoint(t, hc, endpointParameters{
		clusterName:  clusterBName,
		endpointName: neutron1EndpointName,
		privateURL:   neutronB1Private.URL,
		publicURL:    neutronB1Public.URL,
	})
	defer cleanupEB1()

	synchronizeProxyEndpoints(server)
	verifyNeutronReadRequests(t, hc, clusterAName)
	verifyNeutronReadRequests(t, hc, clusterBName)

	setIncorrectEndpointURLs(t, hc, clusterAName, neutron1EndpointName)
	setIncorrectEndpointURLs(t, hc, clusterAName, neutron2EndpointName)

	synchronizeProxyEndpoints(server)
	verifyNeutronReadRequestsFail(t, hc, clusterAName)
	verifyNeutronReadRequests(t, hc, clusterBName)

	integration.DeleteEndpoint(t, hc, endpointUUID(clusterAName, neutron1EndpointName))
	integration.DeleteEndpoint(t, hc, endpointUUID(clusterAName, neutron2EndpointName))
	cleanupEA1 = createNeutronEndpoint(t, hc, endpointParameters{
		clusterName:  clusterAName,
		endpointName: neutron1EndpointName,
		privateURL:   neutronA1Private.URL,
		publicURL:    neutronA1Public.URL,
	})
	defer cleanupEA1()

	synchronizeProxyEndpoints(server)
	verifyNeutronReadRequests(t, hc, clusterAName)
	verifyNeutronReadRequests(t, hc, clusterBName)
}

func TestDynamicProxyServiceWithClosedTargetServers(t *testing.T) {
	// arrange
	clusterName := contrailClusterName(t, "")
	const ok1EndpointName, ok2EndpointName = "neutron-ok1", "neutron-ok2"
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	cleanupCC := createContrailCluster(t, hc, clusterName)
	defer cleanupCC()

	ok1Neutron := newNeutronServerStub(http.StatusOK)
	defer ok1Neutron.Close()
	cleanupE1 := createNeutronEndpoint(t, hc, endpointParameters{
		clusterName:  clusterName,
		endpointName: ok1EndpointName,
		privateURL:   ok1Neutron.URL,
		publicURL:    ok1Neutron.URL,
	})
	defer cleanupE1()

	ok2Neutron := newNeutronServerStub(http.StatusOK)
	defer ok2Neutron.Close()
	cleanupE2 := createNeutronEndpoint(t, hc, endpointParameters{
		clusterName:  clusterName,
		endpointName: ok2EndpointName,
		privateURL:   ok2Neutron.URL,
		publicURL:    ok2Neutron.URL,
	})
	defer cleanupE2()

	// act/assert
	server.ForceProxyUpdate()
	verifyMultipleNeutronReadRequests(t, hc, clusterName)

	ok1Neutron.Close()
	verifyMultipleNeutronReadRequests(t, hc, clusterName)

	ok2Neutron.Close()
	verifyNeutronReadRequestsFail(t, hc, clusterName)
}

func TestDynamicProxyServiceWithUnavailableTargetServers(t *testing.T) {
	// arrange
	clusterName := contrailClusterName(t, "")
	const okEndpointName, badGatewayEndpointName = "neutron-ok", "neutron-bad-gateway"
	const unavailableEndpointName = "neutron-unavailable"
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	cleanupCC := createContrailCluster(t, hc, clusterName)
	defer cleanupCC()

	okNeutron := newNeutronServerStub(http.StatusOK)
	defer okNeutron.Close()
	cleanupE := createNeutronEndpoint(t, hc, endpointParameters{
		clusterName:  clusterName,
		endpointName: okEndpointName,
		privateURL:   okNeutron.URL,
		publicURL:    okNeutron.URL,
	})
	defer cleanupE()

	// act/assert
	server.ForceProxyUpdate()
	verifyMultipleNeutronReadRequests(t, hc, clusterName)

	// TODO(Daniel): proxy to other targets when 502/503 received from target
	badGatewayNeutron := newNeutronServerStub(http.StatusBadGateway)
	defer badGatewayNeutron.Close()
	cleanupE = createNeutronEndpoint(t, hc, endpointParameters{
		clusterName:  clusterName,
		endpointName: badGatewayEndpointName,
		privateURL:   badGatewayNeutron.URL,
		publicURL:    badGatewayNeutron.URL,
	})
	defer cleanupE()

	server.ForceProxyUpdate()
	verifyMultipleNeutronReadRequestsStatus(t, hc, clusterName, []int{http.StatusOK, http.StatusBadGateway})

	unavailableNeutron := newNeutronServerStub(http.StatusServiceUnavailable)
	defer unavailableNeutron.Close()
	cleanupE = createNeutronEndpoint(t, hc, endpointParameters{
		clusterName:  clusterName,
		endpointName: unavailableEndpointName,
		privateURL:   unavailableNeutron.URL,
		publicURL:    unavailableNeutron.URL,
	})
	defer cleanupE()

	server.ForceProxyUpdate()
	verifyMultipleNeutronReadRequestsStatus(
		t,
		hc,
		clusterName,
		[]int{http.StatusOK, http.StatusBadGateway, http.StatusServiceUnavailable},
	)
}

func contrailClusterName(t *testing.T, suffix string) string {
	return strings.ReplaceAll(t.Name(), "/", "_") + "-cluster" + suffix
}

func createContrailCluster(t *testing.T, hc *integration.HTTPAPIClient, clusterName string) (cleanup func()) {
	integration.CreateContrailCluster(t, hc, &models.ContrailCluster{
		UUID:       contrailClusterUUID(clusterName),
		FQName:     []string{basemodels.DefaultNameForKind(models.KindGlobalSystemConfig), clusterName},
		ParentType: models.KindGlobalSystemConfig,
	})
	return func() {
		hc.EnsureContrailClusterDeleted(t, contrailClusterUUID(clusterName))
	}
}
func createNeutronEndpoint(t *testing.T, hc *integration.HTTPAPIClient, ep endpointParameters) (cleanup func()) {
	integration.CreateEndpoint(t, hc, &models.Endpoint{
		UUID:       endpointUUID(ep.clusterName, ep.endpointName),
		Name:       fullEndpointName(ep.clusterName, ep.endpointName),
		ParentType: models.KindContrailCluster,
		ParentUUID: contrailClusterUUID(ep.clusterName),
		Prefix:     neutronEndpointPrefix,
		PrivateURL: ep.privateURL,
		PublicURL:  ep.publicURL,
	})
	return func() {
		hc.EnsureEndpointDeleted(t, endpointUUID(ep.clusterName, ep.endpointName))
	}
}

type endpointParameters struct {
	clusterName  string
	endpointName string
	privateURL   string
	publicURL    string
}

func setIncorrectEndpointURLs(t *testing.T, hc *integration.HTTPAPIClient, clusterName, endpointName string) {
	integration.UpdateEndpoint(t, hc, &models.Endpoint{
		UUID:       endpointUUID(clusterName, endpointName),
		PrivateURL: "http://127.0.0.1",
		PublicURL:  "http://127.0.0.1",
	})
}

func createNeutronServers(clusterName string) (publicS *httptest.Server, privateS *httptest.Server, close func()) {
	privateS = newNeutronPrivateServerStub(clusterName)
	publicS = newNeutronPublicServerStub(clusterName)

	return publicS, privateS, func() {
		privateS.Close()
		publicS.Close()
	}
}

func newNeutronPrivateServerStub(clusterName string) *httptest.Server {
	return newTestHTTPServer(routes{
		portsPath: func(c echo.Context) error {
			return c.JSON(http.StatusOK, &portsResponse{Foo: fooValueOnPrivateURL(clusterName)})
		},
	})
}

func fooValueOnPrivateURL(clusterName string) string {
	return clusterName + "-private-url"
}

func newNeutronPublicServerStub(clusterName string) *httptest.Server {
	return newTestHTTPServer(routes{
		portsPath: func(c echo.Context) error {
			clusterID := c.Request().Header.Get(apisrv.XClusterIDKey)
			if clusterID != contrailClusterUUID(clusterName) {
				return c.JSON(http.StatusBadRequest, "cluster ID not found in header")
			}
			return c.JSON(http.StatusOK, &portsResponse{Foo: fooValueOnPublicURL(clusterName)})
		},
	})
}

func fooValueOnPublicURL(clusterName string) string {
	return clusterName + "-public-url"
}

func newNeutronServerStub(statusToReturn int) *httptest.Server {
	return newTestHTTPServer(routes{
		portsPath: func(c echo.Context) error {
			return c.JSON(statusToReturn, &portsResponse{Foo: fooValueWithStatus(statusToReturn)})
		},
	})
}

func fooValueWithStatus(status int) string {
	return strconv.Itoa(status)
}

type portsResponse struct {
	Foo string `json:"foo"`
}

func newTestHTTPServer(r routes) *httptest.Server {
	e := echo.New()
	for route, handler := range r {
		e.GET(route, handler)
	}
	return httptest.NewServer(e)
}

type routes map[string]echo.HandlerFunc

func verifyNeutronReadRequests(t *testing.T, c *integration.HTTPAPIClient, clusterName string) {
	verifyNeutronReadRequest(t, c, neutronPortsPrivatePath(clusterName), fooValueOnPrivateURL(clusterName))
	verifyNeutronReadRequest(t, c, neutronPortsPublicPath(clusterName), fooValueOnPublicURL(clusterName))
}

func verifyMultipleNeutronReadRequests(t *testing.T, c *integration.HTTPAPIClient, clusterName string) {
	for i := 0; i < 3; i++ {
		verifyNeutronReadRequest(t, c, neutronPortsPrivatePath(clusterName), fooValueWithStatus(http.StatusOK))
		verifyNeutronReadRequest(t, c, neutronPortsPublicPath(clusterName), fooValueWithStatus(http.StatusOK))
	}
}

func verifyNeutronReadRequest(t *testing.T, c *integration.HTTPAPIClient, path, expectedValue string) {
	var response portsResponse
	_, err := c.Read(context.Background(), path, &response)

	assert.NoError(t, err, fmt.Sprintf("path: %v, response: %+v", path, response))
	assert.Equal(t, portsResponse{Foo: expectedValue}, response)
}

func verifyMultipleNeutronReadRequestsStatus(
	t *testing.T, c *integration.HTTPAPIClient, clusterName string, expectedStatuses []int,
) {
	for i := 0; i < 3; i++ {
		verifyNeutronReadRequestWithStatus(t, c, neutronPortsPrivatePath(clusterName), expectedStatuses)
		verifyNeutronReadRequestWithStatus(t, c, neutronPortsPublicPath(clusterName), expectedStatuses)
	}
}

func verifyNeutronReadRequestWithStatus(
	t *testing.T, c *integration.HTTPAPIClient, path string, expectedStatuses []int,
) {
	var response portsResponse
	r, err := c.Do(context.Background(), echo.GET, path, nil, nil, &response, expectedStatuses)

	assert.NoError(t, err, fmt.Sprintf("path: %v, response: %+v", path, response))
	assert.Equal(t, portsResponse{Foo: fooValueWithStatus(r.StatusCode)}, response)
}

func verifyNeutronReadRequestsFail(t *testing.T, c *integration.HTTPAPIClient, clusterName string) {
	verifyNeutronReadRequestFail(t, c, neutronPortsPrivatePath(clusterName))
	verifyNeutronReadRequestFail(t, c, neutronPortsPublicPath(clusterName))
}

func verifyNeutronReadRequestFail(t *testing.T, c *integration.HTTPAPIClient, path string) {
	var response map[string]interface{}
	r, err := c.Read(context.Background(), path, &response)

	assert.Error(t, err, fmt.Sprintf("path: %v, response: %+v", path, response))
	assert.Equal(t, http.StatusBadGateway, r.StatusCode)
}

func neutronPortsPrivatePath(clusterName string) string {
	return path.Join(
		"/",
		apisrv.DefaultDynamicProxyPath,
		contrailClusterUUID(clusterName),
		neutronEndpointPrefix,
		endpoint.PrivateURLScope,
		portsPath,
	)
}

func neutronPortsPublicPath(clusterName string) string {
	return path.Join(
		"/",
		apisrv.DefaultDynamicProxyPath,
		contrailClusterUUID(clusterName),
		neutronEndpointPrefix,
		portsPath,
	)
}

func contrailClusterUUID(clusterName string) string {
	return withUUIDSuffix(clusterName)
}

func endpointUUID(clusterName, endpointName string) string {
	return withUUIDSuffix(fullEndpointName(clusterName, endpointName))
}

func fullEndpointName(clusterName, endpointName string) string {
	return fmt.Sprintf("%s_%s", clusterName, endpointName)
}

func withUUIDSuffix(s string) string {
	return fmt.Sprintf("%s_uuid", s)
}

////////////////////
// Keystone tests //
////////////////////

func TestKeystoneEndpoint(t *testing.T) {
	keystoneAuthURL := viper.GetString("keystone.authurl")

	clusterCName := t.Name() + "_clusterC"
	clusterCUser := clusterCName + "_admin"
	ksPrivate := integration.MockServerWithKeystoneTestUser("", keystoneAuthURL, clusterCUser, clusterCUser)
	defer ksPrivate.Close()

	ksPublic := integration.MockServerWithKeystoneTestUser("", keystoneAuthURL, clusterCUser, clusterCUser)
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
	ctx := context.Background()
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

////////////////////////////////////
// Dynamic Proxy WebSockets tests //
////////////////////////////////////

func TestDynamicProxyServiceWebSocketsSupport(t *testing.T) {
	clusterName := t.Name() + "_cluster"
	endpointPrefix := "websocket-server"

	target := echoWebsocketServer(t)
	target.Start()
	defer target.Close()
	cleanup := createClusterAndEndpoint(t, target, clusterName, endpointPrefix)
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

func createClusterAndEndpoint(
	t *testing.T, target *httptest.Server, clusterName, endpointPrefix string,
) (cleanup func()) {
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
