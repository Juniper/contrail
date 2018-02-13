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

//RESTBGPAsAServiceUpdateRequest for update request for REST.
type RESTBGPAsAServiceUpdateRequest struct {
    Data map[string]interface{} `json:"bgp-as-a-service"`
}

//RESTCreateBGPAsAService handle a Create REST service.
func (service *ContrailService) RESTCreateBGPAsAService(c echo.Context) error {
    requestData := &models.BGPAsAServiceCreateRequest{
        BGPAsAService: models.MakeBGPAsAService(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bgp_as_a_service",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateBGPAsAService(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateBGPAsAService handle a Create API
func (service *ContrailService) CreateBGPAsAService(
    ctx context.Context, 
    request *models.BGPAsAServiceCreateRequest) (*models.BGPAsAServiceCreateResponse, error) {
    model := request.BGPAsAService
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
            return db.CreateBGPAsAService(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bgp_as_a_service",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.BGPAsAServiceCreateResponse{
        BGPAsAService: request.BGPAsAService,
    }, nil
}

//RESTUpdateBGPAsAService handles a REST Update request.
func (service *ContrailService) RESTUpdateBGPAsAService(c echo.Context) error {
    id := c.Param("id")
    request := &models.BGPAsAServiceUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "bgp_as_a_service",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateBGPAsAService(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateBGPAsAService handles a Update request.
func (service *ContrailService) UpdateBGPAsAService(ctx context.Context, request *models.BGPAsAServiceUpdateRequest) (*models.BGPAsAServiceUpdateResponse, error) {
    id = request.ID
    model = request.BGPAsAService
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
            return db.UpdateBGPAsAService(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "bgp_as_a_service",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.BGPAsAService.UpdateResponse{
        BGPAsAService: model,
    }, nil
}

//RESTDeleteBGPAsAService delete a resource using REST service.
func (service *ContrailService) RESTDeleteBGPAsAService(c echo.Context) error {
    id := c.Param("id")
    request := &models.BGPAsAServiceDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteBGPAsAService(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteBGPAsAService delete a resource.
func (service *ContrailService) DeleteBGPAsAService(ctx context.Context, request *models.BGPAsAServiceDeleteRequest) (*models.BGPAsAServiceDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteBGPAsAService(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.BGPAsAServiceDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowBGPAsAService a REST Show request.
func (service *ContrailService) RESTShowBGPAsAService(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.BGPAsAService
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBGPAsAService(tx, &common.ListSpec{
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
        "bgp_as_a_service": result,
    })
}

//RESTListBGPAsAService handles a List REST service Request.
func (service *ContrailService) RESTListBGPAsAService(c echo.Context) (error) {
    var result []*models.BGPAsAService
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListBGPAsAService(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "bgp-as-a-services": result,
    })
}