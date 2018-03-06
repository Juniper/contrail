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

//RESTCreateFirewallRule handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateFirewallRule(c echo.Context) error {
	requestData := &models.CreateFirewallRuleRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "firewall_rule",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateFirewallRule(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateFirewallRule handle a Create API
// nolint
func (service *ContrailService) CreateFirewallRule(
	ctx context.Context,
	request *models.CreateFirewallRuleRequest) (*models.CreateFirewallRuleResponse, error) {
	model := request.FirewallRule
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

	return service.Next().CreateFirewallRule(ctx, request)
}

//RESTUpdateFirewallRule handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateFirewallRule(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateFirewallRuleRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "firewall_rule",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateFirewallRule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateFirewallRule handles a Update request.
// nolint
func (service *ContrailService) UpdateFirewallRule(
	ctx context.Context,
	request *models.UpdateFirewallRuleRequest) (*models.UpdateFirewallRuleResponse, error) {
	model := request.FirewallRule
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateFirewallRule(ctx, request)
}

//RESTDeleteFirewallRule delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteFirewallRule(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteFirewallRuleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteFirewallRule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetFirewallRule a REST Get request.
// nolint
func (service *ContrailService) RESTGetFirewallRule(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetFirewallRuleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetFirewallRule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListFirewallRule handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListFirewallRule(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListFirewallRuleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListFirewallRule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
