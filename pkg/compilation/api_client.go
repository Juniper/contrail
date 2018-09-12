package compilation

import (
	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/compilation/config"
)

func newAPIClient(config config.Config) *client.HTTP {
	c := config.APIClientConfig
	restClient := client.NewHTTP(
		c.URL,
		c.AuthURL,
		c.ID,
		c.Password,
		c.Domain,
		c.Insecure,
		&keystone.Scope{
			Project: &keystone.Project{
				ID:     c.Project,
				Domain: &keystone.Domain{},
			},
		},
	)
	restClient.Init()

	return restClient
}
