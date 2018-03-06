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

//RESTCreateDiscoveryServiceAssignment handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateDiscoveryServiceAssignment(c echo.Context) error {
	requestData := &models.CreateDiscoveryServiceAssignmentRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "discovery_service_assignment",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateDiscoveryServiceAssignment(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateDiscoveryServiceAssignment handle a Create API
// nolint
func (service *ContrailService) CreateDiscoveryServiceAssignment(
	ctx context.Context,
	request *models.CreateDiscoveryServiceAssignmentRequest) (*models.CreateDiscoveryServiceAssignmentResponse, error) {
	model := request.DiscoveryServiceAssignment
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

	return service.Next().CreateDiscoveryServiceAssignment(ctx, request)
}

//RESTUpdateDiscoveryServiceAssignment handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateDiscoveryServiceAssignment(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateDiscoveryServiceAssignmentRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "discovery_service_assignment",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateDiscoveryServiceAssignment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateDiscoveryServiceAssignment handles a Update request.
// nolint
func (service *ContrailService) UpdateDiscoveryServiceAssignment(
	ctx context.Context,
	request *models.UpdateDiscoveryServiceAssignmentRequest) (*models.UpdateDiscoveryServiceAssignmentResponse, error) {
	model := request.DiscoveryServiceAssignment
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateDiscoveryServiceAssignment(ctx, request)
}

//RESTDeleteDiscoveryServiceAssignment delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteDiscoveryServiceAssignment(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteDiscoveryServiceAssignmentRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteDiscoveryServiceAssignment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetDiscoveryServiceAssignment a REST Get request.
// nolint
func (service *ContrailService) RESTGetDiscoveryServiceAssignment(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetDiscoveryServiceAssignmentRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetDiscoveryServiceAssignment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListDiscoveryServiceAssignment handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListDiscoveryServiceAssignment(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListDiscoveryServiceAssignmentRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListDiscoveryServiceAssignment(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
