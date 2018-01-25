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

//RESTCreateBaremetalPort handle a Create REST service.
func (service *ContrailService) RESTCreateBaremetalPort(c echo.Context) error {
    requestData := &models.BaremetalPortCreateRequest{
        BaremetalPort: models.MakeBaremetalPort(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "baremetal_port",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateBaremetalPort(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateBaremetalPort handle a Create API
func (service *ContrailService) CreateBaremetalPort(
    ctx context.Context, 
    request *models.BaremetalPortCreateRequest) (*models.BaremetalPortCreateResponse, error) {
    model := request.BaremetalPort
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
            return db.CreateBaremetalPort(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "baremetal_port",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.BaremetalPortCreateResponse{
        BaremetalPort: request.BaremetalPort,
    }, nil
}

//RESTUpdateBaremetalPort handles a REST Update request.
func (service *ContrailService) RESTUpdateBaremetalPort(c echo.Context) error {
    id := c.Param("id")
    request := &models.BaremetalPortUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "baremetal_port",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateBaremetalPort(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateBaremetalPort handles a Update request.
func (service *ContrailService) UpdateBaremetalPort(ctx context.Context, request *models.BaremetalPortUpdateRequest) (*models.BaremetalPortUpdateResponse, error) {
    id = request.ID
    model = request.BaremetalPort
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
            return db.UpdateBaremetalPort(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "baremetal_port",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.BaremetalPort.UpdateResponse{
        BaremetalPort: model,
    }, nil
}

//RESTDeleteBaremetalPort delete a resource using REST service.
func (service *ContrailService) RESTDeleteBaremetalPort(c echo.Context) error {
    id := c.Param("id")
    request := &models.BaremetalPortDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteBaremetalPort(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteBaremetalPort delete a resource.
func (service *ContrailService) DeleteBaremetalPort(ctx context.Context, request *models.BaremetalPortDeleteRequest) (*models.BaremetalPortDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteBaremetalPort(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.BaremetalPortDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowBaremetalPort a REST Show request.
func (service *ContrailService) RESTShowBaremetalPort(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.BaremetalPort
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBaremetalPort(tx, &common.ListSpec{
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
        "baremetal_port": result,
    })
}

//RESTListBaremetalPort handles a List REST service Request.
func (service *ContrailService) RESTListBaremetalPort(c echo.Context) (error) {
    var result []*models.BaremetalPort
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBaremetalPort(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "baremetal-ports": result,
    })
}