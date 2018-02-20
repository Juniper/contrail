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
    requestData := &models.CreateForwardingClassRequest{
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
    request *models.CreateForwardingClassRequest) (*models.CreateForwardingClassResponse, error) {
    model := request.ForwardingClass
    if model.UUID == "" {
        model.UUID = uuid.NewV4().String()
    }
    auth := common.GetAuthCTX(ctx)
    if auth == nil {
        return nil, common.ErrorUnauthenticated
    }

    if model.FQName == nil {
        if model.DisplayName == "" {
        return nil, common.ErrorBadRequest("Both of FQName and Display Name is empty") 
        }
        model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
    }
    model.Perms2 = &models.PermType2{}
    model.Perms2.Owner = auth.ProjectID()
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.CreateForwardingClass(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "forwarding_class",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateForwardingClassResponse{
        ForwardingClass: request.ForwardingClass,
    }, nil
}

//RESTUpdateForwardingClass handles a REST Update request.
func (service *ContrailService) RESTUpdateForwardingClass(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateForwardingClassRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "forwarding_class",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateForwardingClass(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateForwardingClass handles a Update request.
func (service *ContrailService) UpdateForwardingClass(
    ctx context.Context, 
    request *models.UpdateForwardingClassRequest) (*models.UpdateForwardingClassResponse, error) {
    model := request.ForwardingClass
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateForwardingClass(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "forwarding_class",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateForwardingClassResponse{
        ForwardingClass: model,
    }, nil
}

//RESTDeleteForwardingClass delete a resource using REST service.
func (service *ContrailService) RESTDeleteForwardingClass(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteForwardingClassRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteForwardingClass(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteForwardingClass delete a resource.
func (service *ContrailService) DeleteForwardingClass(ctx context.Context, request *models.DeleteForwardingClassRequest) (*models.DeleteForwardingClassResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteForwardingClass(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteForwardingClassResponse{
        ID: request.ID,
    }, nil
}

//RESTGetForwardingClass a REST Get request.
func (service *ContrailService) RESTGetForwardingClass(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetForwardingClassRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetForwardingClass(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetForwardingClass a Get request.
func (service *ContrailService) GetForwardingClass(ctx context.Context, request *models.GetForwardingClassRequest) (response *models.GetForwardingClassResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListForwardingClassRequest{
        Spec: spec,
    }
    var result *models.ListForwardingClassResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListForwardingClass(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.ForwardingClasss) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetForwardingClassResponse{
       ForwardingClass: result.ForwardingClasss[0],
    }
    return response, nil
}

//RESTListForwardingClass handles a List REST service Request.
func (service *ContrailService) RESTListForwardingClass(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListForwardingClassRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListForwardingClass(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListForwardingClass handles a List service Request.
func (service *ContrailService) ListForwardingClass(
    ctx context.Context, 
    request *models.ListForwardingClassRequest) (response *models.ListForwardingClassResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListForwardingClass(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}