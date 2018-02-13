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

//RESTVirtualMachineInterfaceUpdateRequest for update request for REST.
type RESTVirtualMachineInterfaceUpdateRequest struct {
    Data map[string]interface{} `json:"virtual-machine-interface"`
}

//RESTCreateVirtualMachineInterface handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualMachineInterface(c echo.Context) error {
    requestData := &models.VirtualMachineInterfaceCreateRequest{
        VirtualMachineInterface: models.MakeVirtualMachineInterface(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_machine_interface",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateVirtualMachineInterface(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateVirtualMachineInterface handle a Create API
func (service *ContrailService) CreateVirtualMachineInterface(
    ctx context.Context, 
    request *models.VirtualMachineInterfaceCreateRequest) (*models.VirtualMachineInterfaceCreateResponse, error) {
    model := request.VirtualMachineInterface
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
            return db.CreateVirtualMachineInterface(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_machine_interface",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.VirtualMachineInterfaceCreateResponse{
        VirtualMachineInterface: request.VirtualMachineInterface,
    }, nil
}

//RESTUpdateVirtualMachineInterface handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualMachineInterface(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualMachineInterfaceUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "virtual_machine_interface",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateVirtualMachineInterface(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateVirtualMachineInterface handles a Update request.
func (service *ContrailService) UpdateVirtualMachineInterface(ctx context.Context, request *models.VirtualMachineInterfaceUpdateRequest) (*models.VirtualMachineInterfaceUpdateResponse, error) {
    id = request.ID
    model = request.VirtualMachineInterface
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
            return db.UpdateVirtualMachineInterface(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_machine_interface",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.VirtualMachineInterface.UpdateResponse{
        VirtualMachineInterface: model,
    }, nil
}

//RESTDeleteVirtualMachineInterface delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualMachineInterface(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualMachineInterfaceDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteVirtualMachineInterface(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualMachineInterface delete a resource.
func (service *ContrailService) DeleteVirtualMachineInterface(ctx context.Context, request *models.VirtualMachineInterfaceDeleteRequest) (*models.VirtualMachineInterfaceDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteVirtualMachineInterface(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.VirtualMachineInterfaceDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowVirtualMachineInterface a REST Show request.
func (service *ContrailService) RESTShowVirtualMachineInterface(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.VirtualMachineInterface
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualMachineInterface(tx, &common.ListSpec{
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
        "virtual_machine_interface": result,
    })
}

//RESTListVirtualMachineInterface handles a List REST service Request.
func (service *ContrailService) RESTListVirtualMachineInterface(c echo.Context) (error) {
    var result []*models.VirtualMachineInterface
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualMachineInterface(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "virtual-machine-interfaces": result,
    })
}