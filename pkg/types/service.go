package types

import (
	"database/sql"

	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"golang.org/x/net/context"
)

const (
	//VirtualNetworkIDPoolKey is a key for id pool for virtual network id.
	VirtualNetworkIDPoolKey = "virtual_network_id"
)

//DBServiceInterface makes mocking DBService possible in type logic tests
type DBServiceInterface interface {
	serviceif.Service
	DB() *sql.DB
	AllocateInt(context.Context, string) (int64, error)
	DeallocateInt(context.Context, string, int64) error
}

//ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	serviceif.BaseService
	DB             DBServiceInterface
	AddressManager ipam.AddressManager
}
