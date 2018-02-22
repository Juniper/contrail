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

//RESTPortTupleUpdateRequest for update request for REST.
type RESTPortTupleUpdateRequest struct {
	Data map[string]interface{} `json:"port-tuple"`
}

//RESTCreatePortTuple handle a Create REST service.
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
			return db.CreatePortTuple(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "port_tuple",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreatePortTupleResponse{
		PortTuple: request.PortTuple,
	}, nil
}

//RESTUpdatePortTuple handles a REST Update request.
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
func (service *ContrailService) UpdatePortTuple(
	ctx context.Context,
	request *models.UpdatePortTupleRequest) (*models.UpdatePortTupleResponse, error) {
	model := request.PortTuple
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdatePortTuple(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "port_tuple",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdatePortTupleResponse{
		PortTuple: model,
	}, nil
}

//RESTDeletePortTuple delete a resource using REST service.
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

//DeletePortTuple delete a resource.
func (service *ContrailService) DeletePortTuple(ctx context.Context, request *models.DeletePortTupleRequest) (*models.DeletePortTupleResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeletePortTuple(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeletePortTupleResponse{
		ID: request.ID,
	}, nil
}

//RESTGetPortTuple a REST Get request.
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

//GetPortTuple a Get request.
func (service *ContrailService) GetPortTuple(ctx context.Context, request *models.GetPortTupleRequest) (response *models.GetPortTupleResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListPortTupleRequest{
		Spec: spec,
	}
	var result *models.ListPortTupleResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListPortTuple(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.PortTuples) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetPortTupleResponse{
		PortTuple: result.PortTuples[0],
	}
	return response, nil
}

//RESTListPortTuple handles a List REST service Request.
func (service *ContrailService) RESTListPortTuple(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
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

//ListPortTuple handles a List service Request.
func (service *ContrailService) ListPortTuple(
	ctx context.Context,
	request *models.ListPortTupleRequest) (response *models.ListPortTupleResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListPortTuple(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
