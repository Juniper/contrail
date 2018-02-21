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

//RESTTagUpdateRequest for update request for REST.
type RESTTagUpdateRequest struct {
	Data map[string]interface{} `json:"tag"`
}

//RESTCreateTag handle a Create REST service.
func (service *ContrailService) RESTCreateTag(c echo.Context) error {
	requestData := &models.CreateTagRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateTag(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateTag handle a Create API
func (service *ContrailService) CreateTag(
	ctx context.Context,
	request *models.CreateTagRequest) (*models.CreateTagResponse, error) {
	model := request.Tag
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
			return db.CreateTag(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateTagResponse{
		Tag: request.Tag,
	}, nil
}

//RESTUpdateTag handles a REST Update request.
func (service *ContrailService) RESTUpdateTag(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateTagRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateTag(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateTag handles a Update request.
func (service *ContrailService) UpdateTag(
	ctx context.Context,
	request *models.UpdateTagRequest) (*models.UpdateTagResponse, error) {
	model := request.Tag
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateTag(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateTagResponse{
		Tag: model,
	}, nil
}

//RESTDeleteTag delete a resource using REST service.
func (service *ContrailService) RESTDeleteTag(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteTagRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteTag(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteTag delete a resource.
func (service *ContrailService) DeleteTag(ctx context.Context, request *models.DeleteTagRequest) (*models.DeleteTagResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteTag(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteTagResponse{
		ID: request.ID,
	}, nil
}

//RESTGetTag a REST Get request.
func (service *ContrailService) RESTGetTag(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetTagRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetTag(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetTag a Get request.
func (service *ContrailService) GetTag(ctx context.Context, request *models.GetTagRequest) (response *models.GetTagResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListTagRequest{
		Spec: spec,
	}
	var result *models.ListTagResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListTag(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Tags) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetTagResponse{
		Tag: result.Tags[0],
	}
	return response, nil
}

//RESTListTag handles a List REST service Request.
func (service *ContrailService) RESTListTag(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListTagRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListTag(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListTag handles a List service Request.
func (service *ContrailService) ListTag(
	ctx context.Context,
	request *models.ListTagRequest) (response *models.ListTagResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListTag(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
