package keystone_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	defaultUser             = "admin"
	defaultPassword         = "contrail123"
	testClusterTokenAPIFile = "./test_data/test_cluster_token_method.yml"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func FetchCommandServerToken(
	t *testing.T, clusterID string, clusterToken string) string {
	dataJSON, err := json.Marshal(&keystone.UnScopedAuthRequest{
		Auth: &keystone.UnScopedAuth{
			Identity: &keystone.Identity{
				Methods: []string{"cluster_token"},
				Cluster: &keystone.Cluster{
					ID: clusterID,
					Token: &keystone.UserToken{
						ID: clusterToken,
					},
				},
			},
		},
	})
	assert.NoError(t, err, "failed to marshal cluster_token request")
	keystoneAuthURL := viper.GetString("keystone.authurl")
	k := &client.Keystone{
		URL: keystoneAuthURL,
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: viper.GetBool("keystone.insecure")},
			},
		},
	}
	request, err := http.NewRequest(
		"POST", keystoneAuthURL+"/auth/tokens", bytes.NewBuffer(dataJSON))
	assert.NoError(t, err, "failed to create new http request")
	request.Header.Set("Content-Type", "application/json")

	resp, err := k.HTTPClient.Do(request)
	defer resp.Body.Close() // nolint: errcheck
	assert.NoError(t, err, "failed to create new http request")

	token := resp.Header.Get("X-Subject-Token")
	return token
}

func TestClusterTokenMethod(t *testing.T) {
	keystoneAuthURL := viper.GetString("keystone.authurl")
	clusterName := "clusterA"
	ksPrivate := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, defaultUser, defaultPassword)
	defer ksPrivate.Close()

	ksPublic := integration.MockServerWithKeystoneTestUser(
		"", keystoneAuthURL, defaultUser, defaultPassword)
	defer ksPublic.Close()
	pContext := pongo2.Context{
		"cluster_name":  clusterName,
		"endpoint_name": clusterName + "_keystone",
		"private_url":   ksPrivate.URL,
		"public_url":    ksPublic.URL,
	}

	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testClusterTokenAPIFile, pContext)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Fetch cluster keystone token
	k := &client.Keystone{
		URL:        ksPrivate.URL + "/v3",
		HTTPClient: &http.Client{},
	}
	resp, err := k.ObtainToken(
		context.Background(), defaultUser, defaultPassword, nil)
	assert.NoError(t, err)
	token := resp.Header.Get("X-Subject-Token")
	assert.NotEmpty(t, token)
	// Fetch command keystone token with cluster keystone token
	commandServerToken := FetchCommandServerToken(t, clusterName+"_uuid", token)
	assert.NotEmpty(t, commandServerToken)

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
