package intent

import (
	"fmt"
	"sync"

	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// Cache cache for intents.
type Cache struct {
	*intentStore
}

// ID identifies an intent in cache.
type ID struct {
	load func(intents) Intent
	str  string
}

// String returns the string representation of an ID.
func (id ID) String() string {
	return id.str
}

// UUID identifies an intent in cache by UUID.
func UUID(uuid string) ID {
	return ID{
		load: func(is intents) Intent {
			return is.idToIntent[uuid]
		},
		str: uuid,
	}
}

// FQName identifies an intent in cache by FQName.
func FQName(fqName []string) ID {
	return ID{
		load: func(is intents) Intent {
			return is.idToIntent[is.fqNameToID[fqNameKey(fqName)]]
		},
		str: fmt.Sprint(fqName),
	}
}

// NewCache creates new cache for intents.
func NewCache() *Cache {
	return &Cache{
		intentStore: newIntentStore(),
	}
}

// Load loads intent from cache. It accepts as type both kebab-case and CamelCase.
func (c *Cache) Load(typeName string, id ID) Intent {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"TypeName": typeName, "ID": id}).Debug("Loading from cache")
	return c.intentStore.load(typeName, id)
}

// Store puts intent into cache.
func (c *Cache) Store(i Intent) {
	typeName := strcase.ToCamel(i.Kind())
	log.WithFields(log.Fields{"TypeName": typeName, "UUID": i.GetUUID()}).Debug("Storing in cache")
	c.intentStore.store(typeName, i)
}

// Delete deletes intent from cache. It accepts as type both kebab-case and CamelCase.
func (c *Cache) Delete(typeName string, id ID) {
	typeName = strcase.ToCamel(typeName)
	log.WithFields(log.Fields{"TypeName": typeName, "ID": id}).Debug("Deleting from cache")
	c.intentStore.delete(typeName, id)
}

type intentStore struct {
	typeToIntents map[string]intents
	sync.RWMutex
}

type intents struct {
	fqNameToID map[string]string
	idToIntent map[string]Intent
}

func newIntentStore() *intentStore {
	return &intentStore{
		typeToIntents: make(map[string]intents),
	}
}

func (s *intentStore) load(typeName string, id ID) Intent {
	s.RLock()
	defer s.RUnlock()
	is, ok := s.typeToIntents[typeName]
	if !ok {
		return nil
	}
	return id.load(is)
}

func (s *intentStore) delete(typeName string, id ID) {
	s.Lock()
	defer s.Unlock()
	is, ok := s.typeToIntents[typeName]
	if !ok {
		return
	}
	i := id.load(is)
	if i == nil {
		return
	}
	delete(is.fqNameToID, fqNameKey(i.GetFQName()))
	delete(is.idToIntent, i.GetUUID())
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
