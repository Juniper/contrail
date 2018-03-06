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

//RESTCreateInstanceIP handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateInstanceIP(c echo.Context) error {
	requestData := &models.CreateInstanceIPRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "instance_ip",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateInstanceIP(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateInstanceIP handle a Create API
// nolint
func (service *ContrailService) CreateInstanceIP(
	ctx context.Context,
	request *models.CreateInstanceIPRequest) (*models.CreateInstanceIPResponse, error) {
	model := request.InstanceIP
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

	return service.Next().CreateInstanceIP(ctx, request)
}

//RESTUpdateInstanceIP handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateInstanceIP(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateInstanceIPRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "instance_ip",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateInstanceIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateInstanceIP handles a Update request.
// nolint
func (service *ContrailService) UpdateInstanceIP(
	ctx context.Context,
	request *models.UpdateInstanceIPRequest) (*models.UpdateInstanceIPResponse, error) {
	model := request.InstanceIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateInstanceIP(ctx, request)
}

//RESTDeleteInstanceIP delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteInstanceIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteInstanceIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteInstanceIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetInstanceIP a REST Get request.
// nolint
func (service *ContrailService) RESTGetInstanceIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetInstanceIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetInstanceIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListInstanceIP handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListInstanceIP(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListInstanceIPRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListInstanceIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
