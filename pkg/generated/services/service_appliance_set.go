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

//RESTServiceApplianceSetUpdateRequest for update request for REST.
type RESTServiceApplianceSetUpdateRequest struct {
    Data map[string]interface{} `json:"service-appliance-set"`
}

//RESTCreateServiceApplianceSet handle a Create REST service.
func (service *ContrailService) RESTCreateServiceApplianceSet(c echo.Context) error {
    requestData := &models.ServiceApplianceSetCreateRequest{
        ServiceApplianceSet: models.MakeServiceApplianceSet(),
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
    request *models.ServiceApplianceSetCreateRequest) (*models.ServiceApplianceSetCreateResponse, error) {
    model := request.ServiceApplianceSet
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
            return db.CreateServiceApplianceSet(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_appliance_set",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ServiceApplianceSetCreateResponse{
        ServiceApplianceSet: request.ServiceApplianceSet,
    }, nil
}

//RESTUpdateServiceApplianceSet handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceApplianceSet(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceApplianceSetUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_appliance_set",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateServiceApplianceSet(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceApplianceSet handles a Update request.
func (service *ContrailService) UpdateServiceApplianceSet(ctx context.Context, request *models.ServiceApplianceSetUpdateRequest) (*models.ServiceApplianceSetUpdateResponse, error) {
    id = request.ID
    model = request.ServiceApplianceSet
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
            return db.UpdateServiceApplianceSet(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_appliance_set",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ServiceApplianceSet.UpdateResponse{
        ServiceApplianceSet: model,
    }, nil
}

//RESTDeleteServiceApplianceSet delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceApplianceSet(c echo.Context) error {
    id := c.Param("id")
    request := &models.ServiceApplianceSetDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteServiceApplianceSet(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceApplianceSet delete a resource.
func (service *ContrailService) DeleteServiceApplianceSet(ctx context.Context, request *models.ServiceApplianceSetDeleteRequest) (*models.ServiceApplianceSetDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceApplianceSet(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ServiceApplianceSetDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowServiceApplianceSet a REST Show request.
func (service *ContrailService) RESTShowServiceApplianceSet(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ServiceApplianceSet
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceApplianceSet(tx, &common.ListSpec{
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
        "service_appliance_set": result,
    })
}

//RESTListServiceApplianceSet handles a List REST service Request.
func (service *ContrailService) RESTListServiceApplianceSet(c echo.Context) (error) {
    var result []*models.ServiceApplianceSet
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceApplianceSet(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "service-appliance-sets": result,
    })
}