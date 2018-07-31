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
	assert.Len(t, vn.NetworkIpamRefs, 1)

	_, err := hc.DeleteVirtualNetworkNetworkIpamRef(
		context.Background(), &services.DeleteVirtualNetworkNetworkIpamRefRequest{
			ID: vnUUID,
			VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
				UUID: niUUID,
			},
		})
	assert.Equal(t, nil, err)

	vn = hc.GetVirtualNetwork(t, vnUUID)
	assert.Len(t, vn.NetworkIpamRefs, 0)

	_, err = hc.CreateVirtualNetworkNetworkIpamRef(
		context.Background(), &services.CreateVirtualNetworkNetworkIpamRefRequest{
			ID: vnUUID,
			VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
				UUID: niUUID,
			},
		})
	assert.Equal(t, nil, err)

	vn = hc.GetVirtualNetwork(t, vnUUID)
	assert.Len(t, vn.NetworkIpamRefs, 1)
}

func project(uuid string) *models.Project {
	p := models.MakeProject()
	p.UUID = uuid
	p.ParentUUID = integration.DefaultDomainUUID
	p.FQName = []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"}
	return p
}

func networkIPAM(uuid, parentUUID string) *models.NetworkIpam {
	ni := models.MakeNetworkIpam()
	ni.UUID = uuid
	ni.ParentType = integration.ProjectType
	ni.ParentUUID = parentUUID
	ni.FQName = []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"}
	// Ignoring this value to pass network ipam validation
	ni.IpamSubnets = nil
	return ni
}

func virtualNetwork(uuid, parentUUID, networkIpamUUID string) *models.VirtualNetwork {
	vn := models.MakeVirtualNetwork()
	vn.UUID = uuid
	vn.ParentType = integration.ProjectType
	vn.ParentUUID = parentUUID
	vn.FQName = []string{integration.DefaultDomainID, integration.AdminProjectID, uuid + "-fq-name"}
	vn.DisplayName = "vn"
	vn.NetworkIpamRefs = append(
		vn.NetworkIpamRefs,
		&models.VirtualNetworkNetworkIpamRef{
			UUID: networkIpamUUID,
			To:   []string{integration.DefaultDomainID, integration.AdminProjectID, networkIpamUUID + "-fq-name"},
		},
	)
	// Ignoring these values to pass virtual network validation
	vn.ProviderProperties = nil
	vn.VirtualNetworkProperties = nil
	vn.MacMoveControl = nil
	vn.MacLimitControl = nil
	return vn
}
