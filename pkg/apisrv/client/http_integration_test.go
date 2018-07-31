package client_test

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
)

func TestCreateRefMethod(t *testing.T) {
	s := integration.NewRunningAPIServer(t, "../../..", db.DriverPostgreSQL)
	defer s.Close(t)
	hc := integration.NewHTTPAPIClient(t, s.URL())

	testID := "test"
	projectUUID := testID + "_project"
	vnUUID := testID + "_virtual-network"
	niUUID := testID + "_network-ipam"
	rtUUID := testID + "_route-target"

	hc.CreateProject(t, project(projectUUID))
	defer hc.DeleteProject(t, projectUUID)

	hc.CreateNetworkIPAM(t, networkIPAM(niUUID, projectUUID))
	defer hc.DeleteNetworkIPAM(t, niUUID)

	hc.CreateVirtualNetwork(t, virtualNetwork(vnUUID, projectUUID, niUUID))
	defer hc.DeleteVirtualNetwork(t, vnUUID)

	hc.CreateRouteTarget(t, routeTarget(rtUUID))
	defer hc.DeleteRouteTarget(t, rtUUID)

	vn := hc.GetVirtualNetwork(t, vnUUID)

	riUUID := vn.GetRoutingInstances()[0].GetUUID()

	// TODO: After creating Routing Instance resource
	// it should be connected to its default route target.
	// After it will be done in intent compiler this
	// won't have 0 RouteTargetRefs after creating anymore.
	ri := hc.GetRoutingInstance(t, riUUID)
	assert.Equal(t, 0, len(ri.RouteTargetRefs))

	_, err := hc.CreateRoutingInstanceRouteTargetRef(context.Background(), createRoutingInstanceRouteTargetRefRequest(riUUID, rtUUID))
	assert.Equal(t, nil, err)

	ri = hc.GetRoutingInstance(t, riUUID)
	assert.Equal(t, 1, len(ri.RouteTargetRefs))

	_, err = hc.DeleteRoutingInstanceRouteTargetRef(context.Background(), deleteRoutingInstanceRouteTargetRefRequest(riUUID, rtUUID))
	assert.Equal(t, nil, err)

	ri = hc.GetRoutingInstance(t, riUUID)
	assert.Equal(t, 0, len(ri.RouteTargetRefs))
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

func routeTarget(uuid string) *models.RouteTarget {
	return &models.RouteTarget{
		UUID:        uuid,
		FQName:      []string{"target:100:100"},
		DisplayName: "target:100:100",
	}
}

func createRoutingInstanceRouteTargetRefRequest(riUUID, rtUUID string) *services.CreateRoutingInstanceRouteTargetRefRequest {

	return &services.CreateRoutingInstanceRouteTargetRefRequest{
		ID: riUUID,
		RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
			UUID: rtUUID,
			Attr: &models.InstanceTargetType{
				ImportExport: "",
			},
		},
	}
}

func deleteRoutingInstanceRouteTargetRefRequest(riUUID, rtUUID string) *services.DeleteRoutingInstanceRouteTargetRefRequest {

	return &services.DeleteRoutingInstanceRouteTargetRefRequest{
		ID: riUUID,
		RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
			UUID: rtUUID,
		},
	}
}
