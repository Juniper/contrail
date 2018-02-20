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

//RESTCreateContrailStorageNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailStorageNode(c echo.Context) error {
    requestData := &models.CreateContrailStorageNodeRequest{
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_storage_node",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateContrailStorageNode(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateContrailStorageNode handle a Create API
func (service *ContrailService) CreateContrailStorageNode(
    ctx context.Context, 
    request *models.CreateContrailStorageNodeRequest) (*models.CreateContrailStorageNodeResponse, error) {
    model := request.ContrailStorageNode
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
            return db.CreateContrailStorageNode(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_storage_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateContrailStorageNodeResponse{
        ContrailStorageNode: request.ContrailStorageNode,
    }, nil
}

//RESTUpdateContrailStorageNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailStorageNode(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateContrailStorageNodeRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "contrail_storage_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateContrailStorageNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateContrailStorageNode handles a Update request.
func (service *ContrailService) UpdateContrailStorageNode(
    ctx context.Context, 
    request *models.UpdateContrailStorageNodeRequest) (*models.UpdateContrailStorageNodeResponse, error) {
    model := request.ContrailStorageNode
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateContrailStorageNode(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_storage_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateContrailStorageNodeResponse{
        ContrailStorageNode: model,
    }, nil
}

//RESTDeleteContrailStorageNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailStorageNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteContrailStorageNodeRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteContrailStorageNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailStorageNode delete a resource.
func (service *ContrailService) DeleteContrailStorageNode(ctx context.Context, request *models.DeleteContrailStorageNodeRequest) (*models.DeleteContrailStorageNodeResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteContrailStorageNode(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteContrailStorageNodeResponse{
        ID: request.ID,
    }, nil
}

//RESTGetContrailStorageNode a REST Get request.
func (service *ContrailService) RESTGetContrailStorageNode(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetContrailStorageNodeRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetContrailStorageNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetContrailStorageNode a Get request.
func (service *ContrailService) GetContrailStorageNode(ctx context.Context, request *models.GetContrailStorageNodeRequest) (response *models.GetContrailStorageNodeResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListContrailStorageNodeRequest{
        Spec: spec,
    }
    var result *models.ListContrailStorageNodeResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailStorageNode(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.ContrailStorageNodes) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetContrailStorageNodeResponse{
       ContrailStorageNode: result.ContrailStorageNodes[0],
    }
    return response, nil
}

//RESTListContrailStorageNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailStorageNode(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListContrailStorageNodeRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListContrailStorageNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListContrailStorageNode handles a List service Request.
func (service *ContrailService) ListContrailStorageNode(
    ctx context.Context, 
    request *models.ListContrailStorageNodeRequest) (response *models.ListContrailStorageNodeResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListContrailStorageNode(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}