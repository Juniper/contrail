package models

import (
	"net"
	"strconv"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/pkg/errors"
	"github.com/twinj/uuid"
)

const (
	flatSubnet = "flat-subnet"
)

// IsFlatSubnet checks if this is flat subnet.
func (m *NetworkIpam) IsFlatSubnet() bool {
	return m.IpamSubnetMethod == flatSubnet
}

// Net returns IPNet object for this subnet.
func (m *SubnetType) Net() (*net.IPNet, error) {
	cidr := m.IPPrefix + "/" + strconv.Itoa(int(m.IPPrefixLen))
	_, n, err := net.ParseCIDR(cidr)
	return n, err
}

// IsInSubnet validates allocation pool is in specific subnet.
func (m *AllocationPoolType) IsInSubnet(subnet *net.IPNet) error {
	err := isIpInSubnet(subnet, m.Start)
	if err != nil {
		return common.ErrorBadRequest("allocation pool start %v" + err.Error())
	}
	err = isIpInSubnet(subnet, m.End)
	if err != nil {
		return common.ErrorBadRequest("allocation pool end %v" + err.Error())
	}
	return nil
}

func parseIpfromString(ipString string) (net.IP, error) {
	ip := net.ParseIP(ipString)
	if ip == nil {
		return nil, errors.Errorf("invalid address:" + ipString)
	}
	return ip, nil
}

func isIpInSubnet(subnet *net.IPNet, ipString string) error {
	ip, err := parseIpfromString(ipString)
	if err != nil {
		return err
	}
	if !subnet.Contains(ip) {
		return errors.Errorf("address is out of cidr:" + subnet.String())
	}
	return nil
}

// IsValid validates ipam subnet configuraion.
func (m *IpamSubnetType) IsValid() error {
	_, err := uuid.Parse(m.SubnetUUID)
	if err != nil {
		return common.ErrorBadRequest("invalid subnet uuid.")
	}

	return m.CheckIfSubnetParamsAreValid()
}

// CheckIfSubnetParamsAreValid validates ipam subnet params.
func (m *IpamSubnetType) CheckIfSubnetParamsAreValid() error {
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
		err = isIpInSubnet(subnet, m.DefaultGateway)
		if err != nil {
			return common.ErrorBadRequest("default gateway " + err.Error())
		}
	}
	if m.DNSServerAddress != "" {
		err = isIpInSubnet(subnet, m.DNSServerAddress)
		if err != nil {
			return common.ErrorBadRequest("DNS server " + err.Error())
		}
	}
	return nil
}

// ValidateFlatSubnet validates this VnSubnet can be used for flat subnet.
func (m *VnSubnetsType) ValidateFlatSubnet() error {
	for _, ipamSubnet := range m.GetIpamSubnets() {
		if ipamSubnet.Subnet.IPPrefix != "" {
			return common.ErrorBadRequest("with flat-subnet, network can not have user-defined subnet")
		}
	}
	return nil
}

// ValidateUserDefined validates user defined subnet.
func (m *VnSubnetsType) ValidateUserDefined() error {
	for _, ipamSubnet := range m.GetIpamSubnets() {
		// check network subnet
		err := ipamSubnet.IsValid()
		if err != nil {
			return err
		}
	}
	return nil
}
