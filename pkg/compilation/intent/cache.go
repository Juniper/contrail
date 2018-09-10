package intent

import (
	"strings"
	"sync"

	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"
)

// Cache cache for intents
type Cache struct {
	s *intentStore
}

// NewCache creates new cache for intents
func NewCache() *Cache {
	return &Cache{
		s: newIntentStore(),
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

// LoadByFQName loads intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) LoadByFQName(typeName string, fqName []string) (Intent, bool) {
	typeName = strcase.ToCamel(typeName)
	log.Debugf("Loading: TypeName: %s, FQName: %v", typeName, fqName)
	c.s.debug()
	return c.s.loadByFQName(typeName, fqName)
}

// Load loads intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Load(typeName, uuid string) (Intent, bool) {
	typeName = strcase.ToCamel(typeName)
	log.Debugf("Loading: TypeName: %s, UUID: %s", typeName, uuid)
	c.s.debug()
	return c.s.load(typeName, uuid)
}

// Store puts intent into cache.
func (c *Cache) Store(i Intent) {
	typeName := strcase.ToCamel(i.Kind())
	log.Debugf("Storing: TypeName: %s, UUID: %s", typeName, i.GetUUID())
	c.s.store(typeName, i)
	c.s.debug()
}

// Delete deletes intent from cache. It accepts as type both snake-case and CamelCase
func (c *Cache) Delete(typeName, uuid string) {
	typeName = strcase.ToCamel(typeName)
	log.Debugf("Deleting: TypeName: %s, UUID: %s", typeName, uuid)
	c.s.delete(typeName, uuid)
	c.s.debug()
}

func (s *intentStore) loadByFQName(typeName string, fqName []string) (Intent, bool) {
	s.RLock()
	defer s.RUnlock()
	is := s.loadIntents(typeName)
	i, ok := is.idToIntent[is.fqNameToID[fqNameKey(fqName)]]
	return i, ok
}

func (s *intentStore) load(typeName, uuid string) (Intent, bool) {
	s.RLock()
	defer s.RUnlock()
	i, ok := s.loadIntents(typeName).idToIntent[uuid]
	return i, ok
}

func (s *intentStore) delete(typeName, uuid string) {
	s.Lock()
	defer s.Unlock()
	is := s.loadIntents(typeName)
	i, found := is.idToIntent[uuid]
	if !found {
		return
	}
	delete(is.fqNameToID, fqNameKey(i.GetFQName()))
	delete(is.idToIntent, uuid)
}

func (s *intentStore) store(typeName string, intent Intent) {
	s.Lock()
	defer s.Unlock()
	is := s.loadIntents(typeName)
	is.idToIntent[intent.GetUUID()] = intent
	is.fqNameToID[fqNameKey(intent.GetFQName())] = intent.GetUUID()
}

func (s *intentStore) loadIntents(typeName string) intents {
	is, found := s.typeToIntents[typeName]
	if !found {
		is = intents{
			fqNameToID: map[string]string{},
			idToIntent: map[string]Intent{},
		}
		s.typeToIntents[typeName] = is
	}
	return is
}

func (s *intentStore) debug() {
	log.Debug("Cache content:")
	s.Lock()
	defer s.Unlock()
	for t, is := range s.typeToIntents {
		for fqname, uuid := range is.fqNameToID {
			log.Debugf("Type: %s, FQName: %s, UUID: %s", t, fqname, uuid)
		}
	}
}

func fqNameKey(fqName []string) string {
	return strings.Join(fqName, ":")
}
