package ipam

import (
	"context"

	"github.com/Juniper/contrail/pkg/db"
)

// AddressManager address manager interface for virtual network
type AddressManager interface {
	// TODO: extend this interface with additional methods if necessary.
	//		Most likely following methods are going to be required:
	//			- network create/delete
	AllocateIP(context.Context, *db.AllocateIPRequest) (address string, err error)
	DeallocateIP(context.Context, *db.DeallocateIPRequest) (err error)
	IsIPAllocated(context.Context, *db.IsIPAllocatedRequest) (isAllocated bool, err error)
	CreateIpamSubnet(context.Context, *db.CreateIpamSubnetRequest) (subnetUUID string, err error)
	DeleteIpamSubnet(context.Context, *db.DeleteIpamSubnetRequest) (err error)
}
