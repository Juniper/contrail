package services

import (
	"context"
	"encoding/json"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// Possible operations of events.
const (
	OperationCreate = "CREATE"
	OperationUpdate = "UPDATE"
	OperationDelete = "DELETE"
)

// EventOption contains options for Event.
type EventOption struct {
	UUID      string
	Operation string
	Kind      string
	Data      map[string]interface{}
	FieldMask *types.FieldMask
}

// ResourceEvent is an event that relates to a resource.
type ResourceEvent interface {
	GetResource() basemodels.Object
	Operation() string
}

// ReferenceEvent is an event that relates to a reference.
type ReferenceEvent interface {
	GetID() string
	GetReference() basemodels.Reference
	Operation() string
}

// NewRefUpdateFromEvent creates RefUpdate from ReferenceEvent.
func NewRefUpdateFromEvent(e ReferenceEvent) RefUpdate {
	ref := e.GetReference()
	u := RefUpdate{
		Operation: ParseRefOperation(e.Operation()),
		Type:      ref.GetFromKind(),
		UUID:      e.GetID(),
		RefType:   ref.GetToKind(),
		RefUUID:   ref.GetUUID(),
	}

	if attr := ref.GetAttribute(); attr != nil {
		u.Attr = attr.ToMap()
	}
	return u
}

// CanProcessService is interface for process service.
type CanProcessService interface {
	Process(ctx context.Context, service Service) (*Event, error)
}

// EventList has multiple rest requests.
type EventList struct {
	Events []*Event `json:"resources" yaml:"resources"`
}

type state int

const (
	notVisited state = iota
	visited
	temporaryVisited
)

// Sort sorts Events by parent-child dependency using Tarjan algorithm.
// It doesn't verify reference cycles.
func (e *EventList) Sort() (err error) {
	var sorted []*Event
	stateGraph := map[string]state{}
	eventMap := map[string]*Event{}
	for _, event := range e.Events {
		uuid := event.GetUUID()
		stateGraph[uuid] = notVisited
		eventMap[uuid] = event
	}
	foundNotVisited := true
	for foundNotVisited {
		foundNotVisited = false
		for _, event := range e.Events {
			uuid := event.GetUUID()
			st := stateGraph[uuid]
			if st == notVisited {
				sorted, err = visitResource(uuid, sorted, eventMap, stateGraph)
				if err != nil {
					return err
				}
				foundNotVisited = true
				break
			}
		}
	}
	e.Events = sorted
	return nil
}

//reorder request using Tarjan's algorithm
func visitResource(uuid string, sorted []*Event,
	eventMap map[string]*Event, stateGraph map[string]state,
) (sortedList []*Event, err error) {
	if stateGraph[uuid] == temporaryVisited {
		return nil, errors.New("dependency loop found in sync request")
	}
	if stateGraph[uuid] == visited {
		return sorted, nil
	}
	stateGraph[uuid] = temporaryVisited
	event, found := eventMap[uuid]
	if !found {
		stateGraph[uuid] = visited
		return sorted, nil
	}
	r := event.GetResource()
	var parentUUID string
	if r != nil {
		parentUUID = r.GetParentUUID()
	}

	sorted, err = visitResource(parentUUID, sorted, eventMap, stateGraph)
	if err != nil {
		return nil, err
	}

	stateGraph[uuid] = visited
	sorted = append(sorted, event)
	return sorted, nil
}

type eventNode struct {
	event *Event
	refs  []*eventNode
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

// Process dispatches resource event to call corresponding service functions.
func (e *Event) Process(ctx context.Context, service Service) (*Event, error) {
	if e == nil {
		return nil, errors.Errorf("can not process a nil event")
	}
	p, ok := e.Request.(CanProcessService)
	if !ok {
		return nil, errors.Errorf(
			"can not process event: %v with request type: %T and operation: %s",
			e, e.Request, e.Operation())
	}
	return p.Process(ctx, service)
}

// Process process list of events.
func (e *EventList) Process(ctx context.Context, service Service) (*EventList, error) {
	var responses []*Event
	for _, event := range e.Events {
		response, err := event.Process(ctx, service)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}
	return &EventList{
		Events: responses,
	}, nil
}

// GetResource returns event on resource.
func (e *Event) GetResource() basemodels.Object {
	if e == nil {
		return nil
	}
	switch r := e.Unwrap().(type) {
	case CreateRequest:
		return r.GetResource()
	case UpdateRequest:
		return r.GetResource()
	default:
		return nil
	}
}

// Unwrap returns request wrapped by isEvent_Request interface.
func (e *Event) Unwrap() Request {
	switch er := e.GetRequest().(type) {
	case createEventRequest:
		return er.GetRequest()
	case updateEventRequest:
		return er.GetRequest()
	case deleteEventRequest:
		return er.GetRequest()
	case createRefEventRequest:
		return er.GetRequest()
	case deleteRefEventRequest:
		return er.GetRequest()
	default:
		return nil
	}
}

// Request interface.
type Request interface {
	proto.Message
}

// GetUUID returns uuid of resource being modified by this event.
func (e *Event) GetUUID() string {
	if e == nil {
		return ""
	}
	switch r := e.Unwrap().(type) {
	case CreateRequest:
		return r.GetResource().GetUUID()
	case UpdateRequest:
		return r.GetResource().GetUUID()
	case DeleteRequest:
		return r.GetID()
	case CreateRefRequest:
		return r.GetID()
	case DeleteRefRequest:
		return r.GetID()
	default:
		return ""
	}
}

// Operation returns operation type.
func (e *Event) Operation() string {
	switch e.Unwrap().(type) {
	case CreateRequest:
		return OperationCreate
	case UpdateRequest:
		return OperationUpdate
	case DeleteRequest:
		return OperationDelete
	case CreateRefRequest:
		return string(RefOperationAdd)
	case DeleteRefRequest:
		return string(RefOperationDelete)
	default:
		logrus.Debugf("cannot get event's operation: %v", e)
		return ""
	}
}

// SetFieldMask sets field mask on request if event is of create or update type.
func (e *Event) SetFieldMask(fm types.FieldMask) {
	type fieldMaskSetter interface {
		SetFieldMask(types.FieldMask)
	}

	s, ok := e.Request.(fieldMaskSetter)
	if !ok {
		return
	}
	s.SetFieldMask(fm)
}

// RefOperation is enum type for ref-update operation.
type RefOperation string

// RefOperation values.
const (
	RefOperationAdd    RefOperation = "ADD"
	RefOperationDelete RefOperation = "DELETE"
)

// ParseRefOperation parses RefOperation from string value.
func ParseRefOperation(s string) (op RefOperation) {
	switch s {
	case OperationCreate, string(RefOperationAdd):
		return RefOperationAdd
	case OperationDelete:
		return RefOperationDelete
	default:
		return RefOperation(s)
	}
}

// RefUpdateOption contains parameters for NewRefUpdateEvent.
type RefUpdateOption struct {
	ReferenceType    string
	FromUUID, ToUUID string
	Operation        RefOperation
	Attr             map[string]interface{}
}

// ExtractRefEvents extracts references and puts them into a newly created EventList.
func (e *Event) ExtractRefEvents() (EventList, error) {
	switch r := e.Unwrap().(type) {
	case CreateRequest:
		return extractRefEvents(r.GetResource(), RefOperationAdd)
	case UpdateRequest:
		return EventList{}, nil
	case DeleteRequest:
		//	TODO: Extract event for removing refs from resource before deleting it
		logrus.Warn("Extracting references from DELETE event is not supported yet.")
		return EventList{}, nil
	default:
		return EventList{}, errors.Errorf("cannot extract refs from event %v.", e)
	}
}

func extractRefEvents(r basemodels.Object, o RefOperation) (EventList, error) {
	el, err := makeRefEventList(r, o)
	r.RemoveReferences()
	return el, err
}

func makeRefEventList(r basemodels.Object, operation RefOperation) (EventList, error) {
	el := EventList{}
	for _, ref := range r.GetReferences() {
		var attrMap map[string]interface{}
		if attr := ref.GetAttribute(); attr != nil {
			attrMap = attr.ToMap()
		}
		e, err := NewRefUpdateEvent(RefUpdateOption{
			ReferenceType: basemodels.ReferenceKind(r.Kind(), ref.GetToKind()),
			FromUUID:      r.GetUUID(),
			ToUUID:        ref.GetUUID(),
			Operation:     operation,
			Attr:          attrMap,
		})
		if err != nil {
			return EventList{}, err
		}
		el.Events = append(el.Events, e)
	}
	return el, nil
}

// MarshalJSON marshal event.
func (e *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.ToMap())
}

// MarshalYAML marshal event to yaml.
func (e *Event) MarshalYAML() (interface{}, error) {
	return e.ToMap(), nil
}

// NewEvent makes event from interface.
func NewEvent(option EventOption) (*Event, error) {
	option.Kind = sanitizeKind(option.Kind)

	switch o := sanitizeOperation(option.Operation); o {
	case OperationCreate:
		return newCreateEvent(option)
	case OperationUpdate:
		return newUpdateEvent(option)
	case OperationDelete:
		return newDeleteEvent(option)
	default:
		return nil, errors.Errorf("operation %s not supported", o)
	}
}

func newCreateEvent(option EventOption) (*Event, error) {
	er, err := newEmptyCreateEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r := er.GetRequest()
	r.GetResource().ApplyMap(option.Data)
	r.SetFieldMask(option.getFieldMask())
	return &Event{
		Request: er,
	}, nil
}

func newUpdateEvent(option EventOption) (*Event, error) {
	er, err := newEmptyUpdateEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r := er.GetRequest()
	r.GetResource().ApplyMap(option.Data)
	r.GetResource().SetUUID(option.UUID)
	r.SetFieldMask(option.getFieldMask())
	return &Event{
		Request: er,
	}, nil
}

func newDeleteEvent(option EventOption) (*Event, error) {
	er, err := newEmptyDeleteEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	er.GetRequest().SetID(option.UUID)
	return &Event{
		Request: er,
	}, nil
}

// NewRefUpdateEvent makes ref update event from RefUpdateOption.
func NewRefUpdateEvent(option RefUpdateOption) (*Event, error) {
	switch option.Operation {
	case RefOperationAdd:
		return newCreateRefEvent(option)
	case RefOperationDelete:
		return newDeleteRefEvent(option)
	default:
		return nil, errors.Errorf("operation %s not supported", option.Operation)
	}
}

func newCreateRefEvent(option RefUpdateOption) (*Event, error) {
	er, err := newEmptyCreateRefEventRequest(option.ReferenceType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r := er.GetRequest()
	r.SetID(option.FromUUID)
	r.GetReference().SetUUID(option.ToUUID)
	if err = format.ApplyMap(option.Attr, r.GetReference().GetAttribute()); err != nil {
		return nil, errors.Wrapf(err, "failed to apply attribute data %v, error %v", option.Attr, err)
	}

	return &Event{
		Request: er,
	}, nil
}

func newDeleteRefEvent(option RefUpdateOption) (*Event, error) {
	er, err := newEmptyDeleteRefEventRequest(option.ReferenceType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r := er.GetRequest()
	r.SetID(option.FromUUID)
	r.GetReference().SetUUID(option.ToUUID)
	return &Event{
		Request: er,
	}, nil
}

// ToMap translates event to map.
func (e *Event) ToMap() map[string]interface{} {
	if e == nil || e.Request == nil {
		return nil
	}
	return map[string]interface{}{
		"operation": e.Operation(),
		"kind":      basemodels.KindToSchemaID(e.Kind()),
		"data":      e.data(),
	}
}

// Kind returns kind of a resource inside event. If the event has no resource it returns empty string.
func (e *Event) Kind() string {
	switch r := e.Unwrap().(type) {
	case CreateRequest:
		return r.GetResource().Kind()
	case UpdateRequest:
		return r.GetResource().Kind()
	case DeleteRequest:
		return r.Kind()
	default:
		return ""
	}
}

func (e *Event) data() interface{} {
	switch r := e.Unwrap().(type) {
	case CreateRequest:
		return r.GetResource()
	case UpdateRequest:
		return r.GetResource()
	case DeleteRequest:
		return map[string]interface{}{
			"uuid": r.GetID(),
		}
	default:
		return nil
	}
}

func sanitizeKind(kind string) string {
	return basemodels.SchemaIDToKind(kind)
}

func sanitizeOperation(operation string) string {
	if operation == "" {
		return OperationCreate
	}
	return operation
}

// CreateEventRequest interface.
type createEventRequest interface {
	isEvent_Request
	GetRequest() CreateRequest
}

// CreateRequest interface.
type CreateRequest interface {
	proto.Message
	GetResource() basemodels.Object
	GetFieldMask() types.FieldMask
	SetFieldMask(types.FieldMask)
	isCreateRequest()
}

type updateEventRequest interface {
	isEvent_Request
	GetRequest() UpdateRequest
}

// UpdateRequest interface.
type UpdateRequest interface {
	proto.Message
	GetResource() basemodels.Object
	GetFieldMask() types.FieldMask
	SetFieldMask(types.FieldMask)
	isUpdateRequest()
}

type deleteEventRequest interface {
	isEvent_Request
	GetRequest() DeleteRequest
}

// DeleteRequest interface.
type DeleteRequest interface {
	proto.Message
	GetID() string
	SetID(string)
	Kind() string
	isDeleteRequest()
}

type createRefEventRequest interface {
	isEvent_Request
	GetRequest() CreateRefRequest
}

// CreateRefRequest interface.
type CreateRefRequest interface {
	proto.Message
	GetID() string
	SetID(string)
	GetReference() basemodels.Reference
	isCreateRefRequest()
}

type deleteRefEventRequest interface {
	isEvent_Request
	GetRequest() DeleteRefRequest
}

// DeleteRefRequest interface.
type DeleteRefRequest interface {
	proto.Message
	GetID() string
	SetID(string)
	GetReference() basemodels.Reference
	isDeleteRefRequest()
}

func (o *EventOption) getFieldMask() types.FieldMask {
	if o.FieldMask == nil {
		return basemodels.MapToFieldMask(o.Data)
	}
	return *o.FieldMask
}

// UnmarshalJSON unmarshalls Event.
func (e *Event) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	return e.ApplyMap(m)
}

// UnmarshalYAML unmarshalls Event.
func (e *Event) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var i interface{}
	if err := unmarshal(&i); err != nil {
		return err
	}
	m, ok := fileutil.YAMLtoJSONCompat(i).(map[string]interface{})
	if !ok {
		return errors.Errorf("failed to unmarshal, got invalid data %v", i)
	}
	return e.ApplyMap(m)
}

// ApplyMap applies map onto event.
func (e *Event) ApplyMap(m map[string]interface{}) error {
	data, ok := m["data"].(map[string]interface{})
	if !ok {
		return errors.Errorf("got invalid data %v", m["data"])
	}
	fm := basemodels.MapToFieldMask(data)
	event, err := NewEvent(EventOption{
		UUID:      format.InterfaceToString(data["uuid"]),
		Operation: format.InterfaceToString(m["operation"]),
		Kind:      format.InterfaceToString(m["kind"]),
		Data:      data,
		FieldMask: &fm,
	})
	if err != nil {
		return err
	}
	*e = *event
	return nil
}
