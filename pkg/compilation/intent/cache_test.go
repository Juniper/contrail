package intent_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/compilation/logic"
	"github.com/Juniper/contrail/pkg/models"
)

func TestCacheLoad(t *testing.T) {
	c := intent.NewCache()

	vn := &models.VirtualNetwork{
		UUID: "hoge",
	}

	// Test empty internal map
	_, ok := c.Load(vn.Kind(), vn.GetUUID())
	assert.False(t, ok)
	_, ok = c.Load(vn.TypeName(), vn.GetUUID())
	assert.False(t, ok)

	// Test invalid type name internal map
	_, ok = c.Load("hoge", vn.GetUUID())
	assert.False(t, ok)

	// Store and load
	c.Store(&logic.VirtualNetworkIntent{
		VirtualNetwork: vn,
	})

	i, ok := c.Load(vn.Kind(), vn.GetUUID())
	assert.True(t, ok)
	_, ok = c.Load(vn.TypeName(), vn.GetUUID())
	assert.True(t, ok)

	actualIntent, ok := i.(*logic.VirtualNetworkIntent)
	if assert.True(t, ok) {
		assert.Equal(t, vn, actualIntent.VirtualNetwork)
	}

	c.Delete(vn.Kind(), vn.GetUUID())

	_, ok = c.Load(vn.Kind(), vn.GetUUID())

	assert.False(t, ok)
}

func TestCacheDelete(t *testing.T) {
	c := intent.NewCache()

	vn := &models.VirtualNetwork{
		UUID: "hoge",
	}

	// Try to delete resource that doesn't exists
	c.Delete(vn.Kind(), vn.GetUUID())
	c.Delete(vn.TypeName(), vn.GetUUID())

	c.Store(&logic.VirtualNetworkIntent{
		VirtualNetwork: vn,
	})

	i, ok := c.Load(vn.Kind(), vn.GetUUID())
	assert.True(t, ok)
	_, ok = c.Load(vn.TypeName(), vn.GetUUID())
	assert.True(t, ok)

	actualIntent, ok := i.(*logic.VirtualNetworkIntent)
	if assert.True(t, ok) {
		assert.Equal(t, vn, actualIntent.VirtualNetwork)
	}

	c.Delete(vn.Kind(), vn.GetUUID())

	_, ok = c.Load(vn.Kind(), vn.GetUUID())

	assert.False(t, ok)
}

func TestDependencyResolution(t *testing.T) {
	cache := intent.NewCache()

	vnBlue := &models.VirtualNetwork{
		UUID: "vn_blue",
	}

	ri1 := &models.RoutingInstance{
		UUID:       "ri_uuid1",
		ParentUUID: vnBlue.GetUUID(),
	}
	ri2 := &models.RoutingInstance{
		UUID:       "ri_uuid2",
		ParentUUID: vnBlue.GetUUID(),
		RoutingInstanceRefs: []*models.RoutingInstanceRoutingInstanceRef{
			{
				UUID: ri1.UUID,
			},
		},
	}

	vnBlueIntent := &logic.VirtualNetworkIntent{
		VirtualNetwork: vnBlue,
	}
	ri1Intent := &logic.RoutingInstanceIntent{
		RoutingInstance: ri1,
	}
	ri2Intent := &logic.RoutingInstanceIntent{
		RoutingInstance: ri2,
	}

	cache.Store(vnBlueIntent)
	vn, ok := logic.LoadVirtualNetworkIntent(cache, vnBlue.UUID)
	if assert.True(t, ok) {
		assert.Equal(t, 0, len(vn.RoutingInstances))
	}

	cache.Store(ri1Intent)
	vn, ok = logic.LoadVirtualNetworkIntent(cache, vnBlue.UUID)
	if assert.True(t, ok) {
		assert.Equal(t, 1, len(vn.RoutingInstances))
	}

	cache.Store(ri2Intent)
	ri, ok := logic.LoadRoutingInstanceIntent(cache, ri1Intent.UUID)
	if assert.True(t, ok) {
		assert.Equal(t, 1, len(ri.RoutingInstanceBackRefs))
	}
	vn, ok = logic.LoadVirtualNetworkIntent(cache, vnBlue.UUID)
	if assert.True(t, ok) {
		assert.Equal(t, 2, len(vn.RoutingInstances))
	}
}
