package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateNetworkPolicy handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreateNetworkPolicy(c echo.Context) error {
	requestData := &models.CreateNetworkPolicyRequest{}
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
// nolint
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
		if model.DisplayName != "" {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
		} else {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.UUID}
		}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()

	return service.Next().CreateNetworkPolicy(ctx, request)
}

//RESTUpdateNetworkPolicy handles a REST Update request.
// nolint
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
// nolint
func (service *ContrailService) UpdateNetworkPolicy(
	ctx context.Context,
	request *models.UpdateNetworkPolicyRequest) (*models.UpdateNetworkPolicyResponse, error) {
	model := request.NetworkPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateNetworkPolicy(ctx, request)
}

//RESTDeleteNetworkPolicy delete a resource using REST service.
// nolint
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

//RESTGetNetworkPolicy a REST Get request.
// nolint
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

//RESTListNetworkPolicy handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListNetworkPolicy(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
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
