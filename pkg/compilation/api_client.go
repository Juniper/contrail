package compilation

import (
	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/compilation/config"
	"github.com/Juniper/contrail/pkg/keystone"
)

func newAPIClient(config config.Config) *client.HTTP {
	c := config.APIClientConfig
	restClient := client.NewHTTP(
		c.URL,
		c.AuthURL,
		c.ID,
		c.Password,
		c.Insecure,
		keystone.GetScope(
			c.DomainID, c.DomainName, c.ProjectID, c.ProjectName,
		),
	)
	restClient.Init()

	return restClient
}
