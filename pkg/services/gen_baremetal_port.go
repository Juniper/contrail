package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateBaremetalPort handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateBaremetalPort(c echo.Context) error {
	requestData := &models.CreateBaremetalPortRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_port",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBaremetalPort(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBaremetalPort handle a Create API
// nolint
func (service *ContrailService) CreateBaremetalPort(
	ctx context.Context,
	request *models.CreateBaremetalPortRequest) (*models.CreateBaremetalPortResponse, error) {
	model := request.BaremetalPort
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}
	auth := common.GetAuthCTX(ctx)
	if auth == nil {
		return nil, common.ErrorUnauthenticated
	}

	if model.FQName == nil {
		if model.DisplayName != "" {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
		} else {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.UUID}
		}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()

	return service.Next().CreateBaremetalPort(ctx, request)
}

//RESTUpdateBaremetalPort handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateBaremetalPort(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBaremetalPortRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_port",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBaremetalPort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBaremetalPort handles a Update request.
// nolint
func (service *ContrailService) UpdateBaremetalPort(
	ctx context.Context,
	request *models.UpdateBaremetalPortRequest) (*models.UpdateBaremetalPortResponse, error) {
	model := request.BaremetalPort
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateBaremetalPort(ctx, request)
}

//RESTDeleteBaremetalPort delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteBaremetalPort(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBaremetalPortRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBaremetalPort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetBaremetalPort a REST Get request.
// nolint
func (service *ContrailService) RESTGetBaremetalPort(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBaremetalPortRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBaremetalPort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListBaremetalPort handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListBaremetalPort(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListBaremetalPortRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBaremetalPort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
