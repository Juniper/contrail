package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

const (
	// VirtualNetworkIDPoolKey is a key for id pool for virtual network id.
	VirtualNetworkIDPoolKey = "virtual_network_id"
)

// InTransactionDoer executes do function atomically.
type InTransactionDoer interface {
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
}

// MetaDataGetter gets metadata from database
type MetaDataGetter interface {
	GetMetaData(ctx context.Context, uuid string, fqName []string) (*models.MetaData, error)
}

// ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	services.BaseService
	ReadService       services.ReadService
	InTransactionDoer InTransactionDoer
	AddressManager    ipam.AddressManager
	IntPoolAllocator  ipam.IntPoolAllocator
	MetaDataGetter    MetaDataGetter
	WriteService      services.WriteService
}
