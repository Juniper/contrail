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

//RESTCreateContrailConfigNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailConfigNode(c echo.Context) error {
	requestData := &models.CreateContrailConfigNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_config_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailConfigNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailConfigNode handle a Create API
func (service *ContrailService) CreateContrailConfigNode(
	ctx context.Context,
	request *models.CreateContrailConfigNodeRequest) (*models.CreateContrailConfigNodeResponse, error) {
	model := request.ContrailConfigNode
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

	return service.Next().CreateContrailConfigNode(ctx, request)
}

//RESTUpdateContrailConfigNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailConfigNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailConfigNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_config_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailConfigNode handles a Update request.
func (service *ContrailService) UpdateContrailConfigNode(
	ctx context.Context,
	request *models.UpdateContrailConfigNodeRequest) (*models.UpdateContrailConfigNodeResponse, error) {
	model := request.ContrailConfigNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateContrailConfigNode(ctx, request)
}

//RESTDeleteContrailConfigNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailConfigNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailConfigNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetContrailConfigNode a REST Get request.
func (service *ContrailService) RESTGetContrailConfigNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailConfigNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListContrailConfigNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailConfigNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailConfigNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailConfigNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
