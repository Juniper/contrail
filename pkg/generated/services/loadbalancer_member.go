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

//RESTCreateLoadbalancerMember handle a Create REST service.
func (service *ContrailService) RESTCreateLoadbalancerMember(c echo.Context) error {
    requestData := &models.LoadbalancerMemberCreateRequest{
        LoadbalancerMember: models.MakeLoadbalancerMember(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_member",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateLoadbalancerMember(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateLoadbalancerMember handle a Create API
func (service *ContrailService) CreateLoadbalancerMember(
    ctx context.Context, 
    request *models.LoadbalancerMemberCreateRequest) (*models.LoadbalancerMemberCreateResponse, error) {
    model := request.LoadbalancerMember
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
            return db.CreateLoadbalancerMember(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_member",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.LoadbalancerMemberCreateResponse{
        LoadbalancerMember: request.LoadbalancerMember,
    }, nil
}

//RESTUpdateLoadbalancerMember handles a REST Update request.
func (service *ContrailService) RESTUpdateLoadbalancerMember(c echo.Context) error {
    id := c.Param("id")
    request := &models.LoadbalancerMemberUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "loadbalancer_member",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateLoadbalancerMember(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateLoadbalancerMember handles a Update request.
func (service *ContrailService) UpdateLoadbalancerMember(ctx context.Context, request *models.LoadbalancerMemberUpdateRequest) (*models.LoadbalancerMemberUpdateResponse, error) {
    id = request.ID
    model = request.LoadbalancerMember
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
            return db.UpdateLoadbalancerMember(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "loadbalancer_member",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.LoadbalancerMember.UpdateResponse{
        LoadbalancerMember: model,
    }, nil
}

//RESTDeleteLoadbalancerMember delete a resource using REST service.
func (service *ContrailService) RESTDeleteLoadbalancerMember(c echo.Context) error {
    id := c.Param("id")
    request := &models.LoadbalancerMemberDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteLoadbalancerMember(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteLoadbalancerMember delete a resource.
func (service *ContrailService) DeleteLoadbalancerMember(ctx context.Context, request *models.LoadbalancerMemberDeleteRequest) (*models.LoadbalancerMemberDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteLoadbalancerMember(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.LoadbalancerMemberDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowLoadbalancerMember a REST Show request.
func (service *ContrailService) RESTShowLoadbalancerMember(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.LoadbalancerMember
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLoadbalancerMember(tx, &common.ListSpec{
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
        "loadbalancer_member": result,
    })
}

//RESTListLoadbalancerMember handles a List REST service Request.
func (service *ContrailService) RESTListLoadbalancerMember(c echo.Context) (error) {
    var result []*models.LoadbalancerMember
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListLoadbalancerMember(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "loadbalancer-members": result,
    })
}