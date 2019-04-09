package services

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

type eventNode struct {
	event *Event
	referencesAndParent  []*eventNode
}

type eventGraph struct {
	nodes []*eventNode
}

func newEventGraph(events []*Event) (*eventGraph, error) {
	nodes, err := eventsToEventNodes(events)
	return &eventGraph{nodes: nodes}, err
}

func eventsToEventNodes(events []*Event) ([]*eventNode, error) {
	eventToEventNode := map[*Event]*eventNode{}
	eventNodes := make([]*eventNode, 0, len(events))

	// Prepare event Nodes and create map that finds eventNode when known event
	for _, e := range events {
		node := &eventNode{
			event: e,
		}
		eventNodes = append(eventNodes, node)
		eventToEventNode[e] = node
	}

	refToEvent, err := newReferenceToEventParser(events)
	if err != nil {
		return nil, err
	}

	return fillEventNodesReferences(eventNodes, refToEvent, eventToEventNode)
}

type referenceToEvent struct {
	fromUUIDToEvent   map[string]*Event
	fromFQNameToEvent map[string]*Event
}

func (g *eventGraph) sortEvents() EventList {
	visited := make(map[*eventNode]bool)
	sorted := make([]*eventNode, 0, len(g.nodes))

	for _, e := range g.nodes {
		if !visited[e] {
			sorted = g.sortUtil(e, sorted, visited)
		}
	}

	list := EventList{}
	for _, s := range sorted {
		list.Events = append(list.Events, s.event)
	}
	return list
}

func (g *eventGraph) sortUtil(node *eventNode, sorted []*eventNode, visited map[*eventNode]bool) []*eventNode {
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

func (g *eventGraph) checkCycle() bool {
	visited := make(map[*eventNode]bool)
	parsingStack := make(map[*eventNode]bool)
	for _, n := range g.nodes {
		if !visited[n] && g.cycleUtil(n, visited, parsingStack) {
			return true
		}
	}
	return false
}

func (g *eventGraph) cycleUtil(node *eventNode, visited, parsingStack map[*eventNode]bool) bool {
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

func newReferenceToEventParser(events []*Event) (*referenceToEvent, error) {
	uuidToEvent := map[string]*Event{}
	fqNameToEvent := map[string]*Event{}

	for _, event := range events {
		res := event.GetResource()
		if res == nil {
			// TODO: Handle other events than Create Resource.
			logrus.Warn("Method eventsToEventNodes() is implemented for CREATE events only.")
			return nil, errors.Errorf("Cannot extract resource from event: %s", event.String())
		}

		if res.GetUUID() != "" {
			uuidToEvent[res.GetUUID()] = event
		}
		if len(res.GetFQName()) != 0 {
			fqNameToEvent[basemodels.FQNameToString(res.GetFQName())] = event
		}

		if res.GetUUID() == "" && len(res.GetFQName()) == 0 {
			// TODO: Handle events with no UUID or FQ Name.
			logrus.Warn("Method eventsToEventNodes() is implemented for CREATE events only")
			return nil, errors.Errorf(
				"Cannot create translator because resource has no uuid and no FQ Name: %s", event.String())
		}
	}

	return &referenceToEvent{
		fromFQNameToEvent: fqNameToEvent,
		fromUUIDToEvent:   uuidToEvent,
	}, nil
}

func fillEventNodesReferences(
	eventNodes []*eventNode, refToEvent *referenceToEvent, eventToEventNode map[*Event]*eventNode,
) ([]*eventNode, error) {
	for _, node := range eventNodes {
		refs, err := node.event.getReferences()
		if err != nil {
			return nil, err
		}

		for _, ref := range refs {
			refNode, err := getEventNode(ref, refToEvent, eventToEventNode)
			if err != nil {
				return nil, err
			}
			if refNode != nil {
				node.referencesAndParent = append(node.referencesAndParent, refNode)
			}
		}
	}
	return eventNodes, nil
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

func getEventNode(ref basemodels.Reference, refToEvent *referenceToEvent, eventToEventNode map[*Event]*eventNode) (*eventNode, error) {
	refEv := refToEvent.toEvent(ref)
	if refEv == nil {
		// Referenced event is not from the chain of events. It can be ignored.
		return nil, nil
	}

	refNode := eventToEventNode[refEv]
	if refNode == nil {
		return nil, errors.Errorf(
			"Cannot resolve Event Node with UUID: %s and FQ Name: %v", ref.GetUUID(), ref.GetTo())
	}
	return refNode, nil
}

func (t *referenceToEvent) toEvent(ref basemodels.Reference) *Event{
	uuid := ref.GetUUID()
	fqname := ref.GetTo()

	var refEv *Event

	if uuid != "" {
		refEv = t.fromUUIDToEvent[uuid]
	}

	if len(fqname) != 0 && refEv == nil {
		refEv = t.fromFQNameToEvent[basemodels.FQNameToString(fqname)]
	}

	return refEv
}
