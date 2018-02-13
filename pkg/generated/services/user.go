package services 

import (
    "context"
    "net/http"
    "database/sql"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/Juniper/contrail/pkg/generated/db"
    "github.com/satori/go.uuid"
    "github.com/labstack/echo"
    "github.com/Juniper/contrail/pkg/common"

	log "github.com/sirupsen/logrus"
)

//RESTUserUpdateRequest for update request for REST.
type RESTUserUpdateRequest struct {
    Data map[string]interface{} `json:"user"`
}

//RESTCreateUser handle a Create REST service.
func (service *ContrailService) RESTCreateUser(c echo.Context) error {
    requestData := &models.UserCreateRequest{
        User: models.MakeUser(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
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
    request *models.UserCreateRequest) (*models.UserCreateResponse, error) {
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
    model.Perms2.Owner = auth.ProjectID()
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.CreateUser(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "user",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.UserCreateResponse{
        User: request.User,
    }, nil
}

//RESTUpdateUser handles a REST Update request.
func (service *ContrailService) RESTUpdateUser(c echo.Context) error {
    id := c.Param("id")
    request := &models.UserUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "user",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateUser(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateUser handles a Update request.
func (service *ContrailService) UpdateUser(ctx context.Context, request *models.UserUpdateRequest) (*models.UserUpdateResponse, error) {
    id = request.ID
    model = request.User
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    auth := common.GetAuthCTX(ctx)
    ok := common.SetValueByPath(model, "Perms2.Owner", ".", auth.ProjectID())
    if !ok {
        return nil, common.ErrorBadRequest("Invalid JSON format")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateUser(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "user",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.User.UpdateResponse{
        User: model,
    }, nil
}

//RESTDeleteUser delete a resource using REST service.
func (service *ContrailService) RESTDeleteUser(c echo.Context) error {
    id := c.Param("id")
    request := &models.UserDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteUser(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteUser delete a resource.
func (service *ContrailService) DeleteUser(ctx context.Context, request *models.UserDeleteRequest) (*models.UserDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteUser(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.UserDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowUser a REST Show request.
func (service *ContrailService) RESTShowUser(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.User
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListUser(tx, &common.ListSpec{
                Limit: 1,
                Auth: auth,
                Filter: common.Filter{
                    "uuid": []string{id},
                },
            })
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "user": result,
    })
}

//RESTListUser handles a List REST service Request.
func (service *ContrailService) RESTListUser(c echo.Context) (error) {
    var result []*models.User
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListUser(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "users": result,
    })
}