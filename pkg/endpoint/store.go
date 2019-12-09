package endpoint

import (
	"errors"
	"reflect"
	"strings"
	"sync"

	"github.com/Juniper/asf/pkg/logutil"
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
func (t *TargetStore) Get(id string) *models.Endpoint {
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
func (t *TargetStore) GetAll(prefix string) (targets []*Endpoint) {
	if prefix != PublicURLScope && prefix != PrivateURLScope {
		return nil
	}

	t.Data.Range(func(id, endpoint interface{}) bool {
		e, ok := endpoint.(*models.Endpoint)
		if !ok {
			return true
		}

		targets = append(targets, newTargetEndpoint(e, prefix))
		return true
	})
	return targets
}

func newTargetEndpoint(e *models.Endpoint, prefix string) *Endpoint {
	switch prefix {
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

// Next reads next endpoint target from memory in round robin fashion.
func (t *TargetStore) Next(prefix string) (endpointData *Endpoint) {
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
		switch prefix {
		case PublicURLScope:
			if ids == t.nextTarget {
				e := endpoint.(*models.Endpoint) // nolint: errcheck
				endpointData = NewEndpoint(e.PublicURL, e.Username, e.Password)
				// let Range iterate till nextServer is identified
				return true
			}
		case PrivateURLScope:
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

func (e *Store) RemoveDeleted(endpoints map[string]*models.Endpoint) {
	log := logutil.NewLogger("endpoint-store")
	e.Data.Range(func(prefix, proxy interface{}) bool {
		ts, ok := proxy.(*TargetStore)
		if !ok {
			log.WithField("prefix", prefix).Error("Unable to read cluster's proxy data from in-memory store")
			return true
		}
		ts.Data.Range(func(id, endpoint interface{}) bool {
			_, ok := endpoint.(*models.Endpoint)
			if !ok {
				log.WithField("id", id).Error("Unable to Read endpoint data from in-memory store")
				return true
			}
			ids, ok := id.(string)
			if !ok {
				log.WithField("id", id).Error("Unable to convert id to string when looking endpointStore")
				return true
			}
			_, ok = endpoints[ids]
			if !ok {
				ts.Remove(ids)
				log.WithField("id", ids).Debug("Deleting dynamic proxy endpoint")
			}
			return true
		})
		if ts.Count() == 0 {
			prefixStr, ok := prefix.(string)
			if !ok {
				log.WithField("prefix", prefix).Error("Unable to convert prefix to string")
			}
			e.Remove(prefixStr)
			log.WithField("prefix", prefixStr).Debug("Deleting dynamic proxy endpoint prefix")
		}
		return true
	})
}

// Read reads endpoint targets store from memory.
func (e *Store) Get(endpointKey string) *TargetStore {
	p, ok := e.Data.Load(endpointKey)
	if !ok {
		return nil
	}

	ts, _ := p.(*TargetStore) // nolint: errcheck
	return ts
}

func (e *Store) GetEndpoints(prefix string, endpointKey string) []*Endpoint {
	targetStore := e.Get(endpointKey)
	if targetStore == nil {
		return nil
	}
	return targetStore.GetAll(prefix)
}

func (e *Store) GetUsername(prefix string, endpointKey string) string {
	targets := e.GetEndpoints(prefix, endpointKey)
	return targets[0].Username
}

func (e *Store) GetPassword(prefix string, endpointKey string) string {
	targets := e.GetEndpoints(prefix, endpointKey)
	return targets[0].Password
}

func (e *Store) GetEndpointURLs(prefix string, endpointKey string) []string {
	var u []string
	for _, target := range e.GetEndpoints(prefix, endpointKey) {
		u = append(u, target.URL)
	}
	return u
}

func (e *Store) Contains(prefix string) bool {
	_, ok := e.Data.Load(prefix)
	return ok
}

// Write writes endpoint targets store in-memory.
func (e *Store) Write(endpointKey string, endpointStore *TargetStore) {
	e.Data.Store(endpointKey, endpointStore)
}

func (e *Store) InitScope(prefix string) {
	if e.Get(prefix) == nil {
		e.Write(prefix, NewTargetStore())
	}
}

func (t *Store) UpdateEndpoint(prefix string, endpoint *models.Endpoint) error {
	ts := t.Get(prefix)
	if ts == nil {
		return errors.New("endpoint store for prefix is not found in-memory store")
	}
	e := ts.Get(endpoint.UUID)
	if !reflect.DeepEqual(e, endpoint) {
		ts.Write(endpoint.UUID, endpoint)
	}
	return nil
}

// Remove removes endpoint target store from memory.
func (e *Store) Remove(prefix string) {
	e.Data.Delete(prefix)
}

// GetEndpoint returns endpoint.
func (e *Store) GetEndpointURL(clusterID, prefix string) (string, bool) {
	if clusterID == "" {
		return "", false
	}
	// TODO(dfurman): "server.dynamic_proxy_path" or DefaultDynamicProxyPath should be used
	targets := e.Get(strings.Join([]string{"/proxy", clusterID, prefix, PrivateURLScope}, "/"))
	if targets == nil {
		return "", false
	}
	return targets.Next(PrivateURLScope).URL, true
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
