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

//RESTCreateTag handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateTag(c echo.Context) error {
	requestData := &models.CreateTagRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateTag(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateTag handle a Create API
// nolint
func (service *ContrailService) CreateTag(
	ctx context.Context,
	request *models.CreateTagRequest) (*models.CreateTagResponse, error) {
	model := request.Tag
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

	return service.Next().CreateTag(ctx, request)
}

//RESTUpdateTag handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateTag(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateTagRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateTag(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateTag handles a Update request.
// nolint
func (service *ContrailService) UpdateTag(
	ctx context.Context,
	request *models.UpdateTagRequest) (*models.UpdateTagResponse, error) {
	model := request.Tag
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateTag(ctx, request)
}

//RESTDeleteTag delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteTag(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteTagRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteTag(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetTag a REST Get request.
// nolint
func (service *ContrailService) RESTGetTag(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetTagRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetTag(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListTag handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListTag(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListTagRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListTag(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
