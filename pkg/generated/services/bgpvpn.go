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

//RESTCreateBGPVPN handle a Create REST service.
func (service *ContrailService) RESTCreateBGPVPN(c echo.Context) error {
    requestData := &models.BGPVPNCreateRequest{
        BGPVPN: models.MakeBGPVPN(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bgpvpn",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateBGPVPN(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateBGPVPN handle a Create API
func (service *ContrailService) CreateBGPVPN(
    ctx context.Context, 
    request *models.BGPVPNCreateRequest) (*models.BGPVPNCreateResponse, error) {
    model := request.BGPVPN
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
            return db.CreateBGPVPN(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bgpvpn",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.BGPVPNCreateResponse{
        BGPVPN: request.BGPVPN,
    }, nil
}

//RESTUpdateBGPVPN handles a REST Update request.
func (service *ContrailService) RESTUpdateBGPVPN(c echo.Context) error {
    id := c.Param("id")
    request := &models.BGPVPNUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "bgpvpn",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateBGPVPN(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateBGPVPN handles a Update request.
func (service *ContrailService) UpdateBGPVPN(ctx context.Context, request *models.BGPVPNUpdateRequest) (*models.BGPVPNUpdateResponse, error) {
    id = request.ID
    model = request.BGPVPN
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
            return db.UpdateBGPVPN(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bgpvpn",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.BGPVPN.UpdateResponse{
        BGPVPN: model,
    }, nil
}

//RESTDeleteBGPVPN delete a resource using REST service.
func (service *ContrailService) RESTDeleteBGPVPN(c echo.Context) error {
    id := c.Param("id")
    request := &models.BGPVPNDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteBGPVPN(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteBGPVPN delete a resource.
func (service *ContrailService) DeleteBGPVPN(ctx context.Context, request *models.BGPVPNDeleteRequest) (*models.BGPVPNDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteBGPVPN(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.BGPVPNDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowBGPVPN a REST Show request.
func (service *ContrailService) RESTShowBGPVPN(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.BGPVPN
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBGPVPN(tx, &common.ListSpec{
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
        "bgpvpn": result,
    })
}

//RESTListBGPVPN handles a List REST service Request.
func (service *ContrailService) RESTListBGPVPN(c echo.Context) (error) {
    var result []*models.BGPVPN
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBGPVPN(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "bgpvpns": result,
    })
}