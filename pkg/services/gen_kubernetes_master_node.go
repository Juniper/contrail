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

//RESTCreateKubernetesMasterNode handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateKubernetesMasterNode(c echo.Context) error {
	requestData := &models.CreateKubernetesMasterNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_master_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateKubernetesMasterNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateKubernetesMasterNode handle a Create API
// nolint
func (service *ContrailService) CreateKubernetesMasterNode(
	ctx context.Context,
	request *models.CreateKubernetesMasterNodeRequest) (*models.CreateKubernetesMasterNodeResponse, error) {
	model := request.KubernetesMasterNode
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

	return service.Next().CreateKubernetesMasterNode(ctx, request)
}

//RESTUpdateKubernetesMasterNode handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateKubernetesMasterNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateKubernetesMasterNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_master_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateKubernetesMasterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateKubernetesMasterNode handles a Update request.
// nolint
func (service *ContrailService) UpdateKubernetesMasterNode(
	ctx context.Context,
	request *models.UpdateKubernetesMasterNodeRequest) (*models.UpdateKubernetesMasterNodeResponse, error) {
	model := request.KubernetesMasterNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateKubernetesMasterNode(ctx, request)
}

//RESTDeleteKubernetesMasterNode delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteKubernetesMasterNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteKubernetesMasterNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteKubernetesMasterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetKubernetesMasterNode a REST Get request.
// nolint
func (service *ContrailService) RESTGetKubernetesMasterNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetKubernetesMasterNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetKubernetesMasterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListKubernetesMasterNode handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListKubernetesMasterNode(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListKubernetesMasterNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListKubernetesMasterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
