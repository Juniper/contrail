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

//RESTCreateLoadbalancerMember handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateLoadbalancerMember(c echo.Context) error {
	requestData := &models.CreateLoadbalancerMemberRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_member",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLoadbalancerMember(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerMember handle a Create API
// nolint
func (service *ContrailService) CreateLoadbalancerMember(
	ctx context.Context,
	request *models.CreateLoadbalancerMemberRequest) (*models.CreateLoadbalancerMemberResponse, error) {
	model := request.LoadbalancerMember
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

	return service.Next().CreateLoadbalancerMember(ctx, request)
}

//RESTUpdateLoadbalancerMember handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateLoadbalancerMember(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLoadbalancerMemberRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_member",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLoadbalancerMember(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerMember handles a Update request.
// nolint
func (service *ContrailService) UpdateLoadbalancerMember(
	ctx context.Context,
	request *models.UpdateLoadbalancerMemberRequest) (*models.UpdateLoadbalancerMemberResponse, error) {
	model := request.LoadbalancerMember
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateLoadbalancerMember(ctx, request)
}

//RESTDeleteLoadbalancerMember delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteLoadbalancerMember(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLoadbalancerMemberRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLoadbalancerMember(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetLoadbalancerMember a REST Get request.
// nolint
func (service *ContrailService) RESTGetLoadbalancerMember(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLoadbalancerMemberRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLoadbalancerMember(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListLoadbalancerMember handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListLoadbalancerMember(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListLoadbalancerMemberRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLoadbalancerMember(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
