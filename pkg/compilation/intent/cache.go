package intent

import (
	"sync"

	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// Cache cache for intents.
type Cache struct {
	*intentStore
}

// NewCache creates new cache for intents.
func NewCache() *Cache {
	return &Cache{
		intentStore: newIntentStore(),
	}
}

func newIntentStore() *intentStore {
	return &intentStore{
		typeToIntents: make(map[string]intents),
	}
}

type intentStore struct {
	typeToIntents map[string]intents
	sync.RWMutex
}

type intents struct {
	fqNameToID map[string]string
	idToIntent map[string]Intent
}

// LoadByFQName loads intent from cache. It accepts as type both kebab-case and CamelCase
func (c *Cache) LoadByFQName(typeName string, fqName []string) (Intent, bool) {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"TypeName": typeName, "FQName": fqName}).Debug("Storing in cache")
	return c.intentStore.loadByFQName(typeName, fqName)
}

// Load loads intent from cache. It accepts as type both kebab-case and CamelCase
func (c *Cache) Load(typeName, uuid string) (Intent, bool) {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"type-name": typeName, "uuid": uuid}).Debug("Loading from cache")
	return c.intentStore.load(typeName, uuid)
}

// Store puts intent into cache.
func (c *Cache) Store(i Intent) {
	typeName := strcase.ToCamel(i.Kind())
	log.WithFields(log.Fields{"type-name": typeName, "uuid": i.GetUUID()}).Debug("Storing in cache")
	c.intentStore.store(typeName, i)
}

// Delete deletes intent from cache. It accepts as type both kebab-case and CamelCase
func (c *Cache) Delete(typeName, uuid string) {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"TypeName": typeName, "UUID": uuid}).Debug("Deleting from cache")
	c.intentStore.delete(typeName, uuid)
}

func (s *intentStore) loadByFQName(typeName string, fqName []string) (Intent, bool) {
	s.RLock()
	defer s.RUnlock()
	is, ok := s.typeToIntents[typeName]
	if !ok {
		return nil, false
	}
	i, ok := is.idToIntent[is.fqNameToID[fqNameKey(fqName)]]
	return i, ok
}

func (s *intentStore) load(typeName, uuid string) (Intent, bool) {
	s.RLock()
	defer s.RUnlock()
	is, ok := s.typeToIntents[typeName]
	if !ok {
		return nil, false
	}
	i, ok := is.idToIntent[uuid]
	return i, ok
}

func (s *intentStore) delete(typeName, uuid string) {
	s.Lock()
	defer s.Unlock()
	is, ok := s.typeToIntents[typeName]
	if !ok {
		return
	}
	i, ok := is.idToIntent[uuid]
	if !ok {
		return
	}
	delete(is.fqNameToID, fqNameKey(i.GetFQName()))
	delete(is.idToIntent, uuid)
}

func (s *intentStore) store(typeName string, intent Intent) {
	s.Lock()
	defer s.Unlock()
	is, ok := s.typeToIntents[typeName]
	if !ok {
		is = intents{
			fqNameToID: map[string]string{},
			idToIntent: map[string]Intent{},
		}
		s.typeToIntents[typeName] = is
	}
	is.idToIntent[intent.GetUUID()] = intent
	is.fqNameToID[fqNameKey(intent.GetFQName())] = intent.GetUUID()
}

func fqNameKey(fqName []string) string {
	return basemodels.FQNameToString(fqName)
}
