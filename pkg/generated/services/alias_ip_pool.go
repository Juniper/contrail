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

//RESTCreateAliasIPPool handle a Create REST service.
func (service *ContrailService) RESTCreateAliasIPPool(c echo.Context) error {
	requestData := &models.CreateAliasIPPoolRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip_pool",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAliasIPPool(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAliasIPPool handle a Create API
func (service *ContrailService) CreateAliasIPPool(
	ctx context.Context,
	request *models.CreateAliasIPPoolRequest) (*models.CreateAliasIPPoolResponse, error) {
	model := request.AliasIPPool
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
			return db.CreateAliasIPPool(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip_pool",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAliasIPPoolResponse{
		AliasIPPool: request.AliasIPPool,
	}, nil
}

//RESTUpdateAliasIPPool handles a REST Update request.
func (service *ContrailService) RESTUpdateAliasIPPool(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAliasIPPoolRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip_pool",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAliasIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAliasIPPool handles a Update request.
func (service *ContrailService) UpdateAliasIPPool(
	ctx context.Context,
	request *models.UpdateAliasIPPoolRequest) (*models.UpdateAliasIPPoolResponse, error) {
	model := request.AliasIPPool
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateAliasIPPool(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip_pool",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAliasIPPoolResponse{
		AliasIPPool: model,
	}, nil
}

//RESTDeleteAliasIPPool delete a resource using REST service.
func (service *ContrailService) RESTDeleteAliasIPPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAliasIPPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAliasIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteAliasIPPool delete a resource.
func (service *ContrailService) DeleteAliasIPPool(ctx context.Context, request *models.DeleteAliasIPPoolRequest) (*models.DeleteAliasIPPoolResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteAliasIPPool(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAliasIPPoolResponse{
		ID: request.ID,
	}, nil
}

//RESTGetAliasIPPool a REST Get request.
func (service *ContrailService) RESTGetAliasIPPool(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAliasIPPoolRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAliasIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetAliasIPPool a Get request.
func (service *ContrailService) GetAliasIPPool(ctx context.Context, request *models.GetAliasIPPoolRequest) (response *models.GetAliasIPPoolResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListAliasIPPoolRequest{
		Spec: spec,
	}
	var result *models.ListAliasIPPoolResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListAliasIPPool(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.AliasIPPools) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAliasIPPoolResponse{
		AliasIPPool: result.AliasIPPools[0],
	}
	return response, nil
}

//RESTListAliasIPPool handles a List REST service Request.
func (service *ContrailService) RESTListAliasIPPool(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAliasIPPoolRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAliasIPPool(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListAliasIPPool handles a List service Request.
func (service *ContrailService) ListAliasIPPool(
	ctx context.Context,
	request *models.ListAliasIPPoolRequest) (response *models.ListAliasIPPoolResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListAliasIPPool(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
