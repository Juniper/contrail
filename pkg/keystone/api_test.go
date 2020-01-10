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

	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	kstypes "github.com/Juniper/asf/pkg/keystone"
)

const (
	defaultUser             = "admin"
	defaultPassword         = "contrail123"
	testClusterTokenAPIFile = "./test_data/test_cluster_token_method.tmpl"
	testBasicAuthFile       = "./test_data/test_basic_auth.tmpl"
	openstack               = "openstack"
)

func TestClusterTokenMethod(t *testing.T) {
	keystoneAuthURL := viper.GetString("keystone.authurl")
	clusterName := t.Name() + "_clusterA"
	ksPrivate := integration.NewKeystoneServerFake(t, keystoneAuthURL, defaultUser, defaultPassword)
	defer ksPrivate.Close()

	ksPublic := integration.NewKeystoneServerFake(t, keystoneAuthURL, defaultUser, defaultPassword)
	defer ksPublic.Close()
	pContext := pongo2.Context{
		"cluster_name":  clusterName,
		"endpoint_name": clusterName + "_keystone",
		"private_url":   ksPrivate.URL,
		"public_url":    ksPublic.URL,
	}

	ts, err := integration.LoadTest(testClusterTokenAPIFile, pContext)
	require.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	server.ForceProxyUpdate()

	// Fetch cluster keystone token
	k := &kstypes.Client{
		URL:      ksPrivate.URL + "/v3",
		HTTPDoer: &http.Client{},
	}
	token, err := k.ObtainToken(context.Background(), defaultUser, defaultPassword, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	// Fetch command keystone token with cluster keystone token
	commandServerToken := fetchCommandServerToken(t, clusterName+"_uuid", token)
	assert.NotEmpty(t, commandServerToken)
	// Verify token
	url := strings.Join([]string{server.URL(), "contrail-cluster", clusterName + "_uuid"}, "/")
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

	// Cleanup test TODO: Fix cleanup to remove all resources
	integration.RunCleanTestScenario(t, ts, server)
}

func fetchCommandServerToken(t *testing.T, clusterID string, clusterToken string) string {
	dataJSON, err := json.Marshal(&kstypes.UnScopedAuthRequest{
		Auth: &kstypes.UnScopedAuth{
			Identity: &kstypes.Identity{
				Methods: []string{"cluster_token"},
				Cluster: &kstypes.Cluster{
					ID: clusterID,
					Token: &kstypes.UserToken{
						ID: clusterToken,
					},
				},
			},
		},
	})
	assert.NoError(t, err, "failed to marshal cluster_token request")
	keystoneAuthURL := viper.GetString("keystone.authurl")
	k := &kstypes.Client{
		URL: keystoneAuthURL,
		HTTPDoer: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: viper.GetBool("keystone.insecure")},
			},
		},
	}
	request, err := http.NewRequest("POST", keystoneAuthURL+"/auth/tokens", bytes.NewBuffer(dataJSON))
	assert.NoError(t, err, "failed to create new http request")
	request.Header.Set("Content-Type", "application/json")

	resp, err := k.HTTPDoer.Do(request)
	assert.NoError(t, err, "failed to create new http request")
	defer resp.Body.Close() // nolint: errcheck

	token := resp.Header.Get("X-Subject-Token")
	return token
}

func TestClusterLogin(t *testing.T) {
	keystoneAuthURL := viper.GetString("keystone.authurl")
	clusterName := t.Name() + "_clusterB"
	ksPrivate := integration.NewKeystoneServerFake(t, keystoneAuthURL, defaultUser, defaultPassword)
	defer ksPrivate.Close()

	ksPublic := integration.NewKeystoneServerFake(t, keystoneAuthURL, defaultUser, defaultPassword)
	defer ksPublic.Close()
	pContext := pongo2.Context{
		"cluster_name":  clusterName,
		"endpoint_name": clusterName + "_keystone",
		"private_url":   ksPrivate.URL,
		"public_url":    ksPublic.URL,
	}

	ts, err := integration.LoadTest(testClusterTokenAPIFile, pContext)
	require.NoError(t, err, "failed to load endpoint create test data")
	cleanup := integration.RunDirtyTestScenario(t, ts, server)
	defer cleanup()

	server.ForceProxyUpdate()

	ctx := context.Background()
	ts, err = integration.LoadTest(testClusterTokenAPIFile, pContext)
	require.NoError(t, err, "failed to load endpoint create test data")
	clients := integration.PrepareClients(ctx, t, ts, server)

	// preserve infra token
	var infraToken string
	for _, client := range clients {
		infraToken = client.AuthToken
		break
	}
	clusterID := clusterName + "_uuid"
	tests := []struct {
		desc     string
		id       string
		password string
		token    string
		wantErr  bool
	}{{
		desc:     "login cluster with correct credentials",
		id:       defaultUser,
		password: defaultPassword,
	}, {
		desc:     "login to cluster with incorrect credentials",
		id:       defaultUser,
		password: "hacker",
		wantErr:  true,
	}, {
		desc:    "login cluster with no credentials and no infra(superuser) token",
		wantErr: true,
	}, {
		desc:  "login cluster with no credentials and with infra(superuser) token",
		token: infraToken,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			for _, client := range clients {
				ctx = auth.WithXClusterID(ctx, clusterID)
				if tt.token != "" {
					ctx = kstypes.WithXAuthToken(ctx, tt.token)
				}
				client.ID = tt.id
				client.Password = tt.password
				client.Scope = nil
				if err := client.Login(ctx); !tt.wantErr {
					assert.NoError(t, err, "unexpected error")
				} else {
					assert.Error(t, err, "got error")
				}
			}
		})
	}

	// Cleanup test
	integration.RunCleanTestScenario(t, ts, server)
}

func TestMultiClusterAuth(t *testing.T) {
	t.Skip("skipping flaky test, will fix it in master.")
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		RepoRootPath: "../../..",
	})
	defer s.CloseT(t)

	openstackClusterName := t.Name() + "_cluster_openstack"
	kubernetesClusterName := t.Name() + "_cluster_kubernetes"
	clusters := map[string]struct {
		id               string
		orchestrator     string
		expectedProjects []string
		expectedDomains  []string
	}{
		openstackClusterName: {
			id:               openstackClusterName + "_uuid",
			orchestrator:     openstack,
			expectedProjects: []string{integration.AdminProjectID, integration.NeutronProjectID},
			expectedDomains:  []string{integration.DefaultDomainID},
		},
		kubernetesClusterName: {
			id:               kubernetesClusterName + "_uuid",
			orchestrator:     "kubernetes",
			expectedProjects: []string{kubernetesClusterName + "_uuid"},
			expectedDomains:  []string{kubernetesClusterName + "_uuid"},
		},
	}

	var err error
	var ts *integration.TestScenario
	// Create clusterA, clusterB and their config endpoint
	for clusterName, cluster := range clusters {
		configService := mockConfigServer(clusterName, cluster.id)
		pContext := pongo2.Context{
			"cluster_name":  clusterName,
			"endpoint_name": clusterName + "_config",
			"private_url":   configService.URL,
			"public_url":    configService.URL,
			"orchestrator":  cluster.orchestrator,
		}
		if cluster.orchestrator == openstack {
			// mock keystone for cluster with openstack orchestrator
			keystoneAuthURL := viper.GetString("keystone.authurl")
			ksPublic := integration.NewKeystoneServerFake(t, keystoneAuthURL, defaultUser, defaultPassword)
			defer ksPublic.Close()
			ksPrivate := integration.NewKeystoneServerFake(t, keystoneAuthURL, defaultUser, defaultPassword)
			defer ksPrivate.Close()
		}
		ts, err = integration.LoadTest(testBasicAuthFile, pContext)
		require.NoError(t, err, "failed to load endpoint create test data")
		cleanup := integration.RunDirtyTestScenario(t, ts, server)
		defer cleanup()
	}

	server.ForceProxyUpdate()

	// verify basic auth
	for clusterName, cluster := range clusters {
		ctx := auth.WithXClusterID(context.Background(), cluster.id)
		url := "/keystone/v3/auth/projects"
		ok := verifyBasicAuthProjects(ctx, t, ts, url, cluster.expectedProjects)
		assert.True(t, ok, "failed to get project list from cluster: %s", clusterName)

		url = "/keystone/v3/auth/domains"
		ok = verifyBasicAuthDomains(ctx, t, ts, url, cluster.expectedDomains)
		assert.True(t, ok, "failed to get domain list from cluster: %s", clusterName)
	}
}

func mockConfigServer(clusterName, clusterID string) *httptest.Server {
	// Echo instance
	e := echo.New()

	// Routes
	for route, handler := range map[string]interface{}{
		"/domains": echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK,
				&keystone.VncDomainListResponse{
					Domains: []*keystone.VncDomain{{
						Domain: &keystone.ConfigDomain{
							Name: clusterName,
							UUID: clusterID,
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
							UUID:   clusterID,
							FQName: []string{clusterName},
						},
					},
					},
				},
			)
		}),
	} {

		e.GET(route, handler.(echo.HandlerFunc))
	}
	mockServer := httptest.NewServer(e)
	return mockServer
}

func verifyBasicAuthProjects(
	ctx context.Context, t *testing.T, testScenario *integration.TestScenario,
	url string, projects []string) bool {
	for _, client := range testScenario.Clients {
		projectList := kstypes.ProjectListResponse{}
		_, err := client.Read(ctx, url, &projectList)
		if err != nil {
			fmt.Printf("Reading: %s, Response: %v", url, err)
			return false
		}
		if len(projectList.Projects) != len(projects) {
			fmt.Printf("Unexpected projects: %v", projectList)
			return false
		}
		for _, project := range projectList.Projects {
			ok := assert.True(
				t, format.ContainsString(projects, project.ID),
				fmt.Sprintf("Unexpected project: %v", project))
			if !ok {
				return ok
			}
		}
	}
	return true
}

func verifyBasicAuthDomains(
	ctx context.Context, t *testing.T, testScenario *integration.TestScenario,
	url string, domains []string) bool {
	for _, client := range testScenario.Clients {
		domainList := kstypes.DomainListResponse{}
		_, err := client.Read(ctx, url, &domainList)
		if err != nil {
			fmt.Printf("Reading: %s, Response: %s", url, err)
			return false
		}
		if len(domainList.Domains) != len(domains) {
			fmt.Printf("Unexpected domains: %v", domainList)
			return false
		}
		for _, domain := range domainList.Domains {
			ok := assert.True(
				t, format.ContainsString(domains, domain.ID),
				fmt.Sprintf("Unexpected domain: %v", domain))
			if !ok {
				return ok
			}
		}
	}
	return true
}
