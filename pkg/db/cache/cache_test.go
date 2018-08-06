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

func addWatcher(t *testing.T, wg *sync.WaitGroup, cache *DBCache) {
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

func notifyEvent(cache *DBCache, version uint64) {
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
func notifyDelete(cache *DBCache, version uint64) {
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
	cache := NewDBCache(1)
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
	cache := NewDBCache(6)

	vnTestUUID := "vn_blue"
	vn := makeTestVirtualNetwork(&baseResourceParams{uuid: vnTestUUID})

	event1, err := cache.processTestEvent(&services.Event{
		Version: 0,
		Request: &services.Event_CreateVirtualNetworkRequest{
			CreateVirtualNetworkRequest: &services.CreateVirtualNetworkRequest{
				VirtualNetwork: vn,
			},
		},
	})
	assert.NoError(t, err)

	assert.Equal(t, vnTestUUID, vn.UUID)
	e := cache.Get(vnTestUUID)
	assert.Equal(t, event1, e)
	assert.Equal(t, "", e.GetResource().GetParentUUID())

	vn.ParentUUID = "domain"
	event2, err := cache.processTestEvent(&services.Event{
		Version: 1,
		Request: &services.Event_UpdateVirtualNetworkRequest{
			UpdateVirtualNetworkRequest: &services.UpdateVirtualNetworkRequest{
				VirtualNetwork: vn,
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, e.GetResource().GetParentUUID(), "domain")

	e = cache.Get(vnTestUUID)
	assert.Equal(t, event2, e)
	assert.NotEqual(t, event1, event2)

	riUUID1 := "ri_uuid1"
	ri := makeTestRoutingInstance(&riParams{
		baseResourceParams: baseResourceParams{
			uuid: riUUID1, parentUUID: vn.GetUUID(),
		},
	})

	_, err = cache.processTestEvent(&services.Event{
		Version: 2,
		Request: &services.Event_CreateRoutingInstanceRequest{
			CreateRoutingInstanceRequest: &services.CreateRoutingInstanceRequest{
				RoutingInstance: ri,
			},
		},
	})
	assert.NoError(t, err)

	e = cache.Get(vnTestUUID)

	vn = e.GetUpdateVirtualNetworkRequest().GetVirtualNetwork()
	assert.Len(t, vn.RoutingInstances, 1)
	assert.Equal(t, riUUID1, vn.RoutingInstances[0].UUID)

	riUUID2 := "ri_uuid2"
	ri = makeTestRoutingInstance(&riParams{
		baseResourceParams: baseResourceParams{
			uuid: riUUID2, parentUUID: vn.GetUUID()}, riRefs: []*models.RoutingInstanceRoutingInstanceRef{
			{UUID: riUUID1},
		},
	})

	_, err = cache.processTestEvent(&services.Event{
		Version: 3,
		Request: &services.Event_CreateRoutingInstanceRequest{
			CreateRoutingInstanceRequest: &services.CreateRoutingInstanceRequest{
				RoutingInstance: ri,
			},
		},
	})
	assert.NoError(t, err)
	e = cache.Get(vnTestUUID)

	vn = e.GetUpdateVirtualNetworkRequest().GetVirtualNetwork()
	assert.Len(t, vn.RoutingInstances, 2)
	assert.Equal(t, riUUID2, vn.RoutingInstances[1].UUID)

	e = cache.Get(riUUID1)
	ri = e.GetCreateRoutingInstanceRequest().GetRoutingInstance()
	assert.Len(t, ri.RoutingInstanceBackRefs, 1)
	assert.Equal(t, riUUID2, ri.RoutingInstanceBackRefs[0].UUID)

	event4, err := cache.processTestEvent(&services.Event{
		Version: 4,
		Request: &services.Event_DeleteRoutingInstanceRequest{
			DeleteRoutingInstanceRequest: &services.DeleteRoutingInstanceRequest{
				ID: riUUID2,
			},
		},
	})
	assert.NoError(t, err)

	e = cache.Get(riUUID2)
	assert.Equal(t, event4, e)

	e = cache.Get(vnTestUUID)
	vn = e.GetUpdateVirtualNetworkRequest().GetVirtualNetwork()
	assert.Len(t, vn.RoutingInstances, 1)
	assert.Equal(t, riUUID1, vn.RoutingInstances[0].UUID)

	e = cache.Get(riUUID1)
	ri = e.GetCreateRoutingInstanceRequest().GetRoutingInstance()
	assert.Len(t, ri.RoutingInstanceBackRefs, 0)

	event5, err := cache.processTestEvent(&services.Event{
		Version: 5,
		Request: &services.Event_DeleteVirtualNetworkRequest{
			DeleteVirtualNetworkRequest: &services.DeleteVirtualNetworkRequest{
				ID: vnTestUUID,
			},
		},
	})
	assert.NoError(t, err)

	e = cache.Get(vnTestUUID)
	r := e.GetResource()
	assert.Equal(t, event5, e)
	assert.Equal(t, services.OperationDelete, e.Operation())
	assert.NotEqual(t, vn.ParentUUID, r.GetParentUUID())
}

type baseResourceParams struct {
	uuid       string
	parentUUID string
}

func makeTestVirtualNetwork(vnParams *baseResourceParams) *models.VirtualNetwork {
	vn := models.MakeVirtualNetwork()
	vn.UUID = vnParams.uuid
	vn.ParentUUID = vnParams.parentUUID
	return vn
}

type riParams struct {
	baseResourceParams
	riRefs []*models.RoutingInstanceRoutingInstanceRef
}

func makeTestRoutingInstance(rip *riParams) *models.RoutingInstance {
	ri := models.MakeRoutingInstance()
	ri.UUID = rip.uuid
	ri.ParentUUID = rip.parentUUID
	ri.RoutingInstanceRefs = rip.riRefs
	return ri
}

func (cache *DBCache) processTestEvent(event *services.Event) (*services.Event, error) {
	return cache.Process(context.Background(), event)
}
