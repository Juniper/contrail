package db

import (
	"context"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

// CreateIpamSubnet creates IPAM subnet
func (db *Service) CreateIpamSubnet(
	ctx context.Context, request *ipam.CreateIpamSubnetRequest,
) (subnetUUID string, err error) {
	subnetUUID, err = obtainSubnetUUID(request.IpamSubnet)
	if err != nil {
		return "", err
	}

	// TODO: check allocation pool
	// TODO: check and reserve gw
	// TODO: check and reserve service addr
	// TODO: check and reserve dns nameservers
	// TODO: check allocation units

	err = db.deleteIPPools(ctx, &ipPool{
		key: subnetUUID,
	})
	if err != nil {
		return "", err
	}

	for _, pool := range request.IpamSubnet.GetAllocationPools() {
		poolRequest := ipPool{
			key:   subnetUUID,
			start: net.ParseIP(pool.Start),
			end:   net.ParseIP(pool.End),
		}

		err = db.createIPPool(ctx, &poolRequest)
		if err != nil {
			return "", err
		}
	}

	return subnetUUID, err
}

// DeleteIpamSubnet deletes IPAM subnet
func (db *Service) DeleteIpamSubnet(ctx context.Context, request *ipam.DeleteIpamSubnetRequest) (err error) {
	if request.SubnetUUID == "" {
		return errors.Errorf("empty subnet uuid in DeleteIpamSubnet")
	}

	return db.deleteIPPools(ctx, &ipPool{
		key: request.SubnetUUID,
	})
}

// AllocateIP allocates ip
func (db *Service) AllocateIP(
	ctx context.Context, request *ipam.AllocateIPRequest) (address string, subnetUUID string, err error) {
	// TODO: Implement:
	//		- IPAM based ip allocation from flat subnet where instance ip is directly referring
	//		  IPAM for internal ip address
	//		- IPAM based ip allocation from flat subnet where instance ip is referring to vrouter which has
	// 		  allocation pools on vrouter->ipam link.
	//		- Virtual network based ip allocation from flat-subnet ipam
	//		- Virtual network based ip allocation from user-defined connected ipam

	// TODO: Handle allocation methods:
	//		- user-defined-subnet-preferred
	//		- flat-subnet-only
	//		- flat-subnet-preferred

	// This is a simple Virtual network based ip allocation from user-defined subnet

	virtualNetwork := request.VirtualNetwork
	allocationMethod := virtualNetwork.GetAddressAllocationMethod()
	if allocationMethod != "user-defined-subnet-only" {
		return "", "", errors.Errorf("allocation mode %v is not supported", allocationMethod)
	}

	subnetUUIDs := virtualNetwork.GetSubnetUUIDs()
	if request.SubnetUUID != "" {
		if !common.ContainsString(subnetUUIDs, request.SubnetUUID) {
			return "", "", errors.Errorf("could not find subnet %s in in virtual network %v", request.SubnetUUID,
				virtualNetwork.GetUUID())
		}
		subnetUUIDs = append([]string(nil), request.SubnetUUID)
	}

	for _, subnetUUID := range subnetUUIDs {
		addr, err := db.allocateIPForSubnetUUID(ctx, subnetUUID, request.IPAddress)
		if addr != "" && err == nil {
			return addr, subnetUUID, nil
		}
	}

	return "", "", errors.Errorf("could not allocate address %s in any available subnets in virtual network %v",
		request.IPAddress, virtualNetwork.GetUUID())
}

// DeallocateIP deallocates ip
func (db *Service) DeallocateIP(ctx context.Context, request *ipam.DeallocateIPRequest) (err error) {
	for _, subnet := range request.VirtualNetwork.GetSubnets() {
		hit, err := subnet.Contains(net.ParseIP(request.IPAddress))
		if err != nil {
			return err
		}
		if hit {
			return db.deallocateIP(ctx, subnet.SubnetUUID, net.ParseIP(request.IPAddress))
		}
	}

	return errors.Errorf("could not deallocate address %s from any of available subnets in virtual network %v",
		request.IPAddress, request.VirtualNetwork.GetUUID())
}

// IsIPAllocated checks if ip is allocated
func (db *Service) IsIPAllocated(
	ctx context.Context, request *ipam.IsIPAllocatedRequest,
) (isAllocated bool, err error) {
	for _, subnet := range request.VirtualNetwork.GetSubnets() {
		hit, err := subnet.Contains(net.ParseIP(request.IPAddress))
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
		"provided ip %v, doesn't belong to any subnet in virtual network %v",
		request.IPAddress, request.VirtualNetwork.GetUUID())
}

func (db *Service) allocateIPForSubnetUUID(
	ctx context.Context, subnetUUID string, ipRequested string,
) (address string, err error) {
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

// obtainSubnetUUID returns uuid, if no uuid is set, then allocates a new one.
func obtainSubnetUUID(ipamSubnet *models.IpamSubnetType) (string, error) {
	if ipamSubnet == nil {
		return "", errors.Errorf("can't get subnetUUID for nil subnet")
	}

	if ipamSubnet.SubnetUUID != "" {
		return ipamSubnet.SubnetUUID, nil
	}
	return uuid.NewV4().String(), nil
}
