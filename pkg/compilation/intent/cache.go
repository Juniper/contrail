package intent

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"

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
	logrus.WithFields(logrus.Fields{"kind": kind, "query": q}).Debug("Loading from cache")
	return c.intentStore.load(kind, q)
}

// Store puts intent into cache.
func (c *Cache) Store(i Intent) {
	logrus.WithFields(logrus.Fields{"kind": i.Kind(), "uuid": i.GetUUID()}).Debug("Storing in cache")
	c.intentStore.store(i.Kind(), i)
}

// Delete deletes intent from cache. It accepts kebab-case or CamelCase type name.
func (c *Cache) Delete(kind string, q Query) {
	logrus.WithFields(logrus.Fields{"kind": kind, "query": q}).Debug("Deleting from cache")
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
	backRefs := i.GetBackReferences()
	children := i.GetChildren()
	s.removeDependencies(typeName, ByUUID(i.GetUUID()))
	s.storeInternal(i)
	s.addDependencies(i, backRefs, children)
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

func (s *intentStore) storeInternal(i Intent) {
	is, ok := s.typeNameToIntents[i.Kind()]
	if !ok {
		is = intents{
			fqNameToUUID: map[string]string{},
			uuidToIntent: map[string]Intent{},
		}
		s.typeNameToIntents[i.Kind()] = is
	}
	is.uuidToIntent[i.GetUUID()] = i
	is.fqNameToUUID[fqNameKey(i.GetFQName())] = i.GetUUID()
	s.uuidToType[i.GetUUID()] = i.Kind()
}

func (s *intentStore) addDependencies(i Intent, backRefs, children []basemodels.Object) {
	for _, backRef := range backRefs {
		i.AddBackReference(backRef)
	}
	for _, child := range children {
		i.AddChild(child)
	}
	for _, ref := range i.GetReferences() {
		t, ok := s.uuidToType[ref.GetUUID()]
		if !ok {
			continue
		}
		dependentIntent := s.loadInternal(t, ByUUID(ref.GetUUID()))
		if dependentIntent != nil {
			i.AddDependentIntent(dependentIntent)
			dependentIntent.AddBackReference(i.GetObject())
			dependentIntent.AddDependentIntent(i)
			s.storeInternal(dependentIntent)
		}
	}
	if i.GetParentUUID() != "" {
		t, ok := s.uuidToType[i.GetParentUUID()]
		if !ok {
			return
		}
		dependentIntent := s.loadInternal(t, ByUUID(i.GetParentUUID()))
		if dependentIntent != nil {
			i.AddDependentIntent(dependentIntent)
			dependentIntent.AddChild(i.GetObject())
			dependentIntent.AddDependentIntent(i)
			s.storeInternal(dependentIntent)
		}
	}
}

func (s *intentStore) removeDependencies(typeName string, q Query) {
	i := s.loadInternal(typeName, q)
	if i == nil {
		return
	}
	for _, backRef := range i.GetObject().GetBackReferences() {
		i.RemoveBackReference(backRef)
	}
	for _, child := range i.GetObject().GetChildren() {
		i.RemoveChild(child)
	}
	for _, ref := range i.GetObject().GetReferences() {
		t, ok := s.uuidToType[ref.GetUUID()]
		if !ok {
			continue
		}
		dependentIntent := s.loadInternal(t, ByUUID(ref.GetUUID()))
		if dependentIntent != nil {
			i.RemoveDependentIntent(dependentIntent)
			dependentIntent.RemoveBackReference(i.GetObject())
			dependentIntent.RemoveDependentIntent(i)
		}
	}
	if i.GetParentUUID() != "" {
		t, ok := s.uuidToType[i.GetParentUUID()]
		if !ok {
			return
		}
		dependentIntent := s.loadInternal(t, ByUUID(i.GetParentUUID()))
		if dependentIntent != nil {
			i.RemoveDependentIntent(dependentIntent)
			dependentIntent.RemoveChild(i.GetObject())
			dependentIntent.RemoveDependentIntent(i)
		}
	}
}

func fqNameKey(fqName []string) string {
	return basemodels.FQNameToString(fqName)
}
