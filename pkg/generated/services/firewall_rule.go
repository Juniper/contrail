package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
)

//RESTCreateFirewallRule handle a Create REST service.
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
		if model.DisplayName == "" {
			return nil, common.ErrorBadRequest("Both of FQName and Display Name is empty")
		}
		model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateFirewallRule(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "firewall_rule",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateFirewallRuleResponse{
		FirewallRule: request.FirewallRule,
	}, nil
}

//RESTUpdateFirewallRule handles a REST Update request.
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
func (service *ContrailService) UpdateFirewallRule(
	ctx context.Context,
	request *models.UpdateFirewallRuleRequest) (*models.UpdateFirewallRuleResponse, error) {
	model := request.FirewallRule
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateFirewallRule(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "firewall_rule",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateFirewallRuleResponse{
		FirewallRule: model,
	}, nil
}

//RESTDeleteFirewallRule delete a resource using REST service.
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

//DeleteFirewallRule delete a resource.
func (service *ContrailService) DeleteFirewallRule(ctx context.Context, request *models.DeleteFirewallRuleRequest) (*models.DeleteFirewallRuleResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteFirewallRule(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteFirewallRuleResponse{
		ID: request.ID,
	}, nil
}

//RESTGetFirewallRule a REST Get request.
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

//GetFirewallRule a Get request.
func (service *ContrailService) GetFirewallRule(ctx context.Context, request *models.GetFirewallRuleRequest) (response *models.GetFirewallRuleResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListFirewallRuleRequest{
		Spec: spec,
	}
	var result *models.ListFirewallRuleResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListFirewallRule(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.FirewallRules) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetFirewallRuleResponse{
		FirewallRule: result.FirewallRules[0],
	}
	return response, nil
}

//RESTListFirewallRule handles a List REST service Request.
func (service *ContrailService) RESTListFirewallRule(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
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

//ListFirewallRule handles a List service Request.
func (service *ContrailService) ListFirewallRule(
	ctx context.Context,
	request *models.ListFirewallRuleRequest) (response *models.ListFirewallRuleResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListFirewallRule(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
