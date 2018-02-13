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

//RESTDatabaseNodeUpdateRequest for update request for REST.
type RESTDatabaseNodeUpdateRequest struct {
    Data map[string]interface{} `json:"database-node"`
}

//RESTCreateDatabaseNode handle a Create REST service.
func (service *ContrailService) RESTCreateDatabaseNode(c echo.Context) error {
    requestData := &models.DatabaseNodeCreateRequest{
        DatabaseNode: models.MakeDatabaseNode(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "database_node",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateDatabaseNode(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateDatabaseNode handle a Create API
func (service *ContrailService) CreateDatabaseNode(
    ctx context.Context, 
    request *models.DatabaseNodeCreateRequest) (*models.DatabaseNodeCreateResponse, error) {
    model := request.DatabaseNode
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
            return db.CreateDatabaseNode(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "database_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.DatabaseNodeCreateResponse{
        DatabaseNode: request.DatabaseNode,
    }, nil
}

//RESTUpdateDatabaseNode handles a REST Update request.
func (service *ContrailService) RESTUpdateDatabaseNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.DatabaseNodeUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "database_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateDatabaseNode(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateDatabaseNode handles a Update request.
func (service *ContrailService) UpdateDatabaseNode(ctx context.Context, request *models.DatabaseNodeUpdateRequest) (*models.DatabaseNodeUpdateResponse, error) {
    id = request.ID
    model = request.DatabaseNode
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
            return db.UpdateDatabaseNode(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "database_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.DatabaseNode.UpdateResponse{
        DatabaseNode: model,
    }, nil
}

//RESTDeleteDatabaseNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteDatabaseNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.DatabaseNodeDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteDatabaseNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteDatabaseNode delete a resource.
func (service *ContrailService) DeleteDatabaseNode(ctx context.Context, request *models.DatabaseNodeDeleteRequest) (*models.DatabaseNodeDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteDatabaseNode(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DatabaseNodeDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowDatabaseNode a REST Show request.
func (service *ContrailService) RESTShowDatabaseNode(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.DatabaseNode
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDatabaseNode(tx, &common.ListSpec{
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
        "database_node": result,
    })
}

//RESTListDatabaseNode handles a List REST service Request.
func (service *ContrailService) RESTListDatabaseNode(c echo.Context) (error) {
    var result []*models.DatabaseNode
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDatabaseNode(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "database-nodes": result,
    })
}