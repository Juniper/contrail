package cache

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

const numEvent = 4
const timeOut = 10 * time.Second

func addWatcher(t *testing.T, wg *sync.WaitGroup, cache *DB) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	watcher, _ := cache.AddWatcher(ctx, 0)

	go func() {
		wg.Add(1)
		defer wg.Done()
		defer cancel()
		for i := 0; i < numEvent; i++ {
			select {
			case <-ctx.Done():
				log.Debugf("[watcher %d] time out on test", watcher.id)
				assert.Fail(t, "timeout")
			case e := <-watcher.ch:
				log.Debugf("[watcher %d] got event version %d", watcher.id, e.Version)
				assert.Equal(t, uint64(i), e.Version)
			}
		}
	}()
}

func notifyEvent(cache *DB, version uint64) {
	event := &services.Event{
		Version: version,
		Request: &services.Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
				VirtualNetwork: &models.VirtualNetwork{
					UUID: "vn" + strconv.FormatUint(version, 10),
				},
			},
		},
	}
	cache.Process(context.Background(), event) // nolint: errcheck
}

// nolint: unused
func notifyDelete(cache *DB, version uint64) {
	event := &services.Event{
		Version: version,
		Request: &services.Event_DeleteVirtualNetworkRequest{
			DeleteVirtualNetworkRequest: &services.DeleteVirtualNetworkRequest{
				ID: "vn" + strconv.FormatUint(version, 10),
			},
		},
	}
	cache.Process(context.Background(), event) // nolint: errcheck
}

func TestCache(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	cache := NewDB(1)
	wg := &sync.WaitGroup{}

	addWatcher(t, wg, cache)
	addWatcher(t, wg, cache)

	notifyEvent(cache, 0)
	notifyEvent(cache, 1)

	addWatcher(t, wg, cache)
	addWatcher(t, wg, cache)
	// test cancelation of channel.
	// expect no panic or blocking.
	ctx2, cancel := context.WithCancel(context.Background())
	cache.AddWatcher(ctx2, 0) // nolint: errcheck
	cancel()

	//timeout watcher
	//Don't actually receiving events.
	ctx3 := context.Background()
	cache.AddWatcher(ctx3, 0) // nolint: errcheck

	notifyEvent(cache, 2)
	notifyEvent(cache, 3)

	addWatcher(t, wg, cache)

	wg.Wait()

	// notifyDelete(t, cache, 0)
	// notifyDelete(t, cache, 1)
	// notifyDelete(t, cache, 2)
	// notifyDelete(t, cache, 3)

	// _, ok := cache.idMap["vn0"]
	// assert.Equal(t, false, ok, "compaction failed")
}

type dependencyTestAssertion func(
	t *testing.T,
	event *services.Event,
	result *services.Event,
	cache *DB,
)

type dependecyTestScenario struct {
	name      string
	event     *services.Event
	assertion dependencyTestAssertion
}

func TestDependencyResolution(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	cache := NewDB(2)

	riUUID1 := "ri_uuid1"
	riUUID2 := "ri_uuid2"

	vnBlue := &models.VirtualNetwork{
		UUID: "vn_blue",
	}
	vnWithParent := &models.VirtualNetwork{
		UUID:       vnBlue.UUID,
		Name:       "blue",
		ParentUUID: "domain",
	}
	ri1 := &models.RoutingInstance{
		UUID:       riUUID1,
		ParentUUID: vnBlue.GetUUID(),
	}
	ri2 := &models.RoutingInstance{
		UUID:       riUUID2,
		ParentUUID: vnBlue.GetUUID(),
		RoutingInstanceRefs: []*models.RoutingInstanceRoutingInstanceRef{
			{
				UUID: riUUID1,
			},
		},
	}

	tests := []dependecyTestScenario{
		{
			name: "Create vn_blue",
			event: &services.Event{
				Version: 0,
				Request: &services.Event_CreateVirtualNetworkRequest{
					CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
						VirtualNetwork: vnBlue,
					},
				},
			},
			assertion: func(t *testing.T, event *services.Event, result *services.Event, cache *DB) {
				e := cache.Get(vnBlue.UUID)
				assert.Equal(t, result, e)
				assert.Equal(t, "", e.GetResource().GetParentUUID())
			},
		},
		{
			name: "Update vn_blue with parent uuid",
			event: &services.Event{
				Version: 1,
				Request: &services.Event_UpdateVirtualNetworkRequest{
					UpdateVirtualNetworkRequest: &services.UpdateVirtualNetworkRequest{
						VirtualNetwork: vnWithParent,
					},
				},
			},
			assertion: func(t *testing.T, event *services.Event, result *services.Event, cache *DB) {
				e := cache.Get(vnBlue.UUID)
				assert.Equal(t, result, e)
				assert.Equal(t, services.OperationUpdate, e.Operation())
				assert.Equal(t, e.GetResource().GetParentUUID(), "domain")
			},
		},
		{
			name: "Create routing instance 1",
			event: &services.Event{
				Version: 2,
				Request: &services.Event_CreateRoutingInstanceRequest{
					CreateRoutingInstanceRequest: &services.CreateRoutingInstanceRequest{
						RoutingInstance: ri1,
					},
				},
			},
			assertion: func(t *testing.T, event *services.Event, result *services.Event, cache *DB) {
				e := cache.Get(vnBlue.UUID)
				v := e.GetUpdateVirtualNetworkRequest().GetVirtualNetwork()
				assert.Len(t, v.RoutingInstances, 1)
				assert.Equal(t, riUUID1, v.RoutingInstances[0].UUID)
			},
		},
		{
			name: "Create routing instance 2",
			event: &services.Event{
				Version: 3,
				Request: &services.Event_CreateRoutingInstanceRequest{
					CreateRoutingInstanceRequest: &services.CreateRoutingInstanceRequest{
						RoutingInstance: ri2,
					},
				},
			},
			assertion: func(t *testing.T, event *services.Event, result *services.Event, cache *DB) {
				e := cache.Get(vnBlue.UUID)
				v := e.GetUpdateVirtualNetworkRequest().GetVirtualNetwork()
				assert.Len(t, v.RoutingInstances, 2)
				assert.Equal(t, riUUID2, v.RoutingInstances[1].UUID)

				e = cache.Get(riUUID1)
				ri := e.GetCreateRoutingInstanceRequest().GetRoutingInstance()
				assert.Len(t, ri.RoutingInstanceBackRefs, 1)
				assert.Equal(t, riUUID2, ri.RoutingInstanceBackRefs[0].UUID)
			},
		},
		{
			name: "Delete routing instance 2",
			event: &services.Event{
				Version: 4,
				Request: &services.Event_DeleteRoutingInstanceRequest{
					DeleteRoutingInstanceRequest: &services.DeleteRoutingInstanceRequest{
						ID: riUUID2,
					},
				},
			},
			assertion: func(t *testing.T, event *services.Event, result *services.Event, cache *DB) {
				e := cache.Get(vnBlue.UUID)
				v := e.GetUpdateVirtualNetworkRequest().GetVirtualNetwork()
				assert.Len(t, v.RoutingInstances, 1)
				assert.Equal(t, riUUID1, v.RoutingInstances[0].UUID)

				e = cache.Get(riUUID1)
				ri := e.GetCreateRoutingInstanceRequest().GetRoutingInstance()
				assert.Len(t, ri.RoutingInstanceBackRefs, 0)
			},
		},
		{
			name: "Delete routing instance 1",
			event: &services.Event{
				Version: 5,
				Request: &services.Event_DeleteRoutingInstanceRequest{
					DeleteRoutingInstanceRequest: &services.DeleteRoutingInstanceRequest{
						ID: riUUID1,
					},
				},
			},
			assertion: func(t *testing.T, event *services.Event, result *services.Event, cache *DB) {
				e := cache.Get(vnBlue.UUID)
				vnBlue = e.GetUpdateVirtualNetworkRequest().GetVirtualNetwork()
				assert.Len(t, vnBlue.RoutingInstances, 0)
			},
		},
		{
			name: "Delete virtual network",
			event: &services.Event{
				Version: 6,
				Request: &services.Event_DeleteVirtualNetworkRequest{
					DeleteVirtualNetworkRequest: &services.DeleteVirtualNetworkRequest{
						ID: vnBlue.UUID,
					},
				},
			},
			assertion: func(t *testing.T, event *services.Event, result *services.Event, cache *DB) {
				e := cache.Get(vnBlue.UUID)
				r := e.GetResource()
				assert.Equal(t, result, e)
				assert.Equal(t, services.OperationDelete, e.Operation())
				assert.NotEqual(t, vnBlue.ParentUUID, r.GetParentUUID())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := cache.Process(context.Background(), tt.event)
			assert.NoError(t, err)
			assert.Equal(t, e, tt.event)
			tt.assertion(t, tt.event, e, cache)
		})
	}
}
