package intent

import (
	"fmt"
	"sync"

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
		internal: make(map[string]map[string]Intent),
	}
}

type intentMap struct {
	internal map[string]map[string]Intent
	sync.RWMutex
}

func (im *intentMap) loadTypeMap(typeName string) map[string]Intent {
	objMap, found := im.internal[typeName]
	if !found {
		objMap = map[string]Intent{}
		im.internal[typeName] = objMap
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

func (m *intentMap) debug() {
	log.Debug("Cache content:")
	m.Lock()
	defer m.Unlock()
	for t, v := range m.internal {
		for uuid, _ := range v {
			log.Debugf("Type: %s, UUID: %s", t, uuid)
		}
	}
}

// Load loads intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Load(typeName, uuid string) (Intent, bool) {
	typeName = strcase.ToCamel(typeName)
	log.Debugf("Loading: TypeName: %s, UUID: %s", typeName, uuid)
	c.m.debug()
	return c.m.load(typeName, uuid)
}

// Delete puts intent into cache.
func (c *Cache) Store(i Intent) {
	typeName := strcase.ToCamel(i.Kind())
	uuid := i.GetUUID()
	log.Debugf("Storing: TypeName: %s, UUID: %s", typeName, uuid)
	c.m.store(typeName, uuid, i)
	c.m.debug()
}

// Delete deletes intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Delete(typeName, uuid string) {
	typeName = strcase.ToCamel(typeName)
	log.Debugf("Deleting: TypeName: %s, UUID: %s", typeName, uuid)
	c.m.delete(typeName, uuid)
	c.m.debug()
}

func wrongCacheTypeError(got, expected interface{}) error {
	return fmt.Errorf("got wrong type from cache. expected %T, got %T", expected, got)
}
