package apisrv_test

import (
	"fmt"
	"path"
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
}
