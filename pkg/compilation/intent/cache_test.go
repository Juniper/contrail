package intent

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

type TestIntent struct {
	BaseIntent
	*models.VirtualNetwork
}

func TestCacheLoad(t *testing.T) {
	c := NewCache()

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
	c.Store(&TestIntent{
		VirtualNetwork: vn,
	})

	intent, ok := c.Load(vn.Kind(), vn.GetUUID())
	assert.True(t, ok)
	_, ok = c.Load(vn.TypeName(), vn.GetUUID())
	assert.True(t, ok)

	actualIntent, ok := intent.(*TestIntent)
	if assert.True(t, ok) {
		assert.Equal(t, vn, actualIntent.VirtualNetwork)
	}

	c.Delete(vn.Kind(), vn.GetUUID())

	_, ok = c.Load(vn.Kind(), vn.GetUUID())

	assert.False(t, ok)
}

func TestCacheDelete(t *testing.T) {
	c := NewCache()

	vn := &models.VirtualNetwork{
		UUID: "hoge",
	}

	c.Store(&TestIntent{
		VirtualNetwork: vn,
	})

	intent, ok := c.Load(vn.Kind(), vn.GetUUID())
	assert.True(t, ok)
	_, ok = c.Load(models.KindVirtualNetwork, vn.GetUUID())
	assert.True(t, ok)

	actualIntent, ok := intent.(*TestIntent)
	if assert.True(t, ok) {
		assert.Equal(t, vn, actualIntent.VirtualNetwork)
	}

	c.Delete(vn.Kind(), vn.GetUUID())

	_, ok = c.Load(vn.Kind(), vn.GetUUID())

	assert.False(t, ok)
}
