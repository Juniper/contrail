package types

import (
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

const (
	// VirtualNetworkIDPoolKey is a key for id pool for virtual network id.
	VirtualNetworkIDPoolKey = "virtual_network_id"
)

// InTransactionDoer makes transaction mocking possible in type logic tests
type InTransactionDoer interface {
	DoInTransaction(ctx context.Context, do func(context.Context) error) error
}

// FQNameUUIDTranslator translates given fq-name to corresponding uuid and vice versa
type FQNameUUIDTranslator interface {
	TranslateBetweenFQNameUUID(ctx context.Context, uuid string, fqName []string) (*db.MetaData, error)
}

// ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	services.BaseService
	DataService          services.Service
	InTransactionDoer    InTransactionDoer
	AddressManager       ipam.AddressManager
	IntPoolAllocator     ipam.IntPoolAllocator
	FQNameUUIDTranslator FQNameUUIDTranslator
}
