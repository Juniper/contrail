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

//RESTCreateQosQueue handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateQosQueue(c echo.Context) error {
	requestData := &models.CreateQosQueueRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_queue",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateQosQueue(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateQosQueue handle a Create API
// nolint
func (service *ContrailService) CreateQosQueue(
	ctx context.Context,
	request *models.CreateQosQueueRequest) (*models.CreateQosQueueResponse, error) {
	model := request.QosQueue
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

	return service.Next().CreateQosQueue(ctx, request)
}

//RESTUpdateQosQueue handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdateQosQueue(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateQosQueueRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_queue",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateQosQueue(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateQosQueue handles a Update request.
// nolint
func (service *ContrailService) UpdateQosQueue(
	ctx context.Context,
	request *models.UpdateQosQueueRequest) (*models.UpdateQosQueueResponse, error) {
	model := request.QosQueue
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateQosQueue(ctx, request)
}

//RESTDeleteQosQueue delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeleteQosQueue(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteQosQueueRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteQosQueue(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetQosQueue a REST Get request.
// nolint
func (service *ContrailService) RESTGetQosQueue(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetQosQueueRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetQosQueue(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListQosQueue handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListQosQueue(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListQosQueueRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListQosQueue(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
