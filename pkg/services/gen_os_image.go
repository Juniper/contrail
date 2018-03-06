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

//RESTCreateOsImage handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateOsImage(c echo.Context) error {
	requestData := &models.CreateOsImageRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "os_image",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateOsImage(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateOsImage handle a Create API
// nolint
func (service *ContrailService) CreateOsImage(
	ctx context.Context,
	request *models.CreateOsImageRequest) (*models.CreateOsImageResponse, error) {
	model := request.OsImage
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

	return service.Next().CreateOsImage(ctx, request)
}

//RESTUpdateOsImage handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateOsImage(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateOsImageRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "os_image",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateOsImage(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateOsImage handles a Update request.
// nolint
func (service *ContrailService) UpdateOsImage(
	ctx context.Context,
	request *models.UpdateOsImageRequest) (*models.UpdateOsImageResponse, error) {
	model := request.OsImage
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateOsImage(ctx, request)
}

//RESTDeleteOsImage delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteOsImage(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteOsImageRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteOsImage(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetOsImage a REST Get request.
// nolint
func (service *ContrailService) RESTGetOsImage(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetOsImageRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetOsImage(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListOsImage handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListOsImage(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListOsImageRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListOsImage(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
