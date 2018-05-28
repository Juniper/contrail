package ipam

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
)

type AddressManagerErrorCode int

const (
	SubnetExhausted = AddressManagerErrorCode(iota)
)

type AddressManagerError interface {
	GetAddressManagementErrorCode() AddressManagerErrorCode
}

type AllocateIPParams struct {
	VirtualNetwork *models.VirtualNetwork
	SubnetUUID     string
	IPAddress      string
	AllocatorID    string
}

type DeallocateIPParams struct {
	VirtualNetwork *models.VirtualNetwork
	IPAddress      string
	AllocatorID    string
}

type IsIPAllocatedParams struct {
	VirtualNetwork *models.VirtualNetwork
	IPAddress      string
}

type AddressManager interface {
	AllocateIP(context.Context, *AllocateIPParams) (address string, subnetUUID string, err error)
	DeallocateIP(context.Context, *DeallocateIPParams) (err error)
	IsIPAllocated(context.Context, *IsIPAllocatedParams) (isAllocated bool, err error)
}
