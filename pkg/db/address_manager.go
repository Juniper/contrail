package db

import (
	"context"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

// CreateIpamSubnet creates IPAM subnet.
func (db *Service) CreateIpamSubnet(
	ctx context.Context, request *ipam.CreateIpamSubnetRequest) (subnetUUID string, err error) {
	subnetUUID, err = obtainSubnetUUID(request.IpamSubnet)
	if err != nil {
		return "", err
	}

	// TODO: check allocation pool
	// TODO: check and reserve gw
	// TODO: check and reserve service addr
	// TODO: check and reserve dns nameservers
	// TODO: check allocation units

	err = db.clearPools(ctx, subnetUUID)
	if err != nil {
		return "", err
	}

	for _, pool := range request.IpamSubnet.AllocationPools {
		poolRequest := ipPool{subnetUUID, net.ParseIP(pool.Start), net.ParseIP(pool.End)}
		err = db.createIPPool(ctx, &poolRequest)
		if err != nil {
			return "", err
		}
	}

	return subnetUUID, err
}

// DeleteIpamSubnet deletes IPAM subnet.
func (db *Service) DeleteIpamSubnet(ctx context.Context, request *ipam.DeleteIpamSubnetRequest) (err error) {
	if request == nil || request.SubnetUUID == "" {
		return errors.Errorf("Invalid request in DeleteIpamSubnet")
	}

	return db.clearPools(ctx, request.SubnetUUID)
}

// AllocateIP allocates ip
func (db *Service) AllocateIP(
	ctx context.Context, request *ipam.AllocateIPRequest) (address string, subnetUUID string, err error) {

	// TODO: Implement
	// 		- IPAM based ip allocation from flat subnet where instance ip is directly referring
	//		  IPAM for internal ip address
	//		- Virtual network based ip allocation from flat-subnet ipam
	//		- Virtual network based ip allocation from user-defined connected ipam

	// TODO: Handle allocation methods:
	//		- user-defined-subnet-preferred
	//		- flat-subnet-only
	//		- flat-subnet-preferred

	// This is a simple Virtual network based ip allocation either from user-defined connected ipam

	virtualNetwork := request.VirtualNetwork
	// TODO: Enums strings should be generated from schema

	allocationMethod := getAddressAllocationMode(virtualNetwork)
	if allocationMethod != "user-defined-subnet-only" {
		return "", "", errors.Errorf("Allocation mode %v is not supported", allocationMethod)
	}

	if request.SubnetUUID != "" {
		address, err = db.allocateIPForSubnetUUID(ctx, request.SubnetUUID, request.IPAddress)
		return address, request.SubnetUUID, err
	}

	subnetUUIDs, err := db.getSubnetUUIDsForVN(request.VirtualNetwork)
	if err != nil {
		return "", "", err
	}

	for _, subnetUUID := range subnetUUIDs {
		addr, err := db.allocateIPForSubnetUUID(ctx, subnetUUID, request.IPAddress)
		if addr != "" && err == nil {
			return addr, subnetUUID, nil
		}
	}

	return "", "", errors.Errorf("Could not allocate address %s in any available subnet", request.IPAddress)
}

// DeallocateIP deallocate ip
func (db *Service) DeallocateIP(ctx context.Context, request *ipam.DeallocateIPRequest) (err error) {
	subnets, err := db.getSubnetsForVN(request.VirtualNetwork)
	if err != nil {
		return err
	}

	for _, subnet := range subnets {
		hit, err := db.isIPFromSubnet(subnet, net.ParseIP(request.IPAddress))
		if err != nil {
			return err
		}
		if hit {
			return db.deallocateIP(ctx, subnet.SubnetUUID, net.ParseIP(request.IPAddress))
		}
	}

	return errors.Errorf("Could not deallocate address %s from any of available subnets", request.IPAddress)
}

// IsIPAllocated is ip allocated
func (db *Service) IsIPAllocated(
	ctx context.Context, request *ipam.IsIPAllocatedRequest) (isAllocated bool, err error) {
	subnets, err := db.getSubnetsForVN(request.VirtualNetwork)
	if err != nil {
		return false, err
	}

	for _, subnet := range subnets {
		hit, err := db.isIPFromSubnet(subnet, net.ParseIP(request.IPAddress))
		if err != nil {
			return false, err
		}
		if hit {
			ip := net.ParseIP(request.IPAddress)
			reqPool := ipPool{subnet.SubnetUUID, ip, cidr.Inc(ip)}
			res, err := db.getIPPools(ctx, &reqPool)
			if err != nil {
				return false, err
			}
			return len(res) == 0, nil
		}
	}
	return false, errors.Errorf(
		"Provided ip %v, doesn't belong to any subnet in virtual network %v",
		request.IPAddress, request.VirtualNetwork.GetUUID())
}

func (db *Service) allocateIPForSubnetUUID(
	ctx context.Context, subnetUUID string, ipRequested string) (address string, err error) {

	if ipRequested != "" {
		err = db.setIP(ctx, subnetUUID, net.ParseIP(ipRequested))
		if err != nil {
			return "", err
		}
		return ipRequested, nil
	}

	ip, err := db.allocateIP(ctx, subnetUUID)
	if err != nil {
		return "", err
	}
	return ip.String(), nil

}

func (db *Service) isIPFromSubnet(subnet *models.IpamSubnetType, ip net.IP) (bool, error) {
	if subnet == nil {
		return false, errors.Errorf("Nil subnet in isIpFromSubnet")
	}

	for _, pool := range subnet.AllocationPools {
		startIP := net.ParseIP(pool.Start)
		endIP := cidr.Dec(net.ParseIP(pool.End))

		if ipMin(startIP, ip).Equal(startIP) && ipMax(endIP, ip).Equal(endIP) {
			return true, nil
		}
	}

	return false, nil
}

// Returns list of pool subnetUUIDs for all the subnets related to the given VN
func (db *Service) getSubnetUUIDsForVN(vn *models.VirtualNetwork) ([]string, error) {
	var result []string

	subnets, err := db.getSubnetsForVN(vn)
	if err != nil {
		return nil, err
	}

	for _, subnet := range subnets {
		result = append(result, subnet.SubnetUUID)
	}

	return result, nil
}

// Returns list of subnets related to the given VN
func (db *Service) getSubnetsForVN(vn *models.VirtualNetwork) ([]*models.IpamSubnetType, error) {
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
	}
	return result, nil
}

// obtainSubnetUUID returns uuid, if no uuid is set, then allocates a new one.
func obtainSubnetUUID(ipamSubnet *models.IpamSubnetType) (string, error) {
	if ipamSubnet == nil {
		return "", errors.Errorf("Can't get subnetUUID for nil subnet")
	}

	if ipamSubnet.SubnetUUID != "" {
		return ipamSubnet.SubnetUUID, nil
	}
	return uuid.NewV4().String(), nil
}

func getAddressAllocationMode(virtualNetwork *models.VirtualNetwork) string {
	allocationMethod := "user-defined-subnet-preferred"
	if len(virtualNetwork.GetAddressAllocationMode()) > 0 {
		allocationMethod = virtualNetwork.GetAddressAllocationMode()
	}
	return allocationMethod
}

func (db *Service) clearPools(ctx context.Context, subnetUUID string) error {
	deleteRequestPool := ipPool{subnetUUID, net.IP{}, net.IP{}}
	return db.deleteIPPools(ctx, &deleteRequestPool)
}
