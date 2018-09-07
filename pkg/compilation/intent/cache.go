package intent

import (
	"fmt"
	"sync"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Cache cache for intents
type Cache struct {
	m *sync.Map
}

// NewCache creates new cache for intents
func NewCache() *Cache {
	return &Cache{
		m: &sync.Map{},
	}
}

func (c *Cache) debug() {
	log.Debug("Cache content:")
	c.m.Range(func(k1, v1 interface{}) bool {
		typeName := k1.(string)
		typeMap := v1.(*sync.Map)
		typeMap.Range(func(k2, v2 interface{}) bool {
			objUUID := k2.(string)
			log.Debugf("Type: %s, UUID: %s", typeName, objUUID)
			return true
		})
		return true
	})
}

// Load loads intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Load(typeName, uuid string) (Intent, error) {
	typeName = strcase.ToCamel(typeName)
	log.Debugf("Loading: TypeName: %s, UUID: %s", typeName, uuid)
	typeMap, err := c.loadTypeMap(typeName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get intent from cache")
	}
	tmp, ok := typeMap.Load(uuid)
	if !ok {
		return nil, errors.Wrap(err, "failed to get intent from cache")
	}
	if tmp == nil {
		return nil, fmt.Errorf("intent not found. uuid: %s, type: %s", uuid, typeName)
	}
	intent, ok := tmp.(Intent)
	if !ok {
		return nil, wrongCacheTypeError(tmp, intent)
	}
	c.debug()
	return intent, nil
}

// Delete puts intent into cache.
func (c *Cache) Store(i Intent) error {
	typeName := strcase.ToCamel(i.Kind())
	log.Debugf("Storing: TypeName: %s, UUID: %s", typeName, i.GetUUID())
	typeMap, err := c.loadTypeMap(typeName)
	if err != nil {
		return errors.Wrap(err, "failed to store intent in cache")
	}
	typeMap.Store(i.GetUUID(), i)
	c.debug()
	return nil
}

// Delete deletes intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Delete(typeName, uuid string) error {
	typeName = strcase.ToCamel(typeName)
	log.Debugf("Deleting: TypeName: %s, UUID: %s", typeName, uuid)
	tmp, found := c.m.Load(typeName)
	if found {
		objMap, ok := tmp.(*sync.Map)
		if !ok {
			return wrongCacheTypeError(tmp, objMap)
		}
		c.m.Delete(uuid)
	}
	c.debug()
	return nil
}

func (c *Cache) loadTypeMap(typeName string) (*sync.Map, error) {
	var ok bool
	var objMap *sync.Map

	tmp, found := c.m.Load(typeName)
	if tmp != nil && found {
		objMap, ok = tmp.(*sync.Map)
		if !ok {
			return nil, wrongCacheTypeError(tmp, objMap)
		}
	} else {
		objMap = &sync.Map{}
		c.m.Store(typeName, objMap)
	}
	return objMap, nil
}

func wrongCacheTypeError(got, expected interface{}) error {
	return fmt.Errorf("got wrong type from cache. expected %T, got %T", expected, got)
}
