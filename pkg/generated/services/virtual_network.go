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

//RESTVirtualNetworkUpdateRequest for update request for REST.
type RESTVirtualNetworkUpdateRequest struct {
    Data map[string]interface{} `json:"virtual-network"`
}

//RESTCreateVirtualNetwork handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualNetwork(c echo.Context) error {
    requestData := &models.VirtualNetworkCreateRequest{
        VirtualNetwork: models.MakeVirtualNetwork(),
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
    request *models.VirtualNetworkCreateRequest) (*models.VirtualNetworkCreateResponse, error) {
    model := request.VirtualNetwork
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
            return db.CreateVirtualNetwork(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_network",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.VirtualNetworkCreateResponse{
        VirtualNetwork: request.VirtualNetwork,
    }, nil
}

//RESTUpdateVirtualNetwork handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualNetwork(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualNetworkUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "virtual_network",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateVirtualNetwork(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateVirtualNetwork handles a Update request.
func (service *ContrailService) UpdateVirtualNetwork(ctx context.Context, request *models.VirtualNetworkUpdateRequest) (*models.VirtualNetworkUpdateResponse, error) {
    id = request.ID
    model = request.VirtualNetwork
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
            return db.UpdateVirtualNetwork(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "virtual_network",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.VirtualNetwork.UpdateResponse{
        VirtualNetwork: model,
    }, nil
}

//RESTDeleteVirtualNetwork delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualNetwork(c echo.Context) error {
    id := c.Param("id")
    request := &models.VirtualNetworkDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteVirtualNetwork(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualNetwork delete a resource.
func (service *ContrailService) DeleteVirtualNetwork(ctx context.Context, request *models.VirtualNetworkDeleteRequest) (*models.VirtualNetworkDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteVirtualNetwork(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.VirtualNetworkDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowVirtualNetwork a REST Show request.
func (service *ContrailService) RESTShowVirtualNetwork(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.VirtualNetwork
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualNetwork(tx, &common.ListSpec{
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
        "virtual_network": result,
    })
}

//RESTListVirtualNetwork handles a List REST service Request.
func (service *ContrailService) RESTListVirtualNetwork(c echo.Context) (error) {
    var result []*models.VirtualNetwork
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVirtualNetwork(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "virtual-networks": result,
    })
}