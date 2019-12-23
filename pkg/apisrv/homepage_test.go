package apisrv_test

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHomepageResources(t *testing.T) {
	c := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	var response map[string]interface{}
	r, err := c.Read(context.Background(), "/", &response)
	assert.NoError(t, err, fmt.Sprintf("GET failed\n response: %+v", r))

	addr := server.URL()
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
	c := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	var response map[string]interface{}
	r, err := c.Read(context.Background(), "/", &response)
	assert.NoError(t, err, fmt.Sprintf("GET failed\n response: %+v", r))

	routes := routeSet{
		set: make(map[string]struct{}),
	}

	excludedRoutesRegexes, err := compileRegexStrings(
		[]string{
			"^/neutron/*",
		},
	)
	assert.NoError(t, err)

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
	if viper.GetBool("cache.enabled") {
		routes.add(apisrv.WatchPath)
	}

	for _, r := range []string{
		"/",

		"/keystone/v3/projects",
		"/keystone/v3/domains",
		"/keystone/v3/projects/:id",
		"/keystone/v3/auth/projects", // TODO: Remove this, since "/keystone/v3/projects" is a keystone endpoint
		"/keystone/v3/auth/domains",  // TODO: Remove this, since "/keystone/v3/domains" is a keystone endpoint
		"/keystone/v3/auth/tokens",
		services.UploadCloudKeysPath,
	} {
		routes.add(r)
	}

	// Service resources are registered in server.go:setupHomepage().
	{
		contrailService := services.ContrailService{}
		contrailService.RegisterRESTAPI(&routes)
	}

	// Action resources are registered in server.go:setupActionResources()
	// and in service_common.tmpl:RegisterRESTAPI().
	routes.add(apisrv.FQNameToIDPath)
	routes.add(apisrv.IDToFQNamePath)
	routes.add(apisrv.UserAgentKVPath)
	routes.add(services.PropCollectionUpdatePath)

	for _, route := range server.APIServer.BaseServer.Echo.Routes() {
		var isPathExcludedFromHomepage bool
		for _, excludedRegex := range excludedRoutesRegexes {
			if excludedRegex.MatchString(route.Path) {
				isPathExcludedFromHomepage = true
				break
			}
		}
		if !isPathExcludedFromHomepage {
			assert.Truef(t, routes.contains(route.Path),
				"Route %s has no corresponding link in homepage discovery."+
					" Register it in APIServer setup code or add it to the set of excluded routes in the test.",
				route.Path)
		}
	}
}

type routeSet struct {
	set map[string]struct{}
}

func (r *routeSet) add(path string) {
	r.set[resolve(path)] = struct{}{}
}

func (r *routeSet) contains(path string) bool {
	_, result := r.set[resolve(path)]
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

func compileRegexStrings(stringsToCompile []string) ([]*regexp.Regexp, error) {
	compiledRegexes := make([]*regexp.Regexp, len(stringsToCompile))
	var err error
	for i, str := range stringsToCompile {
		compiledRegexes[i], err = regexp.Compile(str)
		if err != nil {
			return nil, err
		}
	}
	return compiledRegexes, nil
}
