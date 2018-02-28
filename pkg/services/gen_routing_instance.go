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

//RESTCreateRoutingInstance handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateRoutingInstance(c echo.Context) error {
	requestData := &models.CreateRoutingInstanceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_instance",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateRoutingInstance(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateRoutingInstance handle a Create API
// nolint
func (service *ContrailService) CreateRoutingInstance(
	ctx context.Context,
	request *models.CreateRoutingInstanceRequest) (*models.CreateRoutingInstanceResponse, error) {
	model := request.RoutingInstance
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

	return service.Next().CreateRoutingInstance(ctx, request)
}

//RESTUpdateRoutingInstance handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateRoutingInstance(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateRoutingInstanceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_instance",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateRoutingInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateRoutingInstance handles a Update request.
// nolint
func (service *ContrailService) UpdateRoutingInstance(
	ctx context.Context,
	request *models.UpdateRoutingInstanceRequest) (*models.UpdateRoutingInstanceResponse, error) {
	model := request.RoutingInstance
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateRoutingInstance(ctx, request)
}

//RESTDeleteRoutingInstance delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteRoutingInstance(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteRoutingInstanceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteRoutingInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetRoutingInstance a REST Get request.
// nolint
func (service *ContrailService) RESTGetRoutingInstance(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetRoutingInstanceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetRoutingInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListRoutingInstance handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListRoutingInstance(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListRoutingInstanceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListRoutingInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
