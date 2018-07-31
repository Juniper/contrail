package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/models/basemodels"
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

type MetadataGetter interface {
	GetMetaData(ctx context.Context, uuid string, fqName []string) (*basemodels.MetaData, error)
	ListMetadata(ctx context.Context, fqNameUUIDPairs []*basemodels.FQNameUUIDPair) ([]*basemodels.MetaData, error)
}

// ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	services.BaseService
	ReadService       services.ReadService
	InTransactionDoer InTransactionDoer
	AddressManager    ipam.AddressManager
	IntPoolAllocator  ipam.IntPoolAllocator
	WriteService      services.WriteService
	MetadataGetter    MetadataGetter
}
