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

//RESTCreatePhysicalInterface handle a Create REST service.
func (service *ContrailService) RESTCreatePhysicalInterface(c echo.Context) error {
	requestData := &models.CreatePhysicalInterfaceRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_interface",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreatePhysicalInterface(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreatePhysicalInterface handle a Create API
func (service *ContrailService) CreatePhysicalInterface(
	ctx context.Context,
	request *models.CreatePhysicalInterfaceRequest) (*models.CreatePhysicalInterfaceResponse, error) {
	model := request.PhysicalInterface
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
			return db.CreatePhysicalInterface(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_interface",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreatePhysicalInterfaceResponse{
		PhysicalInterface: request.PhysicalInterface,
	}, nil
}

//RESTUpdatePhysicalInterface handles a REST Update request.
func (service *ContrailService) RESTUpdatePhysicalInterface(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdatePhysicalInterfaceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_interface",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdatePhysicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdatePhysicalInterface handles a Update request.
func (service *ContrailService) UpdatePhysicalInterface(
	ctx context.Context,
	request *models.UpdatePhysicalInterfaceRequest) (*models.UpdatePhysicalInterfaceResponse, error) {
	model := request.PhysicalInterface
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdatePhysicalInterface(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_interface",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdatePhysicalInterfaceResponse{
		PhysicalInterface: model,
	}, nil
}

//RESTDeletePhysicalInterface delete a resource using REST service.
func (service *ContrailService) RESTDeletePhysicalInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeletePhysicalInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeletePhysicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeletePhysicalInterface delete a resource.
func (service *ContrailService) DeletePhysicalInterface(ctx context.Context, request *models.DeletePhysicalInterfaceRequest) (*models.DeletePhysicalInterfaceResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeletePhysicalInterface(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeletePhysicalInterfaceResponse{
		ID: request.ID,
	}, nil
}

//RESTGetPhysicalInterface a REST Get request.
func (service *ContrailService) RESTGetPhysicalInterface(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetPhysicalInterfaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetPhysicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetPhysicalInterface a Get request.
func (service *ContrailService) GetPhysicalInterface(ctx context.Context, request *models.GetPhysicalInterfaceRequest) (response *models.GetPhysicalInterfaceResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListPhysicalInterfaceRequest{
		Spec: spec,
	}
	var result *models.ListPhysicalInterfaceResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListPhysicalInterface(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.PhysicalInterfaces) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetPhysicalInterfaceResponse{
		PhysicalInterface: result.PhysicalInterfaces[0],
	}
	return response, nil
}

//RESTListPhysicalInterface handles a List REST service Request.
func (service *ContrailService) RESTListPhysicalInterface(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListPhysicalInterfaceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListPhysicalInterface(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListPhysicalInterface handles a List service Request.
func (service *ContrailService) ListPhysicalInterface(
	ctx context.Context,
	request *models.ListPhysicalInterfaceRequest) (response *models.ListPhysicalInterfaceResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListPhysicalInterface(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
