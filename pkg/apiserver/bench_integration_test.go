package apiserver_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func BenchmarkVirtualNetworkCreate(b *testing.B) {
	tests := []struct{ numberOfExistingVNs int }{
		{0}, {100}, {200}, {300}, {400}, {500},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("create virtual network with %v previously existing VNs", tt.numberOfExistingVNs)
		b.Run(testName, func(b *testing.B) {
			withServerAndClient(b, func(hc *client.HTTP) {
				ctx := context.Background()

				projectUUID := "test_project_uuid"
				if _, err := hc.CreateProject(ctx, &services.CreateProjectRequest{
					Project: &models.Project{
						UUID:       projectUUID,
						ParentType: integration.DomainType,
						ParentUUID: integration.DefaultDomainUUID,
						Name:       "test_project",
						Quota:      &models.QuotaType{},
					},
				}); err != nil {
					b.Fatal("Failed to create project: ", err)
				}
				defer hc.DeleteProject(ctx, &services.DeleteProjectRequest{ID: projectUUID}) // nolint: errcheck

				vn := &models.VirtualNetwork{
					ParentType: integration.ProjectType,
					ParentUUID: projectUUID,
				}

				for i := 0; i < tt.numberOfExistingVNs; i++ {
					vn.UUID = "test_exisiting_vn_uuid_" + strconv.Itoa(i)
					vn.Name = "test_exisiting_vn_" + strconv.Itoa(i)

					if _, err := hc.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{VirtualNetwork: vn}); err != nil {
						b.Fatal("Failed to create VN: ", err)
					}

					defer hc.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{ID: vn.UUID}) // nolint: errcheck
				}

				vn.UUID = "test_vn_uuid"
				vn.Name = "test_vn"

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					if _, err := hc.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{VirtualNetwork: vn}); err != nil {
						b.Fatal("Failed to create VN: ", err)
					}

					// Do not include delete times in the benchmark
					b.StopTimer()
					if _, err := hc.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{ID: vn.UUID}); err != nil {
						b.Fatal("Failed to delete VN: ", err)
					}
					b.StartTimer()
				}

				b.StopTimer()
			})
		})
	}
}

func withServerAndClient(t testing.TB, test func(*client.HTTP)) {
	s, err := integration.NewRunningServer(&integration.APIServerConfig{
		RepoRootPath:  "../../..",
		LogLevel:      "warn",
		DisableLogAPI: true,
	})
	if err != nil {
		t.Fatal("creating API server failed", err)
	}

	defer func() {
		if err = s.Close(); err != nil {
			t.Error("closing API Server failed", err)
		}
	}()

	hc, err := integration.NewHTTPClient(s.URL())
	if err != nil {
		t.Fatal("connecting to API Server failed")
	}

	test(hc)
}
