package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
)

// RESTSync handles Sync API request.
func (service *ContrailService) RESTSync(c echo.Context) error {
	events := &EventList{}
	if err := c.Bind(events); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	if len(events.Events) == 0 {
		return c.JSON(http.StatusOK, events.Events)
	}

	events, err := sortEvents(c.Request().Context(), service, events)
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	responses, err := events.Process(c.Request().Context(), service, service.InTransactionDoer)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, responses.Events)
}

func sortEvents(ctx context.Context, service *ContrailService, events *EventList) (*EventList, error) {
	switch events.OperationType() {
	case OperationDelete:
		return sortDelete(ctx, service, events)
	case OperationCreate, OperationMixed:
		return sortMixed(events)
	default:
		return events, nil
	}
}

func sortDelete(ctx context.Context, service *ContrailService, events *EventList) (*EventList, error) {
	refMap, err := getRefMapFromObjects(ctx, service, events.Events)
	if err != nil {
		return nil, err
	}
	return syncSort(events, refMap)
}

func getRefMapFromObjects(
	ctx context.Context, service *ContrailService, events []*Event,
) (map[*Event]basemodels.References, error) {
	refMap := make(map[*Event]basemodels.References)
	for i, ev := range events {
		obj, _, err := service.getObjectAndType(ctx, ev.GetUUID())
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

func syncSort(events *EventList, refMap map[*Event]basemodels.References) (*EventList, error) {
	g := NewEventGraph(events.Events, refMap)
	if g.HasCycle() {
		return events, errors.New("cycle found in reference graph")
	}

	if !g.IsSortRequired(events, refMap) {
		return events, nil
	}

	return g.SortEvents(), nil
}

func sortMixed(events *EventList) (*EventList, error) {
	createsList, updatesList, deletesList := events.separateListByOperation()

	if len(deletesList.Events) != 0 {
		logrus.Warn("Sort for events mixed with deletes is not supported.")
		return events, nil
	}

	if len(createsList.Events) != 0 {
		var err error
		refMap := getRefMapFromEvents(createsList.Events)
		events, err = syncSort(createsList, refMap)
		if err != nil {
			return nil, err
		}
	}

	if len(updatesList.Events) != 0 {
		events.Events = append(events.Events, updatesList.Events...)
	}
	return events, nil
}

func (e *EventList) separateListByOperation() (*EventList, *EventList, *EventList) {
	createList, updateList, deleteList := &EventList{}, &EventList{}, &EventList{}
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

func getRefMapFromEvents(events []*Event) map[*Event]basemodels.References {
	refMap := map[*Event]basemodels.References{}
	for _, ev := range events {
		refMap[ev] = ev.getReferences()
	}
	return refMap
}
