package services

import (
	"context"
	"encoding/json"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

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
		u.Attr = json.RawMessage(format.MustJSON(attr))
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

// GetResource returns event on resource.
func (e *Event) GetResource() basemodels.Object {
	if e == nil {
		return nil
	}
	switch er := e.Request.(type) {
	case CreateEventRequest:
		return er.GetRequest().GetResource()
	case UpdateEventRequest:
		return er.GetRequest().GetResource()
	default:
		return nil
	}
}

// GetUUID returns event on resource.
func (e *Event) GetUUID() string {
	if e == nil {
		return ""
	}
	switch er := e.Request.(type) {
	case CreateEventRequest:
		return er.GetRequest().GetResource().GetUUID()
	case UpdateEventRequest:
		return er.GetRequest().GetResource().GetUUID()
	case DeleteEventRequest:
		return er.GetRequest().GetID()
	case CreateRefEventRequest:
		return er.GetRequest().GetID()
	case DeleteRefEventRequest:
		return er.GetRequest().GetID()
	default:
		return ""
	}
}

// Operation returns operation type.
func (e *Event) Operation() string {
	if e == nil {
		return ""
	}
	switch e.Request.(type) {
	case CreateEventRequest:
		return OperationCreate
	case UpdateEventRequest:
		return OperationUpdate
	case DeleteEventRequest:
		return OperationDelete
	case CreateRefEventRequest:
		return string(RefOperationAdd)
	case DeleteRefEventRequest:
		return string(RefOperationDelete)
	default:
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
	AttrData         json.RawMessage
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

func extractRefEvents(r basemodels.Object, o RefOperation) (EventList, error) {
	el, err := makeRefEventList(r, o)
	r.RemoveReferences()
	return el, err
}

func makeRefEventList(r basemodels.Object, operation RefOperation) (EventList, error) {
	el := EventList{}
	for _, ref := range r.GetReferences() {
		e, err := NewRefUpdateEvent(&RefUpdateOption{
			ReferenceType: basemodels.ReferenceKind(r.Kind(), ref.GetToKind()),
			FromUUID:      r.GetUUID(),
			ToUUID:        ref.GetUUID(),
			Operation:     operation,
			AttrData:      json.RawMessage(format.MustJSON(ref.GetAttribute())),
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
func NewRefUpdateEvent(option *RefUpdateOption) (*Event, error) {
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
		"kind":      e.kind(),
		"data":      e.data(),
	}
}

func (e *Event) kind() string {
	var kind string
	switch er := e.Request.(type) {
	case CreateEventRequest:
		kind =  er.GetRequest().GetResource().Kind()
	case UpdateEventRequest:
		kind =  er.GetRequest().GetResource().Kind()
	case DeleteEventRequest:
		kind =  er.GetRequest().Kind()
	}
	return basemodels.KindToSchemaID(kind)
}

func (e *Event) data() interface{} {
	switch er := e.Request.(type) {
	case CreateEventRequest:
		return er.GetRequest().GetResource()
	case UpdateEventRequest:
		return er.GetRequest().GetResource()
	case DeleteEventRequest:
		return map[string]interface{}{
			"uuid": er.GetRequest().GetID(),
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

// CreateRefEventRequest interface.
type CreateRefEventRequest interface {
	isEvent_Request
	GetRequest() CreateRefRequest
}

// DeleteRefEventRequest interface.
type DeleteRefEventRequest interface {
	isEvent_Request
	GetRequest() DeleteRefRequest
}

type CreateRequest interface {
	GetResource() basemodels.Object
	GetFieldMask() types.FieldMask
	SetFieldMask(types.FieldMask)
	isCreateRequest()
}

type UpdateRequest interface {
	GetResource() basemodels.Object
	GetFieldMask() types.FieldMask
	SetFieldMask(types.FieldMask)
	isUpdateRequest()
}

type DeleteRequest interface {
	GetID() string
	SetID(string)
	Kind() string
	isDeleteRequest()
}

type CreateRefRequest interface {
	GetID() string
	SetID(string)
	SetRefUUID(string)
	GetAttr() basemodels.RefAttribute
	isCreateRefRequest()
}

type DeleteRefRequest interface {
	GetID() string
	SetID(string)
	SetRefUUID(string)
	isDeleteRefRequest()
}

// NewCreateEvent creates new create event.
func NewCreateEvent(option *EventOption) (*Event, error) {
	er, err := NewEmptyCreateEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	er.GetRequest().GetResource().ApplyMap(option.Data)
	er.GetRequest().SetFieldMask(option.getFieldMask())
	return &Event{
		Request: er,
	}, nil
}

// NewUpdateEvent creates new update event.
func NewUpdateEvent(option *EventOption) (*Event, error) {
	er, err := NewEmptyUpdateEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	er.GetRequest().GetResource().ApplyMap(option.Data)
	er.GetRequest().GetResource().SetUUID(option.UUID)
	er.GetRequest().SetFieldMask(option.getFieldMask())
	return &Event{
		Request: er,
	}, nil
}

// NewDeleteEvent creates new delete event.
func NewDeleteEvent(option *EventOption) (*Event, error) {
	er, err := NewEmptyDeleteEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	er.GetRequest().SetID(option.UUID)
	return &Event{
		Request: er,
	}, nil
}

func NewRefAddEvent(option *RefUpdateOption) (*Event, error) {
	r, err := NewEmptyRefAddEventEventRequest(option.ReferenceType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r.GetRequest().SetID(option.FromUUID)
	r.GetRequest().SetRefUUID(option.ToUUID)
	if attr, ok := r.GetRequest().GetAttr().(basemodels.RefAttribute); ok && option.AttrData != nil {
		err = json.Unmarshal(option.AttrData, attr)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal reference attribute data %v, attr %v, error %v", option.AttrData, attr, err)
		}
	}

	return &Event{
		Request: r,
	}, nil
}

func NewRefDeleteEvent(option *RefUpdateOption) (*Event, error) {
	r, err := NewEmptyRefDeleteEventEventRequest(option.ReferenceType)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	r.GetRequest().SetID(option.FromUUID)
	r.GetRequest().SetRefUUID(option.ToUUID)
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
