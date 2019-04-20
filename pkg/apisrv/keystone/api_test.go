package keystone_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/auth"
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
	assert.NoError(t, err, "failed to create new http request")
	defer resp.Body.Close() // nolint: errcheck

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
	// Verfiy token
	url := strings.Join(
		[]string{server.URL(), "contrail-cluster", clusterName + "_uuid"}, "/")
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: viper.GetBool("keystone.insecure")},
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	req.Header.Set("X-Auth-Token", commandServerToken)
	res, err := c.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close() // nolint: errcheck
	assert.NoError(t, err)
	contents, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, contents)

	// Cleanup test
	integration.RunCleanTestScenario(t, &testScenario, server)
}

func TestClusterLogin(t *testing.T) {
	keystoneAuthURL := viper.GetString("keystone.authurl")
	clusterName := "clusterB"
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

	ctx := context.Background()
	var clientScenario integration.TestScenario
	err = integration.LoadTestScenario(&clientScenario, testClusterTokenAPIFile, pContext)
	assert.NoError(t, err, "failed to load endpoint create test data")
	clients := integration.PrepareClients(ctx, t, &clientScenario, server)

	// preserve  infra token
	var infraToken string
	for _, client := range clients {
		infraToken = client.AuthToken
		break
	}
	clusterID := clusterName + "_uuid"
	t.Run(
		"login cluster with correct credentials",
		testClusterLoginWithCorrectCredential(ctx, clients, clusterID),
	)
	t.Run(
		"login to cluster with incorrect credentials",
		testClusterLoginWithIncorrectCredential(ctx, clients, clusterID),
	)
	t.Run(
		"login cluster with no credentials and no infra(superuser) token",
		testClusterLoginWithoutCredentialAndSuperUserToken(ctx, clients, clusterID),
	)
	t.Run(
		"login cluster with no credentials and with infra(superuser) token",
		testClusterLoginWithSuperUserToken(ctx, clients, clusterID, infraToken),
	)

	// Cleanup test
	integration.RunCleanTestScenario(t, &testScenario, server)
}

func testClusterLoginWithCorrectCredential(
	ctx context.Context, clients integration.ClientsList, clusterID string,
) func(*testing.T) {
	return func(t *testing.T) {
		for _, client := range clients {
			ctx = auth.WithXClusterID(ctx, clusterID)
			client.ID = defaultUser
			client.Password = defaultPassword
			client.Scope = nil
			_, err := client.Login(ctx)
			assert.NoError(t, err, "client failed to login cluster keystone")
		}
	}
}

func testClusterLoginWithIncorrectCredential(
	ctx context.Context, clients integration.ClientsList, clusterID string,
) func(*testing.T) {
	return func(t *testing.T) {
		for _, client := range clients {
			ctx = auth.WithXClusterID(ctx, clusterID)
			client.ID = defaultUser
			client.Password = "hacker"
			client.Scope = nil
			_, err := client.Login(ctx)
			assert.Error(t, err, "hacker logged in to cluster keystone !!")
		}
	}
}

func testClusterLoginWithoutCredentialAndSuperUserToken(
	ctx context.Context, clients integration.ClientsList, clusterID string,
) func(*testing.T) {
	return func(t *testing.T) {
		for _, client := range clients {
			ctx = auth.WithXClusterID(ctx, clusterID)
			client.ID = ""
			client.Password = ""
			client.Scope = nil
			_, err := client.Login(ctx)
			assert.Error(t, err, "hacker logged in to cluster keystone without credentials!!")
		}
	}
}

func testClusterLoginWithSuperUserToken(
	ctx context.Context, clients integration.ClientsList, clusterID, token string,
) func(*testing.T) {
	return func(t *testing.T) {
		for _, client := range clients {
			ctx = auth.WithXClusterID(ctx, clusterID)
			ctx = auth.WithXAuthToken(ctx, token)
			client.ID = ""
			client.Password = ""
			client.Scope = nil
			_, err := client.Login(ctx)
			assert.NoError(t, err, "client failed to login cluster keystone with token")
		}
	}
}

func TestBasicAuth(t *testing.T) {
	pContext := pongo2.Context{
		"cluster_name":  clusterName,
		"endpoint_name": clusterName + "_config",
		"private_url":   ksPrivate.URL,
		"public_url":    ksPublic.URL,
	}
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: false,
	})
	defer s.CloseT(t)

	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testClusterTokenAPIFile, pContext)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()

	server.ForceProxyUpdate()
}
