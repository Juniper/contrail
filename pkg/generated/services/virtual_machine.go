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

//RESTCreateVirtualMachine handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualMachine(c echo.Context) error {
    requestData := &models.VirtualMachineCreateRequest{
        VirtualMachine: models.MakeVirtualMachine(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_machine",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateVirtualMachine(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateVirtualMachine handle a Create API
func (service *ContrailService) CreateVirtualMachine(
    ctx context.Context, 
    request *models.VirtualMachineCreateRequest) (*models.VirtualMachineCreateResponse, error) {
    model := request.VirtualMachine
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
            return db.CreateVirtualMachine(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_machine",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.VirtualMachineCreateResponse{
        VirtualMachine: request.VirtualMachine,
    }, nil
}

//RESTUpdateVirtualMachine handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualMachine(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualMachineUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "virtual_machine",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateVirtualMachine(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateVirtualMachine handles a Update request.
func (service *ContrailService) UpdateVirtualMachine(ctx context.Context, request *models.VirtualMachineUpdateRequest) (*models.VirtualMachineUpdateResponse, error) {
    id = request.ID
    model = request.VirtualMachine
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
            return db.UpdateVirtualMachine(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_machine",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.VirtualMachine.UpdateResponse{
        VirtualMachine: model,
    }, nil
}

//RESTDeleteVirtualMachine delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualMachine(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualMachineDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteVirtualMachine(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualMachine delete a resource.
func (service *ContrailService) DeleteVirtualMachine(ctx context.Context, request *models.VirtualMachineDeleteRequest) (*models.VirtualMachineDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteVirtualMachine(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.VirtualMachineDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowVirtualMachine a REST Show request.
func (service *ContrailService) RESTShowVirtualMachine(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.VirtualMachine
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualMachine(tx, &common.ListSpec{
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
        "virtual_machine": result,
    })
}

//RESTListVirtualMachine handles a List REST service Request.
func (service *ContrailService) RESTListVirtualMachine(c echo.Context) (error) {
    var result []*models.VirtualMachine
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualMachine(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "virtual-machines": result,
    })
}