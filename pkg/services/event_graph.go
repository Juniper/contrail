package services

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

type eventNode struct {
	event               *Event
	referencesAndParent []*eventNode
}

//EventGraph is a directed graph of events.
type EventGraph struct {
	nodes        []*eventNode
	nodeByUUID   map[string]*eventNode
	nodeByFQName map[string]*eventNode
}

//NewEventGraph creates EventGraph from list of Events.
func NewEventGraph(events []*Event, referencesMap map[*Event]basemodels.References) *EventGraph {
	g := &EventGraph{}
	g.initNodes(events)
	g.fillGraph(referencesMap)
	return g
}

func (g *EventGraph) initNodes(events []*Event) {
	g.nodes = make([]*eventNode, 0, len(events))
	g.nodeByUUID = make(map[string]*eventNode)
	g.nodeByFQName = make(map[string]*eventNode)

	for _, e := range events {
		node := &eventNode{event: e}
		g.nodes = append(g.nodes, node)
		if e.GetUUID() != "" {
			g.nodeByUUID[e.GetUUID()] = node
		}
		if e.GetResource() == nil {
			logrus.Warnf("event: %v got nil resource", e)
			continue
		}
		fqName := e.GetResource().GetFQName()
		if len(fqName) != 0 {
			g.nodeByFQName[basemodels.FQNameToString(fqName)] = node
		}
	}
	return
}

func (g *EventGraph) fillGraph(refMap map[*Event]basemodels.References) {
	for _, node := range g.nodes {
		refs := refMap[node.event]
		for _, ref := range refs {
			if n := g.getNodeByReference(ref); n != nil {
				node.referencesAndParent = append(node.referencesAndParent, n)
			}
		}
	}
}

func (g *EventGraph) getNodeByReference(ref basemodels.Reference) *eventNode {
	if g == nil || ref == nil {
		return nil
	}
	node := g.getNodeByUUID(ref.GetUUID())
	if node == nil {
		node = g.getNodeByFQName(ref.GetTo())
	}
	return node
}

func (g *EventGraph) getNodeByUUID(uuid string) *eventNode {
	if g == nil || uuid == "" {
		return nil
	}
	return g.nodeByUUID[uuid]
}

func (g *EventGraph) getNodeByFQName(fqName []string) *eventNode {
	if g == nil || len(fqName) == 0 {
		return nil
	}
	return g.nodeByFQName[basemodels.FQNameToString(fqName)]
}

//CheckOperationType checks if all operations have the same type
func (e *EventList) CheckOperationType() string {
	if len(e.Events) == 0 {
		logrus.Warn("Unhandled situation for empty event list.")
		return "EMPTY"
	}

	operation := e.Events[0].Operation()
	for _, ev := range e.Events {
		if operation != ev.Operation() {
			return "MIXED"
		}
	}
	return operation
}

func (e *EventList) isSortRequired(g *EventGraph, refMap map[*Event]basemodels.References) bool {
	operation := e.CheckOperationType()
	parsedEvents := make(map[*Event]bool)
	for _, event := range e.Events {
		for _, ref := range refMap[event] {
			if n := g.getNodeByReference(ref); n != nil {
				switch {
				case !parsedEvents[n.event] && operation == OperationCreate:
					return true
				case parsedEvents[n.event] && operation == OperationDelete:
					return true
				}
			}
		}
		parsedEvents[event] = true
	}
	return false
}

//SortEvents sorts events.
func (g *EventGraph) SortEvents() *EventList {
	visited := make(map[*eventNode]bool)
	sorted := make([]*eventNode, 0, len(g.nodes))

	for _, e := range g.nodes {
		if !visited[e] {
			sorted = g.sortUtil(e, sorted, visited)
		}
	}

	list := &EventList{}
	for _, s := range sorted {
		list.Events = append(list.Events, s.event)
	}

	if list.CheckOperationType() == OperationDelete {
		list.Events = reverseList(list.Events)
	}

	return list
}

func reverseList(events []*Event) []*Event {
	for i := len(events)/2 - 1; i >= 0; i-- {
		opp := len(events) - 1 - i
		events[i], events[opp] = events[opp], events[i]
	}
	return events
}

func (g *EventGraph) sortUtil(node *eventNode, sorted []*eventNode, visited map[*eventNode]bool) []*eventNode {
	if visited[node] {
		return sorted
	}
	visited[node] = true

	if len(node.referencesAndParent) == 0 {
		return append(sorted, node)
	}

	for _, r := range node.referencesAndParent {
		sorted = g.sortUtil(r, sorted, visited)
	}
	return append(sorted, node)
}

//CheckCycle checks if there is cycle in graph.
func (g *EventGraph) CheckCycle() bool {
	visited := make(map[*eventNode]bool)
	parsingStack := make(map[*eventNode]bool)
	for _, n := range g.nodes {
		if !visited[n] && g.cycleUtil(n, visited, parsingStack) {
			return true
		}
	}
	return false
}

func (g *EventGraph) cycleUtil(node *eventNode, visited, parsingStack map[*eventNode]bool) bool {
	visited[node] = true
	parsingStack[node] = true
	for _, neighbour := range node.referencesAndParent {
		if parsingStack[neighbour] {
			return true
		}

		if !visited[neighbour] && g.cycleUtil(neighbour, visited, parsingStack) {
			return true
		}
	}
	parsingStack[node] = false
	return false
}

func (e *Event) getReferences() (basemodels.References, error) {
	res := e.GetResource()
	if res == nil {
		// TODO: Handle other events than Create Resource.
		logrus.Warn("Method eventsToEventNodes() is implemented for CREATE events only.")
		return nil, errors.Errorf("Invalid event. Cannot extract resource from event: %s", e.String())
	}

	refs := res.GetReferences()

	if parentRef := extractParentAsRef(res); parentRef != nil {
		refs = append(refs, parentRef)
	}
	return refs, nil
}

func extractParentAsRef(o basemodels.Object) basemodels.Reference {
	parentUUID := o.GetParentUUID()
	parentFQName := basemodels.ParentFQName(o.GetFQName())
	if parentUUID == "" && len(parentFQName) == 0 {
		return nil
	}

	parentType := o.GetParentType()
	return basemodels.NewReference(parentUUID, parentFQName, parentType)
}
