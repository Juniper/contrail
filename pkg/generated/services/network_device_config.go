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

//RESTNetworkDeviceConfigUpdateRequest for update request for REST.
type RESTNetworkDeviceConfigUpdateRequest struct {
    Data map[string]interface{} `json:"network-device-config"`
}

//RESTCreateNetworkDeviceConfig handle a Create REST service.
func (service *ContrailService) RESTCreateNetworkDeviceConfig(c echo.Context) error {
    requestData := &models.NetworkDeviceConfigCreateRequest{
        NetworkDeviceConfig: models.MakeNetworkDeviceConfig(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "network_device_config",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateNetworkDeviceConfig(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateNetworkDeviceConfig handle a Create API
func (service *ContrailService) CreateNetworkDeviceConfig(
    ctx context.Context, 
    request *models.NetworkDeviceConfigCreateRequest) (*models.NetworkDeviceConfigCreateResponse, error) {
    model := request.NetworkDeviceConfig
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
            return db.CreateNetworkDeviceConfig(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "network_device_config",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.NetworkDeviceConfigCreateResponse{
        NetworkDeviceConfig: request.NetworkDeviceConfig,
    }, nil
}

//RESTUpdateNetworkDeviceConfig handles a REST Update request.
func (service *ContrailService) RESTUpdateNetworkDeviceConfig(c echo.Context) error {
    id := c.Param("id")
    request := &models.NetworkDeviceConfigUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "network_device_config",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateNetworkDeviceConfig(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateNetworkDeviceConfig handles a Update request.
func (service *ContrailService) UpdateNetworkDeviceConfig(ctx context.Context, request *models.NetworkDeviceConfigUpdateRequest) (*models.NetworkDeviceConfigUpdateResponse, error) {
    id = request.ID
    model = request.NetworkDeviceConfig
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
            return db.UpdateNetworkDeviceConfig(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "network_device_config",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.NetworkDeviceConfig.UpdateResponse{
        NetworkDeviceConfig: model,
    }, nil
}

//RESTDeleteNetworkDeviceConfig delete a resource using REST service.
func (service *ContrailService) RESTDeleteNetworkDeviceConfig(c echo.Context) error {
    id := c.Param("id")
    request := &models.NetworkDeviceConfigDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteNetworkDeviceConfig(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteNetworkDeviceConfig delete a resource.
func (service *ContrailService) DeleteNetworkDeviceConfig(ctx context.Context, request *models.NetworkDeviceConfigDeleteRequest) (*models.NetworkDeviceConfigDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteNetworkDeviceConfig(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.NetworkDeviceConfigDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowNetworkDeviceConfig a REST Show request.
func (service *ContrailService) RESTShowNetworkDeviceConfig(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.NetworkDeviceConfig
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListNetworkDeviceConfig(tx, &common.ListSpec{
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
        "network_device_config": result,
    })
}

//RESTListNetworkDeviceConfig handles a List REST service Request.
func (service *ContrailService) RESTListNetworkDeviceConfig(c echo.Context) (error) {
    var result []*models.NetworkDeviceConfig
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListNetworkDeviceConfig(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "network-device-configs": result,
    })
}