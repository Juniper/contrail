package services

import (
	"net/http"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// nolint
type ContrailService struct {
	serviceif.BaseService
}

//RESTSync handle a bluk Create REST service.
func (service *ContrailService) RESTSync(c echo.Context) error {
	requestData := &models.EventList{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debug("bind failed on sync")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	responses, err := serviceif.ProcessEvents(ctx, service, requestData)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, responses.Events)
}
