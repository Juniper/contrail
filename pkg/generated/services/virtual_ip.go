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

//RESTVirtualIPUpdateRequest for update request for REST.
type RESTVirtualIPUpdateRequest struct {
	Data map[string]interface{} `json:"virtual-ip"`
}

//RESTCreateVirtualIP handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualIP(c echo.Context) error {
	requestData := &models.CreateVirtualIPRequest{
		VirtualIP: models.MakeVirtualIP(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_ip",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualIP(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualIP handle a Create API
func (service *ContrailService) CreateVirtualIP(
	ctx context.Context,
	request *models.CreateVirtualIPRequest) (*models.CreateVirtualIPResponse, error) {
	model := request.VirtualIP
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
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateVirtualIP(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_ip",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualIPResponse{
		VirtualIP: request.VirtualIP,
	}, nil
}

//RESTUpdateVirtualIP handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualIP(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualIPRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_ip",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualIP handles a Update request.
func (service *ContrailService) UpdateVirtualIP(
	ctx context.Context,
	request *models.UpdateVirtualIPRequest) (*models.UpdateVirtualIPResponse, error) {
	model := request.VirtualIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateVirtualIP(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_ip",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualIPResponse{
		VirtualIP: model,
	}, nil
}

//RESTDeleteVirtualIP delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualIP delete a resource.
func (service *ContrailService) DeleteVirtualIP(ctx context.Context, request *models.DeleteVirtualIPRequest) (*models.DeleteVirtualIPResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteVirtualIP(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualIPResponse{
		ID: request.ID,
	}, nil
}

//RESTGetVirtualIP a REST Get request.
func (service *ContrailService) RESTGetVirtualIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetVirtualIP a Get request.
func (service *ContrailService) GetVirtualIP(ctx context.Context, request *models.GetVirtualIPRequest) (response *models.GetVirtualIPResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListVirtualIPRequest{
		Spec: spec,
	}
	var result *models.ListVirtualIPResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListVirtualIP(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualIPs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualIPResponse{
		VirtualIP: result.VirtualIPs[0],
	}
	return response, nil
}

//RESTListVirtualIP handles a List REST service Request.
func (service *ContrailService) RESTListVirtualIP(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVirtualIPRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListVirtualIP handles a List service Request.
func (service *ContrailService) ListVirtualIP(
	ctx context.Context,
	request *models.ListVirtualIPRequest) (response *models.ListVirtualIPResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListVirtualIP(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
