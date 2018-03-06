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

//RESTCreateNetworkIpam handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateNetworkIpam(c echo.Context) error {
	requestData := &models.CreateNetworkIpamRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_ipam",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateNetworkIpam(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateNetworkIpam handle a Create API
// nolint
func (service *ContrailService) CreateNetworkIpam(
	ctx context.Context,
	request *models.CreateNetworkIpamRequest) (*models.CreateNetworkIpamResponse, error) {
	model := request.NetworkIpam
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

	return service.Next().CreateNetworkIpam(ctx, request)
}

//RESTUpdateNetworkIpam handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateNetworkIpam(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateNetworkIpamRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_ipam",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateNetworkIpam(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateNetworkIpam handles a Update request.
// nolint
func (service *ContrailService) UpdateNetworkIpam(
	ctx context.Context,
	request *models.UpdateNetworkIpamRequest) (*models.UpdateNetworkIpamResponse, error) {
	model := request.NetworkIpam
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateNetworkIpam(ctx, request)
}

//RESTDeleteNetworkIpam delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteNetworkIpam(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteNetworkIpamRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteNetworkIpam(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetNetworkIpam a REST Get request.
// nolint
func (service *ContrailService) RESTGetNetworkIpam(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetNetworkIpamRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetNetworkIpam(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListNetworkIpam handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListNetworkIpam(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListNetworkIpamRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListNetworkIpam(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
