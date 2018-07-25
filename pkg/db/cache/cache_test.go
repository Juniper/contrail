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
	cache := New(1)
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

func TestDependencyResolution(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	cache := New(3)

	vn := models.MakeVirtualNetwork()
	vn.UUID = "vn_blue"

	event1, err := cache.Process(context.Background(), &services.Event{
		Version: 0,
		Request: &services.Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
				VirtualNetwork: vn,
			},
		},
	})

	assert.NoError(t, err)

	assert.Equal(t, "vn_blue", vn.UUID)
	e := cache.Get("vn_blue")
	assert.Equal(t, e, event1)
	assert.Equal(t, e.GetResource().GetParentUUID(), "")

	vn.ParentUUID = "hoge"
	event2, err := cache.Process(context.Background(), &services.Event{
		Version: 1,
		Request: &services.Event_UpdateVirtualNetworkRequest{
			UpdateVirtualNetworkRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: vn,
			},
		},
	})
	assert.Equal(t, e.GetResource().GetParentUUID(), "hoge")

	e = cache.Get("vn_blue")
	assert.Equal(t, e, event2)
	assert.NotEqual(t, event1, event2)

	ri := models.MakeRoutingInstance()
	ri.UUID = "ri_uuid1"
	ri.ParentUUID = vn.UUID

	_, err = cache.Process(context.Background(), &services.Event{
		Version: 2,
		Request: &services.Event_CreateRoutingInstanceRequest{
			CreateRoutingInstanceRequest: &services.CreateRoutingInstanceRequest{
				RoutingInstance: ri,
			},
		},
	})

	e = cache.Get("vn_blue")

	vn = e.GetUpdateVirtualNetworkRequest().GetVirtualNetwork()
	assert.Len(t, vn.RoutingInstances, 1)
	assert.Equal(t, vn.RoutingInstances[0].UUID, "ri_uuid1")

	ri = models.MakeRoutingInstance()
	ri.UUID = "ri_uuid2"
	ri.ParentUUID = vn.UUID
	ri.RoutingInstanceRefs = append(ri.RoutingInstanceRefs, &models.RoutingInstanceRoutingInstanceRef{UUID: "ri_uuid1"})

	_, err = cache.Process(context.Background(), &services.Event{
		Version: 2,
		Request: &services.Event_CreateRoutingInstanceRequest{
			CreateRoutingInstanceRequest: &services.CreateRoutingInstanceRequest{
				RoutingInstance: ri,
			},
		},
	})
	e = cache.Get("vn_blue")
	vn = e.GetUpdateVirtualNetworkRequest().GetVirtualNetwork()
	assert.Len(t, vn.RoutingInstances, 2)
	assert.Equal(t, vn.RoutingInstances[1].UUID, "ri_uuid2")

	e = cache.Get("ri_uuid1")
	ri = e.GetCreateRoutingInstanceRequest().GetRoutingInstance()
	assert.Len(t, ri.RoutingInstanceBackRefs, 1)
	assert.Equal(t, ri.RoutingInstanceBackRefs[0].UUID, "ri_uuid2")

	event4, err := cache.Process(context.Background(), &services.Event{
		Version: 3,
		Request: &services.Event_DeleteVirtualNetworkRequest{
			DeleteVirtualNetworkRequest: &services.DeleteVirtualNetworkRequest{
				ID: "vn_blue",
			},
		},
	})

	e = cache.Get("vn_blue")
	r := e.GetResource()
	assert.Equal(t, e, event4)
	assert.True(t, services.OperationDelete == e.Operation())
	assert.NotEqual(t, r.GetParentUUID(), vn.ParentUUID)
}
