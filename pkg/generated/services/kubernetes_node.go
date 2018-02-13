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

//RESTKubernetesNodeUpdateRequest for update request for REST.
type RESTKubernetesNodeUpdateRequest struct {
    Data map[string]interface{} `json:"kubernetes-node"`
}

//RESTCreateKubernetesNode handle a Create REST service.
func (service *ContrailService) RESTCreateKubernetesNode(c echo.Context) error {
    requestData := &models.KubernetesNodeCreateRequest{
        KubernetesNode: models.MakeKubernetesNode(),
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
    request *models.KubernetesNodeCreateRequest) (*models.KubernetesNodeCreateResponse, error) {
    model := request.KubernetesNode
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
            return db.CreateKubernetesNode(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "kubernetes_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.KubernetesNodeCreateResponse{
        KubernetesNode: request.KubernetesNode,
    }, nil
}

//RESTUpdateKubernetesNode handles a REST Update request.
func (service *ContrailService) RESTUpdateKubernetesNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.KubernetesNodeUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "kubernetes_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateKubernetesNode(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateKubernetesNode handles a Update request.
func (service *ContrailService) UpdateKubernetesNode(ctx context.Context, request *models.KubernetesNodeUpdateRequest) (*models.KubernetesNodeUpdateResponse, error) {
    id = request.ID
    model = request.KubernetesNode
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
            return db.UpdateKubernetesNode(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "kubernetes_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.KubernetesNode.UpdateResponse{
        KubernetesNode: model,
    }, nil
}

//RESTDeleteKubernetesNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteKubernetesNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.KubernetesNodeDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteKubernetesNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteKubernetesNode delete a resource.
func (service *ContrailService) DeleteKubernetesNode(ctx context.Context, request *models.KubernetesNodeDeleteRequest) (*models.KubernetesNodeDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteKubernetesNode(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.KubernetesNodeDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowKubernetesNode a REST Show request.
func (service *ContrailService) RESTShowKubernetesNode(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.KubernetesNode
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListKubernetesNode(tx, &common.ListSpec{
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
        "kubernetes_node": result,
    })
}

//RESTListKubernetesNode handles a List REST service Request.
func (service *ContrailService) RESTListKubernetesNode(c echo.Context) (error) {
    var result []*models.KubernetesNode
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListKubernetesNode(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "kubernetes-nodes": result,
    })
}