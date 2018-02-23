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

//RESTCreateServiceApplianceSet handle a Create REST service.
func (service *ContrailService) RESTCreateServiceApplianceSet(c echo.Context) error {
    requestData := &models.CreateServiceApplianceSetRequest{
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_appliance_set",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateServiceApplianceSet(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateServiceApplianceSet handle a Create API
func (service *ContrailService) CreateServiceApplianceSet(
    ctx context.Context, 
    request *models.CreateServiceApplianceSetRequest) (*models.CreateServiceApplianceSetResponse, error) {
    model := request.ServiceApplianceSet
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
            return db.CreateServiceApplianceSet(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_appliance_set",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateServiceApplianceSetResponse{
        ServiceApplianceSet: request.ServiceApplianceSet,
    }, nil
}

//RESTUpdateServiceApplianceSet handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceApplianceSet(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateServiceApplianceSetRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_appliance_set",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateServiceApplianceSet(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceApplianceSet handles a Update request.
func (service *ContrailService) UpdateServiceApplianceSet(
    ctx context.Context, 
    request *models.UpdateServiceApplianceSetRequest) (*models.UpdateServiceApplianceSetResponse, error) {
    model := request.ServiceApplianceSet
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateServiceApplianceSet(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_appliance_set",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateServiceApplianceSetResponse{
        ServiceApplianceSet: model,
    }, nil
}

//RESTDeleteServiceApplianceSet delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceApplianceSet(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteServiceApplianceSetRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteServiceApplianceSet(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceApplianceSet delete a resource.
func (service *ContrailService) DeleteServiceApplianceSet(ctx context.Context, request *models.DeleteServiceApplianceSetRequest) (*models.DeleteServiceApplianceSetResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceApplianceSet(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteServiceApplianceSetResponse{
        ID: request.ID,
    }, nil
}

//RESTGetServiceApplianceSet a REST Get request.
func (service *ContrailService) RESTGetServiceApplianceSet(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetServiceApplianceSetRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetServiceApplianceSet(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetServiceApplianceSet a Get request.
func (service *ContrailService) GetServiceApplianceSet(ctx context.Context, request *models.GetServiceApplianceSetRequest) (response *models.GetServiceApplianceSetResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListServiceApplianceSetRequest{
        Spec: spec,
    }
    var result *models.ListServiceApplianceSetResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceApplianceSet(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.ServiceApplianceSets) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetServiceApplianceSetResponse{
       ServiceApplianceSet: result.ServiceApplianceSets[0],
    }
    return response, nil
}

//RESTListServiceApplianceSet handles a List REST service Request.
func (service *ContrailService) RESTListServiceApplianceSet(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListServiceApplianceSetRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListServiceApplianceSet(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListServiceApplianceSet handles a List service Request.
func (service *ContrailService) ListServiceApplianceSet(
    ctx context.Context, 
    request *models.ListServiceApplianceSetRequest) (response *models.ListServiceApplianceSetResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListServiceApplianceSet(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}