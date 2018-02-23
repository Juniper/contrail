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

//RESTCreateLoadbalancerHealthmonitor handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancerHealthmonitor(c echo.Context) error {
	requestData := &models.CreateLoadbalancerHealthmonitorRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_healthmonitor",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLoadbalancerHealthmonitor(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerHealthmonitor handle a Create API
func (service *ContrailService) CreateLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.CreateLoadbalancerHealthmonitorRequest) (*models.CreateLoadbalancerHealthmonitorResponse, error) {
	model := request.LoadbalancerHealthmonitor
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
			return db.CreateLoadbalancerHealthmonitor(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_healthmonitor",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLoadbalancerHealthmonitorResponse{
		LoadbalancerHealthmonitor: request.LoadbalancerHealthmonitor,
	}, nil
}

//RESTUpdateLoadbalancerHealthmonitor handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancerHealthmonitor(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLoadbalancerHealthmonitorRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_healthmonitor",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLoadbalancerHealthmonitor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerHealthmonitor handles a Update request.
func (service *ContrailService) UpdateLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.UpdateLoadbalancerHealthmonitorRequest) (*models.UpdateLoadbalancerHealthmonitorResponse, error) {
	model := request.LoadbalancerHealthmonitor
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateLoadbalancerHealthmonitor(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_healthmonitor",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLoadbalancerHealthmonitorResponse{
		LoadbalancerHealthmonitor: model,
	}, nil
}

//RESTDeleteLoadbalancerHealthmonitor delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancerHealthmonitor(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLoadbalancerHealthmonitorRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLoadbalancerHealthmonitor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteLoadbalancerHealthmonitor delete a resource.
func (service *ContrailService) DeleteLoadbalancerHealthmonitor(ctx context.Context, request *models.DeleteLoadbalancerHealthmonitorRequest) (*models.DeleteLoadbalancerHealthmonitorResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteLoadbalancerHealthmonitor(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLoadbalancerHealthmonitorResponse{
		ID: request.ID,
	}, nil
}

//RESTGetLoadbalancerHealthmonitor a REST Get request.
func (service *ContrailService) RESTGetLoadbalancerHealthmonitor(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLoadbalancerHealthmonitorRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLoadbalancerHealthmonitor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetLoadbalancerHealthmonitor a Get request.
func (service *ContrailService) GetLoadbalancerHealthmonitor(ctx context.Context, request *models.GetLoadbalancerHealthmonitorRequest) (response *models.GetLoadbalancerHealthmonitorResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListLoadbalancerHealthmonitorRequest{
		Spec: spec,
	}
	var result *models.ListLoadbalancerHealthmonitorResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListLoadbalancerHealthmonitor(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.LoadbalancerHealthmonitors) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLoadbalancerHealthmonitorResponse{
		LoadbalancerHealthmonitor: result.LoadbalancerHealthmonitors[0],
	}
	return response, nil
}

//RESTListLoadbalancerHealthmonitor handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancerHealthmonitor(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLoadbalancerHealthmonitorRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLoadbalancerHealthmonitor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListLoadbalancerHealthmonitor handles a List service Request.
func (service *ContrailService) ListLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.ListLoadbalancerHealthmonitorRequest) (response *models.ListLoadbalancerHealthmonitorResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListLoadbalancerHealthmonitor(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
