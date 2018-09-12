package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/stretchr/testify/require"
)

func CreateProject(t *testing.T, s services.WriteService, obj *models.Project) *models.Project {
	resp, err := s.CreateProject(context.Background(), &services.CreateProjectRequest{Project: obj})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating Project failed\n requested: %+v\n "+
			"response: %+v\n", obj, resp),
	)
	return resp.GetProject()
}

func DeleteProject(t *testing.T, s services.WriteService, id string) {
	resp, err := s.DeleteProject(context.Background(), &services.DeleteProjectRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting Project failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

func GetProject(t *testing.T, s services.ReadService, id string) *models.Project {
	resp, err := s.GetProject(context.Background(), &services.GetProjectRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating Project failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetProject()
}

func CreateNetworkIpam(t *testing.T, s services.WriteService, obj *models.NetworkIpam) *models.NetworkIpam {
	resp, err := s.CreateNetworkIpam(context.Background(), &services.CreateNetworkIpamRequest{NetworkIpam: obj})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating NetworkIpam failed\n requested: %+v\n "+
			"response: %+v\n", obj, resp),
	)
	return resp.GetNetworkIpam()
}

func DeleteNetworkIpam(t *testing.T, s services.WriteService, id string) {
	resp, err := s.DeleteNetworkIpam(context.Background(), &services.DeleteNetworkIpamRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting NetworkIpam failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

func GetNetworkIpam(t *testing.T, s services.ReadService, id string) *models.NetworkIpam {
	resp, err := s.GetNetworkIpam(context.Background(), &services.GetNetworkIpamRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating NetworkIpam failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetNetworkIpam()
}

func CreateVirtualNetwork(t *testing.T, s services.WriteService, obj *models.VirtualNetwork) *models.VirtualNetwork {
	resp, err := s.CreateVirtualNetwork(context.Background(), &services.CreateVirtualNetworkRequest{VirtualNetwork: obj})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating VirtualNetwork failed\n requested: %+v\n "+
			"response: %+v\n", obj, resp),
	)
	return resp.GetVirtualNetwork()
}

func DeleteVirtualNetwork(t *testing.T, s services.WriteService, id string) {
	resp, err := s.DeleteVirtualNetwork(context.Background(), &services.DeleteVirtualNetworkRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting VirtualNetwork failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

func GetVirtualNetwork(t *testing.T, s services.ReadService, id string) *models.VirtualNetwork {
	resp, err := s.GetVirtualNetwork(context.Background(), &services.GetVirtualNetworkRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating VirtualNetwork failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetVirtualNetwork()
}

//hc.CreateNetworkIpam(ctx, networkIPAM(networkIPAMUUID, projectUUID))
//defer hc.DeleteNetworkIpam(ctx, networkIPAMUUID)
//defer ec.DeleteNetworkIpam(ctx, networkIPAMUUID)

//hc.CreateVirtualNetwork(ctx, virtualNetworkRed(vnRedUUID, projectUUID, networkIPAMUUID))
//hc.CreateVirtualNetwork(ctx, virtualNetworkGreen(vnGreenUUID, projectUUID, networkIPAMUUID))
//hc.CreateVirtualNetwork(ctx, virtualNetworkBlue(vnBlueUUID, projectUUID, networkIPAMUUID))
//defer deleteVirtualNetworksFromAPIServer(t, hc, vnUUIDs)
//defer ec.DeleteKey(t, integration.JSONEtcdKey(integrationetcd.VirtualNetworkSchemaID, ""),
//clientv3.WithPrefix()) // delete all VNs

//vnRed := hc.GetVirtualNetwork(ctx, vnRedUUID)
//vnGreen := hc.GetVirtualNetwork(ctx, vnGreenUUID)
//vnBlue := hc.GetVirtualNetwork(ctx, vnBlueUUID)

//closeSync := integration.RunSyncService(t)
//defer closeSync()

//redEvent := integration.RetrieveCreateEvent(redCtx, t, vnRedWatch)
//greenEvent := integration.RetrieveCreateEvent(greenCtx, t, vnGreenWatch)
//blueEvent := integration.RetrieveCreateEvent(blueCtx, t, vnBlueWatch)

//checkSyncedVirtualNetwork(t, redEvent, vnRed)
//checkSyncedVirtualNetwork(t, greenEvent, vnGreen)
//checkSyncedVirtualNetwork(t, blueEvent, vnBlue)
