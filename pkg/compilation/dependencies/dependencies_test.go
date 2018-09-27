package dependencies_test

import (
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/compilation/dependencies"
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/compilation/logic"
)

func TestReturnsRefs(t *testing.T) {
	c := intent.NewCache()
	reactions, err := dependencies.ParseReactions("./test_reactions.yml", "test")
	require.NoError(t, err)
	d := dependencies.NewDependencyProcessor(reactions)
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
	Vn2 := models.VirtualNetwork{
		UUID: "Virtual-Network-2",
		NetworkPolicyRefs: []*models.VirtualNetworkNetworkPolicyRef{
			Np2Ref,
		},
	}

	storeTestNetworkPolicyIntent(c, &Np1)
	storeTestNetworkPolicyIntent(c, &Np2)
	storeTestVirtualNetworkIntent(c, &Vn1)
	storeTestVirtualNetworkIntent(c, &Vn2)

	vn1Intent := loadIntentByResource(c, &Vn1)
	l := d.GetDependencies(c, vn1Intent, "self")

	assert.Contains(t, l, "Virtual-Network-1")
	assert.Contains(t, l, "Virtual-Network-2")
	assert.Contains(t, l, "Network-policy-1")
	assert.Contains(t, l, "Network-policy-1")
}

func loadIntentByResource(c *intent.Cache, r basemodels.Object) intent.Intent {
	return c.Load(r.Kind(), intent.ByUUID(r.GetUUID()))
}

func TestReturnsSelf(t *testing.T) {
	c := intent.NewCache()
	reactions, err := dependencies.ParseReactions("./test_reactions.yml", "test")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	d := dependencies.NewDependencyProcessor(reactions)

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
