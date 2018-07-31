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
	s := integration.NewRunningAPIServer(t, "../../..", db.DriverPostgreSQL)
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())

	testID := "test"
	projectUUID := testID + "_project"
	vnUUID := testID + "_virtual-network"
	niUUID := testID + "_network-ipam"

	hc.CreateProject(t, project(projectUUID))
	defer hc.DeleteProject(t, projectUUID)

	hc.CreateNetworkIPAM(t, networkIPAM(niUUID, projectUUID))
	defer hc.DeleteNetworkIPAM(t, niUUID)

	hc.CreateVirtualNetwork(t, virtualNetwork(vnUUID, projectUUID, niUUID))
	defer hc.DeleteVirtualNetwork(t, vnUUID)

	// After creating VirtualNetwork it is already connected to networkIpam
	vn := hc.GetVirtualNetwork(t, vnUUID)
	assert.Equal(t, 1, len(vn.NetworkIpamRefs))

	_, err := hc.DeleteVirtualNetworkNetworkIpamRef(
		context.Background(), &services.DeleteVirtualNetworkNetworkIpamRefRequest{
			ID: vnUUID,
			VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
				UUID: niUUID,
			},
		})
	assert.Equal(t, nil, err)

	vn = hc.GetVirtualNetwork(t, vnUUID)
	assert.Equal(t, 0, len(vn.NetworkIpamRefs))

	_, err = hc.CreateVirtualNetworkNetworkIpamRef(
		context.Background(), &services.CreateVirtualNetworkNetworkIpamRefRequest{
			ID: vnUUID,
			VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
				UUID: niUUID,
			},
		})
	assert.Equal(t, nil, err)

	vn = hc.GetVirtualNetwork(t, vnUUID)
	assert.Equal(t, 1, len(vn.NetworkIpamRefs))
}

func project(uuid string) *models.Project {
	return &models.Project{
		UUID:       uuid,
		ParentType: integration.DomainType,
		ParentUUID: integration.DefaultDomainUUID,
		FQName:     []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
		Quota:      &models.QuotaType{},
	}
}

func networkIPAM(uuid, parentUUID string) *models.NetworkIpam {
	return &models.NetworkIpam{
		UUID:       uuid,
		ParentType: integration.ProjectType,
		ParentUUID: parentUUID,
		FQName:     []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
	}
}

func virtualNetwork(uuid, parentUUID, networkIpamUUID string) *models.VirtualNetwork {
	return &models.VirtualNetwork{
		UUID:        uuid,
		ParentType:  integration.ProjectType,
		ParentUUID:  parentUUID,
		FQName:      []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"},
		Perms2:      &models.PermType2{Owner: integration.AdminUserID},
		DisplayName: "vn",
		NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{{
			UUID: networkIpamUUID,
			To:   []string{integration.DefaultDomainID, integration.AdminProjectID, networkIpamUUID + "-fq-name"},
		}},
	}
}
