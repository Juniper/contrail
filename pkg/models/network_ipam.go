package models

import (
	"bytes"
	"net"
	"strconv"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/pkg/errors"
	"github.com/twinj/uuid"
)

// IPAM subnet methods.
const (
	UserDefinedSubnet = "user-defined-subnet"
	AutoSubnet        = "auto-subnet"
	FlatSubnet        = "flat-subnet"
)

// IsFlatSubnet checks if this is flat subnet.
func (m *NetworkIpam) IsFlatSubnet() bool {
	return m.IpamSubnetMethod == FlatSubnet
}

// Net returns IPNet object for this subnet.
func (m *SubnetType) Net() (*net.IPNet, error) {
	cidr := m.IPPrefix + "/" + strconv.Itoa(int(m.IPPrefixLen))
	_, n, err := net.ParseCIDR(cidr)
	return n, errors.Wrap(err, "couldn't parse cidr")
}

// IsInSubnet validates allocation pool is in specific subnet.
func (m *AllocationPoolType) IsInSubnet(subnet *net.IPNet) error {
	err := isIPInSubnet(subnet, m.Start)
	if err != nil {
		return common.ErrorBadRequest("allocation pool start " + err.Error())
	}
	err = isIPInSubnet(subnet, m.End)
	if err != nil {
		return common.ErrorBadRequest("allocation pool end " + err.Error())
	}
	return nil
}

// Contains checks if ip address belongs to allocation pool
func (m *AllocationPoolType) Contains(ip net.IP) (bool, error) {
	startIP := net.ParseIP(m.Start)
	if startIP == nil {
		return false, errors.Errorf("couldn't parse start ip address: %v", m.Start)
	}

	endIP := net.ParseIP(m.End)
	if endIP == nil {
		return false, errors.Errorf("couldn't parse end ip address: %v", m.End)
	}

	if bytes.Compare(startIP.To16(), ip.To16()) <= 0 && bytes.Compare(ip.To16(), endIP.To16()) <= 0 {
		return true, nil
	}
	return false, nil
}

func isIPInSubnet(subnet *net.IPNet, ipString string) error {
	ip := net.ParseIP(ipString)
	if ip == nil {
		return errors.Errorf("invalid address: " + ipString)
	}
	if !subnet.Contains(ip) {
		return errors.Errorf("address is out of cidr: " + subnet.String())
	}
	return nil
}

// Validate validates ipam subnet configuration.
func (m *IpamSubnetType) Validate() error {
	if m.SubnetUUID != "" {
		if _, err := uuid.Parse(m.SubnetUUID); err != nil {
			return common.ErrorBadRequestf("invalid subnet uuid: %v", err)
		}
	}

	return m.ValidateSubnetParams()
}

// ValidateSubnetParams validates ipam subnet params.
func (m *IpamSubnetType) ValidateSubnetParams() error {
	subnet, err := m.Subnet.Net()
	if err != nil {
		return common.ErrorBadRequest("invalid subnet")
	}

	for _, allocationPool := range m.AllocationPools {
		err = allocationPool.IsInSubnet(subnet)
		if err != nil {
			return err
		}
	}
	if m.DefaultGateway != "" {
		err = isIPInSubnet(subnet, m.DefaultGateway)
		if err != nil {
			return common.ErrorBadRequest("default gateway " + err.Error())
		}
	}
	if m.DNSServerAddress != "" {
		err = isIPInSubnet(subnet, m.DNSServerAddress)
		if err != nil {
			return common.ErrorBadRequest("DNS server " + err.Error())
		}
	}
	return nil
}

// Contains checks if IpamSubnet's AllocationsPools contain provided ip
func (m *IpamSubnetType) Contains(ip net.IP) (bool, error) {
	subnet, err := m.Subnet.Net()
	if err != nil {
		return false, errors.New("invalid subnet")
	}

	if !subnet.Contains(ip) {
		return false, nil
	}

	if len(m.GetAllocationPools()) == 0 {
		return true, nil
	}

	for _, pool := range m.GetAllocationPools() {
		contains, err := pool.Contains(ip)
		if err != nil {
			return false, err
		}
		if contains {
			return true, nil
		}
	}

	return false, nil
}
