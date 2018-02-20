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

//RESTBGPRouterUpdateRequest for update request for REST.
type RESTBGPRouterUpdateRequest struct {
	Data map[string]interface{} `json:"bgp-router"`
}

//RESTCreateBGPRouter handle a Create REST service.
func (service *ContrailService) RESTCreateBGPRouter(c echo.Context) error {
	requestData := &models.CreateBGPRouterRequest{
		BGPRouter: models.MakeBGPRouter(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_router",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBGPRouter(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBGPRouter handle a Create API
func (service *ContrailService) CreateBGPRouter(
	ctx context.Context,
	request *models.CreateBGPRouterRequest) (*models.CreateBGPRouterResponse, error) {
	model := request.BGPRouter
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
			return db.CreateBGPRouter(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_router",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBGPRouterResponse{
		BGPRouter: request.BGPRouter,
	}, nil
}

//RESTUpdateBGPRouter handles a REST Update request.
func (service *ContrailService) RESTUpdateBGPRouter(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBGPRouterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_router",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBGPRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBGPRouter handles a Update request.
func (service *ContrailService) UpdateBGPRouter(
	ctx context.Context,
	request *models.UpdateBGPRouterRequest) (*models.UpdateBGPRouterResponse, error) {
	model := request.BGPRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateBGPRouter(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_router",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBGPRouterResponse{
		BGPRouter: model,
	}, nil
}

//RESTDeleteBGPRouter delete a resource using REST service.
func (service *ContrailService) RESTDeleteBGPRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBGPRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBGPRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteBGPRouter delete a resource.
func (service *ContrailService) DeleteBGPRouter(ctx context.Context, request *models.DeleteBGPRouterRequest) (*models.DeleteBGPRouterResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteBGPRouter(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBGPRouterResponse{
		ID: request.ID,
	}, nil
}

//RESTGetBGPRouter a REST Get request.
func (service *ContrailService) RESTGetBGPRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBGPRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBGPRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetBGPRouter a Get request.
func (service *ContrailService) GetBGPRouter(ctx context.Context, request *models.GetBGPRouterRequest) (response *models.GetBGPRouterResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListBGPRouterRequest{
		Spec: spec,
	}
	var result *models.ListBGPRouterResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListBGPRouter(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BGPRouters) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBGPRouterResponse{
		BGPRouter: result.BGPRouters[0],
	}
	return response, nil
}

//RESTListBGPRouter handles a List REST service Request.
func (service *ContrailService) RESTListBGPRouter(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListBGPRouterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBGPRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListBGPRouter handles a List service Request.
func (service *ContrailService) ListBGPRouter(
	ctx context.Context,
	request *models.ListBGPRouterRequest) (response *models.ListBGPRouterResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListBGPRouter(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
