package logic

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func TestCreateVirtualNetworkEvaluatesDependencies(t *testing.T) {
	t.Skip("TODO") // TODO: fix

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	s := NewService(servicesmock.NewMockWriteService(mockCtrl))

	_, err := s.CreateVirtualNetwork(context.Background(), &services.CreateVirtualNetworkRequest{
		VirtualNetwork: &models.VirtualNetwork{},
	})

	assert.NoError(t, err)
}

func TestVirtualNetworkIntentEvaluate(t *testing.T) {
	tests := []struct {
		name           string
		virtualNetwork *models.VirtualNetwork
		expectedRI     *models.RoutingInstance
		fails          bool
	}{
		{
			name: "adds routing instance route target refs given import and export route targets",
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
			expectedRI: &models.RoutingInstance{
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
			vnIntent := VirtualNetworkIntent{VirtualNetwork: tt.virtualNetwork}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockService := servicesmock.NewMockWriteService(mockCtrl)
			//mockService.EXPECT().UpdateRoutingInstance(
			//	testutil.NotNil(),
			//	testutil.NotNil(),
			//).Return(&services.UpdateRoutingInstanceResponse{}, nil).Times(1)

			err := vnIntent.Evaluate(
				context.Background(),
				&EvaluateContext{
					WriteService: NewService(mockService),
				},
			)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				//assert.Equal(t, tt.expected, actual) // TODO: check updated RI
			}
		})
	}
}
