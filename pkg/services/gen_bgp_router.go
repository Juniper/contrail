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

//RESTCreateBGPRouter handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateBGPRouter(c echo.Context) error {
	requestData := &models.CreateBGPRouterRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_router",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBGPRouter(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBGPRouter handle a Create API
// nolint
func (service *ContrailService) CreateBGPRouter(
	ctx context.Context,
	request *models.CreateBGPRouterRequest) (*models.CreateBGPRouterResponse, error) {
	model := request.BGPRouter
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

	return service.Next().CreateBGPRouter(ctx, request)
}

//RESTUpdateBGPRouter handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateBGPRouter(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBGPRouterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_router",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBGPRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBGPRouter handles a Update request.
// nolint
func (service *ContrailService) UpdateBGPRouter(
	ctx context.Context,
	request *models.UpdateBGPRouterRequest) (*models.UpdateBGPRouterResponse, error) {
	model := request.BGPRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateBGPRouter(ctx, request)
}

//RESTDeleteBGPRouter delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteBGPRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBGPRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBGPRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetBGPRouter a REST Get request.
// nolint
func (service *ContrailService) RESTGetBGPRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBGPRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBGPRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListBGPRouter handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListBGPRouter(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListBGPRouterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBGPRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
