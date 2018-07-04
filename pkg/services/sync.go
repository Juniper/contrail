package services

import (
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// nolint
type ContrailService struct {
	BaseService
	TypeValidator *models.TypeValidator
}

//RESTSync handles a bulk create request.
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
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, responses.Events)
}
