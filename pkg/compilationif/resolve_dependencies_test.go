package compilationif

import (
	"sync"
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveDependenciesReturnsSelf(t *testing.T) {
	compiler := NewCompilationService()

	network := &models.VirtualNetwork{
		UUID: "test_uuid",
	}
	putIntoCache(t, ObjsCache, "VirtualNetwork", network)

	event := services.NewEvent(&services.EventOption{
		UUID:      network.GetUUID(),
		Operation: services.OperationCreate,
		Kind:      "virtual_network", // TODO Can't use network.Kind() here
		Data:      network.ToMap(),
	})
	require.NotNil(t, event)
	require.NotNil(t, event.GetResource())

	allEvents := compiler.resolveDependencies(event)
	assert.Contains(t, allEvents, event)
}

func putIntoCache(t *testing.T, cache *sync.Map, kind string, resource services.Resource) {
	resourceMap := ensureSubmap(t, cache, kind)
	resourceMap.Store(resource.GetUUID(), resource)
}

func ensureSubmap(t *testing.T, cache *sync.Map, key string) (submap *sync.Map) {
	rawSubmap, ok := cache.Load(key)
	if !ok {
		submap = &sync.Map{}
		cache.Store(key, submap)
		return
	}
	submap, ok = rawSubmap.(*sync.Map)
	require.True(t, ok)
	return
}
