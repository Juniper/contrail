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

//RESTVirtualMachineUpdateRequest for update request for REST.
type RESTVirtualMachineUpdateRequest struct {
	Data map[string]interface{} `json:"virtual-machine"`
}

//RESTCreateVirtualMachine handle a Create REST service.
func (service *ContrailService) RESTCreateVirtualMachine(c echo.Context) error {
	requestData := &models.CreateVirtualMachineRequest{
		VirtualMachine: models.MakeVirtualMachine(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateVirtualMachine(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateVirtualMachine handle a Create API
func (service *ContrailService) CreateVirtualMachine(
	ctx context.Context,
	request *models.CreateVirtualMachineRequest) (*models.CreateVirtualMachineResponse, error) {
	model := request.VirtualMachine
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
			return db.CreateVirtualMachine(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualMachineResponse{
		VirtualMachine: request.VirtualMachine,
	}, nil
}

//RESTUpdateVirtualMachine handles a REST Update request.
func (service *ContrailService) RESTUpdateVirtualMachine(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateVirtualMachineRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateVirtualMachine(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateVirtualMachine handles a Update request.
func (service *ContrailService) UpdateVirtualMachine(
	ctx context.Context,
	request *models.UpdateVirtualMachineRequest) (*models.UpdateVirtualMachineResponse, error) {
	model := request.VirtualMachine
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateVirtualMachine(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_machine",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualMachineResponse{
		VirtualMachine: model,
	}, nil
}

//RESTDeleteVirtualMachine delete a resource using REST service.
func (service *ContrailService) RESTDeleteVirtualMachine(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteVirtualMachineRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteVirtualMachine(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteVirtualMachine delete a resource.
func (service *ContrailService) DeleteVirtualMachine(ctx context.Context, request *models.DeleteVirtualMachineRequest) (*models.DeleteVirtualMachineResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteVirtualMachine(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualMachineResponse{
		ID: request.ID,
	}, nil
}

//RESTGetVirtualMachine a REST Get request.
func (service *ContrailService) RESTGetVirtualMachine(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetVirtualMachineRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetVirtualMachine(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetVirtualMachine a Get request.
func (service *ContrailService) GetVirtualMachine(ctx context.Context, request *models.GetVirtualMachineRequest) (response *models.GetVirtualMachineResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListVirtualMachineRequest{
		Spec: spec,
	}
	var result *models.ListVirtualMachineResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListVirtualMachine(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualMachines) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualMachineResponse{
		VirtualMachine: result.VirtualMachines[0],
	}
	return response, nil
}

//RESTListVirtualMachine handles a List REST service Request.
func (service *ContrailService) RESTListVirtualMachine(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListVirtualMachineRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListVirtualMachine(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListVirtualMachine handles a List service Request.
func (service *ContrailService) ListVirtualMachine(
	ctx context.Context,
	request *models.ListVirtualMachineRequest) (response *models.ListVirtualMachineResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListVirtualMachine(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
