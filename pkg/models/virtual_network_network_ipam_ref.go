package models

// RemoveSubnet removes IpamSubnetType from IpamSubnets with specified id.
func (m *VirtualNetworkNetworkIpamRef) RemoveSubnet(id string) {
	m.Attr.IpamSubnets = IpamSubnets{
		Subnets: m.Attr.IpamSubnets,
	}.Filter(func(s *IpamSubnetType) bool {
		return s.SubnetUUID != id
	}).Subnets
}
