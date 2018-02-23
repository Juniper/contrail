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

//RESTCreateServiceGroup handle a Create REST service.
func (service *ContrailService) RESTCreateServiceGroup(c echo.Context) error {
    requestData := &models.CreateServiceGroupRequest{
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_group",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateServiceGroup(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateServiceGroup handle a Create API
func (service *ContrailService) CreateServiceGroup(
    ctx context.Context, 
    request *models.CreateServiceGroupRequest) (*models.CreateServiceGroupResponse, error) {
    model := request.ServiceGroup
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
            return db.CreateServiceGroup(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_group",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.CreateServiceGroupResponse{
        ServiceGroup: request.ServiceGroup,
    }, nil
}

//RESTUpdateServiceGroup handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceGroup(c echo.Context) error {
    //id := c.Param("id")
    request := &models.UpdateServiceGroupRequest{
    }
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "service_group",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    ctx := c.Request().Context()
    response, err := service.UpdateServiceGroup(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateServiceGroup handles a Update request.
func (service *ContrailService) UpdateServiceGroup(
    ctx context.Context, 
    request *models.UpdateServiceGroupRequest) (*models.UpdateServiceGroupResponse, error) {
    model := request.ServiceGroup
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateServiceGroup(ctx, tx, request)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "service_group",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.UpdateServiceGroupResponse{
        ServiceGroup: model,
    }, nil
}

//RESTDeleteServiceGroup delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceGroup(c echo.Context) error {
    id := c.Param("id")
    request := &models.DeleteServiceGroupRequest{
        ID: id,
    } 
    ctx := c.Request().Context()
    _, err := service.DeleteServiceGroup(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceGroup delete a resource.
func (service *ContrailService) DeleteServiceGroup(ctx context.Context, request *models.DeleteServiceGroupRequest) (*models.DeleteServiceGroupResponse, error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteServiceGroup(ctx, tx, request)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DeleteServiceGroupResponse{
        ID: request.ID,
    }, nil
}

//RESTGetServiceGroup a REST Get request.
func (service *ContrailService) RESTGetServiceGroup(c echo.Context) (error) {
    id := c.Param("id")
    request := &models.GetServiceGroupRequest{
        ID: id,
    } 
    ctx := c.Request().Context() 
    response, err := service.GetServiceGroup(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//GetServiceGroup a Get request.
func (service *ContrailService) GetServiceGroup(ctx context.Context, request *models.GetServiceGroupRequest) (response *models.GetServiceGroupResponse, err error) {
    spec := &models.ListSpec{
                Limit: 1,
                Filters: []*models.Filter{
                    &models.Filter{
                        Key: "uuid",
                        Values: []string{request.ID},
                    },
                },
    }
    listRequest := &models.ListServiceGroupRequest{
        Spec: spec,
    }
    var result *models.ListServiceGroupResponse 
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListServiceGroup(ctx, tx, listRequest)
            return err
        }); err != nil {
        return nil, common.ErrorInternal 
    }
    if len(result.ServiceGroups) == 0 {
        return nil, common.ErrorNotFound
    }
    response = &models.GetServiceGroupResponse{
       ServiceGroup: result.ServiceGroups[0],
    }
    return response, nil
}

//RESTListServiceGroup handles a List REST service Request.
func (service *ContrailService) RESTListServiceGroup(c echo.Context) (error) {
    var err error
    spec := common.GetListSpec(c)
    request := &models.ListServiceGroupRequest{
        Spec: spec,
    }
    ctx := c.Request().Context()
    response, err := service.ListServiceGroup(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//ListServiceGroup handles a List service Request.
func (service *ContrailService) ListServiceGroup(
    ctx context.Context, 
    request *models.ListServiceGroupRequest) (response *models.ListServiceGroupResponse, err error) {
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            response, err = db.ListServiceGroup(ctx, tx, request)
            return err
        }); err != nil {
        return nil, common.ErrorInternal
    }
    return response, nil
}