package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Juniper/contrail/pkg/services"

	log "github.com/sirupsen/logrus"
)

//DB is cache db implemenation.
type DB struct {
	first        *node
	last         *node
	lastIndex    uint64
	versionMap   map[uint64]*node
	idMap        map[string]*node
	watchers     map[uint64]*Watcher
	mutex        sync.RWMutex
	watcherMutex sync.RWMutex
}

//Watcher watches cache update.
type Watcher struct {
	ch       chan *services.Event
	updateCh chan bool
	id       uint64
	node     *node
	context  context.Context
}

func (w *Watcher) notify() {
	select {
	case w.updateCh <- true:
		log.Debugf("[Watcher %d] sent notify", w.id)
	default:
		log.Debugf("[Watcher %d] no notify", w.id)
	}
}

//Chan returns event streams.
func (w *Watcher) Chan() chan *services.Event {
	return w.ch
}

func (w *Watcher) watch(db *DB, timeout time.Duration) {
	defer func() {
		log.Debugf("[Watcher %d] watch stopped", w.id)

		db.watcherMutex.Lock()
		defer db.watcherMutex.Unlock()
		_, ok := db.watchers[w.id]
		if ok {
			delete(db.watchers, w.id)
		}
		close(w.ch)
		close(w.updateCh)
	}()
	log.Debugf("[Watcher %d] watch started", w.id)
	if w.node == nil {
		log.Debugf("[Watcher %d] waiting for first event", w.id)
		select {
		case <-w.updateCh:
			log.Debugf("[Watcher %d] got first event", w.id)
			w.node = db.GetFirst()
		case <-w.context.Done():
			log.Debugf("[Watcher %d] canceled by context", w.id)
			return
		}
	}
	for {
		// send out for event.
		var next *node
		if w.node != nil {
			log.Debugf("[Watcher %d] send event %v", w.id, w.node)
			select {
			case w.ch <- w.node.event:
			case <-w.context.Done():
				log.Debugf("[Watcher %d] canceled by context", w.id)
				return
			case <-time.After(timeout):
				log.Warnf("[Watcher %d] timeout on dump", w.id)
				return
			}
			next = w.node.getNext()
		}

		if next == nil {
			log.Debugf("[Watcher %d] waiting for next event", w.id)
			select {
			case <-w.updateCh:
				next = w.node.getNext()
				log.Debugf("[Watcher %d] got next event %v", w.id, next)
			case <-w.context.Done():
				log.Debugf("[Watcher %d] canceled by context", w.id)
				return
			}
		}
		w.node = next
	}
}

//Process call event handler.
func (w *Watcher) Process(ctx context.Context, service services.Service) {
	for {
		select {
		case e := <-w.ch:
			log.Debugf("[Watcher %d] process event %v", w.id, e)
			_, err := e.Process(ctx, service)
			if err != nil {
				log.Warn(err)
			}
		case <-w.context.Done():
			log.Debugf("[Watcher %d] canceled by context", w.id)
			return
		}
	}
}

//New makes cache db.
func New() *DB {
	//TODO(nati) db dump
	return &DB{
		versionMap:   map[uint64]*node{},
		idMap:        map[string]*node{},
		watchers:     map[uint64]*Watcher{},
		mutex:        sync.RWMutex{},
		watcherMutex: sync.RWMutex{},
	}
}

//AddWatcher registers new watcher.
func (db *DB) AddWatcher(ctx context.Context, versionID uint64, timeout time.Duration) *Watcher {
	watcherID := uint64(len(db.watchers))
	watcher := &Watcher{
		id:       watcherID,
		ch:       make(chan *services.Event),
		updateCh: make(chan bool),
		node:     db.getNodeByVersion(versionID),
		context:  ctx,
	}
	go watcher.watch(db, timeout)

	db.watcherMutex.Lock()
	db.watchers[watcherID] = watcher
	db.watcherMutex.Unlock()

	return watcher
}

func (db *DB) update(event *services.Event, timeout time.Duration) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	n := &node{
		prev:  db.last,
		event: event,
	}
	resource := event.GetResource()

	db.versionMap[db.lastIndex] = n
	n.version = db.lastIndex
	db.lastIndex++

	existingNode, ok := db.idMap[resource.GetUUID()]
	if ok {
		log.Debugf("compact existing map for event %d", event.Version)
		if existingNode == db.first {
			db.first = existingNode.getNext()
		} else {
			existingNode.pop()
		}
	}
	db.idMap[resource.GetUUID()] = n

	if db.first == nil {
		db.first = n
	}

	db.last.setNext(n)
	db.last = n
	return nil
}

//Process updates cache data.
func (db *DB) Process(event *services.Event, timeout time.Duration) error {
	db.update(event, timeout)
	db.watcherMutex.RLock()
	defer db.watcherMutex.RUnlock()
	for _, watcher := range db.watchers {
		watcher.notify()
	}
	return nil
}

//GetFirst returns the first node.
func (db *DB) GetFirst() *node {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	return db.first
}

//Get returns a resource for id.
func (db *DB) Get(id string) *services.Event {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	node, ok := db.idMap[id]
	if !ok {
		return nil
	}
	return node.event
}

func (db *DB) getNodeByVersion(version uint64) *node {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	node, ok := db.versionMap[version]
	if !ok {
		return db.first
	}
	return node
}

type node struct {
	prev    *node
	version uint64
	event   *services.Event
	next    *node
}

//String for loggging.
func (n *node) String() string {
	return fmt.Sprintf("[node %d]", n.version)
}

func (n *node) getNext() *node {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *node) setPrev(p *node) {
	if n == nil {
		return
	}
	n.prev = p
}

func (n *node) setNext(next *node) {
	if n == nil {
		return
	}
	n.next = next
}

func (n *node) getPrev() *node {
	if n == nil {
		return nil
	}
	return n.prev
}

func (n *node) pop() {
	prev := n.getPrev()
	next := n.getNext()
	prev.setNext(next)
	next.setPrev(prev)
}
