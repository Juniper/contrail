package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// RESTSync handles Sync API request.
func (service *ContrailService) RESTSync(c echo.Context) error {
	events := &EventList{}
	if err := c.Bind(events); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	refMap := make(map[string]basemodels.References)

	switch events.CheckOperationType() {
	case OperationCreate:
		logrus.Warn("CREATE operation not supported yet")
	case OperationDelete:
		for _, ev := range events.Events {
			obj, kind, err := service.getObjectAndType(c.Request().Context(), ev.GetUUID())
			if err != nil {
				return errutil.ToHTTPError(err)
			}
			refMap[ev.GetUUID()] = obj.GetReferences()
			logrus.Info(obj, kind)
		}
		if val, ok := refMap["abc"]; ok == true {
			logrus.Info("ok ", val)
		} else {
			return errutil.ErrorInternal
		}
		logrus.Warn("DELETE operation not supported yet")
	default:
		logrus.Warn("Operation of other type not supported")
	}
	// TODO: Call events.Sort()

	responses, err := events.Process(c.Request().Context(), service)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, responses.Events)
}
