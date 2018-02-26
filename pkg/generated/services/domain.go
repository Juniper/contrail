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

//RESTCreateDomain handle a Create REST service.
func (service *ContrailService) RESTCreateDomain(c echo.Context) error {
	requestData := &models.CreateDomainRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "domain",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateDomain(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateDomain handle a Create API
func (service *ContrailService) CreateDomain(
	ctx context.Context,
	request *models.CreateDomainRequest) (*models.CreateDomainResponse, error) {
	model := request.Domain
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
			return db.CreateDomain(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "domain",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateDomainResponse{
		Domain: request.Domain,
	}, nil
}

//RESTUpdateDomain handles a REST Update request.
func (service *ContrailService) RESTUpdateDomain(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateDomainRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "domain",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateDomain handles a Update request.
func (service *ContrailService) UpdateDomain(
	ctx context.Context,
	request *models.UpdateDomainRequest) (*models.UpdateDomainResponse, error) {
	model := request.Domain
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateDomain(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "domain",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateDomainResponse{
		Domain: model,
	}, nil
}

//RESTDeleteDomain delete a resource using REST service.
func (service *ContrailService) RESTDeleteDomain(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteDomainRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteDomain delete a resource.
func (service *ContrailService) DeleteDomain(ctx context.Context, request *models.DeleteDomainRequest) (*models.DeleteDomainResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteDomain(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteDomainResponse{
		ID: request.ID,
	}, nil
}

//RESTGetDomain a REST Get request.
func (service *ContrailService) RESTGetDomain(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetDomainRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetDomain a Get request.
func (service *ContrailService) GetDomain(ctx context.Context, request *models.GetDomainRequest) (response *models.GetDomainResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListDomainRequest{
		Spec: spec,
	}
	var result *models.ListDomainResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListDomain(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Domains) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetDomainResponse{
		Domain: result.Domains[0],
	}
	return response, nil
}

//RESTListDomain handles a List REST service Request.
func (service *ContrailService) RESTListDomain(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListDomainRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListDomain handles a List service Request.
func (service *ContrailService) ListDomain(
	ctx context.Context,
	request *models.ListDomainRequest) (response *models.ListDomainResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListDomain(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
