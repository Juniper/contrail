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
    requestData := &models.CustomerAttachmentCreateRequest{
        CustomerAttachment: models.MakeCustomerAttachment(),
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
    request *models.CustomerAttachmentCreateRequest) (*models.CustomerAttachmentCreateResponse, error) {
    model := request.CustomerAttachment
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
            return db.CreateCustomerAttachment(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "customer_attachment",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CustomerAttachmentCreateResponse{
        CustomerAttachment: request.CustomerAttachment,
    }, nil
}

//RESTUpdateCustomerAttachment handles a REST Update request.
func (service *ContrailService) RESTUpdateCustomerAttachment(c echo.Context) error {
    id := c.Param("id")
    request := &models.CustomerAttachmentUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "customer_attachment",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateCustomerAttachment(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateCustomerAttachment handles a Update request.
func (service *ContrailService) UpdateCustomerAttachment(ctx context.Context, request *models.CustomerAttachmentUpdateRequest) (*models.CustomerAttachmentUpdateResponse, error) {
    id = request.ID
    model = request.CustomerAttachment
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
            return db.UpdateCustomerAttachment(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "customer_attachment",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.CustomerAttachment.UpdateResponse{
        CustomerAttachment: model,
    }, nil
}

//RESTDeleteCustomerAttachment delete a resource using REST service.
func (service *ContrailService) RESTDeleteCustomerAttachment(c echo.Context) error {
    id := c.Param("id")
    request := &models.CustomerAttachmentDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteCustomerAttachment(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteCustomerAttachment delete a resource.
func (service *ContrailService) DeleteCustomerAttachment(ctx context.Context, request *models.CustomerAttachmentDeleteRequest) (*models.CustomerAttachmentDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteCustomerAttachment(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.CustomerAttachmentDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowCustomerAttachment a REST Show request.
func (service *ContrailService) RESTShowCustomerAttachment(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.CustomerAttachment
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListCustomerAttachment(tx, &common.ListSpec{
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
        "customer_attachment": result,
    })
}

//RESTListCustomerAttachment handles a List REST service Request.
func (service *ContrailService) RESTListCustomerAttachment(c echo.Context) (error) {
    var result []*models.CustomerAttachment
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListCustomerAttachment(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "customer-attachments": result,
    })
}