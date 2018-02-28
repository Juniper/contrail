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

//RESTCreateFlavor handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateFlavor(c echo.Context) error {
	requestData := &models.CreateFlavorRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "flavor",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateFlavor(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateFlavor handle a Create API
// nolint
func (service *ContrailService) CreateFlavor(
	ctx context.Context,
	request *models.CreateFlavorRequest) (*models.CreateFlavorResponse, error) {
	model := request.Flavor
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

	return service.Next().CreateFlavor(ctx, request)
}

//RESTUpdateFlavor handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateFlavor(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateFlavorRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "flavor",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateFlavor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateFlavor handles a Update request.
// nolint
func (service *ContrailService) UpdateFlavor(
	ctx context.Context,
	request *models.UpdateFlavorRequest) (*models.UpdateFlavorResponse, error) {
	model := request.Flavor
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateFlavor(ctx, request)
}

//RESTDeleteFlavor delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteFlavor(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteFlavorRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteFlavor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetFlavor a REST Get request.
// nolint
func (service *ContrailService) RESTGetFlavor(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetFlavorRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetFlavor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListFlavor handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListFlavor(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListFlavorRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListFlavor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
