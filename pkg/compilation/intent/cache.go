package intent

import (
	"sync"

	"github.com/Juniper/contrail/pkg/models/basemodels"
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
	objMap, ok := m.internal[typeName]
	if !ok {
		objMap = map[string]Intent{}
		m.internal[typeName] = objMap
	}
	return objMap
}

func (m *intentMap) load(typeName, uuid string) (value Intent, ok bool) {
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
	typeMap := m.loadTypeMap(typeName)
	delete(typeMap, uuid)
	delete(m.uuidToType, uuid)
}

func (m *intentMap) Delete(typeName, uuid string) {
	m.Lock()
	defer m.Unlock()
	m.delete(typeName, uuid)
}

func (m *intentMap) store(typeName, uuid string, intent Intent) {
	typeMap := m.loadTypeMap(typeName)
	typeMap[uuid] = intent
	m.uuidToType[uuid] = typeName
}

func (m *intentMap) Store(typeName, uuid string, i Intent) {
	m.Lock()
	defer m.Unlock()
	oldIntent, ok := m.load(typeName, uuid)
	if ok {
		m.removeDependencies(oldIntent.GetObject())
	}
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

func (m *intentMap) removeDependencies(resource basemodels.Object) {
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
	typeName = strcase.ToCamel(typeName)
	log.Debugf("Loading: TypeName: %s, UUID: %s", typeName, uuid)
	return c.m.Load(typeName, uuid)
}

// Store puts intent into cache.
func (c *Cache) Store(i Intent) {
	typeName := strcase.ToCamel(i.Kind())
	uuid := i.GetUUID()
	log.Debugf("Storing: TypeName: %s, UUID: %s", typeName, uuid)
	c.m.Store(typeName, uuid, i)
}

// Delete deletes intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Delete(typeName, uuid string) {
	typeName = strcase.ToCamel(typeName)
	log.Debugf("Deleting: TypeName: %s, UUID: %s", typeName, uuid)
	c.m.Delete(typeName, uuid)
}
