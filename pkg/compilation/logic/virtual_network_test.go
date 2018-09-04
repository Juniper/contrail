package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

func TestVirtualNetworkIntentEvaluate(t *testing.T) {
	tests := []struct {
		name           string
		virtualNetwork *models.VirtualNetwork
		expected       *models.RoutingInstance
		fails          bool
	}{
		{
			name: "adds routing_instance.route_target_refs given import and export route targets",
			virtualNetwork: &models.VirtualNetwork{
				Name:   "test-network",
				FQName: []string{"test-network"},
				UUID:   "test-network-uuid",
				ImportRouteTargetList: &models.RouteTargetList{
					RouteTarget: []string{
						"target:111:1111",
						"target:111:2222",
					},
				},
				ExportRouteTargetList: &models.RouteTargetList{
					RouteTarget: []string{
						"target:111:3333",
						"target:111:4444",
						"target:111:5555",
					},
				},
			},
			expected: &models.RoutingInstance{
				Name:                     "test-network",
				FQName:                   []string{"test-network", "test-network"},
				ParentUUID:               "test-network-uuid",
				RoutingInstanceIsDefault: true,
				RouteTargetRefs: []*models.RoutingInstanceRouteTargetRef{
					{
						To:   []string{"target:111:1111"},
						Attr: &models.InstanceTargetType{ImportExport: "import"},
					},
					{
						To:   []string{"target:111:2222"},
						Attr: &models.InstanceTargetType{ImportExport: "import"},
					},
					{
						To:   []string{"target:111:3333"},
						Attr: &models.InstanceTargetType{ImportExport: "export"},
					},
					{
						To:   []string{"target:111:4444"},
						Attr: &models.InstanceTargetType{ImportExport: "export"},
					},
					{
						To:   []string{"target:111:5555"},
						Attr: &models.InstanceTargetType{ImportExport: "export"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vn := VirtualNetworkIntent{VirtualNetwork: tt.virtualNetwork}

			err := vn.Evaluate(context.Background(), &EvaluateContext{})

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				//assert.Equal(t, tt.expected, actual) // TODO
			}
		})
	}
}
