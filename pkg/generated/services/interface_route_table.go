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

//RESTInterfaceRouteTableUpdateRequest for update request for REST.
type RESTInterfaceRouteTableUpdateRequest struct {
	Data map[string]interface{} `json:"interface-route-table"`
}

//RESTCreateInterfaceRouteTable handle a Create REST service.
func (service *ContrailService) RESTCreateInterfaceRouteTable(c echo.Context) error {
	requestData := &models.CreateInterfaceRouteTableRequest{
		InterfaceRouteTable: models.MakeInterfaceRouteTable(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "interface_route_table",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateInterfaceRouteTable(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateInterfaceRouteTable handle a Create API
func (service *ContrailService) CreateInterfaceRouteTable(
	ctx context.Context,
	request *models.CreateInterfaceRouteTableRequest) (*models.CreateInterfaceRouteTableResponse, error) {
	model := request.InterfaceRouteTable
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
			return db.CreateInterfaceRouteTable(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "interface_route_table",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateInterfaceRouteTableResponse{
		InterfaceRouteTable: request.InterfaceRouteTable,
	}, nil
}

//RESTUpdateInterfaceRouteTable handles a REST Update request.
func (service *ContrailService) RESTUpdateInterfaceRouteTable(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateInterfaceRouteTableRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "interface_route_table",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateInterfaceRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateInterfaceRouteTable handles a Update request.
func (service *ContrailService) UpdateInterfaceRouteTable(
	ctx context.Context,
	request *models.UpdateInterfaceRouteTableRequest) (*models.UpdateInterfaceRouteTableResponse, error) {
	model := request.InterfaceRouteTable
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateInterfaceRouteTable(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "interface_route_table",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateInterfaceRouteTableResponse{
		InterfaceRouteTable: model,
	}, nil
}

//RESTDeleteInterfaceRouteTable delete a resource using REST service.
func (service *ContrailService) RESTDeleteInterfaceRouteTable(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteInterfaceRouteTableRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteInterfaceRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteInterfaceRouteTable delete a resource.
func (service *ContrailService) DeleteInterfaceRouteTable(ctx context.Context, request *models.DeleteInterfaceRouteTableRequest) (*models.DeleteInterfaceRouteTableResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteInterfaceRouteTable(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteInterfaceRouteTableResponse{
		ID: request.ID,
	}, nil
}

//RESTGetInterfaceRouteTable a REST Get request.
func (service *ContrailService) RESTGetInterfaceRouteTable(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetInterfaceRouteTableRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetInterfaceRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetInterfaceRouteTable a Get request.
func (service *ContrailService) GetInterfaceRouteTable(ctx context.Context, request *models.GetInterfaceRouteTableRequest) (response *models.GetInterfaceRouteTableResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListInterfaceRouteTableRequest{
		Spec: spec,
	}
	var result *models.ListInterfaceRouteTableResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListInterfaceRouteTable(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.InterfaceRouteTables) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetInterfaceRouteTableResponse{
		InterfaceRouteTable: result.InterfaceRouteTables[0],
	}
	return response, nil
}

//RESTListInterfaceRouteTable handles a List REST service Request.
func (service *ContrailService) RESTListInterfaceRouteTable(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListInterfaceRouteTableRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListInterfaceRouteTable(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListInterfaceRouteTable handles a List service Request.
func (service *ContrailService) ListInterfaceRouteTable(
	ctx context.Context,
	request *models.ListInterfaceRouteTableRequest) (response *models.ListInterfaceRouteTableResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListInterfaceRouteTable(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
