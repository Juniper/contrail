package services

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
)

//RESTBridgeDomainUpdateRequest for update request for REST.
type RESTBridgeDomainUpdateRequest struct {
	Data map[string]interface{} `json:"bridge-domain"`
}

//RESTCreateBridgeDomain handle a Create REST service.
func (service *ContrailService) RESTCreateBridgeDomain(c echo.Context) error {
	requestData := &models.CreateBridgeDomainRequest{
		BridgeDomain: models.MakeBridgeDomain(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bridge_domain",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBridgeDomain(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBridgeDomain handle a Create API
func (service *ContrailService) CreateBridgeDomain(
	ctx context.Context,
	request *models.CreateBridgeDomainRequest) (*models.CreateBridgeDomainResponse, error) {
	model := request.BridgeDomain
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}
	auth := common.GetAuthCTX(ctx)
	if auth == nil {
		return nil, common.ErrorUnauthenticated
	}

	if model.FQName == nil {
		if model.DisplayName == "" {
			return nil, common.ErrorBadRequest("Both of FQName and Display Name is empty")
		}
		model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateBridgeDomain(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bridge_domain",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBridgeDomainResponse{
		BridgeDomain: request.BridgeDomain,
	}, nil
}

//RESTUpdateBridgeDomain handles a REST Update request.
func (service *ContrailService) RESTUpdateBridgeDomain(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBridgeDomainRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bridge_domain",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBridgeDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBridgeDomain handles a Update request.
func (service *ContrailService) UpdateBridgeDomain(
	ctx context.Context,
	request *models.UpdateBridgeDomainRequest) (*models.UpdateBridgeDomainResponse, error) {
	model := request.BridgeDomain
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateBridgeDomain(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bridge_domain",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBridgeDomainResponse{
		BridgeDomain: model,
	}, nil
}

//RESTDeleteBridgeDomain delete a resource using REST service.
func (service *ContrailService) RESTDeleteBridgeDomain(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBridgeDomainRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBridgeDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteBridgeDomain delete a resource.
func (service *ContrailService) DeleteBridgeDomain(ctx context.Context, request *models.DeleteBridgeDomainRequest) (*models.DeleteBridgeDomainResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteBridgeDomain(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBridgeDomainResponse{
		ID: request.ID,
	}, nil
}

//RESTGetBridgeDomain a REST Get request.
func (service *ContrailService) RESTGetBridgeDomain(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBridgeDomainRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBridgeDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetBridgeDomain a Get request.
func (service *ContrailService) GetBridgeDomain(ctx context.Context, request *models.GetBridgeDomainRequest) (response *models.GetBridgeDomainResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListBridgeDomainRequest{
		Spec: spec,
	}
	var result *models.ListBridgeDomainResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListBridgeDomain(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BridgeDomains) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBridgeDomainResponse{
		BridgeDomain: result.BridgeDomains[0],
	}
	return response, nil
}

//RESTListBridgeDomain handles a List REST service Request.
func (service *ContrailService) RESTListBridgeDomain(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListBridgeDomainRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBridgeDomain(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListBridgeDomain handles a List service Request.
func (service *ContrailService) ListBridgeDomain(
	ctx context.Context,
	request *models.ListBridgeDomainRequest) (response *models.ListBridgeDomainResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListBridgeDomain(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
