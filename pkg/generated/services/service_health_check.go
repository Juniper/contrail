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

//RESTServiceHealthCheckUpdateRequest for update request for REST.
type RESTServiceHealthCheckUpdateRequest struct {
    Data map[string]interface{} `json:"service-health-check"`
}

//RESTCreateServiceHealthCheck handle a Create REST service.
func (service *ContrailService) RESTCreateServiceHealthCheck(c echo.Context) error {
    requestData := &models.ServiceHealthCheckCreateRequest{
        ServiceHealthCheck: models.MakeServiceHealthCheck(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_health_check",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateServiceHealthCheck(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateServiceHealthCheck handle a Create API
func (service *ContrailService) CreateServiceHealthCheck(
    ctx context.Context, 
    request *models.ServiceHealthCheckCreateRequest) (*models.ServiceHealthCheckCreateResponse, error) {
    model := request.ServiceHealthCheck
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
            return db.CreateServiceHealthCheck(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_health_check",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ServiceHealthCheckCreateResponse{
        ServiceHealthCheck: request.ServiceHealthCheck,
    }, nil
}

//RESTUpdateServiceHealthCheck handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceHealthCheck(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceHealthCheckUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_health_check",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateServiceHealthCheck(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceHealthCheck handles a Update request.
func (service *ContrailService) UpdateServiceHealthCheck(ctx context.Context, request *models.ServiceHealthCheckUpdateRequest) (*models.ServiceHealthCheckUpdateResponse, error) {
    id = request.ID
    model = request.ServiceHealthCheck
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
            return db.UpdateServiceHealthCheck(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_health_check",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ServiceHealthCheck.UpdateResponse{
        ServiceHealthCheck: model,
    }, nil
}

//RESTDeleteServiceHealthCheck delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceHealthCheck(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceHealthCheckDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteServiceHealthCheck(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceHealthCheck delete a resource.
func (service *ContrailService) DeleteServiceHealthCheck(ctx context.Context, request *models.ServiceHealthCheckDeleteRequest) (*models.ServiceHealthCheckDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceHealthCheck(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ServiceHealthCheckDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowServiceHealthCheck a REST Show request.
func (service *ContrailService) RESTShowServiceHealthCheck(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ServiceHealthCheck
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceHealthCheck(tx, &common.ListSpec{
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
        "service_health_check": result,
    })
}

//RESTListServiceHealthCheck handles a List REST service Request.
func (service *ContrailService) RESTListServiceHealthCheck(c echo.Context) (error) {
    var result []*models.ServiceHealthCheck
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceHealthCheck(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "service-health-checks": result,
    })
}