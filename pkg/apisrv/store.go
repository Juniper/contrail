package apisrv

import (
	"reflect"
	"sync"

	"github.com/Juniper/contrail/pkg/models"
)

// Endpoint related constants.
const (
	PublicURLScope  = "public"
	PrivateURLScope = "private"
)

// TargetStore is used to store service specific endpoint targets in-memory.
type TargetStore struct {
	Data       *sync.Map
	nextTarget string
}

// NewTargetStore returns new TargetStore.
func NewTargetStore() *TargetStore {
	return &TargetStore{
		Data:       new(sync.Map),
		nextTarget: "",
	}
}

// Read endpoint target from memory.
func (t *TargetStore) Read(id string) *models.Endpoint {
	rawE, ok := t.Data.Load(id)
	if !ok {
		return nil
	}
	e, ok := rawE.(*models.Endpoint)
	if !ok {
		return nil
	}
	return e
}

// ReadAll reads all endpoint targets from target store using given scope.
// Order of returned targets is not deterministic.
func (t *TargetStore) ReadAll(scope string) (targets []*Endpoint) {
	if scope != PublicURLScope && scope != PrivateURLScope {
		return nil
	}

	t.Data.Range(func(id, endpoint interface{}) bool {
		e, ok := endpoint.(*models.Endpoint)
		if !ok {
			return true
		}

		targets = append(targets, newTargetEndpoint(e, scope))
		return true
	})
	return targets
}

func newTargetEndpoint(e *models.Endpoint, scope string) *Endpoint {
	switch scope {
	case PublicURLScope:
		return NewEndpoint(e.PublicURL, e.Username, e.Password)
	case PrivateURLScope:
		if e.PrivateURL != "" {
			return NewEndpoint(e.PrivateURL, e.Username, e.Password)
		}
		return NewEndpoint(e.PublicURL, e.Username, e.Password)
	}
	return nil
}

// Write writes endpoint target in-memory.
func (t *TargetStore) Write(id string, endpoint *models.Endpoint) {
	t.Data.Store(id, endpoint)
}

// Remove removes endpoint target from memory.
func (t *TargetStore) Remove(endpointKey string) {
	if endpointKey == t.nextTarget {
		// Reset the next target before deleting the endpoint
		t.nextTarget = ""
	}
	t.Data.Delete(endpointKey)
}

// Count returns number of endpoint targets in memory.
func (t *TargetStore) Count() int {
	count := 0
	t.Data.Range(func(id, endpoint interface{}) bool {
		count++
		return true
	})
	return count
}

// Store is used to store cluster specific endpoint targets store in-memory.
type Store struct {
	Data *sync.Map
}

// NewStore returns new Store.
func NewStore() *Store {
	return &Store{
		Data: new(sync.Map),
	}
}

func (e *Store) GetData() *sync.Map {
	return e.Data
}

// Read reads endpoint targets store from memory.
func (e *Store) Read(endpointKey string) *TargetStore {
	p, ok := e.Data.Load(endpointKey)
	if !ok {
		return nil
	}

	ts, _ := p.(*TargetStore) // nolint: errcheck
	return ts
}

func (e *Store) InStore(endpointKey string) bool {

	_, ok := e.Data.Load(endpointKey)
	return ok
}

func (t *Store) UpdateEndpoint(endpointKey string, endpoint *models.Endpoint) {

	targetStore := t.Read(endpointKey)
	if targetStore == nil {
		return
	}

	e := targetStore.Read(endpoint.UUID)
	if !reflect.DeepEqual(e, endpoint) {
		// proxy endpoint not in memory store or
		// proxy endpoint updated
		targetStore.Write(endpoint.UUID, endpoint)
	}
}

func (e *Store) ReadEndpoints(scope string, endpointKey string) []*Endpoint {

	targetStore := e.Read(endpointKey)
	if targetStore == nil {
		return nil
	}

	targets := targetStore.ReadAll(scope)
	return targets
}

// Write writes endpoint targets store in-memory.
func (e *Store) Write(endpointKey string, endpointStore *TargetStore) {
	e.Data.Store(endpointKey, endpointStore)
}

func (e *Store) WriteNewTargetStore(endpointKey string) {
	targetStore := e.Read(endpointKey)
	if targetStore == nil {
		e.Write(endpointKey, NewTargetStore())
	}
}

// Remove removes endpoint target store from memory.
func (e *Store) Remove(prefix string) {
	e.Data.Delete(prefix)
}

// Endpoint represents an endpoint url with its credentials.
type Endpoint struct {
	URL      string
	Username string
	Password string
}

// NewEndpoint returns new Endpoint.
func NewEndpoint(url, user, password string) *Endpoint {
	return &Endpoint{
		URL:      url,
		Username: user,
		Password: password,
	}
}
