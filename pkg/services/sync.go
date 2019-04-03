package services

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/errutil"
)

// RESTSync handles Sync API request.
func (service *ContrailService) RESTSync(c echo.Context) error {
	events := &EventList{}
	if err := c.Bind(events); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	responses, err := service.Sync(c.Request().Context(), events)
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, responses.Events)
}

func (service *ContrailService) Sync(ctx context.Context, events *EventList) (*EventList, error) {
	// TODO: Call events.Sort()

	responses, err := events.Process(ctx, service)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to process events")
	}

	return responses, err
}
