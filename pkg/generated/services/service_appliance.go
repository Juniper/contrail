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

//RESTCreateServiceAppliance handle a Create REST service.
func (service *ContrailService) RESTCreateServiceAppliance(c echo.Context) error {
	requestData := &models.CreateServiceApplianceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateServiceAppliance(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateServiceAppliance handle a Create API
func (service *ContrailService) CreateServiceAppliance(
	ctx context.Context,
	request *models.CreateServiceApplianceRequest) (*models.CreateServiceApplianceResponse, error) {
	model := request.ServiceAppliance
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
			return db.CreateServiceAppliance(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateServiceApplianceResponse{
		ServiceAppliance: request.ServiceAppliance,
	}, nil
}

//RESTUpdateServiceAppliance handles a REST Update request.
func (service *ContrailService) RESTUpdateServiceAppliance(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateServiceApplianceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateServiceAppliance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateServiceAppliance handles a Update request.
func (service *ContrailService) UpdateServiceAppliance(
	ctx context.Context,
	request *models.UpdateServiceApplianceRequest) (*models.UpdateServiceApplianceResponse, error) {
	model := request.ServiceAppliance
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateServiceAppliance(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "service_appliance",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateServiceApplianceResponse{
		ServiceAppliance: model,
	}, nil
}

//RESTDeleteServiceAppliance delete a resource using REST service.
func (service *ContrailService) RESTDeleteServiceAppliance(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteServiceApplianceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteServiceAppliance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteServiceAppliance delete a resource.
func (service *ContrailService) DeleteServiceAppliance(ctx context.Context, request *models.DeleteServiceApplianceRequest) (*models.DeleteServiceApplianceResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteServiceAppliance(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteServiceApplianceResponse{
		ID: request.ID,
	}, nil
}

//RESTGetServiceAppliance a REST Get request.
func (service *ContrailService) RESTGetServiceAppliance(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetServiceApplianceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetServiceAppliance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetServiceAppliance a Get request.
func (service *ContrailService) GetServiceAppliance(ctx context.Context, request *models.GetServiceApplianceRequest) (response *models.GetServiceApplianceResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListServiceApplianceRequest{
		Spec: spec,
	}
	var result *models.ListServiceApplianceResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListServiceAppliance(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ServiceAppliances) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetServiceApplianceResponse{
		ServiceAppliance: result.ServiceAppliances[0],
	}
	return response, nil
}

//RESTListServiceAppliance handles a List REST service Request.
func (service *ContrailService) RESTListServiceAppliance(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListServiceApplianceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListServiceAppliance(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListServiceAppliance handles a List service Request.
func (service *ContrailService) ListServiceAppliance(
	ctx context.Context,
	request *models.ListServiceApplianceRequest) (response *models.ListServiceApplianceResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListServiceAppliance(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
