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

//RESTCreateSubnet handle a Create REST service.
func (service *ContrailService) RESTCreateSubnet(c echo.Context) error {
	requestData := &models.CreateSubnetRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "subnet",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateSubnet(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateSubnet handle a Create API
func (service *ContrailService) CreateSubnet(
	ctx context.Context,
	request *models.CreateSubnetRequest) (*models.CreateSubnetResponse, error) {
	model := request.Subnet
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

	return service.Next().CreateSubnet(ctx, request)
}

//RESTUpdateSubnet handles a REST Update request.
func (service *ContrailService) RESTUpdateSubnet(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateSubnetRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "subnet",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateSubnet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateSubnet handles a Update request.
func (service *ContrailService) UpdateSubnet(
	ctx context.Context,
	request *models.UpdateSubnetRequest) (*models.UpdateSubnetResponse, error) {
	model := request.Subnet
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateSubnet(ctx, request)
}

//RESTDeleteSubnet delete a resource using REST service.
func (service *ContrailService) RESTDeleteSubnet(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteSubnetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteSubnet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetSubnet a REST Get request.
func (service *ContrailService) RESTGetSubnet(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetSubnetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetSubnet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListSubnet handles a List REST service Request.
func (service *ContrailService) RESTListSubnet(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListSubnetRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListSubnet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
