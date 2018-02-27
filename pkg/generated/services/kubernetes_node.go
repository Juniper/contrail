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

//RESTCreateKubernetesNode handle a Create REST service.
func (service *ContrailService) RESTCreateKubernetesNode(c echo.Context) error {
	requestData := &models.CreateKubernetesNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateKubernetesNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateKubernetesNode handle a Create API
func (service *ContrailService) CreateKubernetesNode(
	ctx context.Context,
	request *models.CreateKubernetesNodeRequest) (*models.CreateKubernetesNodeResponse, error) {
	model := request.KubernetesNode
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

	return service.Next().CreateKubernetesNode(ctx, request)
}

//RESTUpdateKubernetesNode handles a REST Update request.
func (service *ContrailService) RESTUpdateKubernetesNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateKubernetesNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateKubernetesNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateKubernetesNode handles a Update request.
func (service *ContrailService) UpdateKubernetesNode(
	ctx context.Context,
	request *models.UpdateKubernetesNodeRequest) (*models.UpdateKubernetesNodeResponse, error) {
	model := request.KubernetesNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateKubernetesNode(ctx, request)
}

//RESTDeleteKubernetesNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteKubernetesNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteKubernetesNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteKubernetesNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetKubernetesNode a REST Get request.
func (service *ContrailService) RESTGetKubernetesNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetKubernetesNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetKubernetesNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListKubernetesNode handles a List REST service Request.
func (service *ContrailService) RESTListKubernetesNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListKubernetesNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListKubernetesNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
