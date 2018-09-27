package intent_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/compilation/logic"
	"github.com/Juniper/contrail/pkg/models"
)

func TestCacheLoad(t *testing.T) {
	c := intent.NewCache()

	vn := &models.VirtualNetwork{
		UUID:   "hoge",
		FQName: []string{"dead", "beef"},
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
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s (query by UUID)", tt.name), func(t *testing.T) {
			if tt.storeResource {
				c.Store(&logic.VirtualNetworkIntent{
					VirtualNetwork: vn,
				})
			}

			i := c.Load(tt.typeName, intent.ByUUID(vn.GetUUID()))
			assert.Equal(t, tt.expectedOk, i != nil)
		})

		t.Run(fmt.Sprintf("%s (query by FQName)", tt.name), func(t *testing.T) {
			if tt.storeResource {
				c.Store(&logic.VirtualNetworkIntent{
					VirtualNetwork: vn,
				})
			}

			i := c.Load(tt.typeName, intent.ByFQName(vn.GetFQName()))
			assert.Equal(t, tt.expectedOk, i != nil)
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
			name:          "deletesStoredResource",
			typeName:      vn.Kind(),
			storeResource: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s (query by UUID)", tt.name), func(t *testing.T) {
			if tt.storeResource {
				c.Store(&logic.VirtualNetworkIntent{
					VirtualNetwork: vn,
				})
			}

			c.Delete(tt.typeName, intent.ByUUID(vn.GetUUID()))

			i := c.Load(tt.typeName, intent.ByUUID(vn.GetUUID()))
			assert.Nil(t, i)
		})

		t.Run(fmt.Sprintf("%s (query by FQName)", tt.name), func(t *testing.T) {
			if tt.storeResource {
				c.Store(&logic.VirtualNetworkIntent{
					VirtualNetwork: vn,
				})
			}

			c.Delete(tt.typeName, intent.ByFQName(vn.GetFQName()))

			i := c.Load(tt.typeName, intent.ByFQName(vn.GetFQName()))
			assert.Nil(t, i)
		})
	}
}

func TestDependencyResolution(t *testing.T) {
	c := intent.NewCache()

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

	c.Store(vnBlueIntent)
	vn := logic.LoadVirtualNetworkIntent(c, vnBlue.UUID)
	if assert.NotNil(t, vn) {
		assert.Equal(t, 0, len(vn.RoutingInstances))
	}

	c.Store(ri1Intent)
	vn = logic.LoadVirtualNetworkIntent(c, vnBlue.UUID)
	if assert.NotNil(t, vn) {
		dependencies := vn.GetDependencies()
		if assert.Contains(t, dependencies, "routing-instance") {
			assert.Contains(t, dependencies["routing-instance"], "ri_uuid1")
		}
		assert.Equal(t, 1, len(vn.RoutingInstances))
	}

	c.Store(ri2Intent)
	ri := logic.LoadRoutingInstanceIntent(c, ri1Intent.UUID)
	if assert.NotNil(t, ri) {
		dependencies := ri.GetDependencies()
		if assert.Contains(t, dependencies, "routing-instance") {
			assert.Contains(t, dependencies["routing-instance"], "ri_uuid2")
		}
		assert.Equal(t, 1, len(ri.RoutingInstanceBackRefs))
	}
	ri = logic.LoadRoutingInstanceIntent(c, ri2Intent.UUID)
	if assert.NotNil(t, ri) {
		assert.Equal(t, 0, len(ri.RoutingInstanceBackRefs))
	}
	vn = logic.LoadVirtualNetworkIntent(c, vnBlue.UUID)
	if assert.NotNil(t, vn) {
		assert.Equal(t, 2, len(vn.RoutingInstances))
	}

	c.Delete(ri2Intent.Kind(), intent.ByUUID(ri2Intent.GetUUID()))
	ri = logic.LoadRoutingInstanceIntent(c, ri1Intent.UUID)
	if assert.NotNil(t, ri) {
		assert.Equal(t, 0, len(ri.RoutingInstanceBackRefs))
	}
	vn = logic.LoadVirtualNetworkIntent(c, vnBlue.UUID)
	if assert.NotNil(t, vn) {
		assert.Equal(t, 1, len(vn.RoutingInstances))
	}
}
