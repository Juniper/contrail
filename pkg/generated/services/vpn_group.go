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

//RESTVPNGroupUpdateRequest for update request for REST.
type RESTVPNGroupUpdateRequest struct {
	Data map[string]interface{} `json:"vpn-group"`
}

//RESTCreateVPNGroup handle a Create REST service.
func (service *ContrailService) RESTCreateVPNGroup(c echo.Context) error {
	requestData := &models.CreateVPNGroupRequest{
		VPNGroup: models.MakeVPNGroup(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "vpn_group",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVPNGroup(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVPNGroup handle a Create API
func (service *ContrailService) CreateVPNGroup(
	ctx context.Context,
	request *models.CreateVPNGroupRequest) (*models.CreateVPNGroupResponse, error) {
	model := request.VPNGroup
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
			return db.CreateVPNGroup(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "vpn_group",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVPNGroupResponse{
		VPNGroup: request.VPNGroup,
	}, nil
}

//RESTUpdateVPNGroup handles a REST Update request.
func (service *ContrailService) RESTUpdateVPNGroup(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVPNGroupRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "vpn_group",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVPNGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVPNGroup handles a Update request.
func (service *ContrailService) UpdateVPNGroup(
	ctx context.Context,
	request *models.UpdateVPNGroupRequest) (*models.UpdateVPNGroupResponse, error) {
	model := request.VPNGroup
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateVPNGroup(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "vpn_group",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVPNGroupResponse{
		VPNGroup: model,
	}, nil
}

//RESTDeleteVPNGroup delete a resource using REST service.
func (service *ContrailService) RESTDeleteVPNGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVPNGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVPNGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteVPNGroup delete a resource.
func (service *ContrailService) DeleteVPNGroup(ctx context.Context, request *models.DeleteVPNGroupRequest) (*models.DeleteVPNGroupResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteVPNGroup(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVPNGroupResponse{
		ID: request.ID,
	}, nil
}

//RESTGetVPNGroup a REST Get request.
func (service *ContrailService) RESTGetVPNGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVPNGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVPNGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetVPNGroup a Get request.
func (service *ContrailService) GetVPNGroup(ctx context.Context, request *models.GetVPNGroupRequest) (response *models.GetVPNGroupResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListVPNGroupRequest{
		Spec: spec,
	}
	var result *models.ListVPNGroupResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListVPNGroup(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VPNGroups) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVPNGroupResponse{
		VPNGroup: result.VPNGroups[0],
	}
	return response, nil
}

//RESTListVPNGroup handles a List REST service Request.
func (service *ContrailService) RESTListVPNGroup(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVPNGroupRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVPNGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListVPNGroup handles a List service Request.
func (service *ContrailService) ListVPNGroup(
	ctx context.Context,
	request *models.ListVPNGroupRequest) (response *models.ListVPNGroupResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListVPNGroup(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
