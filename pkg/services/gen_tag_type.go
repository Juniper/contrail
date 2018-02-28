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

//RESTCreateTagType handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateTagType(c echo.Context) error {
	requestData := &models.CreateTagTypeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag_type",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateTagType(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateTagType handle a Create API
// nolint
func (service *ContrailService) CreateTagType(
	ctx context.Context,
	request *models.CreateTagTypeRequest) (*models.CreateTagTypeResponse, error) {
	model := request.TagType
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

	return service.Next().CreateTagType(ctx, request)
}

//RESTUpdateTagType handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateTagType(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateTagTypeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag_type",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateTagType(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateTagType handles a Update request.
// nolint
func (service *ContrailService) UpdateTagType(
	ctx context.Context,
	request *models.UpdateTagTypeRequest) (*models.UpdateTagTypeResponse, error) {
	model := request.TagType
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateTagType(ctx, request)
}

//RESTDeleteTagType delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteTagType(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteTagTypeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteTagType(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetTagType a REST Get request.
// nolint
func (service *ContrailService) RESTGetTagType(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetTagTypeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetTagType(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListTagType handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListTagType(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListTagTypeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListTagType(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
