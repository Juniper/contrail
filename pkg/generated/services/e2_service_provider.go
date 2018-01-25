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

//RESTCreateE2ServiceProvider handle a Create REST service.
func (service *ContrailService) RESTCreateE2ServiceProvider(c echo.Context) error {
    requestData := &models.E2ServiceProviderCreateRequest{
        E2ServiceProvider: models.MakeE2ServiceProvider(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "e2_service_provider",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateE2ServiceProvider(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateE2ServiceProvider handle a Create API
func (service *ContrailService) CreateE2ServiceProvider(
    ctx context.Context, 
    request *models.E2ServiceProviderCreateRequest) (*models.E2ServiceProviderCreateResponse, error) {
    model := request.E2ServiceProvider
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
            return db.CreateE2ServiceProvider(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "e2_service_provider",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.E2ServiceProviderCreateResponse{
        E2ServiceProvider: request.E2ServiceProvider,
    }, nil
}

//RESTUpdateE2ServiceProvider handles a REST Update request.
func (service *ContrailService) RESTUpdateE2ServiceProvider(c echo.Context) error {
    id := c.Param("id")
    request := &models.E2ServiceProviderUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "e2_service_provider",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateE2ServiceProvider(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateE2ServiceProvider handles a Update request.
func (service *ContrailService) UpdateE2ServiceProvider(ctx context.Context, request *models.E2ServiceProviderUpdateRequest) (*models.E2ServiceProviderUpdateResponse, error) {
    id = request.ID
    model = request.E2ServiceProvider
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
            return db.UpdateE2ServiceProvider(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "e2_service_provider",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.E2ServiceProvider.UpdateResponse{
        E2ServiceProvider: model,
    }, nil
}

//RESTDeleteE2ServiceProvider delete a resource using REST service.
func (service *ContrailService) RESTDeleteE2ServiceProvider(c echo.Context) error {
    id := c.Param("id")
    request := &models.E2ServiceProviderDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteE2ServiceProvider(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteE2ServiceProvider delete a resource.
func (service *ContrailService) DeleteE2ServiceProvider(ctx context.Context, request *models.E2ServiceProviderDeleteRequest) (*models.E2ServiceProviderDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteE2ServiceProvider(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.E2ServiceProviderDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowE2ServiceProvider a REST Show request.
func (service *ContrailService) RESTShowE2ServiceProvider(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.E2ServiceProvider
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListE2ServiceProvider(tx, &common.ListSpec{
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
        "e2_service_provider": result,
    })
}

//RESTListE2ServiceProvider handles a List REST service Request.
func (service *ContrailService) RESTListE2ServiceProvider(c echo.Context) (error) {
    var result []*models.E2ServiceProvider
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListE2ServiceProvider(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "e2-service-providers": result,
    })
}