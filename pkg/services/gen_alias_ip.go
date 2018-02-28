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

//RESTCreateAliasIP handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateAliasIP(c echo.Context) error {
	requestData := &models.CreateAliasIPRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAliasIP(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAliasIP handle a Create API
// nolint
func (service *ContrailService) CreateAliasIP(
	ctx context.Context,
	request *models.CreateAliasIPRequest) (*models.CreateAliasIPResponse, error) {
	model := request.AliasIP
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

	return service.Next().CreateAliasIP(ctx, request)
}

//RESTUpdateAliasIP handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateAliasIP(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAliasIPRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAliasIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAliasIP handles a Update request.
// nolint
func (service *ContrailService) UpdateAliasIP(
	ctx context.Context,
	request *models.UpdateAliasIPRequest) (*models.UpdateAliasIPResponse, error) {
	model := request.AliasIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateAliasIP(ctx, request)
}

//RESTDeleteAliasIP delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteAliasIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAliasIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAliasIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetAliasIP a REST Get request.
// nolint
func (service *ContrailService) RESTGetAliasIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAliasIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAliasIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListAliasIP handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListAliasIP(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListAliasIPRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAliasIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
