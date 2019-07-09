package common

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
	pathSep = "/"
)

// GetClusterIDFromProxyURL parses the proxy url to retrieve clusterID
func GetClusterIDFromProxyURL(url string) (clusterID string) {
	paths := strings.Split(url, pathSep)
	if len(paths) > 3 && paths[1] == "proxy" {
		clusterID = paths[2]
	}
	return clusterID
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

// TargetStore is used to store service specific endpoint targets in-memory
type TargetStore struct {
	Data       *sync.Map
	nextTarget string
}

// EndpointStore is used to store cluster specific endpoints in-memory
type EndpointStore struct {
	Data *sync.Map
}

//MakeTargetStore is used to make an in-memory endpoint target store.
func MakeTargetStore() *TargetStore {
	return &TargetStore{
		Data:       new(sync.Map),
		nextTarget: "",
	}
}

//MakeEndpointStore is used to make an in-memory endpoint store.
func MakeEndpointStore() *EndpointStore {
	return &EndpointStore{
		Data: new(sync.Map),
	}
}

//Read endpoint target from memory
func (t *TargetStore) Read(id string) *models.Endpoint {
	ep, ok := t.Data.Load(id)
	if !ok {
		return nil
	}
	endpoint, ok := ep.(*models.Endpoint)
	if !ok {
		return nil
	}
	return endpoint
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

//Read endpoint targets store from memory
func (e *EndpointStore) Read(endpointKey string) *TargetStore {
	p, ok := e.Data.Load(endpointKey)
	if !ok {
		return nil
	}
	endpointStore, _ := p.(*TargetStore) // nolint: errcheck
	return endpointStore

}

//Write endpoint targets store in-memory
func (e *EndpointStore) Write(endpointKey string, endpointStore *TargetStore) {
	e.Data.Store(endpointKey, endpointStore)
}

//Remove endpoint target store from memory
func (e *EndpointStore) Remove(prefix string) {
	e.Data.Delete(prefix)
}

//GetEndpoint by prefix
func (e *EndpointStore) GetEndpoint(clusterID, prefix string) (endpoint *Endpoint) {
	if clusterID != "" {
		targets := e.Read(
			strings.Join(
				[]string{"/proxy", clusterID, prefix, Private}, "/"))
		if targets == nil {
			return nil
		}
		return targets.Next(scope)
	}
	return endpoint
}
