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

//RESTCreateFirewallPolicy handle a Create REST service.
func (service *ContrailService) RESTCreateFirewallPolicy(c echo.Context) error {
	requestData := &models.CreateFirewallPolicyRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "firewall_policy",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateFirewallPolicy(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateFirewallPolicy handle a Create API
func (service *ContrailService) CreateFirewallPolicy(
	ctx context.Context,
	request *models.CreateFirewallPolicyRequest) (*models.CreateFirewallPolicyResponse, error) {
	model := request.FirewallPolicy
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

	return service.Next().CreateFirewallPolicy(ctx, request)
}

//RESTUpdateFirewallPolicy handles a REST Update request.
func (service *ContrailService) RESTUpdateFirewallPolicy(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateFirewallPolicyRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "firewall_policy",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateFirewallPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateFirewallPolicy handles a Update request.
func (service *ContrailService) UpdateFirewallPolicy(
	ctx context.Context,
	request *models.UpdateFirewallPolicyRequest) (*models.UpdateFirewallPolicyResponse, error) {
	model := request.FirewallPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateFirewallPolicy(ctx, request)
}

//RESTDeleteFirewallPolicy delete a resource using REST service.
func (service *ContrailService) RESTDeleteFirewallPolicy(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteFirewallPolicyRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteFirewallPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetFirewallPolicy a REST Get request.
func (service *ContrailService) RESTGetFirewallPolicy(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetFirewallPolicyRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetFirewallPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListFirewallPolicy handles a List REST service Request.
func (service *ContrailService) RESTListFirewallPolicy(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListFirewallPolicyRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListFirewallPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
