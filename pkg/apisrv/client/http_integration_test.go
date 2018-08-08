package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestCreateRefMethod(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           db.DriverPostgreSQL,
		EnableEtcdNotifier: false,
		RepoRootPath:       "../../..",
	})
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())

	testID := "test"
	projectUUID := testID + "_project"
	vnUUID := testID + "_virtual-network"
	niUUID := testID + "_network-ipam"

	hc.CreateProject(
		t,
		&models.Project{
			UUID:       projectUUID,
			ParentType: integration.DomainType,
			ParentUUID: integration.DefaultDomainUUID,
			Name:       "testProject",
			Quota:      &models.QuotaType{},
		},
	)
	defer hc.DeleteProject(t, projectUUID)

	hc.CreateNetworkIPAM(
		t,
		&models.NetworkIpam{
			UUID:       niUUID,
			ParentType: integration.ProjectType,
			ParentUUID: projectUUID,
			Name:       "testIpam",
		},
	)
	defer hc.DeleteNetworkIPAM(t, niUUID)

	hc.CreateVirtualNetwork(
		t,
		&models.VirtualNetwork{
			UUID:       vnUUID,
			ParentType: integration.ProjectType,
			ParentUUID: projectUUID,
			Name:       "testVN",
			NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
				{
					UUID: niUUID,
				},
			},
		},
	)
	defer hc.DeleteVirtualNetwork(t, vnUUID)

	// After creating VirtualNetwork it is already connected to networkIpam
	vn := hc.GetVirtualNetwork(t, vnUUID)
	assert.Len(t, vn.NetworkIpamRefs, 1)

	_, err := hc.DeleteVirtualNetworkNetworkIpamRef(
		context.Background(),
		&services.DeleteVirtualNetworkNetworkIpamRefRequest{
			ID: vnUUID,
			VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
				UUID: niUUID,
			},
		},
	)
	assert.Equal(t, nil, err)

	vn = hc.GetVirtualNetwork(t, vnUUID)
	assert.Len(t, vn.NetworkIpamRefs, 0)

	_, err = hc.CreateVirtualNetworkNetworkIpamRef(
		context.Background(),
		&services.CreateVirtualNetworkNetworkIpamRefRequest{
			ID: vnUUID,
			VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
				UUID: niUUID,
			},
		},
	)
	assert.Equal(t, nil, err)

	vn = hc.GetVirtualNetwork(t, vnUUID)
	assert.Len(t, vn.NetworkIpamRefs, 1)
}
