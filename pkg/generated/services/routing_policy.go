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

//RESTCreateRoutingPolicy handle a Create REST service.
func (service *ContrailService) RESTCreateRoutingPolicy(c echo.Context) error {
	requestData := &models.CreateRoutingPolicyRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_policy",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateRoutingPolicy(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateRoutingPolicy handle a Create API
func (service *ContrailService) CreateRoutingPolicy(
	ctx context.Context,
	request *models.CreateRoutingPolicyRequest) (*models.CreateRoutingPolicyResponse, error) {
	model := request.RoutingPolicy
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
			return db.CreateRoutingPolicy(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_policy",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateRoutingPolicyResponse{
		RoutingPolicy: request.RoutingPolicy,
	}, nil
}

//RESTUpdateRoutingPolicy handles a REST Update request.
func (service *ContrailService) RESTUpdateRoutingPolicy(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateRoutingPolicyRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_policy",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateRoutingPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateRoutingPolicy handles a Update request.
func (service *ContrailService) UpdateRoutingPolicy(
	ctx context.Context,
	request *models.UpdateRoutingPolicyRequest) (*models.UpdateRoutingPolicyResponse, error) {
	model := request.RoutingPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateRoutingPolicy(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "routing_policy",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateRoutingPolicyResponse{
		RoutingPolicy: model,
	}, nil
}

//RESTDeleteRoutingPolicy delete a resource using REST service.
func (service *ContrailService) RESTDeleteRoutingPolicy(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteRoutingPolicyRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteRoutingPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteRoutingPolicy delete a resource.
func (service *ContrailService) DeleteRoutingPolicy(ctx context.Context, request *models.DeleteRoutingPolicyRequest) (*models.DeleteRoutingPolicyResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteRoutingPolicy(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteRoutingPolicyResponse{
		ID: request.ID,
	}, nil
}

//RESTGetRoutingPolicy a REST Get request.
func (service *ContrailService) RESTGetRoutingPolicy(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetRoutingPolicyRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetRoutingPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetRoutingPolicy a Get request.
func (service *ContrailService) GetRoutingPolicy(ctx context.Context, request *models.GetRoutingPolicyRequest) (response *models.GetRoutingPolicyResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListRoutingPolicyRequest{
		Spec: spec,
	}
	var result *models.ListRoutingPolicyResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListRoutingPolicy(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.RoutingPolicys) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetRoutingPolicyResponse{
		RoutingPolicy: result.RoutingPolicys[0],
	}
	return response, nil
}

//RESTListRoutingPolicy handles a List REST service Request.
func (service *ContrailService) RESTListRoutingPolicy(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListRoutingPolicyRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListRoutingPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListRoutingPolicy handles a List service Request.
func (service *ContrailService) ListRoutingPolicy(
	ctx context.Context,
	request *models.ListRoutingPolicyRequest) (response *models.ListRoutingPolicyResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListRoutingPolicy(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
