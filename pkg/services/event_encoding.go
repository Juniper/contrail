package services

import (
	"context"
	"encoding/json"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

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

// HasResource defines methods that might be implemented by Event.
type HasResource interface {
	GetResource() basemodels.Object
	Operation() string
	ExtractRefsEventFromEvent() (*Event, error)
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
	parentUUID := event.GetResource().GetParentUUID()

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
		uuid := event.GetResource().GetUUID()
		stateGraph[uuid] = notVisited
		eventMap[uuid] = event
	}
	foundNotVisited := true
	for foundNotVisited {
		foundNotVisited = false
		for _, event := range e.Events {
			uuid := event.GetResource().GetUUID()
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
	resourceEvent, ok := e.Request.(HasResource)
	if !ok {
		return nil
	}
	return resourceEvent.GetResource()
}

// Operation returns operation type.
func (e *Event) Operation() string {
	if e == nil {
		return ""
	}
	resourceEvent, ok := e.Request.(HasResource)
	if !ok {
		return ""
	}
	return resourceEvent.Operation()
}

// ExtractRefsEventFromEvent extracts references and puts them into a newly created event.
func (e *Event) ExtractRefsEventFromEvent() (*Event, error) {
	if e.Request.(HasResource) == nil {
		return nil, errors.Errorf("Cannot extract refs from event %v.", e.ToMap())
	}
	refEvent, err := e.Request.(HasResource).ExtractRefsEventFromEvent()
	if err != nil {
		return nil, errors.Wrap(err, "extracting references update from event failed")
	}
	return refEvent, nil
}

//MarshalJSON marshal event.
func (e *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.ToMap())
}

//MarshalYAML marshal event to yaml.
func (e *Event) MarshalYAML() (interface{}, error) {
	return e.ToMap(), nil
}

//NewEvent makes event from interface.
func NewEvent(option *EventOption) (*Event, error) {
	option.sanitizeKind()

	switch option.getOperationOrDefault() {
	case OperationCreate:
		return NewCreateEvent(option)
	case OperationUpdate:
		return NewUpdateEvent(option)
	case OperationDelete:
		return NewDeleteEvent(option)
	default:
		return nil, errors.Errorf("operation %s not supported", option.getOperationOrDefault())
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
	return basemodels.KindToSchemaID(e.GetResource().Kind())
}

func (e *Event) data() interface{} {
	switch e.Operation() {
	case OperationCreate:
		return e.GetResource()
	case OperationUpdate:
		return e.GetResource()
	case OperationDelete:
		return map[string]interface{}{
			"uuid": e.GetResource().GetUUID(),
		}
	default:
		return nil
	}
}

// CreateEventRequest interface.
type CreateEventRequest interface {
	isEvent_Request
	GetResource() basemodels.Object
}

// UpdateEventRequest interface.
type UpdateEventRequest interface {
	isEvent_Request
	GetResource() basemodels.Object
	SetFieldMask(types.FieldMask)
}

// DeleteEventRequest interface.
type DeleteEventRequest interface {
	isEvent_Request
	SetID(string)
}

// NewCreateEvent creates new create event.
func NewCreateEvent(option *EventOption) (*Event, error) {
	request, err := NewEmptyCreateEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	request.GetResource().ApplyMap(option.Data)
	return &Event{
		Request: request,
	}, nil
}

// NewUpdateEvent creates new update event.
func NewUpdateEvent(option *EventOption) (*Event, error) {
	request, err := NewEmptyUpdateEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	request.GetResource().ApplyMap(option.Data)
	request.GetResource().SetUUID(option.UUID)
	request.SetFieldMask(option.getFieldMask())
	return &Event{
		Request: request,
	}, nil
}

// NewDeleteEvent creates new delete event.
func NewDeleteEvent(option *EventOption) (*Event, error) {
	request, err := NewEmptyDeleteEventRequest(option.Kind)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create event from option %v", option)
	}
	request.SetID(option.UUID)
	return &Event{
		Request: request,
	}, nil
}

func (o *EventOption) getOperationOrDefault() string {
	if o.Operation == "" {
		return OperationCreate
	}
	return o.Operation
}

func (o *EventOption) getFieldMask() types.FieldMask {
	if o.FieldMask == nil {
		return basemodels.MapToFieldMask(o.Data)
	}
	return *o.FieldMask
}

func (o *EventOption) sanitizeKind() {
	o.Kind = basemodels.SchemaIDToKind(o.Kind)
}
