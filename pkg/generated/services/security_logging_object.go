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
    requestData := &models.CreateSecurityLoggingObjectRequest{
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
    request *models.CreateSecurityLoggingObjectRequest) (*models.CreateSecurityLoggingObjectResponse, error) {
    model := request.SecurityLoggingObject
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
        func (tx *sql.Tx) error {
            return db.CreateSecurityLoggingObject(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "security_logging_object",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateSecurityLoggingObjectResponse{
        SecurityLoggingObject: request.SecurityLoggingObject,
    }, nil
}

//RESTUpdateSecurityLoggingObject handles a REST Update request.
func (service *ContrailService) RESTUpdateSecurityLoggingObject(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateSecurityLoggingObjectRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "security_logging_object",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateSecurityLoggingObject(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateSecurityLoggingObject handles a Update request.
func (service *ContrailService) UpdateSecurityLoggingObject(
    ctx context.Context, 
    request *models.UpdateSecurityLoggingObjectRequest) (*models.UpdateSecurityLoggingObjectResponse, error) {
    model := request.SecurityLoggingObject
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateSecurityLoggingObject(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "security_logging_object",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateSecurityLoggingObjectResponse{
        SecurityLoggingObject: model,
    }, nil
}

//RESTDeleteSecurityLoggingObject delete a resource using REST service.
func (service *ContrailService) RESTDeleteSecurityLoggingObject(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteSecurityLoggingObjectRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteSecurityLoggingObject(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteSecurityLoggingObject delete a resource.
func (service *ContrailService) DeleteSecurityLoggingObject(ctx context.Context, request *models.DeleteSecurityLoggingObjectRequest) (*models.DeleteSecurityLoggingObjectResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteSecurityLoggingObject(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteSecurityLoggingObjectResponse{
        ID: request.ID,
    }, nil
}

//RESTGetSecurityLoggingObject a REST Get request.
func (service *ContrailService) RESTGetSecurityLoggingObject(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetSecurityLoggingObjectRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetSecurityLoggingObject(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetSecurityLoggingObject a Get request.
func (service *ContrailService) GetSecurityLoggingObject(ctx context.Context, request *models.GetSecurityLoggingObjectRequest) (response *models.GetSecurityLoggingObjectResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListSecurityLoggingObjectRequest{
        Spec: spec,
    }
    var result *models.ListSecurityLoggingObjectResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListSecurityLoggingObject(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.SecurityLoggingObjects) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetSecurityLoggingObjectResponse{
       SecurityLoggingObject: result.SecurityLoggingObjects[0],
    }
    return response, nil
}

//RESTListSecurityLoggingObject handles a List REST service Request.
func (service *ContrailService) RESTListSecurityLoggingObject(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListSecurityLoggingObjectRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListSecurityLoggingObject(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListSecurityLoggingObject handles a List service Request.
func (service *ContrailService) ListSecurityLoggingObject(
    ctx context.Context, 
    request *models.ListSecurityLoggingObjectRequest) (response *models.ListSecurityLoggingObjectResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListSecurityLoggingObject(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}