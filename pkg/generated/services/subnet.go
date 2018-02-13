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

//RESTSubnetUpdateRequest for update request for REST.
type RESTSubnetUpdateRequest struct {
    Data map[string]interface{} `json:"subnet"`
}

//RESTCreateSubnet handle a Create REST service.
func (service *ContrailService) RESTCreateSubnet(c echo.Context) error {
    requestData := &models.SubnetCreateRequest{
        Subnet: models.MakeSubnet(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "subnet",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateSubnet(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateSubnet handle a Create API
func (service *ContrailService) CreateSubnet(
    ctx context.Context, 
    request *models.SubnetCreateRequest) (*models.SubnetCreateResponse, error) {
    model := request.Subnet
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
            return db.CreateSubnet(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "subnet",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.SubnetCreateResponse{
        Subnet: request.Subnet,
    }, nil
}

//RESTUpdateSubnet handles a REST Update request.
func (service *ContrailService) RESTUpdateSubnet(c echo.Context) error {
    id := c.Param("id")
    request := &models.SubnetUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "subnet",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateSubnet(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateSubnet handles a Update request.
func (service *ContrailService) UpdateSubnet(ctx context.Context, request *models.SubnetUpdateRequest) (*models.SubnetUpdateResponse, error) {
    id = request.ID
    model = request.Subnet
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
            return db.UpdateSubnet(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "subnet",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.Subnet.UpdateResponse{
        Subnet: model,
    }, nil
}

//RESTDeleteSubnet delete a resource using REST service.
func (service *ContrailService) RESTDeleteSubnet(c echo.Context) error {
    id := c.Param("id")
    request := &models.SubnetDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteSubnet(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteSubnet delete a resource.
func (service *ContrailService) DeleteSubnet(ctx context.Context, request *models.SubnetDeleteRequest) (*models.SubnetDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteSubnet(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.SubnetDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowSubnet a REST Show request.
func (service *ContrailService) RESTShowSubnet(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.Subnet
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListSubnet(tx, &common.ListSpec{
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
        "subnet": result,
    })
}

//RESTListSubnet handles a List REST service Request.
func (service *ContrailService) RESTListSubnet(c echo.Context) (error) {
    var result []*models.Subnet
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListSubnet(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "subnets": result,
    })
}