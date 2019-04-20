package keystone_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	privatePortList  = "private_port_list"
	publicPortList   = "public_port_list"
	testEndpointFile = "../test_data/test_endpoint.tmpl"
)

func TestRemoteAuthenticate(t *testing.T) {
	// Test to verify the token validation using
	// the remote keystone configured in the endopoint
	ctx := context.Background()
	keystoneAuthURL := viper.GetString("keystone.authurl")
	clusterXName := "clusterX"
	clusterXUser := clusterXName + "_admin"
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterXUser, clusterXUser)
	defer ksPrivate.Close()

	ksPublic := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, clusterXUser, clusterXUser)
	defer ksPublic.Close()

	pContext := pongo2.Context{
		"cluster_name":    clusterXName,
		"endpoint_name":   clusterXName + "_keystone",
		"endpoint_prefix": "keystone",
		"private_url":     ksPrivate.URL,
		"public_url":      ksPublic.URL,
		"manage_parent":   true,
		"admin_user":      clusterXUser,
	}

	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testEndpointFile, pContext)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Delete the clusterX's keystone endpoint
	for _, client := range testScenario.Clients {
		ctx = context.Background()
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterXName)
		_, err := client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete clusterX's keystone endpoint")
		break
	}
	server.ForceProxyUpdate()
}
