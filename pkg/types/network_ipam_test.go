package types

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"github.com/Juniper/contrail/pkg/types/ipam/mock"
)

type testNetIpamParams struct {
	uuid             string
	ipamSubnetMethod string
	ipamSubnets      *models.IpamSubnets
}

func createTestNetworkIpam(testParams *testNetIpamParams) *models.NetworkIpam {
	networkIpam := models.MakeNetworkIpam()
	networkIpam.UUID = testParams.uuid
	networkIpam.IpamSubnetMethod = testParams.ipamSubnetMethod
	networkIpam.IpamSubnets = testParams.ipamSubnets
	return networkIpam
}

func networkIpamNextServMocks(service *ContrailTypeLogicService) {
	ipamMock := service.AddressManager.(*ipammock.MockAddressManager)
	ipamMock.EXPECT().CreateIpamSubnet(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, request *ipam.CreateIpamSubnetRequest) (subnetUUID string, err error) {
			return "uuuu-uuuu-iiii-dddd", nil
		}).AnyTimes()
	ipamMock.EXPECT().DeleteIpamSubnet(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

}

func networkIpamIPAMMocks(service *ContrailTypeLogicService) {
	nextServiceMock := service.Next().(*servicesmock.MockService)
	nextServiceMock.EXPECT().CreateNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.CreateNetworkIpamRequest,
		) (response *services.CreateNetworkIpamResponse, err error) {
			return &services.CreateNetworkIpamResponse{NetworkIpam: request.NetworkIpam}, nil
		}).AnyTimes()
	nextServiceMock.EXPECT().DeleteNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.DeleteNetworkIpamRequest,
		) (response *services.DeleteNetworkIpamResponse, err error) {
			return &services.DeleteNetworkIpamResponse{ID: request.ID}, nil
		}).AnyTimes()

}

func TestCreateNetworkIpam(t *testing.T) {
	tests := []struct {
		name              string
		testNetIpamParams *testNetIpamParams
		fails             bool
		changedSubnetUUID bool
		errorCode         codes.Code
	}{
		{
			name: "Try to create network ipam with empty ipam_subnets list",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "notFlat",
			},
			fails:             false,
			changedSubnetUUID: false,
		},
		{
			name: "Try to create network ipam with empty flat subnet",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets:      &models.IpamSubnets{},
			},
			fails:             false,
			changedSubnetUUID: false,
		},
		{
			name: "Try to create network ipam with non-empty ipam_subnets list and not flat subnet",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "notFlat",
				ipamSubnets:      &models.IpamSubnets{},
			},
			fails:             true,
			errorCode:         codes.InvalidArgument,
			changedSubnetUUID: false,
		},
		{
			name: "Try to create network ipam with specific ipam_subnets",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				}}},
			},
			fails:             false,
			changedSubnetUUID: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			service := makeMockedContrailTypeLogicService(mockCtrl)
			networkIpamIPAMMocks(service)
			networkIpamNextServMocks(service)

			ctx := context.Background()

			networkIpam := createTestNetworkIpam(tt.testNetIpamParams)
			createNetworkIpamRequest := &services.CreateNetworkIpamRequest{NetworkIpam: networkIpam}
			createNetworkIpamResponse, err := service.CreateNetworkIpam(ctx, createNetworkIpamRequest)

			if tt.fails {
				assert.Error(t, err, "create succeeded but shouldn't")
				assert.Nil(t, createNetworkIpamResponse)

				if tt.errorCode != codes.OK {
					status, ok := status.FromError(err)
					assert.True(t, ok)
					assert.Equal(t, tt.errorCode, status.Code())
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, createNetworkIpamResponse)
			}

			if tt.changedSubnetUUID {
				receivedUUID := createNetworkIpamResponse.NetworkIpam.IpamSubnets.Subnets[0].SubnetUUID
				assert.Equal(t, "uuuu-uuuu-iiii-dddd", receivedUUID)
			}
			mockCtrl.Finish()
		})
	}

}

func TestDeleteNetworkIpam(t *testing.T) {
	deleteIpamDBMock := func(service *ContrailTypeLogicService, getNetworkIpamResponse *services.GetNetworkIpamResponse) {
		dataService := service.DataService.(*servicesmock.MockService)
		dataService.EXPECT().GetNetworkIpam(gomock.Any(), gomock.Any()).Return(getNetworkIpamResponse, nil)
	}

	tests := []struct {
		name              string
		testNetIpamParams *testNetIpamParams
	}{
		{
			name: "Delete network ipam without ipam_subnets",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "notFlat",
			},
		},
		{
			name: "Delete network ipam with non-empty ipam_subnets list",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24},
				}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			service := makeMockedContrailTypeLogicService(mockCtrl)
			networkIpamIPAMMocks(service)
			networkIpamNextServMocks(service)

			ctx := context.Background()

			networkIpam := createTestNetworkIpam(tt.testNetIpamParams)
			getNetworkIpamResponse := &services.GetNetworkIpamResponse{NetworkIpam: networkIpam}
			deleteIpamDBMock(service, getNetworkIpamResponse)
			deleteNetworkIpamRequest := &services.DeleteNetworkIpamRequest{
				ID: getNetworkIpamResponse.GetNetworkIpam().GetUUID(),
			}
			deleteNetworkIpamResponse, err := service.DeleteNetworkIpam(ctx, deleteNetworkIpamRequest)

			assert.NoError(t, err)
			assert.NotNil(t, deleteNetworkIpamResponse)
			mockCtrl.Finish()
		})
	}
}
