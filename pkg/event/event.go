package event

import (
	"fmt"

	"github.com/pkg/errors"
)

//Operation represents operation type.
type Operation int

const (
	//Create operation
	Create Operation = iota
	//Update operation
	Update
	//Delete operation
	Delete
)

const (
	//OperationCreate for create operation.
	OperationCreate = "CREATE"
	//OperationUpdate for update operation.
	OperationUpdate = "UPDATE"
	//OperationDelete for delete operation.
	OperationDelete = "DELETE"
)

// String returns string form of operation.
func (o Operation) String() string {
	switch o {
	case Create:
		return OperationCreate
	case Update:
		return OperationUpdate
	case Delete:
		return OperationDelete
	}
	return ""
}

// StringToOperation get Operation from string.
func StringToOperation(s string) Operation {
	switch s {
	case OperationCreate:
		return Create
	case OperationUpdate:
		return Update
	case OperationDelete:
		return Delete
	}
	return 0
}

// Option contains options for Event.
type Option struct {
	UUID      string
	Operation string
	Kind      string
	Data      map[string]interface{}
}

// Event defines methods that might be implemented by Event.
type Event interface {
	GetResource() Resource
	Operation() string
}

// Events is a list of events.
type Events []Event

// Resource is a generic resource interface.
type Resource interface {
	GetUUID() string
	GetParentUUID() string
	ToMap() map[string]interface{}
	// Kind returns id of schema. For non resource objects it returns empty string
	Kind() string
	// Depends returns UUIDs of children and back references
	Depends() []string
	// AddDependency adds child/backref to model
	AddDependency(i interface{})
	// RemoveDependency removes child/backref from model
	RemoveDependency(i interface{})
}

type state int

const (
	notVisited state = iota
	visited
	temporaryVisited
)

//reorder request using Tarjan's algorithm
func visitResource(uuid string, sorted Events,
	eventMap map[string]Event, stateGraph map[string]state,
) (sortedList Events, err error) {
	if stateGraph[uuid] == temporaryVisited {
		return nil, errors.New("dependency loop found in sync request")
	}
	if stateGraph[uuid] == visited {
		return sorted, nil
	}
	stateGraph[uuid] = temporaryVisited
	event, found := eventMap[uuid]
	if !found {
		return nil, fmt.Errorf("Resource with uuid: %s not found in eventMap", uuid)
	}
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

// Sort sorts Events by dependency using Tarjan's algorithm.
// TODO: support parent-child relationship while checking dependencies.
func (e Events) Sort() (sorted Events, err error) {
	stateGraph := map[string]state{}
	eventMap := map[string]Event{}
	for _, event := range e {
		uuid := event.GetResource().GetUUID()
		stateGraph[uuid] = notVisited
		eventMap[uuid] = event
	}
	foundNotVisited := true
	for foundNotVisited {
		foundNotVisited = false
		for _, event := range e {
			uuid := event.GetResource().GetUUID()
			state := stateGraph[uuid]
			if state == notVisited {
				sorted, err = visitResource(uuid, sorted, eventMap, stateGraph)
				if err != nil {
					return nil, err
				}
				foundNotVisited = true
				break
			}
		}
	}
	return sorted, nil
}
