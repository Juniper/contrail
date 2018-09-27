package intent

import (
	"fmt"
	"sync"

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

// Load loads intent from cache.
func (c *Cache) Load(kind string, q Query) Intent {
	log.WithFields(log.Fields{"kind": kind, "query": q}).Debug("Loading from cache")
	return c.intentStore.load(kind, q)
}

// Store puts intent into cache.
func (c *Cache) Store(i Intent) {
	log.WithFields(log.Fields{"kind": i.Kind(), "uuid": i.GetUUID()}).Debug("Storing in cache")
	c.intentStore.store(i.Kind(), i)
}

// Delete deletes intent from cache. It accepts kebab-case or CamelCase type name.
func (c *Cache) Delete(kind string, q Query) {
	log.WithFields(log.Fields{"kind": kind, "query": q}).Debug("Deleting from cache")
	c.intentStore.delete(kind, q)
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
	s.addDependencies(i)
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

func (s *intentStore) addDependencies(i Intent) {
	for _, ref := range i.Depends() {
		s.addDependency(ref.GetUUID(), i)
	}
	if i.GetParentUUID() != "" {
		s.addDependency(i.GetParentUUID(), i)
	}
}

func (s *intentStore) removeDependencies(typeName string, q Query) {
	i := s.loadInternal(typeName, q)
	if i != nil {
		for _, ref := range i.GetObject().Depends() {
			s.removeDependency(ref.GetUUID(), i)
		}
		if i.GetParentUUID() != "" {
			s.removeDependency(i.GetParentUUID(), i)
		}
	}
}

func (s *intentStore) addDependency(uuid string, i Intent) {
	t, ok := s.uuidToType[uuid]
	if !ok {
		return
	}
	dependentIntent := s.loadInternal(t, ByUUID(uuid))
	if dependentIntent != nil {
		i.AddDependentIntent(dependentIntent)
		dependentIntent.AddDependency(i.GetObject())
		dependentIntent.AddDependentIntent(i)
	}
}

func (s *intentStore) removeDependency(uuid string, i Intent) {
	t, ok := s.uuidToType[uuid]
	if !ok {
		return
	}
	dependentIntent := s.loadInternal(t, ByUUID(uuid))
	if dependentIntent != nil {
		i.RemoveDependentIntent(dependentIntent)
		dependentIntent.RemoveDependency(i.GetObject())
		dependentIntent.RemoveDependentIntent(i)
	}
}

func fqNameKey(fqName []string) string {
	return basemodels.FQNameToString(fqName)
}
