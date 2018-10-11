package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateProject creates a project resource in given service.
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

// DeleteProject deletes a project resource using given service.
func DeleteProject(t *testing.T, s services.WriteService, id string) {
	resp, err := s.DeleteProject(context.Background(), &services.DeleteProjectRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting Project failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

// GetProject gets a project resource from given service.
func GetProject(t *testing.T, s services.ReadService, id string) *models.Project {
	resp, err := s.GetProject(context.Background(), &services.GetProjectRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("getting Project failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetProject()
}

// CreateNetworkIpam creates a network IPAM resource in given service.
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

// DeleteNetworkIpam deletes a network IPAM resource from given service.
func DeleteNetworkIpam(t *testing.T, s services.WriteService, id string) {
	resp, err := s.DeleteNetworkIpam(context.Background(), &services.DeleteNetworkIpamRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting NetworkIpam failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

// GetNetworkIpam gets a network IPAM resource from given service.
func GetNetworkIpam(t *testing.T, s services.ReadService, id string) *models.NetworkIpam {
	resp, err := s.GetNetworkIpam(context.Background(), &services.GetNetworkIpamRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("getting NetworkIpam failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetNetworkIpam()
}

// CreateVirtualNetwork creates a virtual network resource from given service.
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

// DeleteVirtualNetwork deletes a virtual network resource from given service.
func DeleteVirtualNetwork(t *testing.T, s services.WriteService, id string) {
	resp, err := s.DeleteVirtualNetwork(context.Background(), &services.DeleteVirtualNetworkRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting VirtualNetwork failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

// GetVirtualNetwork gets a virtual network resource from given service.
func GetVirtualNetwork(t *testing.T, s services.ReadService, id string) *models.VirtualNetwork {
	resp, err := s.GetVirtualNetwork(context.Background(), &services.GetVirtualNetworkRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("getting VirtualNetwork failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetVirtualNetwork()
}

// DeleteSecurityGroup deletes a security group resource from given service.
func DeleteSecurityGroup(t *testing.T, s services.WriteService, id string) {
	resp, err := s.DeleteSecurityGroup(context.Background(), &services.DeleteSecurityGroupRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting SecurityGroup failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

// DeleteAccessControlList deletes an access control list resource from given service.
func DeleteAccessControlList(t *testing.T, s services.WriteService, id string) {
	resp, err := s.DeleteAccessControlList(context.Background(), &services.DeleteAccessControlListRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting AccessControlList failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}
