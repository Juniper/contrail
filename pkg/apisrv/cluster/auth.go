package cluster

import (
	"context"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	openstack    = "openstack"
	basicAuth    = "basic-auth"
	keystoneAuth = "keystone"
)

// GetAuthType from the cluster configuration
func GetAuthType(clusterID string) (authType string, err error) {
	h := client.NewHTTP(&client.HTTPConfig{
		ID:       viper.GetString("client.id"),
		Password: viper.GetString("client.password"),
		Endpoint: viper.GetString("client.endpoint"),
		AuthURL:  viper.GetString("keystone.authurl"),
		Scope: keystone.NewScope(
			viper.GetString("client.domain_id"),
			viper.GetString("client.domain_name"),
			viper.GetString("client.project_id"),
			viper.GetString("client.project_name"),
		),
		Insecure: viper.GetBool("insecure"),
	})

	request := &services.GetContrailClusterRequest{
		ID: clusterID,
	}
	var resp *services.GetContrailClusterResponse
	resp, err = h.GetContrailCluster(context.Background(), request)
	if err != nil {
		return "", err
	}

	if resp.GetContrailCluster().GetOrchestrator() != openstack {
		authType = basicAuth
	} else {
		authType = keystoneAuth
	}
	return authType, nil
}
