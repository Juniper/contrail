package services

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

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

	translator, err := newTranslator(events)
	if err != nil {
		return nil, err
	}

	return fillEventNodesReferences(eventNodes, translator, eventToEventNode)
}

type eventNode struct {
	event *Event
	referencesAndParent  []*eventNode
}

func fillEventNodesReferences(
	eventNodes []*eventNode, translator *eventTranslator, eventToEventNode map[*Event]*eventNode,
) ([]*eventNode, error) {
	for _, node := range eventNodes {
		refs, err := node.event.getReferences()
		if err != nil {
			return nil, err
		}

		for _, ref := range refs {
			refNode, err := getEventNode(ref, translator, eventToEventNode)
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

type eventTranslator struct {
	fromUUIDToEvent   map[string]*Event
	fromFQNameToEvent map[string]*Event
}

func newTranslator(events []*Event) (*eventTranslator, error) {
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

	translator := &eventTranslator{
		fromFQNameToEvent: fqNameToEvent,
		fromUUIDToEvent:   uuidToEvent,
	}

	return translator, nil
}

func (t *eventTranslator) uuidToEvent(uuid string) *Event {
	return t.fromUUIDToEvent[uuid]
}

func (t *eventTranslator) fqNameToEvent(fqname []string) *Event {
	return t.fromFQNameToEvent[basemodels.FQNameToString(fqname)]
}

func (e *Event) getReferences() (basemodels.References, error) {
	res := e.GetResource()
	if res == nil {
		// TODO: Handle other events than Create Resource.
		logrus.Warn("Method eventsToEventNodes() is implemented for CREATE events only.")
		return nil, errors.Errorf("Invalid event. Cannot extract resource from event: %s", e.String())
	}

	refs := res.GetReferences()

	parentUUID := res.GetParentUUID()
	parentFQName := basemodels.ParentFQName(res.GetFQName())

	if parentUUID != "" || len(parentFQName) != 0 {
		parentType := res.GetParentType()
		parentRef := basemodels.NewReference(parentUUID, parentFQName, parentType)
		refs = append(refs, parentRef)
	}

	return refs, nil
}

func getEventNode(ref basemodels.Reference, translator *eventTranslator, eventToEventNode map[*Event]*eventNode) (*eventNode, error) {
	uuid := ref.GetUUID()
	fqname := ref.GetTo()

	var refEv *Event

	if uuid != "" {
		refEv = translator.uuidToEvent(uuid)
	}

	if len(fqname) != 0 && refEv == nil {
		refEv = translator.fqNameToEvent(fqname)
	}

	if refEv == nil {
		// Referenced event is not from the chain of events. It can be ignored.
		return nil, nil
	}

	refNode := eventToEventNode[refEv]
	if refNode == nil {
		return nil, errors.Errorf(
			"Cannot resolve Event Node with UUID: %s and FQ Name: %v", uuid, fqname)
	}
	return refNode, nil
}