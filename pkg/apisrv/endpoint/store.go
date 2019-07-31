package endpoint

import (
	"strings"
	"sync"

	"github.com/Juniper/contrail/pkg/models"
)

const (
	// Public scope of the endpoint url
	Public = "public"
	// Private scope of the endpoint url
	Private = "private"
)

// TargetStore is used to store service specific endpoint targets in-memory
type TargetStore struct {
	Data       *sync.Map
	nextTarget string
}

// NewTargetStore is used to make an in-memory endpoint target store.
func NewTargetStore() *TargetStore {
	return &TargetStore{
		Data:       new(sync.Map),
		nextTarget: "",
	}
}

//Read endpoint target from memory
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

// ReadAll reads all targets from target store using given scope.
// TODO: return all endpoints
func (t *TargetStore) ReadAll(scope string) (targets []*Endpoint) {
	var nextEndpoint *Endpoint
	t.Data.Range(func(id, endpoint interface{}) bool {
		e, ok := endpoint.(*models.Endpoint)
		if !ok {
			return false
		}

		if t.nextTarget == e.UUID {
			nextEndpoint = newTargetEndpoint(e, scope)
		} else {
			targets = append(targets, newTargetEndpoint(e, scope))
		}
		return true
	})

	// Return the next target as first entry in the list so that the proxy service will load balance the
	// requests among available endpoints starting from the next target.
	if nextEndpoint != nil {
		targets = append([]*Endpoint{nextEndpoint}, targets...)
	}
	return targets
}

func newTargetEndpoint(e *models.Endpoint, scope string) *Endpoint {
	var t *Endpoint
	switch scope {
	case Public:
		t = NewEndpoint(e.PublicURL, e.Username, e.Password)
	case Private:
		if e.PrivateURL != "" {
			t = NewEndpoint(e.PrivateURL, e.Username, e.Password)
		} else {
			t = NewEndpoint(e.PublicURL, e.Username, e.Password)
		}
	}
	return t
}

//Write endpoint target in-memory
func (t *TargetStore) Write(id string, endpoint *models.Endpoint) {
	t.Data.Store(id, endpoint)
}

//Next endpoint target from memory is read(roundrobin)
func (t *TargetStore) Next(scope string) (endpointData *Endpoint) {
	t.Data.Range(func(id, endpoint interface{}) bool {
		ids := id.(string) // nolint: errcheck
		if t.nextTarget == "" {
			t.nextTarget = ids
		}
		if endpointData != nil {
			t.nextTarget = ids
			// exit Range iteration as next target is identified
			return false
		}
		switch scope {
		case Public:
			if ids == t.nextTarget {
				e := endpoint.(*models.Endpoint) // nolint: errcheck
				endpointData = NewEndpoint(e.PublicURL, e.Username, e.Password)
				// let Range iterate till nextServer is identified
				return true
			}
		case Private:
			if ids == t.nextTarget {
				e := endpoint.(*models.Endpoint) // nolint: errcheck
				endpointData = NewEndpoint(e.PrivateURL, e.Username, e.Password)
				if endpointData == nil {
					// no private url configured, use public url
					e := endpoint.(*models.Endpoint) // nolint: errcheck
					endpointData = NewEndpoint(e.PublicURL, e.Username, e.Password)
				}
				// let Range iterate till nextServer is identified
				return true
			}
		}
		return true
	})
	return endpointData
}

//Remove endpoint target from memory
func (t *TargetStore) Remove(endpointKey string) {
	if endpointKey == t.nextTarget {
		// Reset the next target before deleting the endpoint
		t.nextTarget = ""
	}
	t.Data.Delete(endpointKey)
}

//Count endpoint target from memory
func (t *TargetStore) Count() int {
	count := 0
	t.Data.Range(func(id, endpoint interface{}) bool {
		count++
		return true
	})
	return count
}

// Store is used to store cluster specific endpoints in-memory
type Store struct {
	Data *sync.Map
}

// NewStore is used to make an in-memory endpoint store.
func NewStore() *Store {
	return &Store{
		Data: new(sync.Map),
	}
}

//Read endpoint targets store from memory
func (e *Store) Read(endpointKey string) *TargetStore {
	p, ok := e.Data.Load(endpointKey)
	if !ok {
		return nil
	}

	ts, _ := p.(*TargetStore) // nolint: errcheck
	return ts
}

//Write endpoint targets store in-memory
func (e *Store) Write(endpointKey string, endpointStore *TargetStore) {
	e.Data.Store(endpointKey, endpointStore)
}

//Remove endpoint target store from memory
func (e *Store) Remove(prefix string) {
	e.Data.Delete(prefix)
}

//GetEndpoint by prefix
func (e *Store) GetEndpoint(clusterID, prefix string) (endpoint *Endpoint) {
	if clusterID == "" {
		return nil
	}
	targets := e.Read(strings.Join([]string{"/proxy", clusterID, prefix, Private}, "/"))
	if targets == nil {
		return nil
	}
	return targets.Next(Private)
}

// Endpoint represents an endpoint url with its credentials
type Endpoint struct {
	URL      string
	Username string
	Password string
}

// NewEndpoint returns endoint struct with credential
func NewEndpoint(url, user, password string) *Endpoint {
	return &Endpoint{
		URL:      url,
		Username: user,
		Password: password,
	}
}
