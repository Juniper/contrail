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

//RESTCreateKeypair handle a Create REST service.
func (service *ContrailService) RESTCreateKeypair(c echo.Context) error {
	requestData := &models.CreateKeypairRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "keypair",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateKeypair(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateKeypair handle a Create API
func (service *ContrailService) CreateKeypair(
	ctx context.Context,
	request *models.CreateKeypairRequest) (*models.CreateKeypairResponse, error) {
	model := request.Keypair
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
			return db.CreateKeypair(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "keypair",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateKeypairResponse{
		Keypair: request.Keypair,
	}, nil
}

//RESTUpdateKeypair handles a REST Update request.
func (service *ContrailService) RESTUpdateKeypair(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateKeypairRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "keypair",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateKeypair(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateKeypair handles a Update request.
func (service *ContrailService) UpdateKeypair(
	ctx context.Context,
	request *models.UpdateKeypairRequest) (*models.UpdateKeypairResponse, error) {
	model := request.Keypair
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateKeypair(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "keypair",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateKeypairResponse{
		Keypair: model,
	}, nil
}

//RESTDeleteKeypair delete a resource using REST service.
func (service *ContrailService) RESTDeleteKeypair(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteKeypairRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteKeypair(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteKeypair delete a resource.
func (service *ContrailService) DeleteKeypair(ctx context.Context, request *models.DeleteKeypairRequest) (*models.DeleteKeypairResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteKeypair(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteKeypairResponse{
		ID: request.ID,
	}, nil
}

//RESTGetKeypair a REST Get request.
func (service *ContrailService) RESTGetKeypair(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetKeypairRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetKeypair(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetKeypair a Get request.
func (service *ContrailService) GetKeypair(ctx context.Context, request *models.GetKeypairRequest) (response *models.GetKeypairResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListKeypairRequest{
		Spec: spec,
	}
	var result *models.ListKeypairResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListKeypair(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Keypairs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetKeypairResponse{
		Keypair: result.Keypairs[0],
	}
	return response, nil
}

//RESTListKeypair handles a List REST service Request.
func (service *ContrailService) RESTListKeypair(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListKeypairRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListKeypair(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListKeypair handles a List service Request.
func (service *ContrailService) ListKeypair(
	ctx context.Context,
	request *models.ListKeypairRequest) (response *models.ListKeypairResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListKeypair(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
