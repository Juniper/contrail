package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// Runner can be run and return an error.
type Runner interface {
	Run() error
}

// RunConcurrently runs runner in separate goroutine and returns a channel to read Run error.
func RunConcurrently(r Runner) <-chan error {
	runError := make(chan error)
	go func() {
		runError <- r.Run()
	}()

	return runError
}

// Closer can be closed.
type Closer interface {
	Close()
}

// CloseNoError calls close and expects that error channel is closed without an error.
func CloseNoError(t *testing.T, c Closer, errChan <-chan error) {
	c.Close()
	assert.NoError(t, <-errChan, "unexpected error while closing")
}

// CloseFatalIfError calls close and calls log.Fatal if error channel returns an error.
func CloseFatalIfError(c Closer, errChan <-chan error) {
	c.Close()
	if err := <-errChan; err != nil {
		logutil.FatalWithStackTrace(errors.Wrap(err, "unexpected error while closing"))
	}
}

// RunCloser is a Runner that is also a Closer.
type RunCloser interface {
	Runner
	Closer
}

// RunNoError runs RunCloser concurrently and returns callback for stopping
// the goroutine that also expects no error is returned from Run.
func RunNoError(t *testing.T, rc RunCloser) (close func(*testing.T)) {
	errChan := RunConcurrently(rc)
	return func(*testing.T) { CloseNoError(t, rc, errChan) }
}

///////////////////////////////////
// HTTP API CRUD request helpers //
///////////////////////////////////

// DeleteAccessControlList deletes an access control list resource from given service.
func DeleteAccessControlList(t *testing.T, s services.WriteService, uuid string) {
	resp, err := s.DeleteAccessControlList(context.Background(), &services.DeleteAccessControlListRequest{ID: uuid})
	require.NoError(t, err, deleteFailureMessage(uuid, resp))
}

// CreateContrailCluster creates a Contrail cluster resource in given service.
func CreateContrailCluster(
	t *testing.T, s services.WriteService, obj *models.ContrailCluster,
) *models.ContrailCluster {
	resp, err := s.CreateContrailCluster(
		context.Background(),
		&services.CreateContrailClusterRequest{ContrailCluster: obj},
	)
	require.NoError(t, err, createFailureMessage(obj, resp))
	return resp.GetContrailCluster()
}

// DeleteContrailCluster deletes a Contrail cluster resource from given service.
func DeleteContrailCluster(t *testing.T, s services.WriteService, uuid string) {
	resp, err := s.DeleteContrailCluster(context.Background(), &services.DeleteContrailClusterRequest{ID: uuid})
	require.NoError(t, err, deleteFailureMessage(uuid, resp))
}

// CreateEndpoint creates an endpoint resource in given service.
func CreateEndpoint(t *testing.T, s services.WriteService, obj *models.Endpoint) *models.Endpoint {
	resp, err := s.CreateEndpoint(context.Background(), &services.CreateEndpointRequest{Endpoint: obj})
	require.NoError(t, err, createFailureMessage(obj, resp))
	return resp.GetEndpoint()
}

// DeleteEndpoint deletes a network IPAM resource from given service.
func DeleteEndpoint(t *testing.T, s services.WriteService, uuid string) {
	resp, err := s.DeleteEndpoint(context.Background(), &services.DeleteEndpointRequest{ID: uuid})
	require.NoError(t, err, deleteFailureMessage(uuid, resp))
}

// CreateNetworkIpam creates a network IPAM resource in given service.
func CreateNetworkIpam(t *testing.T, s services.WriteService, obj *models.NetworkIpam) *models.NetworkIpam {
	resp, err := s.CreateNetworkIpam(context.Background(), &services.CreateNetworkIpamRequest{NetworkIpam: obj})
	require.NoError(t, err, createFailureMessage(obj, resp))
	return resp.GetNetworkIpam()
}

// GetNetworkIpam gets a network IPAM resource from given service.
func GetNetworkIpam(t *testing.T, s services.ReadService, uuid string) *models.NetworkIpam {
	resp, err := s.GetNetworkIpam(context.Background(), &services.GetNetworkIpamRequest{ID: uuid})
	require.NoError(t, err, getFailureMessage(uuid, resp))
	return resp.GetNetworkIpam()
}

// DeleteNetworkIpam deletes a network IPAM resource from given service.
func DeleteNetworkIpam(t *testing.T, s services.WriteService, uuid string) {
	resp, err := s.DeleteNetworkIpam(context.Background(), &services.DeleteNetworkIpamRequest{ID: uuid})
	require.NoError(t, err, deleteFailureMessage(uuid, resp))
}

// CreateNetworkPolicy creates a network policy resource in given service.
func CreateNetworkPolicy(
	ctx context.Context, t *testing.T, s services.WriteService, obj *models.NetworkPolicy,
) *models.NetworkPolicy {
	resp, err := s.CreateNetworkPolicy(ctx, &services.CreateNetworkPolicyRequest{NetworkPolicy: obj})
	require.NoError(t, err, createFailureMessage(obj, resp))
	return resp.GetNetworkPolicy()
}

// GetNetworkPolicy gets a network policy resource from given service.
func GetNetworkPolicy(ctx context.Context, t *testing.T, s services.ReadService, uuid string) *models.NetworkPolicy {
	resp, err := s.GetNetworkPolicy(ctx, &services.GetNetworkPolicyRequest{ID: uuid})
	require.NoError(t, err, getFailureMessage(uuid, resp))
	return resp.GetNetworkPolicy()
}

// DeleteNetworkPolicy deletes a network policy resource from given service.
func DeleteNetworkPolicy(
	ctx context.Context, t *testing.T, s services.WriteService, uuid string,
) {
	resp, err := s.DeleteNetworkPolicy(ctx, &services.DeleteNetworkPolicyRequest{ID: uuid})
	require.NoError(t, err, deleteFailureMessage(uuid, resp))
}

// CreateProject creates a project resource in given service.
func CreateProject(t *testing.T, s services.WriteService, obj *models.Project) *models.Project {
	resp, err := s.CreateProject(context.Background(), &services.CreateProjectRequest{Project: obj})
	require.NoError(t, err, createFailureMessage(obj, resp))
	return resp.GetProject()
}

// GetProject gets a project resource from given service.
func GetProject(t *testing.T, s services.ReadService, uuid string) *models.Project {
	resp, err := s.GetProject(context.Background(), &services.GetProjectRequest{ID: uuid})
	require.NoError(t, err, getFailureMessage(uuid, resp))
	return resp.GetProject()
}

// DeleteProject deletes a project resource using given service.
func DeleteProject(t *testing.T, s services.WriteService, uuid string) {
	resp, err := s.DeleteProject(context.Background(), &services.DeleteProjectRequest{ID: uuid})
	require.NoError(t, err, deleteFailureMessage(uuid, resp))
}

// CreateRoutingInstance creates a routing instance resource from given service.
func CreateRoutingInstance(t *testing.T, s services.WriteService, obj *models.RoutingInstance) *models.RoutingInstance {
	resp, err := s.CreateRoutingInstance(context.Background(), &services.CreateRoutingInstanceRequest{
		RoutingInstance: obj,
	})
	require.NoError(t, err, createFailureMessage(obj, resp))
	return resp.GetRoutingInstance()
}

// GetRoutingInstance gets a routing instance resource from given service.
func GetRoutingInstance(t *testing.T, s services.ReadService, uuid string) *models.RoutingInstance {
	resp, err := s.GetRoutingInstance(context.Background(), &services.GetRoutingInstanceRequest{ID: uuid})
	require.NoError(t, err, getFailureMessage(uuid, resp))
	return resp.GetRoutingInstance()
}

// DeleteRoutingInstance deletes a routing instance resource from given service.
func DeleteRoutingInstance(t *testing.T, s services.WriteService, uuid string) {
	resp, err := s.DeleteRoutingInstance(context.Background(), &services.DeleteRoutingInstanceRequest{ID: uuid})
	require.NoError(t, err, deleteFailureMessage(uuid, resp))
}

// CreateRouteTarget creates a route target resource from given service.
func CreateRouteTarget(t *testing.T, s services.WriteService, obj *models.RouteTarget) *models.RouteTarget {
	resp, err := s.CreateRouteTarget(context.Background(), &services.CreateRouteTargetRequest{RouteTarget: obj})
	require.NoError(t, err, createFailureMessage(obj, resp))
	return resp.GetRouteTarget()
}

// GetRouteTarget gets a route target resource from given service.
func GetRouteTarget(t *testing.T, s services.ReadService, uuid string) *models.RouteTarget {
	resp, err := s.GetRouteTarget(context.Background(), &services.GetRouteTargetRequest{ID: uuid})
	require.NoError(t, err, getFailureMessage(uuid, resp))
	return resp.GetRouteTarget()
}

// DeleteRouteTarget deletes a route target resource from given service.
func DeleteRouteTarget(t *testing.T, s services.WriteService, uuid string) {
	resp, err := s.DeleteRouteTarget(context.Background(), &services.DeleteRouteTargetRequest{ID: uuid})
	require.NoError(t, err, deleteFailureMessage(uuid, resp))
}

// DeleteSecurityGroup deletes a security group resource from given service.
func DeleteSecurityGroup(t *testing.T, s services.WriteService, uuid string) {
	resp, err := s.DeleteSecurityGroup(context.Background(), &services.DeleteSecurityGroupRequest{ID: uuid})
	require.NoError(t, err, deleteFailureMessage(uuid, resp))
}

// CreateVirtualNetwork creates a virtual network resource from given service.
func CreateVirtualNetwork(t *testing.T, s services.WriteService, obj *models.VirtualNetwork) *models.VirtualNetwork {
	resp, err := s.CreateVirtualNetwork(context.Background(), &services.CreateVirtualNetworkRequest{VirtualNetwork: obj})
	require.NoError(t, err, createFailureMessage(obj, resp))
	return resp.GetVirtualNetwork()
}

// GetVirtualNetwork gets a virtual network resource from given service.
func GetVirtualNetwork(t *testing.T, s services.ReadService, uuid string) *models.VirtualNetwork {
	resp, err := s.GetVirtualNetwork(context.Background(), &services.GetVirtualNetworkRequest{ID: uuid})
	require.NoError(t, err, getFailureMessage(uuid, resp))
	return resp.GetVirtualNetwork()
}

// DeleteVirtualNetwork deletes a virtual network resource from given service.
func DeleteVirtualNetwork(t *testing.T, s services.WriteService, uuid string) {
	resp, err := s.DeleteVirtualNetwork(context.Background(), &services.DeleteVirtualNetworkRequest{ID: uuid})
	require.NoError(t, err, deleteFailureMessage(uuid, resp))
}

func createFailureMessage(obj, response interface{}) string {
	return fmt.Sprintf("create failed\nrequest object: %+v\nresponse: %+v\n", obj, response)
}

func getFailureMessage(uuid string, response interface{}) string {
	return fmt.Sprintf("get failed\nUUID: %+v\nresponse: %+v\n", uuid, response)
}

func deleteFailureMessage(uuid string, response interface{}) string {
	return fmt.Sprintf("delete failed\nUUID: %+v\nresponse: %+v\n", uuid, response)
}
