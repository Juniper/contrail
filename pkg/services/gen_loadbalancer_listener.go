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

//RESTCreateLoadbalancerListener handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateLoadbalancerListener(c echo.Context) error {
	requestData := &models.CreateLoadbalancerListenerRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_listener",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLoadbalancerListener(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerListener handle a Create API
// nolint
func (service *ContrailService) CreateLoadbalancerListener(
	ctx context.Context,
	request *models.CreateLoadbalancerListenerRequest) (*models.CreateLoadbalancerListenerResponse, error) {
	model := request.LoadbalancerListener
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

	return service.Next().CreateLoadbalancerListener(ctx, request)
}

//RESTUpdateLoadbalancerListener handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateLoadbalancerListener(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLoadbalancerListenerRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_listener",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLoadbalancerListener(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerListener handles a Update request.
// nolint
func (service *ContrailService) UpdateLoadbalancerListener(
	ctx context.Context,
	request *models.UpdateLoadbalancerListenerRequest) (*models.UpdateLoadbalancerListenerResponse, error) {
	model := request.LoadbalancerListener
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateLoadbalancerListener(ctx, request)
}

//RESTDeleteLoadbalancerListener delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteLoadbalancerListener(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLoadbalancerListenerRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLoadbalancerListener(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetLoadbalancerListener a REST Get request.
// nolint
func (service *ContrailService) RESTGetLoadbalancerListener(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLoadbalancerListenerRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLoadbalancerListener(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListLoadbalancerListener handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListLoadbalancerListener(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListLoadbalancerListenerRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLoadbalancerListener(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
