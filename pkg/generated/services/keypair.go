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

//RESTCreateKeypair handle a Create REST service.
func (service *ContrailService) RESTCreateKeypair(c echo.Context) error {
    requestData := &models.KeypairCreateRequest{
        Keypair: models.MakeKeypair(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "keypair",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateKeypair(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateKeypair handle a Create API
func (service *ContrailService) CreateKeypair(
    ctx context.Context, 
    request *models.KeypairCreateRequest) (*models.KeypairCreateResponse, error) {
    model := request.Keypair
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
            return db.CreateKeypair(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "keypair",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.KeypairCreateResponse{
        Keypair: request.Keypair,
    }, nil
}

//RESTUpdateKeypair handles a REST Update request.
func (service *ContrailService) RESTUpdateKeypair(c echo.Context) error {
    id := c.Param("id")
    request := &models.KeypairUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "keypair",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateKeypair(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateKeypair handles a Update request.
func (service *ContrailService) UpdateKeypair(ctx context.Context, request *models.KeypairUpdateRequest) (*models.KeypairUpdateResponse, error) {
    id = request.ID
    model = request.Keypair
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
            return db.UpdateKeypair(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "keypair",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.Keypair.UpdateResponse{
        Keypair: model,
    }, nil
}

//RESTDeleteKeypair delete a resource using REST service.
func (service *ContrailService) RESTDeleteKeypair(c echo.Context) error {
    id := c.Param("id")
    request := &models.KeypairDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteKeypair(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteKeypair delete a resource.
func (service *ContrailService) DeleteKeypair(ctx context.Context, request *models.KeypairDeleteRequest) (*models.KeypairDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteKeypair(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.KeypairDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowKeypair a REST Show request.
func (service *ContrailService) RESTShowKeypair(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.Keypair
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListKeypair(tx, &common.ListSpec{
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
        "keypair": result,
    })
}

//RESTListKeypair handles a List REST service Request.
func (service *ContrailService) RESTListKeypair(c echo.Context) (error) {
    var result []*models.Keypair
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListKeypair(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "keypairs": result,
    })
}