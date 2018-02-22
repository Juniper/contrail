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

//RESTVirtualRouterUpdateRequest for update request for REST.
type RESTVirtualRouterUpdateRequest struct {
	Data map[string]interface{} `json:"virtual-router"`
}

//RESTCreateVirtualRouter handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualRouter(c echo.Context) error {
	requestData := &models.CreateVirtualRouterRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_router",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualRouter(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualRouter handle a Create API
func (service *ContrailService) CreateVirtualRouter(
	ctx context.Context,
	request *models.CreateVirtualRouterRequest) (*models.CreateVirtualRouterResponse, error) {
	model := request.VirtualRouter
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
			return db.CreateVirtualRouter(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_router",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualRouterResponse{
		VirtualRouter: request.VirtualRouter,
	}, nil
}

//RESTUpdateVirtualRouter handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualRouter(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualRouterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_router",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualRouter handles a Update request.
func (service *ContrailService) UpdateVirtualRouter(
	ctx context.Context,
	request *models.UpdateVirtualRouterRequest) (*models.UpdateVirtualRouterResponse, error) {
	model := request.VirtualRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateVirtualRouter(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_router",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualRouterResponse{
		VirtualRouter: model,
	}, nil
}

//RESTDeleteVirtualRouter delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualRouter delete a resource.
func (service *ContrailService) DeleteVirtualRouter(ctx context.Context, request *models.DeleteVirtualRouterRequest) (*models.DeleteVirtualRouterResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteVirtualRouter(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualRouterResponse{
		ID: request.ID,
	}, nil
}

//RESTGetVirtualRouter a REST Get request.
func (service *ContrailService) RESTGetVirtualRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetVirtualRouter a Get request.
func (service *ContrailService) GetVirtualRouter(ctx context.Context, request *models.GetVirtualRouterRequest) (response *models.GetVirtualRouterResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListVirtualRouterRequest{
		Spec: spec,
	}
	var result *models.ListVirtualRouterResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListVirtualRouter(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualRouters) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualRouterResponse{
		VirtualRouter: result.VirtualRouters[0],
	}
	return response, nil
}

//RESTListVirtualRouter handles a List REST service Request.
func (service *ContrailService) RESTListVirtualRouter(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVirtualRouterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListVirtualRouter handles a List service Request.
func (service *ContrailService) ListVirtualRouter(
	ctx context.Context,
	request *models.ListVirtualRouterRequest) (response *models.ListVirtualRouterResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListVirtualRouter(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
