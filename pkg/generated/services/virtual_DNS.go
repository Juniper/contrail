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

//RESTCreateVirtualDNS handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualDNS(c echo.Context) error {
    requestData := &models.VirtualDNSCreateRequest{
        VirtualDNS: models.MakeVirtualDNS(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_DNS",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateVirtualDNS(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateVirtualDNS handle a Create API
func (service *ContrailService) CreateVirtualDNS(
    ctx context.Context, 
    request *models.VirtualDNSCreateRequest) (*models.VirtualDNSCreateResponse, error) {
    model := request.VirtualDNS
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
            return db.CreateVirtualDNS(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_DNS",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.VirtualDNSCreateResponse{
        VirtualDNS: request.VirtualDNS,
    }, nil
}

//RESTUpdateVirtualDNS handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualDNS(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualDNSUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "virtual_DNS",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateVirtualDNS(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateVirtualDNS handles a Update request.
func (service *ContrailService) UpdateVirtualDNS(ctx context.Context, request *models.VirtualDNSUpdateRequest) (*models.VirtualDNSUpdateResponse, error) {
    id = request.ID
    model = request.VirtualDNS
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
            return db.UpdateVirtualDNS(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_DNS",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.VirtualDNS.UpdateResponse{
        VirtualDNS: model,
    }, nil
}

//RESTDeleteVirtualDNS delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualDNS(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualDNSDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteVirtualDNS(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualDNS delete a resource.
func (service *ContrailService) DeleteVirtualDNS(ctx context.Context, request *models.VirtualDNSDeleteRequest) (*models.VirtualDNSDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteVirtualDNS(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.VirtualDNSDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowVirtualDNS a REST Show request.
func (service *ContrailService) RESTShowVirtualDNS(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.VirtualDNS
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualDNS(tx, &common.ListSpec{
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
        "virtual_DNS": result,
    })
}

//RESTListVirtualDNS handles a List REST service Request.
func (service *ContrailService) RESTListVirtualDNS(c echo.Context) (error) {
    var result []*models.VirtualDNS
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualDNS(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "virtual-DNSs": result,
    })
}