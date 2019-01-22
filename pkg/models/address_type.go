package models

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

// IsSecurityGroupNameAReference checks if the Security Group name in an address
// is a reference to other security group.
func (m *AddressType) IsSecurityGroupNameAReference() bool {
	return m.SecurityGroup != AnySecurityGroup &&
		m.SecurityGroup != LocalSecurityGroup && m.SecurityGroup != UnspecifiedSecurityGroup
}

func (m *AddressType) isSecurityGroupLocal() bool {
	return m.SecurityGroup == LocalSecurityGroup
}
