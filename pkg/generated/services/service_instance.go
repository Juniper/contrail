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

//RESTCreateServiceInstance handle a Create REST service.
func (service *ContrailService) RESTCreateServiceInstance(c echo.Context) error {
    requestData := &models.ServiceInstanceCreateRequest{
        ServiceInstance: models.MakeServiceInstance(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_instance",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateServiceInstance(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateServiceInstance handle a Create API
func (service *ContrailService) CreateServiceInstance(
    ctx context.Context, 
    request *models.ServiceInstanceCreateRequest) (*models.ServiceInstanceCreateResponse, error) {
    model := request.ServiceInstance
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
            return db.CreateServiceInstance(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_instance",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ServiceInstanceCreateResponse{
        ServiceInstance: request.ServiceInstance,
    }, nil
}

//RESTUpdateServiceInstance handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceInstance(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceInstanceUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_instance",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateServiceInstance(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceInstance handles a Update request.
func (service *ContrailService) UpdateServiceInstance(ctx context.Context, request *models.ServiceInstanceUpdateRequest) (*models.ServiceInstanceUpdateResponse, error) {
    id = request.ID
    model = request.ServiceInstance
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
            return db.UpdateServiceInstance(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_instance",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ServiceInstance.UpdateResponse{
        ServiceInstance: model,
    }, nil
}

//RESTDeleteServiceInstance delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceInstance(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceInstanceDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteServiceInstance(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceInstance delete a resource.
func (service *ContrailService) DeleteServiceInstance(ctx context.Context, request *models.ServiceInstanceDeleteRequest) (*models.ServiceInstanceDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceInstance(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ServiceInstanceDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowServiceInstance a REST Show request.
func (service *ContrailService) RESTShowServiceInstance(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ServiceInstance
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceInstance(tx, &common.ListSpec{
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
        "service_instance": result,
    })
}

//RESTListServiceInstance handles a List REST service Request.
func (service *ContrailService) RESTListServiceInstance(c echo.Context) (error) {
    var result []*models.ServiceInstance
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceInstance(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "service-instances": result,
    })
}