package services

import (
	"net/http"

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
	requestData := &ResourceList{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "{{ schema.ID }}",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	responses, err := requestData.Process(ctx, service)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, responses)
}
