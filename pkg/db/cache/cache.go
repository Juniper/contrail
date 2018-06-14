package cache

import (
	"context"
	"sync"
	"time"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/siddontang/go/log"
)

//DB is cache db implemenation.
type DB struct {
	first      *node
	last       *node
	lastIndex  int64
	versionMap map[int64]*node
	idMap      map[string]*node
	watchers   map[context.Context]chan *services.Event
	mutex      sync.RWMutex
}

func (db *DB) NewDB() *DB {
	//TODO(nati) db dump
	return &DB{
		versionMap: map[int64]*node{},
		mutex:      sync.RWMutex{},
	}
}

func (db *DB) AddWatcher(ctx context.Context) chan *services.Event {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	ch := make(chan *services.Event)
	db.watchers[ctx] = ch
	go func() {
		<-ctx.Done()
		db.CancelWatcher(ctx)
	}()
	return ch
}

func (db *DB) CancelWatcher(ctx context.Context) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	ch, ok := db.watchers[ctx]
	if !ok {
		return
	}
	close(ch)
	delete(db.watchers, ctx)
}

func (db *DB) Process(event *services.Event, timeout time.Duration) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	n := &node{
		prev:  db.last,
		event: event,
	}
	resource := event.GetResource()

	db.versionMap[db.lastIndex] = n

	existingNode, ok := db.idMap[resource.GetUUID()]
	if !ok {
		db.idMap[resource.GetUUID()] = n
		existingNode.pop()
	}

	if db.first == nil {
		db.first = n
	}
	db.last.setNext(n)
	db.last = n

	for ctx, ch := range db.watchers {
		select {
		case ch <- n.event:
		case <-time.After(timeout):
			log.Errorf("watcher for %v timeout", ctx)
		}
	}
}

func (db *DB) GetFirst() *node {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	return db.first
}

func (db *DB) Get(id string) *services.Event {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	node, ok := db.idMap[id]
	if !ok {
		return nil
	}
	return node.event
}

type node struct {
	prev    *node
	version int64
	event   *services.Event
	next    *node
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
