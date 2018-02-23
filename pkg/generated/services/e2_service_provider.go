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

//RESTCreateE2ServiceProvider handle a Create REST service.
func (service *ContrailService) RESTCreateE2ServiceProvider(c echo.Context) error {
	requestData := &models.CreateE2ServiceProviderRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "e2_service_provider",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateE2ServiceProvider(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateE2ServiceProvider handle a Create API
func (service *ContrailService) CreateE2ServiceProvider(
	ctx context.Context,
	request *models.CreateE2ServiceProviderRequest) (*models.CreateE2ServiceProviderResponse, error) {
	model := request.E2ServiceProvider
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
			return db.CreateE2ServiceProvider(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "e2_service_provider",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateE2ServiceProviderResponse{
		E2ServiceProvider: request.E2ServiceProvider,
	}, nil
}

//RESTUpdateE2ServiceProvider handles a REST Update request.
func (service *ContrailService) RESTUpdateE2ServiceProvider(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateE2ServiceProviderRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "e2_service_provider",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateE2ServiceProvider(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateE2ServiceProvider handles a Update request.
func (service *ContrailService) UpdateE2ServiceProvider(
	ctx context.Context,
	request *models.UpdateE2ServiceProviderRequest) (*models.UpdateE2ServiceProviderResponse, error) {
	model := request.E2ServiceProvider
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateE2ServiceProvider(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "e2_service_provider",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateE2ServiceProviderResponse{
		E2ServiceProvider: model,
	}, nil
}

//RESTDeleteE2ServiceProvider delete a resource using REST service.
func (service *ContrailService) RESTDeleteE2ServiceProvider(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteE2ServiceProviderRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteE2ServiceProvider(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteE2ServiceProvider delete a resource.
func (service *ContrailService) DeleteE2ServiceProvider(ctx context.Context, request *models.DeleteE2ServiceProviderRequest) (*models.DeleteE2ServiceProviderResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteE2ServiceProvider(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteE2ServiceProviderResponse{
		ID: request.ID,
	}, nil
}

//RESTGetE2ServiceProvider a REST Get request.
func (service *ContrailService) RESTGetE2ServiceProvider(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetE2ServiceProviderRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetE2ServiceProvider(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetE2ServiceProvider a Get request.
func (service *ContrailService) GetE2ServiceProvider(ctx context.Context, request *models.GetE2ServiceProviderRequest) (response *models.GetE2ServiceProviderResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListE2ServiceProviderRequest{
		Spec: spec,
	}
	var result *models.ListE2ServiceProviderResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListE2ServiceProvider(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.E2ServiceProviders) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetE2ServiceProviderResponse{
		E2ServiceProvider: result.E2ServiceProviders[0],
	}
	return response, nil
}

//RESTListE2ServiceProvider handles a List REST service Request.
func (service *ContrailService) RESTListE2ServiceProvider(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListE2ServiceProviderRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListE2ServiceProvider(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListE2ServiceProvider handles a List service Request.
func (service *ContrailService) ListE2ServiceProvider(
	ctx context.Context,
	request *models.ListE2ServiceProviderRequest) (response *models.ListE2ServiceProviderResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListE2ServiceProvider(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
