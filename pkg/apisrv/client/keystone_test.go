package client_test

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestGetProject(t *testing.T) {
	defer runAPIServer(t).CloseT(t)
	keystone, token := keystoneClientAndToken(t)

	p, err := keystone.GetProject(context.Background(), token, integration.AdminProjectID)
	assert.NoError(t, err)
	assert.Equal(t, integration.AdminProjectID, p.ID)
	assert.Equal(t, integration.AdminProjectName, p.Name)
}

func TestGetProjects(t *testing.T) {
	defer runAPIServer(t).CloseT(t)
	keystone, token := keystoneClientAndToken(t)

	projects, err := keystone.GetProjects(context.Background(), token)
	assert.NoError(t, err)
	assert.Len(t, projects, 1)
	assert.Equal(t, integration.AdminProjectID, projects[0].ID)
	assert.Equal(t, integration.AdminProjectName, projects[0].Name)
}

func runAPIServer(t *testing.T) *integration.APIServer {
	return integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: false,
	})
}

func keystoneClientAndToken(t *testing.T) (*client.Keystone, string) {
	k := &client.Keystone{
		URL: viper.GetString("keystone.authurl"),
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: viper.GetBool("keystone.insecure")},
			},
		},
	}

	resp, err := k.ObtainToken(context.Background(), integration.AdminUserID, integration.AdminUserPassword, nil)
	assert.NoError(t, err)
	token := resp.Header.Get("X-Subject-Token")
	assert.NotEmpty(t, token)

	return k, token
}
