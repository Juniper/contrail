package apisrv_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/Juniper/asf/pkg/apisrv/baseapisrv"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
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

			map[string]interface{}{
				"link": map[string]interface{}{
					"href":   resolve(addr, "fqname-to-id"),
					"name":   "name-to-id",
					"rel":    "action",
					"method": "POST",
				},
			},
			map[string]interface{}{
				"link": map[string]interface{}{
					"href":   resolve(addr, "id-to-fqname"),
					"name":   "id-to-name",
					"rel":    "action",
					"method": "POST",
				},
			},
			defaultLink(addr, "POST", services.SyncPath),
			defaultLink(addr, "POST", services.UserAgentKVPath),
			defaultLink(addr, "POST", services.RefRelaxForDeletePath),
			defaultLink(addr, "POST", services.PropCollectionUpdatePath),
			defaultLink(addr, "POST", services.SetTagPath),
			defaultLink(addr, "POST", services.ChownPath),
			defaultLink(addr, "GET", services.IntPoolPath),
			defaultLink(addr, "POST", services.IntPoolPath),
			defaultLink(addr, "DELETE", services.IntPoolPath),
			defaultLink(addr, "POST", services.IntPoolsPath),
			defaultLink(addr, "DELETE", services.IntPoolsPath),
			defaultLink(addr, "GET", services.ObjPerms),

			defaultLink(addr, "POST", services.UploadCloudKeysPath),

			map[string]interface{}{
				"link": map[string]interface{}{
					"href": resolve(addr, apisrv.DefaultDynamicProxyPath),
					"name": apisrv.DefaultDynamicProxyPath,
					// TODO Use a different "rel"?
					"rel":    "action",
					"method": nil,
				},
			},
			// static proxy
			map[string]interface{}{
				"link": map[string]interface{}{
					"href": resolve(addr, "contrail"),
					"name": "contrail",
					// TODO Use a different "rel"?
					"rel":    "action",
					"method": nil,
				},
			},

			// TODO Use a different "rel"?
			defaultLink(addr, "GET", cache.WatchPath),

			defaultLink(addr, "POST", "keystone/v3/auth/tokens"),
			defaultLink(addr, "GET", "keystone/v3/auth/tokens"),
			defaultLink(addr, "GET", "keystone/v3/auth/projects"),
			defaultLink(addr, "GET", "keystone/v3/auth/domains"),
			// TODO Check whether there are duplicates.
			// TODO Use a different "rel"?
			defaultLink(addr, "GET", "keystone/v3/projects"),
			defaultLink(addr, "GET", "keystone/v3/domains"),
			defaultLink(addr, "GET", "keystone/v3/users"),
		},
	}

	require.Contains(t, response, "href")
	assert.Equal(t, expected["href"], response["href"])

	require.Contains(t, response, "links")
	assert.Subset(t, response["links"], expected["links"])
}

func defaultLink(serverURL, method, path string) map[string]interface{} {
	return map[string]interface{}{
		"link": map[string]interface{}{
			"href":   resolve(serverURL, path),
			"name":   path,
			"rel":    "action",
			"method": method,
		},
	}
}

func TestRoutesAreRegistered(t *testing.T) {
	c := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

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

	// TODO(Witaut): Use staticProxyPlugin directly instead.
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
	} {
		routes.add(r)
	}

	for _, plugin := range []baseapisrv.APIPlugin{
		&cache.DB{},
		&keystone.Keystone{},
		services.UploadCloudKeysPlugin{},
		&services.ContrailService{},
	} {
		plugin.RegisterHTTPAPI(&routes)
	}

	// TODO(Witaut): Don't use Echo - an internal detail of Server.
	for _, route := range server.APIServer.Server.Echo.Routes() {
		if !assert.Truef(t, routes.contains(route.Path),
			"Route %s has no corresponding link in homepage discovery."+
				" Register it in APIServer setup code or add it to the set of excluded routes in the test.",
			route.Path) {
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
func (r *routeSet) GET(path string, _ baseapisrv.HandlerFunc, _ ...baseapisrv.RouteOption) {
	r.add(path)
}

func (r *routeSet) POST(path string, _ baseapisrv.HandlerFunc, _ ...baseapisrv.RouteOption) {
	r.add(path)
}

func (r *routeSet) PUT(path string, _ baseapisrv.HandlerFunc, _ ...baseapisrv.RouteOption) {
	r.add(path)
}

func (r *routeSet) DELETE(path string, _ baseapisrv.HandlerFunc, _ ...baseapisrv.RouteOption) {
	r.add(path)
}

func (r *routeSet) Add(_, path string, _ baseapisrv.HandlerFunc, _ ...baseapisrv.RouteOption) {
	r.add(path)
}

func (r *routeSet) Use(_ ...baseapisrv.MiddlewareFunc) {
}

func (r *routeSet) Group(prefix string, _ ...baseapisrv.RouteOption) {
	r.add(prefix)
}

func (r *routeSet) Register(_, _, _, _ string) {
}

func resolve(base string, parts ...string) string {
	base = strings.TrimSuffix(base, "/")
	return strings.Join(append([]string{base}, parts...), "/")
}
