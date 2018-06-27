package db

import (
	"context"

	"net"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types/ipam"
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/graph-gophers/graphql-go/errors"
	"github.com/satori/go.uuid"
)

const (
	SubnetPoolPrefix = "subnet_pool:"
)

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

func (db *Service) CreateIpamSubnet(ctx context.Context, request *ipam.CreateIpamSubnetRequest) (subnetUUID string, err error) {
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

func (db *Service) DeleteIpamSubnet(ctx context.Context, request *ipam.DeleteIpamSubnetRequest) (err error) {
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
func (db *Service) getSubnetKeysForVN(ctx context.Context, vn *models.VirtualNetwork) ([]string, error) {
	var result []string

	subnets, err := db.getSubnetsForVN(ctx, vn)
	if err != nil {
		return nil, err
	}

	for _, subnet := range subnets {
		result = append(result, subnetUUIDToPoolKey(subnet.SubnetUUID))
	}

	return result, nil
}

// Returns list of subnets related to the given VN
func (db *Service) getSubnetsForVN(ctx context.Context, vn *models.VirtualNetwork) ([]*models.IpamSubnetType, error) {
	var result []*models.IpamSubnetType

	// Take attr subnets
	for _, networkIpam := range vn.NetworkIpamRefs {
		if networkIpam.Attr != nil {
			for _, subnet := range networkIpam.Attr.IpamSubnets {
				if subnet != nil {
					result = append(result, subnet)
				}
			}
		}
		if networkIpam.UUID == "" {
			continue
		}
		req := services.GetNetworkIpamRequest{ID: networkIpam.UUID}
		res, err := db.GetNetworkIpam(ctx, &req)

		if err != nil {
			return nil, err
		}
		if res.NetworkIpam == nil || res.NetworkIpam.IpamSubnets == nil {
			continue
		}
		for _, subnet := range res.NetworkIpam.IpamSubnets.Subnets {
			if subnet != nil {
				result = append(result, subnet)
			}
		}
	}
	return result, nil
}

func (db *Service) AllocateIP(ctx context.Context, request *ipam.AllocateIPRequest) (address string, err error) {
	if request.SubnetUUID != "" {
		key := subnetUUIDToPoolKey(request.SubnetUUID)
		return db.allocateIPForKey(ctx, key, request.IPAddress)
	}

	keys, err := db.getSubnetKeysForVN(ctx, request.VirtualNetwork)
	if err != nil {
		return "", err
	}

	for _, subnetKey := range keys {
		addr, err := db.allocateIPForKey(ctx, subnetKey, request.IPAddress)
		if addr != "" && err == nil {
			return addr, nil
		}
	}

	return "", errors.Errorf("Could not allocate address %s in any available subnet", request.IPAddress)
}

func (db *Service) isIPFromSubnet(ctx context.Context, subnet *models.IpamSubnetType, ip net.IP) (bool, error) {
	if subnet == nil {
		return false, errors.Errorf("Nil subnet in isIpFromSubnet")
	}

	for _, pool := range subnet.AllocationPools {
		startIp := net.ParseIP(pool.Start)
		endIp := cidr.Dec(net.ParseIP(pool.End))

		if ipMin(startIp, ip).Equal(startIp) && ipMax(endIp, ip).Equal(endIp) {
			return true, nil
		}
	}

	return false, nil
}

func (db *Service) DeallocateIP(ctx context.Context, request *ipam.DeallocateIPRequest) (err error) {
	subnets, err := db.getSubnetsForVN(ctx, request.VirtualNetwork)
	if err != nil {
		return err
	}

	for _, subnet := range subnets {
		hit, err := db.isIPFromSubnet(ctx, subnet, net.ParseIP(request.IPAddress))
		if err != nil {
			return err
		}
		if hit {
			err := db.deallocateIP(ctx, subnetUUIDToPoolKey(subnet.SubnetUUID), net.ParseIP(request.IPAddress))
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.Errorf("Could not deallocate address %s from any of available subnets", request.IPAddress)
}

func (db *Service) IsIPAllocated(ctx context.Context, request *ipam.IsIPAllocatedRequest) (isAllocated bool, err error) {
	subnets, err := db.getSubnetsForVN(ctx, request.VirtualNetwork)
	if err != nil {
		return false, err
	}

	for _, subnet := range subnets {
		hit, err := db.isIPFromSubnet(ctx, subnet, net.ParseIP(request.IPAddress))
		if err != nil {
			return false, err
		}
		if hit {
			ip := net.ParseIP(request.IPAddress)
			reqPool := ipPool{subnetUUIDToPoolKey(subnet.SubnetUUID), ip, cidr.Inc(ip)}
			res, err := db.getIPPools(ctx, &reqPool)
			if err != nil {
				return false, err
			}
			return len(res) == 1, nil
		}
	}
	return false, nil
}
