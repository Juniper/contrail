package services

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/errutil"
)

// RESTSync handles Sync API request.
func (service *ContrailService) RESTSync(c echo.Context) error {
	events := &EventList{}
	if err := c.Bind(events); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	switch events.CheckOperationType()  {
	case OperationCreate:
		err := events.SortCreateNoCycle()
		if err != nil {
			return errutil.ToHTTPError(err)
		}
	case OperationDelete:
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
