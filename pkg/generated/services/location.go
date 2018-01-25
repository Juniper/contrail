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

//RESTCreateLocation handle a Create REST service.
func (service *ContrailService) RESTCreateLocation(c echo.Context) error {
    requestData := &models.LocationCreateRequest{
        Location: models.MakeLocation(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "location",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateLocation(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateLocation handle a Create API
func (service *ContrailService) CreateLocation(
    ctx context.Context, 
    request *models.LocationCreateRequest) (*models.LocationCreateResponse, error) {
    model := request.Location
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
            return db.CreateLocation(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "location",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.LocationCreateResponse{
        Location: request.Location,
    }, nil
}

//RESTUpdateLocation handles a REST Update request.
func (service *ContrailService) RESTUpdateLocation(c echo.Context) error {
    id := c.Param("id")
    request := &models.LocationUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "location",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateLocation(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateLocation handles a Update request.
func (service *ContrailService) UpdateLocation(ctx context.Context, request *models.LocationUpdateRequest) (*models.LocationUpdateResponse, error) {
    id = request.ID
    model = request.Location
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
            return db.UpdateLocation(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "location",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.Location.UpdateResponse{
        Location: model,
    }, nil
}

//RESTDeleteLocation delete a resource using REST service.
func (service *ContrailService) RESTDeleteLocation(c echo.Context) error {
    id := c.Param("id")
    request := &models.LocationDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteLocation(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteLocation delete a resource.
func (service *ContrailService) DeleteLocation(ctx context.Context, request *models.LocationDeleteRequest) (*models.LocationDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteLocation(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.LocationDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowLocation a REST Show request.
func (service *ContrailService) RESTShowLocation(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.Location
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLocation(tx, &common.ListSpec{
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
        "location": result,
    })
}

//RESTListLocation handles a List REST service Request.
func (service *ContrailService) RESTListLocation(c echo.Context) (error) {
    var result []*models.Location
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLocation(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "locations": result,
    })
}