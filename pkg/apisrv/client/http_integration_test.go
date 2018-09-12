package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestCreateRefMethod(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		EnableEtcdNotifier: false,
		RepoRootPath:       "../../..",
	})
	defer s.CloseT(t)
	hc := integration.NewTestingHTTPClient(t, s.URL())

	testID := "test"
	projectUUID := testID + "_project"
	vnUUID := testID + "_virtual-network"
	niUUID := testID + "_network-ipam"

	ctx := context.Background()

	hc.CreateProject(
		ctx,
		&services.CreateProjectRequest{
			Project: &models.Project{
				UUID:       projectUUID,
				ParentType: integration.DomainType,
				ParentUUID: integration.DefaultDomainUUID,
				Name:       "testProject",
				Quota:      &models.QuotaType{},
			},
		},
	)
	defer hc.DeleteProject(ctx, &services.DeleteProjectRequest{ID: projectUUID})

	hc.CreateNetworkIpam(
		ctx,
		&services.CreateNetworkIpamRequest{
			NetworkIpam: &models.NetworkIpam{
				UUID:       niUUID,
				ParentType: integration.ProjectType,
				ParentUUID: projectUUID,
				Name:       "testIpam",
			},
		},
	)
	defer hc.DeleteNetworkIpam(ctx, &services.DeleteNetworkIpamRequest{ID: niUUID})

	hc.CreateVirtualNetwork(
		ctx,
		&services.CreateVirtualNetworkRequest{
			VirtualNetwork: &models.VirtualNetwork{
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
		},
	)
	defer hc.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{ID: vnUUID})

	// After creating VirtualNetwork it is already connected to networkIpam
	vnResp, err := hc.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: vnUUID})
	assert.NoError(t, err)

	vn := vnResp.GetVirtualNetwork()
	assert.Len(t, vn.NetworkIpamRefs, 1)

	_, err = hc.DeleteVirtualNetworkNetworkIpamRef(
		context.Background(),
		&services.DeleteVirtualNetworkNetworkIpamRefRequest{
			ID: vnUUID,
			VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
				UUID: niUUID,
			},
		},
	)
	assert.NoError(t, err)

	vnResp, err = hc.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: vnUUID})
	assert.NoError(t, err)

	vn = vnResp.GetVirtualNetwork()
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
	assert.NoError(t, err)

	vnResp, err = hc.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: vnUUID})
	assert.NoError(t, err)
	vn = vnResp.GetVirtualNetwork()
	assert.Len(t, vn.NetworkIpamRefs, 1)
}

func TestRemoteIntPoolMethods(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		EnableEtcdNotifier: false,
		RepoRootPath:       "../../..",
	})
	defer s.CloseT(t)
	hc := integration.NewTestingHTTPClient(t, s.URL())
	err := hc.Login(context.Background())
	require.NoError(t, err)

	rt, err := hc.AllocateInt(context.Background(), "route_target_number")
	defer func() {
		err = hc.DeallocateInt(context.Background(), "route_target_number", rt)
		assert.NoError(t, err)
	}()

	assert.NoError(t, err)
	assert.True(t, rt > 8000001)

	err = hc.SetInt(context.Background(), "route_target_number", rt+1)
	defer func() {
		err = hc.DeallocateInt(context.Background(), "route_target_number", rt+1)
		assert.NoError(t, err)
	}()

	assert.NoError(t, err)
}
