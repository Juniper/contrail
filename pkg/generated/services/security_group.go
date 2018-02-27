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

//RESTCreateSecurityGroup handle a Create REST service.
func (service *ContrailService) RESTCreateSecurityGroup(c echo.Context) error {
	requestData := &models.CreateSecurityGroupRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateSecurityGroup(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateSecurityGroup handle a Create API
func (service *ContrailService) CreateSecurityGroup(
	ctx context.Context,
	request *models.CreateSecurityGroupRequest) (*models.CreateSecurityGroupResponse, error) {
	model := request.SecurityGroup
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

	return service.Next().CreateSecurityGroup(ctx, request)
}

//RESTUpdateSecurityGroup handles a REST Update request.
func (service *ContrailService) RESTUpdateSecurityGroup(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateSecurityGroupRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateSecurityGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateSecurityGroup handles a Update request.
func (service *ContrailService) UpdateSecurityGroup(
	ctx context.Context,
	request *models.UpdateSecurityGroupRequest) (*models.UpdateSecurityGroupResponse, error) {
	model := request.SecurityGroup
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateSecurityGroup(ctx, request)
}

//RESTDeleteSecurityGroup delete a resource using REST service.
func (service *ContrailService) RESTDeleteSecurityGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteSecurityGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteSecurityGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetSecurityGroup a REST Get request.
func (service *ContrailService) RESTGetSecurityGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetSecurityGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetSecurityGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListSecurityGroup handles a List REST service Request.
func (service *ContrailService) RESTListSecurityGroup(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListSecurityGroupRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListSecurityGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
