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

//RESTCreateAddressGroup handle a Create REST service.
func (service *ContrailService) RESTCreateAddressGroup(c echo.Context) error {
    requestData := &models.AddressGroupCreateRequest{
        AddressGroup: models.MakeAddressGroup(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "address_group",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateAddressGroup(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateAddressGroup handle a Create API
func (service *ContrailService) CreateAddressGroup(
    ctx context.Context, 
    request *models.AddressGroupCreateRequest) (*models.AddressGroupCreateResponse, error) {
    model := request.AddressGroup
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
            return db.CreateAddressGroup(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "address_group",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.AddressGroupCreateResponse{
        AddressGroup: request.AddressGroup,
    }, nil
}

//RESTUpdateAddressGroup handles a REST Update request.
func (service *ContrailService) RESTUpdateAddressGroup(c echo.Context) error {
    id := c.Param("id")
    request := &models.AddressGroupUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "address_group",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateAddressGroup(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateAddressGroup handles a Update request.
func (service *ContrailService) UpdateAddressGroup(ctx context.Context, request *models.AddressGroupUpdateRequest) (*models.AddressGroupUpdateResponse, error) {
    id = request.ID
    model = request.AddressGroup
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
            return db.UpdateAddressGroup(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "address_group",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.AddressGroup.UpdateResponse{
        AddressGroup: model,
    }, nil
}

//RESTDeleteAddressGroup delete a resource using REST service.
func (service *ContrailService) RESTDeleteAddressGroup(c echo.Context) error {
    id := c.Param("id")
    request := &models.AddressGroupDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteAddressGroup(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteAddressGroup delete a resource.
func (service *ContrailService) DeleteAddressGroup(ctx context.Context, request *models.AddressGroupDeleteRequest) (*models.AddressGroupDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteAddressGroup(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.AddressGroupDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowAddressGroup a REST Show request.
func (service *ContrailService) RESTShowAddressGroup(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.AddressGroup
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAddressGroup(tx, &common.ListSpec{
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
        "address_group": result,
    })
}

//RESTListAddressGroup handles a List REST service Request.
func (service *ContrailService) RESTListAddressGroup(c echo.Context) (error) {
    var result []*models.AddressGroup
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAddressGroup(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "address-groups": result,
    })
}