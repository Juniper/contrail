package compilation

import (
	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/compilation/config"
)

func newAPIClient(config config.Config) *client.HTTP {
	c := config.APIClientConfig
	restClient := client.NewHTTP(
		c.URL,
		c.AuthURL,
		c.ID,
		c.Password,
		c.Insecure,
		client.GetKeystoneScope(c.DomainID, c.DomainName,
			c.ProjectID, c.ProjectName),
	)
	restClient.Init()

	return restClient
}
