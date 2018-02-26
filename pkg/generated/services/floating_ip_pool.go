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

//RESTCreateFloatingIPPool handle a Create REST service.
func (service *ContrailService) RESTCreateFloatingIPPool(c echo.Context) error {
	requestData := &models.CreateFloatingIPPoolRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip_pool",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateFloatingIPPool(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateFloatingIPPool handle a Create API
func (service *ContrailService) CreateFloatingIPPool(
	ctx context.Context,
	request *models.CreateFloatingIPPoolRequest) (*models.CreateFloatingIPPoolResponse, error) {
	model := request.FloatingIPPool
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
			return db.CreateFloatingIPPool(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip_pool",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateFloatingIPPoolResponse{
		FloatingIPPool: request.FloatingIPPool,
	}, nil
}

//RESTUpdateFloatingIPPool handles a REST Update request.
func (service *ContrailService) RESTUpdateFloatingIPPool(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateFloatingIPPoolRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip_pool",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateFloatingIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateFloatingIPPool handles a Update request.
func (service *ContrailService) UpdateFloatingIPPool(
	ctx context.Context,
	request *models.UpdateFloatingIPPoolRequest) (*models.UpdateFloatingIPPoolResponse, error) {
	model := request.FloatingIPPool
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateFloatingIPPool(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip_pool",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateFloatingIPPoolResponse{
		FloatingIPPool: model,
	}, nil
}

//RESTDeleteFloatingIPPool delete a resource using REST service.
func (service *ContrailService) RESTDeleteFloatingIPPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteFloatingIPPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteFloatingIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteFloatingIPPool delete a resource.
func (service *ContrailService) DeleteFloatingIPPool(ctx context.Context, request *models.DeleteFloatingIPPoolRequest) (*models.DeleteFloatingIPPoolResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteFloatingIPPool(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteFloatingIPPoolResponse{
		ID: request.ID,
	}, nil
}

//RESTGetFloatingIPPool a REST Get request.
func (service *ContrailService) RESTGetFloatingIPPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetFloatingIPPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetFloatingIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetFloatingIPPool a Get request.
func (service *ContrailService) GetFloatingIPPool(ctx context.Context, request *models.GetFloatingIPPoolRequest) (response *models.GetFloatingIPPoolResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListFloatingIPPoolRequest{
		Spec: spec,
	}
	var result *models.ListFloatingIPPoolResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListFloatingIPPool(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.FloatingIPPools) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetFloatingIPPoolResponse{
		FloatingIPPool: result.FloatingIPPools[0],
	}
	return response, nil
}

//RESTListFloatingIPPool handles a List REST service Request.
func (service *ContrailService) RESTListFloatingIPPool(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListFloatingIPPoolRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListFloatingIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListFloatingIPPool handles a List service Request.
func (service *ContrailService) ListFloatingIPPool(
	ctx context.Context,
	request *models.ListFloatingIPPoolRequest) (response *models.ListFloatingIPPoolResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListFloatingIPPool(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
