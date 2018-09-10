package intent

import (
	"fmt"
	"sync"

	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// Cache stores intents.
type Cache struct {
	*intentStore
}

// Query finds an intent in cache.
type Query struct {
	load        func(intents) Intent
	description string
}

// String returns the description of a Query.
func (q Query) String() string {
	return q.description
}

// ByUUID returns a query to find an intent by UUID.
func ByUUID(uuid string) Query {
	return Query{
		load: func(is intents) Intent {
			return is.uuidToIntent[uuid]
		},
		description: fmt.Sprintf("by UUID: %v", uuid),
	}
}

// ByFQName returns a query to find an intent by FQName.
func ByFQName(fqName []string) Query {
	return Query{
		load: func(is intents) Intent {
			return is.uuidToIntent[is.fqNameToUUID[fqNameKey(fqName)]]
		},
		description: fmt.Sprintf("by FQName: %v", fqName),
	}
}

// NewCache creates new cache for intents.
func NewCache() *Cache {
	return &Cache{
		intentStore: newIntentStore(),
	}
}

// Load loads intent from cache. It accepts kebab-case or CamelCase type name.
func (c *Cache) Load(typeName string, q Query) Intent {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"type-name": typeName, "query": q}).Debug("Loading from cache")
	return c.intentStore.load(typeName, q)
}

// Store puts intent into cache.
func (c *Cache) Store(i Intent) {
	typeName := strcase.ToCamel(i.Kind())
	log.WithFields(log.Fields{"type-name": typeName, "uuid": i.GetUUID()}).Debug("Storing in cache")
	c.intentStore.store(typeName, i)
}

// Delete deletes intent from cache. It accepts kebab-case or CamelCase type name.
func (c *Cache) Delete(typeName string, q Query) {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"type-name": typeName, "query": q}).Debug("Deleting from cache")
	c.intentStore.delete(typeName, q)
}

type intentStore struct {
	typeNameToIntents map[string]intents
	sync.RWMutex
}

type intents struct {
	fqNameToUUID map[string]string
	uuidToIntent map[string]Intent
}

func newIntentStore() *intentStore {
	return &intentStore{
		typeNameToIntents: make(map[string]intents),
	}
}

func (s *intentStore) load(typeName string, q Query) Intent {
	s.RLock()
	defer s.RUnlock()
	is, ok := s.typeNameToIntents[typeName]
	if !ok {
		return nil
	}
	return q.load(is)
}

func (s *intentStore) delete(typeName string, q Query) {
	s.Lock()
	defer s.Unlock()
	is, ok := s.typeNameToIntents[typeName]
	if !ok {
		return
	}
	intent := q.load(is)
	if intent == nil {
		return
	}
	delete(is.fqNameToUUID, fqNameKey(intent.GetFQName()))
	delete(is.uuidToIntent, intent.GetUUID())
}

func (s *intentStore) store(typeName string, i Intent) {
	s.Lock()
	defer s.Unlock()
	is, ok := s.typeNameToIntents[typeName]
	if !ok {
		is = intents{
			fqNameToUUID: map[string]string{},
			uuidToIntent: map[string]Intent{},
		}
		s.typeNameToIntents[typeName] = is
	}
	is.uuidToIntent[i.GetUUID()] = i
	is.fqNameToUUID[fqNameKey(i.GetFQName())] = i.GetUUID()
}

func fqNameKey(fqName []string) string {
	return basemodels.FQNameToString(fqName)
}
