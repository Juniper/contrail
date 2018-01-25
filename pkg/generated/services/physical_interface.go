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

//RESTCreatePhysicalInterface handle a Create REST service.
func (service *ContrailService) RESTCreatePhysicalInterface(c echo.Context) error {
    requestData := &models.PhysicalInterfaceCreateRequest{
        PhysicalInterface: models.MakePhysicalInterface(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "physical_interface",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreatePhysicalInterface(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreatePhysicalInterface handle a Create API
func (service *ContrailService) CreatePhysicalInterface(
    ctx context.Context, 
    request *models.PhysicalInterfaceCreateRequest) (*models.PhysicalInterfaceCreateResponse, error) {
    model := request.PhysicalInterface
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
            return db.CreatePhysicalInterface(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "physical_interface",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.PhysicalInterfaceCreateResponse{
        PhysicalInterface: request.PhysicalInterface,
    }, nil
}

//RESTUpdatePhysicalInterface handles a REST Update request.
func (service *ContrailService) RESTUpdatePhysicalInterface(c echo.Context) error {
    id := c.Param("id")
    request := &models.PhysicalInterfaceUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "physical_interface",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdatePhysicalInterface(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdatePhysicalInterface handles a Update request.
func (service *ContrailService) UpdatePhysicalInterface(ctx context.Context, request *models.PhysicalInterfaceUpdateRequest) (*models.PhysicalInterfaceUpdateResponse, error) {
    id = request.ID
    model = request.PhysicalInterface
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
            return db.UpdatePhysicalInterface(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "physical_interface",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.PhysicalInterface.UpdateResponse{
        PhysicalInterface: model,
    }, nil
}

//RESTDeletePhysicalInterface delete a resource using REST service.
func (service *ContrailService) RESTDeletePhysicalInterface(c echo.Context) error {
    id := c.Param("id")
    request := &models.PhysicalInterfaceDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeletePhysicalInterface(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeletePhysicalInterface delete a resource.
func (service *ContrailService) DeletePhysicalInterface(ctx context.Context, request *models.PhysicalInterfaceDeleteRequest) (*models.PhysicalInterfaceDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeletePhysicalInterface(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.PhysicalInterfaceDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowPhysicalInterface a REST Show request.
func (service *ContrailService) RESTShowPhysicalInterface(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.PhysicalInterface
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListPhysicalInterface(tx, &common.ListSpec{
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
        "physical_interface": result,
    })
}

//RESTListPhysicalInterface handles a List REST service Request.
func (service *ContrailService) RESTListPhysicalInterface(c echo.Context) (error) {
    var result []*models.PhysicalInterface
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListPhysicalInterface(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "physical-interfaces": result,
    })
}