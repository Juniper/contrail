package proxy_test

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

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/proxy"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/websocket"

	asfclient "github.com/Juniper/asf/pkg/client"
	asfkeystone "github.com/Juniper/asf/pkg/keystone"
)

const (
	keystoneEndpointPrefix = "keystone"
	neutronEndpointPrefix  = "neutron"
	swiftEndpointPrefix    = "swift"
	portsPath              = "/ports"
)

//////////////////////////////
// Dynamic proxy HTTP tests //
//////////////////////////////

func TestDynamicProxyHTTPSupport(t *testing.T) {
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
				time.Sleep(proxy.SyncInterval)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			testDynamicProxyHTTPSupport(t, tt.synchronizeProxyEndpoints)
		})
	}
}

func testDynamicProxyHTTPSupport(t *testing.T, synchronizeProxyEndpoints func(s *integration.APIServer)) {
	// arrange
	const neutron1EndpointName, neutron2EndpointName = "neutron1", "neutron2"
	clusterAName, clusterBName := contrailClusterName(t, "A"), contrailClusterName(t, "B")
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	cleanupCCA := createContrailCluster(t, hc, clusterAName)
	defer cleanupCCA()

	neutronA1Public, neutronA1Private, closeNeutronA1 := createNeutronServers(clusterAName)
	defer closeNeutronA1()

	cleanupEA1 := createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterAName,
		endpointName:   neutron1EndpointName,
		endpointPrefix: neutronEndpointPrefix,
		privateURL:     neutronA1Private.URL,
		publicURL:      neutronA1Public.URL,
	})
	defer cleanupEA1()

	neutronA2Public, neutronA2Private, closeNeutronA2 := createNeutronServers(clusterAName)
	defer closeNeutronA2()

	cleanupEA2 := createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterAName,
		endpointName:   neutron2EndpointName,
		endpointPrefix: neutronEndpointPrefix,
		privateURL:     neutronA2Private.URL,
		publicURL:      neutronA2Public.URL,
	})
	defer cleanupEA2()

	synchronizeProxyEndpoints(server)

	// act/assert
	verifyNeutronReadRequests(t, hc, clusterAName)

	cleanupCCB := createContrailCluster(t, hc, clusterBName)
	defer cleanupCCB()

	neutronB1Public, neutronB1Private, closeNeutronB1 := createNeutronServers(clusterBName)
	defer closeNeutronB1()

	cleanupEB1 := createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterBName,
		endpointName:   neutron1EndpointName,
		endpointPrefix: neutronEndpointPrefix,
		privateURL:     neutronB1Private.URL,
		publicURL:      neutronB1Public.URL,
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
	cleanupEA1 = createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterAName,
		endpointName:   neutron1EndpointName,
		endpointPrefix: neutronEndpointPrefix,
		privateURL:     neutronA1Private.URL,
		publicURL:      neutronA1Public.URL,
	})
	defer cleanupEA1()
	synchronizeProxyEndpoints(server)

	verifyNeutronReadRequests(t, hc, clusterAName)
	verifyNeutronReadRequests(t, hc, clusterBName)
}

func TestDynamicProxyRetriesRequestWhenTargetServerIsClosed(t *testing.T) {
	// arrange
	const ok1EndpointName, ok2EndpointName = "neutron-ok1", "neutron-ok2"
	clusterName := contrailClusterName(t, "")
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	cleanupCC := createContrailCluster(t, hc, clusterName)
	defer cleanupCC()

	ok1Neutron := newNeutronServerStub(http.StatusOK)
	defer ok1Neutron.Close()

	cleanupE1 := createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterName,
		endpointName:   ok1EndpointName,
		endpointPrefix: neutronEndpointPrefix,
		privateURL:     ok1Neutron.URL,
		publicURL:      ok1Neutron.URL,
	})
	defer cleanupE1()

	ok2Neutron := newNeutronServerStub(http.StatusOK)
	defer ok2Neutron.Close()

	cleanupE2 := createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterName,
		endpointName:   ok2EndpointName,
		endpointPrefix: neutronEndpointPrefix,
		privateURL:     ok2Neutron.URL,
		publicURL:      ok2Neutron.URL,
	})
	defer cleanupE2()
	server.ForceProxyUpdate()

	// act/assert
	verifyFiveNeutronReadRequests(t, hc, clusterName)

	ok1Neutron.Close()
	verifyFiveNeutronReadRequests(t, hc, clusterName)

	ok2Neutron.Close()
	verifyNeutronReadRequestsFail(t, hc, clusterName)
}

func TestDynamicProxyRetriesRequestWhenTargetServerReturnsServerError(t *testing.T) {
	// arrange
	const (
		healthyEndpointName, unhealthyEndpointName = "healthy-neutron", "unhealthy-neutron"
	)
	clusterName := contrailClusterName(t, "")
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	cleanupCC := createContrailCluster(t, hc, clusterName)
	defer cleanupCC()

	healthyNeutron := newNeutronServerStub(http.StatusOK)
	defer healthyNeutron.Close()

	cleanupE := createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterName,
		endpointName:   healthyEndpointName,
		endpointPrefix: neutronEndpointPrefix,
		privateURL:     healthyNeutron.URL,
		publicURL:      healthyNeutron.URL,
	})
	defer cleanupE()
	server.ForceProxyUpdate()

	// act/assert
	verifyFiveNeutronReadRequests(t, hc, clusterName)

	for status := http.StatusInternalServerError; status <= http.StatusNetworkAuthenticationRequired; status++ {
		unhealthyNeutron := newNeutronServerStub(status)
		defer unhealthyNeutron.Close()

		cleanupE = createEndpoint(t, hc, endpointParameters{
			clusterName:    clusterName,
			endpointName:   unhealthyEndpointName,
			endpointPrefix: neutronEndpointPrefix,
			privateURL:     unhealthyNeutron.URL,
			publicURL:      unhealthyNeutron.URL,
		})
		defer cleanupE() // ensure that test does not leave dirty DB
		server.ForceProxyUpdate()

		verifyFiveNeutronReadRequests(t, hc, clusterName)

		cleanupE() // next iteration with clean endpoints
	}
}

func TestDynamicProxyInjectsServiceToken(t *testing.T) {
	// arrange
	const (
		keystoneEndpointName  = "keystone"
		swiftEndpointName     = "some-swift"
		unrelatedEndpointName = "neutron"

		// Defined in ../../sample/test_config.yml
		serviceUserUsername = "goapi"
		serviceUserPassword = "goapi"
	)

	clusterName := contrailClusterName(t, "")
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	cleanupCC := createContrailCluster(t, hc, clusterName)
	defer cleanupCC()

	authURL := server.URL() + keystone.LocalAuthPath
	clusterKeystone := integration.NewKeystoneServerFake(t, authURL, serviceUserUsername, serviceUserPassword)
	defer clusterKeystone.Close()

	swift := newCollectingServerStub()
	defer swift.Close()

	unrelated := newCollectingServerStub()
	defer unrelated.Close()

	cleanupKeystoneEndpoint := createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterName,
		endpointName:   keystoneEndpointName,
		endpointPrefix: keystoneEndpointPrefix,
		privateURL:     clusterKeystone.URL,
		publicURL:      clusterKeystone.URL,
		username:       serviceUserUsername,
		password:       serviceUserPassword,
	})
	defer cleanupKeystoneEndpoint()

	cleanupSwiftEndpoint := createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterName,
		endpointName:   swiftEndpointName,
		endpointPrefix: swiftEndpointPrefix,
		privateURL:     swift.URL,
		publicURL:      swift.URL,
	})
	defer cleanupSwiftEndpoint()

	cleanupUnrelatedEndpoint := createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterName,
		endpointName:   unrelatedEndpointName,
		endpointPrefix: neutronEndpointPrefix,
		privateURL:     unrelated.URL,
		publicURL:      unrelated.URL,
	})
	defer cleanupUnrelatedEndpoint()

	server.ForceProxyUpdate()

	// act/assert
	verifyNeutronReadRequest(t, hc, swiftPrivatePath(clusterName), fooValueWithStatus(http.StatusOK))
	if assert.Len(t, swift.Requests, 1, "1 request should be sent to swift") {
		_, ok := swift.Requests[0].Header["X-Service-Token"]
		assert.True(t, ok, "an X-Service-Token header should be added to the request to swift")
	}

	verifyNeutronReadRequest(t, hc, neutronPortsPrivatePath(clusterName), fooValueWithStatus(http.StatusOK))
	if assert.Len(t, unrelated.Requests, 1, "1 request should be sent to the unrelated endpoint") {
		_, ok := unrelated.Requests[0].Header["X-Service-Token"]
		assert.False(t, ok, "no X-Service-Token header should be added to the request to the unrelated endpoint")
	}
}

type collectingHTTPServer struct {
	*httptest.Server
	Requests []*http.Request
}

func newCollectingServerStub() *collectingHTTPServer {
	cs := &collectingHTTPServer{}
	cs.Server = newTestHTTPServer(routes{
		portsPath: func(c echo.Context) error {
			cs.Requests = append(cs.Requests, c.Request())
			return c.JSON(http.StatusOK, &portsResponse{Foo: fooValueWithStatus(http.StatusOK)})
		},
	})
	return cs
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
		portsPath: func(ctx echo.Context) error {
			return ctx.JSON(http.StatusOK, &portsResponse{Foo: fooValueOnPrivateURL(clusterName)})
		},
	})
}

func fooValueOnPrivateURL(clusterName string) string {
	return clusterName + "-private-url"
}

func newNeutronPublicServerStub(clusterName string) *httptest.Server {
	return newTestHTTPServer(routes{
		portsPath: func(ctx echo.Context) error {
			clusterID := ctx.Request().Header.Get(proxy.XClusterIDKey)
			if clusterID != contrailClusterUUID(clusterName) {
				return ctx.JSON(http.StatusBadRequest, "cluster ID not found in header")
			}
			return ctx.JSON(http.StatusOK, &portsResponse{Foo: fooValueOnPublicURL(clusterName)})
		},
	})
}

func fooValueOnPublicURL(clusterName string) string {
	return clusterName + "-public-url"
}

func newNeutronServerStub(statusToReturn int) *httptest.Server {
	return newTestHTTPServer(routes{
		portsPath: func(ctx echo.Context) error {
			return ctx.JSON(statusToReturn, &portsResponse{Foo: fooValueWithStatus(statusToReturn)})
		},
	})
}

func newTestHTTPServer(r routes) *httptest.Server {
	e := echo.New()
	for route, handler := range r {
		e.GET(route, handler)
	}
	return httptest.NewServer(e)
}

type routes map[string]echo.HandlerFunc

type portsResponse struct {
	Foo string `json:"foo"`
}

func verifyNeutronReadRequests(t *testing.T, c *integration.HTTPAPIClient, clusterName string) {
	verifyNeutronReadRequest(t, c, neutronPortsPrivatePath(clusterName), fooValueOnPrivateURL(clusterName))
	verifyNeutronReadRequest(t, c, neutronPortsPublicPath(clusterName), fooValueOnPublicURL(clusterName))
}

func verifyFiveNeutronReadRequests(t *testing.T, c *integration.HTTPAPIClient, clusterName string) {
	for i := 0; i < 5; i++ {
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
		proxy.DefaultPath,
		contrailClusterUUID(clusterName),
		neutronEndpointPrefix,
		proxy.PrivateURLScope,
		portsPath,
	)
}

func neutronPortsPublicPath(clusterName string) string {
	return path.Join(
		"/",
		proxy.DefaultPath,
		contrailClusterUUID(clusterName),
		neutronEndpointPrefix,
		portsPath,
	)
}

func swiftPrivatePath(clusterName string) string {
	return path.Join(
		"/",
		proxy.DefaultPath,
		contrailClusterUUID(clusterName),
		swiftEndpointPrefix,
		proxy.PrivateURLScope,
		portsPath,
	)
}

////////////////////
// Keystone tests //
////////////////////

func TestKeystoneProxy(t *testing.T) {
	// arrange
	const (
		endpointName         = "keystone-endpoint"
		usernameC, usernameD = "test-user-C", "test-user-D"
		passwordC, passwordD = "test-password-C", "test-password-D"
	)
	clusterCName, clusterDName := contrailClusterName(t, "C"), contrailClusterName(t, "D")
	authURL := server.URL() + keystone.LocalAuthPath
	hcBob := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)
	hcUserC := client.NewHTTP(&asfclient.HTTPConfig{
		ID:       usernameC,
		Password: passwordC,
		Endpoint: server.URL(),
		AuthURL:  authURL,
		Insecure: true,
	})
	hcUserD := client.NewHTTP(&asfclient.HTTPConfig{
		ID:       usernameD,
		Password: passwordD,
		Endpoint: server.URL(),
		AuthURL:  authURL,
		Insecure: true,
	})

	cleanupCCC := createContrailCluster(t, hcBob, clusterCName)
	defer cleanupCCC()
	cleanupCCD := createContrailCluster(t, hcBob, clusterDName)
	defer cleanupCCD()

	ksPrivateC := integration.NewKeystoneServerFake(t, authURL, usernameC, passwordC)
	defer ksPrivateC.Close()
	ksPublicC := integration.NewKeystoneServerFake(t, authURL, usernameC, passwordC)
	defer ksPublicC.Close()
	ksPrivateD := integration.NewKeystoneServerFake(t, authURL, usernameD, passwordD)
	defer ksPrivateD.Close()
	ksPublicD := integration.NewKeystoneServerFake(t, authURL, usernameD, passwordD)
	defer ksPublicD.Close()

	cleanupEC := createEndpoint(t, hcBob, endpointParameters{
		clusterName:    clusterCName,
		endpointName:   endpointName,
		endpointPrefix: keystoneEndpointPrefix,
		privateURL:     ksPrivateC.URL,
		publicURL:      ksPublicC.URL,
		username:       usernameC,
		password:       passwordC,
	})
	defer cleanupEC()

	cleanupED := createEndpoint(t, hcBob, endpointParameters{
		clusterName:    clusterDName,
		endpointName:   endpointName,
		endpointPrefix: keystoneEndpointPrefix,
		privateURL:     ksPrivateD.URL,
		publicURL:      ksPublicD.URL,
		username:       usernameD,
		password:       passwordD,
	})
	defer cleanupED()
	server.ForceProxyUpdate()

	// act/assert
	// When multiple cluster endpoints are present auth middleware should find the keystone endpoint
	// with X-Cluster-ID in the header.
	verifyCreateTokenRequest(ctxWithXClusterID(clusterCName), t, hcUserC, "with X-Cluster-ID")
	verifyCreateTokenRequest(ctxWithXClusterID(clusterDName), t, hcUserD, "with X-Cluster-ID")
	verifyReadTokenRequest(ctxWithXClusterID(clusterCName), t, hcUserC, "with X-Cluster-ID")
	verifyReadTokenRequest(ctxWithXClusterID(clusterDName), t, hcUserD, "with X-Cluster-ID")

	// When multiple cluster endpoints are present auth middleware cannot not find keystone endpoint
	// without X-Cluster-ID in the header.
	verifyCreateTokenRequestFails(context.Background(), t, hcUserC, "without X-Cluster-ID")
	verifyCreateTokenRequestFails(context.Background(), t, hcUserD, "without X-Cluster-ID")
	verifyReadTokenRequestFails(context.Background(), t, hcUserC, "without X-Cluster-ID")
	verifyReadTokenRequestFails(context.Background(), t, hcUserD, "without X-Cluster-ID")
}

func TestKeystoneProxyRetriesRequestWhenTargetKeystoneServerIsClosed(t *testing.T) {
	// arrange
	const (
		endpointNameOne, endpointNameTwo = "keystone-endpoint-one", "keystone-endpoint-two"
		username                         = "test-user"
		password                         = "test-password"
	)
	clusterName := contrailClusterName(t, "")
	authURL := server.URL() + keystone.LocalAuthPath
	hcBob := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)
	hcTest := client.NewHTTP(&asfclient.HTTPConfig{
		ID:       username,
		Password: password,
		Endpoint: server.URL(),
		AuthURL:  authURL,
		Insecure: true,
	})

	cleanupCC := createContrailCluster(t, hcBob, clusterName)
	defer cleanupCC()

	ksPrivateOne := integration.NewKeystoneServerFake(t, authURL, username, password)
	defer ksPrivateOne.Close()
	ksPublicOne := integration.NewKeystoneServerFake(t, authURL, username, password)
	defer ksPublicOne.Close()
	ksPrivateTwo := integration.NewKeystoneServerFake(t, authURL, username, password)
	defer ksPrivateTwo.Close()
	ksPublicTwo := integration.NewKeystoneServerFake(t, authURL, username, password)
	defer ksPublicTwo.Close()

	cleanupEOne := createEndpoint(t, hcBob, endpointParameters{
		clusterName:    clusterName,
		endpointName:   endpointNameOne,
		endpointPrefix: keystoneEndpointPrefix,
		privateURL:     ksPrivateOne.URL,
		publicURL:      ksPublicOne.URL,
		username:       username,
		password:       password,
	})
	defer cleanupEOne()

	cleanupETwo := createEndpoint(t, hcBob, endpointParameters{
		clusterName:    clusterName,
		endpointName:   endpointNameTwo,
		endpointPrefix: keystoneEndpointPrefix,
		privateURL:     ksPrivateTwo.URL,
		publicURL:      ksPublicTwo.URL,
		username:       username,
		password:       password,
	})
	defer cleanupETwo()
	server.ForceProxyUpdate()

	// act/assert
	verifyFiveCreateTokenRequests(ctxWithXClusterID(clusterName), t, hcTest, "0/2 Keystone servers closed")
	// TODO(dfurman): find and fix race condition
	//verifyFiveReadTokenRequests(ctxWithXClusterID(clusterName), t, hcTest, "0/2 Keystone servers closed")

	ksPrivateOne.Close()
	ksPublicOne.Close()

	verifyFiveCreateTokenRequests(ctxWithXClusterID(clusterName), t, hcTest, "1/2 Keystone servers closed")
	// TODO(dfurman): find and fix race condition
	//verifyFiveReadTokenRequests(ctxWithXClusterID(clusterName), t, hcTest, "1/2 Keystone servers closed")

	ksPrivateTwo.Close()
	ksPublicTwo.Close()

	verifyFiveCreateTokenRequestsFail(ctxWithXClusterID(clusterName), t, hcTest, "all Keystone servers closed")
	verifyFiveReadTokenRequestsFail(ctxWithXClusterID(clusterName), t, hcTest, "all Keystone servers closed")
}

func TestKeystoneProxyRetriesRequestWhenTargetKeystoneServerReturnsServerError(t *testing.T) {
	// arrange
	const (
		healthyEndpointName   = "healthy-keystone-endpoint"
		unhealthyEndpointName = "unhealthy-keystone-endpoint"
		username              = "test-user"
		password              = "test-password"
	)
	clusterName := contrailClusterName(t, "")
	authURL := server.URL() + keystone.LocalAuthPath
	hcBob := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)
	hcTest := client.NewHTTP(&asfclient.HTTPConfig{
		ID:       username,
		Password: password,
		Endpoint: server.URL(),
		AuthURL:  authURL,
		Insecure: true,
	})

	cleanupCC := createContrailCluster(t, hcBob, clusterName)
	defer cleanupCC()

	ksPrivateOne := integration.NewKeystoneServerFake(t, authURL, username, password)
	defer ksPrivateOne.Close()
	ksPublicOne := integration.NewKeystoneServerFake(t, authURL, username, password)
	defer ksPublicOne.Close()

	cleanupEOne := createEndpoint(t, hcBob, endpointParameters{
		clusterName:    clusterName,
		endpointName:   healthyEndpointName,
		endpointPrefix: keystoneEndpointPrefix,
		privateURL:     ksPrivateOne.URL,
		publicURL:      ksPublicOne.URL,
		username:       username,
		password:       password,
	})
	defer cleanupEOne()
	server.ForceProxyUpdate()

	// act/assert
	verifyFiveCreateTokenRequests(ctxWithXClusterID(clusterName), t, hcTest, "only healthy Keystone server up")
	verifyFiveReadTokenRequests(ctxWithXClusterID(clusterName), t, hcTest, "only healthy Keystone servers up")

	for status := http.StatusInternalServerError; status <= http.StatusNetworkAuthenticationRequired; status++ {
		ksBadGateway := newKeystoneServerStub(status)
		defer ksBadGateway.Close()

		cleanupEBadGateway := createEndpoint(t, hcBob, endpointParameters{
			clusterName:    clusterName,
			endpointName:   unhealthyEndpointName,
			endpointPrefix: keystoneEndpointPrefix,
			privateURL:     ksBadGateway.URL,
			publicURL:      ksBadGateway.URL,
			username:       username,
			password:       password,
		})
		defer cleanupEBadGateway() // ensure that test does not leave dirty DB
		server.ForceProxyUpdate()

		verifyFiveCreateTokenRequests(ctxWithXClusterID(clusterName), t, hcTest,
			"both healthy and unhealthy Keystone servers up",
		)
		verifyFiveReadTokenRequests(ctxWithXClusterID(clusterName), t, hcTest,
			"both healthy and unhealthy Keystone servers up",
		)

		cleanupEBadGateway() // next iteration with clean endpoints
	}
}

func newKeystoneServerStub(statusToReturn int) *httptest.Server {
	e := echo.New()
	e.Any("/v3/auth/tokens", func(ctx echo.Context) error {
		return ctx.JSON(statusToReturn, &portsResponse{Foo: fooValueWithStatus(statusToReturn)})
	})
	return httptest.NewServer(e)
}
func ctxWithXClusterID(clusterName string) context.Context {
	return keystone.WithXClusterID(context.Background(), contrailClusterUUID(clusterName))
}

func verifyFiveCreateTokenRequests(ctx context.Context, t *testing.T, hc *client.HTTP, msg string) {
	for i := 0; i < 5; i++ {
		verifyCreateTokenRequest(ctx, t, hc, msg)
	}
}

func verifyCreateTokenRequest(ctx context.Context, t *testing.T, hc *client.HTTP, msg string) {
	err := hc.Login(ctx)
	assert.NoError(t, err, "%s, HTTP client ID: %s", msg, hc.ID)
}

func verifyFiveCreateTokenRequestsFail(ctx context.Context, t *testing.T, hc *client.HTTP, msg string) {
	for i := 0; i < 5; i++ {
		verifyCreateTokenRequestFails(ctx, t, hc, msg)
	}
}

func verifyCreateTokenRequestFails(ctx context.Context, t *testing.T, hc *client.HTTP, msg string) {
	err := hc.Login(ctx)
	assert.Error(t, err, "%s, HTTP client ID: %s", msg, hc.ID)
}

func verifyFiveReadTokenRequests(ctx context.Context, t *testing.T, hc *client.HTTP, msg string) {
	for i := 0; i < 5; i++ {
		verifyReadTokenRequest(ctx, t, hc, msg)
	}
}

func verifyReadTokenRequest(ctx context.Context, t *testing.T, hc *client.HTTP, msg string) {
	var response asfkeystone.ValidateTokenResponse
	_, err := hc.Read(ctx, path.Join(keystone.LocalAuthPath, "auth/tokens"), &response)

	msg = fmt.Sprintf("%s, HTTP client ID: %s", msg, hc.ID)
	assert.NoError(t, err, msg)
	if assert.NotNil(t, response.Token, msg) {
		assert.NotNil(t, response.Token.Domain, msg)
		assert.NotNil(t, response.Token.User, msg)
		assert.NotNil(t, response.Token.Roles, msg)
	}
}

func verifyFiveReadTokenRequestsFail(ctx context.Context, t *testing.T, hc *client.HTTP, msg string) {
	for i := 0; i < 5; i++ {
		verifyReadTokenRequestFails(ctx, t, hc, msg)
	}
}

func verifyReadTokenRequestFails(ctx context.Context, t *testing.T, hc *client.HTTP, msg string) {
	var response interface{}
	_, err := hc.Read(ctx, path.Join(keystone.LocalAuthPath, "auth/tokens"), &response)

	assert.Error(t, err, "%s, HTTP client ID: %s", msg, hc.ID)
	assert.Nil(t, response, "%s, HTTP client ID: %s", msg, hc.ID)
}

////////////////////////////////////
// Dynamic Proxy WebSockets tests //
////////////////////////////////////

func TestDynamicProxyWebSocketsSupport(t *testing.T) {
	// arrange
	clusterName := contrailClusterName(t, "")
	const endpointPrefix, endpointName = "websocket-prefix", "websocket-endpoint"
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	cleanupCC := createContrailCluster(t, hc, clusterName)
	defer cleanupCC()

	target := echoWebsocketServer(t)
	defer target.Close()

	cleanupE := createEndpoint(t, hc, endpointParameters{
		clusterName:    clusterName,
		endpointName:   endpointName,
		endpointPrefix: endpointPrefix,
		privateURL:     target.URL,
		publicURL:      target.URL,
	})
	defer cleanupE()
	server.ForceProxyUpdate()

	config, err := websocket.NewConfig(requestURL(clusterName, endpointPrefix), "http://localhost/")
	assert.NoError(t, err, "failed to create websocket config from proxy URL")
	config.TlsConfig = &tls.Config{InsecureSkipVerify: true}

	// act/assert
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
	return httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) {
		if _, err := io.Copy(ws, ws); err != nil {
			t.Error("Failed to echo the message back to the client: ", err)
		}
	}))
}

func requestURL(clusterName, endpointPrefix string) string {
	wsURLBase := strings.ReplaceAll(server.URL(), "https://", "wss://")
	return fmt.Sprintf(
		"%s/%s/%s/%s",
		wsURLBase,
		proxy.DefaultPath,
		contrailClusterUUID(clusterName),
		endpointPrefix,
	)
}

///////////////////////////
// Common test utilities //
///////////////////////////

func fooValueWithStatus(status int) string {
	return strconv.Itoa(status)
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

func createEndpoint(t *testing.T, hc *integration.HTTPAPIClient, ep endpointParameters) (cleanup func()) {
	integration.CreateEndpoint(t, hc, &models.Endpoint{
		UUID:       endpointUUID(ep.clusterName, ep.endpointName),
		Name:       ep.endpointName,
		ParentType: models.KindContrailCluster,
		ParentUUID: contrailClusterUUID(ep.clusterName),
		Prefix:     ep.endpointPrefix,
		PrivateURL: ep.privateURL,
		PublicURL:  ep.publicURL,
		Username:   ep.username,
		Password:   ep.password,
	})
	return func() {
		hc.EnsureEndpointDeleted(t, endpointUUID(ep.clusterName, ep.endpointName))
		server.ForceProxyUpdate()
	}
}

type endpointParameters struct {
	clusterName    string
	endpointName   string
	endpointPrefix string
	privateURL     string
	publicURL      string
	username       string
	password       string
}

func contrailClusterUUID(clusterName string) string {
	return withUUIDSuffix(clusterName)
}

// endpointUUID returns UUID. It prevents endpoint UUID collisions between tests.
func endpointUUID(clusterName, endpointName string) string {
	return withUUIDSuffix(fmt.Sprintf("%s_%s", clusterName, endpointName))
}

func withUUIDSuffix(s string) string {
	return fmt.Sprintf("%s_uuid", s)
}
