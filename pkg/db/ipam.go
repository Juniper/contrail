package db

import (
	"context"

	"github.com/Juniper/contrail/pkg/types/ipam"
)

// AllocateIP allocates ip address.
func (db *Service) AllocateIP(ctx context.Context, request *ipam.AllocateIPRequest) (address string, subnetUUID string, err error) {
	// TODO: Implement IPAM
	return request.IPAddress, "", nil
}

// DeallocateIP deallocates ip address.
func (db *Service) DeallocateIP(ctx context.Context, request *ipam.DeallocateIPRequest) (err error) {
	// TODO: Implement IPAM
	return nil
}

// IsIPAllocated checks if ip is already allocated.
func (db *Service) IsIPAllocated(ctx context.Context, request *ipam.IsIPAllocatedRequest) (isAllocated bool, err error) {
	// TODO: Implement IPAM
	return false, nil
}
