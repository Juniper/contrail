package services

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/ngaut/log"
)

const (
	//OperationCreate for create operaion.
	OperationCreate = "CREATE"
	//OperationUpdate for update operaion.
	OperationUpdate = "UPDATE"
	//OperationDelete for delete operaion.
	OperationDelete = "DELETE"
)

//CanProcessService is interface for process service.
type CanProcessService interface {
	Process(ctx context.Context, service Service) (*Event, error)
}

//Resource is a generic resource interface.
type Resource interface {
	GetUUID() string
	Kind() string
	Depends() []string
	ToMap() map[string]interface{}
}

//EventList has multiple rest requests.
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
	eventMap map[string]*Event, stateGraph map[string]state) (sortedList []*Event, err error) {
	if stateGraph[uuid] == temporaryVisited {
		return nil, fmt.Errorf("dependency loop found in sync request")
	}
	if stateGraph[uuid] == visited {
		return sorted, nil
	}
	stateGraph[uuid] = temporaryVisited
	event := eventMap[uuid]
	depends := event.GetResource().Depends()
	for _, refUUID := range depends {
		sorted, err = visitResource(refUUID, sorted, eventMap, stateGraph)
		if err != nil {
			return nil, err
		}
		break
	}
	stateGraph[uuid] = visited
	sorted = append(sorted, event)
	return sorted, nil
}

//Sort sorts request by dependency using Tarjan's algorithm
func (e *EventList) Sort() (err error) {
	sorted := []*Event{}
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

//Process dispatches resource event to call correcponding service functions.
func (e *Event) Process(ctx context.Context, service Service) (*Event, error) {
	return e.Request.(CanProcessService).Process(ctx, service)
}

//Process process list of events.
func (e *EventList) Process(ctx context.Context, service Service) (*EventList, error) {
	responses := []*Event{}
	for _, event := range e.Events {
		response, err := event.Process(ctx, service)
		if err != nil {
			log.Debug(response, err)
			return nil, err
		}
		responses = append(responses, response)
	}
	return &EventList{
		Events: responses,
	}, nil
}

//GetResource returns event on resource.
func (e *Event) GetResource() Resource {
	if e == nil {
		return nil
	}
	resourceEvent, ok := e.Request.(HasResource)
	if !ok {
		pp.Println(e)
	}
	return resourceEvent.GetResource()
}
