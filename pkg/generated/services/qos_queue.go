package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
)

//RESTQosQueueUpdateRequest for update request for REST.
type RESTQosQueueUpdateRequest struct {
	Data map[string]interface{} `json:"qos-queue"`
}

//RESTCreateQosQueue handle a Create REST service.
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
		if model.DisplayName == "" {
			return nil, common.ErrorBadRequest("Both of FQName and Display Name is empty")
		}
		model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateQosQueue(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_queue",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateQosQueueResponse{
		QosQueue: request.QosQueue,
	}, nil
}

//RESTUpdateQosQueue handles a REST Update request.
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
func (service *ContrailService) UpdateQosQueue(
	ctx context.Context,
	request *models.UpdateQosQueueRequest) (*models.UpdateQosQueueResponse, error) {
	model := request.QosQueue
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateQosQueue(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_queue",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateQosQueueResponse{
		QosQueue: model,
	}, nil
}

//RESTDeleteQosQueue delete a resource using REST service.
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

//DeleteQosQueue delete a resource.
func (service *ContrailService) DeleteQosQueue(ctx context.Context, request *models.DeleteQosQueueRequest) (*models.DeleteQosQueueResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteQosQueue(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteQosQueueResponse{
		ID: request.ID,
	}, nil
}

//RESTGetQosQueue a REST Get request.
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

//GetQosQueue a Get request.
func (service *ContrailService) GetQosQueue(ctx context.Context, request *models.GetQosQueueRequest) (response *models.GetQosQueueResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListQosQueueRequest{
		Spec: spec,
	}
	var result *models.ListQosQueueResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListQosQueue(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.QosQueues) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetQosQueueResponse{
		QosQueue: result.QosQueues[0],
	}
	return response, nil
}

//RESTListQosQueue handles a List REST service Request.
func (service *ContrailService) RESTListQosQueue(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
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

//ListQosQueue handles a List service Request.
func (service *ContrailService) ListQosQueue(
	ctx context.Context,
	request *models.ListQosQueueRequest) (response *models.ListQosQueueResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListQosQueue(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
