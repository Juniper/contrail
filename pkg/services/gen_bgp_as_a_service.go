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

//RESTCreateBGPAsAService handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateBGPAsAService(c echo.Context) error {
	requestData := &models.CreateBGPAsAServiceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_as_a_service",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBGPAsAService(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBGPAsAService handle a Create API
// nolint
func (service *ContrailService) CreateBGPAsAService(
	ctx context.Context,
	request *models.CreateBGPAsAServiceRequest) (*models.CreateBGPAsAServiceResponse, error) {
	model := request.BGPAsAService
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

	return service.Next().CreateBGPAsAService(ctx, request)
}

//RESTUpdateBGPAsAService handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateBGPAsAService(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBGPAsAServiceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_as_a_service",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBGPAsAService(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBGPAsAService handles a Update request.
// nolint
func (service *ContrailService) UpdateBGPAsAService(
	ctx context.Context,
	request *models.UpdateBGPAsAServiceRequest) (*models.UpdateBGPAsAServiceResponse, error) {
	model := request.BGPAsAService
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateBGPAsAService(ctx, request)
}

//RESTDeleteBGPAsAService delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteBGPAsAService(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBGPAsAServiceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBGPAsAService(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetBGPAsAService a REST Get request.
// nolint
func (service *ContrailService) RESTGetBGPAsAService(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBGPAsAServiceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBGPAsAService(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListBGPAsAService handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListBGPAsAService(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListBGPAsAServiceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBGPAsAService(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
