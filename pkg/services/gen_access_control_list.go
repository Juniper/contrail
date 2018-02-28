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

//RESTCreateAccessControlList handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateAccessControlList(c echo.Context) error {
	requestData := &models.CreateAccessControlListRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "access_control_list",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAccessControlList(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAccessControlList handle a Create API
// nolint
func (service *ContrailService) CreateAccessControlList(
	ctx context.Context,
	request *models.CreateAccessControlListRequest) (*models.CreateAccessControlListResponse, error) {
	model := request.AccessControlList
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

	return service.Next().CreateAccessControlList(ctx, request)
}

//RESTUpdateAccessControlList handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateAccessControlList(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAccessControlListRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "access_control_list",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAccessControlList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAccessControlList handles a Update request.
// nolint
func (service *ContrailService) UpdateAccessControlList(
	ctx context.Context,
	request *models.UpdateAccessControlListRequest) (*models.UpdateAccessControlListResponse, error) {
	model := request.AccessControlList
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateAccessControlList(ctx, request)
}

//RESTDeleteAccessControlList delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteAccessControlList(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAccessControlListRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAccessControlList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetAccessControlList a REST Get request.
// nolint
func (service *ContrailService) RESTGetAccessControlList(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAccessControlListRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAccessControlList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListAccessControlList handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListAccessControlList(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListAccessControlListRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAccessControlList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
