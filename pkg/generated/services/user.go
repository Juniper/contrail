package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateUser handle a Create REST service.
func (service *ContrailService) RESTCreateUser(c echo.Context) error {
	requestData := &models.CreateUserRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "user",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateUser(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateUser handle a Create API
func (service *ContrailService) CreateUser(
	ctx context.Context,
	request *models.CreateUserRequest) (*models.CreateUserResponse, error) {
	model := request.User
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

	return service.Next().CreateUser(ctx, request)
}

//RESTUpdateUser handles a REST Update request.
func (service *ContrailService) RESTUpdateUser(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateUserRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "user",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateUser(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateUser handles a Update request.
func (service *ContrailService) UpdateUser(
	ctx context.Context,
	request *models.UpdateUserRequest) (*models.UpdateUserResponse, error) {
	model := request.User
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdateUser(ctx, request)
}

//RESTDeleteUser delete a resource using REST service.
func (service *ContrailService) RESTDeleteUser(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteUserRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteUser(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetUser a REST Get request.
func (service *ContrailService) RESTGetUser(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetUserRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetUser(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListUser handles a List REST service Request.
func (service *ContrailService) RESTListUser(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListUserRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListUser(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
