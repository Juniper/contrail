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

//RESTConfigNodeUpdateRequest for update request for REST.
type RESTConfigNodeUpdateRequest struct {
    Data map[string]interface{} `json:"config-node"`
}

//RESTCreateConfigNode handle a Create REST service.
func (service *ContrailService) RESTCreateConfigNode(c echo.Context) error {
    requestData := &models.ConfigNodeCreateRequest{
        ConfigNode: models.MakeConfigNode(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "config_node",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateConfigNode(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateConfigNode handle a Create API
func (service *ContrailService) CreateConfigNode(
    ctx context.Context, 
    request *models.ConfigNodeCreateRequest) (*models.ConfigNodeCreateResponse, error) {
    model := request.ConfigNode
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
            return db.CreateConfigNode(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "config_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ConfigNodeCreateResponse{
        ConfigNode: request.ConfigNode,
    }, nil
}

//RESTUpdateConfigNode handles a REST Update request.
func (service *ContrailService) RESTUpdateConfigNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.ConfigNodeUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "config_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateConfigNode(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateConfigNode handles a Update request.
func (service *ContrailService) UpdateConfigNode(ctx context.Context, request *models.ConfigNodeUpdateRequest) (*models.ConfigNodeUpdateResponse, error) {
    id = request.ID
    model = request.ConfigNode
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
            return db.UpdateConfigNode(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "config_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ConfigNode.UpdateResponse{
        ConfigNode: model,
    }, nil
}

//RESTDeleteConfigNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteConfigNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.ConfigNodeDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteConfigNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteConfigNode delete a resource.
func (service *ContrailService) DeleteConfigNode(ctx context.Context, request *models.ConfigNodeDeleteRequest) (*models.ConfigNodeDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteConfigNode(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ConfigNodeDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowConfigNode a REST Show request.
func (service *ContrailService) RESTShowConfigNode(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ConfigNode
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListConfigNode(tx, &common.ListSpec{
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
        "config_node": result,
    })
}

//RESTListConfigNode handles a List REST service Request.
func (service *ContrailService) RESTListConfigNode(c echo.Context) (error) {
    var result []*models.ConfigNode
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListConfigNode(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "config-nodes": result,
    })
}