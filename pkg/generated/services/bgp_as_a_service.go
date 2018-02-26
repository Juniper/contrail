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

//RESTCreateBGPAsAService handle a Create REST service.
func (service *ContrailService) RESTCreateBGPAsAService(c echo.Context) error {
	requestData := &models.CreateBGPAsAServiceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_as_a_service",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBGPAsAService(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBGPAsAService handle a Create API
func (service *ContrailService) CreateBGPAsAService(
	ctx context.Context,
	request *models.CreateBGPAsAServiceRequest) (*models.CreateBGPAsAServiceResponse, error) {
	model := request.BGPAsAService
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
			return db.CreateBGPAsAService(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_as_a_service",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBGPAsAServiceResponse{
		BGPAsAService: request.BGPAsAService,
	}, nil
}

//RESTUpdateBGPAsAService handles a REST Update request.
func (service *ContrailService) RESTUpdateBGPAsAService(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBGPAsAServiceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_as_a_service",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBGPAsAService(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBGPAsAService handles a Update request.
func (service *ContrailService) UpdateBGPAsAService(
	ctx context.Context,
	request *models.UpdateBGPAsAServiceRequest) (*models.UpdateBGPAsAServiceResponse, error) {
	model := request.BGPAsAService
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateBGPAsAService(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgp_as_a_service",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBGPAsAServiceResponse{
		BGPAsAService: model,
	}, nil
}

//RESTDeleteBGPAsAService delete a resource using REST service.
func (service *ContrailService) RESTDeleteBGPAsAService(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBGPAsAServiceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBGPAsAService(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteBGPAsAService delete a resource.
func (service *ContrailService) DeleteBGPAsAService(ctx context.Context, request *models.DeleteBGPAsAServiceRequest) (*models.DeleteBGPAsAServiceResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteBGPAsAService(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBGPAsAServiceResponse{
		ID: request.ID,
	}, nil
}

//RESTGetBGPAsAService a REST Get request.
func (service *ContrailService) RESTGetBGPAsAService(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBGPAsAServiceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBGPAsAService(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetBGPAsAService a Get request.
func (service *ContrailService) GetBGPAsAService(ctx context.Context, request *models.GetBGPAsAServiceRequest) (response *models.GetBGPAsAServiceResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListBGPAsAServiceRequest{
		Spec: spec,
	}
	var result *models.ListBGPAsAServiceResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListBGPAsAService(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BGPAsAServices) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBGPAsAServiceResponse{
		BGPAsAService: result.BGPAsAServices[0],
	}
	return response, nil
}

//RESTListBGPAsAService handles a List REST service Request.
func (service *ContrailService) RESTListBGPAsAService(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListBGPAsAServiceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBGPAsAService(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListBGPAsAService handles a List service Request.
func (service *ContrailService) ListBGPAsAService(
	ctx context.Context,
	request *models.ListBGPAsAServiceRequest) (response *models.ListBGPAsAServiceResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListBGPAsAService(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
