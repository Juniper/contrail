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

//RESTCreateContrailCluster handle a Create REST service.
func (service *ContrailService) RESTCreateContrailCluster(c echo.Context) error {
	requestData := &models.CreateContrailClusterRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_cluster",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailCluster(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailCluster handle a Create API
func (service *ContrailService) CreateContrailCluster(
	ctx context.Context,
	request *models.CreateContrailClusterRequest) (*models.CreateContrailClusterResponse, error) {
	model := request.ContrailCluster
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

	return service.Next().CreateContrailCluster(ctx, request)
}

//RESTUpdateContrailCluster handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailCluster(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailClusterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_cluster",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailCluster handles a Update request.
func (service *ContrailService) UpdateContrailCluster(
	ctx context.Context,
	request *models.UpdateContrailClusterRequest) (*models.UpdateContrailClusterResponse, error) {
	model := request.ContrailCluster
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateContrailCluster(ctx, request)
}

//RESTDeleteContrailCluster delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailCluster(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailClusterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetContrailCluster a REST Get request.
func (service *ContrailService) RESTGetContrailCluster(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailClusterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListContrailCluster handles a List REST service Request.
func (service *ContrailService) RESTListContrailCluster(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailClusterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
