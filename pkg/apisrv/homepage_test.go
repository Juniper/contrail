package apisrv_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestHomepageResources(t *testing.T) {
	c := integration.NewHTTPAPIClient(t, apisrv.TestServer.URL)

	var response map[string]interface{}
	r, err := c.Read("/", &response)
	assert.NoError(t, err, fmt.Sprintf("GET failed\n response: %+v", r))

	addr := strings.TrimSuffix(viper.GetString("server.address"), "/")
	expected := map[string]interface{}{
		"href": addr,
		"links": []interface{}{
			map[string]interface{}{
				"href":   strings.Join([]string{addr, "virtual-network"}, "/"),
				"name":   "virtual-network",
				"rel":    "resource-base",
				"method": nil,
			},
			map[string]interface{}{
				"href":   strings.Join([]string{addr, "virtual-networks"}, "/"),
				"name":   "virtual-network",
				"rel":    "collection",
				"method": nil,
			},
		},
	}

	require.Contains(t, response, "href")
	assert.Equal(t, expected["href"], response["href"])

	require.Contains(t, response, "links")
	assert.Subset(t, response["links"], expected["links"])
}

type routeSet struct {
	set map[string]struct{}
}

func (r *routeSet) add(path string) {
	r.set[path] = struct{}{}
}

func (r *routeSet) contains(path string) bool {
	_, result := r.set[path]
	return result
}

func TestRoutesAreRegistered(t *testing.T) {
	c := integration.NewHTTPAPIClient(t, apisrv.TestServer.URL)

	var response map[string]interface{}
	r, err := c.Read("/", &response)
	assert.NoError(t, err, fmt.Sprintf("GET failed\n response: %+v", r))

	routes := routeSet{
		set: make(map[string]struct{}),
	}

	for _, r := range []string{
		"/",
		"proxy",
		"proxy/*",
		"/keystone/v3/auth/projects",
		"/keystone/v3/auth/tokens",

		// TODO WIP: should register?
		//"public",
		//"public/*",
		//"/contrail",
		//"/contrail/*",
	} {
		routes.add(r)
	}

	{
		contrailService := services.ContrailService{}
		contrailService.RegisterRESTAPI(&routes)
	}

	rs := apisrv.APIServer.Echo.Routes()
	data, _ := json.MarshalIndent(rs, "", "  ")
	fmt.Printf("%s\n", data)

	for _, route := range rs {
		assert.Truef(t, routes.contains(route.Path),
			"Route %s has no corresponding link in homepage discovery."+
				" Register it in APIServer setup code and add it to the set of excluded routes in the test.",
			route.Path)
	}
}

// mock an Echo server
func (r *routeSet) GET(path string, _ echo.HandlerFunc, _ ...echo.MiddlewareFunc) *echo.Route {
	r.add(path)
	return nil
}

func (r *routeSet) POST(path string, _ echo.HandlerFunc, _ ...echo.MiddlewareFunc) *echo.Route {
	r.add(path)
	return nil
}

func (r *routeSet) PUT(path string, _ echo.HandlerFunc, _ ...echo.MiddlewareFunc) *echo.Route {
	r.add(path)
	return nil
}

func (r *routeSet) DELETE(path string, _ echo.HandlerFunc, _ ...echo.MiddlewareFunc) *echo.Route {
	r.add(path)
	return nil
}
