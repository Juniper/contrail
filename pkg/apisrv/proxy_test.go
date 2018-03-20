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
		"extra_tasks":         extraTasks,
		"cluster_name":        clusterName,
		"neutron_private_url": neutronPrivate.URL,
		"neutron_public_url":  neutronPublic.URL,
	}

	testFile := GetTestFromTemplate(t, "./test_data/test_endpoint.tmpl", context)

	var testScenario TestScenario
	LoadTestScenario(&testScenario, testFile)
	RunTestScenario(t, &testScenario)

	return testFile, &testScenario, neutronPublic, neutronPrivate
}

func verifyProxy(t *testing.T, testScenario *TestScenario, url string,
	clusterName string, expected string) bool {
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		_, err := client.Read(url, &response)
		assert.NoError(t, err, "failed to proxy %s", url)
		ok := common.AssertEqual(t,
			map[string]interface{}{"name": clusterName + expected},
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
		"/neutron/ports", clusterAName, publicPortList)
	if !ok {
		return
	}
	ok = verifyProxy(t, testScenario,
		"/neutron/private/ports", clusterAName, privatePortList)
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
		"/"+clusterBName+"_uuid/neutron/ports", clusterBName, publicPortList)
	if !ok {
		return
	}
	ok = verifyProxy(t, testScenario,
		"/"+clusterBName+"_uuid/neutron/private/ports", clusterBName, privatePortList)
	if !ok {
		return
	}

	// verify existing proxies, make sure the proxy prefix is updated with cluster id
	ok = verifyProxy(t, testScenario,
		"/"+clusterAName+"_uuid/neutron/ports", clusterAName, publicPortList)
	if !ok {
		return
	}
	ok = verifyProxy(t, testScenario,
		"/"+clusterAName+"_uuid/neutron/private/ports", clusterAName, privatePortList)
	if !ok {
		return
	}
}
