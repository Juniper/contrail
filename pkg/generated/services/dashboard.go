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

//RESTCreateDashboard handle a Create REST service.
func (service *ContrailService) RESTCreateDashboard(c echo.Context) error {
	requestData := &models.CreateDashboardRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "dashboard",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateDashboard(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateDashboard handle a Create API
func (service *ContrailService) CreateDashboard(
	ctx context.Context,
	request *models.CreateDashboardRequest) (*models.CreateDashboardResponse, error) {
	model := request.Dashboard
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
			return db.CreateDashboard(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "dashboard",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateDashboardResponse{
		Dashboard: request.Dashboard,
	}, nil
}

//RESTUpdateDashboard handles a REST Update request.
func (service *ContrailService) RESTUpdateDashboard(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateDashboardRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "dashboard",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateDashboard(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateDashboard handles a Update request.
func (service *ContrailService) UpdateDashboard(
	ctx context.Context,
	request *models.UpdateDashboardRequest) (*models.UpdateDashboardResponse, error) {
	model := request.Dashboard
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateDashboard(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "dashboard",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateDashboardResponse{
		Dashboard: model,
	}, nil
}

//RESTDeleteDashboard delete a resource using REST service.
func (service *ContrailService) RESTDeleteDashboard(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteDashboardRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteDashboard(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteDashboard delete a resource.
func (service *ContrailService) DeleteDashboard(ctx context.Context, request *models.DeleteDashboardRequest) (*models.DeleteDashboardResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteDashboard(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteDashboardResponse{
		ID: request.ID,
	}, nil
}

//RESTGetDashboard a REST Get request.
func (service *ContrailService) RESTGetDashboard(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetDashboardRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetDashboard(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetDashboard a Get request.
func (service *ContrailService) GetDashboard(ctx context.Context, request *models.GetDashboardRequest) (response *models.GetDashboardResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListDashboardRequest{
		Spec: spec,
	}
	var result *models.ListDashboardResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListDashboard(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Dashboards) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetDashboardResponse{
		Dashboard: result.Dashboards[0],
	}
	return response, nil
}

//RESTListDashboard handles a List REST service Request.
func (service *ContrailService) RESTListDashboard(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListDashboardRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListDashboard(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListDashboard handles a List service Request.
func (service *ContrailService) ListDashboard(
	ctx context.Context,
	request *models.ListDashboardRequest) (response *models.ListDashboardResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListDashboard(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
