package intent

import (
	"sync"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"
)

// Cache cache for intents
type Cache struct {
	m *intentMap
}

// NewCache creates new cache for intents
func NewCache() *Cache {
	return &Cache{
		m: newIntentMap(),
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

func (m *intentMap) loadTypeMap(typeName string) map[string]Intent {
	objMap, found := m.internal[typeName]
	if !found {
		objMap = map[string]Intent{}
		m.internal[typeName] = objMap
	}
	return objMap
}

func (m *intentMap) load(typeName, uuid string) (value Intent, ok bool) {
	typeName = strcase.ToCamel(typeName)
	typeMap := m.loadTypeMap(typeName)
	intent, ok := typeMap[uuid]
	return intent, ok
}

func (m *intentMap) Load(typeName, uuid string) (value Intent, ok bool) {
	m.RLock()
	defer m.RUnlock()
	return m.load(typeName, uuid)
}

func (m *intentMap) delete(typeName, uuid string) {
	typeName = strcase.ToCamel(typeName)
	typeMap := m.loadTypeMap(typeName)
	delete(typeMap, uuid)
	delete(m.uuidToType, uuid)
}

func (m *intentMap) Delete(typeName, uuid string) {
	m.Lock()
	defer m.Unlock()
	m.delete(typeName, uuid)
}

func (m *intentMap) Store(typeName, uuid string, i Intent) {
	m.Lock()
	defer m.Unlock()
	oldIntent, ok := m.load(typeName, uuid)
	if ok {
		m.removeDependencies(oldIntent)
	}
	m.store(typeName, uuid, i)
	m.addDependencies(i)
}

func (m *intentMap) translateUUIDToType(uuid string) (string, bool) {
	typeName, ok := m.uuidToType[uuid]
	if !ok {
		return "", false
	}
	return strcase.ToCamel(typeName), true
}

func (m *intentMap) store(typeName, uuid string, intent Intent) {
	typeName = strcase.ToCamel(typeName)
	typeMap := m.loadTypeMap(typeName)
	typeMap[uuid] = intent
	m.uuidToType[uuid] = typeName
}

func (m *intentMap) debug() {
	log.Debug("Cache content:")
	m.Lock()
	defer m.Unlock()
	for t, v := range m.internal {
		for uuid := range v {
			log.Debugf("Type: %s, UUID: %s", t, uuid)
		}
	}
}

func (m *intentMap) addDependencies(resource services.Resource) {
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

func (m *intentMap) removeDependencies(resource services.Resource) {
	dependencies := resource.Depends()
	for _, dependencyID := range dependencies {
		t, ok := m.translateUUIDToType(dependencyID)
		if !ok {
			continue
		}
		dependentIntent, ok := m.load(t, dependencyID)
		if ok {
			dependentIntent.RemoveDependency(resource)
		}
	}
}

// Load loads intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Load(typeName, uuid string) (Intent, bool) {
	log.Debugf("Loading: TypeName: %s, UUID: %s", typeName, uuid)
	c.m.debug()
	return c.m.Load(typeName, uuid)
}

// Store puts intent into cache.
func (c *Cache) Store(i Intent) {
	uuid := i.GetUUID()
	kind := i.Kind()
	log.Debugf("Storing: TypeName: %s, UUID: %s", kind, uuid)
	c.m.Store(kind, uuid, i)
	c.m.debug()
}

// Delete deletes intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Delete(typeName, uuid string) {
	log.Debugf("Deleting: TypeName: %s, UUID: %s", typeName, uuid)
	c.m.Delete(typeName, uuid)
	c.m.debug()
}
