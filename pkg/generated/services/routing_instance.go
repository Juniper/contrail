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

//RESTRoutingInstanceUpdateRequest for update request for REST.
type RESTRoutingInstanceUpdateRequest struct {
	Data map[string]interface{} `json:"routing-instance"`
}

//RESTCreateRoutingInstance handle a Create REST service.
func (service *ContrailService) RESTCreateRoutingInstance(c echo.Context) error {
	requestData := &models.CreateRoutingInstanceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_instance",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateRoutingInstance(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateRoutingInstance handle a Create API
func (service *ContrailService) CreateRoutingInstance(
	ctx context.Context,
	request *models.CreateRoutingInstanceRequest) (*models.CreateRoutingInstanceResponse, error) {
	model := request.RoutingInstance
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
			return db.CreateRoutingInstance(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_instance",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateRoutingInstanceResponse{
		RoutingInstance: request.RoutingInstance,
	}, nil
}

//RESTUpdateRoutingInstance handles a REST Update request.
func (service *ContrailService) RESTUpdateRoutingInstance(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateRoutingInstanceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_instance",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateRoutingInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateRoutingInstance handles a Update request.
func (service *ContrailService) UpdateRoutingInstance(
	ctx context.Context,
	request *models.UpdateRoutingInstanceRequest) (*models.UpdateRoutingInstanceResponse, error) {
	model := request.RoutingInstance
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateRoutingInstance(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_instance",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateRoutingInstanceResponse{
		RoutingInstance: model,
	}, nil
}

//RESTDeleteRoutingInstance delete a resource using REST service.
func (service *ContrailService) RESTDeleteRoutingInstance(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteRoutingInstanceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteRoutingInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteRoutingInstance delete a resource.
func (service *ContrailService) DeleteRoutingInstance(ctx context.Context, request *models.DeleteRoutingInstanceRequest) (*models.DeleteRoutingInstanceResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteRoutingInstance(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteRoutingInstanceResponse{
		ID: request.ID,
	}, nil
}

//RESTGetRoutingInstance a REST Get request.
func (service *ContrailService) RESTGetRoutingInstance(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetRoutingInstanceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetRoutingInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetRoutingInstance a Get request.
func (service *ContrailService) GetRoutingInstance(ctx context.Context, request *models.GetRoutingInstanceRequest) (response *models.GetRoutingInstanceResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListRoutingInstanceRequest{
		Spec: spec,
	}
	var result *models.ListRoutingInstanceResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListRoutingInstance(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.RoutingInstances) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetRoutingInstanceResponse{
		RoutingInstance: result.RoutingInstances[0],
	}
	return response, nil
}

//RESTListRoutingInstance handles a List REST service Request.
func (service *ContrailService) RESTListRoutingInstance(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListRoutingInstanceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListRoutingInstance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListRoutingInstance handles a List service Request.
func (service *ContrailService) ListRoutingInstance(
	ctx context.Context,
	request *models.ListRoutingInstanceRequest) (response *models.ListRoutingInstanceResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListRoutingInstance(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
