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

//RESTCreateAccessControlList handle a Create REST service.
func (service *ContrailService) RESTCreateAccessControlList(c echo.Context) error {
	requestData := &models.CreateAccessControlListRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "access_control_list",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAccessControlList(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAccessControlList handle a Create API
func (service *ContrailService) CreateAccessControlList(
	ctx context.Context,
	request *models.CreateAccessControlListRequest) (*models.CreateAccessControlListResponse, error) {
	model := request.AccessControlList
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
			return db.CreateAccessControlList(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "access_control_list",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAccessControlListResponse{
		AccessControlList: request.AccessControlList,
	}, nil
}

//RESTUpdateAccessControlList handles a REST Update request.
func (service *ContrailService) RESTUpdateAccessControlList(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAccessControlListRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "access_control_list",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAccessControlList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAccessControlList handles a Update request.
func (service *ContrailService) UpdateAccessControlList(
	ctx context.Context,
	request *models.UpdateAccessControlListRequest) (*models.UpdateAccessControlListResponse, error) {
	model := request.AccessControlList
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateAccessControlList(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "access_control_list",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAccessControlListResponse{
		AccessControlList: model,
	}, nil
}

//RESTDeleteAccessControlList delete a resource using REST service.
func (service *ContrailService) RESTDeleteAccessControlList(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAccessControlListRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAccessControlList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteAccessControlList delete a resource.
func (service *ContrailService) DeleteAccessControlList(ctx context.Context, request *models.DeleteAccessControlListRequest) (*models.DeleteAccessControlListResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteAccessControlList(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAccessControlListResponse{
		ID: request.ID,
	}, nil
}

//RESTGetAccessControlList a REST Get request.
func (service *ContrailService) RESTGetAccessControlList(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAccessControlListRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAccessControlList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetAccessControlList a Get request.
func (service *ContrailService) GetAccessControlList(ctx context.Context, request *models.GetAccessControlListRequest) (response *models.GetAccessControlListResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListAccessControlListRequest{
		Spec: spec,
	}
	var result *models.ListAccessControlListResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListAccessControlList(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.AccessControlLists) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAccessControlListResponse{
		AccessControlList: result.AccessControlLists[0],
	}
	return response, nil
}

//RESTListAccessControlList handles a List REST service Request.
func (service *ContrailService) RESTListAccessControlList(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAccessControlListRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAccessControlList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListAccessControlList handles a List service Request.
func (service *ContrailService) ListAccessControlList(
	ctx context.Context,
	request *models.ListAccessControlListRequest) (response *models.ListAccessControlListResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListAccessControlList(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
