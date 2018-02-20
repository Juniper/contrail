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

//RESTAPIAccessListUpdateRequest for update request for REST.
type RESTAPIAccessListUpdateRequest struct {
	Data map[string]interface{} `json:"api-access-list"`
}

//RESTCreateAPIAccessList handle a Create REST service.
func (service *ContrailService) RESTCreateAPIAccessList(c echo.Context) error {
	requestData := &models.CreateAPIAccessListRequest{
		APIAccessList: models.MakeAPIAccessList(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "api_access_list",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAPIAccessList(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAPIAccessList handle a Create API
func (service *ContrailService) CreateAPIAccessList(
	ctx context.Context,
	request *models.CreateAPIAccessListRequest) (*models.CreateAPIAccessListResponse, error) {
	model := request.APIAccessList
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
			return db.CreateAPIAccessList(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "api_access_list",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAPIAccessListResponse{
		APIAccessList: request.APIAccessList,
	}, nil
}

//RESTUpdateAPIAccessList handles a REST Update request.
func (service *ContrailService) RESTUpdateAPIAccessList(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAPIAccessListRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "api_access_list",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAPIAccessList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAPIAccessList handles a Update request.
func (service *ContrailService) UpdateAPIAccessList(
	ctx context.Context,
	request *models.UpdateAPIAccessListRequest) (*models.UpdateAPIAccessListResponse, error) {
	model := request.APIAccessList
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateAPIAccessList(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "api_access_list",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAPIAccessListResponse{
		APIAccessList: model,
	}, nil
}

//RESTDeleteAPIAccessList delete a resource using REST service.
func (service *ContrailService) RESTDeleteAPIAccessList(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAPIAccessListRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAPIAccessList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteAPIAccessList delete a resource.
func (service *ContrailService) DeleteAPIAccessList(ctx context.Context, request *models.DeleteAPIAccessListRequest) (*models.DeleteAPIAccessListResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteAPIAccessList(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAPIAccessListResponse{
		ID: request.ID,
	}, nil
}

//RESTGetAPIAccessList a REST Get request.
func (service *ContrailService) RESTGetAPIAccessList(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAPIAccessListRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAPIAccessList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetAPIAccessList a Get request.
func (service *ContrailService) GetAPIAccessList(ctx context.Context, request *models.GetAPIAccessListRequest) (response *models.GetAPIAccessListResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListAPIAccessListRequest{
		Spec: spec,
	}
	var result *models.ListAPIAccessListResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListAPIAccessList(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.APIAccessLists) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAPIAccessListResponse{
		APIAccessList: result.APIAccessLists[0],
	}
	return response, nil
}

//RESTListAPIAccessList handles a List REST service Request.
func (service *ContrailService) RESTListAPIAccessList(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAPIAccessListRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAPIAccessList(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListAPIAccessList handles a List service Request.
func (service *ContrailService) ListAPIAccessList(
	ctx context.Context,
	request *models.ListAPIAccessListRequest) (response *models.ListAPIAccessListResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListAPIAccessList(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
