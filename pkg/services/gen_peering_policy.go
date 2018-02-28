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

//RESTCreatePeeringPolicy handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreatePeeringPolicy(c echo.Context) error {
	requestData := &models.CreatePeeringPolicyRequest{}
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
// nolint
func (service *ContrailService) CreatePeeringPolicy(
	ctx context.Context,
	request *models.CreatePeeringPolicyRequest) (*models.CreatePeeringPolicyResponse, error) {
	model := request.PeeringPolicy
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

	return service.Next().CreatePeeringPolicy(ctx, request)
}

//RESTUpdatePeeringPolicy handles a REST Update request.
// nolint
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
// nolint
func (service *ContrailService) UpdatePeeringPolicy(
	ctx context.Context,
	request *models.UpdatePeeringPolicyRequest) (*models.UpdatePeeringPolicyResponse, error) {
	model := request.PeeringPolicy
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdatePeeringPolicy(ctx, request)
}

//RESTDeletePeeringPolicy delete a resource using REST service.
// nolint
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

//RESTGetPeeringPolicy a REST Get request.
// nolint
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

//RESTListPeeringPolicy handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListPeeringPolicy(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
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
