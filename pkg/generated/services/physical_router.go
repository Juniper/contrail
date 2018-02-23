package services

import (
	"context"
	"database/sql"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreatePhysicalRouter handle a Create REST service.
func (service *ContrailService) RESTCreatePhysicalRouter(c echo.Context) error {
	requestData := &models.CreatePhysicalRouterRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_router",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreatePhysicalRouter(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreatePhysicalRouter handle a Create API
func (service *ContrailService) CreatePhysicalRouter(
	ctx context.Context,
	request *models.CreatePhysicalRouterRequest) (*models.CreatePhysicalRouterResponse, error) {
	model := request.PhysicalRouter
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
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreatePhysicalRouter(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_router",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreatePhysicalRouterResponse{
		PhysicalRouter: request.PhysicalRouter,
	}, nil
}

//RESTUpdatePhysicalRouter handles a REST Update request.
func (service *ContrailService) RESTUpdatePhysicalRouter(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdatePhysicalRouterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_router",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdatePhysicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdatePhysicalRouter handles a Update request.
func (service *ContrailService) UpdatePhysicalRouter(
	ctx context.Context,
	request *models.UpdatePhysicalRouterRequest) (*models.UpdatePhysicalRouterResponse, error) {
	model := request.PhysicalRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdatePhysicalRouter(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_router",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdatePhysicalRouterResponse{
		PhysicalRouter: model,
	}, nil
}

//RESTDeletePhysicalRouter delete a resource using REST service.
func (service *ContrailService) RESTDeletePhysicalRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeletePhysicalRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeletePhysicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeletePhysicalRouter delete a resource.
func (service *ContrailService) DeletePhysicalRouter(ctx context.Context, request *models.DeletePhysicalRouterRequest) (*models.DeletePhysicalRouterResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeletePhysicalRouter(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeletePhysicalRouterResponse{
		ID: request.ID,
	}, nil
}

//RESTGetPhysicalRouter a REST Get request.
func (service *ContrailService) RESTGetPhysicalRouter(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetPhysicalRouterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetPhysicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetPhysicalRouter a Get request.
func (service *ContrailService) GetPhysicalRouter(ctx context.Context, request *models.GetPhysicalRouterRequest) (response *models.GetPhysicalRouterResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListPhysicalRouterRequest{
		Spec: spec,
	}
	var result *models.ListPhysicalRouterResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListPhysicalRouter(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.PhysicalRouters) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetPhysicalRouterResponse{
		PhysicalRouter: result.PhysicalRouters[0],
	}
	return response, nil
}

//RESTListPhysicalRouter handles a List REST service Request.
func (service *ContrailService) RESTListPhysicalRouter(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListPhysicalRouterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListPhysicalRouter(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListPhysicalRouter handles a List service Request.
func (service *ContrailService) ListPhysicalRouter(
	ctx context.Context,
	request *models.ListPhysicalRouterRequest) (response *models.ListPhysicalRouterResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListPhysicalRouter(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
