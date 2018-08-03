package models

// Returns true if routing instance FQName fits IP Fabric
func (ri *RoutingInstance) IsIPFabric() bool {
	fq := []string{"default-domain", "default-project", "ip-fabric", "__default__"}
	return FQNameEquals(fq, ri.GetFQName())
}

// Returns true if routing instance FQName fits Link Local
func (ri *RoutingInstance) IsLinkLocal() bool {
	fq := []string{"default-domain", "default-project", "__link_local__", "__link_local__"}
	return FQNameEquals(fq, ri.GetFQName())
}
