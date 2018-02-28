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

//RESTCreateDsaRule handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateDsaRule(c echo.Context) error {
	requestData := &models.CreateDsaRuleRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "dsa_rule",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateDsaRule(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateDsaRule handle a Create API
// nolint
func (service *ContrailService) CreateDsaRule(
	ctx context.Context,
	request *models.CreateDsaRuleRequest) (*models.CreateDsaRuleResponse, error) {
	model := request.DsaRule
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

	return service.Next().CreateDsaRule(ctx, request)
}

//RESTUpdateDsaRule handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateDsaRule(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateDsaRuleRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "dsa_rule",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateDsaRule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateDsaRule handles a Update request.
// nolint
func (service *ContrailService) UpdateDsaRule(
	ctx context.Context,
	request *models.UpdateDsaRuleRequest) (*models.UpdateDsaRuleResponse, error) {
	model := request.DsaRule
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateDsaRule(ctx, request)
}

//RESTDeleteDsaRule delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteDsaRule(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteDsaRuleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteDsaRule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetDsaRule a REST Get request.
// nolint
func (service *ContrailService) RESTGetDsaRule(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetDsaRuleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetDsaRule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListDsaRule handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListDsaRule(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListDsaRuleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListDsaRule(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
