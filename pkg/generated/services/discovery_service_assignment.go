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

//RESTDiscoveryServiceAssignmentUpdateRequest for update request for REST.
type RESTDiscoveryServiceAssignmentUpdateRequest struct {
    Data map[string]interface{} `json:"discovery-service-assignment"`
}

//RESTCreateDiscoveryServiceAssignment handle a Create REST service.
func (service *ContrailService) RESTCreateDiscoveryServiceAssignment(c echo.Context) error {
    requestData := &models.DiscoveryServiceAssignmentCreateRequest{
        DiscoveryServiceAssignment: models.MakeDiscoveryServiceAssignment(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "discovery_service_assignment",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateDiscoveryServiceAssignment(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateDiscoveryServiceAssignment handle a Create API
func (service *ContrailService) CreateDiscoveryServiceAssignment(
    ctx context.Context, 
    request *models.DiscoveryServiceAssignmentCreateRequest) (*models.DiscoveryServiceAssignmentCreateResponse, error) {
    model := request.DiscoveryServiceAssignment
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
            return db.CreateDiscoveryServiceAssignment(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "discovery_service_assignment",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.DiscoveryServiceAssignmentCreateResponse{
        DiscoveryServiceAssignment: request.DiscoveryServiceAssignment,
    }, nil
}

//RESTUpdateDiscoveryServiceAssignment handles a REST Update request.
func (service *ContrailService) RESTUpdateDiscoveryServiceAssignment(c echo.Context) error {
    id := c.Param("id")
    request := &models.DiscoveryServiceAssignmentUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "discovery_service_assignment",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateDiscoveryServiceAssignment(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateDiscoveryServiceAssignment handles a Update request.
func (service *ContrailService) UpdateDiscoveryServiceAssignment(ctx context.Context, request *models.DiscoveryServiceAssignmentUpdateRequest) (*models.DiscoveryServiceAssignmentUpdateResponse, error) {
    id = request.ID
    model = request.DiscoveryServiceAssignment
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
            return db.UpdateDiscoveryServiceAssignment(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "discovery_service_assignment",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.DiscoveryServiceAssignment.UpdateResponse{
        DiscoveryServiceAssignment: model,
    }, nil
}

//RESTDeleteDiscoveryServiceAssignment delete a resource using REST service.
func (service *ContrailService) RESTDeleteDiscoveryServiceAssignment(c echo.Context) error {
    id := c.Param("id")
    request := &models.DiscoveryServiceAssignmentDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteDiscoveryServiceAssignment(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteDiscoveryServiceAssignment delete a resource.
func (service *ContrailService) DeleteDiscoveryServiceAssignment(ctx context.Context, request *models.DiscoveryServiceAssignmentDeleteRequest) (*models.DiscoveryServiceAssignmentDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteDiscoveryServiceAssignment(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DiscoveryServiceAssignmentDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowDiscoveryServiceAssignment a REST Show request.
func (service *ContrailService) RESTShowDiscoveryServiceAssignment(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.DiscoveryServiceAssignment
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDiscoveryServiceAssignment(tx, &common.ListSpec{
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
        "discovery_service_assignment": result,
    })
}

//RESTListDiscoveryServiceAssignment handles a List REST service Request.
func (service *ContrailService) RESTListDiscoveryServiceAssignment(c echo.Context) (error) {
    var result []*models.DiscoveryServiceAssignment
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDiscoveryServiceAssignment(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "discovery-service-assignments": result,
    })
}