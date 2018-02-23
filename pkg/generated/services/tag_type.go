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

//RESTCreateTagType handle a Create REST service.
func (service *ContrailService) RESTCreateTagType(c echo.Context) error {
	requestData := &models.CreateTagTypeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag_type",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateTagType(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateTagType handle a Create API
func (service *ContrailService) CreateTagType(
	ctx context.Context,
	request *models.CreateTagTypeRequest) (*models.CreateTagTypeResponse, error) {
	model := request.TagType
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
			return db.CreateTagType(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag_type",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateTagTypeResponse{
		TagType: request.TagType,
	}, nil
}

//RESTUpdateTagType handles a REST Update request.
func (service *ContrailService) RESTUpdateTagType(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateTagTypeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag_type",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateTagType(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateTagType handles a Update request.
func (service *ContrailService) UpdateTagType(
	ctx context.Context,
	request *models.UpdateTagTypeRequest) (*models.UpdateTagTypeResponse, error) {
	model := request.TagType
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateTagType(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "tag_type",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateTagTypeResponse{
		TagType: model,
	}, nil
}

//RESTDeleteTagType delete a resource using REST service.
func (service *ContrailService) RESTDeleteTagType(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteTagTypeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteTagType(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteTagType delete a resource.
func (service *ContrailService) DeleteTagType(ctx context.Context, request *models.DeleteTagTypeRequest) (*models.DeleteTagTypeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteTagType(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteTagTypeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetTagType a REST Get request.
func (service *ContrailService) RESTGetTagType(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetTagTypeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetTagType(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetTagType a Get request.
func (service *ContrailService) GetTagType(ctx context.Context, request *models.GetTagTypeRequest) (response *models.GetTagTypeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListTagTypeRequest{
		Spec: spec,
	}
	var result *models.ListTagTypeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListTagType(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.TagTypes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetTagTypeResponse{
		TagType: result.TagTypes[0],
	}
	return response, nil
}

//RESTListTagType handles a List REST service Request.
func (service *ContrailService) RESTListTagType(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListTagTypeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListTagType(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListTagType handles a List service Request.
func (service *ContrailService) ListTagType(
	ctx context.Context,
	request *models.ListTagTypeRequest) (response *models.ListTagTypeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListTagType(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
