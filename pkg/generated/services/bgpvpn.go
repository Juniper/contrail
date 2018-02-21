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

//RESTBGPVPNUpdateRequest for update request for REST.
type RESTBGPVPNUpdateRequest struct {
	Data map[string]interface{} `json:"bgpvpn"`
}

//RESTCreateBGPVPN handle a Create REST service.
func (service *ContrailService) RESTCreateBGPVPN(c echo.Context) error {
	requestData := &models.CreateBGPVPNRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgpvpn",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBGPVPN(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBGPVPN handle a Create API
func (service *ContrailService) CreateBGPVPN(
	ctx context.Context,
	request *models.CreateBGPVPNRequest) (*models.CreateBGPVPNResponse, error) {
	model := request.BGPVPN
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
			return db.CreateBGPVPN(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgpvpn",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBGPVPNResponse{
		BGPVPN: request.BGPVPN,
	}, nil
}

//RESTUpdateBGPVPN handles a REST Update request.
func (service *ContrailService) RESTUpdateBGPVPN(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBGPVPNRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgpvpn",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBGPVPN(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBGPVPN handles a Update request.
func (service *ContrailService) UpdateBGPVPN(
	ctx context.Context,
	request *models.UpdateBGPVPNRequest) (*models.UpdateBGPVPNResponse, error) {
	model := request.BGPVPN
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateBGPVPN(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bgpvpn",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBGPVPNResponse{
		BGPVPN: model,
	}, nil
}

//RESTDeleteBGPVPN delete a resource using REST service.
func (service *ContrailService) RESTDeleteBGPVPN(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBGPVPNRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBGPVPN(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteBGPVPN delete a resource.
func (service *ContrailService) DeleteBGPVPN(ctx context.Context, request *models.DeleteBGPVPNRequest) (*models.DeleteBGPVPNResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteBGPVPN(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBGPVPNResponse{
		ID: request.ID,
	}, nil
}

//RESTGetBGPVPN a REST Get request.
func (service *ContrailService) RESTGetBGPVPN(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBGPVPNRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBGPVPN(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetBGPVPN a Get request.
func (service *ContrailService) GetBGPVPN(ctx context.Context, request *models.GetBGPVPNRequest) (response *models.GetBGPVPNResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListBGPVPNRequest{
		Spec: spec,
	}
	var result *models.ListBGPVPNResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListBGPVPN(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BGPVPNs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBGPVPNResponse{
		BGPVPN: result.BGPVPNs[0],
	}
	return response, nil
}

//RESTListBGPVPN handles a List REST service Request.
func (service *ContrailService) RESTListBGPVPN(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListBGPVPNRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBGPVPN(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListBGPVPN handles a List service Request.
func (service *ContrailService) ListBGPVPN(
	ctx context.Context,
	request *models.ListBGPVPNRequest) (response *models.ListBGPVPNResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListBGPVPN(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
