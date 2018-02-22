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

//RESTAnalyticsNodeUpdateRequest for update request for REST.
type RESTAnalyticsNodeUpdateRequest struct {
	Data map[string]interface{} `json:"analytics-node"`
}

//RESTCreateAnalyticsNode handle a Create REST service.
func (service *ContrailService) RESTCreateAnalyticsNode(c echo.Context) error {
	requestData := &models.CreateAnalyticsNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "analytics_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAnalyticsNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAnalyticsNode handle a Create API
func (service *ContrailService) CreateAnalyticsNode(
	ctx context.Context,
	request *models.CreateAnalyticsNodeRequest) (*models.CreateAnalyticsNodeResponse, error) {
	model := request.AnalyticsNode
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
			return db.CreateAnalyticsNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "analytics_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAnalyticsNodeResponse{
		AnalyticsNode: request.AnalyticsNode,
	}, nil
}

//RESTUpdateAnalyticsNode handles a REST Update request.
func (service *ContrailService) RESTUpdateAnalyticsNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAnalyticsNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "analytics_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAnalyticsNode handles a Update request.
func (service *ContrailService) UpdateAnalyticsNode(
	ctx context.Context,
	request *models.UpdateAnalyticsNodeRequest) (*models.UpdateAnalyticsNodeResponse, error) {
	model := request.AnalyticsNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateAnalyticsNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "analytics_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAnalyticsNodeResponse{
		AnalyticsNode: model,
	}, nil
}

//RESTDeleteAnalyticsNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteAnalyticsNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAnalyticsNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteAnalyticsNode delete a resource.
func (service *ContrailService) DeleteAnalyticsNode(ctx context.Context, request *models.DeleteAnalyticsNodeRequest) (*models.DeleteAnalyticsNodeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteAnalyticsNode(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAnalyticsNodeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetAnalyticsNode a REST Get request.
func (service *ContrailService) RESTGetAnalyticsNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAnalyticsNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetAnalyticsNode a Get request.
func (service *ContrailService) GetAnalyticsNode(ctx context.Context, request *models.GetAnalyticsNodeRequest) (response *models.GetAnalyticsNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListAnalyticsNodeRequest{
		Spec: spec,
	}
	var result *models.ListAnalyticsNodeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListAnalyticsNode(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.AnalyticsNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAnalyticsNodeResponse{
		AnalyticsNode: result.AnalyticsNodes[0],
	}
	return response, nil
}

//RESTListAnalyticsNode handles a List REST service Request.
func (service *ContrailService) RESTListAnalyticsNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAnalyticsNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAnalyticsNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListAnalyticsNode handles a List service Request.
func (service *ContrailService) ListAnalyticsNode(
	ctx context.Context,
	request *models.ListAnalyticsNodeRequest) (response *models.ListAnalyticsNodeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListAnalyticsNode(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
