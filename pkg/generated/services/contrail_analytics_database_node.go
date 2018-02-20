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

//RESTCreateContrailAnalyticsDatabaseNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailAnalyticsDatabaseNode(c echo.Context) error {
    requestData := &models.CreateContrailAnalyticsDatabaseNodeRequest{
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_analytics_database_node",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateContrailAnalyticsDatabaseNode(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateContrailAnalyticsDatabaseNode handle a Create API
func (service *ContrailService) CreateContrailAnalyticsDatabaseNode(
    ctx context.Context, 
    request *models.CreateContrailAnalyticsDatabaseNodeRequest) (*models.CreateContrailAnalyticsDatabaseNodeResponse, error) {
    model := request.ContrailAnalyticsDatabaseNode
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
            return db.CreateContrailAnalyticsDatabaseNode(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_analytics_database_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateContrailAnalyticsDatabaseNodeResponse{
        ContrailAnalyticsDatabaseNode: request.ContrailAnalyticsDatabaseNode,
    }, nil
}

//RESTUpdateContrailAnalyticsDatabaseNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailAnalyticsDatabaseNode(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateContrailAnalyticsDatabaseNodeRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "contrail_analytics_database_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateContrailAnalyticsDatabaseNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateContrailAnalyticsDatabaseNode handles a Update request.
func (service *ContrailService) UpdateContrailAnalyticsDatabaseNode(
    ctx context.Context, 
    request *models.UpdateContrailAnalyticsDatabaseNodeRequest) (*models.UpdateContrailAnalyticsDatabaseNodeResponse, error) {
    model := request.ContrailAnalyticsDatabaseNode
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateContrailAnalyticsDatabaseNode(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_analytics_database_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateContrailAnalyticsDatabaseNodeResponse{
        ContrailAnalyticsDatabaseNode: model,
    }, nil
}

//RESTDeleteContrailAnalyticsDatabaseNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailAnalyticsDatabaseNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteContrailAnalyticsDatabaseNodeRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteContrailAnalyticsDatabaseNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailAnalyticsDatabaseNode delete a resource.
func (service *ContrailService) DeleteContrailAnalyticsDatabaseNode(ctx context.Context, request *models.DeleteContrailAnalyticsDatabaseNodeRequest) (*models.DeleteContrailAnalyticsDatabaseNodeResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteContrailAnalyticsDatabaseNode(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteContrailAnalyticsDatabaseNodeResponse{
        ID: request.ID,
    }, nil
}

//RESTGetContrailAnalyticsDatabaseNode a REST Get request.
func (service *ContrailService) RESTGetContrailAnalyticsDatabaseNode(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetContrailAnalyticsDatabaseNodeRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetContrailAnalyticsDatabaseNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetContrailAnalyticsDatabaseNode a Get request.
func (service *ContrailService) GetContrailAnalyticsDatabaseNode(ctx context.Context, request *models.GetContrailAnalyticsDatabaseNodeRequest) (response *models.GetContrailAnalyticsDatabaseNodeResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListContrailAnalyticsDatabaseNodeRequest{
        Spec: spec,
    }
    var result *models.ListContrailAnalyticsDatabaseNodeResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailAnalyticsDatabaseNode(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.ContrailAnalyticsDatabaseNodes) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetContrailAnalyticsDatabaseNodeResponse{
       ContrailAnalyticsDatabaseNode: result.ContrailAnalyticsDatabaseNodes[0],
    }
    return response, nil
}

//RESTListContrailAnalyticsDatabaseNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailAnalyticsDatabaseNode(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListContrailAnalyticsDatabaseNodeRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListContrailAnalyticsDatabaseNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListContrailAnalyticsDatabaseNode handles a List service Request.
func (service *ContrailService) ListContrailAnalyticsDatabaseNode(
    ctx context.Context, 
    request *models.ListContrailAnalyticsDatabaseNodeRequest) (response *models.ListContrailAnalyticsDatabaseNodeResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListContrailAnalyticsDatabaseNode(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}