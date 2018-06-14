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
const timeOut = time.Second

func addWatcher(t *testing.T, wg *sync.WaitGroup, cache *DB) {
	ctx, _ := context.WithTimeout(context.Background(), timeOut)
	watcher := cache.AddWatcher(ctx, 0)

	go func() {
		wg.Add(1)
		defer wg.Done()
		for i := 0; i < numEvent; i++ {
			select {
			case <-ctx.Done():
				log.Debugf("[watcher %d] time out on test", watcher.id)
				assert.Fail(t, "watcher timed out")
			case e := <-watcher.ch:
				log.Debugf("[watcher %d] got event version %d", watcher.id, e.Version)
				assert.Equal(t, uint64(i), e.Version)
			}
		}
	}()
}

func notifyEvent(t *testing.T, cache *DB, version uint64) {
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
	cache.Process(context.Background(), event)
}

func TestCache(t *testing.T) {
	log.SetLevel(logrus.DebugLevel)
	cache := New()
	wg := &sync.WaitGroup{}

	addWatcher(t, wg, cache)
	addWatcher(t, wg, cache)

	notifyEvent(t, cache, 0)
	notifyEvent(t, cache, 1)

	addWatcher(t, wg, cache)
	addWatcher(t, wg, cache)
	// test cancelation of channel.
	// expect no panic or blocking.
	ctx2, cancel := context.WithCancel(context.Background())
	cache.AddWatcher(ctx2, 0)
	cancel()

	//timeout watcher
	//Don't actually receiving events.
	ctx3 := context.Background()
	cache.AddWatcher(ctx3, 0)

	notifyEvent(t, cache, 2)
	notifyEvent(t, cache, 3)

	addWatcher(t, wg, cache)

	wg.Wait()
}
