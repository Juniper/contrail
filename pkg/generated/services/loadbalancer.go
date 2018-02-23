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

//RESTCreateLoadbalancer handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancer(c echo.Context) error {
	requestData := &models.CreateLoadbalancerRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLoadbalancer(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancer handle a Create API
func (service *ContrailService) CreateLoadbalancer(
	ctx context.Context,
	request *models.CreateLoadbalancerRequest) (*models.CreateLoadbalancerResponse, error) {
	model := request.Loadbalancer
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
			return db.CreateLoadbalancer(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLoadbalancerResponse{
		Loadbalancer: request.Loadbalancer,
	}, nil
}

//RESTUpdateLoadbalancer handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancer(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLoadbalancerRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLoadbalancer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancer handles a Update request.
func (service *ContrailService) UpdateLoadbalancer(
	ctx context.Context,
	request *models.UpdateLoadbalancerRequest) (*models.UpdateLoadbalancerResponse, error) {
	model := request.Loadbalancer
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateLoadbalancer(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLoadbalancerResponse{
		Loadbalancer: model,
	}, nil
}

//RESTDeleteLoadbalancer delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancer(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLoadbalancerRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLoadbalancer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteLoadbalancer delete a resource.
func (service *ContrailService) DeleteLoadbalancer(ctx context.Context, request *models.DeleteLoadbalancerRequest) (*models.DeleteLoadbalancerResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteLoadbalancer(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLoadbalancerResponse{
		ID: request.ID,
	}, nil
}

//RESTGetLoadbalancer a REST Get request.
func (service *ContrailService) RESTGetLoadbalancer(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLoadbalancerRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLoadbalancer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetLoadbalancer a Get request.
func (service *ContrailService) GetLoadbalancer(ctx context.Context, request *models.GetLoadbalancerRequest) (response *models.GetLoadbalancerResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListLoadbalancerRequest{
		Spec: spec,
	}
	var result *models.ListLoadbalancerResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListLoadbalancer(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Loadbalancers) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLoadbalancerResponse{
		Loadbalancer: result.Loadbalancers[0],
	}
	return response, nil
}

//RESTListLoadbalancer handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancer(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLoadbalancerRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLoadbalancer(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListLoadbalancer handles a List service Request.
func (service *ContrailService) ListLoadbalancer(
	ctx context.Context,
	request *models.ListLoadbalancerRequest) (response *models.ListLoadbalancerResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListLoadbalancer(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
