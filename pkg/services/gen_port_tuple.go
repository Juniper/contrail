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

//RESTCreatePortTuple handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreatePortTuple(c echo.Context) error {
	requestData := &models.CreatePortTupleRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "port_tuple",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreatePortTuple(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreatePortTuple handle a Create API
// nolint
func (service *ContrailService) CreatePortTuple(
	ctx context.Context,
	request *models.CreatePortTupleRequest) (*models.CreatePortTupleResponse, error) {
	model := request.PortTuple
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

	return service.Next().CreatePortTuple(ctx, request)
}

//RESTUpdatePortTuple handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdatePortTuple(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdatePortTupleRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "port_tuple",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdatePortTuple(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdatePortTuple handles a Update request.
// nolint
func (service *ContrailService) UpdatePortTuple(
	ctx context.Context,
	request *models.UpdatePortTupleRequest) (*models.UpdatePortTupleResponse, error) {
	model := request.PortTuple
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdatePortTuple(ctx, request)
}

//RESTDeletePortTuple delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeletePortTuple(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeletePortTupleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeletePortTuple(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetPortTuple a REST Get request.
// nolint
func (service *ContrailService) RESTGetPortTuple(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetPortTupleRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetPortTuple(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListPortTuple handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListPortTuple(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListPortTupleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListPortTuple(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
