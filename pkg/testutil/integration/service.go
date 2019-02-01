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

// CreateNetworkPolicy creates a network policy resource in given service.
func CreateNetworkPolicy(
	ctx context.Context, t *testing.T, s services.WriteService, obj *models.NetworkPolicy,
) *models.NetworkPolicy {
	resp, err := s.CreateNetworkPolicy(ctx, &services.CreateNetworkPolicyRequest{NetworkPolicy: obj})
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating NetworkPolicy failed\n requested: %+v\n "+
			"response: %+v\n", obj, resp),
	)
	return resp.GetNetworkPolicy()
}

// DeleteNetworkPolicy deletes a network policy resource from given service.
func DeleteNetworkPolicy(
	ctx context.Context, t *testing.T, s services.WriteService, id string,
) {
	resp, err := s.DeleteNetworkPolicy(ctx, &services.DeleteNetworkPolicyRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("deleting NetworkPolicy failed\n UUID: %+v\n "+
			"response: %+v\n", id, resp),
	)
}

// GetNetworkPolicy gets a network policy resource from given service.
func GetNetworkPolicy(
	ctx context.Context, t *testing.T, s services.ReadService, id string,
) *models.NetworkPolicy {
	resp, err := s.GetNetworkPolicy(ctx, &services.GetNetworkPolicyRequest{ID: id})
	require.NoError(
		t,
		err,
		fmt.Sprintf("getting NetworkPolicy failed\n id: %+v\n "+
			"response: %+v\n", id, resp),
	)
	return resp.GetNetworkPolicy()
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
