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

//RESTCreateBridgeDomain handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateBridgeDomain(c echo.Context) error {
	requestData := &models.CreateBridgeDomainRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bridge_domain",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBridgeDomain(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBridgeDomain handle a Create API
// nolint
func (service *ContrailService) CreateBridgeDomain(
	ctx context.Context,
	request *models.CreateBridgeDomainRequest) (*models.CreateBridgeDomainResponse, error) {
	model := request.BridgeDomain
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

	return service.Next().CreateBridgeDomain(ctx, request)
}

//RESTUpdateBridgeDomain handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateBridgeDomain(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBridgeDomainRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bridge_domain",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBridgeDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBridgeDomain handles a Update request.
// nolint
func (service *ContrailService) UpdateBridgeDomain(
	ctx context.Context,
	request *models.UpdateBridgeDomainRequest) (*models.UpdateBridgeDomainResponse, error) {
	model := request.BridgeDomain
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateBridgeDomain(ctx, request)
}

//RESTDeleteBridgeDomain delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteBridgeDomain(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBridgeDomainRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBridgeDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetBridgeDomain a REST Get request.
// nolint
func (service *ContrailService) RESTGetBridgeDomain(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBridgeDomainRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBridgeDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListBridgeDomain handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListBridgeDomain(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListBridgeDomainRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBridgeDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
