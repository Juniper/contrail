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

//RESTDashboardUpdateRequest for update request for REST.
type RESTDashboardUpdateRequest struct {
    Data map[string]interface{} `json:"dashboard"`
}

//RESTCreateDashboard handle a Create REST service.
func (service *ContrailService) RESTCreateDashboard(c echo.Context) error {
    requestData := &models.DashboardCreateRequest{
        Dashboard: models.MakeDashboard(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "dashboard",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateDashboard(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateDashboard handle a Create API
func (service *ContrailService) CreateDashboard(
    ctx context.Context, 
    request *models.DashboardCreateRequest) (*models.DashboardCreateResponse, error) {
    model := request.Dashboard
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
            return db.CreateDashboard(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "dashboard",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.DashboardCreateResponse{
        Dashboard: request.Dashboard,
    }, nil
}

//RESTUpdateDashboard handles a REST Update request.
func (service *ContrailService) RESTUpdateDashboard(c echo.Context) error {
    id := c.Param("id")
    request := &models.DashboardUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "dashboard",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateDashboard(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateDashboard handles a Update request.
func (service *ContrailService) UpdateDashboard(ctx context.Context, request *models.DashboardUpdateRequest) (*models.DashboardUpdateResponse, error) {
    id = request.ID
    model = request.Dashboard
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
            return db.UpdateDashboard(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "dashboard",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.Dashboard.UpdateResponse{
        Dashboard: model,
    }, nil
}

//RESTDeleteDashboard delete a resource using REST service.
func (service *ContrailService) RESTDeleteDashboard(c echo.Context) error {
    id := c.Param("id")
    request := &models.DashboardDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteDashboard(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteDashboard delete a resource.
func (service *ContrailService) DeleteDashboard(ctx context.Context, request *models.DashboardDeleteRequest) (*models.DashboardDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteDashboard(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DashboardDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowDashboard a REST Show request.
func (service *ContrailService) RESTShowDashboard(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.Dashboard
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDashboard(tx, &common.ListSpec{
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
        "dashboard": result,
    })
}

//RESTListDashboard handles a List REST service Request.
func (service *ContrailService) RESTListDashboard(c echo.Context) (error) {
    var result []*models.Dashboard
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDashboard(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "dashboards": result,
    })
}