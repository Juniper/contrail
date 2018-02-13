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

//RESTNetworkIpamUpdateRequest for update request for REST.
type RESTNetworkIpamUpdateRequest struct {
    Data map[string]interface{} `json:"network-ipam"`
}

//RESTCreateNetworkIpam handle a Create REST service.
func (service *ContrailService) RESTCreateNetworkIpam(c echo.Context) error {
    requestData := &models.NetworkIpamCreateRequest{
        NetworkIpam: models.MakeNetworkIpam(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "network_ipam",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateNetworkIpam(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateNetworkIpam handle a Create API
func (service *ContrailService) CreateNetworkIpam(
    ctx context.Context, 
    request *models.NetworkIpamCreateRequest) (*models.NetworkIpamCreateResponse, error) {
    model := request.NetworkIpam
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
            return db.CreateNetworkIpam(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "network_ipam",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.NetworkIpamCreateResponse{
        NetworkIpam: request.NetworkIpam,
    }, nil
}

//RESTUpdateNetworkIpam handles a REST Update request.
func (service *ContrailService) RESTUpdateNetworkIpam(c echo.Context) error {
    id := c.Param("id")
    request := &models.NetworkIpamUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "network_ipam",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateNetworkIpam(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateNetworkIpam handles a Update request.
func (service *ContrailService) UpdateNetworkIpam(ctx context.Context, request *models.NetworkIpamUpdateRequest) (*models.NetworkIpamUpdateResponse, error) {
    id = request.ID
    model = request.NetworkIpam
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
            return db.UpdateNetworkIpam(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "network_ipam",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.NetworkIpam.UpdateResponse{
        NetworkIpam: model,
    }, nil
}

//RESTDeleteNetworkIpam delete a resource using REST service.
func (service *ContrailService) RESTDeleteNetworkIpam(c echo.Context) error {
    id := c.Param("id")
    request := &models.NetworkIpamDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteNetworkIpam(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteNetworkIpam delete a resource.
func (service *ContrailService) DeleteNetworkIpam(ctx context.Context, request *models.NetworkIpamDeleteRequest) (*models.NetworkIpamDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteNetworkIpam(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.NetworkIpamDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowNetworkIpam a REST Show request.
func (service *ContrailService) RESTShowNetworkIpam(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.NetworkIpam
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListNetworkIpam(tx, &common.ListSpec{
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
        "network_ipam": result,
    })
}

//RESTListNetworkIpam handles a List REST service Request.
func (service *ContrailService) RESTListNetworkIpam(c echo.Context) (error) {
    var result []*models.NetworkIpam
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListNetworkIpam(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "network-ipams": result,
    })
}