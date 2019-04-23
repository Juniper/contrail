package services

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// RESTSync handles Sync API request.
func (service *ContrailService) RESTSync(c echo.Context) error {
	events := &EventList{}
	if err := c.Bind(events); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	events, err := service.sortEvents(c, events)
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	responses, err := events.Process(c.Request().Context(), service)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, responses.Events)
}

func (service *ContrailService) sortEvents(c echo.Context, events *EventList) (*EventList, error) {
	switch events.CheckOperationType() {
	case OperationCreate:
		return sortCreate(events)
	case OperationDelete:
		return service.sortDelete(c, events)
	case "MIXED":
		return sortMixed(events)
	default:
		logrus.Warn("Sort for operation of other type than CREATE or DELETE is not supported")
		return events, nil
	}
}

func sortCreate(e *EventList) (*EventList, error) {
	refMap, err := getRefMapFromEvents(e.Events)
	if err != nil {
		return nil, err
	}
	events, err := syncSort(e, refMap)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func getRefMapFromEvents(events []*Event) (map[*Event]basemodels.References, error) {
	refMap := make(map[*Event]basemodels.References)
	var err error
	for _, ev := range events {
		if refMap[ev], err = ev.getReferences(); err != nil {
			return nil, err
		}
	}
	return refMap, nil
}

func syncSort(events *EventList, refMap map[*Event]basemodels.References) (*EventList, error) {
	g := NewEventGraph(events.Events, refMap)
	if g.CheckCycle() {
		return events, errors.New("cycle found in reference graph")
	}

	if !events.isSortRequired(g, refMap) {
		return events, nil
	}

	return g.SortEvents(), nil
}

func (service *ContrailService) sortDelete(c echo.Context, e *EventList) (*EventList, error) {
	logrus.Info("sorting delete")
	refMap, err := service.getRefMapFromObjects(c, e.Events)
	if err != nil {
		return nil, err
	}
	events, err := syncSort(e, refMap)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (service *ContrailService) getRefMapFromObjects(
	c echo.Context, events []*Event,
) (map[*Event]basemodels.References, error) {
	refMap := make(map[*Event]basemodels.References)
	for i, ev := range events {
		obj, _, err := service.getObjectAndType(c.Request().Context(), ev.GetUUID())
		if err != nil {
			return nil, errors.Wrapf(err,
				"failed to retrieve object for event at index: %v, operation: '%v', kind '%v', uuid '%v'",
				i, ev.Operation(), ev.Kind(), ev.GetUUID())
		}
		refMap[ev] = obj.GetReferences()
		if parentRef := extractParentAsRef(obj); parentRef != nil {
			refMap[ev] = append(refMap[ev], parentRef)
		}
	}
	return refMap, nil
}

func sortMixed(events *EventList) (*EventList, error) {
	createsList, updatesList, deletesList := events.separateListByOperation()
	var err error

	if len(deletesList.Events) != 0 {
		logrus.Warn("Sort for events mixed with deletes is not supported.")
		return events, nil
	}

	if len(createsList.Events) != 0 {
		events, err = sortCreate(&createsList)
		if err != nil {
			return nil, err
		}
	}

	if len(updatesList.Events) != 0 {
		events.Events = append(events.Events, updatesList.Events...)
	}
	return events, nil
}

func (e *EventList) separateListByOperation() (EventList, EventList, EventList) {
	var createList, updateList, deleteList EventList
	for _, event := range e.Events {
		switch event.Operation() {
		case OperationCreate:
			createList.Events = append(createList.Events, event)
		case OperationUpdate:
			updateList.Events = append(updateList.Events, event)
		case OperationDelete:
			deleteList.Events = append(deleteList.Events, event)
		}
	}
	return createList, updateList, deleteList
}
