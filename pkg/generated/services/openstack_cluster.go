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

//RESTOpenstackClusterUpdateRequest for update request for REST.
type RESTOpenstackClusterUpdateRequest struct {
    Data map[string]interface{} `json:"openstack-cluster"`
}

//RESTCreateOpenstackCluster handle a Create REST service.
func (service *ContrailService) RESTCreateOpenstackCluster(c echo.Context) error {
    requestData := &models.OpenstackClusterCreateRequest{
        OpenstackCluster: models.MakeOpenstackCluster(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "openstack_cluster",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateOpenstackCluster(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateOpenstackCluster handle a Create API
func (service *ContrailService) CreateOpenstackCluster(
    ctx context.Context, 
    request *models.OpenstackClusterCreateRequest) (*models.OpenstackClusterCreateResponse, error) {
    model := request.OpenstackCluster
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
            return db.CreateOpenstackCluster(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "openstack_cluster",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.OpenstackClusterCreateResponse{
        OpenstackCluster: request.OpenstackCluster,
    }, nil
}

//RESTUpdateOpenstackCluster handles a REST Update request.
func (service *ContrailService) RESTUpdateOpenstackCluster(c echo.Context) error {
    id := c.Param("id")
    request := &models.OpenstackClusterUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "openstack_cluster",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateOpenstackCluster(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateOpenstackCluster handles a Update request.
func (service *ContrailService) UpdateOpenstackCluster(ctx context.Context, request *models.OpenstackClusterUpdateRequest) (*models.OpenstackClusterUpdateResponse, error) {
    id = request.ID
    model = request.OpenstackCluster
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
            return db.UpdateOpenstackCluster(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "openstack_cluster",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.OpenstackCluster.UpdateResponse{
        OpenstackCluster: model,
    }, nil
}

//RESTDeleteOpenstackCluster delete a resource using REST service.
func (service *ContrailService) RESTDeleteOpenstackCluster(c echo.Context) error {
    id := c.Param("id")
    request := &models.OpenstackClusterDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteOpenstackCluster(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteOpenstackCluster delete a resource.
func (service *ContrailService) DeleteOpenstackCluster(ctx context.Context, request *models.OpenstackClusterDeleteRequest) (*models.OpenstackClusterDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteOpenstackCluster(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.OpenstackClusterDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowOpenstackCluster a REST Show request.
func (service *ContrailService) RESTShowOpenstackCluster(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.OpenstackCluster
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListOpenstackCluster(tx, &common.ListSpec{
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
        "openstack_cluster": result,
    })
}

//RESTListOpenstackCluster handles a List REST service Request.
func (service *ContrailService) RESTListOpenstackCluster(c echo.Context) (error) {
    var result []*models.OpenstackCluster
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListOpenstackCluster(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "openstack-clusters": result,
    })
}