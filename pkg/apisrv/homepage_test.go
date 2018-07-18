package apisrv_test

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestHomepage(t *testing.T) {
	c := integration.NewHTTPAPIClient(t, apisrv.TestServer.URL)

	var response map[string]interface{}
	r, err := c.Read("/", &response)
	assert.NoError(t, err, fmt.Sprintf("GET failed\n response: %+v", r))

	addr := viper.GetString("server.address")
	expected := map[string]interface{}{
		"href": addr,
		"links": []interface{}{
			map[string]interface{}{
				"href":   path.Join(addr, "virtual-network"),
				"name":   "virtual-network",
				"rel":    "resource-base",
				"method": nil,
			},
			map[string]interface{}{
				"href":   path.Join(addr, "virtual-networks"),
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

	routes := apisrv.APIServer.Echo.Routes()
	data, _ := json.MarshalIndent(routes, "", "  ")
	fmt.Printf("%s\n", data)

	excludedRoutes := []string{
		"/",
		"proxy",
		"proxy/*",
		"/keystone/v3/auth/projects",
		"/keystone/v3/auth/tokens",

		// TODO WIP: should register?
		"public",
		"public/*",
		"/contrail",
		"/contrail/*",
		"sync",
	}
	discoveredHrefs := linksToHrefs(response["links"].([]interface{}), addr)
	for _, route := range routes {
		if contains(excludedRoutes, route.Path) {
			continue
		}

		assert.Truef(t, containsPrefix(discoveredHrefs, route.Path),
			"Endpoint %s has no corresponding link in homepage discovery. Register it in APIServer setup code.",
			route.Path)
	}
}

func linksToHrefs(links []interface{}, addr string) []string {
	var result []string
	addr = strings.TrimSuffix(addr, "/")
	for _, link := range links {
		href := strings.TrimPrefix(link.(map[string]interface{})["href"].(string), addr)
		result = append(result, href)
	}
	return result
}

func contains(list []string, s string) bool {
	for _, str := range list {
		if str == s {
			return true
		}
	}
	return false
}

func containsPrefix(list []string, of string) bool {
	for _, str := range list {
		if strings.HasPrefix(of, str) {
			return true
		}
	}
	return false
}
