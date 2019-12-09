package endpoint

import (
	"errors"
	"reflect"
	"strings"
	"sync"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/sirupsen/logrus"
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

// Get endpoint target from memory.
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

// GetAll reads all endpoint targets from target store using given scope.
// Order of returned targets is not deterministic.
func (t *TargetStore) GetAll(scope string) (targets []*Endpoint) {
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

// Next reads next endpoint target from memory in round robin fashion.
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
	log  *logrus.Entry
}

// NewStore returns new Store.
func NewStore() *Store {
	return &Store{
		Data: new(sync.Map),
		log:  logutil.NewLogger("endpoint-store"),
	}
}

// RemoveDeleted removes all endpoints that do not match the given ones.
func (s *Store) RemoveDeleted(endpoints map[string]*models.Endpoint) {
	s.Data.Range(func(prefix, targetStore interface{}) bool {
		ts, ok := targetStore.(*TargetStore)
		if !ok {
			s.log.WithField("prefix", prefix).Error("Unable to read cluster's proxy data from in-memory store")
			return true
		}
		ts.Data.Range(func(id, endpoint interface{}) bool {
			_, ok := endpoint.(*models.Endpoint)
			if !ok {
				s.log.WithField("id", id).Error("Unable to Read endpoint data from in-memory store")
				return true
			}
			ids, ok := id.(string)
			if !ok {
				s.log.WithField("id", id).Error("Unable to convert id to string when looking endpointStore")
				return true
			}
			_, ok = endpoints[ids]
			if !ok {
				ts.Remove(ids)
				s.log.WithField("id", ids).Debug("Deleting dynamic proxy endpoint")
			}
			return true
		})
		if ts.Count() == 0 {
			prefixStr, ok := prefix.(string)
			if !ok {
				s.log.WithField("prefix", prefix).Error("Unable to convert prefix to string")
			}
			s.Remove(prefixStr)
			s.log.WithField("prefix", prefixStr).Debug("Deleting dynamic proxy endpoint prefix")
		}
		return true
	})
}

// Get reads endpoint targets store from memory.
func (s *Store) Get(endpointKey string) *TargetStore {
	p, ok := s.Data.Load(endpointKey)
	if !ok {
		return nil
	}

	ts, _ := p.(*TargetStore) // nolint: errcheck
	return ts
}

func (s *Store) getEndpoints(prefix string, endpointKey string) []*Endpoint {
	targetStore := s.Get(endpointKey)
	if targetStore == nil {
		return nil
	}
	return targetStore.GetAll(prefix)
}

// GetUsername reads the username of the first endpoint target of the given store.
func (s *Store) GetUsername(prefix string, endpointKey string) string {
	targets := s.getEndpoints(prefix, endpointKey)
	if len(targets) == 0 {
		return ""
	}
	return targets[0].Username
}

// GetPassword reads the password of the first endpoint target of the given store.
func (s *Store) GetPassword(prefix string, endpointKey string) string {
	targets := s.getEndpoints(prefix, endpointKey)
	if len(targets) == 0 {
		return ""
	}
	return targets[0].Password
}

// GetEndpointURLs reads the urls of endpoint targets of the given store.
func (s *Store) GetEndpointURLs(prefix string, endpointKey string) []string {
	var u []string
	for _, target := range s.getEndpoints(prefix, endpointKey) {
		u = append(u, target.URL)
	}
	return u
}

// Contains check if the prefix exists in store.
func (s *Store) Contains(prefix string) bool {
	_, ok := s.Data.Load(prefix)
	return ok
}

// Write writes endpoint targets store in-memory.
func (s *Store) Write(endpointKey string, endpointStore *TargetStore) {
	s.Data.Store(endpointKey, endpointStore)
}

// InitScope creates a new target store for a given prefix if it doesn't yet exist.
func (s *Store) InitScope(prefix string) {
	if s.Get(prefix) == nil {
		s.Write(prefix, NewTargetStore())
	}
}

// UpdateEndpoint writes endpoint target in memory if proxy endpoint not in store or updated.
func (s *Store) UpdateEndpoint(prefix string, endpoint *models.Endpoint) error {
	ts := s.Get(prefix)
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
func (s *Store) Remove(prefix string) {
	s.Data.Delete(prefix)
}

// GetEndpointURL returns endpoint.
func (s *Store) GetEndpointURL(clusterID, prefix string) (string, bool) {
	if clusterID == "" {
		return "", false
	}
	// TODO(dfurman): "server.dynamic_proxy_path" or DefaultDynamicProxyPath should be used
	targets := s.Get(strings.Join([]string{"/proxy", clusterID, prefix, PrivateURLScope}, "/"))
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
