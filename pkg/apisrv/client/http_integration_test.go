package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestCreateRefMethod(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: false,
	})
	defer s.CloseT(t)
	hc := integration.NewTestingHTTPClient(t, s.URL(), integration.BobUserID)

	testID := "test"
	projectUUID := testID + "_project"
	vnUUID := testID + "_virtual-network"
	niUUID := testID + "_network-ipam"

	integration.CreateProject(
		t,
		hc,
		&models.Project{
			UUID:       projectUUID,
			ParentType: integration.DomainType,
			ParentUUID: integration.DefaultDomainUUID,
			Name:       "testProject",
			Quota:      &models.QuotaType{},
		},
	)
	defer integration.DeleteProject(t, hc, projectUUID)

	integration.CreateNetworkIpam(
		t,
		hc,
		&models.NetworkIpam{
			UUID:       niUUID,
			ParentType: integration.ProjectType,
			ParentUUID: projectUUID,
			Name:       "testIpam",
		},
	)
	defer integration.DeleteNetworkIpam(t, hc, niUUID)

	integration.CreateVirtualNetwork(
		t,
		hc,
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

	defer integration.DeleteVirtualNetwork(t, hc, vnUUID)

	// After creating VirtualNetwork it is already connected to networkIpam
	vn := integration.GetVirtualNetwork(t, hc, vnUUID)

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
	assert.NoError(t, err)

	vn = integration.GetVirtualNetwork(t, hc, vnUUID)

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

	vn = integration.GetVirtualNetwork(t, hc, vnUUID)
	assert.Len(t, vn.NetworkIpamRefs, 1)
}

func TestRemoteIntPoolMethods(t *testing.T) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: false,
	})
	defer s.CloseT(t)
	hc := integration.NewTestingHTTPClient(t, s.URL(), integration.BobUserID)
	_, err := hc.Login(context.Background())
	require.NoError(t, err)

	err = hc.CreateIntPool(context.Background(), "test_int_pool_806f099f3", 8000100, 8000200)
	require.NoError(t, err)
	defer func() {
		err = hc.DeleteIntPool(context.Background(), "test_int_pool_806f099f3")
		assert.NoError(t, err)
	}()

	val, err := hc.AllocateInt(context.Background(), "test_int_pool_806f099f3", "test_owner_806f099f3")
	defer func() {
		err = hc.DeallocateInt(context.Background(), "test_int_pool_806f099f3", val)
		assert.NoError(t, err)
	}()
	assert.NoError(t, err)
	assert.True(t, val > 8000099)

	owner, err := hc.GetIntOwner(context.Background(), "test_int_pool_806f099f3", val)
	assert.NoError(t, err)
	assert.Equal(t, "test_owner_806f099f3", owner)

	err = hc.SetInt(context.Background(), "test_int_pool_806f099f3", val+1, db.EmptyIntOwner)
	defer func() {
		err = hc.DeallocateInt(context.Background(), "test_int_pool_806f099f3", val+1)
		assert.NoError(t, err)
	}()

	assert.NoError(t, err)
}
