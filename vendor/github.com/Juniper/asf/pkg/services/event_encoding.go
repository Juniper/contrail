package services

import (
	"context"
	"encoding/json"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/models/basemodels"
)

// Possible operations of events.
const (
	OperationCreate = "CREATE"
	OperationUpdate = "UPDATE"
	OperationDelete = "DELETE"
	OperationMixed  = "MIXED"
	EmptyEventList  = "EMPTY"
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

//OperationType checks if all operations have the same type.
func (e *EventList) OperationType() string {
	if len(e.Events) == 0 {
		return EmptyEventList
	}

	operation := e.Events[0].Operation()
	for _, ev := range e.Events {
		if operation != ev.Operation() {
			return OperationMixed
		}
	}
	return operation
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
func (e *EventList) Process(ctx context.Context, service Service, doer InTransactionDoer) (*EventList, error) {
	var responses []*Event
	err := doer.DoInTransaction(ctx, func(ctx context.Context) error {
		for i, event := range e.Events {
			response, err := event.Process(ctx, service)
			if err != nil {
				return errors.Wrapf(err, "failed to process event at index: %v, operation: '%v', kind '%v', uuid '%v'",
					i, event.Operation(), event.Kind(), event.GetUUID())
			}
			responses = append(responses, response)
		}
		return nil
	})

	if err != nil {
		return &EventList{Events: []*Event{}}, err
	}
	return &EventList{Events: responses}, nil
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

func (e *Event) getReferences() basemodels.References {
	res := e.GetResource()
	if res == nil {
		return nil
	}

	refs := res.GetReferences()

	if parentRef := extractParentAsRef(res); parentRef != nil {
		refs = append(refs, parentRef)
	}
	return refs
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
	if err = r.GetResource().ApplyMap(option.Data); err != nil {
		return nil, err
	}
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
	err = r.GetResource().ApplyMap(option.Data)
	r.GetResource().SetUUID(option.UUID)
	r.SetFieldMask(option.getFieldMask())
	return &Event{
		Request: er,
	}, err
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
	var uuid, operation, kind string
	var err error
	uuid, err = format.InterfaceToStringE(data["uuid"])
	if err != nil {
		return err
	}
	operation, err = format.InterfaceToStringE(m["operation"])
	if err != nil {
		return err
	}
	kind, err = format.InterfaceToStringE(m["kind"])
	if err != nil {
		return err
	}
	fm := basemodels.MapToFieldMask(data)
	event, err := NewEvent(EventOption{
		UUID:      uuid,
		Operation: operation,
		Kind:      kind,
		Data:      data,
		FieldMask: &fm,
	})
	if event != nil && e != nil {
		*e = *event
	}
	return err
}
