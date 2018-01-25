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

//RESTCreateBaremetalNode handle a Create REST service.
func (service *ContrailService) RESTCreateBaremetalNode(c echo.Context) error {
    requestData := &models.BaremetalNodeCreateRequest{
        BaremetalNode: models.MakeBaremetalNode(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "baremetal_node",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateBaremetalNode(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateBaremetalNode handle a Create API
func (service *ContrailService) CreateBaremetalNode(
    ctx context.Context, 
    request *models.BaremetalNodeCreateRequest) (*models.BaremetalNodeCreateResponse, error) {
    model := request.BaremetalNode
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
            return db.CreateBaremetalNode(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "baremetal_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.BaremetalNodeCreateResponse{
        BaremetalNode: request.BaremetalNode,
    }, nil
}

//RESTUpdateBaremetalNode handles a REST Update request.
func (service *ContrailService) RESTUpdateBaremetalNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.BaremetalNodeUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "baremetal_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateBaremetalNode(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateBaremetalNode handles a Update request.
func (service *ContrailService) UpdateBaremetalNode(ctx context.Context, request *models.BaremetalNodeUpdateRequest) (*models.BaremetalNodeUpdateResponse, error) {
    id = request.ID
    model = request.BaremetalNode
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
            return db.UpdateBaremetalNode(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "baremetal_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.BaremetalNode.UpdateResponse{
        BaremetalNode: model,
    }, nil
}

//RESTDeleteBaremetalNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteBaremetalNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.BaremetalNodeDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteBaremetalNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteBaremetalNode delete a resource.
func (service *ContrailService) DeleteBaremetalNode(ctx context.Context, request *models.BaremetalNodeDeleteRequest) (*models.BaremetalNodeDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteBaremetalNode(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.BaremetalNodeDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowBaremetalNode a REST Show request.
func (service *ContrailService) RESTShowBaremetalNode(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.BaremetalNode
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBaremetalNode(tx, &common.ListSpec{
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
        "baremetal_node": result,
    })
}

//RESTListBaremetalNode handles a List REST service Request.
func (service *ContrailService) RESTListBaremetalNode(c echo.Context) (error) {
    var result []*models.BaremetalNode
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBaremetalNode(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "baremetal-nodes": result,
    })
}