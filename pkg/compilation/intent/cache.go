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
	uuidToType        map[string]string
	sync.RWMutex
}

type intents struct {
	fqNameToUUID map[string]string
	uuidToIntent map[string]Intent
}

func newIntentStore() *intentStore {
	return &intentStore{
		typeNameToIntents: make(map[string]intents),
		uuidToType:        make(map[string]string),
	}
}

func (s *intentStore) load(typeName string, q Query) Intent {
	s.RLock()
	defer s.RUnlock()
	return s.loadInternal(typeName, q)
}

func (s *intentStore) delete(typeName string, q Query) {
	s.Lock()
	defer s.Unlock()
	s.removeDependencies(typeName, q)
	s.deleteInternal(typeName, q)
}

func (s *intentStore) store(typeName string, i Intent) {
	s.Lock()
	defer s.Unlock()
	s.removeDependencies(typeName, ByUUID(i.GetUUID()))
	s.storeInternal(typeName, i)
	s.addDependencies(i.GetObject())
}

func (s *intentStore) loadInternal(typeName string, q Query) Intent {
	is, ok := s.typeNameToIntents[typeName]
	if !ok {
		return nil
	}
	return q.load(is)
}

func (s *intentStore) deleteInternal(typeName string, q Query) {
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
	delete(s.uuidToType, intent.GetUUID())
}

func (s *intentStore) storeInternal(typeName string, i Intent) {
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
	s.uuidToType[i.GetUUID()] = typeName
}

func (s *intentStore) addDependencies(resource basemodels.Object) {
	dependencies := resource.Depends()
	for _, dependencyID := range dependencies {
		t, ok := s.uuidToType[dependencyID]
		if !ok {
			continue
		}
		dependentIntent := s.loadInternal(t, ByUUID(dependencyID))
		if dependentIntent != nil {
			dependentIntent.AddDependency(resource)
		}
	}
}

func (s *intentStore) removeDependencies(typeName string, q Query) {
	i := s.loadInternal(typeName, q)
	if i != nil {
		dependencies := i.GetObject().Depends()
		for _, dependencyID := range dependencies {
			t, ok := s.uuidToType[dependencyID]
			if !ok {
				continue
			}
			dependentIntent := s.loadInternal(t, ByUUID(dependencyID))
			if ok {
				dependentIntent.RemoveDependency(i.GetObject())
			}
		}
	}
}

func fqNameKey(fqName []string) string {
	return basemodels.FQNameToString(fqName)
}
