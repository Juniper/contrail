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

//RESTCreateKubernetesNode handle a Create REST service.
func (service *ContrailService) RESTCreateKubernetesNode(c echo.Context) error {
    requestData := &models.CreateKubernetesNodeRequest{
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "kubernetes_node",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateKubernetesNode(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateKubernetesNode handle a Create API
func (service *ContrailService) CreateKubernetesNode(
    ctx context.Context, 
    request *models.CreateKubernetesNodeRequest) (*models.CreateKubernetesNodeResponse, error) {
    model := request.KubernetesNode
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
            return db.CreateKubernetesNode(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "kubernetes_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateKubernetesNodeResponse{
        KubernetesNode: request.KubernetesNode,
    }, nil
}

//RESTUpdateKubernetesNode handles a REST Update request.
func (service *ContrailService) RESTUpdateKubernetesNode(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateKubernetesNodeRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "kubernetes_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateKubernetesNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateKubernetesNode handles a Update request.
func (service *ContrailService) UpdateKubernetesNode(
    ctx context.Context, 
    request *models.UpdateKubernetesNodeRequest) (*models.UpdateKubernetesNodeResponse, error) {
    model := request.KubernetesNode
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateKubernetesNode(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "kubernetes_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateKubernetesNodeResponse{
        KubernetesNode: model,
    }, nil
}

//RESTDeleteKubernetesNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteKubernetesNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteKubernetesNodeRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteKubernetesNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteKubernetesNode delete a resource.
func (service *ContrailService) DeleteKubernetesNode(ctx context.Context, request *models.DeleteKubernetesNodeRequest) (*models.DeleteKubernetesNodeResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteKubernetesNode(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteKubernetesNodeResponse{
        ID: request.ID,
    }, nil
}

//RESTGetKubernetesNode a REST Get request.
func (service *ContrailService) RESTGetKubernetesNode(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetKubernetesNodeRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetKubernetesNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetKubernetesNode a Get request.
func (service *ContrailService) GetKubernetesNode(ctx context.Context, request *models.GetKubernetesNodeRequest) (response *models.GetKubernetesNodeResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListKubernetesNodeRequest{
        Spec: spec,
    }
    var result *models.ListKubernetesNodeResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListKubernetesNode(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.KubernetesNodes) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetKubernetesNodeResponse{
       KubernetesNode: result.KubernetesNodes[0],
    }
    return response, nil
}

//RESTListKubernetesNode handles a List REST service Request.
func (service *ContrailService) RESTListKubernetesNode(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListKubernetesNodeRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListKubernetesNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListKubernetesNode handles a List service Request.
func (service *ContrailService) ListKubernetesNode(
    ctx context.Context, 
    request *models.ListKubernetesNodeRequest) (response *models.ListKubernetesNodeResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListKubernetesNode(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}