package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/services"

	log "github.com/sirupsen/logrus"
)

//DB is cache db implemenation.
type DB struct {
	first        *node
	last         *node
	lastIndex    uint64
	maxHistory   uint64
	deleted      []*node
	versionMap   map[uint64]*node
	idMap        map[string]*node
	resources    map[string]map[string]*node
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
}

type node struct {
	prev    *node
	version uint64
	event   *services.Event
	next    *node
}

func (w *Watcher) notify() {
	w.updateCh <- true
}

//Chan returns event streams.
func (w *Watcher) Chan() chan *services.Event {
	return w.ch
}

//nolint: gocyclo
func (w *Watcher) watch(ctx context.Context, db *DB) {
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

	for w.node == nil {
		log.Debugf("[Watcher %d] waiting for first event", w.id)
		select {
		case <-w.updateCh:
			log.Debugf("[Watcher %d] got first event", w.id)
			w.node = db.getFirst()
		case <-ctx.Done():
			log.Debugf("[Watcher %d] canceled by context", w.id)
			return
		}
	}

	for {
		// send out for event.
		log.Debugf("[Watcher %d] send event %v", w.id, w.node)

		select {
		case w.ch <- w.node.event:
		case <-ctx.Done():
			log.Debugf("[Watcher %d] canceled by context", w.id)
			return
		}

		next := w.node.getNext()

		for next == nil {
			log.Debugf("[Watcher %d] waiting for next event", w.id)
			select {
			case <-w.updateCh:
				next = w.node.getNext()
				log.Debugf("[Watcher %d] got next event %v", w.id, next)
			case <-ctx.Done():
				log.Debugf("[Watcher %d] canceled by context", w.id)
				return
			}
		}
		w.node = next
	}
}

//NewDB makes cache db.
func NewDB(maxHistory uint64) *DB {
	//TODO(nati) db dump
	db := &DB{
		versionMap:   map[uint64]*node{},
		idMap:        map[string]*node{},
		watchers:     map[uint64]*Watcher{},
		deleted:      []*node{},
		mutex:        sync.RWMutex{},
		watcherMutex: sync.RWMutex{},
		maxHistory:   maxHistory,
		resources:    map[string]map[string]*node{},
	}
	return db
}

//AddWatcher registers new watcher.
func (db *DB) AddWatcher(ctx context.Context, versionID uint64) (*Watcher, error) {
	watcherID := uint64(len(db.watchers))
	node := db.getNodeByVersion(versionID)
	if versionID != 0 && node == nil {
		return nil, common.ErrorBadRequest("requested version is compacted.")
	}
	watcher := &Watcher{
		id:       watcherID,
		ch:       make(chan *services.Event),
		updateCh: make(chan bool, 1000),
		node:     db.getNodeByVersion(versionID),
	}
	go watcher.watch(ctx, db)

	db.watcherMutex.Lock()
	db.watchers[watcherID] = watcher
	db.watcherMutex.Unlock()

	return watcher, nil
}

func (db *DB) update(event *services.Event) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	n := &node{
		prev:  db.last,
		event: event,
	}
	resource := event.GetResource()

	existingNode, ok := db.idMap[resource.GetUUID()]
	if ok {
		// We should consider throwing error if operation is create here
		oldResource := existingNode.event.GetResource()
		db.removeDependencies(oldResource)
		log.Debugf("Update id map for key: %s,  event version: %d", resource.GetUUID(), event.Version)
		if existingNode == db.first {
			db.first = existingNode.getNext()
		}
		delete(db.idMap, existingNode.event.GetResource().GetUUID())
		delete(db.versionMap, existingNode.version)
		existingNode.pop()
	}

	if event.Operation() == services.OperationDelete {
		db.deleted = append(db.deleted, n)
	}

	db.append(n)
	// first update version so that in case of maxHistory exceeded we remove that version
	db.updateDBVersion(n)
	db.removeTooOldNodesIfNeeded()
	// This is done after removing too old objects as the uuid is same regardless of operation type
	db.idMap[resource.GetUUID()] = n

	db.updateDependentNodes(event)
	db.handleNode(n)

	log.Debugf("node %v updated", n.version)
}

func (db *DB) updateDBVersion(n *node) {
	db.versionMap[db.lastIndex] = n
	n.version = db.lastIndex
	db.lastIndex++
}

func (db *DB) append(n *node) {
	if db.first == nil {
		db.first = n
	}
	db.last.setNext(n)
	n.setPrev(db.last)
	db.last = n
}

func (db *DB) removeTooOldNodesIfNeeded() {
	if len(db.deleted) > 0 && db.lastIndex-db.deleted[0].version > db.maxHistory {
		compactedNode := db.deleted[0]
		db.deleted = db.deleted[1:]
		delete(db.idMap, compactedNode.event.GetResource().GetUUID())
		delete(db.versionMap, compactedNode.version)
		compactedNode.pop()
	}
}

func (db *DB) addDependencies(resource services.Resource) {
	dependencies := resource.Depends()
	for _, dependencyID := range dependencies {
		dependentNode, ok := db.idMap[dependencyID]
		if ok {
			dependentNode.event.GetResource().AddDependency(resource)
			dependentNode.pop()
			db.append(dependentNode)
		}
	}
}

func (db *DB) removeDependencies(resource services.Resource) {
	dependencies := resource.Depends()
	for _, dependencyID := range dependencies {
		dependentNode, ok := db.idMap[dependencyID]
		if ok {
			dependentNode.event.GetResource().RemoveDependency(resource)
			dependentNode.pop()
			db.append(dependentNode)
		}
	}
}

func (db *DB) updateDependentNodes(event services.HasResource) {
	if event.Operation() == services.OperationCreate || event.Operation() == services.OperationUpdate {
		db.addDependencies(event.GetResource())
	}
}

//Process updates cache data.
func (db *DB) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	db.update(event)
	db.watcherMutex.RLock()
	defer db.watcherMutex.RUnlock()
	for _, watcher := range db.watchers {
		watcher.notify()
	}
	return event, nil
}

func (db *DB) getFirst() *node {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	return db.first
}

//Get returns a resource by id.
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

	if version == 0 {
		return db.first
	}

	node, ok := db.versionMap[version]
	if !ok {
		return nil
	}
	return node
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
	if n != nil {
		n.prev = p
	}
}

func (n *node) setNext(next *node) {
	if n != nil {
		n.next = next
	}
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
	n.next = nil
	n.prev = nil
}

func (db *DB) handleNode(n *node) {
	e := n.event
	r := e.GetResource()
	k := r.Kind()
	uuid := r.GetUUID()

	switch e.Operation() {
	case services.OperationCreate, services.OperationUpdate:
		if db.resources[k] == nil {
			db.resources[k] = map[string]*node{}
		}
		db.resources[k][uuid] = n
	case services.OperationDelete:
		delete(db.resources[k], uuid)
	}
}
