package apisrv_test

import (
	"context"
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
	r, err := c.Read(context.Background(), "/", &response)
	assert.NoError(t, err, fmt.Sprintf("GET failed\n response: %+v", r))

	addr := resolve(viper.GetString("server.address"))
	expected := map[string]interface{}{
		"href": addr,
		"links": []interface{}{
			map[string]interface{}{
				"link": map[string]interface{}{
					"href":   resolve(addr, "virtual-network"),
					"name":   "virtual-network",
					"rel":    "resource-base",
					"method": nil,
				},
			},
			map[string]interface{}{
				"link": map[string]interface{}{
					"href":   resolve(addr, "virtual-networks"),
					"name":   "virtual-network",
					"rel":    "collection",
					"method": nil,
				},
			},
		},
	}

	require.Contains(t, response, "href")
	assert.Equal(t, expected["href"], response["href"])

	require.Contains(t, response, "links")
	assert.Subset(t, response["links"], expected["links"])
}

func TestRoutesAreRegistered(t *testing.T) {
	c := integration.NewHTTPAPIClient(t, apisrv.TestServer.URL)

	var response map[string]interface{}
	r, err := c.Read(context.Background(), "/", &response)
	assert.NoError(t, err, fmt.Sprintf("GET failed\n response: %+v", r))

	routes := routeSet{
		set: make(map[string]struct{}),
	}

	// Not for discovery
	for _, configKey := range []string{
		"server.static_files",
		"server.proxy",
	} {
		for path := range viper.GetStringMapString(configKey) {
			routes.add(resolve(path))
			routes.add(resolve(path, "*"))
		}
	}

	{
		proxyPath := apisrv.DefaultDynamicProxyPath
		if p := viper.GetString("server.dynamic_proxy_path"); p != "" {
			proxyPath = p
		}

		routes.add(resolve(proxyPath))
		routes.add(resolve(proxyPath, "*"))
	}

	for _, r := range []string{
		"/",

		"/keystone/v3/auth/projects",
		"/keystone/v3/auth/tokens",
	} {
		routes.add(r)
	}

	// Service resources are registered in server.go:setupHomepage().
	{
		contrailService := services.ContrailService{}
		contrailService.RegisterRESTAPI(&routes)
	}

	routes.add("/fqname-to-id")

	for _, route := range apisrv.APIServer.Echo.Routes() {
		assert.Truef(t, routes.contains(route.Path),
			"Route %s has no corresponding link in homepage discovery."+
				" Register it in APIServer setup code and add it to the set of excluded routes in the test.",
			route.Path)
	}
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

func resolve(base string, parts ...string) string {
	base = strings.TrimSuffix(base, "/")
	return strings.Join(append([]string{base}, parts...), "/")
}
