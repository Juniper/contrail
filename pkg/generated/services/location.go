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

//RESTCreateLocation handle a Create REST service.
func (service *ContrailService) RESTCreateLocation(c echo.Context) error {
	requestData := &models.CreateLocationRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "location",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateLocation(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateLocation handle a Create API
func (service *ContrailService) CreateLocation(
	ctx context.Context,
	request *models.CreateLocationRequest) (*models.CreateLocationResponse, error) {
	model := request.Location
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
			return db.CreateLocation(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "location",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLocationResponse{
		Location: request.Location,
	}, nil
}

//RESTUpdateLocation handles a REST Update request.
func (service *ContrailService) RESTUpdateLocation(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateLocationRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "location",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateLocation(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateLocation handles a Update request.
func (service *ContrailService) UpdateLocation(
	ctx context.Context,
	request *models.UpdateLocationRequest) (*models.UpdateLocationResponse, error) {
	model := request.Location
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateLocation(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "location",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLocationResponse{
		Location: model,
	}, nil
}

//RESTDeleteLocation delete a resource using REST service.
func (service *ContrailService) RESTDeleteLocation(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteLocationRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteLocation(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteLocation delete a resource.
func (service *ContrailService) DeleteLocation(ctx context.Context, request *models.DeleteLocationRequest) (*models.DeleteLocationResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteLocation(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLocationResponse{
		ID: request.ID,
	}, nil
}

//RESTGetLocation a REST Get request.
func (service *ContrailService) RESTGetLocation(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetLocationRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetLocation(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetLocation a Get request.
func (service *ContrailService) GetLocation(ctx context.Context, request *models.GetLocationRequest) (response *models.GetLocationResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListLocationRequest{
		Spec: spec,
	}
	var result *models.ListLocationResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListLocation(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Locations) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLocationResponse{
		Location: result.Locations[0],
	}
	return response, nil
}

//RESTListLocation handles a List REST service Request.
func (service *ContrailService) RESTListLocation(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListLocationRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListLocation(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListLocation handles a List service Request.
func (service *ContrailService) ListLocation(
	ctx context.Context,
	request *models.ListLocationRequest) (response *models.ListLocationResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListLocation(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
