package services_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	projectBluePath        = "test_data/project_blue.yml"
	virtualNetworkBluePath = "test_data/virtual_network_blue.yml"
	networkIpamBluePath    = "test_data/network_ipam_blue.yml"
)

type uuidAndKind interface {
	GetUUID() string
	Kind() string
}

func TestApplyHref(t *testing.T) {
	runTest(t, func(t *testing.T, client *integration.HTTPAPIClient, server *integration.APIServer) {
		ctx := context.Background()

		p, deleteProject := createResource(t, ctx, client, loadProject(t, projectBluePath))
		defer deleteProject()
		_, deleteNetworkIpam := createResource(t, ctx, client, loadNetworkIpam(t, networkIpamBluePath))
		defer deleteNetworkIpam()
		vn, deleteVirtualNetwork := createResource(t, ctx, client, loadVirtualNetwork(t, virtualNetworkBluePath))
		defer deleteVirtualNetwork()

		virtualNetwork := integration.GetVirtualNetwork(t, client, vn.GetUUID())
		require.Equal(t, 1, len(virtualNetwork.NetworkIpamRefs))
		ipamRef := virtualNetwork.NetworkIpamRefs[0]

		project := integration.GetProject(t, client, p.GetUUID())
		require.Equal(t, 1, len(project.VirtualNetworks))
		childNetwork := project.VirtualNetworks[0]

		networkIpam := integration.GetNetworkIpam(t, client, ipamRef.GetUUID())
		require.Equal(t, 1, len(networkIpam.VirtualNetworkBackRefs))
		backReferencedVirtualNetwork := networkIpam.VirtualNetworkBackRefs[0]

		tests := []struct {
			name        string
			uuidAndKind uuidAndKind
			actual      string
		}{
			{
				name:        "resource href",
				uuidAndKind: virtualNetwork,
				actual:      virtualNetwork.Href,
			},
			{
				name:        "network ipam reference href",
				uuidAndKind: newRefToUUIDAndKindAdapter(ipamRef),
				actual:      ipamRef.Href,
			},
			{
				name:        "project child vn href",
				uuidAndKind: childNetwork,
				actual:      childNetwork.Href,
			},
			{
				name:        "network ipam back-referenced vn href",
				uuidAndKind: backReferencedVirtualNetwork,
				actual:      backReferencedVirtualNetwork.Href,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				expectedHref := server.URL() + "/" + tt.uuidAndKind.Kind() + "/" + tt.uuidAndKind.GetUUID()
				assert.Equal(t, expectedHref, tt.actual)
			})
		}
	})
}

func deleteResource(t *testing.T, ctx context.Context, client services.Service, object basemodels.Object) {
	e, err := services.NewEvent(&services.EventOption{
		UUID:      object.GetUUID(),
		Kind:      object.Kind(),
		Operation: services.OperationDelete,
	})
	require.NoError(t, err)
	_, err = e.Process(ctx, client)
	require.NoError(t, err)
}

func createResource(t *testing.T, ctx context.Context, client services.Service, object basemodels.Object) (basemodels.Object, func()) {
	e, err := services.NewEvent(&services.EventOption{
		UUID: object.GetUUID(),
		Kind: object.Kind(),
		Data: object.ToMap(),
	})
	require.NoError(t, err)
	resp, err := e.Process(ctx, client)
	require.NoError(t, err)
	return resp.GetResource(), func() {
		deleteResource(t, ctx, client, resp.GetResource())
	}
}

func loadProject(t *testing.T, path string) (project *models.Project) {
	require.NoError(t, fileutil.LoadFile(path, &project))
	return project
}

func loadVirtualNetwork(t *testing.T, path string) (vn *models.VirtualNetwork) {
	require.NoError(t, fileutil.LoadFile(path, &vn))
	return vn
}

func loadNetworkIpam(t *testing.T, path string) (ni *models.NetworkIpam) {
	require.NoError(t, fileutil.LoadFile(path, &ni))
	return ni
}

func runTest(t *testing.T, test func(*testing.T, *integration.HTTPAPIClient, *integration.APIServer)) {
	for _, driver := range []string{basedb.DriverMySQL, basedb.DriverPostgreSQL} {
		func() {
			s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
				DBDriver:     driver,
				RepoRootPath: "../../..",
			})
			defer func() {
				assert.NoError(t, s.Close())
			}()

			test(t, integration.NewTestingHTTPClient(t, s.URL()), s)
		}()
	}
}

type refToUUIDAndKindAdapter struct {
	basemodels.Reference
}

func newRefToUUIDAndKindAdapter(r basemodels.Reference) *refToUUIDAndKindAdapter {
	return &refToUUIDAndKindAdapter{
		Reference: r,
	}
}

func (r *refToUUIDAndKindAdapter) Kind() string {
	return r.GetReferredKind()
}
