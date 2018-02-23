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

//RESTCreateFlavor handle a Create REST service.
func (service *ContrailService) RESTCreateFlavor(c echo.Context) error {
	requestData := &models.CreateFlavorRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "flavor",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateFlavor(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateFlavor handle a Create API
func (service *ContrailService) CreateFlavor(
	ctx context.Context,
	request *models.CreateFlavorRequest) (*models.CreateFlavorResponse, error) {
	model := request.Flavor
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
			return db.CreateFlavor(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "flavor",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateFlavorResponse{
		Flavor: request.Flavor,
	}, nil
}

//RESTUpdateFlavor handles a REST Update request.
func (service *ContrailService) RESTUpdateFlavor(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateFlavorRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "flavor",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateFlavor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateFlavor handles a Update request.
func (service *ContrailService) UpdateFlavor(
	ctx context.Context,
	request *models.UpdateFlavorRequest) (*models.UpdateFlavorResponse, error) {
	model := request.Flavor
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateFlavor(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "flavor",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateFlavorResponse{
		Flavor: model,
	}, nil
}

//RESTDeleteFlavor delete a resource using REST service.
func (service *ContrailService) RESTDeleteFlavor(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteFlavorRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteFlavor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteFlavor delete a resource.
func (service *ContrailService) DeleteFlavor(ctx context.Context, request *models.DeleteFlavorRequest) (*models.DeleteFlavorResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteFlavor(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteFlavorResponse{
		ID: request.ID,
	}, nil
}

//RESTGetFlavor a REST Get request.
func (service *ContrailService) RESTGetFlavor(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetFlavorRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetFlavor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetFlavor a Get request.
func (service *ContrailService) GetFlavor(ctx context.Context, request *models.GetFlavorRequest) (response *models.GetFlavorResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListFlavorRequest{
		Spec: spec,
	}
	var result *models.ListFlavorResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListFlavor(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Flavors) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetFlavorResponse{
		Flavor: result.Flavors[0],
	}
	return response, nil
}

//RESTListFlavor handles a List REST service Request.
func (service *ContrailService) RESTListFlavor(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListFlavorRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListFlavor(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListFlavor handles a List service Request.
func (service *ContrailService) ListFlavor(
	ctx context.Context,
	request *models.ListFlavorRequest) (response *models.ListFlavorResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListFlavor(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
