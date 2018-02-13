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

//RESTContrailAnalyticsDatabaseNodeRoleUpdateRequest for update request for REST.
type RESTContrailAnalyticsDatabaseNodeRoleUpdateRequest struct {
    Data map[string]interface{} `json:"contrail-analytics-database-node-role"`
}

//RESTCreateContrailAnalyticsDatabaseNodeRole handle a Create REST service.
func (service *ContrailService) RESTCreateContrailAnalyticsDatabaseNodeRole(c echo.Context) error {
    requestData := &models.ContrailAnalyticsDatabaseNodeRoleCreateRequest{
        ContrailAnalyticsDatabaseNodeRole: models.MakeContrailAnalyticsDatabaseNodeRole(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_analytics_database_node_role",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateContrailAnalyticsDatabaseNodeRole(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateContrailAnalyticsDatabaseNodeRole handle a Create API
func (service *ContrailService) CreateContrailAnalyticsDatabaseNodeRole(
    ctx context.Context, 
    request *models.ContrailAnalyticsDatabaseNodeRoleCreateRequest) (*models.ContrailAnalyticsDatabaseNodeRoleCreateResponse, error) {
    model := request.ContrailAnalyticsDatabaseNodeRole
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
            return db.CreateContrailAnalyticsDatabaseNodeRole(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_analytics_database_node_role",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ContrailAnalyticsDatabaseNodeRoleCreateResponse{
        ContrailAnalyticsDatabaseNodeRole: request.ContrailAnalyticsDatabaseNodeRole,
    }, nil
}

//RESTUpdateContrailAnalyticsDatabaseNodeRole handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailAnalyticsDatabaseNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.ContrailAnalyticsDatabaseNodeRoleUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "contrail_analytics_database_node_role",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateContrailAnalyticsDatabaseNodeRole(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateContrailAnalyticsDatabaseNodeRole handles a Update request.
func (service *ContrailService) UpdateContrailAnalyticsDatabaseNodeRole(ctx context.Context, request *models.ContrailAnalyticsDatabaseNodeRoleUpdateRequest) (*models.ContrailAnalyticsDatabaseNodeRoleUpdateResponse, error) {
    id = request.ID
    model = request.ContrailAnalyticsDatabaseNodeRole
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
            return db.UpdateContrailAnalyticsDatabaseNodeRole(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_analytics_database_node_role",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ContrailAnalyticsDatabaseNodeRole.UpdateResponse{
        ContrailAnalyticsDatabaseNodeRole: model,
    }, nil
}

//RESTDeleteContrailAnalyticsDatabaseNodeRole delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailAnalyticsDatabaseNodeRole(c echo.Context) error {
    id := c.Param("id")
    request := &models.ContrailAnalyticsDatabaseNodeRoleDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteContrailAnalyticsDatabaseNodeRole(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailAnalyticsDatabaseNodeRole delete a resource.
func (service *ContrailService) DeleteContrailAnalyticsDatabaseNodeRole(ctx context.Context, request *models.ContrailAnalyticsDatabaseNodeRoleDeleteRequest) (*models.ContrailAnalyticsDatabaseNodeRoleDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteContrailAnalyticsDatabaseNodeRole(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ContrailAnalyticsDatabaseNodeRoleDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowContrailAnalyticsDatabaseNodeRole a REST Show request.
func (service *ContrailService) RESTShowContrailAnalyticsDatabaseNodeRole(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ContrailAnalyticsDatabaseNodeRole
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailAnalyticsDatabaseNodeRole(tx, &common.ListSpec{
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
        "contrail_analytics_database_node_role": result,
    })
}

//RESTListContrailAnalyticsDatabaseNodeRole handles a List REST service Request.
func (service *ContrailService) RESTListContrailAnalyticsDatabaseNodeRole(c echo.Context) (error) {
    var result []*models.ContrailAnalyticsDatabaseNodeRole
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailAnalyticsDatabaseNodeRole(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "contrail-analytics-database-node-roles": result,
    })
}