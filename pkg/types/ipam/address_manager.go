package ipam

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
)

// AddressManagerErrorCode address management error code.
type AddressManagerErrorCode int

const (
	// SubnetExhausted indicates that the subnet is exhausted.
	SubnetExhausted = AddressManagerErrorCode(iota)
)

// AddressManagerError IPAM specific error.
type AddressManagerError interface {
	GetAddressManagementErrorCode() AddressManagerErrorCode
}

// AllocateIPRequest arguments for AllocateIP methods.
type AllocateIPRequest struct {
	VirtualNetwork *models.VirtualNetwork
	SubnetUUID     string
	IPAddress      string
	AllocatorID    string
}

// DeallocateIPRequest arguments for DeallocateIP methods.
type DeallocateIPRequest struct {
	VirtualNetwork *models.VirtualNetwork
	IPAddress      string
	AllocatorID    string
}

// IsIPAllocatedRequest arguments for IsIPAllocated methods.
type IsIPAllocatedRequest struct {
	VirtualNetwork *models.VirtualNetwork
	IPAddress      string
}

// AddressManager address manager interface for virtual network
type AddressManager interface {
	// TODO: extend this interface with additional methods if necessary.
	//		Most likely following methods are going to be required:
	//			- network create/delete
	//			- subnet create/delete
	AllocateIP(context.Context, *AllocateIPRequest) (address string, subnetUUID string, err error)
	DeallocateIP(context.Context, *DeallocateIPRequest) (err error)
	IsIPAllocated(context.Context, *IsIPAllocatedRequest) (isAllocated bool, err error)
}
