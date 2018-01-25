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

//RESTCreateVPNGroup handle a Create REST service.
func (service *ContrailService) RESTCreateVPNGroup(c echo.Context) error {
    requestData := &models.VPNGroupCreateRequest{
        VPNGroup: models.MakeVPNGroup(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "vpn_group",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateVPNGroup(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateVPNGroup handle a Create API
func (service *ContrailService) CreateVPNGroup(
    ctx context.Context, 
    request *models.VPNGroupCreateRequest) (*models.VPNGroupCreateResponse, error) {
    model := request.VPNGroup
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
            return db.CreateVPNGroup(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "vpn_group",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.VPNGroupCreateResponse{
        VPNGroup: request.VPNGroup,
    }, nil
}

//RESTUpdateVPNGroup handles a REST Update request.
func (service *ContrailService) RESTUpdateVPNGroup(c echo.Context) error {
    id := c.Param("id")
    request := &models.VPNGroupUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "vpn_group",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateVPNGroup(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateVPNGroup handles a Update request.
func (service *ContrailService) UpdateVPNGroup(ctx context.Context, request *models.VPNGroupUpdateRequest) (*models.VPNGroupUpdateResponse, error) {
    id = request.ID
    model = request.VPNGroup
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
            return db.UpdateVPNGroup(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "vpn_group",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.VPNGroup.UpdateResponse{
        VPNGroup: model,
    }, nil
}

//RESTDeleteVPNGroup delete a resource using REST service.
func (service *ContrailService) RESTDeleteVPNGroup(c echo.Context) error {
    id := c.Param("id")
    request := &models.VPNGroupDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteVPNGroup(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteVPNGroup delete a resource.
func (service *ContrailService) DeleteVPNGroup(ctx context.Context, request *models.VPNGroupDeleteRequest) (*models.VPNGroupDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteVPNGroup(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.VPNGroupDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowVPNGroup a REST Show request.
func (service *ContrailService) RESTShowVPNGroup(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.VPNGroup
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVPNGroup(tx, &common.ListSpec{
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
        "vpn_group": result,
    })
}

//RESTListVPNGroup handles a List REST service Request.
func (service *ContrailService) RESTListVPNGroup(c echo.Context) (error) {
    var result []*models.VPNGroup
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListVPNGroup(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "vpn-groups": result,
    })
}