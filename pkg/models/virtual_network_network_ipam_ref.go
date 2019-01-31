package models

// RemoveSubnet removes IpamSubnetType from IpamSubnets with specified id.
func (m *VirtualNetworkNetworkIpamRef) RemoveSubnet(id string) {
	m.Attr.IpamSubnets = IpamSubnets{
		Subnets: m.Attr.IpamSubnets,
	}.Filter(func(s *IpamSubnetType) bool {
		return s.SubnetUUID != id
	}).Subnets
}

// FindSubnet removes IpamSubnetType from IpamSubnets with specified id.
func (m *VirtualNetworkNetworkIpamRef) FindSubnet(id string) *IpamSubnetType {
	return (&(IpamSubnets{
		Subnets: m.Attr.IpamSubnets,
	})).Find(func(s *IpamSubnetType) bool {
		return s.SubnetUUID == id
	})
}
