package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
)

func bgpRouterSetupReadServiceMocks(s *ContrailTypeLogicService, subCluster *models.SubCluster) {
	readService := s.ReadService.(*servicesmock.MockReadService)

	readService.EXPECT().GetSubCluster(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetSubClusterResponse{
			SubCluster: subCluster,
		},
		nil,
	).AnyTimes()

}

func TestCreateBGPRouter(t *testing.T) {
	tests := []struct {
		name          string
		testBGPRouter *models.BGPRouter
		subCluster    *models.SubCluster
		errorCode     codes.Code
	}{
		{
			name:          "without parameters",
			testBGPRouter: &models.BGPRouter{},
		},
		{
			name: "with asn only in subcluster",
			testBGPRouter: &models.BGPRouter{
				SubClusterRefs: []*models.BGPRouterSubClusterRef{
					{
						UUID: "sub-cluster-1",
					},
				},
			},
			subCluster: &models.SubCluster{
				SubClusterAsn: 64512,
			},
		},
		{
			name: "with equal asn in bgp router and subcluster",
			testBGPRouter: &models.BGPRouter{
				BGPRouterParameters: &models.BgpRouterParams{
					AutonomousSystem: 64512,
				},
				SubClusterRefs: []*models.BGPRouterSubClusterRef{
					{
						UUID: "sub-cluster-1",
					},
				},
			},
			subCluster: &models.SubCluster{
				SubClusterAsn: 64512,
			},
		},
		{
			name: "with not equal asn in bgp router and subcluster",
			testBGPRouter: &models.BGPRouter{
				BGPRouterParameters: &models.BgpRouterParams{
					AutonomousSystem: 64512,
				},
				SubClusterRefs: []*models.BGPRouterSubClusterRef{
					{
						UUID: "sub-cluster-1",
					},
				},
			},
			subCluster: &models.SubCluster{
				SubClusterAsn: 11111,
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "without asn in BGPRouterParameters",
			testBGPRouter: &models.BGPRouter{
				BGPRouterParameters: &models.BgpRouterParams{
					Vendor: "contrail",
				},
				SubClusterRefs: []*models.BGPRouterSubClusterRef{
					{
						UUID: "sub-cluster-1",
					},
				},
			},
			subCluster: &models.SubCluster{
				SubClusterAsn: 64512,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			bgpRouterSetupReadServiceMocks(service, tt.subCluster)

			ctx := context.Background()

			paramRequest := services.CreateBGPRouterRequest{BGPRouter: tt.testBGPRouter}
			expectedResponse := services.CreateBGPRouterResponse{BGPRouter: tt.testBGPRouter}

			createBGPRCall := service.Next().(*servicesmock.MockService).EXPECT().CreateBGPRouter(
				gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
			).DoAndReturn(
				func(_ context.Context, request *services.CreateBGPRouterRequest,
				) (response *services.CreateBGPRouterResponse, err error) {
					return &services.CreateBGPRouterResponse{BGPRouter: tt.testBGPRouter}, nil
				},
			)

			if tt.errorCode != codes.OK {
				createBGPRCall.MaxTimes(1)
			} else {
				createBGPRCall.Times(1)
			}

			createBGPRouterResponse, err := service.CreateBGPRouter(ctx, &paramRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err)
				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &expectedResponse, createBGPRouterResponse)
			}
		})
	}
}
