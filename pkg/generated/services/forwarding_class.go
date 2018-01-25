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

//RESTCreateForwardingClass handle a Create REST service.
func (service *ContrailService) RESTCreateForwardingClass(c echo.Context) error {
    requestData := &models.ForwardingClassCreateRequest{
        ForwardingClass: models.MakeForwardingClass(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "forwarding_class",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateForwardingClass(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateForwardingClass handle a Create API
func (service *ContrailService) CreateForwardingClass(
    ctx context.Context, 
    request *models.ForwardingClassCreateRequest) (*models.ForwardingClassCreateResponse, error) {
    model := request.ForwardingClass
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
            return db.CreateForwardingClass(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "forwarding_class",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ForwardingClassCreateResponse{
        ForwardingClass: request.ForwardingClass,
    }, nil
}

//RESTUpdateForwardingClass handles a REST Update request.
func (service *ContrailService) RESTUpdateForwardingClass(c echo.Context) error {
    id := c.Param("id")
    request := &models.ForwardingClassUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "forwarding_class",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateForwardingClass(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateForwardingClass handles a Update request.
func (service *ContrailService) UpdateForwardingClass(ctx context.Context, request *models.ForwardingClassUpdateRequest) (*models.ForwardingClassUpdateResponse, error) {
    id = request.ID
    model = request.ForwardingClass
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
            return db.UpdateForwardingClass(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "forwarding_class",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ForwardingClass.UpdateResponse{
        ForwardingClass: model,
    }, nil
}

//RESTDeleteForwardingClass delete a resource using REST service.
func (service *ContrailService) RESTDeleteForwardingClass(c echo.Context) error {
    id := c.Param("id")
    request := &models.ForwardingClassDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteForwardingClass(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteForwardingClass delete a resource.
func (service *ContrailService) DeleteForwardingClass(ctx context.Context, request *models.ForwardingClassDeleteRequest) (*models.ForwardingClassDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteForwardingClass(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ForwardingClassDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowForwardingClass a REST Show request.
func (service *ContrailService) RESTShowForwardingClass(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ForwardingClass
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListForwardingClass(tx, &common.ListSpec{
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
        "forwarding_class": result,
    })
}

//RESTListForwardingClass handles a List REST service Request.
func (service *ContrailService) RESTListForwardingClass(c echo.Context) (error) {
    var result []*models.ForwardingClass
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListForwardingClass(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "forwarding-classs": result,
    })
}