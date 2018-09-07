package intent

import (
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/stretchr/testify/assert"
)

type TestIntent struct {
	BaseIntent
	*models.VirtualNetwork
}

func TestCache(t *testing.T) {
	c := NewCache()

	vn := &models.VirtualNetwork{
		UUID: "hoge",
	}

	_, ok := c.Load(vn.Kind(), vn.GetUUID())

	assert.False(t, ok)

	c.Store(&TestIntent{
		VirtualNetwork: vn,
	})

	intent, ok := c.Load(vn.Kind(), vn.GetUUID())
	assert.True(t, ok)
	actualVN, ok := intent.(*TestIntent)
	if assert.True(t, ok) {
		assert.Equal(t, vn, actualVN.VirtualNetwork)
	}

	c.Delete(vn.Kind(), vn.GetUUID())

	_, ok = c.Load(vn.Kind(), vn.GetUUID())

	assert.False(t, ok)
}
