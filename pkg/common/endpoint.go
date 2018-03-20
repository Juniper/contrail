package common

import "sync"

// EndpointStore is used to store cluster specific endpoints in-memory
type EndpointStore struct {
	Data *sync.Map
}

//MakeEndpointStore is used to make a in memory endpoint store.
func MakeEndpointStore() *EndpointStore {
	return &EndpointStore{
		Data: new(sync.Map),
	}
}
