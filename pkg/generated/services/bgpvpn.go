package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateBGPVPN handle a Create REST service.
func (service *ContrailService) RESTCreateBGPVPN(c echo.Context) error {
	requestData := &models.CreateBGPVPNRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgpvpn",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBGPVPN(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBGPVPN handle a Create API
func (service *ContrailService) CreateBGPVPN(
	ctx context.Context,
	request *models.CreateBGPVPNRequest) (*models.CreateBGPVPNResponse, error) {
	model := request.BGPVPN
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

	return service.Next().CreateBGPVPN(ctx, request)
}

//RESTUpdateBGPVPN handles a REST Update request.
func (service *ContrailService) RESTUpdateBGPVPN(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBGPVPNRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgpvpn",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBGPVPN(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBGPVPN handles a Update request.
func (service *ContrailService) UpdateBGPVPN(
	ctx context.Context,
	request *models.UpdateBGPVPNRequest) (*models.UpdateBGPVPNResponse, error) {
	model := request.BGPVPN
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateBGPVPN(ctx, request)
}

//RESTDeleteBGPVPN delete a resource using REST service.
func (service *ContrailService) RESTDeleteBGPVPN(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBGPVPNRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBGPVPN(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetBGPVPN a REST Get request.
func (service *ContrailService) RESTGetBGPVPN(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBGPVPNRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBGPVPN(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListBGPVPN handles a List REST service Request.
func (service *ContrailService) RESTListBGPVPN(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListBGPVPNRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBGPVPN(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
