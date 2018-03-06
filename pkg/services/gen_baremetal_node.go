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

//RESTCreateBaremetalNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateBaremetalNode(c echo.Context) error {
	requestData := &models.CreateBaremetalNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBaremetalNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBaremetalNode handle a Create API
// nolint
func (service *ContrailService) CreateBaremetalNode(
	ctx context.Context,
	request *models.CreateBaremetalNodeRequest) (*models.CreateBaremetalNodeResponse, error) {
	model := request.BaremetalNode
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

	return service.Next().CreateBaremetalNode(ctx, request)
}

//RESTUpdateBaremetalNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateBaremetalNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBaremetalNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBaremetalNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBaremetalNode handles a Update request.
// nolint
func (service *ContrailService) UpdateBaremetalNode(
	ctx context.Context,
	request *models.UpdateBaremetalNodeRequest) (*models.UpdateBaremetalNodeResponse, error) {
	model := request.BaremetalNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateBaremetalNode(ctx, request)
}

//RESTDeleteBaremetalNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteBaremetalNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBaremetalNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBaremetalNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetBaremetalNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetBaremetalNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBaremetalNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBaremetalNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListBaremetalNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListBaremetalNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListBaremetalNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBaremetalNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
