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

//RESTRouteTargetUpdateRequest for update request for REST.
type RESTRouteTargetUpdateRequest struct {
	Data map[string]interface{} `json:"route-target"`
}

//RESTCreateRouteTarget handle a Create REST service.
func (service *ContrailService) RESTCreateRouteTarget(c echo.Context) error {
	requestData := &models.CreateRouteTargetRequest{
		RouteTarget: models.MakeRouteTarget(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_target",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateRouteTarget(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateRouteTarget handle a Create API
func (service *ContrailService) CreateRouteTarget(
	ctx context.Context,
	request *models.CreateRouteTargetRequest) (*models.CreateRouteTargetResponse, error) {
	model := request.RouteTarget
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
			return db.CreateRouteTarget(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_target",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateRouteTargetResponse{
		RouteTarget: request.RouteTarget,
	}, nil
}

//RESTUpdateRouteTarget handles a REST Update request.
func (service *ContrailService) RESTUpdateRouteTarget(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateRouteTargetRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_target",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateRouteTarget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateRouteTarget handles a Update request.
func (service *ContrailService) UpdateRouteTarget(
	ctx context.Context,
	request *models.UpdateRouteTargetRequest) (*models.UpdateRouteTargetResponse, error) {
	model := request.RouteTarget
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateRouteTarget(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "route_target",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateRouteTargetResponse{
		RouteTarget: model,
	}, nil
}

//RESTDeleteRouteTarget delete a resource using REST service.
func (service *ContrailService) RESTDeleteRouteTarget(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteRouteTargetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteRouteTarget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteRouteTarget delete a resource.
func (service *ContrailService) DeleteRouteTarget(ctx context.Context, request *models.DeleteRouteTargetRequest) (*models.DeleteRouteTargetResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteRouteTarget(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteRouteTargetResponse{
		ID: request.ID,
	}, nil
}

//RESTGetRouteTarget a REST Get request.
func (service *ContrailService) RESTGetRouteTarget(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetRouteTargetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetRouteTarget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetRouteTarget a Get request.
func (service *ContrailService) GetRouteTarget(ctx context.Context, request *models.GetRouteTargetRequest) (response *models.GetRouteTargetResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListRouteTargetRequest{
		Spec: spec,
	}
	var result *models.ListRouteTargetResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListRouteTarget(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.RouteTargets) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetRouteTargetResponse{
		RouteTarget: result.RouteTargets[0],
	}
	return response, nil
}

//RESTListRouteTarget handles a List REST service Request.
func (service *ContrailService) RESTListRouteTarget(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListRouteTargetRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListRouteTarget(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListRouteTarget handles a List service Request.
func (service *ContrailService) ListRouteTarget(
	ctx context.Context,
	request *models.ListRouteTargetRequest) (response *models.ListRouteTargetResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListRouteTarget(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
