package apisrv

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)

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

func TestProxyEndpoint(t *testing.T) {
	routes := map[string]interface{}{
		"/ports": echo.HandlerFunc(func(c echo.Context) error { return c.String(http.StatusOK, "custer_a_neutron_ports_via_private_url") }),
	}
	neutronPrivate := mockServer(routes)
	defer neutronPrivate.Close()
	routes = map[string]interface{}{
		"/ports": echo.HandlerFunc(func(c echo.Context) error { return c.String(http.StatusOK, "custer_a_neutron_ports_via_public_url") }),
	}
	neutronPublic := mockServer(routes)
	defer neutronPublic.Close()

	context := pongo2.Context{
		"cluster_a_neutron_private_url": neutronPrivate.URL,
		"cluster_a_neutron_public_url":  neutronPublic.URL,
	}

	testFile := GetTestFromTemplate(t, "./test_data/test_endpoint.tmpl", context)
	defer os.Remove(testFile) // remove tempfile after test

	var testScenario TestScenario
	LoadTestScenario(&testScenario, testFile)
	RunTestScenario(t, &testScenario)

	for _, client := range testScenario.Clients {
		var publicRes map[string]interface{}
		client.Read("/neutron/ports", &publicRes)
		var privateRes map[string]interface{}
		client.Read("/neutron/private/ports", &privateRes)

	}
}
