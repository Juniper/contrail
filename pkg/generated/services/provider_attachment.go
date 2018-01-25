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

//RESTCreateProviderAttachment handle a Create REST service.
func (service *ContrailService) RESTCreateProviderAttachment(c echo.Context) error {
    requestData := &models.ProviderAttachmentCreateRequest{
        ProviderAttachment: models.MakeProviderAttachment(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "provider_attachment",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateProviderAttachment(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateProviderAttachment handle a Create API
func (service *ContrailService) CreateProviderAttachment(
    ctx context.Context, 
    request *models.ProviderAttachmentCreateRequest) (*models.ProviderAttachmentCreateResponse, error) {
    model := request.ProviderAttachment
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
            return db.CreateProviderAttachment(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "provider_attachment",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ProviderAttachmentCreateResponse{
        ProviderAttachment: request.ProviderAttachment,
    }, nil
}

//RESTUpdateProviderAttachment handles a REST Update request.
func (service *ContrailService) RESTUpdateProviderAttachment(c echo.Context) error {
    id := c.Param("id")
    request := &models.ProviderAttachmentUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "provider_attachment",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateProviderAttachment(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateProviderAttachment handles a Update request.
func (service *ContrailService) UpdateProviderAttachment(ctx context.Context, request *models.ProviderAttachmentUpdateRequest) (*models.ProviderAttachmentUpdateResponse, error) {
    id = request.ID
    model = request.ProviderAttachment
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
            return db.UpdateProviderAttachment(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "provider_attachment",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ProviderAttachment.UpdateResponse{
        ProviderAttachment: model,
    }, nil
}

//RESTDeleteProviderAttachment delete a resource using REST service.
func (service *ContrailService) RESTDeleteProviderAttachment(c echo.Context) error {
    id := c.Param("id")
    request := &models.ProviderAttachmentDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteProviderAttachment(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteProviderAttachment delete a resource.
func (service *ContrailService) DeleteProviderAttachment(ctx context.Context, request *models.ProviderAttachmentDeleteRequest) (*models.ProviderAttachmentDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteProviderAttachment(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ProviderAttachmentDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowProviderAttachment a REST Show request.
func (service *ContrailService) RESTShowProviderAttachment(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ProviderAttachment
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListProviderAttachment(tx, &common.ListSpec{
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
        "provider_attachment": result,
    })
}

//RESTListProviderAttachment handles a List REST service Request.
func (service *ContrailService) RESTListProviderAttachment(c echo.Context) (error) {
    var result []*models.ProviderAttachment
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListProviderAttachment(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "provider-attachments": result,
    })
}