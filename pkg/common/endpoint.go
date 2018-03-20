package common

import "sync"

// EndpointStore is used to store cluster specific endpoints in-memory
type EndpointStore struct {
	Data *sync.Map
}
