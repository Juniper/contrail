package types

import (
	"context"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

const (
	// VirtualNetworkIDPoolKey identifies the int pool of virtual network IDs.
	VirtualNetworkIDPoolKey = "virtual_network_id"
	// SecurityGroupIDPoolKey identifies the int pool of security group IDs.
	SecurityGroupIDPoolKey = "security_group_id"
)

// InTransactionDoer executes do function atomically.
type InTransactionDoer interface {
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
}

// IntPoolAllocator (de)allocates integers in an integer pool.
type IntPoolAllocator interface {
	CreateIntPool(context.Context, string, int64, int64) error
	DeleteIntPool(context.Context, string) error
	AllocateInt(context.Context, string) (int64, error)
	SetInt(context.Context, string, int64) error
	DeallocateInt(context.Context, string, int64) error
}

//MetadataGetter provides getter for metadata.
type MetadataGetter interface {
	GetMetadata(ctx context.Context, requested basemodels.Metadata) (*basemodels.Metadata, error)
	ListMetadata(ctx context.Context, requested []*basemodels.Metadata) ([]*basemodels.Metadata, error)
}

// ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	services.BaseService
	ReadService       services.ReadService
	InTransactionDoer InTransactionDoer
	AddressManager    ipam.AddressManager
	IntPoolAllocator  IntPoolAllocator
	MetadataGetter    baseservices.MetadataGetter
	WriteService      services.WriteService
}
