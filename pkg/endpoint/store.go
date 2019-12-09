package endpoint

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

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
}

// NewStore returns new Store.
func NewStore() *Store {
	return &Store{
		Data: new(sync.Map),
	}
}

func (e *Store) Range(f func(scope, value interface{}) bool) {
	e.Data.Range(f)
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

func (e *Store) ReadEndpoints(scope string, endpointKey string) []*Endpoint {

	targetStore := e.Read(endpointKey)
	if targetStore == nil {
		return nil
	}

	return targetStore.ReadAll(scope)
}

func (e *Store) ReadEndpointsUrls(scope string, endpointKey string) []string {

	targetStore := e.Read(endpointKey)
	if targetStore == nil {
		return nil
	}

	targets := targetStore.ReadAll(scope)

	var u []string
	for _, target := range targets {
		u = append(u, target.URL)
	}
	return u
}

func (e *Store) InStore(scope string) bool {

	_, ok := e.Data.Load(scope)
	return ok
}

// Write writes endpoint targets store in-memory.
func (e *Store) Write(endpointKey string, endpointStore *TargetStore) {
	e.Data.Store(endpointKey, endpointStore)
}

func (e *Store) InitScope(scope string) {
	targetStore := e.Read(scope)
	if targetStore == nil {
		e.Write(scope, NewTargetStore())
	}
}

func (t *Store) UpdateEndpoint(scope string, endpoint *models.Endpoint) error {

	targetStore := t.Read(scope)
	if targetStore == nil {
		return errors.New("Endpoint store for prefix is not found in-memory store")
	}

	e := targetStore.Read(endpoint.UUID)
	if !reflect.DeepEqual(e, endpoint) {
		// proxy endpoint not in memory store or
		// proxy endpoint updated
		targetStore.Write(endpoint.UUID, endpoint)
	}
	return nil
}

// Remove removes endpoint target store from memory.
func (e *Store) Remove(prefix string) {
	e.Data.Delete(prefix)
}

func (e *Store) RemoveDeleted(scope, proxy interface{}, endpoints map[string]*models.Endpoint, log *logrus.Entry) bool {

	s, ok := proxy.(*TargetStore)
	if !ok {
		log.WithField("prefix", scope).Error("Unable to read cluster's proxy data from in-memory store")
		return true
	}
	s.Data.Range(func(id, endpoint interface{}) bool {
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
			s.Remove(ids)
			log.WithField("id", ids).Debug("Deleting dynamic proxy endpoint")
		}
		return true
	})
	if s.Count() == 0 {
		prefixStr, ok := scope.(string)
		if !ok {
			log.WithField("prefix", scope).Error("Unable to convert prefix to string")
		}
		e.Remove(prefixStr)
		log.WithField("prefix", prefixStr).Debug("Deleting dynamic proxy endpoint prefix")
	}
	return true
}

// GetEndpoint returns endpoint.
func (e *Store) GetEndpoint(clusterID, prefix string) (*Endpoint, bool) {
	if clusterID == "" {
		return nil, false
	}
	// TODO(dfurman): "server.dynamic_proxy_path" or DefaultDynamicProxyPath should be used
	targets := e.Read(strings.Join([]string{"/proxy", clusterID, prefix, PrivateURLScope}, "/"))
	if targets == nil {
		return nil, false
	}
	return targets.Next(PrivateURLScope), true
}

func (e *Store) GetAuthEndpoint(scope string, endpointKey string) (*Endpoint, error) {
	keystoneTargets := e.Read(endpointKey)
	if keystoneTargets == nil {
		err := fmt.Errorf("keystone targets not found for: %s", endpointKey)
		return nil, err
	}
	authEndpoint := keystoneTargets.Next(scope)
	if authEndpoint == nil {
		err := fmt.Errorf("unable to get keystone endpoint for: %s", endpointKey)
		return nil, err
	}
	return authEndpoint, nil
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
