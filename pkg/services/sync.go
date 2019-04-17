package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
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

func (service *ContrailService) sortEvents(c echo.Context, events *EventList) (*EventList, error) {
	var err error
	var refMap map[*Event]basemodels.References

	switch events.CheckOperationType() {
	case OperationCreate:
		refMap, err = getRefMapFromEvents(events.Events)
		if err != nil {
			return nil, errutil.ToHTTPError(err)
		}
		return syncSort(events, refMap)
	case OperationDelete:
		refMap, err = service.getRefMapFromObjects(c, events.Events)
		if err != nil {
			return nil, errutil.ToHTTPError(err)
		}
		return syncSort(events, refMap)
	default:
		logrus.Warn("Sort for operation of other type than CREATE or DELETE is not supported")
	}
	return nil, err
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
