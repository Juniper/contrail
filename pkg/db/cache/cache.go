package cache

import (
	"sync"

	"github.com/Juniper/contrail/pkg/services"

	"github.com/Juniper/contrail/pkg/serviceif"
)

//Cache watch upstream db and dispatch it for service.
type Cache struct {
	db            sync.Map
	backReference sync.Map
	service       serviceif.Service
	watcher       Wather
}

//Watcher interface for watcher
type Watcher interface {
	//Listen listens upstream update and call hanlder func
	Listen(callback func(resource *services.ResourceEvent) error) error
}

//New makes cache module for the service.
func New(service serviceif.Service, watcher Wather) *Cache {
	cache := &Cache{
		db:            sync.Map{},
		backReference: sync.Map{},
		service:       service,
		watcher:       watcher,
	}
	watcher.Listen(cache.handler)
	return cache
}

//I don't like this implmenation because it has many overhead for cast.
//We need generate nessesary code to make this faster
func (cache *Cache) handler(resource *services.ResourceEvent) error {
	uuid := resource.UUID()
	switch resource.Operation {
	case services.EventCreate:
		cache.db.Store(uuid, resource)
		for _, refs := resource.Depend() {
			//[TODO] update backreference list
			//We may not use sync.Map here because we need append atomic way.
		}
	case services.EventUpdate:
		//[TODO] we need merge object
	case services.EventDelete:
		cache.db.Delete(uuid)
	}
	err := resource.Process(cache.service)
	//[TODO] need think about error handling of this error..
	//Depends should cache dependency list.
	//One idea is generate a service which resovles dependency call.
	for _, ref := resource.Depend() {
		resourceInterface := cache.db.Load(ref)
		resource, ok := resourceInterface.(*services.ResourceEvent)
		if !ok {
			//[TODO] error handling
			continue
		}
		resource.Operation = EventUpdate // or we should add EventRefUpdate
		err := resource.Process(cache.service)
		//[TODO] error handling
	}
}
