package keystone

import (
	"context"

	"github.com/Juniper/contrail/pkg/services"
)

const (
	openstack    = "openstack"
	keystoneAuth = "keystone"
)

// GetAuthType from the cluster configuration
func (keystone *Keystone) GetAuthType(clusterID string) (authType string, err error) {
	resp, err := keystone.APIClient.GetContrailCluster(
		context.Background(),
		&services.GetContrailClusterRequest{
			ID: clusterID,
		})
	if err != nil {
		return "", err
	}

	authType = keystoneAuth
	if resp.GetContrailCluster().GetOrchestrator() != openstack {
		authType = basicAuth
	}
	return authType, nil
}
