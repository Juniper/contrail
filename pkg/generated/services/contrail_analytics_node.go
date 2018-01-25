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

//RESTContrailAnalyticsNodeUpdateRequest for update request for REST.
type RESTContrailAnalyticsNodeUpdateRequest struct {
	Data map[string]interface{} `json:"contrail-analytics-node"`
}

//RESTCreateContrailAnalyticsNode handle a Create REST service.
func (service *ContrailService) RESTCreateContrailAnalyticsNode(c echo.Context) error {
	requestData := &models.CreateContrailAnalyticsNodeRequest{
		ContrailAnalyticsNode: models.MakeContrailAnalyticsNode(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailAnalyticsNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailAnalyticsNode handle a Create API
func (service *ContrailService) CreateContrailAnalyticsNode(
	ctx context.Context,
	request *models.CreateContrailAnalyticsNodeRequest) (*models.CreateContrailAnalyticsNodeResponse, error) {
	model := request.ContrailAnalyticsNode
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
			return db.CreateContrailAnalyticsNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateContrailAnalyticsNodeResponse{
		ContrailAnalyticsNode: request.ContrailAnalyticsNode,
	}, nil
}

//RESTUpdateContrailAnalyticsNode handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailAnalyticsNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailAnalyticsNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailAnalyticsNode handles a Update request.
func (service *ContrailService) UpdateContrailAnalyticsNode(
	ctx context.Context,
	request *models.UpdateContrailAnalyticsNodeRequest) (*models.UpdateContrailAnalyticsNodeResponse, error) {
	model := request.ContrailAnalyticsNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateContrailAnalyticsNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_analytics_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateContrailAnalyticsNodeResponse{
		ContrailAnalyticsNode: model,
	}, nil
}

//RESTDeleteContrailAnalyticsNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailAnalyticsNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailAnalyticsNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailAnalyticsNode delete a resource.
func (service *ContrailService) DeleteContrailAnalyticsNode(ctx context.Context, request *models.DeleteContrailAnalyticsNodeRequest) (*models.DeleteContrailAnalyticsNodeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteContrailAnalyticsNode(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteContrailAnalyticsNodeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetContrailAnalyticsNode a REST Get request.
func (service *ContrailService) RESTGetContrailAnalyticsNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailAnalyticsNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetContrailAnalyticsNode a Get request.
func (service *ContrailService) GetContrailAnalyticsNode(ctx context.Context, request *models.GetContrailAnalyticsNodeRequest) (response *models.GetContrailAnalyticsNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListContrailAnalyticsNodeRequest{
		Spec: spec,
	}
	var result *models.ListContrailAnalyticsNodeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailAnalyticsNode(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ContrailAnalyticsNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetContrailAnalyticsNodeResponse{
		ContrailAnalyticsNode: result.ContrailAnalyticsNodes[0],
	}
	return response, nil
}

//RESTListContrailAnalyticsNode handles a List REST service Request.
func (service *ContrailService) RESTListContrailAnalyticsNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailAnalyticsNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListContrailAnalyticsNode handles a List service Request.
func (service *ContrailService) ListContrailAnalyticsNode(
	ctx context.Context,
	request *models.ListContrailAnalyticsNodeRequest) (response *models.ListContrailAnalyticsNodeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListContrailAnalyticsNode(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
