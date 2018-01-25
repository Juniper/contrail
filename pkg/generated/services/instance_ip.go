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

//RESTCreateInstanceIP handle a Create REST service.
func (service *ContrailService) RESTCreateInstanceIP(c echo.Context) error {
    requestData := &models.InstanceIPCreateRequest{
        InstanceIP: models.MakeInstanceIP(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "instance_ip",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateInstanceIP(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateInstanceIP handle a Create API
func (service *ContrailService) CreateInstanceIP(
    ctx context.Context, 
    request *models.InstanceIPCreateRequest) (*models.InstanceIPCreateResponse, error) {
    model := request.InstanceIP
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
            return db.CreateInstanceIP(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "instance_ip",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.InstanceIPCreateResponse{
        InstanceIP: request.InstanceIP,
    }, nil
}

//RESTUpdateInstanceIP handles a REST Update request.
func (service *ContrailService) RESTUpdateInstanceIP(c echo.Context) error {
    id := c.Param("id")
    request := &models.InstanceIPUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "instance_ip",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateInstanceIP(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateInstanceIP handles a Update request.
func (service *ContrailService) UpdateInstanceIP(ctx context.Context, request *models.InstanceIPUpdateRequest) (*models.InstanceIPUpdateResponse, error) {
    id = request.ID
    model = request.InstanceIP
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
            return db.UpdateInstanceIP(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "instance_ip",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.InstanceIP.UpdateResponse{
        InstanceIP: model,
    }, nil
}

//RESTDeleteInstanceIP delete a resource using REST service.
func (service *ContrailService) RESTDeleteInstanceIP(c echo.Context) error {
    id := c.Param("id")
    request := &models.InstanceIPDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteInstanceIP(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteInstanceIP delete a resource.
func (service *ContrailService) DeleteInstanceIP(ctx context.Context, request *models.InstanceIPDeleteRequest) (*models.InstanceIPDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteInstanceIP(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.InstanceIPDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowInstanceIP a REST Show request.
func (service *ContrailService) RESTShowInstanceIP(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.InstanceIP
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListInstanceIP(tx, &common.ListSpec{
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
        "instance_ip": result,
    })
}

//RESTListInstanceIP handles a List REST service Request.
func (service *ContrailService) RESTListInstanceIP(c echo.Context) (error) {
    var result []*models.InstanceIP
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListInstanceIP(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "instance-ips": result,
    })
}