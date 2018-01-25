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

//RESTCreateServiceObject handle a Create REST service.
func (service *ContrailService) RESTCreateServiceObject(c echo.Context) error {
    requestData := &models.ServiceObjectCreateRequest{
        ServiceObject: models.MakeServiceObject(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_object",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateServiceObject(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateServiceObject handle a Create API
func (service *ContrailService) CreateServiceObject(
    ctx context.Context, 
    request *models.ServiceObjectCreateRequest) (*models.ServiceObjectCreateResponse, error) {
    model := request.ServiceObject
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
            return db.CreateServiceObject(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_object",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ServiceObjectCreateResponse{
        ServiceObject: request.ServiceObject,
    }, nil
}

//RESTUpdateServiceObject handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceObject(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceObjectUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_object",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateServiceObject(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceObject handles a Update request.
func (service *ContrailService) UpdateServiceObject(ctx context.Context, request *models.ServiceObjectUpdateRequest) (*models.ServiceObjectUpdateResponse, error) {
    id = request.ID
    model = request.ServiceObject
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
            return db.UpdateServiceObject(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_object",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ServiceObject.UpdateResponse{
        ServiceObject: model,
    }, nil
}

//RESTDeleteServiceObject delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceObject(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceObjectDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteServiceObject(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceObject delete a resource.
func (service *ContrailService) DeleteServiceObject(ctx context.Context, request *models.ServiceObjectDeleteRequest) (*models.ServiceObjectDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceObject(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ServiceObjectDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowServiceObject a REST Show request.
func (service *ContrailService) RESTShowServiceObject(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ServiceObject
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceObject(tx, &common.ListSpec{
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
        "service_object": result,
    })
}

//RESTListServiceObject handles a List REST service Request.
func (service *ContrailService) RESTListServiceObject(c echo.Context) (error) {
    var result []*models.ServiceObject
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceObject(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "service-objects": result,
    })
}