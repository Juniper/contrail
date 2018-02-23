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

//RESTCreateSecurityGroup handle a Create REST service.
func (service *ContrailService) RESTCreateSecurityGroup(c echo.Context) error {
	requestData := &models.CreateSecurityGroupRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateSecurityGroup(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateSecurityGroup handle a Create API
func (service *ContrailService) CreateSecurityGroup(
	ctx context.Context,
	request *models.CreateSecurityGroupRequest) (*models.CreateSecurityGroupResponse, error) {
	model := request.SecurityGroup
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
			return db.CreateSecurityGroup(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateSecurityGroupResponse{
		SecurityGroup: request.SecurityGroup,
	}, nil
}

//RESTUpdateSecurityGroup handles a REST Update request.
func (service *ContrailService) RESTUpdateSecurityGroup(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateSecurityGroupRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateSecurityGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateSecurityGroup handles a Update request.
func (service *ContrailService) UpdateSecurityGroup(
	ctx context.Context,
	request *models.UpdateSecurityGroupRequest) (*models.UpdateSecurityGroupResponse, error) {
	model := request.SecurityGroup
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateSecurityGroup(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateSecurityGroupResponse{
		SecurityGroup: model,
	}, nil
}

//RESTDeleteSecurityGroup delete a resource using REST service.
func (service *ContrailService) RESTDeleteSecurityGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteSecurityGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteSecurityGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteSecurityGroup delete a resource.
func (service *ContrailService) DeleteSecurityGroup(ctx context.Context, request *models.DeleteSecurityGroupRequest) (*models.DeleteSecurityGroupResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteSecurityGroup(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteSecurityGroupResponse{
		ID: request.ID,
	}, nil
}

//RESTGetSecurityGroup a REST Get request.
func (service *ContrailService) RESTGetSecurityGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetSecurityGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetSecurityGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetSecurityGroup a Get request.
func (service *ContrailService) GetSecurityGroup(ctx context.Context, request *models.GetSecurityGroupRequest) (response *models.GetSecurityGroupResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListSecurityGroupRequest{
		Spec: spec,
	}
	var result *models.ListSecurityGroupResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListSecurityGroup(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.SecurityGroups) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetSecurityGroupResponse{
		SecurityGroup: result.SecurityGroups[0],
	}
	return response, nil
}

//RESTListSecurityGroup handles a List REST service Request.
func (service *ContrailService) RESTListSecurityGroup(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListSecurityGroupRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListSecurityGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListSecurityGroup handles a List service Request.
func (service *ContrailService) ListSecurityGroup(
	ctx context.Context,
	request *models.ListSecurityGroupRequest) (response *models.ListSecurityGroupResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListSecurityGroup(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
