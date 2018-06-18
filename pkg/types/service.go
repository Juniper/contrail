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

//DBer makes mocking DB possible in type logic tests
type DBer interface {
	DB() *sql.DB
}

//ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	services.BaseService
	DataService      services.Service
	DBer             DBer
	AddressManager   ipam.AddressManager
	IntPoolAllocator ipam.IntPoolAllocator
}
