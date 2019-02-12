package types

import (
	"context"
	"github.com/Juniper/contrail/pkg/types/ipam/mock"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

type testNetIpamParams struct {
	uuid                   string
	ipamSubnetMethod       string
	ipamSubnets            *models.IpamSubnets
	networkIpamMGMT        *models.IpamType
	virtualNetworkBackRefs []*models.VirtualNetwork
}

func createTestNetworkIpam(testParams *testNetIpamParams) *models.NetworkIpam {
	networkIpam := models.MakeNetworkIpam()
	networkIpam.UUID = testParams.uuid
	networkIpam.IpamSubnetMethod = testParams.ipamSubnetMethod
	networkIpam.IpamSubnets = testParams.ipamSubnets
	networkIpam.NetworkIpamMGMT = testParams.networkIpamMGMT
	if len(testParams.virtualNetworkBackRefs) > 0 {
		networkIpam.VirtualNetworkBackRefs = testParams.virtualNetworkBackRefs
	}
	return networkIpam
}

func networkIpamNextServMocks(service *ContrailTypeLogicService) {
	nextServiceMock := service.Next().(*servicesmock.MockService) //nolint: errcheck
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
	nextServiceMock.EXPECT().UpdateNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.UpdateNetworkIpamRequest,
		) (response *services.UpdateNetworkIpamResponse, err error) {
			return &services.UpdateNetworkIpamResponse{NetworkIpam: request.NetworkIpam}, nil
		}).AnyTimes()
}


func mockAddressManager(service *ContrailTypeLogicService) {
	service.AddressManager.(*ipammock.MockAddressManager). // nolint: errcheck
		EXPECT().AllocateIP(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return("", "", nil).AnyTimes()
}

func TestCreateNetworkIpam(t *testing.T) {
	tests := []struct {
		name              string
		testNetIpamParams *testNetIpamParams
		errorCode         codes.Code
	}{
		{
			name: "Try to create network ipam with empty ipam_subnets list",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "notFlat",
			},
		},
		{
			name: "Try to create network ipam with empty flat subnet",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets:      &models.IpamSubnets{},
			},
		},
		{
			name: "Try to create network ipam with non-empty ipam_subnets list and not flat subnet",
			testNetIpamParams: &testNetIpamParams{
				uuid:             "uuid",
				ipamSubnetMethod: "notFlat",
				ipamSubnets:      &models.IpamSubnets{},
			},
			errorCode: codes.InvalidArgument,
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			service := makeMockedContrailTypeLogicService(mockCtrl)
			networkIpamNextServMocks(service)
			mockAddressManager(service)

			ctx := context.Background()

			networkIpam := createTestNetworkIpam(tt.testNetIpamParams)
			createNetworkIpamRequest := &services.CreateNetworkIpamRequest{NetworkIpam: networkIpam}
			createNetworkIpamResponse, err := service.CreateNetworkIpam(ctx, createNetworkIpamRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err, "create succeeded but shouldn't")
				assert.Nil(t, createNetworkIpamResponse)

				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, createNetworkIpamResponse)
			}

			mockCtrl.Finish()
		})
	}

}

func TestUpdateNetworkIpam(t *testing.T) {
	updateIpamDBMock := func(service *ContrailTypeLogicService, getNetworkIpamResponse *services.GetNetworkIpamResponse) {
		readService := service.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
		readService.EXPECT().GetNetworkIpam(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetNetworkIpamRequest,
			) (response *services.GetNetworkIpamResponse, err error) {
				if getNetworkIpamResponse.GetNetworkIpam() == nil {
					return nil, grpc.Errorf(codes.NotFound, "ipam not found")
				}
				return getNetworkIpamResponse, nil
			})

		readService.EXPECT().GetVirtualNetwork(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetVirtualNetworkRequest,
			) (response *services.GetVirtualNetworkResponse, err error) {
				if getNetworkIpamResponse.GetNetworkIpam() == nil {
					return nil, grpc.Errorf(codes.NotFound, "vn not found")
				}
				vn := getNetworkIpamResponse.GetNetworkIpam().VirtualNetworkBackRefs[0]
				return &services.GetVirtualNetworkResponse{VirtualNetwork: vn}, nil
			}).AnyTimes()

		readService.EXPECT().GetFloatingIPPool(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetFloatingIPPoolRequest,
			) (response *services.GetFloatingIPPoolResponse, err error) {
				if getNetworkIpamResponse.GetNetworkIpam() == nil {
					return nil, grpc.Errorf(codes.NotFound, "fipp not found")
				}
				vn := getNetworkIpamResponse.GetNetworkIpam().VirtualNetworkBackRefs[0]
				fipp := vn.GetFloatingIPPools()[0]
				return &services.GetFloatingIPPoolResponse{FloatingIPPool: fipp}, nil
			}).AnyTimes()

		readService.EXPECT().GetAliasIPPool(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetAliasIPPoolRequest,
			) (response *services.GetAliasIPPoolResponse, err error) {
				if getNetworkIpamResponse.GetNetworkIpam() == nil {
					return nil, grpc.Errorf(codes.NotFound, "aipp not found")
				}
				vn := getNetworkIpamResponse.GetNetworkIpam().VirtualNetworkBackRefs[0]
				aipp := vn.GetAliasIPPools()[0]
				return &services.GetAliasIPPoolResponse{AliasIPPool: aipp}, nil
			}).AnyTimes()

		readService.EXPECT().GetInstanceIP(gomock.Any(), gomock.Any()).DoAndReturn(
			func(
				_ context.Context, _ *services.GetInstanceIPRequest,
			) (response *services.GetInstanceIPResponse, err error) {
				if getNetworkIpamResponse.GetNetworkIpam() == nil {
					return nil, grpc.Errorf(codes.NotFound, "ipp not found")
				}
				vn := getNetworkIpamResponse.GetNetworkIpam().VirtualNetworkBackRefs[0]
				ipp := vn.GetInstanceIPBackRefs()[0]
				return &services.GetInstanceIPResponse{InstanceIP: ipp}, nil
			}).AnyTimes()
	}

	tests := []struct {
		name                 string
		oldNetworkIpamParams *testNetIpamParams
		newNetworkIpamParams *testNetIpamParams
		errorCode            codes.Code
	}{
		{
			name: "Update network ipam which does not exist",
			newNetworkIpamParams: &testNetIpamParams{
				uuid: "uuid-1",
			},
			errorCode: codes.NotFound,
		},
		{
			name: "Update network ipam subnet method",
			oldNetworkIpamParams: &testNetIpamParams{
				ipamSubnetMethod: "flat-subnet",
			},
			newNetworkIpamParams: &testNetIpamParams{
				ipamSubnetMethod: "user-defined",
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Update network ipam where dns method is changed from default",
			oldNetworkIpamParams: &testNetIpamParams{
				networkIpamMGMT: &models.IpamType{IpamDNSMethod: "default-dns-server"},
				virtualNetworkBackRefs: []*models.VirtualNetwork{{UUID: "vnUUID",
					VirtualMachineInterfaceBackRefs: []*models.VirtualMachineInterface{{UUID: "vmUUID"}}}},
			},
			newNetworkIpamParams: &testNetIpamParams{
				networkIpamMGMT: &models.IpamType{IpamDNSMethod: "tenant-dns-server"},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Update network ipam where dns method is changed from virtual",
			oldNetworkIpamParams: &testNetIpamParams{
				networkIpamMGMT: &models.IpamType{IpamDNSMethod: "virtual-dns-server"},
				virtualNetworkBackRefs: []*models.VirtualNetwork{{UUID: "vnUUID",
					VirtualMachineInterfaceBackRefs: []*models.VirtualMachineInterface{{UUID: "vmUUID"}}}},
			},
			newNetworkIpamParams: &testNetIpamParams{
				networkIpamMGMT: &models.IpamType{IpamDNSMethod: "tenant-dns-server"},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Update network ipam where dns method is changed from tenant",
			oldNetworkIpamParams: &testNetIpamParams{
				networkIpamMGMT: &models.IpamType{IpamDNSMethod: "tenant-dns-server"},
				virtualNetworkBackRefs: []*models.VirtualNetwork{{UUID: "vnUUID",
					VirtualMachineInterfaceBackRefs: []*models.VirtualMachineInterface{{UUID: "vmUUID"}}}},
			},
			newNetworkIpamParams: &testNetIpamParams{
				networkIpamMGMT: &models.IpamType{IpamDNSMethod: "virtual-dns-server"},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Update network ipam where dns method is changed but there was no mgmt before",
			oldNetworkIpamParams: &testNetIpamParams{
				uuid: "ipamUUID",
			},
			newNetworkIpamParams: &testNetIpamParams{
				uuid:            "ipamUUID",
				networkIpamMGMT: &models.IpamType{IpamDNSMethod: "tenant-dns-server"},
			},
		},
		{
			name: "Update network ipam where ipam subnet method is not flat subnet",
			oldNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "user-defined",
			},
			newNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "user-defined",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Update network ipam with overlapping subnets",
			oldNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
			},
			newNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}},
					{Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Update network ipam with new ipam subnets",
			oldNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
			},
			newNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "11.0.0.0", IPPrefixLen: 24}}}},
			},
		},
		{
			name: "Update network ipam with new ipam subnets when one vnref includes instance ip",
			oldNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
				virtualNetworkBackRefs: []*models.VirtualNetwork{{
					InstanceIPBackRefs: []*models.InstanceIP{{InstanceIPAddress: "10.0.0.5"}}}},
			},
			newNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "11.0.0.0", IPPrefixLen: 24}}}},
			},
			errorCode: codes.Aborted,
		},
		{
			name: "Update network ipam with new ipam subnets when one vn ref includes floating ip",
			oldNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
				virtualNetworkBackRefs: []*models.VirtualNetwork{{FloatingIPPools: []*models.FloatingIPPool{{
					FloatingIPs: []*models.FloatingIP{{FloatingIPAddress: "10.0.0.5"}}}}}},
			},
			newNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "11.0.0.0", IPPrefixLen: 24}}}},
			},
			errorCode: codes.Aborted,
		},
		{
			name: "Update network ipam with new ipam subnets when one vn ref includes alias ip",
			oldNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
				virtualNetworkBackRefs: []*models.VirtualNetwork{{AliasIPPools: []*models.AliasIPPool{{
					AliasIPs: []*models.AliasIP{{AliasIPAddress: "10.0.0.5"}}}}}},
			},
			newNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "11.0.0.0", IPPrefixLen: 24}}}},
			},
			errorCode: codes.Aborted,
		},
		{
			name: "Update network ipam with ref to vn with subnet which overlaps with new ipam_subnet",
			oldNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				virtualNetworkBackRefs: []*models.VirtualNetwork{{UUID: "vnUUID",
					NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{{UUID: "testIpam",
						Attr: &models.VnSubnetsType{IpamSubnets: []*models.IpamSubnetType{{
							Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}},
						}}},
				}},
			},
			newNetworkIpamParams: &testNetIpamParams{
				uuid: "ipamUUID",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					Subnet: &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Update network ipam by changing default gateway to none in subnet",
			oldNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					SubnetName:     "test",
					DefaultGateway: "10.0.0.1",
					Subnet:         &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
			},
			newNetworkIpamParams: &testNetIpamParams{
				uuid:             "ipamUUID",
				ipamSubnetMethod: "flat-subnet",
				ipamSubnets: &models.IpamSubnets{Subnets: []*models.IpamSubnetType{{
					SubnetName: "test",
					Subnet:     &models.SubnetType{IPPrefix: "10.0.0.0", IPPrefixLen: 24}}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			service := makeMockedContrailTypeLogicService(mockCtrl)
			networkIpamNextServMocks(service)

			ctx := context.Background()
			var fMask types.FieldMask
			var oldNetworkIpam, newNetworkIpam *models.NetworkIpam
			if tt.oldNetworkIpamParams != nil {
				oldNetworkIpam = createTestNetworkIpam(tt.oldNetworkIpamParams)
			}
			if tt.newNetworkIpamParams != nil {
				newNetworkIpam = createTestNetworkIpam(tt.newNetworkIpamParams)
				fMask = basemodels.MapToFieldMask(newNetworkIpam.ToMap())
			}

			getNetworkIpamResponse := &services.GetNetworkIpamResponse{NetworkIpam: oldNetworkIpam}
			updateIpamDBMock(service, getNetworkIpamResponse)
			updateNetworkIpamRequest := &services.UpdateNetworkIpamRequest{NetworkIpam: newNetworkIpam, FieldMask: fMask}
			updateNetworkIpamResponse, err := service.UpdateNetworkIpam(ctx, updateNetworkIpamRequest)

			if tt.errorCode != codes.OK {
				assert.Error(t, err, "update succeeded but shouldn't")
				assert.Nil(t, updateNetworkIpamResponse)

				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, updateNetworkIpamResponse)
			}

			mockCtrl.Finish()
		})
	}
}
