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

//RESTCreateDomain handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateDomain(c echo.Context) error {
	requestData := &models.CreateDomainRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "domain",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateDomain(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateDomain handle a Create API
// nolint
func (service *ContrailService) CreateDomain(
	ctx context.Context,
	request *models.CreateDomainRequest) (*models.CreateDomainResponse, error) {
	model := request.Domain
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

	return service.Next().CreateDomain(ctx, request)
}

//RESTUpdateDomain handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateDomain(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateDomainRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "domain",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateDomain handles a Update request.
// nolint
func (service *ContrailService) UpdateDomain(
	ctx context.Context,
	request *models.UpdateDomainRequest) (*models.UpdateDomainResponse, error) {
	model := request.Domain
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateDomain(ctx, request)
}

//RESTDeleteDomain delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteDomain(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteDomainRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetDomain a REST Get request.
// nolint
func (service *ContrailService) RESTGetDomain(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetDomainRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListDomain handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListDomain(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListDomainRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
