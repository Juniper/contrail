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

//RESTCreateVirtualNetwork handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualNetwork(c echo.Context) error {
    requestData := &models.CreateVirtualNetworkRequest{
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_network",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateVirtualNetwork(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateVirtualNetwork handle a Create API
func (service *ContrailService) CreateVirtualNetwork(
    ctx context.Context, 
    request *models.CreateVirtualNetworkRequest) (*models.CreateVirtualNetworkResponse, error) {
    model := request.VirtualNetwork
    if model.UUID == "" {
        model.UUID = uuid.NewV4().String()
    }
    auth := common.GetAuthCTX(ctx)
    if auth == nil {
        return nil, common.ErrorUnauthenticated
    }

    if model.FQName == nil {
        if model.DisplayName == "" {
        return nil, common.ErrorBadRequest("Both of FQName and Display Name is empty") 
        }
        model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
    }
    model.Perms2 = &models.PermType2{}
    model.Perms2.Owner = auth.ProjectID()
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.CreateVirtualNetwork(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_network",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateVirtualNetworkResponse{
        VirtualNetwork: request.VirtualNetwork,
    }, nil
}

//RESTUpdateVirtualNetwork handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualNetwork(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateVirtualNetworkRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "virtual_network",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateVirtualNetwork(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateVirtualNetwork handles a Update request.
func (service *ContrailService) UpdateVirtualNetwork(
    ctx context.Context, 
    request *models.UpdateVirtualNetworkRequest) (*models.UpdateVirtualNetworkResponse, error) {
    model := request.VirtualNetwork
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateVirtualNetwork(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_network",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateVirtualNetworkResponse{
        VirtualNetwork: model,
    }, nil
}

//RESTDeleteVirtualNetwork delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualNetwork(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteVirtualNetworkRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteVirtualNetwork(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualNetwork delete a resource.
func (service *ContrailService) DeleteVirtualNetwork(ctx context.Context, request *models.DeleteVirtualNetworkRequest) (*models.DeleteVirtualNetworkResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteVirtualNetwork(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteVirtualNetworkResponse{
        ID: request.ID,
    }, nil
}

//RESTGetVirtualNetwork a REST Get request.
func (service *ContrailService) RESTGetVirtualNetwork(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetVirtualNetworkRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetVirtualNetwork(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetVirtualNetwork a Get request.
func (service *ContrailService) GetVirtualNetwork(ctx context.Context, request *models.GetVirtualNetworkRequest) (response *models.GetVirtualNetworkResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListVirtualNetworkRequest{
        Spec: spec,
    }
    var result *models.ListVirtualNetworkResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualNetwork(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.VirtualNetworks) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetVirtualNetworkResponse{
       VirtualNetwork: result.VirtualNetworks[0],
    }
    return response, nil
}

//RESTListVirtualNetwork handles a List REST service Request.
func (service *ContrailService) RESTListVirtualNetwork(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListVirtualNetworkRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListVirtualNetwork(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListVirtualNetwork handles a List service Request.
func (service *ContrailService) ListVirtualNetwork(
    ctx context.Context, 
    request *models.ListVirtualNetworkRequest) (response *models.ListVirtualNetworkResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListVirtualNetwork(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}