package intent

import (
	"fmt"
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
		t.Run(fmt.Sprintf("%s (query by UUID)", tt.name), func(t *testing.T) {
			if tt.storeResource {
				c.Store(&TestIntent{
					VirtualNetwork: vn,
				})
			}

			i := c.Load(tt.typeName, ByUUID(vn.GetUUID()))
			assert.Equal(t, tt.expectedOk, i != nil)
		})

		t.Run(fmt.Sprintf("%s (query by FQName)", tt.name), func(t *testing.T) {
			if tt.storeResource {
				c.Store(&TestIntent{
					VirtualNetwork: vn,
				})
			}

			i := c.Load(tt.typeName, ByFQName(vn.GetFQName()))
			assert.Equal(t, tt.expectedOk, i != nil)
		})
	}
}

func TestCacheDelete(t *testing.T) {
	c := NewCache()

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
		t.Run(fmt.Sprintf("%s (query by UUID)", tt.name), func(t *testing.T) {
			if tt.storeResource {
				c.Store(&TestIntent{
					VirtualNetwork: vn,
				})
			}

			c.Delete(tt.typeName, ByUUID(vn.GetUUID()))

			i := c.Load(tt.typeName, ByUUID(vn.GetUUID()))
			assert.Nil(t, i)
		})

		t.Run(fmt.Sprintf("%s (query by FQName)", tt.name), func(t *testing.T) {
			if tt.storeResource {
				c.Store(&TestIntent{
					VirtualNetwork: vn,
				})
			}

			c.Delete(tt.typeName, ByFQName(vn.GetFQName()))

			i := c.Load(tt.typeName, ByFQName(vn.GetFQName()))
			assert.Nil(t, i)
		})
	}
}
