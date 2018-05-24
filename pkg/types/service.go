package types

import (
	"database/sql"

	"github.com/Juniper/contrail/pkg/serviceif"
	"golang.org/x/net/context"
)

const (
	//VirtualNetworkIDPoolKey is a key for id pool for virutal network id.
	VirtualNetworkIDPoolKey = "virtual_network_id"
)

//ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	serviceif.BaseService
	DB DBServiceInterface
}

//DBServiceInterface makes mocking DBService possible in type logic tests
type DBServiceInterface interface {
	serviceif.Service
	DB() *sql.DB
	AllocateInt(context.Context, string) (int64, error)
	DeallocateInt(context.Context, string, int64) error
}
