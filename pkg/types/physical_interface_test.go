package types

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/models"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/mock/gomock"

	"github.com/Juniper/contrail/pkg/services"
	servicesmock "github.com/Juniper/contrail/pkg/services/mock"
)

type testPhysicalInterfaceParams struct {
	UUID                      string
	DisplayName               string
	EthernetSegmentIdentifier string
}

func physicalInterfaceNextServMocks(service *ContrailTypeLogicService) {
	nextServiceMock := service.Next().(*servicesmock.MockService) //nolint: errcheck
	nextServiceMock.EXPECT().CreatePhysicalInterface(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.CreatePhysicalInterfaceRequest,
		) (response *services.CreatePhysicalInterfaceResponse, err error) {
			return &services.CreatePhysicalInterfaceResponse{PhysicalInterface: request.PhysicalInterface}, nil
		}).AnyTimes()
	nextServiceMock.EXPECT().UpdatePhysicalInterface(gomock.Any(), gomock.Any()).DoAndReturn(
		func(
			_ context.Context, request *services.UpdatePhysicalInterfaceRequest,
		) (response *services.UpdatePhysicalInterfaceResponse, err error) {
			return &services.UpdatePhysicalInterfaceResponse{PhysicalInterface: request.PhysicalInterface}, nil
		}).AnyTimes()
}

func physicalInterfaceReadServiceMocks(
	s *ContrailTypeLogicService,
	parentRouter *models.PhysicalRouter,
	listPhysicalInterface []*models.PhysicalInterface,
) {
	readService := s.ReadService.(*servicesmock.MockReadService) //nolint: errcheck
	// Use empty physical router structure instead nil parent physical rounter
	if parentRouter == nil {
		parentRouter = new(models.PhysicalRouter)
	}
	readService.EXPECT().GetPhysicalRouter(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.GetPhysicalRouterResponse{
			PhysicalRouter: parentRouter,
		}, nil).AnyTimes()
	// Try to find requested physical interface within parent router data
	readService.EXPECT().GetPhysicalInterface(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(
			_ context.Context, request *services.GetPhysicalInterfaceRequest,
		) (response *services.GetPhysicalInterfaceResponse, err error) {
			for _, pi := range parentRouter.PhysicalInterfaces {
				if pi.UUID == request.ID {
					return &services.GetPhysicalInterfaceResponse{PhysicalInterface: pi}, nil
				}
			}
			return nil, grpc.Errorf(codes.NotFound, "physical interface with uuid %s not found", request.ID)
		}).AnyTimes()
	// Use empty physical interface slice instead nil others physical interface list
	readService.EXPECT().ListPhysicalInterface(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).Return(
		&services.ListPhysicalInterfaceResponse{
			PhysicalInterfaces: listPhysicalInterface,
		}, nil).AnyTimes()
	// Try to find requested logical interface within current physical router and/or list of physical interface data
	readService.EXPECT().GetLogicalInterface(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(
			_ context.Context, request *services.GetLogicalInterfaceRequest,
		) (response *services.GetLogicalInterfaceResponse, err error) {
			for _, pi := range parentRouter.PhysicalInterfaces {
				for _, li := range pi.LogicalInterfaces {
					if li.UUID == request.ID {
						return &services.GetLogicalInterfaceResponse{LogicalInterface: li}, nil
					}
				}
			}
			for _, pi := range listPhysicalInterface {
				for _, li := range pi.LogicalInterfaces {
					if li.UUID == request.ID {
						return &services.GetLogicalInterfaceResponse{LogicalInterface: li}, nil
					}
				}
			}
			return nil, grpc.Errorf(codes.NotFound, "logical interface with uuid %s not found", request.ID)
		}).AnyTimes()
}

func TestCreatePhysicalInterface(t *testing.T) {
	tests := []struct {
		name          string
		createRequest *services.CreatePhysicalInterfaceRequest
		parentRouter  *models.PhysicalRouter
		errorCode     codes.Code
	}{
		{
			name: "Try create physical interface with empty display name and ESI",
			createRequest: &services.CreatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID: "uuid",
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try create physical interface with valid display name and ESI",
			createRequest: &services.CreatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:                      "uuid",
				DisplayName:               "unique",
				EthernetSegmentIdentifier: "01:23:45:67:89:1A:BC:DE:F0:01",
			}},
			parentRouter: &models.PhysicalRouter{
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:        "uuid",
					DisplayName: "other",
				}},
			},
			errorCode: codes.OK,
		},
		{
			name: "Try create physical interface with same display name",
			createRequest: &services.CreatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:        "uuid",
				DisplayName: "double",
			}},
			parentRouter: &models.PhysicalRouter{
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:        "uuid",
					DisplayName: "double",
				}},
			},
			errorCode: codes.AlreadyExists,
		},
		{
			name: "Try create physical interface with wrong ESI",
			createRequest: &services.CreatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:                      "uuid",
				EthernetSegmentIdentifier: "0123:4567:891A:BCDE:F001",
			}},
			errorCode: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			physicalInterfaceNextServMocks(service)
			ctx := context.Background()

			physicalInterfaceReadServiceMocks(service, tt.parentRouter, nil)

			createPhysicalInterfaceResponse, err := service.CreatePhysicalInterface(ctx, tt.createRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err, "create succeeded but shouldn't")
				assert.Nil(t, createPhysicalInterfaceResponse)

				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, createPhysicalInterfaceResponse)
			}
		})
	}
}

func TestUpdatePhysicalInterface(t *testing.T) {
	tests := []struct {
		name                  string
		updateRequest         *services.UpdatePhysicalInterfaceRequest
		parentRouter          *models.PhysicalRouter
		listPhysicalInterface []*models.PhysicalInterface
		errorCode             codes.Code
	}{
		{
			name: "Try update physical interface without display name and ESI",
			updateRequest: &services.UpdatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID: "uuid",
			}},
			parentRouter: &models.PhysicalRouter{
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID: "uuid",
				}},
			},
			errorCode: codes.OK,
		},
		{
			name: "Try update physical interface with different display name",
			updateRequest: &services.UpdatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:        "uuid",
				DisplayName: "second",
			}},
			parentRouter: &models.PhysicalRouter{
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:        "uuid",
					DisplayName: "first",
				}},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try update physical interface with valid ESI name",
			updateRequest: &services.UpdatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:                      "uuid",
				EthernetSegmentIdentifier: "00:11:22:33:44:55:66:77:88:99",
			}},
			parentRouter: &models.PhysicalRouter{
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:                      "uuid",
					EthernetSegmentIdentifier: "01:02:03:04:05:06:07:08:09:10",
				}},
			},
			errorCode: codes.OK,
		},
		{
			name: "Try update physical interface with wrong ESI name",
			updateRequest: &services.UpdatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:                      "uuid",
				EthernetSegmentIdentifier: "AAAA:BBBB:CCCC:DDDD:EEEE",
			}},
			parentRouter: &models.PhysicalRouter{
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:                      "uuid",
					EthernetSegmentIdentifier: "01:02:03:04:05:06:07:08:09:10",
				}},
			},
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Try update physical interface with same VMI refs",
			updateRequest: &services.UpdatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:                      "uuid",
				EthernetSegmentIdentifier: "01:02:03:04:05:06:07:08:09:10",
			}},
			parentRouter: &models.PhysicalRouter{
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:                      "uuid",
					EthernetSegmentIdentifier: "01:02:03:04:05:06:07:08:09:10",
					LogicalInterfaces: []*models.LogicalInterface{{
						UUID:                    "uuid_0",
						LogicalInterfaceVlanTag: 1024,
						VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
							UUID: "vmi0",
						}},
					}},
				}},
			},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID:                      "uuid",
				EthernetSegmentIdentifier: "01:02:03:04:05:06:07:08:09:10",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID:                    "uuid_1",
					LogicalInterfaceVlanTag: 1024,
					VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
						UUID: "vmi0",
					}},
				}},
			}},
			errorCode: codes.OK,
		},
		{
			name: "Try update physical interface with different VMI refs",
			updateRequest: &services.UpdatePhysicalInterfaceRequest{PhysicalInterface: &models.PhysicalInterface{
				UUID:                      "uuid",
				EthernetSegmentIdentifier: "00:11:22:33:44:55:66:77:88:99",
			}},
			parentRouter: &models.PhysicalRouter{
				PhysicalInterfaces: []*models.PhysicalInterface{{
					UUID:                      "uuid",
					EthernetSegmentIdentifier: "01:02:03:04:05:06:07:08:09:10",
					LogicalInterfaces: []*models.LogicalInterface{{
						UUID:                    "uuid_0",
						LogicalInterfaceVlanTag: 1024,
						VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
							UUID: "vmi0_0",
						}},
					},
						{
							UUID:                    "uuid_1",
							LogicalInterfaceVlanTag: 2048,
							VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
								UUID: "vmi1_0",
							}},
						}},
				}},
			},
			listPhysicalInterface: []*models.PhysicalInterface{{
				UUID:                      "uuid",
				EthernetSegmentIdentifier: "00:11:22:33:44:55:66:77:88:99",
				LogicalInterfaces: []*models.LogicalInterface{{
					UUID:                    "uuid_2",
					LogicalInterfaceVlanTag: 1024,
					VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
						UUID: "vmi0_0",
					}},
				},
					{
						UUID:                    "uuid_3",
						LogicalInterfaceVlanTag: 2048,
						VirtualMachineInterfaceRefs: []*models.LogicalInterfaceVirtualMachineInterfaceRef{{
							UUID: "vmi1_9",
						}},
					}},
			}},
			errorCode: codes.AlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			service := makeMockedContrailTypeLogicService(mockCtrl)
			physicalInterfaceNextServMocks(service)
			ctx := context.Background()

			physicalInterfaceReadServiceMocks(service, tt.parentRouter, tt.listPhysicalInterface)

			updatePhysicalInterfaceResponse, err := service.UpdatePhysicalInterface(ctx, tt.updateRequest)
			if tt.errorCode != codes.OK {
				assert.Error(t, err, "update succeeded but shouldn't")
				assert.Nil(t, updatePhysicalInterfaceResponse)

				status, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.errorCode, status.Code())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, updatePhysicalInterfaceResponse)
			}
		})
	}
}
