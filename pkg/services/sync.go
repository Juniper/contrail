package services

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// nolint
type ContrailService struct {
	BaseService
}

//RESTSync handle a bluk Create REST service.
func (service *ContrailService) RESTSync(c echo.Context) error {
	events := &EventList{}
	if err := c.Bind(events); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("bind failed on sync")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	responses, err := events.Process(ctx, service)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, responses.Events)
}
