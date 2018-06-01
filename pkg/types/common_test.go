package types

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/types/ipam/mock"
	"github.com/Juniper/contrail/pkg/types/mock"
)

func makeMockedContrailTypeLogicService(t *testing.T, controller *gomock.Controller) *ContrailTypeLogicService {
	service := &ContrailTypeLogicService{
		AddressManager: ipammock.NewMockAddressManager(controller),
		DB:             typesmock.NewMockDBServiceInterface(controller),
	}
	service.SetNext(servicesmock.NewMockService(controller))

	return service
}
