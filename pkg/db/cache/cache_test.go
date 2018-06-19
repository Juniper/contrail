package cache

import (
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const numEvent = 4
const timeOut = 10 * time.Second

func addWatcher(t *testing.T, wg *sync.WaitGroup, cache *DB) {
	ctx, _ := context.WithTimeout(context.Background(), timeOut)
	watcher, _ := cache.AddWatcher(ctx, 0)

	go func() {
		wg.Add(1)
		defer wg.Done()
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
	log.SetLevel(logrus.DebugLevel)
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
