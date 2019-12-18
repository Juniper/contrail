package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
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
		logrus.Debugf("[Watcher %d] watch stopped", w.id)

		db.watcherMutex.Lock()
		defer db.watcherMutex.Unlock()
		_, ok := db.watchers[w.id]
		if ok {
			delete(db.watchers, w.id)
		}
		close(w.ch)
		close(w.updateCh)
	}()
	logrus.Debugf("[Watcher %d] watch started", w.id)

	for w.node == nil {
		logrus.Debugf("[Watcher %d] waiting for first event", w.id)
		select {
		case <-w.updateCh:
			logrus.Debugf("[Watcher %d] got first event", w.id)
			w.node = db.getFirst()
		case <-ctx.Done():
			logrus.Debugf("[Watcher %d] canceled by context", w.id)
			return
		}
	}

	for {
		// send out for event.
		logrus.Debugf("[Watcher %d] send event %v", w.id, w.node)

		select {
		case w.ch <- w.node.event:
		case <-ctx.Done():
			logrus.Debugf("[Watcher %d] canceled by context", w.id)
			return
		}

		next := w.node.getNext()

		for next == nil {
			logrus.Debugf("[Watcher %d] waiting for next event", w.id)
			select {
			case <-w.updateCh:
				next = w.node.getNext()
				logrus.Debugf("[Watcher %d] got next event %v", w.id, next)
			case <-ctx.Done():
				logrus.Debugf("[Watcher %d] canceled by context", w.id)
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
		return nil, errutil.ErrorBadRequest("requested version is compacted.")
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

	uuid := event.GetUUID()
	if uuid == "" {
		logrus.WithField("event", event).Debug("Skipping event")
		return
	}

	existingNode, ok := db.idMap[uuid]
	var backRefs, children []basemodels.Object
	if ok {
		// We should consider throwing error if operation is create here
		oldResource := existingNode.event.GetResource()
		if oldResource != nil {
			backRefs = oldResource.GetBackReferences()
			children = oldResource.GetChildren()
			db.removeDependencies(oldResource)
		}
		logrus.Debugf("Update id map for key: %s,  event version: %d", uuid, event.Version)
		if existingNode == db.first {
			db.first = existingNode.getNext()
		}
		delete(db.idMap, existingNode.event.GetUUID())
		delete(db.versionMap, existingNode.version)
		existingNode.pop()
	}

	n := &node{
		prev:  db.last,
		event: event,
	}

	if event.Operation() == services.OperationDelete {
		db.deleted = append(db.deleted, n)
	}

	db.append(n)
	// first update version so that in case of maxHistory exceeded we remove that version
	db.updateDBVersion(n)
	db.removeTooOldNodesIfNeeded()
	// This is done after removing too old objects as the uuid is same regardless of operation type
	db.idMap[uuid] = n

	db.updateDependentNodes(event, backRefs, children)
	db.handleNode(n)
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
		delete(db.idMap, compactedNode.event.GetUUID())
		delete(db.versionMap, compactedNode.version)
		compactedNode.pop()
	}
}

func (db *DB) addDependencies(
	resource basemodels.Object,
	backRefs, children []basemodels.Object,
) {
	for _, backRef := range backRefs {
		resource.AddBackReference(backRef)
	}
	for _, child := range children {
		resource.AddChild(child)
	}
	for _, ref := range resource.GetReferences() {
		dependentNode, ok := db.idMap[ref.GetUUID()]
		if ok {
			if r := dependentNode.event.GetResource(); r != nil {
				r.AddBackReference(resource)
			}
		}
	}
	if resource.GetParentUUID() != "" {
		dependentNode, ok := db.idMap[resource.GetParentUUID()]
		if ok {
			if r := dependentNode.event.GetResource(); r != nil {
				r.AddChild(resource)
			}
		}
	}
}

func (db *DB) removeDependencies(resource basemodels.Object) {
	for _, backRef := range resource.GetBackReferences() {
		resource.RemoveBackReference(backRef)
	}
	for _, child := range resource.GetChildren() {
		resource.RemoveChild(child)
	}
	for _, ref := range resource.GetReferences() {
		dependentNode, ok := db.idMap[ref.GetUUID()]
		if ok {
			if r := dependentNode.event.GetResource(); r != nil {
				r.RemoveBackReference(resource)
			}
			dependentNode.pop()
			db.append(dependentNode)
		}
	}
	if resource.GetParentUUID() != "" {
		dependentNode, ok := db.idMap[resource.GetParentUUID()]
		if ok {
			if r := dependentNode.event.GetResource(); r != nil {
				r.RemoveChild(resource)
			}
			dependentNode.pop()
			db.append(dependentNode)
		}
	}
}

func (db *DB) updateDependentNodes(
	event services.ResourceEvent,
	backRefs, children []basemodels.Object,
) {
	if event.Operation() == services.OperationCreate || event.Operation() == services.OperationUpdate {
		db.addDependencies(event.GetResource(), backRefs, children)
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

	switch e.Operation() {
	case services.OperationCreate, services.OperationUpdate:
		if db.resources[e.Kind()] == nil {
			db.resources[e.Kind()] = map[string]*node{}
		}
		db.resources[e.Kind()][e.GetUUID()] = n
	case services.OperationDelete:
		delete(db.resources[e.Kind()], e.GetUUID())
	}
}
