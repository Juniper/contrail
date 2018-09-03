package types

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam/mock"
	"github.com/Juniper/contrail/pkg/types/mock"
)

func makeMockedContrailTypeLogicService(controller *gomock.Controller) *ContrailTypeLogicService {
	service := &ContrailTypeLogicService{
		AddressManager:    ipammock.NewMockAddressManager(controller),
		ReadService:       servicesmock.NewMockReadService(controller),
		IntPoolAllocator:  typesmock.NewMockIntPoolAllocator(controller),
		InTransactionDoer: typesmock.NewMockInTransactionDoer(controller),
		MetadataGetter:    typesmock.NewMockMetadataGetter(controller),
		WriteService:      servicesmock.NewMockWriteService(controller),
	}
	service.SetNext(servicesmock.NewMockService(controller))

	service.InTransactionDoer.(*typesmock.MockInTransactionDoer).
		EXPECT().DoInTransaction(gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(ctx context.Context, do func(context.Context) error) error {
			return do(ctx)
		},
	).AnyTimes()

	return service
}

func mockedReadServiceAddVirtualNetwork(s *ContrailTypeLogicService, virtualNetwork *models.VirtualNetwork) {
	readServiceMock := s.ReadService.(*servicesmock.MockReadService)

	readServiceMock.EXPECT().GetVirtualNetwork(gomock.Not(gomock.Nil()),
		&services.GetVirtualNetworkRequest{
			ID: virtualNetwork.UUID,
		}).Return(
		&services.GetVirtualNetworkResponse{
			VirtualNetwork: virtualNetwork,
		}, nil).AnyTimes()
}

func runTest(t *testing.T, name string, test func(t *testing.T, sv *ContrailTypeLogicService)) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sv := makeMockedContrailTypeLogicService(ctrl)

	t.Run(name, func(t *testing.T) {
		test(t, sv)
	})
}
