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

//RESTCreateContrailCluster handle a Create REST service.
func (service *ContrailService) RESTCreateContrailCluster(c echo.Context) error {
    requestData := &models.ContrailClusterCreateRequest{
        ContrailCluster: models.MakeContrailCluster(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_cluster",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateContrailCluster(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateContrailCluster handle a Create API
func (service *ContrailService) CreateContrailCluster(
    ctx context.Context, 
    request *models.ContrailClusterCreateRequest) (*models.ContrailClusterCreateResponse, error) {
    model := request.ContrailCluster
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
            return db.CreateContrailCluster(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_cluster",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.ContrailClusterCreateResponse{
        ContrailCluster: request.ContrailCluster,
    }, nil
}

//RESTUpdateContrailCluster handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailCluster(c echo.Context) error {
    id := c.Param("id")
    request := &models.ContrailClusterUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "contrail_cluster",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateContrailCluster(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateContrailCluster handles a Update request.
func (service *ContrailService) UpdateContrailCluster(ctx context.Context, request *models.ContrailClusterUpdateRequest) (*models.ContrailClusterUpdateResponse, error) {
    id = request.ID
    model = request.ContrailCluster
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
            return db.UpdateContrailCluster(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "contrail_cluster",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.ContrailCluster.UpdateResponse{
        ContrailCluster: model,
    }, nil
}

//RESTDeleteContrailCluster delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailCluster(c echo.Context) error {
    id := c.Param("id")
    request := &models.ContrailClusterDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteContrailCluster(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailCluster delete a resource.
func (service *ContrailService) DeleteContrailCluster(ctx context.Context, request *models.ContrailClusterDeleteRequest) (*models.ContrailClusterDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteContrailCluster(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.ContrailClusterDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowContrailCluster a REST Show request.
func (service *ContrailService) RESTShowContrailCluster(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.ContrailCluster
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailCluster(tx, &common.ListSpec{
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
        "contrail_cluster": result,
    })
}

//RESTListContrailCluster handles a List REST service Request.
func (service *ContrailService) RESTListContrailCluster(c echo.Context) (error) {
    var result []*models.ContrailCluster
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListContrailCluster(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "contrail-clusters": result,
    })
}