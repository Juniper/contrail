package models

import (
	"fmt"
)

const (
	//OperationCreate for create operaion.
	OperationCreate = "CREATE"
	//OperationUpdate for update operaion.
	OperationUpdate = "UPDATE"
	//OperationDelete for delete operaion.
	OperationDelete = "DELETE"
)

//Resource is a generic resource interface.
type Resource interface {
	GetUUID() string
	Kind() string
	Depends() []string
	ToMap() map[string]interface{}
}

//EventList has multiple rest requests.
type EventList struct {
	Events []*Event `json:"resources"`
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
