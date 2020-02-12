package keystone_test

// TODO(mblotniak): refactor and move to ASF

import (
	"context"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testEndpointFile = "../apiserver/test_data/test_endpoint.tmpl"
)

func TestRemoteAuthenticate(t *testing.T) {
	// Test to verify the token validation using
	// the remote keystone configured in the endopoint
	keystoneAuthURL := viper.GetString("keystone.authurl")
	clusterXName := "clusterX"
	clusterXUser := clusterXName + "_admin"
	ksPrivate := integration.NewKeystoneServerFake(t, keystoneAuthURL, clusterXUser, clusterXUser)
	defer ksPrivate.Close()

	ksPublic := integration.NewKeystoneServerFake(t, keystoneAuthURL, clusterXUser, clusterXUser)
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

	ts, err := integration.LoadTest(testEndpointFile, pContext)
	require.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Delete the clusterX's keystone endpoint
	for _, client := range ts.Clients {
		ctx := context.Background()
		var response map[string]interface{}
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterXName)
		_, err := client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete clusterX's keystone endpoint")
		break
	}
	server.ForceProxyUpdate()
}
