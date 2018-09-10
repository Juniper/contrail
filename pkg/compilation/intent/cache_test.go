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

	tests := []struct {
		name          string
		typeName      string
		storeResource bool
		expectedOk    bool
	}{
		{
			name:     "failsWhenResourceIsNotInCache",
			typeName: vn.Kind(),
		},
		{
			name:     "failsWhenResourceIsNotInCacheUsingTypeName",
			typeName: vn.TypeName(),
		},
		{
			name:          "failsGivenIncorrectTypeName",
			typeName:      "hoge",
			storeResource: true,
		},
		{
			name:          "loadsStoredResource",
			typeName:      vn.Kind(),
			storeResource: true,
			expectedOk:    true,
		},
		{
			name:          "loadsStoredResourceUsingTypeName",
			typeName:      vn.TypeName(),
			storeResource: true,
			expectedOk:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.storeResource {
				c.Store(&logic.VirtualNetworkIntent{
					VirtualNetwork: vn,
				})
			}

			_, ok := c.Load(tt.typeName, vn.GetUUID())
			assert.Equal(t, tt.expectedOk, ok)
		})
	}
}

func TestCacheDelete(t *testing.T) {
	c := intent.NewCache()

	vn := &models.VirtualNetwork{
		UUID: "hoge",
	}

	tests := []struct {
		name          string
		typeName      string
		storeResource bool
	}{
		{
			name:     "doesNothingGivenInvalidUUID",
			typeName: vn.Kind(),
		},
		{
			name:     "doesNothingGivenInvalidUUIDUsingTypeName",
			typeName: vn.TypeName(),
		},
		{
			name:          "deletesStoredResource",
			typeName:      vn.Kind(),
			storeResource: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.storeResource {
				c.Store(&logic.VirtualNetworkIntent{
					VirtualNetwork: vn,
				})
			}

			c.Delete(tt.typeName, vn.GetUUID())

			_, ok := c.Load(tt.typeName, vn.GetUUID())
			assert.False(t, ok)
		})
	}
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
