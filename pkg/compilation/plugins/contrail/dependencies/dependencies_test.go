package dependencies_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/compilation/logic"
	"github.com/Juniper/contrail/pkg/compilation/plugins/contrail/dependencies"
	"github.com/Juniper/contrail/pkg/models"
)

func TestReturnsRefs(t *testing.T) {
	cache := intent.NewCache()

	// Create VirtualNetworks in the Cache
	Vn1 := models.VirtualNetwork{
		UUID: "Virtual-Network-1",
	}
	Vn2 := models.VirtualNetwork{
		UUID: "Virtual-Network-2",
	}

	// Create NetworkPolicys in the Cache
	Np1 := models.NetworkPolicy{
		UUID: "Network-policy-1",
	}
	Np2 := models.NetworkPolicy{
		UUID: "Network-policy-2",
	}

	Np1Ref := models.VirtualNetworkNetworkPolicyRef{}
	Np1Ref.UUID = Np1.UUID
	Np1Ref.To = []string{"default-domain", "default-project", "Network-Policy-1"}
	Vn1.NetworkPolicyRefs = append(Vn1.NetworkPolicyRefs, &Np1Ref)

	Np2Ref := models.VirtualNetworkNetworkPolicyRef{}
	Np2Ref.UUID = Np2.UUID
	Np2Ref.To = []string{"default-domain", "default-project", "Network-Policy-2"}
	Vn1.NetworkPolicyRefs = append(Vn1.NetworkPolicyRefs, &Np2Ref)

	Vn2.NetworkPolicyRefs = append(Vn2.NetworkPolicyRefs, &Np2Ref)

	Np1.VirtualNetworkBackRefs = append(Np1.VirtualNetworkBackRefs, &Vn1)
	Np2.VirtualNetworkBackRefs = append(Np1.VirtualNetworkBackRefs, &Vn1)

	Np2.VirtualNetworkBackRefs = append(Np1.VirtualNetworkBackRefs, &Vn2)

	storeTestVirtualNetworkIntent(cache, &Vn1)
	storeTestVirtualNetworkIntent(cache, &Vn2)

	storeTestNetworkPolicyIntent(cache, &Np1)
	storeTestNetworkPolicyIntent(cache, &Np2)

	d := dependencies.NewDependencyProcessor(cache)
	if d != nil {
		d.Evaluate(Vn1, "VirtualNetwork", "Self")
		resources := d.GetResources()

		networks := mustLoad(t, resources, "VirtualNetwork")
		mustLoad(t, networks, Vn1.UUID)
		mustLoad(t, networks, Vn2.UUID)

		policies := mustLoad(t, resources, "NetworkPolicy")
		mustLoad(t, policies, Np1.UUID)
		mustLoad(t, policies, Np2.UUID)
	}
}

func TestReturnsSelf(t *testing.T) {
	cache := intent.NewCache()
	vn := &models.VirtualNetwork{
		UUID: "Virtual-Network-1",
	}
	vnIntent := storeTestVirtualNetworkIntent(cache, vn)

	d := dependencies.NewDependencyProcessor(cache)
	if d != nil {
		d.Evaluate(vnIntent, "VirtualNetwork", "Self")
		resources := d.GetResources()

		networks := mustLoad(t, resources, "VirtualNetwork")
		mustLoad(t, networks, vnIntent.GetUUID())
	}
}

func mustLoad(t *testing.T, rawSyncMap interface{}, key string) interface{} {
	syncMap, ok := rawSyncMap.(*sync.Map)
	assert.Truef(t, ok, "%v should be a sync.Map", rawSyncMap)

	value, ok := syncMap.Load(key)
	assert.Truef(t, ok, "The map should have a '%s' key", key)
	return value
}

func storeTestVirtualNetworkIntent(
	cache *intent.Cache,
	vn *models.VirtualNetwork,
) intent.Intent {
	i := &logic.VirtualNetworkIntent{
		VirtualNetwork: vn,
	}
	cache.Store(i)
	return i
}

func storeTestNetworkPolicyIntent(
	cache *intent.Cache,
	np *models.NetworkPolicy,
) intent.Intent {
	i := &logic.NetworkPolicyIntent{
		NetworkPolicy: np,
	}
	cache.Store(i)
	return i
}
