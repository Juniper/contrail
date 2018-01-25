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

//RESTCreateTagType handle a Create REST service.
func (service *ContrailService) RESTCreateTagType(c echo.Context) error {
    requestData := &models.TagTypeCreateRequest{
        TagType: models.MakeTagType(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "tag_type",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateTagType(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateTagType handle a Create API
func (service *ContrailService) CreateTagType(
    ctx context.Context, 
    request *models.TagTypeCreateRequest) (*models.TagTypeCreateResponse, error) {
    model := request.TagType
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
            return db.CreateTagType(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "tag_type",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.TagTypeCreateResponse{
        TagType: request.TagType,
    }, nil
}

//RESTUpdateTagType handles a REST Update request.
func (service *ContrailService) RESTUpdateTagType(c echo.Context) error {
    id := c.Param("id")
    request := &models.TagTypeUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "tag_type",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateTagType(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateTagType handles a Update request.
func (service *ContrailService) UpdateTagType(ctx context.Context, request *models.TagTypeUpdateRequest) (*models.TagTypeUpdateResponse, error) {
    id = request.ID
    model = request.TagType
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
            return db.UpdateTagType(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "tag_type",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.TagType.UpdateResponse{
        TagType: model,
    }, nil
}

//RESTDeleteTagType delete a resource using REST service.
func (service *ContrailService) RESTDeleteTagType(c echo.Context) error {
    id := c.Param("id")
    request := &models.TagTypeDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteTagType(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteTagType delete a resource.
func (service *ContrailService) DeleteTagType(ctx context.Context, request *models.TagTypeDeleteRequest) (*models.TagTypeDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteTagType(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.TagTypeDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowTagType a REST Show request.
func (service *ContrailService) RESTShowTagType(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.TagType
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListTagType(tx, &common.ListSpec{
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
        "tag_type": result,
    })
}

//RESTListTagType handles a List REST service Request.
func (service *ContrailService) RESTListTagType(c echo.Context) (error) {
    var result []*models.TagType
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListTagType(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "tag-types": result,
    })
}