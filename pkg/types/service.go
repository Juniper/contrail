package types

import (
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

const (
	//VirtualNetworkIDPoolKey is a key for id pool for virtual network id.
	VirtualNetworkIDPoolKey = "virtual_network_id"
)

//DBService makes mocking DB possible in type logic tests
type DBService interface {
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
}

//ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	services.BaseService
	DataService      services.Service
	DBService        DBService
	AddressManager   ipam.AddressManager
	IntPoolAllocator ipam.IntPoolAllocator
}
