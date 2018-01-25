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

//RESTCreateDomain handle a Create REST service.
func (service *ContrailService) RESTCreateDomain(c echo.Context) error {
    requestData := &models.DomainCreateRequest{
        Domain: models.MakeDomain(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "domain",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateDomain(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateDomain handle a Create API
func (service *ContrailService) CreateDomain(
    ctx context.Context, 
    request *models.DomainCreateRequest) (*models.DomainCreateResponse, error) {
    model := request.Domain
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
            return db.CreateDomain(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "domain",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.DomainCreateResponse{
        Domain: request.Domain,
    }, nil
}

//RESTUpdateDomain handles a REST Update request.
func (service *ContrailService) RESTUpdateDomain(c echo.Context) error {
    id := c.Param("id")
    request := &models.DomainUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "domain",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateDomain(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateDomain handles a Update request.
func (service *ContrailService) UpdateDomain(ctx context.Context, request *models.DomainUpdateRequest) (*models.DomainUpdateResponse, error) {
    id = request.ID
    model = request.Domain
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
            return db.UpdateDomain(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "domain",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.Domain.UpdateResponse{
        Domain: model,
    }, nil
}

//RESTDeleteDomain delete a resource using REST service.
func (service *ContrailService) RESTDeleteDomain(c echo.Context) error {
    id := c.Param("id")
    request := &models.DomainDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteDomain(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteDomain delete a resource.
func (service *ContrailService) DeleteDomain(ctx context.Context, request *models.DomainDeleteRequest) (*models.DomainDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteDomain(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.DomainDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowDomain a REST Show request.
func (service *ContrailService) RESTShowDomain(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.Domain
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDomain(tx, &common.ListSpec{
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
        "domain": result,
    })
}

//RESTListDomain handles a List REST service Request.
func (service *ContrailService) RESTListDomain(c echo.Context) (error) {
    var result []*models.Domain
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListDomain(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "domains": result,
    })
}