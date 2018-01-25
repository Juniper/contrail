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

//RESTCreateServiceGroup handle a Create REST service.
func (service *ContrailService) RESTCreateServiceGroup(c echo.Context) error {
    requestData := &models.ServiceGroupCreateRequest{
        ServiceGroup: models.MakeServiceGroup(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_group",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateServiceGroup(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateServiceGroup handle a Create API
func (service *ContrailService) CreateServiceGroup(
    ctx context.Context, 
    request *models.ServiceGroupCreateRequest) (*models.ServiceGroupCreateResponse, error) {
    model := request.ServiceGroup
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
            return db.CreateServiceGroup(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_group",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ServiceGroupCreateResponse{
        ServiceGroup: request.ServiceGroup,
    }, nil
}

//RESTUpdateServiceGroup handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceGroup(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceGroupUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_group",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateServiceGroup(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceGroup handles a Update request.
func (service *ContrailService) UpdateServiceGroup(ctx context.Context, request *models.ServiceGroupUpdateRequest) (*models.ServiceGroupUpdateResponse, error) {
    id = request.ID
    model = request.ServiceGroup
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
            return db.UpdateServiceGroup(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_group",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ServiceGroup.UpdateResponse{
        ServiceGroup: model,
    }, nil
}

//RESTDeleteServiceGroup delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceGroup(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceGroupDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteServiceGroup(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceGroup delete a resource.
func (service *ContrailService) DeleteServiceGroup(ctx context.Context, request *models.ServiceGroupDeleteRequest) (*models.ServiceGroupDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceGroup(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ServiceGroupDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowServiceGroup a REST Show request.
func (service *ContrailService) RESTShowServiceGroup(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ServiceGroup
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceGroup(tx, &common.ListSpec{
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
        "service_group": result,
    })
}

//RESTListServiceGroup handles a List REST service Request.
func (service *ContrailService) RESTListServiceGroup(c echo.Context) (error) {
    var result []*models.ServiceGroup
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceGroup(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "service-groups": result,
    })
}