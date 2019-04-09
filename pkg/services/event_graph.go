package services

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

//EventGraph is a directed graphs of events.
type EventGraph struct {
	nodes []*eventNode
}

type eventNode struct {
	event *Event
	refs  []*eventNode
}

//NewEventGraph creates new event graph from event list.
func NewEventGraph(events []*Event) (*EventGraph, error) {
	nodes, err := eventsToEventNodes(events)
	return &EventGraph{nodes: nodes}, err
}

func eventsToEventNodes(events []*Event) ([]*eventNode, error) {
	translator, err := getTranslatorToEvents(events)
	if err != nil {
		return nil, err
	}

	eventToEventNode := make(map[*Event]*eventNode)
	eventNodes := []*eventNode{}

	// Prepare event Nodes and create map that finds eventNode when known event
	for id, e := range events {
		node := &eventNode{
			event: events[id],
		}
		eventNodes = append(eventNodes, node)
		eventToEventNode[e] = node
	}

	return fillEventNodesReferences(eventNodes, translator, eventToEventNode)
}

//SortEvents sorts events by reference chain.
func (g *EventGraph) SortEvents() EventList {
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

func (g *EventGraph) sortUtil(node *eventNode, sorted []*eventNode, visited map[*eventNode]bool) []*eventNode {
	if visited[node] {
		return sorted
	}
	visited[node] = true

	if len(node.refs) == 0 {
		return append(sorted, node)
	}

	for _, r := range node.refs {
		sorted = g.sortUtil(r, sorted, visited)
	}
	return append(sorted, node)
}

//CheckCycle checks if there is cycle in events references.
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
	for _, neighbour := range node.refs {
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

type eventTranslator struct {
	fromUUIDToEvent   map[string]*Event
	fromFQNameToEvent map[string]*Event
}

func (t *eventTranslator) uuidToEvent(uuid string) *Event {
	return t.fromUUIDToEvent[uuid]
}

func (t *eventTranslator) fqNameToEvent(fqname []string) *Event {
	return t.fromFQNameToEvent[basemodels.FQNameToString(fqname)]
}

type ref struct {
	toUUID   string
	toFQNAME []string
}

func (r *ref) isEmpty() bool {
	return r.toUUID == "" && len(r.toFQNAME) == 0
}

func getTranslatorToEvents(events []*Event) (*eventTranslator, error) {
	uuidToEvent := make(map[string]*Event)
	fqNameToEvent := make(map[string]*Event)

	for id, e := range events {
		res := e.GetResource()
		if res == nil {
			// TODO: Handle other events than Create Resource.
			logrus.Warn("Method eventsToEventNodes() is implemented for CREATE events only.")
			return nil, errors.Errorf("Cannot extract resource from event: %s", e.String())
		}

		if res.GetUUID() != "" {
			uuidToEvent[res.GetUUID()] = events[id]
		}
		if len(res.GetFQName()) != 0 {
			fqNameToEvent[basemodels.FQNameToString(res.GetFQName())] = events[id]
		}

		if res.GetUUID() == "" && len(res.GetFQName()) == 0 {
			// TODO: Handle events with no UUID or FQ Name.
			logrus.Warn("Method eventsToEventNodes() is implemented for CREATE events only")
			return nil, errors.Errorf(
				"Cannot create translator because resource has no uuid and no FQ Name: %s", e.String())
		}
	}

	translator := &eventTranslator{
		fromFQNameToEvent: fqNameToEvent,
		fromUUIDToEvent:   uuidToEvent,
	}

	return translator, nil
}

func fillEventNodesReferences(
	eventNodes []*eventNode, translator *eventTranslator, eventToEventNode map[*Event]*eventNode,
) ([]*eventNode, error) {
	for id, node := range eventNodes {
		event := node.event
		refs, err := event.getRefs()
		if err != nil {
			return nil, err
		}

		for _, ref := range refs {
			refNode, err := ref.getEventNode(translator, eventToEventNode)
			if err != nil {
				return nil, err
			}
			if refNode != nil {
				eventNodes[id].refs = append(eventNodes[id].refs, refNode)
			}
		}
	}
	return eventNodes, nil
}

func (r *ref) getEventNode(translator *eventTranslator, eventToEventNode map[*Event]*eventNode) (*eventNode, error) {
	if r.isEmpty() {
		return nil, errors.New("event Node has invalid reference (no UUID and no FQ Name)")
	}

	var refEv *Event

	if r.toUUID != "" {
		refEv = translator.uuidToEvent(r.toUUID)
	}

	if len(r.toFQNAME) != 0 && refEv == nil {
		refEv = translator.fqNameToEvent(r.toFQNAME)
	}

	if refEv == nil {
		// Referenced event is not from the chain of events. It can be ignored.
		return nil, nil
	}

	refNode := eventToEventNode[refEv]
	if refNode == nil {
		return nil, errors.Errorf(
			"Cannot resolve Event Node with UUID: %s and FQ Name: %v", r.toUUID, r.toFQNAME)
	}
	return refNode, nil
}

func (e *Event) getRefs() ([]*ref, error) {
	res := e.GetResource()
	if res == nil {
		// TODO: Handle other events than Create Resource.
		logrus.Warn("Method eventsToEventNodes() is implemented for CREATE events only.")
		return nil, errors.Errorf("Invalid event. Cannot extract resource from event: %s", e.String())
	}

	result := []*ref{}

	parentRef := &ref{
		toUUID:   res.GetParentUUID(),
		toFQNAME: basemodels.ParentFQName(res.GetFQName()),
	}
	if !parentRef.isEmpty() {
		result = append(result, parentRef)
	}

	refs, err := parseRefs(res.GetReferences())
	if err != nil {
		return nil, err
	}

	result = append(result, refs...)
	return result, nil
}

func parseRefs(refs basemodels.References) ([]*ref, error) {
	result := []*ref{}
	for _, r := range refs {
		parsedRef := &ref{
			toUUID:   r.GetUUID(),
			toFQNAME: r.GetTo(),
		}
		if !parsedRef.isEmpty() {
			result = append(result, parsedRef)
		} else {
			return nil, errors.Errorf("Cannot get reference UUID or FQName: %v", r)
		}
	}
	return result, nil
}
