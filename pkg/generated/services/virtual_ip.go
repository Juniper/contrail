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

//RESTCreateVirtualIP handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualIP(c echo.Context) error {
    requestData := &models.VirtualIPCreateRequest{
        VirtualIP: models.MakeVirtualIP(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_ip",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateVirtualIP(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateVirtualIP handle a Create API
func (service *ContrailService) CreateVirtualIP(
    ctx context.Context, 
    request *models.VirtualIPCreateRequest) (*models.VirtualIPCreateResponse, error) {
    model := request.VirtualIP
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
            return db.CreateVirtualIP(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_ip",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.VirtualIPCreateResponse{
        VirtualIP: request.VirtualIP,
    }, nil
}

//RESTUpdateVirtualIP handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualIP(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualIPUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "virtual_ip",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateVirtualIP(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateVirtualIP handles a Update request.
func (service *ContrailService) UpdateVirtualIP(ctx context.Context, request *models.VirtualIPUpdateRequest) (*models.VirtualIPUpdateResponse, error) {
    id = request.ID
    model = request.VirtualIP
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
            return db.UpdateVirtualIP(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_ip",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.VirtualIP.UpdateResponse{
        VirtualIP: model,
    }, nil
}

//RESTDeleteVirtualIP delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualIP(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualIPDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteVirtualIP(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualIP delete a resource.
func (service *ContrailService) DeleteVirtualIP(ctx context.Context, request *models.VirtualIPDeleteRequest) (*models.VirtualIPDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteVirtualIP(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.VirtualIPDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowVirtualIP a REST Show request.
func (service *ContrailService) RESTShowVirtualIP(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.VirtualIP
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualIP(tx, &common.ListSpec{
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
        "virtual_ip": result,
    })
}

//RESTListVirtualIP handles a List REST service Request.
func (service *ContrailService) RESTListVirtualIP(c echo.Context) (error) {
    var result []*models.VirtualIP
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualIP(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "virtual-ips": result,
    })
}