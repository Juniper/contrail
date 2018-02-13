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

//RESTVirtualDNSRecordUpdateRequest for update request for REST.
type RESTVirtualDNSRecordUpdateRequest struct {
    Data map[string]interface{} `json:"virtual-DNS-record"`
}

//RESTCreateVirtualDNSRecord handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualDNSRecord(c echo.Context) error {
    requestData := &models.VirtualDNSRecordCreateRequest{
        VirtualDNSRecord: models.MakeVirtualDNSRecord(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_DNS_record",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateVirtualDNSRecord(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateVirtualDNSRecord handle a Create API
func (service *ContrailService) CreateVirtualDNSRecord(
    ctx context.Context, 
    request *models.VirtualDNSRecordCreateRequest) (*models.VirtualDNSRecordCreateResponse, error) {
    model := request.VirtualDNSRecord
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
            return db.CreateVirtualDNSRecord(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_DNS_record",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.VirtualDNSRecordCreateResponse{
        VirtualDNSRecord: request.VirtualDNSRecord,
    }, nil
}

//RESTUpdateVirtualDNSRecord handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualDNSRecord(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualDNSRecordUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "virtual_DNS_record",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateVirtualDNSRecord(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateVirtualDNSRecord handles a Update request.
func (service *ContrailService) UpdateVirtualDNSRecord(ctx context.Context, request *models.VirtualDNSRecordUpdateRequest) (*models.VirtualDNSRecordUpdateResponse, error) {
    id = request.ID
    model = request.VirtualDNSRecord
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
            return db.UpdateVirtualDNSRecord(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_DNS_record",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.VirtualDNSRecord.UpdateResponse{
        VirtualDNSRecord: model,
    }, nil
}

//RESTDeleteVirtualDNSRecord delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualDNSRecord(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualDNSRecordDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteVirtualDNSRecord(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualDNSRecord delete a resource.
func (service *ContrailService) DeleteVirtualDNSRecord(ctx context.Context, request *models.VirtualDNSRecordDeleteRequest) (*models.VirtualDNSRecordDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteVirtualDNSRecord(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.VirtualDNSRecordDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowVirtualDNSRecord a REST Show request.
func (service *ContrailService) RESTShowVirtualDNSRecord(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.VirtualDNSRecord
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualDNSRecord(tx, &common.ListSpec{
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
        "virtual_DNS_record": result,
    })
}

//RESTListVirtualDNSRecord handles a List REST service Request.
func (service *ContrailService) RESTListVirtualDNSRecord(c echo.Context) (error) {
    var result []*models.VirtualDNSRecord
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualDNSRecord(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "virtual-DNS-records": result,
    })
}