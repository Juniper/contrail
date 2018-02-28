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

//RESTCreateFloatingIPPool handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateFloatingIPPool(c echo.Context) error {
	requestData := &models.CreateFloatingIPPoolRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip_pool",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateFloatingIPPool(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateFloatingIPPool handle a Create API
// nolint
func (service *ContrailService) CreateFloatingIPPool(
	ctx context.Context,
	request *models.CreateFloatingIPPoolRequest) (*models.CreateFloatingIPPoolResponse, error) {
	model := request.FloatingIPPool
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

	return service.Next().CreateFloatingIPPool(ctx, request)
}

//RESTUpdateFloatingIPPool handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateFloatingIPPool(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateFloatingIPPoolRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip_pool",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateFloatingIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateFloatingIPPool handles a Update request.
// nolint
func (service *ContrailService) UpdateFloatingIPPool(
	ctx context.Context,
	request *models.UpdateFloatingIPPoolRequest) (*models.UpdateFloatingIPPoolResponse, error) {
	model := request.FloatingIPPool
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateFloatingIPPool(ctx, request)
}

//RESTDeleteFloatingIPPool delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteFloatingIPPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteFloatingIPPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteFloatingIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetFloatingIPPool a REST Get request.
// nolint
func (service *ContrailService) RESTGetFloatingIPPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetFloatingIPPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetFloatingIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListFloatingIPPool handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListFloatingIPPool(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListFloatingIPPoolRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListFloatingIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
