package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// GRPCCreateProject creates a project resource in given service.
func GRPCCreateProject(t *testing.T, s services.ContrailServiceClient, ctx context.Context, obj *models.Project) *models.Project {
	resp, err := s.CreateProject(ctx, &services.CreateProjectRequest{Project: obj})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating Project failed\n requested: %+v\n "+
			"response: %+v\n", obj, resp),
	)
	return resp.GetProject()
}

// GRPCDeleteProject deletes a project resource using given service.
func GRPCDeleteProject(t *testing.T, s services.ContrailServiceClient, ctx context.Context, id string) {
	resp, err := s.DeleteProject(ctx, &services.DeleteProjectRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting Project failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

// GRPCGetProject gets a project resource from given service.
func GRPCGetProject(t *testing.T, s services.ReadService, ctx context.Context, id string) *models.Project {
	resp, err := s.GetProject(ctx, &services.GetProjectRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("getting Project failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetProject()
}

// GRPCCreateNetworkIpam creates a network IPAM resource in given service.
func GRPCCreateNetworkIpam(t *testing.T, s services.ContrailServiceClient, ctx context.Context, obj *models.NetworkIpam) *models.NetworkIpam {
	resp, err := s.CreateNetworkIpam(ctx, &services.CreateNetworkIpamRequest{NetworkIpam: obj})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating NetworkIpam failed\n requested: %+v\n "+
			"response: %+v\n", obj, resp),
	)
	return resp.GetNetworkIpam()
}

// GRPCDeleteNetworkIpam deletes a network IPAM resource from given service.
func GRPCDeleteNetworkIpam(t *testing.T, s services.ContrailServiceClient, ctx context.Context, id string) {
	resp, err := s.DeleteNetworkIpam(ctx, &services.DeleteNetworkIpamRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting NetworkIpam failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

// GRPCGetNetworkIpam gets a network IPAM resource from given service.
func GRPCGetNetworkIpam(t *testing.T, s services.ReadService, ctx context.Context, id string) *models.NetworkIpam {
	resp, err := s.GetNetworkIpam(ctx, &services.GetNetworkIpamRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("getting NetworkIpam failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetNetworkIpam()
}

// GRPCCreateNetworkPolicy creates a network policy resource in given service.
func GRPCCreateNetworkPolicy(t *testing.T, s services.ContrailServiceClient, ctx context.Context, obj *models.NetworkPolicy) *models.NetworkPolicy {
	resp, err := s.CreateNetworkPolicy(ctx, &services.CreateNetworkPolicyRequest{NetworkPolicy: obj})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating NetworkPolicy failed\n requested: %+v\n "+
			"response: %+v\n", obj, resp),
	)
	return resp.GetNetworkPolicy()
}

// GRPCDeleteNetworkPolicy deletes a network policy resource from given service.
func GRPCDeleteNetworkPolicy(t *testing.T, s services.ContrailServiceClient, ctx context.Context, id string) {
	resp, err := s.DeleteNetworkPolicy(ctx, &services.DeleteNetworkPolicyRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting NetworkPolicy failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

// GRPCGetNetworkPolicy gets a network policy resource from given service.
func GRPCGetNetworkPolicy(t *testing.T, s services.ReadService, ctx context.Context, id string) *models.NetworkPolicy {
	resp, err := s.GetNetworkPolicy(ctx, &services.GetNetworkPolicyRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("getting NetworkPolicy failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetNetworkPolicy()
}

// GRPCCreateVirtualNetwork creates a virtual network resource from given service.
func GRPCCreateVirtualNetwork(t *testing.T, s services.ContrailServiceClient, ctx context.Context, obj *models.VirtualNetwork) *models.VirtualNetwork {
	resp, err := s.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{VirtualNetwork: obj})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating VirtualNetwork failed\n requested: %+v\n "+
			"response: %+v\n", obj, resp),
	)
	return resp.GetVirtualNetwork()
}

// GRPCDeleteVirtualNetwork deletes a virtual network resource from given service.
func GRPCDeleteVirtualNetwork(t *testing.T, s services.ContrailServiceClient, ctx context.Context, id string) {
	resp, err := s.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting VirtualNetwork failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

// GRPCGetVirtualNetwork gets a virtual network resource from given service.
func GRPCGetVirtualNetwork(t *testing.T, s services.ReadService, ctx context.Context, id string) *models.VirtualNetwork {
	resp, err := s.GetVirtualNetwork(ctx, &services.GetVirtualNetworkRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("getting VirtualNetwork failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetVirtualNetwork()
}

// GRPCDeleteSecurityGroup deletes a security group resource from given service.
func GRPCDeleteSecurityGroup(t *testing.T, s services.ContrailServiceClient, ctx context.Context, id string) {
	resp, err := s.DeleteSecurityGroup(ctx, &services.DeleteSecurityGroupRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting SecurityGroup failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

// GRPCDeleteAccessControlList deletes an access control list resource from given service.
func GRPCDeleteAccessControlList(t *testing.T, s services.ContrailServiceClient, ctx context.Context, id string) {
	resp, err := s.DeleteAccessControlList(ctx, &services.DeleteAccessControlListRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting AccessControlList failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}
