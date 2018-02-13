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

//RESTServiceConnectionModuleUpdateRequest for update request for REST.
type RESTServiceConnectionModuleUpdateRequest struct {
    Data map[string]interface{} `json:"service-connection-module"`
}

//RESTCreateServiceConnectionModule handle a Create REST service.
func (service *ContrailService) RESTCreateServiceConnectionModule(c echo.Context) error {
    requestData := &models.ServiceConnectionModuleCreateRequest{
        ServiceConnectionModule: models.MakeServiceConnectionModule(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_connection_module",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateServiceConnectionModule(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateServiceConnectionModule handle a Create API
func (service *ContrailService) CreateServiceConnectionModule(
    ctx context.Context, 
    request *models.ServiceConnectionModuleCreateRequest) (*models.ServiceConnectionModuleCreateResponse, error) {
    model := request.ServiceConnectionModule
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
            return db.CreateServiceConnectionModule(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_connection_module",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ServiceConnectionModuleCreateResponse{
        ServiceConnectionModule: request.ServiceConnectionModule,
    }, nil
}

//RESTUpdateServiceConnectionModule handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceConnectionModule(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceConnectionModuleUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_connection_module",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateServiceConnectionModule(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceConnectionModule handles a Update request.
func (service *ContrailService) UpdateServiceConnectionModule(ctx context.Context, request *models.ServiceConnectionModuleUpdateRequest) (*models.ServiceConnectionModuleUpdateResponse, error) {
    id = request.ID
    model = request.ServiceConnectionModule
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
            return db.UpdateServiceConnectionModule(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_connection_module",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ServiceConnectionModule.UpdateResponse{
        ServiceConnectionModule: model,
    }, nil
}

//RESTDeleteServiceConnectionModule delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceConnectionModule(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceConnectionModuleDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteServiceConnectionModule(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceConnectionModule delete a resource.
func (service *ContrailService) DeleteServiceConnectionModule(ctx context.Context, request *models.ServiceConnectionModuleDeleteRequest) (*models.ServiceConnectionModuleDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceConnectionModule(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ServiceConnectionModuleDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowServiceConnectionModule a REST Show request.
func (service *ContrailService) RESTShowServiceConnectionModule(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ServiceConnectionModule
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceConnectionModule(tx, &common.ListSpec{
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
        "service_connection_module": result,
    })
}

//RESTListServiceConnectionModule handles a List REST service Request.
func (service *ContrailService) RESTListServiceConnectionModule(c echo.Context) (error) {
    var result []*models.ServiceConnectionModule
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceConnectionModule(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "service-connection-modules": result,
    })
}