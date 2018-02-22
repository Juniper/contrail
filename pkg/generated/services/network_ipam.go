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

//RESTNetworkIpamUpdateRequest for update request for REST.
type RESTNetworkIpamUpdateRequest struct {
	Data map[string]interface{} `json:"network-ipam"`
}

//RESTCreateNetworkIpam handle a Create REST service.
func (service *ContrailService) RESTCreateNetworkIpam(c echo.Context) error {
	requestData := &models.CreateNetworkIpamRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_ipam",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateNetworkIpam(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateNetworkIpam handle a Create API
func (service *ContrailService) CreateNetworkIpam(
	ctx context.Context,
	request *models.CreateNetworkIpamRequest) (*models.CreateNetworkIpamResponse, error) {
	model := request.NetworkIpam
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
			return db.CreateNetworkIpam(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_ipam",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateNetworkIpamResponse{
		NetworkIpam: request.NetworkIpam,
	}, nil
}

//RESTUpdateNetworkIpam handles a REST Update request.
func (service *ContrailService) RESTUpdateNetworkIpam(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateNetworkIpamRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_ipam",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateNetworkIpam(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateNetworkIpam handles a Update request.
func (service *ContrailService) UpdateNetworkIpam(
	ctx context.Context,
	request *models.UpdateNetworkIpamRequest) (*models.UpdateNetworkIpamResponse, error) {
	model := request.NetworkIpam
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateNetworkIpam(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_ipam",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateNetworkIpamResponse{
		NetworkIpam: model,
	}, nil
}

//RESTDeleteNetworkIpam delete a resource using REST service.
func (service *ContrailService) RESTDeleteNetworkIpam(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteNetworkIpamRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteNetworkIpam(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteNetworkIpam delete a resource.
func (service *ContrailService) DeleteNetworkIpam(ctx context.Context, request *models.DeleteNetworkIpamRequest) (*models.DeleteNetworkIpamResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteNetworkIpam(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteNetworkIpamResponse{
		ID: request.ID,
	}, nil
}

//RESTGetNetworkIpam a REST Get request.
func (service *ContrailService) RESTGetNetworkIpam(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetNetworkIpamRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetNetworkIpam(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetNetworkIpam a Get request.
func (service *ContrailService) GetNetworkIpam(ctx context.Context, request *models.GetNetworkIpamRequest) (response *models.GetNetworkIpamResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListNetworkIpamRequest{
		Spec: spec,
	}
	var result *models.ListNetworkIpamResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListNetworkIpam(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.NetworkIpams) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetNetworkIpamResponse{
		NetworkIpam: result.NetworkIpams[0],
	}
	return response, nil
}

//RESTListNetworkIpam handles a List REST service Request.
func (service *ContrailService) RESTListNetworkIpam(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListNetworkIpamRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListNetworkIpam(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListNetworkIpam handles a List service Request.
func (service *ContrailService) ListNetworkIpam(
	ctx context.Context,
	request *models.ListNetworkIpamRequest) (response *models.ListNetworkIpamResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListNetworkIpam(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
