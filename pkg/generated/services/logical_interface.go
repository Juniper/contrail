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

//RESTLogicalInterfaceUpdateRequest for update request for REST.
type RESTLogicalInterfaceUpdateRequest struct {
	Data map[string]interface{} `json:"logical-interface"`
}

//RESTCreateLogicalInterface handle a Create REST service.
func (service *ContrailService) RESTCreateLogicalInterface(c echo.Context) error {
	requestData := &models.CreateLogicalInterfaceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_interface",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLogicalInterface(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLogicalInterface handle a Create API
func (service *ContrailService) CreateLogicalInterface(
	ctx context.Context,
	request *models.CreateLogicalInterfaceRequest) (*models.CreateLogicalInterfaceResponse, error) {
	model := request.LogicalInterface
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
			return db.CreateLogicalInterface(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_interface",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLogicalInterfaceResponse{
		LogicalInterface: request.LogicalInterface,
	}, nil
}

//RESTUpdateLogicalInterface handles a REST Update request.
func (service *ContrailService) RESTUpdateLogicalInterface(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLogicalInterfaceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_interface",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLogicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLogicalInterface handles a Update request.
func (service *ContrailService) UpdateLogicalInterface(
	ctx context.Context,
	request *models.UpdateLogicalInterfaceRequest) (*models.UpdateLogicalInterfaceResponse, error) {
	model := request.LogicalInterface
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateLogicalInterface(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "logical_interface",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLogicalInterfaceResponse{
		LogicalInterface: model,
	}, nil
}

//RESTDeleteLogicalInterface delete a resource using REST service.
func (service *ContrailService) RESTDeleteLogicalInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLogicalInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLogicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteLogicalInterface delete a resource.
func (service *ContrailService) DeleteLogicalInterface(ctx context.Context, request *models.DeleteLogicalInterfaceRequest) (*models.DeleteLogicalInterfaceResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteLogicalInterface(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLogicalInterfaceResponse{
		ID: request.ID,
	}, nil
}

//RESTGetLogicalInterface a REST Get request.
func (service *ContrailService) RESTGetLogicalInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLogicalInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLogicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetLogicalInterface a Get request.
func (service *ContrailService) GetLogicalInterface(ctx context.Context, request *models.GetLogicalInterfaceRequest) (response *models.GetLogicalInterfaceResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListLogicalInterfaceRequest{
		Spec: spec,
	}
	var result *models.ListLogicalInterfaceResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListLogicalInterface(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.LogicalInterfaces) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLogicalInterfaceResponse{
		LogicalInterface: result.LogicalInterfaces[0],
	}
	return response, nil
}

//RESTListLogicalInterface handles a List REST service Request.
func (service *ContrailService) RESTListLogicalInterface(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLogicalInterfaceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLogicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListLogicalInterface handles a List service Request.
func (service *ContrailService) ListLogicalInterface(
	ctx context.Context,
	request *models.ListLogicalInterfaceRequest) (response *models.ListLogicalInterfaceResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListLogicalInterface(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
