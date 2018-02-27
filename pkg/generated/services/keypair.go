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

//RESTCreateKeypair handle a Create REST service.
func (service *ContrailService) RESTCreateKeypair(c echo.Context) error {
	requestData := &models.CreateKeypairRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "keypair",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateKeypair(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateKeypair handle a Create API
func (service *ContrailService) CreateKeypair(
	ctx context.Context,
	request *models.CreateKeypairRequest) (*models.CreateKeypairResponse, error) {
	model := request.Keypair
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

	return service.Next().CreateKeypair(ctx, request)
}

//RESTUpdateKeypair handles a REST Update request.
func (service *ContrailService) RESTUpdateKeypair(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateKeypairRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "keypair",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateKeypair(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateKeypair handles a Update request.
func (service *ContrailService) UpdateKeypair(
	ctx context.Context,
	request *models.UpdateKeypairRequest) (*models.UpdateKeypairResponse, error) {
	model := request.Keypair
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateKeypair(ctx, request)
}

//RESTDeleteKeypair delete a resource using REST service.
func (service *ContrailService) RESTDeleteKeypair(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteKeypairRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteKeypair(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetKeypair a REST Get request.
func (service *ContrailService) RESTGetKeypair(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetKeypairRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetKeypair(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListKeypair handles a List REST service Request.
func (service *ContrailService) RESTListKeypair(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListKeypairRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListKeypair(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
