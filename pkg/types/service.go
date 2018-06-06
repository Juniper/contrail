package types

import (
	"database/sql"

	"github.com/Juniper/contrail/pkg/serviceif"
	"golang.org/x/net/context"
	"github.com/Juniper/contrail/pkg/models"
)

const (
	//VirtualNetworkIDPoolKey is a key for id pool for virutal network id.
	VirtualNetworkIDPoolKey = "virtual_network_id"
)

//DBServiceInterface makes mocking DBService possible in type logic tests
type DBServiceInterface interface {
	serviceif.Service
	DB() *sql.DB
	AllocateInt(context.Context, string) (int64, error)
	DeallocateInt(context.Context, string, int64) error
	GetLinkedBGPVPNviaRouter(ctx context.Context, virtualNetwork *models.VirtualNetwork) ([]string, error)
}

//ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	serviceif.BaseService
	DB             DBServiceInterface
}
