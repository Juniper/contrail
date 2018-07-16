package ipam

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
)

// ErrSubnetExhausted signals that address cannot be allocated since subnet is exhausted
type ErrSubnetExhausted interface {
	SubnetExhausted()
}

// AllocateIPRequest arguments for AllocateIP methods.
type AllocateIPRequest struct {
	VirtualNetwork  *models.VirtualNetwork
	SubnetUUID      string
	IPAddress       string
	IPFamily        string
	IpamRefs        []*models.VirtualRouterNetworkIpamRef
	AllocationPools []*models.AllocationPoolType
}

// DeallocateIPRequest arguments for DeallocateIP methods.
type DeallocateIPRequest struct {
	VirtualNetwork *models.VirtualNetwork
	IPAddress      string
	IpamRefs       []*models.InstanceIPNetworkIpamRef
}

// IsIPAllocatedRequest arguments for IsIPAllocated methods.
type IsIPAllocatedRequest struct {
	VirtualNetwork *models.VirtualNetwork
	IPAddress      string
}

// CreateIpamSubnetRequest arguments for CreateIpamSubnet methods.
type CreateIpamSubnetRequest struct {
	IpamSubnet      *models.IpamSubnetType
	NetworkIpamUUID string
}

// DeleteIpamSubnetRequest arguments for DeleteIpamSubnet methods.
type DeleteIpamSubnetRequest struct {
	SubnetUUID string
}

// AddressManager address manager interface for virtual network
type AddressManager interface {
	// TODO: extend this interface with additional methods if necessary.
	//		Most likely following methods are going to be required:
	//			- network create/delete
	AllocateIP(context.Context, *AllocateIPRequest) (address string, subnetUUID string, err error)
	DeallocateIP(context.Context, *DeallocateIPRequest) (err error)
	IsIPAllocated(context.Context, *IsIPAllocatedRequest) (isAllocated bool, err error)
	CreateIpamSubnet(context.Context, *CreateIpamSubnetRequest) (subnetUUID string, err error)
	DeleteIpamSubnet(context.Context, *DeleteIpamSubnetRequest) (err error)
}
