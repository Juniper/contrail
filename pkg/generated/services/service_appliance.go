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

//RESTCreateServiceAppliance handle a Create REST service.
func (service *ContrailService) RESTCreateServiceAppliance(c echo.Context) error {
    requestData := &models.ServiceApplianceCreateRequest{
        ServiceAppliance: models.MakeServiceAppliance(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_appliance",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateServiceAppliance(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateServiceAppliance handle a Create API
func (service *ContrailService) CreateServiceAppliance(
    ctx context.Context, 
    request *models.ServiceApplianceCreateRequest) (*models.ServiceApplianceCreateResponse, error) {
    model := request.ServiceAppliance
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
            return db.CreateServiceAppliance(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_appliance",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ServiceApplianceCreateResponse{
        ServiceAppliance: request.ServiceAppliance,
    }, nil
}

//RESTUpdateServiceAppliance handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceAppliance(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceApplianceUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_appliance",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateServiceAppliance(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceAppliance handles a Update request.
func (service *ContrailService) UpdateServiceAppliance(ctx context.Context, request *models.ServiceApplianceUpdateRequest) (*models.ServiceApplianceUpdateResponse, error) {
    id = request.ID
    model = request.ServiceAppliance
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
            return db.UpdateServiceAppliance(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_appliance",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ServiceAppliance.UpdateResponse{
        ServiceAppliance: model,
    }, nil
}

//RESTDeleteServiceAppliance delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceAppliance(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceApplianceDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteServiceAppliance(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceAppliance delete a resource.
func (service *ContrailService) DeleteServiceAppliance(ctx context.Context, request *models.ServiceApplianceDeleteRequest) (*models.ServiceApplianceDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceAppliance(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ServiceApplianceDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowServiceAppliance a REST Show request.
func (service *ContrailService) RESTShowServiceAppliance(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ServiceAppliance
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceAppliance(tx, &common.ListSpec{
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
        "service_appliance": result,
    })
}

//RESTListServiceAppliance handles a List REST service Request.
func (service *ContrailService) RESTListServiceAppliance(c echo.Context) (error) {
    var result []*models.ServiceAppliance
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceAppliance(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "service-appliances": result,
    })
}