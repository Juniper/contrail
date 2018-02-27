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

//RESTCreateNamespace handle a Create REST service.
func (service *ContrailService) RESTCreateNamespace(c echo.Context) error {
	requestData := &models.CreateNamespaceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "namespace",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateNamespace(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateNamespace handle a Create API
func (service *ContrailService) CreateNamespace(
	ctx context.Context,
	request *models.CreateNamespaceRequest) (*models.CreateNamespaceResponse, error) {
	model := request.Namespace
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

	return service.Next().CreateNamespace(ctx, request)
}

//RESTUpdateNamespace handles a REST Update request.
func (service *ContrailService) RESTUpdateNamespace(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateNamespaceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "namespace",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateNamespace(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateNamespace handles a Update request.
func (service *ContrailService) UpdateNamespace(
	ctx context.Context,
	request *models.UpdateNamespaceRequest) (*models.UpdateNamespaceResponse, error) {
	model := request.Namespace
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateNamespace(ctx, request)
}

//RESTDeleteNamespace delete a resource using REST service.
func (service *ContrailService) RESTDeleteNamespace(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteNamespaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteNamespace(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetNamespace a REST Get request.
func (service *ContrailService) RESTGetNamespace(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetNamespaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetNamespace(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListNamespace handles a List REST service Request.
func (service *ContrailService) RESTListNamespace(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListNamespaceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListNamespace(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
