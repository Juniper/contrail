package models

import (
	"net"
	"strconv"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/twinj/uuid"
)

const (
	flatSubnet = "flat-subnet"
)

//IsFlatSubnet checks if this is flat subnet.
func (m *NetworkIpam) IsFlatSubnet() bool {
	return m.IpamSubnetMethod == flatSubnet
}

//Net returns IPNet object for this subnet.
func (m *SubnetType) Net() (*net.IPNet, error) {
	cidr := m.IPPrefix + "/" + strconv.Itoa(int(m.IPPrefixLen))
	_, n, err := net.ParseCIDR(cidr)
	return n, err
}

//Validate validates allocation pool is in specific subnet.
func (m *AllocationPoolType) Validate(subnet *net.IPNet) error {
	err := validateIPinSubnet(subnet, "allocation pool start", m.Start)
	if err != nil {
		return err
	}
	err = validateIPinSubnet(subnet, "allocation pool end", m.End)
	if err != nil {
		return err
	}
	return nil
}

func validateIPinSubnet(subnet *net.IPNet, name, ipString string) error {
	ip := net.ParseIP(ipString)
	if ip == nil {
		return common.ErrorBadRequest("Invalid " + name + " address:" + ipString)
	}
	if !subnet.Contains(ip) {
		return common.ErrorBadRequest(name + " address is out of cidr:" + subnet.String())
	}
	return nil
}

//Validate validates ipam subnet configuraion.
func (m *IpamSubnetType) Validate() error {
	subnet, err := m.Subnet.Net()
	if err != nil {
		return common.ErrorBadRequest("Invalid subnet")
	}

	//TODO: add validation on subnet uuid on schema level.
	_, err = uuid.Parse(m.SubnetUUID)
	if err != nil {
		return common.ErrorBadRequest("Invalid subnet uuid.")
	}

	for _, allocationPool := range m.AllocationPools {
		err = allocationPool.Validate(subnet)
		if err != nil {
			return err
		}
	}
	if m.DefaultGateway != "" {
		err := validateIPinSubnet(subnet, "gateway Ip", m.DefaultGateway)
		if err != nil {
			return err
		}
	}
	if m.DNSServerAddress != "" {
		err := validateIPinSubnet(subnet, "Dns Server Ip", m.DNSServerAddress)
		if err != nil {
			return err
		}
	}
	return nil
}

//ValidateFlatSubnet validates this VnSubnet can be used for flat subnet.
func (m *VnSubnetsType) ValidateFlatSubnet() error {
	for _, ipamSubnet := range m.GetIpamSubnets() {
		if ipamSubnet.Subnet.IPPrefix != "" {
			return common.ErrorBadRequest("with flat-subnet, network can not have user-defined subnet")
		}
	}
	return nil
}

//ValidateUserDefined validates user defined subnet.
func (m *VnSubnetsType) ValidateUserDefined() error {
	for _, ipamSubnet := range m.GetIpamSubnets() {
		// check network subnet
		err := ipamSubnet.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
