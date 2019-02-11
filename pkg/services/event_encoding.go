package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// EventOperation
type EventOperation string

// Possible operations of events.
const (
	OperationCreate EventOperation = "CREATE"
	OperationUpdate EventOperation = "UPDATE"
	OperationDelete EventOperation = "DELETE"
)

// EventOption contains options for Event.
type EventOption struct {
	UUID      string
	Operation EventOperation
	Kind      string
	Data      map[string]interface{}
	FieldMask *types.FieldMask
}

// ResourceEvent is an event that relates to a resource.
type ResourceEvent interface {
	GetResource() basemodels.Object
	Operation() EventOperation
}
type RefEventOperation string

const (
	RefOperationAdd    RefEventOperation = "ADD"
	RefOperationDelete RefEventOperation = "DELETE"
)

// EventOption contains options for Event.
type RefEventOption struct {
	FromUUID  string
	ToUUID    string
	Operation RefEventOperation
	RefType   string
	Attr      map[string]interface{}
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
	parentUUID := event.GetParentUUID()

	sorted, err = visitResource(parentUUID, sorted, eventMap, stateGraph)
	if err != nil {
		return nil, err
	}

	stateGraph[uuid] = visited
	sorted = append(sorted, event)
	return sorted, nil
}

func (e *Event) GetUUID() string {
	switch r := e.Request.(type) {
	case CreateEventRequest:
		return r.GetRequest().GetResource().GetUUID()
	case UpdateEventRequest:
		return r.GetRequest().GetResource().GetUUID()
	case DeleteEventRequest:
		return r.GetRequest().GetID()
	default:
		return ""
	}
}

func (e *Event) GetParentUUID() string {
	switch r := e.Request.(type) {
	case CreateEventRequest:
		return r.GetRequest().GetResource().GetParentUUID()
	case UpdateEventRequest:
		return r.GetRequest().GetResource().GetParentUUID()
	default:
		return ""
	}
}

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
			state := stateGraph[uuid]
			if state == notVisited {
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

// Operation returns operation type.
func (e *Event) Operation() EventOperation {
	switch e.Request.(type) {
	case CreateEventRequest:
		return OperationCreate
	case UpdateEventRequest:
		return OperationUpdate
	case DeleteEventRequest:
		return OperationDelete
	default:
		return ""
	}
}

func (e *Event) GetResource() basemodels.Object {
	switch r := e.Request.(type) {
	case CreateEventRequest:
		return r.GetRequest().GetResource()
	case UpdateEventRequest:
		return r.GetRequest().GetResource()
	default:
		return nil
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

// ParseRefOperation parses RefOperation from string value.
func ParseRefOperation(s string) (op RefEventOperation) {
	switch s {
	case string(OperationCreate), string(RefOperationAdd):
		return RefOperationAdd
	case string(OperationDelete):
		return RefOperationDelete
	default:
		return RefEventOperation(s)
	}
}

// ExtractRefEvents extracts references and puts them into a newly created EventList.
func (e *Event) ExtractRefEvents() (EventList, error) {
	switch r := e.Request.(type) {
	case CreateEventRequest:
		return extractRefEvents(r.GetRequest().GetResource(), RefOperationAdd)
	case UpdateEventRequest:
		return EventList{}, nil
	case DeleteEventRequest:
		//	TODO: Extract event for removing refs from resource before deleting it
		return EventList{}, nil
	default:
		return EventList{}, errors.Errorf("cannot extract refs from event %v.", e)
	}
}

func extractRefEvents(r basemodels.Object, o RefEventOperation) (EventList, error) {
	el, err := makeRefEventList(r, o)
	r.RemoveReferences()
	return el, err
}

func makeRefEventList(r basemodels.Object, operation RefEventOperation) (EventList, error) {
	el := EventList{}
	for _, ref := range r.GetReferences() {
		var attrMap map[string]interface{}
		if ref.GetAttribute() != nil {
			attrMap = ref.GetAttribute().ToMap()
		}
		e, err := NewRefEvent(&RefEventOption{
			RefType:   basemodels.ReferenceKind(r.Kind(), ref.GetToKind()),
			FromUUID:  r.GetUUID(),
			ToUUID:    ref.GetUUID(),
			Operation: operation,
			Attr:      attrMap,
		})
		if err != nil {
			return EventList{}, errors.Wrapf(err, "failed to process ref: %v", ref)
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
func NewEvent(option *EventOption) (*Event, error) {
	option.Kind = sanitizeKind(option.Kind)

	switch o := sanitizeOperation(option.Operation); o {
	case OperationCreate:
		return NewCreateEvent(option)
	case OperationUpdate:
		return NewUpdateEvent(option)
	case OperationDelete:
		return NewDeleteEvent(option)
	default:
		return nil, errors.Errorf("operation %s not supported", o)
	}
}

// NewEvent makes event from interface.
func NewRefEvent(option *RefEventOption) (*Event, error) {
	switch option.Operation {
	case RefOperationAdd:
		return NewRefAddEvent(option)
	case RefOperationDelete:
		return NewRefDeleteEvent(option)
	default:
		return nil, errors.Errorf("operation %s not supported", option.Operation)
	}
}

// ToMap translates event to map.
func (e *Event) ToMap() map[string]interface{} {
	if e == nil || e.Request == nil {
		return nil
	}
	return map[string]interface{}{
		"operation": e.Operation(),
		"kind":      e.Kind(),
		"data":      e.Data(),
	}
}

func (e *Event) Kind() string {
	var k string
	switch r := e.Request.(type) {
	case CreateEventRequest:
		k = r.GetRequest().GetResource().Kind()
	case UpdateEventRequest:
		k = r.GetRequest().GetResource().Kind()
	case DeleteEventRequest:
		k = r.GetRequest().Kind()
	}
	return basemodels.KindToSchemaID(k)
}

func (e *Event) Data() interface{} {
	switch r := e.Request.(type) {
	case CreateEventRequest:
		return r.GetRequest().GetResource()
	case UpdateEventRequest:
		return r.GetRequest().GetResource()
	case DeleteEventRequest:
		return map[string]interface{}{
			"uuid": r.GetRequest().GetID(),
		}
	default:
		return nil
	}
}

func sanitizeKind(kind string) string {
	return basemodels.SchemaIDToKind(kind)
}

func sanitizeOperation(operation EventOperation) EventOperation {
	if operation == "" {
		return OperationCreate
	}
	return operation
}

// CreateEventRequest interface.
type CreateEventRequest interface {
	isEvent_Request
	GetRequest() CreateRequest
}

// UpdateEventRequest interface.
type UpdateEventRequest interface {
	isEvent_Request
	GetRequest() UpdateRequest
}

// DeleteEventRequest interface.
type DeleteEventRequest interface {
	isEvent_Request
	GetRequest() DeleteRequest
}

// RefAddEventRequest interface.
type RefAddEventRequest interface {
	isEvent_Request
	GetRequest() RefAddRequest
}

// RefDeleteEventRequest interface.
type RefDeleteEventRequest interface {
	isEvent_Request
	GetRequest() RefDeleteRequest
}

// CreateEventRequest interface.
type CreateRequest interface {
	isCreateRequest()
	GetResource() basemodels.Object
	SetFieldMask(types.FieldMask)
	GetFieldMask() types.FieldMask
}

// UpdateEventRequest interface.
type UpdateRequest interface {
	isUpdateRequest()
	GetResource() basemodels.Object
	SetFieldMask(types.FieldMask)
	GetFieldMask() types.FieldMask
}

// DeleteEventRequest interface.
type DeleteRequest interface {
	isDeleteRequest()
	SetID(string)
	GetID() string
	Kind() string
}

// RefAddRequest interface.
type RefAddRequest interface {
	isRefAddRequest()
	SetID(string)
	SetToUUID(string)
	GetAttr() basemodels.RefAttribute
}

// RefDeleteRequest interface.
type RefDeleteRequest interface {
	isRefDeleteRequest()
	SetID(string)
	SetToUUID(string)
}

// NewCreateEvent creates new create event.
func NewCreateEvent(option *EventOption) (*Event, error) {
	r, err := NewEmptyCreateEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r.GetRequest().GetResource().ApplyMap(option.Data)
	r.GetRequest().SetFieldMask(option.getFieldMask())
	return &Event{
		Request: r,
	}, nil
}

// NewUpdateEvent creates new update event.
func NewUpdateEvent(option *EventOption) (*Event, error) {
	r, err := NewEmptyUpdateEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r.GetRequest().GetResource().ApplyMap(option.Data)
	r.GetRequest().GetResource().SetUUID(option.UUID)
	r.GetRequest().SetFieldMask(option.getFieldMask())
	return &Event{
		Request: r,
	}, nil
}

// NewDeleteEvent creates new delete event.
func NewDeleteEvent(option *EventOption) (*Event, error) {
	r, err := NewEmptyDeleteEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r.GetRequest().SetID(option.UUID)
	return &Event{
		Request: r,
	}, nil
}

func NewRefAddEvent(option *RefEventOption) (*Event, error) {
	r, err := NewEmptyRefAddEventEventRequest(option.RefType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r.GetRequest().SetID(option.FromUUID)
	r.GetRequest().SetToUUID(option.ToUUID)
	if attr, ok := r.GetRequest().GetAttr().(basemodels.RefAttribute); ok {
		fmt.Printf("ATTR: %v\n", option.Attr)
		err = format.ApplyMap(option.Attr, attr)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to apply reference attribute data %v, attr %v, error %v", option.Attr, attr, err)
		}
	}
	fmt.Printf("%v\n", r)

	return &Event{
		Request: r,
	}, nil
}

func NewRefDeleteEvent(option *RefEventOption) (*Event, error) {
	r, err := NewEmptyRefDeleteEventEventRequest(option.RefType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r.GetRequest().SetID(option.FromUUID)
	r.GetRequest().SetToUUID(option.ToUUID)
	return &Event{
		Request: r,
	}, nil
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
	event, err := NewEvent(&EventOption{
		UUID:      format.InterfaceToString(data["uuid"]),
		Operation: EventOperation(format.InterfaceToString(m["operation"])),
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
