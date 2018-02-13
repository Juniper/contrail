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

//RESTServiceTemplateUpdateRequest for update request for REST.
type RESTServiceTemplateUpdateRequest struct {
    Data map[string]interface{} `json:"service-template"`
}

//RESTCreateServiceTemplate handle a Create REST service.
func (service *ContrailService) RESTCreateServiceTemplate(c echo.Context) error {
    requestData := &models.ServiceTemplateCreateRequest{
        ServiceTemplate: models.MakeServiceTemplate(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_template",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateServiceTemplate(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateServiceTemplate handle a Create API
func (service *ContrailService) CreateServiceTemplate(
    ctx context.Context, 
    request *models.ServiceTemplateCreateRequest) (*models.ServiceTemplateCreateResponse, error) {
    model := request.ServiceTemplate
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
            return db.CreateServiceTemplate(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_template",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ServiceTemplateCreateResponse{
        ServiceTemplate: request.ServiceTemplate,
    }, nil
}

//RESTUpdateServiceTemplate handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceTemplate(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceTemplateUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_template",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateServiceTemplate(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceTemplate handles a Update request.
func (service *ContrailService) UpdateServiceTemplate(ctx context.Context, request *models.ServiceTemplateUpdateRequest) (*models.ServiceTemplateUpdateResponse, error) {
    id = request.ID
    model = request.ServiceTemplate
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
            return db.UpdateServiceTemplate(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_template",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ServiceTemplate.UpdateResponse{
        ServiceTemplate: model,
    }, nil
}

//RESTDeleteServiceTemplate delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceTemplate(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceTemplateDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteServiceTemplate(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceTemplate delete a resource.
func (service *ContrailService) DeleteServiceTemplate(ctx context.Context, request *models.ServiceTemplateDeleteRequest) (*models.ServiceTemplateDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceTemplate(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ServiceTemplateDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowServiceTemplate a REST Show request.
func (service *ContrailService) RESTShowServiceTemplate(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ServiceTemplate
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceTemplate(tx, &common.ListSpec{
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
        "service_template": result,
    })
}

//RESTListServiceTemplate handles a List REST service Request.
func (service *ContrailService) RESTListServiceTemplate(c echo.Context) (error) {
    var result []*models.ServiceTemplate
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceTemplate(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "service-templates": result,
    })
}