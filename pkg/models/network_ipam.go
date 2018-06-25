package models

import (
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
	return n, err
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
	_, err := uuid.Parse(m.SubnetUUID)
	if err != nil {
		return common.ErrorBadRequest("invalid subnet uuid")
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
