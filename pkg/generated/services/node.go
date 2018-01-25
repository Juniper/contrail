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

//RESTNodeUpdateRequest for update request for REST.
type RESTNodeUpdateRequest struct {
	Data map[string]interface{} `json:"node"`
}

//RESTCreateNode handle a Create REST service.
func (service *ContrailService) RESTCreateNode(c echo.Context) error {
	requestData := &models.CreateNodeRequest{
		Node: models.MakeNode(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateNode handle a Create API
func (service *ContrailService) CreateNode(
	ctx context.Context,
	request *models.CreateNodeRequest) (*models.CreateNodeResponse, error) {
	model := request.Node
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
			return db.CreateNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateNodeResponse{
		Node: request.Node,
	}, nil
}

//RESTUpdateNode handles a REST Update request.
func (service *ContrailService) RESTUpdateNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateNode handles a Update request.
func (service *ContrailService) UpdateNode(
	ctx context.Context,
	request *models.UpdateNodeRequest) (*models.UpdateNodeResponse, error) {
	model := request.Node
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateNodeResponse{
		Node: model,
	}, nil
}

//RESTDeleteNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteNode delete a resource.
func (service *ContrailService) DeleteNode(ctx context.Context, request *models.DeleteNodeRequest) (*models.DeleteNodeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteNode(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteNodeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetNode a REST Get request.
func (service *ContrailService) RESTGetNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetNode a Get request.
func (service *ContrailService) GetNode(ctx context.Context, request *models.GetNodeRequest) (response *models.GetNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListNodeRequest{
		Spec: spec,
	}
	var result *models.ListNodeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListNode(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Nodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetNodeResponse{
		Node: result.Nodes[0],
	}
	return response, nil
}

//RESTListNode handles a List REST service Request.
func (service *ContrailService) RESTListNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListNode handles a List service Request.
func (service *ContrailService) ListNode(
	ctx context.Context,
	request *models.ListNodeRequest) (response *models.ListNodeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListNode(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
