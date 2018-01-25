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

//RESTCreateServiceEndpoint handle a Create REST service.
func (service *ContrailService) RESTCreateServiceEndpoint(c echo.Context) error {
    requestData := &models.ServiceEndpointCreateRequest{
        ServiceEndpoint: models.MakeServiceEndpoint(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_endpoint",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateServiceEndpoint(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateServiceEndpoint handle a Create API
func (service *ContrailService) CreateServiceEndpoint(
    ctx context.Context, 
    request *models.ServiceEndpointCreateRequest) (*models.ServiceEndpointCreateResponse, error) {
    model := request.ServiceEndpoint
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
            return db.CreateServiceEndpoint(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_endpoint",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ServiceEndpointCreateResponse{
        ServiceEndpoint: request.ServiceEndpoint,
    }, nil
}

//RESTUpdateServiceEndpoint handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceEndpoint(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceEndpointUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_endpoint",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateServiceEndpoint(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceEndpoint handles a Update request.
func (service *ContrailService) UpdateServiceEndpoint(ctx context.Context, request *models.ServiceEndpointUpdateRequest) (*models.ServiceEndpointUpdateResponse, error) {
    id = request.ID
    model = request.ServiceEndpoint
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
            return db.UpdateServiceEndpoint(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_endpoint",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ServiceEndpoint.UpdateResponse{
        ServiceEndpoint: model,
    }, nil
}

//RESTDeleteServiceEndpoint delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceEndpoint(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceEndpointDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteServiceEndpoint(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceEndpoint delete a resource.
func (service *ContrailService) DeleteServiceEndpoint(ctx context.Context, request *models.ServiceEndpointDeleteRequest) (*models.ServiceEndpointDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceEndpoint(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ServiceEndpointDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowServiceEndpoint a REST Show request.
func (service *ContrailService) RESTShowServiceEndpoint(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ServiceEndpoint
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceEndpoint(tx, &common.ListSpec{
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
        "service_endpoint": result,
    })
}

//RESTListServiceEndpoint handles a List REST service Request.
func (service *ContrailService) RESTListServiceEndpoint(c echo.Context) (error) {
    var result []*models.ServiceEndpoint
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceEndpoint(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "service-endpoints": result,
    })
}