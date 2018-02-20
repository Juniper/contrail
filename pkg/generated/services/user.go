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

//RESTUserUpdateRequest for update request for REST.
type RESTUserUpdateRequest struct {
	Data map[string]interface{} `json:"user"`
}

//RESTCreateUser handle a Create REST service.
func (service *ContrailService) RESTCreateUser(c echo.Context) error {
	requestData := &models.CreateUserRequest{
		User: models.MakeUser(),
	}
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
			return db.CreateUser(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "user",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateUserResponse{
		User: request.User,
	}, nil
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
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateUser(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "user",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateUserResponse{
		User: model,
	}, nil
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

//DeleteUser delete a resource.
func (service *ContrailService) DeleteUser(ctx context.Context, request *models.DeleteUserRequest) (*models.DeleteUserResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteUser(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteUserResponse{
		ID: request.ID,
	}, nil
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

//GetUser a Get request.
func (service *ContrailService) GetUser(ctx context.Context, request *models.GetUserRequest) (response *models.GetUserResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListUserRequest{
		Spec: spec,
	}
	var result *models.ListUserResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListUser(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Users) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetUserResponse{
		User: result.Users[0],
	}
	return response, nil
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

//ListUser handles a List service Request.
func (service *ContrailService) ListUser(
	ctx context.Context,
	request *models.ListUserRequest) (response *models.ListUserResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListUser(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
