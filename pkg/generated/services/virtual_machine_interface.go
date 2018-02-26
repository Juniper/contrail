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

//RESTCreateVirtualMachineInterface handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualMachineInterface(c echo.Context) error {
	requestData := &models.CreateVirtualMachineInterfaceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine_interface",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualMachineInterface(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualMachineInterface handle a Create API
func (service *ContrailService) CreateVirtualMachineInterface(
	ctx context.Context,
	request *models.CreateVirtualMachineInterfaceRequest) (*models.CreateVirtualMachineInterfaceResponse, error) {
	model := request.VirtualMachineInterface
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
			return db.CreateVirtualMachineInterface(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine_interface",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualMachineInterfaceResponse{
		VirtualMachineInterface: request.VirtualMachineInterface,
	}, nil
}

//RESTUpdateVirtualMachineInterface handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualMachineInterface(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualMachineInterfaceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine_interface",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualMachineInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualMachineInterface handles a Update request.
func (service *ContrailService) UpdateVirtualMachineInterface(
	ctx context.Context,
	request *models.UpdateVirtualMachineInterfaceRequest) (*models.UpdateVirtualMachineInterfaceResponse, error) {
	model := request.VirtualMachineInterface
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateVirtualMachineInterface(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine_interface",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualMachineInterfaceResponse{
		VirtualMachineInterface: model,
	}, nil
}

//RESTDeleteVirtualMachineInterface delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualMachineInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualMachineInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualMachineInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualMachineInterface delete a resource.
func (service *ContrailService) DeleteVirtualMachineInterface(ctx context.Context, request *models.DeleteVirtualMachineInterfaceRequest) (*models.DeleteVirtualMachineInterfaceResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteVirtualMachineInterface(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualMachineInterfaceResponse{
		ID: request.ID,
	}, nil
}

//RESTGetVirtualMachineInterface a REST Get request.
func (service *ContrailService) RESTGetVirtualMachineInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualMachineInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualMachineInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetVirtualMachineInterface a Get request.
func (service *ContrailService) GetVirtualMachineInterface(ctx context.Context, request *models.GetVirtualMachineInterfaceRequest) (response *models.GetVirtualMachineInterfaceResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListVirtualMachineInterfaceRequest{
		Spec: spec,
	}
	var result *models.ListVirtualMachineInterfaceResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListVirtualMachineInterface(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualMachineInterfaces) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualMachineInterfaceResponse{
		VirtualMachineInterface: result.VirtualMachineInterfaces[0],
	}
	return response, nil
}

//RESTListVirtualMachineInterface handles a List REST service Request.
func (service *ContrailService) RESTListVirtualMachineInterface(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVirtualMachineInterfaceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualMachineInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListVirtualMachineInterface handles a List service Request.
func (service *ContrailService) ListVirtualMachineInterface(
	ctx context.Context,
	request *models.ListVirtualMachineInterfaceRequest) (response *models.ListVirtualMachineInterfaceResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListVirtualMachineInterface(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
