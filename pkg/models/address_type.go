package models

const (
	// Non-reference security group names.
	AnySecurityGroup         = "any"
	LocalSecurityGroup       = "local"
	UnspecifiedSecurityGroup = ""

	IPv4ZeroValue = "0.0.0.0"
	IPv6ZeroValue = "::"
)

// AllIPv4Addresses returns an AddressType with a subnet of all possible IPv4 addresses.
func AllIPv4Addresses() *AddressType {
	return &AddressType{
		Subnet: &SubnetType{
			IPPrefix:    IPv4ZeroValue,
			IPPrefixLen: 0,
		},
	}
}

// AllIPv6Addresses returns an AddressType with a subnet of all possible IPv6 addresses.
func AllIPv6Addresses() *AddressType {
	return &AddressType{
		Subnet: &SubnetType{
			IPPrefix:    IPv6ZeroValue,
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
