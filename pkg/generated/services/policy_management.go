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

//RESTCreatePolicyManagement handle a Create REST service.
func (service *ContrailService) RESTCreatePolicyManagement(c echo.Context) error {
	requestData := &models.CreatePolicyManagementRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "policy_management",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreatePolicyManagement(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreatePolicyManagement handle a Create API
func (service *ContrailService) CreatePolicyManagement(
	ctx context.Context,
	request *models.CreatePolicyManagementRequest) (*models.CreatePolicyManagementResponse, error) {
	model := request.PolicyManagement
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
			return db.CreatePolicyManagement(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "policy_management",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreatePolicyManagementResponse{
		PolicyManagement: request.PolicyManagement,
	}, nil
}

//RESTUpdatePolicyManagement handles a REST Update request.
func (service *ContrailService) RESTUpdatePolicyManagement(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdatePolicyManagementRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "policy_management",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdatePolicyManagement(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdatePolicyManagement handles a Update request.
func (service *ContrailService) UpdatePolicyManagement(
	ctx context.Context,
	request *models.UpdatePolicyManagementRequest) (*models.UpdatePolicyManagementResponse, error) {
	model := request.PolicyManagement
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdatePolicyManagement(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "policy_management",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdatePolicyManagementResponse{
		PolicyManagement: model,
	}, nil
}

//RESTDeletePolicyManagement delete a resource using REST service.
func (service *ContrailService) RESTDeletePolicyManagement(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeletePolicyManagementRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeletePolicyManagement(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeletePolicyManagement delete a resource.
func (service *ContrailService) DeletePolicyManagement(ctx context.Context, request *models.DeletePolicyManagementRequest) (*models.DeletePolicyManagementResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeletePolicyManagement(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeletePolicyManagementResponse{
		ID: request.ID,
	}, nil
}

//RESTGetPolicyManagement a REST Get request.
func (service *ContrailService) RESTGetPolicyManagement(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetPolicyManagementRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetPolicyManagement(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetPolicyManagement a Get request.
func (service *ContrailService) GetPolicyManagement(ctx context.Context, request *models.GetPolicyManagementRequest) (response *models.GetPolicyManagementResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListPolicyManagementRequest{
		Spec: spec,
	}
	var result *models.ListPolicyManagementResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListPolicyManagement(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.PolicyManagements) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetPolicyManagementResponse{
		PolicyManagement: result.PolicyManagements[0],
	}
	return response, nil
}

//RESTListPolicyManagement handles a List REST service Request.
func (service *ContrailService) RESTListPolicyManagement(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListPolicyManagementRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListPolicyManagement(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListPolicyManagement handles a List service Request.
func (service *ContrailService) ListPolicyManagement(
	ctx context.Context,
	request *models.ListPolicyManagementRequest) (response *models.ListPolicyManagementResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListPolicyManagement(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
