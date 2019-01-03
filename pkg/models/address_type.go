package models

import (
	"net"
	"strconv"

	"github.com/Juniper/contrail/pkg/errutil"
)

// Non-reference security group names.
const (
	AnySecurityGroup         = "any"
	LocalSecurityGroup       = "local"
	UnspecifiedSecurityGroup = ""
)

// AllIPv4Addresses returns an AddressType with a subnet of all possible IPv4 addresses.
func AllIPv4Addresses() *AddressType {
	return &AddressType{
		Subnet: &SubnetType{
			IPPrefix:    "0.0.0.0",
			IPPrefixLen: 0,
		},
	}
}

// AllIPv6Addresses returns an AddressType with a subnet of all possible IPv6 addresses.
func AllIPv6Addresses() *AddressType {
	return &AddressType{
		Subnet: &SubnetType{
			IPPrefix:    "::",
			IPPrefixLen: 0,
		},
	}
}

func (m *AddressType) VerifyEthertypeWithSubnet(ethertype string) error {
	if m.GetSubnet() != nil {
		IPPrefix := m.GetSubnet().GetIPPrefix()
		IPPrefixLen := strconv.Itoa(int(m.GetSubnet().GetIPPrefixLen()))
		version, err := resolveIPVersionFromCIDR(IPPrefix, IPPrefixLen)
		if err != nil {
			return err
		}
		if ethertype != version {
			return errutil.ErrorBadRequestf("Rule subnet %v doesn't match ethertype %v",
				IPPrefix+"/"+IPPrefixLen, ethertype)
		}
	}
	return nil
}

func resolveIPVersionFromCIDR(IPPrefix, IPPrefixLen string) (string, error) {
	network, _, err := net.ParseCIDR(IPPrefix + "/" + IPPrefixLen)
	if err != nil {
		return "", errutil.ErrorBadRequestf("Cannot parse address %v/%v. %v.",
			IPPrefix, IPPrefixLen, err)
	}
	switch {
	case network.To4() != nil:
		return "IPv4", nil
	case network.To16() != nil:
		return "IPv6", nil
	default:
		return "", errutil.ErrorBadRequestf("Cannot resolve ip version %v/%v.",
			IPPrefix, IPPrefixLen)
	}
}

// IsSecurityGroupNameAReference checks if the Security Group name in an address
// is a reference to other security group.
func (m *AddressType) IsSecurityGroupNameAReference() bool {
	return m.SecurityGroup != AnySecurityGroup &&
		m.SecurityGroup != LocalSecurityGroup && m.SecurityGroup != UnspecifiedSecurityGroup
}