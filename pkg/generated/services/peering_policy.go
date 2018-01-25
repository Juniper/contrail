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

//RESTPeeringPolicyUpdateRequest for update request for REST.
type RESTPeeringPolicyUpdateRequest struct {
	Data map[string]interface{} `json:"peering-policy"`
}

//RESTCreatePeeringPolicy handle a Create REST service.
func (service *ContrailService) RESTCreatePeeringPolicy(c echo.Context) error {
	requestData := &models.CreatePeeringPolicyRequest{
		PeeringPolicy: models.MakePeeringPolicy(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "peering_policy",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreatePeeringPolicy(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreatePeeringPolicy handle a Create API
func (service *ContrailService) CreatePeeringPolicy(
	ctx context.Context,
	request *models.CreatePeeringPolicyRequest) (*models.CreatePeeringPolicyResponse, error) {
	model := request.PeeringPolicy
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}
	if model.FQName == nil {
		return nil, common.ErrorBadRequest("Missing fq_name")
	}

	auth := common.GetAuthCTX(ctx)
	if auth == nil {
		return nil, common.ErrorUnauthenticated
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreatePeeringPolicy(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "peering_policy",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreatePeeringPolicyResponse{
		PeeringPolicy: request.PeeringPolicy,
	}, nil
}

//RESTUpdatePeeringPolicy handles a REST Update request.
func (service *ContrailService) RESTUpdatePeeringPolicy(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdatePeeringPolicyRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "peering_policy",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdatePeeringPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdatePeeringPolicy handles a Update request.
func (service *ContrailService) UpdatePeeringPolicy(
	ctx context.Context,
	request *models.UpdatePeeringPolicyRequest) (*models.UpdatePeeringPolicyResponse, error) {
	model := request.PeeringPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdatePeeringPolicy(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "peering_policy",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdatePeeringPolicyResponse{
		PeeringPolicy: model,
	}, nil
}

//RESTDeletePeeringPolicy delete a resource using REST service.
func (service *ContrailService) RESTDeletePeeringPolicy(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeletePeeringPolicyRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeletePeeringPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeletePeeringPolicy delete a resource.
func (service *ContrailService) DeletePeeringPolicy(ctx context.Context, request *models.DeletePeeringPolicyRequest) (*models.DeletePeeringPolicyResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeletePeeringPolicy(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeletePeeringPolicyResponse{
		ID: request.ID,
	}, nil
}

//RESTGetPeeringPolicy a REST Get request.
func (service *ContrailService) RESTGetPeeringPolicy(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetPeeringPolicyRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetPeeringPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetPeeringPolicy a Get request.
func (service *ContrailService) GetPeeringPolicy(ctx context.Context, request *models.GetPeeringPolicyRequest) (response *models.GetPeeringPolicyResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListPeeringPolicyRequest{
		Spec: spec,
	}
	var result *models.ListPeeringPolicyResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListPeeringPolicy(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.PeeringPolicys) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetPeeringPolicyResponse{
		PeeringPolicy: result.PeeringPolicys[0],
	}
	return response, nil
}

//RESTListPeeringPolicy handles a List REST service Request.
func (service *ContrailService) RESTListPeeringPolicy(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListPeeringPolicyRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListPeeringPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListPeeringPolicy handles a List service Request.
func (service *ContrailService) ListPeeringPolicy(
	ctx context.Context,
	request *models.ListPeeringPolicyRequest) (response *models.ListPeeringPolicyResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListPeeringPolicy(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
