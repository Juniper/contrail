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

//RESTCreateSecurityLoggingObject handle a Create REST service.
func (service *ContrailService) RESTCreateSecurityLoggingObject(c echo.Context) error {
    requestData := &models.SecurityLoggingObjectCreateRequest{
        SecurityLoggingObject: models.MakeSecurityLoggingObject(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "security_logging_object",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateSecurityLoggingObject(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateSecurityLoggingObject handle a Create API
func (service *ContrailService) CreateSecurityLoggingObject(
    ctx context.Context, 
    request *models.SecurityLoggingObjectCreateRequest) (*models.SecurityLoggingObjectCreateResponse, error) {
    model := request.SecurityLoggingObject
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
            return db.CreateSecurityLoggingObject(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "security_logging_object",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.SecurityLoggingObjectCreateResponse{
        SecurityLoggingObject: request.SecurityLoggingObject,
    }, nil
}

//RESTUpdateSecurityLoggingObject handles a REST Update request.
func (service *ContrailService) RESTUpdateSecurityLoggingObject(c echo.Context) error {
    id := c.Param("id")
    request := &models.SecurityLoggingObjectUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "security_logging_object",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateSecurityLoggingObject(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateSecurityLoggingObject handles a Update request.
func (service *ContrailService) UpdateSecurityLoggingObject(ctx context.Context, request *models.SecurityLoggingObjectUpdateRequest) (*models.SecurityLoggingObjectUpdateResponse, error) {
    id = request.ID
    model = request.SecurityLoggingObject
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
            return db.UpdateSecurityLoggingObject(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "security_logging_object",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.SecurityLoggingObject.UpdateResponse{
        SecurityLoggingObject: model,
    }, nil
}

//RESTDeleteSecurityLoggingObject delete a resource using REST service.
func (service *ContrailService) RESTDeleteSecurityLoggingObject(c echo.Context) error {
    id := c.Param("id")
    request := &models.SecurityLoggingObjectDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteSecurityLoggingObject(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteSecurityLoggingObject delete a resource.
func (service *ContrailService) DeleteSecurityLoggingObject(ctx context.Context, request *models.SecurityLoggingObjectDeleteRequest) (*models.SecurityLoggingObjectDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteSecurityLoggingObject(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.SecurityLoggingObjectDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowSecurityLoggingObject a REST Show request.
func (service *ContrailService) RESTShowSecurityLoggingObject(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.SecurityLoggingObject
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListSecurityLoggingObject(tx, &common.ListSpec{
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
        "security_logging_object": result,
    })
}

//RESTListSecurityLoggingObject handles a List REST service Request.
func (service *ContrailService) RESTListSecurityLoggingObject(c echo.Context) (error) {
    var result []*models.SecurityLoggingObject
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListSecurityLoggingObject(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "security-logging-objects": result,
    })
}