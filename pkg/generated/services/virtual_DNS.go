package services

import (
	"context"
	"database/sql"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreateVirtualDNS handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualDNS(c echo.Context) error {
	requestData := &models.CreateVirtualDNSRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualDNS(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualDNS handle a Create API
func (service *ContrailService) CreateVirtualDNS(
	ctx context.Context,
	request *models.CreateVirtualDNSRequest) (*models.CreateVirtualDNSResponse, error) {
	model := request.VirtualDNS
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}
	auth := common.GetAuthCTX(ctx)
	if auth == nil {
		return nil, common.ErrorUnauthenticated
	}

	if model.FQName == nil {
		if model.DisplayName != "" {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
		} else {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.UUID}
		}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.CreateVirtualDNS(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualDNSResponse{
		VirtualDNS: request.VirtualDNS,
	}, nil
}

//RESTUpdateVirtualDNS handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualDNS(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualDNSRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualDNS(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualDNS handles a Update request.
func (service *ContrailService) UpdateVirtualDNS(
	ctx context.Context,
	request *models.UpdateVirtualDNSRequest) (*models.UpdateVirtualDNSResponse, error) {
	model := request.VirtualDNS
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateVirtualDNS(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_DNS",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualDNSResponse{
		VirtualDNS: model,
	}, nil
}

//RESTDeleteVirtualDNS delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualDNS(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualDNSRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualDNS(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualDNS delete a resource.
func (service *ContrailService) DeleteVirtualDNS(ctx context.Context, request *models.DeleteVirtualDNSRequest) (*models.DeleteVirtualDNSResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteVirtualDNS(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualDNSResponse{
		ID: request.ID,
	}, nil
}

//RESTGetVirtualDNS a REST Get request.
func (service *ContrailService) RESTGetVirtualDNS(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualDNSRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualDNS(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetVirtualDNS a Get request.
func (service *ContrailService) GetVirtualDNS(ctx context.Context, request *models.GetVirtualDNSRequest) (response *models.GetVirtualDNSResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListVirtualDNSRequest{
		Spec: spec,
	}
	var result *models.ListVirtualDNSResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListVirtualDNS(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualDNSs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualDNSResponse{
		VirtualDNS: result.VirtualDNSs[0],
	}
	return response, nil
}

//RESTListVirtualDNS handles a List REST service Request.
func (service *ContrailService) RESTListVirtualDNS(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVirtualDNSRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualDNS(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListVirtualDNS handles a List service Request.
func (service *ContrailService) ListVirtualDNS(
	ctx context.Context,
	request *models.ListVirtualDNSRequest) (response *models.ListVirtualDNSResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListVirtualDNS(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
