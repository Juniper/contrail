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

type mockPortsRespose struct {
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

func TestProxyEndpoint(t *testing.T) {
	routes := map[string]interface{}{
		"/ports": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK,
				&mockPortsRespose{Name: "custer_a_private_port_list"})
		}),
	}
	neutronPrivate := mockServer(routes)
	defer neutronPrivate.Close()
	routes = map[string]interface{}{
		"/ports": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK,
				&mockPortsRespose{Name: "custer_a_public_port_list"})
		}),
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

	// Wait a sec for the dynamic proxy to be created/updated
	time.Sleep(time.Second)
	for _, client := range testScenario.Clients {
		var publicRes map[string]interface{}
		_, err := client.Read("/neutron/ports", &publicRes)
		assert.NoError(t, err, "failed to proxy neutron public url")
		ok := common.AssertEqual(t,
			map[string]interface{}{"name": "custer_a_public_port_list"},
			publicRes,
			fmt.Sprintf("Unexpected Response: %s", publicRes))
		if !ok {
			return
		}
		var privateRes map[string]interface{}
		_, err = client.Read("/neutron/private/ports", &privateRes)
		assert.NoError(t, err, "failed to proxy neutron private url")
		ok = common.AssertEqual(t,
			map[string]interface{}{"name": "custer_a_private_port_list"},
			privateRes,
			fmt.Sprintf("Unexpected Response: %s", privateRes))
		if !ok {
			return
		}

	}
}
