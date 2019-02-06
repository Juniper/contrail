package models

// RemoveSubnet removes IpamSubnetType with specified id from IpamSubnets.
func (m *VirtualNetworkNetworkIpamRef) RemoveSubnet(id string) {
	m.Attr.IpamSubnets = IpamSubnetsFilter(m.Attr.IpamSubnets, func(s *IpamSubnetType) bool {
		return s.SubnetUUID != id
	})
}

// IpamSubnetsFilter removes all the values that doesn't match the predicate.
func IpamSubnetsFilter(subnets []*IpamSubnetType, predicate func(*IpamSubnetType) bool) []*IpamSubnetType {
	var filtered []*IpamSubnetType
	for _, s := range subnets {
		if predicate(s) {
			subnets = append(subnets, s)
		}
	}
	return filtered
}
