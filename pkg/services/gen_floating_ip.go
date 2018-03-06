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

//RESTCreateFloatingIP handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateFloatingIP(c echo.Context) error {
	requestData := &models.CreateFloatingIPRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateFloatingIP(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateFloatingIP handle a Create API
// nolint
func (service *ContrailService) CreateFloatingIP(
	ctx context.Context,
	request *models.CreateFloatingIPRequest) (*models.CreateFloatingIPResponse, error) {
	model := request.FloatingIP
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

	return service.Next().CreateFloatingIP(ctx, request)
}

//RESTUpdateFloatingIP handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateFloatingIP(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateFloatingIPRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateFloatingIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateFloatingIP handles a Update request.
// nolint
func (service *ContrailService) UpdateFloatingIP(
	ctx context.Context,
	request *models.UpdateFloatingIPRequest) (*models.UpdateFloatingIPResponse, error) {
	model := request.FloatingIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateFloatingIP(ctx, request)
}

//RESTDeleteFloatingIP delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteFloatingIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteFloatingIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteFloatingIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetFloatingIP a REST Get request.
// nolint
func (service *ContrailService) RESTGetFloatingIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetFloatingIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetFloatingIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListFloatingIP handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListFloatingIP(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListFloatingIPRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListFloatingIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
