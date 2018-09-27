package dependencies_test

import (
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/compilation/logic"
	"github.com/Juniper/contrail/pkg/compilation/plugins/contrail/dependencies"
)

func TestReturnsRefs(t *testing.T) {
	c := intent.NewCache()
	d := dependencies.NewDependencyProcessor(logic.ReactionMap)

	// Create NetworkPolicys in the Cache
	Np1 := models.NetworkPolicy{
		UUID: "Network-policy-1",
	}
	Np1Ref := &models.VirtualNetworkNetworkPolicyRef{
		UUID: Np1.UUID,
	}
	Np2 := models.NetworkPolicy{
		UUID: "Network-policy-2",
	}
	Np2Ref := &models.VirtualNetworkNetworkPolicyRef{
		UUID: Np2.UUID,
	}
	Vn1 := models.VirtualNetwork{
		UUID: "Virtual-Network-1",
		NetworkPolicyRefs: []*models.VirtualNetworkNetworkPolicyRef{
			Np1Ref,
			Np2Ref,
		},
	}
	Vn1Ref := &models.VirtualNetworkVirtualNetworkRef{
		UUID: Vn1.UUID,
	}
	Vn2 := models.VirtualNetwork{
		UUID: "Virtual-Network-2",
		NetworkPolicyRefs: []*models.VirtualNetworkNetworkPolicyRef{
			Np2Ref,
			Np1Ref,
		},
		VirtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
			Vn1Ref,
		},
	}
	Vn2Ref := &models.VirtualNetworkVirtualNetworkRef{
		UUID: Vn2.UUID,
	}

	storeTestNetworkPolicyIntent(c, &Np1)
	storeTestNetworkPolicyIntent(c, &Np2)
	storeTestVirtualNetworkIntent(c, &Vn1)
	storeTestVirtualNetworkIntent(c, &Vn2)

	Vn1 = models.VirtualNetwork{
		UUID: "Virtual-Network-1",
		NetworkPolicyRefs: []*models.VirtualNetworkNetworkPolicyRef{
			Np1Ref,
			Np2Ref,
		},
		VirtualNetworkRefs: []*models.VirtualNetworkVirtualNetworkRef{
			Vn2Ref,
		},
	}
	vn1Intent := c.Load(Vn1.Kind(), intent.ByUUID(Vn1.GetUUID()))
	storeTestVirtualNetworkIntent(c, &Vn1)

	vn1Intent = c.Load(Vn1.Kind(), intent.ByUUID(Vn1.GetUUID()))

	l := d.GetDependencies(c, vn1Intent, "self")

	assert.Contains(t, l, "Virtual-Network-1")
	assert.Contains(t, l, "Virtual-Network-2")
	assert.Contains(t, l, "Network-policy-1")
	assert.Contains(t, l, "Network-policy-1")
}

func TestReturnsSelf(t *testing.T) {
	c := intent.NewCache()
	d := dependencies.NewDependencyProcessor(logic.ReactionMap)

	vn := &models.VirtualNetwork{
		UUID: "Virtual-Network-1",
	}
	vnIntent := storeTestVirtualNetworkIntent(c, vn)

	l := d.GetDependencies(c, vnIntent, "self")

	assert.Contains(t, l, "Virtual-Network-1")
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
