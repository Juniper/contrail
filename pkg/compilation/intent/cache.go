package intent

import (
	"sync"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"
)

// Cache cache for intents.
type Cache struct {
	intentMap *intentMap
}

// NewCache creates new cache for intents.
func NewCache() *Cache {
	return &Cache{
		intentMap: newIntentMap(),
	}
}

func newIntentMap() *intentMap {
	return &intentMap{
		internal:   make(map[string]map[string]Intent),
		uuidToType: make(map[string]string),
	}
}

type intentMap struct {
	internal   map[string]map[string]Intent
	uuidToType map[string]string
	sync.RWMutex
}

func (m *intentMap) load(typeName, uuid string) (value Intent, ok bool) {
	typeMap, ok := m.internal[typeName]
	if !ok {
		return nil, false
	}
	intent, ok := typeMap[uuid]
	return intent, ok
}

func (m *intentMap) Load(typeName, uuid string) (value Intent, ok bool) {
	m.RLock()
	defer m.RUnlock()
	return m.load(typeName, uuid)
}

func (m *intentMap) delete(typeName, uuid string) {
	typeMap, ok := m.internal[typeName]
	if ok {
		delete(typeMap, uuid)
	}
	delete(m.uuidToType, uuid)
}

func (m *intentMap) Delete(typeName, uuid string) {
	m.Lock()
	defer m.Unlock()
	m.removeDependencies(typeName, uuid)
	m.delete(typeName, uuid)
}

func (m *intentMap) store(typeName, uuid string, intent Intent) {
	typeMap, ok := m.internal[typeName]
	if !ok {
		typeMap = make(map[string]Intent)
		m.internal[typeName] = typeMap
	}
	typeMap[uuid] = intent
	m.uuidToType[uuid] = typeName
}

func (m *intentMap) Store(typeName, uuid string, i Intent) {
	m.Lock()
	defer m.Unlock()
	m.removeDependencies(typeName, uuid)
	m.store(typeName, uuid, i)
	m.addDependencies(i.GetObject())
}

func (m *intentMap) translateUUIDToType(uuid string) (string, bool) {
	typeName, ok := m.uuidToType[uuid]
	return typeName, ok
}

func (m *intentMap) addDependencies(resource basemodels.Object) {
	dependencies := resource.Depends()
	for _, dependencyID := range dependencies {
		t, ok := m.translateUUIDToType(dependencyID)
		if !ok {
			continue
		}
		dependentIntent, ok := m.load(t, dependencyID)
		if ok {
			dependentIntent.AddDependency(resource)
		}
	}
}

func (m *intentMap) removeDependencies(typeName, uuid string) {
	i, ok := m.load(typeName, uuid)
	if ok {
		dependencies := i.GetObject().Depends()
		for _, dependencyID := range dependencies {
			t, ok := m.translateUUIDToType(dependencyID)
			if !ok {
				continue
			}
			dependentIntent, ok := m.load(t, dependencyID)
			if ok {
				dependentIntent.RemoveDependency(i.GetObject())
			}
		}
	}
}

// Load loads intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Load(typeName, uuid string) (Intent, bool) {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"type-name": typeName, "uuid": uuid}).Debug("Loading from cache")
	return c.intentMap.Load(typeName, uuid)
}

// Store puts intent into cache.
func (c *Cache) Store(i Intent) {
	typeName := strcase.ToCamel(i.Kind())
	log.WithFields(log.Fields{"type-name": typeName, "uuid": i.GetUUID()}).Debug("Storing in cache")
	c.intentMap.Store(typeName, i.GetUUID(), i)
}

// Delete deletes intent from cache. It accepts as type both kebab-case and CamelCase
func (c *Cache) Delete(typeName, uuid string) {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"type-name": typeName, "uuid": uuid}).Debug("Deleting from cache")
	c.intentMap.Delete(typeName, uuid)
}
