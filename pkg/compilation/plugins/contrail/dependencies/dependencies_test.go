package dependencies

import (
	"sync"
	"testing"
	"sync"

	"github.com/Juniper/contrail/pkg/models"

	"github.com/stretchr/testify/assert"
)

func TestReturnsRefs(t *testing.T) {
	ObjsCache := &sync.Map{}

	// Create VirtualNetworks in the Cache
	VnObjMap := &sync.Map{}
	Vn1 := models.VirtualNetwork{}
	Vn1.UUID = "Virtual-Network-1"
	VnObjMap.Store(Vn1.UUID, &Vn1)
	Vn2 := models.VirtualNetwork{}
	Vn2.UUID = "Virtual-Network-2"
	VnObjMap.Store(Vn2.UUID, &Vn2)
	ObjsCache.Store("VirtualNetwork", VnObjMap)

	// Create NetworkPolicys in the Cache
	NpObjMap := &sync.Map{}
	Np1 := models.NetworkPolicy{}
	Np1.UUID = "Network-policy-1"
	NpObjMap.Store(Np1.UUID, &Np1)
	Np2 := models.NetworkPolicy{}
	Np2.UUID = "Network-policy-2"
	NpObjMap.Store(Np2.UUID, &Np2)
	ObjsCache.Store("NetworkPolicy", NpObjMap)

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

	d := NewDependencyProcessor(ObjsCache)
	d.Evaluate(Vn1, "VirtualNetwork", "Self")
	resources := d.GetResources()

	networks := mustLoad(t, resources, "VirtualNetwork")
	mustLoad(t, networks, Vn1.UUID)
	mustLoad(t, networks, Vn2.UUID)

	policies := mustLoad(t, resources, "NetworkPolicy")
	mustLoad(t, policies, Np1.UUID)
	mustLoad(t, policies, Np2.UUID)
}

func TestReturnsSelf(t *testing.T) {
	ObjsCache := make(map[string]map[string]interface{})

	ObjsCache["virtual_network"] = make(map[string]interface{})
	Vn1 := models.VirtualNetwork{
		UUID: "Virtual-Network-1",
	}
	ObjsCache["virtual_network"][Vn1.UUID] = &Vn1

	d := NewDependencyProcessor(ObjsCache)
	d.Evaluate(Vn1, "VirtualNetwork", "Self")
	resources := d.GetResources()

	networks := mustLoad(t, resources, "VirtualNetwork")
	mustLoad(t, networks, Vn1.UUID)
}

func mustLoad(t *testing.T, rawSyncMap interface{}, key string) interface{} {
	syncMap, ok := rawSyncMap.(*sync.Map)
	assert.Truef(t, ok, "%v should be a sync.Map", rawSyncMap)

	value, ok := syncMap.Load(key)
	assert.Truef(t, ok, "The map should have a '%s' key", key)
	return value
}
