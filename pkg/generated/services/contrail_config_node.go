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

//RESTCreateContrailConfigNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailConfigNode(c echo.Context) error {
    requestData := &models.CreateContrailConfigNodeRequest{
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_config_node",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateContrailConfigNode(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateContrailConfigNode handle a Create API
func (service *ContrailService) CreateContrailConfigNode(
    ctx context.Context, 
    request *models.CreateContrailConfigNodeRequest) (*models.CreateContrailConfigNodeResponse, error) {
    model := request.ContrailConfigNode
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
            return db.CreateContrailConfigNode(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_config_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateContrailConfigNodeResponse{
        ContrailConfigNode: request.ContrailConfigNode,
    }, nil
}

//RESTUpdateContrailConfigNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailConfigNode(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateContrailConfigNodeRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "contrail_config_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateContrailConfigNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateContrailConfigNode handles a Update request.
func (service *ContrailService) UpdateContrailConfigNode(
    ctx context.Context, 
    request *models.UpdateContrailConfigNodeRequest) (*models.UpdateContrailConfigNodeResponse, error) {
    model := request.ContrailConfigNode
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateContrailConfigNode(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_config_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateContrailConfigNodeResponse{
        ContrailConfigNode: model,
    }, nil
}

//RESTDeleteContrailConfigNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailConfigNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteContrailConfigNodeRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteContrailConfigNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailConfigNode delete a resource.
func (service *ContrailService) DeleteContrailConfigNode(ctx context.Context, request *models.DeleteContrailConfigNodeRequest) (*models.DeleteContrailConfigNodeResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteContrailConfigNode(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteContrailConfigNodeResponse{
        ID: request.ID,
    }, nil
}

//RESTGetContrailConfigNode a REST Get request.
func (service *ContrailService) RESTGetContrailConfigNode(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetContrailConfigNodeRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetContrailConfigNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetContrailConfigNode a Get request.
func (service *ContrailService) GetContrailConfigNode(ctx context.Context, request *models.GetContrailConfigNodeRequest) (response *models.GetContrailConfigNodeResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListContrailConfigNodeRequest{
        Spec: spec,
    }
    var result *models.ListContrailConfigNodeResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailConfigNode(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.ContrailConfigNodes) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetContrailConfigNodeResponse{
       ContrailConfigNode: result.ContrailConfigNodes[0],
    }
    return response, nil
}

//RESTListContrailConfigNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailConfigNode(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListContrailConfigNodeRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListContrailConfigNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListContrailConfigNode handles a List service Request.
func (service *ContrailService) ListContrailConfigNode(
    ctx context.Context, 
    request *models.ListContrailConfigNodeRequest) (response *models.ListContrailConfigNodeResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListContrailConfigNode(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}