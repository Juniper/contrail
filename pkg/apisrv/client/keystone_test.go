package client_test

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestKeystoneClient(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		RepoRootPath: "../../..",
	})
	defer s.CloseT(t)

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

	p, err := k.GetProject(context.Background(), token, integration.AdminProjectID)
	assert.NoError(t, err)
	assert.Equal(t, integration.AdminProjectID, p.ID)
	assert.Equal(t, integration.AdminProjectName, p.Name)
}
