package types

import (
	"database/sql"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

const (
	//VirtualNetworkIDPoolKey is a key for id pool for virtual network id.
	VirtualNetworkIDPoolKey = "virtual_network_id"
)

//DBInterface makes mocking DB possible in type logic tests
type DBInterface interface {
	DB() *sql.DB
}

//ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	services.BaseService
	DBService        services.Service
	DB               DBInterface
	AddressManager   ipam.AddressManager
	IntPoolAllocator ipam.IntPoolAllocator
}
