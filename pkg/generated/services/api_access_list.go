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

//RESTCreateAPIAccessList handle a Create REST service.
func (service *ContrailService) RESTCreateAPIAccessList(c echo.Context) error {
    requestData := &models.APIAccessListCreateRequest{
        APIAccessList: models.MakeAPIAccessList(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "api_access_list",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateAPIAccessList(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateAPIAccessList handle a Create API
func (service *ContrailService) CreateAPIAccessList(
    ctx context.Context, 
    request *models.APIAccessListCreateRequest) (*models.APIAccessListCreateResponse, error) {
    model := request.APIAccessList
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
            return db.CreateAPIAccessList(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "api_access_list",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.APIAccessListCreateResponse{
        APIAccessList: request.APIAccessList,
    }, nil
}

//RESTUpdateAPIAccessList handles a REST Update request.
func (service *ContrailService) RESTUpdateAPIAccessList(c echo.Context) error {
    id := c.Param("id")
    request := &models.APIAccessListUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "api_access_list",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateAPIAccessList(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateAPIAccessList handles a Update request.
func (service *ContrailService) UpdateAPIAccessList(ctx context.Context, request *models.APIAccessListUpdateRequest) (*models.APIAccessListUpdateResponse, error) {
    id = request.ID
    model = request.APIAccessList
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
            return db.UpdateAPIAccessList(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "api_access_list",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.APIAccessList.UpdateResponse{
        APIAccessList: model,
    }, nil
}

//RESTDeleteAPIAccessList delete a resource using REST service.
func (service *ContrailService) RESTDeleteAPIAccessList(c echo.Context) error {
    id := c.Param("id")
    request := &models.APIAccessListDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteAPIAccessList(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteAPIAccessList delete a resource.
func (service *ContrailService) DeleteAPIAccessList(ctx context.Context, request *models.APIAccessListDeleteRequest) (*models.APIAccessListDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteAPIAccessList(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.APIAccessListDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowAPIAccessList a REST Show request.
func (service *ContrailService) RESTShowAPIAccessList(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.APIAccessList
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAPIAccessList(tx, &common.ListSpec{
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
        "api_access_list": result,
    })
}

//RESTListAPIAccessList handles a List REST service Request.
func (service *ContrailService) RESTListAPIAccessList(c echo.Context) (error) {
    var result []*models.APIAccessList
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAPIAccessList(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "api-access-lists": result,
    })
}