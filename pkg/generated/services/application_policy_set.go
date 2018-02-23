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

//RESTCreateApplicationPolicySet handle a Create REST service.
func (service *ContrailService) RESTCreateApplicationPolicySet(c echo.Context) error {
	requestData := &models.CreateApplicationPolicySetRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "application_policy_set",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateApplicationPolicySet(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateApplicationPolicySet handle a Create API
func (service *ContrailService) CreateApplicationPolicySet(
	ctx context.Context,
	request *models.CreateApplicationPolicySetRequest) (*models.CreateApplicationPolicySetResponse, error) {
	model := request.ApplicationPolicySet
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
			return db.CreateApplicationPolicySet(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "application_policy_set",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateApplicationPolicySetResponse{
		ApplicationPolicySet: request.ApplicationPolicySet,
	}, nil
}

//RESTUpdateApplicationPolicySet handles a REST Update request.
func (service *ContrailService) RESTUpdateApplicationPolicySet(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateApplicationPolicySetRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "application_policy_set",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateApplicationPolicySet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateApplicationPolicySet handles a Update request.
func (service *ContrailService) UpdateApplicationPolicySet(
	ctx context.Context,
	request *models.UpdateApplicationPolicySetRequest) (*models.UpdateApplicationPolicySetResponse, error) {
	model := request.ApplicationPolicySet
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateApplicationPolicySet(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "application_policy_set",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateApplicationPolicySetResponse{
		ApplicationPolicySet: model,
	}, nil
}

//RESTDeleteApplicationPolicySet delete a resource using REST service.
func (service *ContrailService) RESTDeleteApplicationPolicySet(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteApplicationPolicySetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteApplicationPolicySet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteApplicationPolicySet delete a resource.
func (service *ContrailService) DeleteApplicationPolicySet(ctx context.Context, request *models.DeleteApplicationPolicySetRequest) (*models.DeleteApplicationPolicySetResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteApplicationPolicySet(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteApplicationPolicySetResponse{
		ID: request.ID,
	}, nil
}

//RESTGetApplicationPolicySet a REST Get request.
func (service *ContrailService) RESTGetApplicationPolicySet(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetApplicationPolicySetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetApplicationPolicySet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetApplicationPolicySet a Get request.
func (service *ContrailService) GetApplicationPolicySet(ctx context.Context, request *models.GetApplicationPolicySetRequest) (response *models.GetApplicationPolicySetResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListApplicationPolicySetRequest{
		Spec: spec,
	}
	var result *models.ListApplicationPolicySetResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListApplicationPolicySet(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ApplicationPolicySets) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetApplicationPolicySetResponse{
		ApplicationPolicySet: result.ApplicationPolicySets[0],
	}
	return response, nil
}

//RESTListApplicationPolicySet handles a List REST service Request.
func (service *ContrailService) RESTListApplicationPolicySet(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListApplicationPolicySetRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListApplicationPolicySet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListApplicationPolicySet handles a List service Request.
func (service *ContrailService) ListApplicationPolicySet(
	ctx context.Context,
	request *models.ListApplicationPolicySetRequest) (response *models.ListApplicationPolicySetResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListApplicationPolicySet(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
