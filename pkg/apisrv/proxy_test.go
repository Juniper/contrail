package apisrv

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const (
	privatePortList  = "private_port_list"
	publicPortList   = "public_port_list"
	privateAuthToken = "private_auth_token"
	publicAuthToken  = "public_auth_token"
)

type mockPortsResponse struct {
	Name string `json:"name"`
}

type mockAuthTokenResponse struct {
	Token string `json:"token"`
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
	clusterName string, key string, expected string) bool {
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		_, err := client.Read(url, &response)
		assert.NoError(t, err, "failed to proxy %s", url)
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

func TestProxyEndpoint(t *testing.T) {
	// Create a cluster and its neutron endpoint
	clusterAName := "clusterA"
	testFile, testScenario, neutronPublic, neutronPrivate := runEndpointTest(
		t, clusterAName, true)
	defer os.Remove(testFile) // remove tempfile after test
	defer neutronPrivate.Close()
	defer neutronPublic.Close()

	// Wait a sec for the dynamic proxy to be created/updated
	time.Sleep(time.Second)

	// verify proxies
	ok := verifyProxy(t, testScenario,
		"/proxy/"+clusterAName+"_uuid/neutron/ports", clusterAName, "name", publicPortList)
	if !ok {
		return
	}
	ok = verifyProxy(t, testScenario,
		"/proxy/"+clusterAName+"_uuid/neutron/private/ports", clusterAName, "name", privatePortList)
	if !ok {
		return
	}

	// create one more cluster/neutron endpoint for new cluster
	clusterBName := "clusterB"
	testFile, testScenario, neutronPublic, neutronPrivate = runEndpointTest(
		t, clusterBName, false)
	defer os.Remove(testFile) // remove tempfile after test
	defer neutronPrivate.Close()
	defer neutronPublic.Close()
	// Wait a sec for the dynamic proxy to be created/updated
	time.Sleep(time.Second)

	// verify new proxies
	ok = verifyProxy(t, testScenario,
		"/proxy/"+clusterBName+"_uuid/neutron/ports",
		clusterBName, "name", publicPortList)
	if !ok {
		return
	}
	ok = verifyProxy(t, testScenario,
		"/proxy/"+clusterBName+"_uuid/neutron/private/ports",
		clusterBName, "name", privatePortList)
	if !ok {
		return
	}

	// verify existing proxies, make sure the proxy prefix is updated with cluster id
	ok = verifyProxy(t, testScenario,
		"/proxy/"+clusterAName+"_uuid/neutron/ports",
		clusterAName, "name", publicPortList)
	if !ok {
		return
	}
	ok = verifyProxy(t, testScenario,
		"/proxy/"+clusterAName+"_uuid/neutron/private/ports",
		clusterAName, "name", privatePortList)
	if !ok {
		return
	}
}

func TestKeystoneEndpoint(t *testing.T) {
	clusterName := "clusterA"
	routes := map[string]interface{}{
		"/v3/auth/tokens": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK,
				&mockAuthTokenResponse{Token: clusterName + privateAuthToken})
		}),
	}
	ksPrivate := mockServer(routes)
	defer ksPrivate.Close()

	routes = map[string]interface{}{
		"/v3/auth/tokens": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK,
				&mockAuthTokenResponse{Token: clusterName + publicAuthToken})
		}),
	}
	ksPublic := mockServer(routes)
	defer ksPublic.Close()

	context := pongo2.Context{
		"extra_tasks":   true,
		"cluster_name":  clusterName,
		"endpoint_name": "keystone",
		"private_url":   ksPrivate.URL,
		"public_url":    ksPublic.URL,
	}

	testFile := GetTestFromTemplate(t, "./test_data/test_endpoint.tmpl", context)
	defer os.Remove(testFile) // remove tempfile after test

	var testScenario TestScenario
	err := LoadTestScenario(&testScenario, testFile)
	assert.NoError(t, err, "failed to load test data")
	RunTestScenario(t, &testScenario)

	// Wait a sec for the dynamic proxy to be created/updated
	time.Sleep(time.Second)

	// verify proxies
	ok := verifyProxy(t, &testScenario,
		"/proxy/"+clusterName+"_uuid/keystone/v3/auth/tokens",
		clusterName, "token", publicAuthToken)
	if !ok {
		return
	}
	ok = verifyProxy(t, &testScenario,
		"/proxy/"+clusterName+"_uuid/keystone/private/v3/auth/tokens",
		clusterName, "token", privateAuthToken)
	if !ok {
		return
	}
}
