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

//RESTLogicalInterfaceUpdateRequest for update request for REST.
type RESTLogicalInterfaceUpdateRequest struct {
    Data map[string]interface{} `json:"logical-interface"`
}

//RESTCreateLogicalInterface handle a Create REST service.
func (service *ContrailService) RESTCreateLogicalInterface(c echo.Context) error {
    requestData := &models.LogicalInterfaceCreateRequest{
        LogicalInterface: models.MakeLogicalInterface(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "logical_interface",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateLogicalInterface(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateLogicalInterface handle a Create API
func (service *ContrailService) CreateLogicalInterface(
    ctx context.Context, 
    request *models.LogicalInterfaceCreateRequest) (*models.LogicalInterfaceCreateResponse, error) {
    model := request.LogicalInterface
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
            return db.CreateLogicalInterface(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "logical_interface",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.LogicalInterfaceCreateResponse{
        LogicalInterface: request.LogicalInterface,
    }, nil
}

//RESTUpdateLogicalInterface handles a REST Update request.
func (service *ContrailService) RESTUpdateLogicalInterface(c echo.Context) error {
    id := c.Param("id")
    request := &models.LogicalInterfaceUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "logical_interface",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateLogicalInterface(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateLogicalInterface handles a Update request.
func (service *ContrailService) UpdateLogicalInterface(ctx context.Context, request *models.LogicalInterfaceUpdateRequest) (*models.LogicalInterfaceUpdateResponse, error) {
    id = request.ID
    model = request.LogicalInterface
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
            return db.UpdateLogicalInterface(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "logical_interface",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.LogicalInterface.UpdateResponse{
        LogicalInterface: model,
    }, nil
}

//RESTDeleteLogicalInterface delete a resource using REST service.
func (service *ContrailService) RESTDeleteLogicalInterface(c echo.Context) error {
    id := c.Param("id")
    request := &models.LogicalInterfaceDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteLogicalInterface(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteLogicalInterface delete a resource.
func (service *ContrailService) DeleteLogicalInterface(ctx context.Context, request *models.LogicalInterfaceDeleteRequest) (*models.LogicalInterfaceDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteLogicalInterface(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.LogicalInterfaceDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowLogicalInterface a REST Show request.
func (service *ContrailService) RESTShowLogicalInterface(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.LogicalInterface
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLogicalInterface(tx, &common.ListSpec{
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
        "logical_interface": result,
    })
}

//RESTListLogicalInterface handles a List REST service Request.
func (service *ContrailService) RESTListLogicalInterface(c echo.Context) (error) {
    var result []*models.LogicalInterface
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLogicalInterface(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "logical-interfaces": result,
    })
}