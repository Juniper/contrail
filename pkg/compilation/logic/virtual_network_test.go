package logic

import (
	"context"
	"sync"
	"testing"

	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func TestCreateVirtualNetworkStoresVNIntent(t *testing.T) {
	compilationif.Init()

	c := gomock.NewController(t)
	defer c.Finish()
	s := NewService(servicesmock.NewMockWriteService(c))

	vn := &models.VirtualNetwork{UUID: "test-network-uuid"}

	_, err := s.CreateVirtualNetwork(context.Background(), &services.CreateVirtualNetworkRequest{
		VirtualNetwork: vn,
	})

	assert.NoError(t, err)
	checkVNIntentStoredInCache(t, vn)
}

func checkVNIntentStoredInCache(t *testing.T, vn *models.VirtualNetwork) {
	syncMap, ok := compilationif.ObjsCache.Load(virtualNetworkIntentKey)
	if assert.True(t, ok, "no expected virtual network intent sync map in cache") {
		v, ok := syncMap.(*sync.Map).Load(vn.GetUUID())
		if assert.True(t, ok, "no expected virtual network intent object in cache") {
			assert.Equal(t, vn, v.(*VirtualNetworkIntent).VirtualNetwork)
		}
	}
}

func TestVirtualNetworkIntentEvaluate(t *testing.T) {
	tests := []struct {
		name             string
		virtualNetwork   *models.VirtualNetwork
		expectedRIUpdate *models.RoutingInstance
		riUpdateFails    bool
		fails            bool
	}{
		{
			name:           "succeeds given no route target list",
			virtualNetwork: &models.VirtualNetwork{UUID: "test-network-uuid"},
		},
		{
			name: "fails given VN with no child routing instances",
			virtualNetwork: &models.VirtualNetwork{
				UUID: "test-network-uuid",
				ImportRouteTargetList: &models.RouteTargetList{
					RouteTarget: []string{
						"target:111:1111",
					},
				},
			},
			fails: true,
		},
		{
			name: "fails when routing instance update fails",
			virtualNetwork: &models.VirtualNetwork{
				UUID: "test-network-uuid",
				ImportRouteTargetList: &models.RouteTargetList{
					RouteTarget: []string{
						"target:111:1111",
					},
				},
				RoutingInstances: []*models.RoutingInstance{{UUID: "test-routing-instance-uuid"}},
			},
			expectedRIUpdate: &models.RoutingInstance{
				UUID: "test-routing-instance-uuid",
				RouteTargetRefs: []*models.RoutingInstanceRouteTargetRef{
					{
						To:   []string{"target:111:1111"},
						Attr: &models.InstanceTargetType{ImportExport: "import"},
					},
				},
			},
			riUpdateFails: true,
			fails:         true,
		},
		{
			name: "adds RI route target refs given VN with export route targets",
			virtualNetwork: &models.VirtualNetwork{
				UUID: "test-network-uuid",
				ExportRouteTargetList: &models.RouteTargetList{
					RouteTarget: []string{
						"target:111:3333",
						"target:111:4444",
						"target:111:5555",
					},
				},
				RoutingInstances: []*models.RoutingInstance{{UUID: "test-routing-instance-uuid"}},
			},
			expectedRIUpdate: &models.RoutingInstance{
				UUID: "test-routing-instance-uuid",
				RouteTargetRefs: []*models.RoutingInstanceRouteTargetRef{
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
		{
			name: "adds RI route target refs given VN with import and export route targets",
			virtualNetwork: &models.VirtualNetwork{
				UUID: "test-network-uuid",
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
				RoutingInstances: []*models.RoutingInstance{{UUID: "test-routing-instance-uuid"}},
			},
			expectedRIUpdate: &models.RoutingInstance{
				UUID: "test-routing-instance-uuid",
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
			compilationif.Init()
			vnIntent := VirtualNetworkIntent{VirtualNetwork: tt.virtualNetwork}

			c := gomock.NewController(t)
			defer c.Finish()
			mockAPIClient := servicesmock.NewMockWriteService(c)

			if tt.expectedRIUpdate != nil {
				expectRoutingInstanceUpdate(mockAPIClient, tt.expectedRIUpdate, tt.riUpdateFails)
			}

			err := vnIntent.Evaluate(
				context.Background(),
				&EvaluateContext{
					WriteService: mockAPIClient,
				},
			)

			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func expectRoutingInstanceUpdate(
	mockAPIClient *servicesmock.MockWriteService, expectedRIUpdate *models.RoutingInstance, riUpdateFails bool,
) {
	riUpdateCall := mockAPIClient.EXPECT().UpdateRoutingInstance(
		testutil.NotNil(),
		gomock.Eq(&services.UpdateRoutingInstanceRequest{
			RoutingInstance: expectedRIUpdate,
		}),
	).Times(1)

	if riUpdateFails {
		riUpdateCall.Return(nil, assert.AnError)
	} else {
		riUpdateCall.Return(nil, nil)
	}
}
