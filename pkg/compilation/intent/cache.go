package intent

import (
	"sync"

	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"
)

// Cache cache for intents
type Cache struct {
	intentMap *intentMap
}

type Loader interface {
	Load(typeName, uuid string) (Intent, bool)
}

// NewCache creates new cache for intents.
func NewCache() *Cache {
	return &Cache{
		intentMap: newIntentMap(),
	}
}

func newIntentMap() *intentMap {
	return &intentMap{
		internal: make(map[string]map[string]Intent),
	}
}

type intentMap struct {
	internal map[string]map[string]Intent
	sync.RWMutex
}

func (m *intentMap) loadTypeMap(typeName string) map[string]Intent {
	objMap, ok := m.internal[typeName]
	if !ok {
		objMap = make(map[string]Intent)
		m.internal[typeName] = objMap
	}
	return objMap
}

func (m *intentMap) load(typeName, uuid string) (value Intent, ok bool) {
	m.RLock()
	defer m.RUnlock()
	typeMap := m.loadTypeMap(typeName)
	intent, ok := typeMap[uuid]
	return intent, ok
}

func (m *intentMap) delete(typeName, uuid string) {
	m.Lock()
	defer m.Unlock()
	typeMap := m.loadTypeMap(typeName)
	delete(typeMap, uuid)
}

func (m *intentMap) store(typeName, uuid string, intent Intent) {
	m.Lock()
	defer m.Unlock()
	typeMap := m.loadTypeMap(typeName)
	typeMap[uuid] = intent
}

// Load loads intent from cache. It accepts as type both kebab-case and CamelCase
func (c *Cache) Load(typeName, uuid string) (Intent, bool) {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"TypeName": typeName, "UUID": uuid}).Debug("Loading from cache")
	return c.intentMap.load(typeName, uuid)
}

// Store puts intent into cache.
func (c *Cache) Store(i Intent) {
	typeName := strcase.ToCamel(i.Kind())
	log.WithFields(log.Fields{"TypeName": typeName, "UUID": i.GetUUID()}).Debug("Storing in cache")
	c.intentMap.store(typeName, i.GetUUID(), i)
}

// Delete deletes intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Delete(typeName, uuid string) {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"TypeName": typeName, "UUID": uuid}).Debug("Deleting from cache")
	c.intentMap.delete(typeName, uuid)
}
