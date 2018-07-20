package compilation

import (
	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/compilation/config"
)

func dialAPIServer(config config.Config) (*client.HTTP, error) {
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
				Name: c.Project,
			},
		},
	)
	restClient.Init()

	err := restClient.Login()
	if err != nil {
		return nil, err
	}
	// TODO conn.Close()

	// TODO Make restClient fulfill services.WriteService
	return restClient, nil
}
