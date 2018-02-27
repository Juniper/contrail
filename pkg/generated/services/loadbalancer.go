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

//RESTCreateLoadbalancer handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancer(c echo.Context) error {
	requestData := &models.CreateLoadbalancerRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLoadbalancer(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancer handle a Create API
func (service *ContrailService) CreateLoadbalancer(
	ctx context.Context,
	request *models.CreateLoadbalancerRequest) (*models.CreateLoadbalancerResponse, error) {
	model := request.Loadbalancer
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

	return service.Next().CreateLoadbalancer(ctx, request)
}

//RESTUpdateLoadbalancer handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancer(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLoadbalancerRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLoadbalancer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancer handles a Update request.
func (service *ContrailService) UpdateLoadbalancer(
	ctx context.Context,
	request *models.UpdateLoadbalancerRequest) (*models.UpdateLoadbalancerResponse, error) {
	model := request.Loadbalancer
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateLoadbalancer(ctx, request)
}

//RESTDeleteLoadbalancer delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancer(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLoadbalancerRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLoadbalancer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetLoadbalancer a REST Get request.
func (service *ContrailService) RESTGetLoadbalancer(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLoadbalancerRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLoadbalancer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListLoadbalancer handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancer(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLoadbalancerRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLoadbalancer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
