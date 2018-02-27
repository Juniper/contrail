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

//RESTCreateLoadbalancerPool handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancerPool(c echo.Context) error {
	requestData := &models.CreateLoadbalancerPoolRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_pool",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLoadbalancerPool(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerPool handle a Create API
func (service *ContrailService) CreateLoadbalancerPool(
	ctx context.Context,
	request *models.CreateLoadbalancerPoolRequest) (*models.CreateLoadbalancerPoolResponse, error) {
	model := request.LoadbalancerPool
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

	return service.Next().CreateLoadbalancerPool(ctx, request)
}

//RESTUpdateLoadbalancerPool handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancerPool(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLoadbalancerPoolRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_pool",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLoadbalancerPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerPool handles a Update request.
func (service *ContrailService) UpdateLoadbalancerPool(
	ctx context.Context,
	request *models.UpdateLoadbalancerPoolRequest) (*models.UpdateLoadbalancerPoolResponse, error) {
	model := request.LoadbalancerPool
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateLoadbalancerPool(ctx, request)
}

//RESTDeleteLoadbalancerPool delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancerPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLoadbalancerPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLoadbalancerPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetLoadbalancerPool a REST Get request.
func (service *ContrailService) RESTGetLoadbalancerPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLoadbalancerPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLoadbalancerPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListLoadbalancerPool handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancerPool(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLoadbalancerPoolRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLoadbalancerPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
