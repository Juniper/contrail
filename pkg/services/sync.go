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

	if len(events.Events) == 0 {
		return c.JSON(http.StatusOK, events.Events)
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
	switch events.OperationType() {
	case OperationCreate:
		refMap := getRefMapFromEvents(events.Events)
		return syncSort(events, refMap)
	case OperationDelete:
		refMap, err := service.getRefMapFromObjects(c, events.Events)
		if err != nil {
			return nil, err
		}
		return syncSort(events, refMap)
	default:
		logrus.Warn("Sort for operation of other type than CREATE or DELETE is not supported")
		return events, nil
	}
}

func getRefMapFromEvents(events []*Event) map[*Event]basemodels.References {
	refMap := map[*Event]basemodels.References{}
	for _, ev := range events {
		refMap[ev] = ev.getReferences()
	}
	return refMap
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
