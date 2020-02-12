package apiserver_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/contrail/pkg/db/cache"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/proxy"
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
			link(addr, "action", "POST", services.SyncPath),
			link(addr, "action", "POST", services.UserAgentKVPath),
			link(addr, "action", "POST", services.RefRelaxForDeletePath),
			link(addr, "action", "POST", services.PropCollectionUpdatePath),
			link(addr, "action", "POST", services.SetTagPath),
			link(addr, "action", "POST", services.ChownPath),
			link(addr, "action", "GET", services.IntPoolPath),
			link(addr, "action", "POST", services.IntPoolPath),
			link(addr, "action", "DELETE", services.IntPoolPath),
			link(addr, "action", "POST", services.IntPoolsPath),
			link(addr, "action", "DELETE", services.IntPoolsPath),
			link(addr, "action", "GET", services.ObjPerms),

			link(addr, "action", "POST", services.UploadCloudKeysPath),

			map[string]interface{}{
				"link": map[string]interface{}{
					"href":   resolve(addr, proxy.DefaultPath),
					"name":   proxy.DefaultPath,
					"rel":    "proxy",
					"method": nil,
				},
			},
			// static proxy
			map[string]interface{}{
				"link": map[string]interface{}{
					"href":   resolve(addr, "contrail"),
					"name":   "contrail",
					"rel":    "proxy",
					"method": nil,
				},
			},

			link(addr, "action", "GET", cache.WatchPath),

			link(addr, "action", "POST", "keystone/v3/auth/tokens"),
			link(addr, "action", "GET", "keystone/v3/auth/tokens"),
			link(addr, "collection", "GET", "keystone/v3/auth/projects"),
			link(addr, "collection", "GET", "keystone/v3/auth/domains"),
			link(addr, "collection", "GET", "keystone/v3/projects"),
			link(addr, "resource-base", "GET", "keystone/v3/projects"),
			link(addr, "collection", "GET", "keystone/v3/domains"),
			link(addr, "collection", "GET", "keystone/v3/users"),
		},
	}

	require.Contains(t, response, "href")
	assert.Equal(t, expected["href"], response["href"])

	require.Contains(t, response, "links")
	for _, link := range expected["links"].([]interface{}) {
		assert.Contains(t, response["links"], link, "the response does not contain link: %#v", link)
	}
}

func link(serverURL, rel, method, path string) map[string]interface{} {
	return map[string]interface{}{
		"link": map[string]interface{}{
			"href":   resolve(serverURL, path),
			"name":   path,
			"rel":    rel,
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
		proxyPath := proxy.ConfigFromViper().Path

		routes.add(resolve(proxyPath))
		routes.add(resolve(proxyPath, "*"))
	}
	for _, r := range []string{
		"/",
	} {
		routes.add(r)
	}

	for _, plugin := range []apiserver.APIPlugin{
		&cache.DB{},
		&keystone.Keystone{},
		services.UploadCloudKeysPlugin{},
		&services.ContrailService{},
	} {
		plugin.RegisterHTTPAPI(&routes)
	}

	// TODO(Witaut): Don't use Echo - an internal detail of Server.
	for _, route := range server.APIServer.Echo.Routes() {
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
func (r *routeSet) GET(path string, _ apiserver.HandlerFunc, _ ...apiserver.RouteOption) {
	r.add(path)
}

func (r *routeSet) POST(path string, _ apiserver.HandlerFunc, _ ...apiserver.RouteOption) {
	r.add(path)
}

func (r *routeSet) PUT(path string, _ apiserver.HandlerFunc, _ ...apiserver.RouteOption) {
	r.add(path)
}

func (r *routeSet) DELETE(path string, _ apiserver.HandlerFunc, _ ...apiserver.RouteOption) {
	r.add(path)
}

func (r *routeSet) Add(_, path string, _ apiserver.HandlerFunc, _ ...apiserver.RouteOption) {
	r.add(path)
}

func (r *routeSet) Use(_ ...apiserver.MiddlewareFunc) {
}

func (r *routeSet) Group(prefix string, _ ...apiserver.RouteOption) {
	r.add(prefix)
}

func (r *routeSet) Register(_, _, _, _ string) {
}

func resolve(base string, parts ...string) string {
	base = strings.TrimSuffix(base, "/")
	return strings.Join(append([]string{base}, parts...), "/")
}
