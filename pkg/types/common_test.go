package types

import (
	"testing"

	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam/mock"
	"github.com/Juniper/contrail/pkg/types/mock"
)

func makeMockedContrailTypeLogicService(t *testing.T, controller *gomock.Controller) *ContrailTypeLogicService {
	service := &ContrailTypeLogicService{
		AddressManager:    ipammock.NewMockAddressManager(controller),
		DataService:       servicesmock.NewMockService(controller),
		IntPoolAllocator:  ipammock.NewMockIntPoolAllocator(controller),
		InTransactionDoer: typesmock.NewMockInTransactionDoer(controller),
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
