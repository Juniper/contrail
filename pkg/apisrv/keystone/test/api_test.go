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
	testClusterTokenApiFile = "./test_data/test_cluster_token_method.yml"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func TestClusterTokenMethod(t *testing.T) {
	keystoneAuthURL := viper.GetString("keystone.authurl")
	ksPrivate := integration.MockServerWithKeystone("", keystoneAuthURL)
	defer ksPrivate.Close()

	ksPublic := integration.MockServerWithKeystone("", keystoneAuthURL)
	defer ksPublic.Close()

	clusterName := "clusterA"
	pContext := pongo2.Context{
		"cluster_name":  clusterName,
		"endpoint_name": clusterName + "_keystone",
		"private_url":   ksPrivate.URL,
		"public_url":    ksPublic.URL,
	}

	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testClusterTokenApiFile, pContext)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Cleanup endpoint test
	ctx := context.Background()
	for _, client := range testScenario.Clients {
		var response map[string]interface{}
		_, err = client.Login(ctx)
		assert.NoError(t, err, "failed to login cluster keystone")
		url := fmt.Sprintf("/endpoint/endpoint_%s_keystone_uuid", clusterName)
		_, err = client.Delete(ctx, url, &response)
		assert.NoError(t, err, "failed to delete keystone endpoint")
		break
	}
}
