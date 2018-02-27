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

//RESTCreateForwardingClass handle a Create REST service.
func (service *ContrailService) RESTCreateForwardingClass(c echo.Context) error {
	requestData := &models.CreateForwardingClassRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "forwarding_class",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateForwardingClass(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateForwardingClass handle a Create API
func (service *ContrailService) CreateForwardingClass(
	ctx context.Context,
	request *models.CreateForwardingClassRequest) (*models.CreateForwardingClassResponse, error) {
	model := request.ForwardingClass
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

	return service.Next().CreateForwardingClass(ctx, request)
}

//RESTUpdateForwardingClass handles a REST Update request.
func (service *ContrailService) RESTUpdateForwardingClass(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateForwardingClassRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "forwarding_class",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateForwardingClass(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateForwardingClass handles a Update request.
func (service *ContrailService) UpdateForwardingClass(
	ctx context.Context,
	request *models.UpdateForwardingClassRequest) (*models.UpdateForwardingClassResponse, error) {
	model := request.ForwardingClass
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateForwardingClass(ctx, request)
}

//RESTDeleteForwardingClass delete a resource using REST service.
func (service *ContrailService) RESTDeleteForwardingClass(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteForwardingClassRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteForwardingClass(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetForwardingClass a REST Get request.
func (service *ContrailService) RESTGetForwardingClass(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetForwardingClassRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetForwardingClass(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListForwardingClass handles a List REST service Request.
func (service *ContrailService) RESTListForwardingClass(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListForwardingClassRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListForwardingClass(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
