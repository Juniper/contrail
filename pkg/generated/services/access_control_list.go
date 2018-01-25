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

//RESTCreateAccessControlList handle a Create REST service.
func (service *ContrailService) RESTCreateAccessControlList(c echo.Context) error {
    requestData := &models.AccessControlListCreateRequest{
        AccessControlList: models.MakeAccessControlList(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "access_control_list",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateAccessControlList(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateAccessControlList handle a Create API
func (service *ContrailService) CreateAccessControlList(
    ctx context.Context, 
    request *models.AccessControlListCreateRequest) (*models.AccessControlListCreateResponse, error) {
    model := request.AccessControlList
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
            return db.CreateAccessControlList(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "access_control_list",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.AccessControlListCreateResponse{
        AccessControlList: request.AccessControlList,
    }, nil
}

//RESTUpdateAccessControlList handles a REST Update request.
func (service *ContrailService) RESTUpdateAccessControlList(c echo.Context) error {
    id := c.Param("id")
    request := &models.AccessControlListUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "access_control_list",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateAccessControlList(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateAccessControlList handles a Update request.
func (service *ContrailService) UpdateAccessControlList(ctx context.Context, request *models.AccessControlListUpdateRequest) (*models.AccessControlListUpdateResponse, error) {
    id = request.ID
    model = request.AccessControlList
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
            return db.UpdateAccessControlList(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "access_control_list",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.AccessControlList.UpdateResponse{
        AccessControlList: model,
    }, nil
}

//RESTDeleteAccessControlList delete a resource using REST service.
func (service *ContrailService) RESTDeleteAccessControlList(c echo.Context) error {
    id := c.Param("id")
    request := &models.AccessControlListDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteAccessControlList(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteAccessControlList delete a resource.
func (service *ContrailService) DeleteAccessControlList(ctx context.Context, request *models.AccessControlListDeleteRequest) (*models.AccessControlListDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteAccessControlList(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.AccessControlListDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowAccessControlList a REST Show request.
func (service *ContrailService) RESTShowAccessControlList(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.AccessControlList
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAccessControlList(tx, &common.ListSpec{
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
        "access_control_list": result,
    })
}

//RESTListAccessControlList handles a List REST service Request.
func (service *ContrailService) RESTListAccessControlList(c echo.Context) (error) {
    var result []*models.AccessControlList
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAccessControlList(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "access-control-lists": result,
    })
}