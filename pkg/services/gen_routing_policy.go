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

//RESTCreateRoutingPolicy handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateRoutingPolicy(c echo.Context) error {
	requestData := &models.CreateRoutingPolicyRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_policy",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateRoutingPolicy(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateRoutingPolicy handle a Create API
// nolint
func (service *ContrailService) CreateRoutingPolicy(
	ctx context.Context,
	request *models.CreateRoutingPolicyRequest) (*models.CreateRoutingPolicyResponse, error) {
	model := request.RoutingPolicy
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

	return service.Next().CreateRoutingPolicy(ctx, request)
}

//RESTUpdateRoutingPolicy handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateRoutingPolicy(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateRoutingPolicyRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_policy",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateRoutingPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateRoutingPolicy handles a Update request.
// nolint
func (service *ContrailService) UpdateRoutingPolicy(
	ctx context.Context,
	request *models.UpdateRoutingPolicyRequest) (*models.UpdateRoutingPolicyResponse, error) {
	model := request.RoutingPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateRoutingPolicy(ctx, request)
}

//RESTDeleteRoutingPolicy delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteRoutingPolicy(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteRoutingPolicyRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteRoutingPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetRoutingPolicy a REST Get request.
// nolint
func (service *ContrailService) RESTGetRoutingPolicy(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetRoutingPolicyRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetRoutingPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListRoutingPolicy handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListRoutingPolicy(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListRoutingPolicyRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListRoutingPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
