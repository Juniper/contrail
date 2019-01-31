package common

import (
	"fmt"
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
func (t *TargetStore) Next(scope string) (endpointURL string) {
	endpointURL = ""
	t.Data.Range(func(id, endpoint interface{}) bool {
		ids := id.(string) // nolint: errcheck
		if t.nextTarget == "" {
			t.nextTarget = ids
		}
		if endpointURL != "" {
			t.nextTarget = ids
			// exit Range iteration as next target is identified
			return false
		}
		switch scope {
		case Public:
			if ids == t.nextTarget {
				endpointURL = endpoint.(*models.Endpoint).PublicURL
				// let Range iterate till nextServer is identified
				return true
			}
		case Private:
			if ids == t.nextTarget {
				endpointURL = endpoint.(*models.Endpoint).PrivateURL
				if endpointURL == "" {
					// no private url configured, use public url
					endpointURL = endpoint.(*models.Endpoint).PublicURL
				}
				// let Range iterate till nextServer is identified
				return true
			}
		}
		return true
	})
	return endpointURL
}

//Remove endpoint target from memory
func (t *TargetStore) Remove(endpointKey string) {
	if endpointKey == t.nextTarget {
		// Reset the next target before deleting the endpoint
		t.nextTarget = ""
	}
	t.Data.Delete(endpointKey)
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

//GetEndpoint by prefix
func (e *EndpointStore) GetEndpoint(prefix string) (endpoint string, err error) {
	endpointCount := 0
	endpoint = ""
	e.Data.Range(func(key, targets interface{}) bool {
		keyString, _ := key.(string) // nolint: errcheck
		keyParts := strings.Split(keyString, pathSep)
		if keyParts[3] != prefix || keyParts[4] != Private {
			endpoints, _ := targets.(*TargetStore) // nolint: errcheck
			endpoint = endpoints.Next(Private)
			if endpoint != "" {
				endpointCount++
			}
			//return true // continue iterating the endpoints
		}
		if endpointCount > 1 {
			err = fmt.Errorf("ambiguious, more than one cluster found")
			return false
		}
		//endpoints, _ := targets.(*TargetStore) // nolint: errcheck
		//endpoint = endpoints.Next(Private)
		return true

	})

	return endpoint, err
}
