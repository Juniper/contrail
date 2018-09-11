package models

// GetNetworkIpamRefUUIDs get UUIDs of IPAM references
func (m *VirtualRouter) GetNetworkIpamRefUUIDs() []string {
	var uuids []string
	for _, ref := range m.GetNetworkIpamRefs() {
		uuids = append(uuids, ref.GetUUID())
	}
	return uuids
}

// GetNetworkIpamRefUUIDs get UUIDs of IPAM references
func (m *VirtualRouter) GetRefUUIDToNetworkIpamRefMap() map[string]*VirtualRouterNetworkIpamRef {
	uuidToRefMap := make(map[string]*VirtualRouterNetworkIpamRef)
	for _, ref := range m.GetNetworkIpamRefs() {
		uuidToRefMap[ref.GetUUID()] = ref
	}
	return uuidToRefMap
}

func (ref *VirtualRouterNetworkIpamRef) GetVrouterSpecificAllocationPools() []*AllocationPoolType {
	var vrSpecificPools []*AllocationPoolType
	for _, vrAllocPool := range ref.GetAttr().GetAllocationPools() {
		if !vrAllocPool.GetVrouterSpecificPool() {
			continue
		}
		vrSpecificPools = append(vrSpecificPools, vrAllocPool)
	}

	return vrSpecificPools
}
