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

//RESTCreateSubnet handle a Create REST service.
func (service *ContrailService) RESTCreateSubnet(c echo.Context) error {
	requestData := &models.CreateSubnetRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "subnet",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateSubnet(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateSubnet handle a Create API
func (service *ContrailService) CreateSubnet(
	ctx context.Context,
	request *models.CreateSubnetRequest) (*models.CreateSubnetResponse, error) {
	model := request.Subnet
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
			return db.CreateSubnet(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "subnet",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateSubnetResponse{
		Subnet: request.Subnet,
	}, nil
}

//RESTUpdateSubnet handles a REST Update request.
func (service *ContrailService) RESTUpdateSubnet(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateSubnetRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "subnet",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateSubnet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateSubnet handles a Update request.
func (service *ContrailService) UpdateSubnet(
	ctx context.Context,
	request *models.UpdateSubnetRequest) (*models.UpdateSubnetResponse, error) {
	model := request.Subnet
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateSubnet(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "subnet",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateSubnetResponse{
		Subnet: model,
	}, nil
}

//RESTDeleteSubnet delete a resource using REST service.
func (service *ContrailService) RESTDeleteSubnet(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteSubnetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteSubnet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteSubnet delete a resource.
func (service *ContrailService) DeleteSubnet(ctx context.Context, request *models.DeleteSubnetRequest) (*models.DeleteSubnetResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteSubnet(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteSubnetResponse{
		ID: request.ID,
	}, nil
}

//RESTGetSubnet a REST Get request.
func (service *ContrailService) RESTGetSubnet(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetSubnetRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetSubnet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetSubnet a Get request.
func (service *ContrailService) GetSubnet(ctx context.Context, request *models.GetSubnetRequest) (response *models.GetSubnetResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListSubnetRequest{
		Spec: spec,
	}
	var result *models.ListSubnetResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListSubnet(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Subnets) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetSubnetResponse{
		Subnet: result.Subnets[0],
	}
	return response, nil
}

//RESTListSubnet handles a List REST service Request.
func (service *ContrailService) RESTListSubnet(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListSubnetRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListSubnet(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListSubnet handles a List service Request.
func (service *ContrailService) ListSubnet(
	ctx context.Context,
	request *models.ListSubnetRequest) (response *models.ListSubnetResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListSubnet(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
