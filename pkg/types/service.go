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

// FQNameUUIDTranslator translates given fq-name to corresponding uuid and vice versa
type FQNameUUIDTranslator interface {
	TranslateBetweenFQNameUUID(ctx context.Context, uuid string, fqName []string) (*models.MetaData, error)
}

// ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	services.BaseService
	ReadService          services.ReadService
	InTransactionDoer    InTransactionDoer
	AddressManager       ipam.AddressManager
	IntPoolAllocator     ipam.IntPoolAllocator
	FQNameUUIDTranslator FQNameUUIDTranslator
	WriteService         services.WriteService
}
