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

//RESTCreateFloatingIP handle a Create REST service.
func (service *ContrailService) RESTCreateFloatingIP(c echo.Context) error {
    requestData := &models.FloatingIPCreateRequest{
        FloatingIP: models.MakeFloatingIP(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "floating_ip",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateFloatingIP(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateFloatingIP handle a Create API
func (service *ContrailService) CreateFloatingIP(
    ctx context.Context, 
    request *models.FloatingIPCreateRequest) (*models.FloatingIPCreateResponse, error) {
    model := request.FloatingIP
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
            return db.CreateFloatingIP(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "floating_ip",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.FloatingIPCreateResponse{
        FloatingIP: request.FloatingIP,
    }, nil
}

//RESTUpdateFloatingIP handles a REST Update request.
func (service *ContrailService) RESTUpdateFloatingIP(c echo.Context) error {
    id := c.Param("id")
    request := &models.FloatingIPUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "floating_ip",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateFloatingIP(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateFloatingIP handles a Update request.
func (service *ContrailService) UpdateFloatingIP(ctx context.Context, request *models.FloatingIPUpdateRequest) (*models.FloatingIPUpdateResponse, error) {
    id = request.ID
    model = request.FloatingIP
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
            return db.UpdateFloatingIP(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "floating_ip",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.FloatingIP.UpdateResponse{
        FloatingIP: model,
    }, nil
}

//RESTDeleteFloatingIP delete a resource using REST service.
func (service *ContrailService) RESTDeleteFloatingIP(c echo.Context) error {
    id := c.Param("id")
    request := &models.FloatingIPDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteFloatingIP(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteFloatingIP delete a resource.
func (service *ContrailService) DeleteFloatingIP(ctx context.Context, request *models.FloatingIPDeleteRequest) (*models.FloatingIPDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteFloatingIP(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.FloatingIPDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowFloatingIP a REST Show request.
func (service *ContrailService) RESTShowFloatingIP(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.FloatingIP
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListFloatingIP(tx, &common.ListSpec{
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
        "floating_ip": result,
    })
}

//RESTListFloatingIP handles a List REST service Request.
func (service *ContrailService) RESTListFloatingIP(c echo.Context) (error) {
    var result []*models.FloatingIP
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListFloatingIP(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "floating-ips": result,
    })
}