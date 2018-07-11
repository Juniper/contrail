package dependencies

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// TestDependencies - test
func TestDependencies(t *testing.T) {
	cache := &sync.Map{}

	Vn1 := &models.VirtualNetwork{}
	Vn1.UUID = "Virtual-Network-1"
	putIntoCache(t, cache, "VirtualNetwork", Vn1)

	Vn2 := &models.VirtualNetwork{}
	Vn2.UUID = "Virtual-Network-2"
	putIntoCache(t, cache, "VirtualNetwork", Vn2)

	Np1 := &models.NetworkPolicy{}
	Np1.UUID = "Network-policy-1"
	putIntoCache(t, cache, "NetworkPolicy", Np1)

	Np2 := &models.NetworkPolicy{}
	Np2.UUID = "Network-policy-2"
	putIntoCache(t, cache, "NetworkPolicy", Np2)

	Np1Ref := &models.VirtualNetworkNetworkPolicyRef{}
	Np1Ref.UUID = Np1.UUID
	Np1Ref.To = []string{"default-domain", "default-project", "Network-Policy-1"}
	Vn1.NetworkPolicyRefs = append(Vn1.NetworkPolicyRefs, Np1Ref)

	Np2Ref := &models.VirtualNetworkNetworkPolicyRef{}
	Np2Ref.UUID = Np2.UUID
	Np2Ref.To = []string{"default-domain", "default-project", "Network-Policy-2"}
	Vn1.NetworkPolicyRefs = append(Vn1.NetworkPolicyRefs, Np2Ref)

	Vn2.NetworkPolicyRefs = append(Vn2.NetworkPolicyRefs, Np2Ref)

	Np1.VirtualNetworkBackRefs = append(Np1.VirtualNetworkBackRefs, Vn1)
	Np2.VirtualNetworkBackRefs = append(Np1.VirtualNetworkBackRefs, Vn1)

	Np2.VirtualNetworkBackRefs = append(Np1.VirtualNetworkBackRefs, Vn2)

	d := NewDependencyProcessor(cache)
	d.Evaluate(Vn1, "VirtualNetwork", "Self")
	resources := d.GetResources()

	networks := mustLoad(t, resources, "VirtualNetwork")
	mustLoad(t, networks, Vn1.UUID)
	mustLoad(t, networks, Vn2.UUID)

	policies := mustLoad(t, resources, "NetworkPolicy")
	mustLoad(t, policies, Np1.UUID)
	mustLoad(t, policies, Np2.UUID)
}

func putIntoCache(t *testing.T, cache *sync.Map, kind string, resource services.Resource) {
	resourceMap := ensureSubmap(t, cache, kind)
	resourceMap.Store(resource.GetUUID(), resource)
}

func ensureSubmap(t *testing.T, cache *sync.Map, key string) (submap *sync.Map) {
	rawSubmap, ok := cache.Load(key)
	if !ok {
		submap = &sync.Map{}
		cache.Store(key, submap)
		return
	}
	submap, ok = rawSubmap.(*sync.Map)
	require.True(t, ok)
	return
}

func mustLoad(t *testing.T, rawSyncMap interface{}, key string) interface{} {
	syncMap, ok := rawSyncMap.(*sync.Map)
	assert.Truef(t, ok, "%v should be a sync.Map", rawSyncMap)

	value, ok := syncMap.Load(key)
	assert.Truef(t, ok, "The map should have a '%s' key", key)
	return value
}
