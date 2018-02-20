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

//RESTCreateCustomerAttachment handle a Create REST service.
func (service *ContrailService) RESTCreateCustomerAttachment(c echo.Context) error {
    requestData := &models.CreateCustomerAttachmentRequest{
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "customer_attachment",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateCustomerAttachment(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateCustomerAttachment handle a Create API
func (service *ContrailService) CreateCustomerAttachment(
    ctx context.Context, 
    request *models.CreateCustomerAttachmentRequest) (*models.CreateCustomerAttachmentResponse, error) {
    model := request.CustomerAttachment
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
            return db.CreateCustomerAttachment(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "customer_attachment",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateCustomerAttachmentResponse{
        CustomerAttachment: request.CustomerAttachment,
    }, nil
}

//RESTUpdateCustomerAttachment handles a REST Update request.
func (service *ContrailService) RESTUpdateCustomerAttachment(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateCustomerAttachmentRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "customer_attachment",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateCustomerAttachment(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateCustomerAttachment handles a Update request.
func (service *ContrailService) UpdateCustomerAttachment(
    ctx context.Context, 
    request *models.UpdateCustomerAttachmentRequest) (*models.UpdateCustomerAttachmentResponse, error) {
    model := request.CustomerAttachment
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateCustomerAttachment(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "customer_attachment",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateCustomerAttachmentResponse{
        CustomerAttachment: model,
    }, nil
}

//RESTDeleteCustomerAttachment delete a resource using REST service.
func (service *ContrailService) RESTDeleteCustomerAttachment(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteCustomerAttachmentRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteCustomerAttachment(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteCustomerAttachment delete a resource.
func (service *ContrailService) DeleteCustomerAttachment(ctx context.Context, request *models.DeleteCustomerAttachmentRequest) (*models.DeleteCustomerAttachmentResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteCustomerAttachment(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteCustomerAttachmentResponse{
        ID: request.ID,
    }, nil
}

//RESTGetCustomerAttachment a REST Get request.
func (service *ContrailService) RESTGetCustomerAttachment(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetCustomerAttachmentRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetCustomerAttachment(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetCustomerAttachment a Get request.
func (service *ContrailService) GetCustomerAttachment(ctx context.Context, request *models.GetCustomerAttachmentRequest) (response *models.GetCustomerAttachmentResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListCustomerAttachmentRequest{
        Spec: spec,
    }
    var result *models.ListCustomerAttachmentResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListCustomerAttachment(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.CustomerAttachments) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetCustomerAttachmentResponse{
       CustomerAttachment: result.CustomerAttachments[0],
    }
    return response, nil
}

//RESTListCustomerAttachment handles a List REST service Request.
func (service *ContrailService) RESTListCustomerAttachment(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListCustomerAttachmentRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListCustomerAttachment(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListCustomerAttachment handles a List service Request.
func (service *ContrailService) ListCustomerAttachment(
    ctx context.Context, 
    request *models.ListCustomerAttachmentRequest) (response *models.ListCustomerAttachmentResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListCustomerAttachment(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}