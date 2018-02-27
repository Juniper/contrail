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

//RESTCreateSecurityLoggingObject handle a Create REST service.
func (service *ContrailService) RESTCreateSecurityLoggingObject(c echo.Context) error {
	requestData := &models.CreateSecurityLoggingObjectRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_logging_object",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateSecurityLoggingObject(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateSecurityLoggingObject handle a Create API
func (service *ContrailService) CreateSecurityLoggingObject(
	ctx context.Context,
	request *models.CreateSecurityLoggingObjectRequest) (*models.CreateSecurityLoggingObjectResponse, error) {
	model := request.SecurityLoggingObject
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

	return service.Next().CreateSecurityLoggingObject(ctx, request)
}

//RESTUpdateSecurityLoggingObject handles a REST Update request.
func (service *ContrailService) RESTUpdateSecurityLoggingObject(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateSecurityLoggingObjectRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_logging_object",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateSecurityLoggingObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateSecurityLoggingObject handles a Update request.
func (service *ContrailService) UpdateSecurityLoggingObject(
	ctx context.Context,
	request *models.UpdateSecurityLoggingObjectRequest) (*models.UpdateSecurityLoggingObjectResponse, error) {
	model := request.SecurityLoggingObject
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateSecurityLoggingObject(ctx, request)
}

//RESTDeleteSecurityLoggingObject delete a resource using REST service.
func (service *ContrailService) RESTDeleteSecurityLoggingObject(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteSecurityLoggingObjectRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteSecurityLoggingObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetSecurityLoggingObject a REST Get request.
func (service *ContrailService) RESTGetSecurityLoggingObject(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetSecurityLoggingObjectRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetSecurityLoggingObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListSecurityLoggingObject handles a List REST service Request.
func (service *ContrailService) RESTListSecurityLoggingObject(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListSecurityLoggingObjectRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListSecurityLoggingObject(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
