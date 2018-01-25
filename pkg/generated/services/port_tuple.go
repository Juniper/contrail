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

//RESTCreatePortTuple handle a Create REST service.
func (service *ContrailService) RESTCreatePortTuple(c echo.Context) error {
    requestData := &models.PortTupleCreateRequest{
        PortTuple: models.MakePortTuple(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "port_tuple",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreatePortTuple(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreatePortTuple handle a Create API
func (service *ContrailService) CreatePortTuple(
    ctx context.Context, 
    request *models.PortTupleCreateRequest) (*models.PortTupleCreateResponse, error) {
    model := request.PortTuple
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
            return db.CreatePortTuple(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "port_tuple",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.PortTupleCreateResponse{
        PortTuple: request.PortTuple,
    }, nil
}

//RESTUpdatePortTuple handles a REST Update request.
func (service *ContrailService) RESTUpdatePortTuple(c echo.Context) error {
    id := c.Param("id")
    request := &models.PortTupleUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "port_tuple",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdatePortTuple(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdatePortTuple handles a Update request.
func (service *ContrailService) UpdatePortTuple(ctx context.Context, request *models.PortTupleUpdateRequest) (*models.PortTupleUpdateResponse, error) {
    id = request.ID
    model = request.PortTuple
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
            return db.UpdatePortTuple(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "port_tuple",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.PortTuple.UpdateResponse{
        PortTuple: model,
    }, nil
}

//RESTDeletePortTuple delete a resource using REST service.
func (service *ContrailService) RESTDeletePortTuple(c echo.Context) error {
    id := c.Param("id")
    request := &models.PortTupleDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeletePortTuple(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeletePortTuple delete a resource.
func (service *ContrailService) DeletePortTuple(ctx context.Context, request *models.PortTupleDeleteRequest) (*models.PortTupleDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeletePortTuple(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.PortTupleDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowPortTuple a REST Show request.
func (service *ContrailService) RESTShowPortTuple(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.PortTuple
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListPortTuple(tx, &common.ListSpec{
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
        "port_tuple": result,
    })
}

//RESTListPortTuple handles a List REST service Request.
func (service *ContrailService) RESTListPortTuple(c echo.Context) (error) {
    var result []*models.PortTuple
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListPortTuple(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "port-tuples": result,
    })
}