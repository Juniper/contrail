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

//RESTCreateTag handle a Create REST service.
func (service *ContrailService) RESTCreateTag(c echo.Context) error {
    requestData := &models.TagCreateRequest{
        Tag: models.MakeTag(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "tag",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateTag(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateTag handle a Create API
func (service *ContrailService) CreateTag(
    ctx context.Context, 
    request *models.TagCreateRequest) (*models.TagCreateResponse, error) {
    model := request.Tag
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
            return db.CreateTag(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "tag",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.TagCreateResponse{
        Tag: request.Tag,
    }, nil
}

//RESTUpdateTag handles a REST Update request.
func (service *ContrailService) RESTUpdateTag(c echo.Context) error {
    id := c.Param("id")
    request := &models.TagUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "tag",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateTag(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateTag handles a Update request.
func (service *ContrailService) UpdateTag(ctx context.Context, request *models.TagUpdateRequest) (*models.TagUpdateResponse, error) {
    id = request.ID
    model = request.Tag
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
            return db.UpdateTag(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "tag",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.Tag.UpdateResponse{
        Tag: model,
    }, nil
}

//RESTDeleteTag delete a resource using REST service.
func (service *ContrailService) RESTDeleteTag(c echo.Context) error {
    id := c.Param("id")
    request := &models.TagDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteTag(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteTag delete a resource.
func (service *ContrailService) DeleteTag(ctx context.Context, request *models.TagDeleteRequest) (*models.TagDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteTag(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.TagDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowTag a REST Show request.
func (service *ContrailService) RESTShowTag(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.Tag
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListTag(tx, &common.ListSpec{
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
        "tag": result,
    })
}

//RESTListTag handles a List REST service Request.
func (service *ContrailService) RESTListTag(c echo.Context) (error) {
    var result []*models.Tag
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListTag(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "tags": result,
    })
}