package db

import (
	"context"

	"net"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/graph-gophers/graphql-go/errors"
	"github.com/satori/go.uuid"
)

const (
	SubnetPoolPrefix = "subnet_pool:"
)

// ErrSubnetExhausted signals that address cannot be allocated since subnet is exhausted
type ErrSubnetExhausted interface {
	SubnetExhausted()
}

// AllocateIPRequest arguments for AllocateIP methods.
type AllocateIPRequest struct {
	VirtualNetwork *models.VirtualNetwork
	SubnetUUID     string
	IPAddress      string
}

// DeallocateIPRequest arguments for DeallocateIP methods.
type DeallocateIPRequest struct {
	VirtualNetwork *models.VirtualNetwork
	IPAddress      string
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

// getSubnetUUID returns uuid, if no uuid is set, then allocates a new one. //TODO Probably should change function name
func getSubnetUUID(ipamSubnet *models.IpamSubnetType) (string, error) {
	if ipamSubnet == nil {
		return "", errors.Errorf("Can't get subnetUUID for nil subnet")
	}

	if ipamSubnet.SubnetUUID != "" {
		return ipamSubnet.SubnetUUID, nil
	}
	return uuid.NewV4().String(), nil
}

func (db *Service) clearPools(ctx context.Context, key string) error {
	deleteRequestPool := ipPool{key, net.IP{}, net.IP{}}
	err := db.deleteIPPools(ctx, &deleteRequestPool)
	if err != nil {
		return err
	}
	return nil
}

func subnetUUIDToPoolKey(uuid string) string {
	return SubnetPoolPrefix + uuid
}

func (db *Service) CreateIpamSubnet(ctx context.Context, request *CreateIpamSubnetRequest) (subnetUUID string, err error) {
	subnetUUID, err = getSubnetUUID(request.IpamSubnet)
	if err != nil {
		return "", err
	}
	subnetPoolKey := subnetUUIDToPoolKey(subnetUUID)

	err = db.clearPools(ctx, subnetPoolKey)
	if err != nil {
		return "", err
	}

	for _, pool := range request.IpamSubnet.AllocationPools {
		poolRequest := ipPool{subnetPoolKey, net.ParseIP(pool.Start), net.ParseIP(pool.End)}
		err = db.createIPPool(ctx, &poolRequest)
		if err != nil {
			return "", err
		}
	}

	// We should probably add some logic from vnc_addr_mgmt.py:770, like default dns setting, default gw etc.
	return subnetUUID, err
}

func (db *Service) DeleteIpamSubnet(ctx context.Context, request *DeleteIpamSubnetRequest) (err error) {
	if request == nil || request.SubnetUUID == "" {
		return errors.Errorf("Invalid request in DeleteIpamSubnet")
	}
	subnetPoolKey := subnetUUIDToPoolKey(request.SubnetUUID)

	return db.clearPools(ctx, subnetPoolKey)
}

func (db *Service) allocateIPForKey(ctx context.Context, key string, ipRequested string) (address string, err error) {
	if ipRequested != "" {
		err = db.setIP(ctx, key, net.ParseIP(ipRequested))
		if err != nil {
			return "", err
		}
		return ipRequested, nil
	} else {
		ip, err := db.allocateIP(ctx, key)
		if err != nil {
			return "", err
		}
		return ip.String(), nil
	}
}

// Returns list of pool keys for all the subnets related to the given VN
func getSubnetKeysForVN(vn *models.VirtualNetwork) []string {
	var result []string

	// Take attr subnets
	for _, networkIpam := range vn.NetworkIpamRefs {
		if networkIpam.Attr == nil {
			continue
		}
		for _, subnet := range networkIpam.Attr.IpamSubnets {
			result = append(result, subnetUUIDToPoolKey(subnet.SubnetUUID))
		}
	}

	// TODO do the same with real ipams, not attr.
	return nil
}

func (db *Service) AllocateIP(ctx context.Context, request *AllocateIPRequest) (address string, err error) {
	if request.SubnetUUID != "" {
		key := subnetUUIDToPoolKey(request.SubnetUUID)
		return db.allocateIPForKey(ctx, key, request.IPAddress)
	}

	keys := getSubnetKeysForVN(request.VirtualNetwork)

	// Is it ok? just ignoring errors can potentially be bad
	for _, subnetKey := range keys {
		addr, err := db.allocateIPForKey(ctx, subnetKey, request.IPAddress)
		if addr != "" && err == nil {
			return addr, nil
		}
	}

	return "", errors.Errorf("Could not allocate address %s in any available subnet", request.IPAddress)
}

func (db *Service) DeallocateIP(ctx context.Context, request *DeallocateIPRequest) (err error) {
	//TODO
	return nil
}

func (db *Service) IsIPAllocated(ctx context.Context, request *IsIPAllocatedRequest) (isAllocated bool, err error) {
	//TODO
	return false, nil
}
