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

//RESTCreateContrailConfigDatabaseNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailConfigDatabaseNode(c echo.Context) error {
    requestData := &models.CreateContrailConfigDatabaseNodeRequest{
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_config_database_node",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateContrailConfigDatabaseNode(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateContrailConfigDatabaseNode handle a Create API
func (service *ContrailService) CreateContrailConfigDatabaseNode(
    ctx context.Context, 
    request *models.CreateContrailConfigDatabaseNodeRequest) (*models.CreateContrailConfigDatabaseNodeResponse, error) {
    model := request.ContrailConfigDatabaseNode
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
            return db.CreateContrailConfigDatabaseNode(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_config_database_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateContrailConfigDatabaseNodeResponse{
        ContrailConfigDatabaseNode: request.ContrailConfigDatabaseNode,
    }, nil
}

//RESTUpdateContrailConfigDatabaseNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailConfigDatabaseNode(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateContrailConfigDatabaseNodeRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "contrail_config_database_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateContrailConfigDatabaseNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateContrailConfigDatabaseNode handles a Update request.
func (service *ContrailService) UpdateContrailConfigDatabaseNode(
    ctx context.Context, 
    request *models.UpdateContrailConfigDatabaseNodeRequest) (*models.UpdateContrailConfigDatabaseNodeResponse, error) {
    model := request.ContrailConfigDatabaseNode
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateContrailConfigDatabaseNode(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_config_database_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateContrailConfigDatabaseNodeResponse{
        ContrailConfigDatabaseNode: model,
    }, nil
}

//RESTDeleteContrailConfigDatabaseNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailConfigDatabaseNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteContrailConfigDatabaseNodeRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteContrailConfigDatabaseNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailConfigDatabaseNode delete a resource.
func (service *ContrailService) DeleteContrailConfigDatabaseNode(ctx context.Context, request *models.DeleteContrailConfigDatabaseNodeRequest) (*models.DeleteContrailConfigDatabaseNodeResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteContrailConfigDatabaseNode(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteContrailConfigDatabaseNodeResponse{
        ID: request.ID,
    }, nil
}

//RESTGetContrailConfigDatabaseNode a REST Get request.
func (service *ContrailService) RESTGetContrailConfigDatabaseNode(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetContrailConfigDatabaseNodeRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetContrailConfigDatabaseNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetContrailConfigDatabaseNode a Get request.
func (service *ContrailService) GetContrailConfigDatabaseNode(ctx context.Context, request *models.GetContrailConfigDatabaseNodeRequest) (response *models.GetContrailConfigDatabaseNodeResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListContrailConfigDatabaseNodeRequest{
        Spec: spec,
    }
    var result *models.ListContrailConfigDatabaseNodeResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailConfigDatabaseNode(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.ContrailConfigDatabaseNodes) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetContrailConfigDatabaseNodeResponse{
       ContrailConfigDatabaseNode: result.ContrailConfigDatabaseNodes[0],
    }
    return response, nil
}

//RESTListContrailConfigDatabaseNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailConfigDatabaseNode(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListContrailConfigDatabaseNodeRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListContrailConfigDatabaseNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListContrailConfigDatabaseNode handles a List service Request.
func (service *ContrailService) ListContrailConfigDatabaseNode(
    ctx context.Context, 
    request *models.ListContrailConfigDatabaseNodeRequest) (response *models.ListContrailConfigDatabaseNodeResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListContrailConfigDatabaseNode(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}