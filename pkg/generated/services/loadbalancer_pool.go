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

//RESTLoadbalancerPoolUpdateRequest for update request for REST.
type RESTLoadbalancerPoolUpdateRequest struct {
	Data map[string]interface{} `json:"loadbalancer-pool"`
}

//RESTCreateLoadbalancerPool handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancerPool(c echo.Context) error {
	requestData := &models.CreateLoadbalancerPoolRequest{
		LoadbalancerPool: models.MakeLoadbalancerPool(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_pool",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLoadbalancerPool(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerPool handle a Create API
func (service *ContrailService) CreateLoadbalancerPool(
	ctx context.Context,
	request *models.CreateLoadbalancerPoolRequest) (*models.CreateLoadbalancerPoolResponse, error) {
	model := request.LoadbalancerPool
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
			return db.CreateLoadbalancerPool(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_pool",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLoadbalancerPoolResponse{
		LoadbalancerPool: request.LoadbalancerPool,
	}, nil
}

//RESTUpdateLoadbalancerPool handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancerPool(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLoadbalancerPoolRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_pool",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLoadbalancerPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerPool handles a Update request.
func (service *ContrailService) UpdateLoadbalancerPool(
	ctx context.Context,
	request *models.UpdateLoadbalancerPoolRequest) (*models.UpdateLoadbalancerPoolResponse, error) {
	model := request.LoadbalancerPool
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateLoadbalancerPool(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_pool",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLoadbalancerPoolResponse{
		LoadbalancerPool: model,
	}, nil
}

//RESTDeleteLoadbalancerPool delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancerPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLoadbalancerPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLoadbalancerPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteLoadbalancerPool delete a resource.
func (service *ContrailService) DeleteLoadbalancerPool(ctx context.Context, request *models.DeleteLoadbalancerPoolRequest) (*models.DeleteLoadbalancerPoolResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteLoadbalancerPool(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLoadbalancerPoolResponse{
		ID: request.ID,
	}, nil
}

//RESTGetLoadbalancerPool a REST Get request.
func (service *ContrailService) RESTGetLoadbalancerPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLoadbalancerPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLoadbalancerPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetLoadbalancerPool a Get request.
func (service *ContrailService) GetLoadbalancerPool(ctx context.Context, request *models.GetLoadbalancerPoolRequest) (response *models.GetLoadbalancerPoolResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListLoadbalancerPoolRequest{
		Spec: spec,
	}
	var result *models.ListLoadbalancerPoolResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListLoadbalancerPool(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.LoadbalancerPools) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLoadbalancerPoolResponse{
		LoadbalancerPool: result.LoadbalancerPools[0],
	}
	return response, nil
}

//RESTListLoadbalancerPool handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancerPool(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLoadbalancerPoolRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLoadbalancerPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListLoadbalancerPool handles a List service Request.
func (service *ContrailService) ListLoadbalancerPool(
	ctx context.Context,
	request *models.ListLoadbalancerPoolRequest) (response *models.ListLoadbalancerPoolResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListLoadbalancerPool(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
