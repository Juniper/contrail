package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateAliasIPPool handle a Create REST service.
func (service *ContrailService) RESTCreateAliasIPPool(c echo.Context) error {
	requestData := &models.CreateAliasIPPoolRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip_pool",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAliasIPPool(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAliasIPPool handle a Create API
func (service *ContrailService) CreateAliasIPPool(
	ctx context.Context,
	request *models.CreateAliasIPPoolRequest) (*models.CreateAliasIPPoolResponse, error) {
	model := request.AliasIPPool
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

	return service.Next().CreateAliasIPPool(ctx, request)
}

//RESTUpdateAliasIPPool handles a REST Update request.
func (service *ContrailService) RESTUpdateAliasIPPool(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAliasIPPoolRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip_pool",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAliasIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAliasIPPool handles a Update request.
func (service *ContrailService) UpdateAliasIPPool(
	ctx context.Context,
	request *models.UpdateAliasIPPoolRequest) (*models.UpdateAliasIPPoolResponse, error) {
	model := request.AliasIPPool
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateAliasIPPool(ctx, request)
}

//RESTDeleteAliasIPPool delete a resource using REST service.
func (service *ContrailService) RESTDeleteAliasIPPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAliasIPPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAliasIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetAliasIPPool a REST Get request.
func (service *ContrailService) RESTGetAliasIPPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAliasIPPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAliasIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListAliasIPPool handles a List REST service Request.
func (service *ContrailService) RESTListAliasIPPool(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAliasIPPoolRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAliasIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
