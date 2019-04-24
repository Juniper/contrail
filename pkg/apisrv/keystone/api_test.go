package keystone_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/db/basedb"
	kscommon "github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	defaultUser             = "admin"
	defaultPassword         = "contrail123"
	testClusterTokenAPIFile = "./test_data/test_cluster_token_method.yml"
	testBasicAuthFile       = "./test_data/test_basic_auth.yml"
)

var server *integration.APIServer

func mockServer(routes map[string]interface{}) *httptest.Server {
	// Echo instance
	e := echo.New()

	// Routes
	for route, handler := range routes {
		e.GET(route, handler.(echo.HandlerFunc))
	}
	mockServer := httptest.NewServer(e)
	return mockServer
}

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func FetchCommandServerToken(
	t *testing.T, clusterID string, clusterToken string) string {
	dataJSON, err := json.Marshal(&kscommon.UnScopedAuthRequest{
		Auth: &kscommon.UnScopedAuth{
			Identity: &kscommon.Identity{
				Methods: []string{"cluster_token"},
				Cluster: &kscommon.Cluster{
					ID: clusterID,
					Token: &kscommon.UserToken{
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

func verifyBasicAuthDomains(
	ctx context.Context, t *testing.T, testScenario *integration.TestScenario,
	url string, clusterName string) bool {
	for _, client := range testScenario.Clients {
		domainList := keystone.DomainListResponse{}
		_, err := client.Read(ctx, url, &domainList)
		if err != nil {
			fmt.Printf("Reading: %s, Response: %s", url, err)
			return false
		}
		if len(domainList.Domains) != 1 {
			fmt.Printf("Unexpected domains: %s", domainList)
			return false
		}
		ok := testutil.AssertEqual(
			t, clusterName, domainList.Domains[0].Name,
			fmt.Sprintf("Unexpected name in domain: %s", domainList))
		if !ok {
			return ok
		}
		ok = testutil.AssertEqual(
			t, clusterName+"_uuid", domainList.Domains[0].ID,
			fmt.Sprintf("Unexpected uuid in domain: %s", domainList))
		if !ok {
			return ok
		}
	}
	return true
}

func verifyBasicAuthProjects(
	ctx context.Context, t *testing.T, testScenario *integration.TestScenario,
	url string, clusterName string) bool {
	for _, client := range testScenario.Clients {
		projectList := keystone.ProjectListResponse{}
		_, err := client.Read(ctx, url, &projectList)
		if err != nil {
			fmt.Printf("Reading: %s, Response: %s", url, err)
			return false
		}
		if len(projectList.Projects) != 1 {
			fmt.Printf("Unexpected projects: %s", projectList)
			return false
		}
		ok := testutil.AssertEqual(
			t, clusterName, projectList.Projects[0].Name,
			fmt.Sprintf("Unexpected name in project: %s", projectList))
		if !ok {
			return ok
		}
		ok = testutil.AssertEqual(
			t, clusterName+"_uuid", projectList.Projects[0].ID,
			fmt.Sprintf("Unexpected uuid in project: %s", projectList))
		if !ok {
			return ok
		}
		ok = testutil.AssertEqual(
			t, clusterName+"_uuid", projectList.Projects[0].ParentID,
			fmt.Sprintf("Unexpected parent_id in project: %s", projectList))
		if !ok {
			return ok
		}
	}
	return true
}

func TestBasicAuth(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: false,
		AuthType:           "basic-auth",
	})
	defer s.CloseT(t)

	clusterName := "clusterBasicAuth"
	routes := map[string]interface{}{
		"/domains": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK,
				&keystone.VncDomainListResponse{
					Domains: []*keystone.VncDomain{{
						Domain: &keystone.ConfigDomain{
							Name: clusterName,
							UUID: clusterName + "_uuid",
						},
					},
					},
				},
			)
		}),
		"/projects": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK,
				&keystone.VncProjectListResponse{
					Projects: []*keystone.VncProject{{
						Project: &keystone.ConfigProject{
							Name:   clusterName,
							UUID:   clusterName + "_uuid",
							FQName: []string{clusterName},
						},
					},
					},
				},
			)
		}),
	}
	configService := mockServer(routes)

	pContext := pongo2.Context{
		"cluster_name":  clusterName,
		"endpoint_name": clusterName + "_config",
		"private_url":   configService.URL,
		"public_url":    configService.URL,
	}
	var testScenario integration.TestScenario
	err := integration.LoadTestScenario(&testScenario, testBasicAuthFile, pContext)
	assert.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, &testScenario, server)
	defer cleanup()

	server.ForceProxyUpdate()

	ctx := context.Background()
	url := "/keystone/v3/auth/projects"
	ok := verifyBasicAuthProjects(ctx, t, &testScenario, url, clusterName)
	assert.True(t, ok, "failed to get project list from config %s", url)

	url = "/keystone/v3/auth/domains"
	ok = verifyBasicAuthDomains(ctx, t, &testScenario, url, clusterName)
	assert.True(t, ok, "failed to get domain list from config %s", url)
}
