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

//RESTNetworkPolicyUpdateRequest for update request for REST.
type RESTNetworkPolicyUpdateRequest struct {
	Data map[string]interface{} `json:"network-policy"`
}

//RESTCreateNetworkPolicy handle a Create REST service.
func (service *ContrailService) RESTCreateNetworkPolicy(c echo.Context) error {
	requestData := &models.CreateNetworkPolicyRequest{
		NetworkPolicy: models.MakeNetworkPolicy(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_policy",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateNetworkPolicy(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateNetworkPolicy handle a Create API
func (service *ContrailService) CreateNetworkPolicy(
	ctx context.Context,
	request *models.CreateNetworkPolicyRequest) (*models.CreateNetworkPolicyResponse, error) {
	model := request.NetworkPolicy
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
			return db.CreateNetworkPolicy(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_policy",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateNetworkPolicyResponse{
		NetworkPolicy: request.NetworkPolicy,
	}, nil
}

//RESTUpdateNetworkPolicy handles a REST Update request.
func (service *ContrailService) RESTUpdateNetworkPolicy(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateNetworkPolicyRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_policy",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateNetworkPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateNetworkPolicy handles a Update request.
func (service *ContrailService) UpdateNetworkPolicy(
	ctx context.Context,
	request *models.UpdateNetworkPolicyRequest) (*models.UpdateNetworkPolicyResponse, error) {
	model := request.NetworkPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateNetworkPolicy(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_policy",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateNetworkPolicyResponse{
		NetworkPolicy: model,
	}, nil
}

//RESTDeleteNetworkPolicy delete a resource using REST service.
func (service *ContrailService) RESTDeleteNetworkPolicy(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteNetworkPolicyRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteNetworkPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteNetworkPolicy delete a resource.
func (service *ContrailService) DeleteNetworkPolicy(ctx context.Context, request *models.DeleteNetworkPolicyRequest) (*models.DeleteNetworkPolicyResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteNetworkPolicy(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteNetworkPolicyResponse{
		ID: request.ID,
	}, nil
}

//RESTGetNetworkPolicy a REST Get request.
func (service *ContrailService) RESTGetNetworkPolicy(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetNetworkPolicyRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetNetworkPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetNetworkPolicy a Get request.
func (service *ContrailService) GetNetworkPolicy(ctx context.Context, request *models.GetNetworkPolicyRequest) (response *models.GetNetworkPolicyResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListNetworkPolicyRequest{
		Spec: spec,
	}
	var result *models.ListNetworkPolicyResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListNetworkPolicy(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.NetworkPolicys) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetNetworkPolicyResponse{
		NetworkPolicy: result.NetworkPolicys[0],
	}
	return response, nil
}

//RESTListNetworkPolicy handles a List REST service Request.
func (service *ContrailService) RESTListNetworkPolicy(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListNetworkPolicyRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListNetworkPolicy(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListNetworkPolicy handles a List service Request.
func (service *ContrailService) ListNetworkPolicy(
	ctx context.Context,
	request *models.ListNetworkPolicyRequest) (response *models.ListNetworkPolicyResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListNetworkPolicy(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
