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

//RESTWidgetUpdateRequest for update request for REST.
type RESTWidgetUpdateRequest struct {
    Data map[string]interface{} `json:"widget"`
}

//RESTCreateWidget handle a Create REST service.
func (service *ContrailService) RESTCreateWidget(c echo.Context) error {
    requestData := &models.WidgetCreateRequest{
        Widget: models.MakeWidget(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "widget",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateWidget(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateWidget handle a Create API
func (service *ContrailService) CreateWidget(
    ctx context.Context, 
    request *models.WidgetCreateRequest) (*models.WidgetCreateResponse, error) {
    model := request.Widget
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
            return db.CreateWidget(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "widget",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.WidgetCreateResponse{
        Widget: request.Widget,
    }, nil
}

//RESTUpdateWidget handles a REST Update request.
func (service *ContrailService) RESTUpdateWidget(c echo.Context) error {
    id := c.Param("id")
    request := &models.WidgetUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "widget",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateWidget(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateWidget handles a Update request.
func (service *ContrailService) UpdateWidget(ctx context.Context, request *models.WidgetUpdateRequest) (*models.WidgetUpdateResponse, error) {
    id = request.ID
    model = request.Widget
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
            return db.UpdateWidget(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "widget",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.Widget.UpdateResponse{
        Widget: model,
    }, nil
}

//RESTDeleteWidget delete a resource using REST service.
func (service *ContrailService) RESTDeleteWidget(c echo.Context) error {
    id := c.Param("id")
    request := &models.WidgetDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteWidget(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteWidget delete a resource.
func (service *ContrailService) DeleteWidget(ctx context.Context, request *models.WidgetDeleteRequest) (*models.WidgetDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteWidget(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.WidgetDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowWidget a REST Show request.
func (service *ContrailService) RESTShowWidget(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.Widget
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListWidget(tx, &common.ListSpec{
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
        "widget": result,
    })
}

//RESTListWidget handles a List REST service Request.
func (service *ContrailService) RESTListWidget(c echo.Context) (error) {
    var result []*models.Widget
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListWidget(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "widgets": result,
    })
}