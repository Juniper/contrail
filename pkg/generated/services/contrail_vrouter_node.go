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

//RESTCreateContrailVrouterNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailVrouterNode(c echo.Context) error {
	requestData := &models.CreateContrailVrouterNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_vrouter_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailVrouterNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailVrouterNode handle a Create API
func (service *ContrailService) CreateContrailVrouterNode(
	ctx context.Context,
	request *models.CreateContrailVrouterNodeRequest) (*models.CreateContrailVrouterNodeResponse, error) {
	model := request.ContrailVrouterNode
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

	return service.Next().CreateContrailVrouterNode(ctx, request)
}

//RESTUpdateContrailVrouterNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailVrouterNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailVrouterNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_vrouter_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailVrouterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailVrouterNode handles a Update request.
func (service *ContrailService) UpdateContrailVrouterNode(
	ctx context.Context,
	request *models.UpdateContrailVrouterNodeRequest) (*models.UpdateContrailVrouterNodeResponse, error) {
	model := request.ContrailVrouterNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateContrailVrouterNode(ctx, request)
}

//RESTDeleteContrailVrouterNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailVrouterNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailVrouterNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailVrouterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetContrailVrouterNode a REST Get request.
func (service *ContrailService) RESTGetContrailVrouterNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailVrouterNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailVrouterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListContrailVrouterNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailVrouterNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailVrouterNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailVrouterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
