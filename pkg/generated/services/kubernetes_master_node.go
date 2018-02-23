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

//RESTCreateKubernetesMasterNode handle a Create REST service.
func (service *ContrailService) RESTCreateKubernetesMasterNode(c echo.Context) error {
	requestData := &models.CreateKubernetesMasterNodeRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_master_node",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateKubernetesMasterNode(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateKubernetesMasterNode handle a Create API
func (service *ContrailService) CreateKubernetesMasterNode(
	ctx context.Context,
	request *models.CreateKubernetesMasterNodeRequest) (*models.CreateKubernetesMasterNodeResponse, error) {
	model := request.KubernetesMasterNode
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
			return db.CreateKubernetesMasterNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_master_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateKubernetesMasterNodeResponse{
		KubernetesMasterNode: request.KubernetesMasterNode,
	}, nil
}

//RESTUpdateKubernetesMasterNode handles a REST Update request.
func (service *ContrailService) RESTUpdateKubernetesMasterNode(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateKubernetesMasterNodeRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_master_node",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateKubernetesMasterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateKubernetesMasterNode handles a Update request.
func (service *ContrailService) UpdateKubernetesMasterNode(
	ctx context.Context,
	request *models.UpdateKubernetesMasterNodeRequest) (*models.UpdateKubernetesMasterNodeResponse, error) {
	model := request.KubernetesMasterNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateKubernetesMasterNode(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_master_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateKubernetesMasterNodeResponse{
		KubernetesMasterNode: model,
	}, nil
}

//RESTDeleteKubernetesMasterNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteKubernetesMasterNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteKubernetesMasterNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteKubernetesMasterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteKubernetesMasterNode delete a resource.
func (service *ContrailService) DeleteKubernetesMasterNode(ctx context.Context, request *models.DeleteKubernetesMasterNodeRequest) (*models.DeleteKubernetesMasterNodeResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteKubernetesMasterNode(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteKubernetesMasterNodeResponse{
		ID: request.ID,
	}, nil
}

//RESTGetKubernetesMasterNode a REST Get request.
func (service *ContrailService) RESTGetKubernetesMasterNode(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetKubernetesMasterNodeRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetKubernetesMasterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetKubernetesMasterNode a Get request.
func (service *ContrailService) GetKubernetesMasterNode(ctx context.Context, request *models.GetKubernetesMasterNodeRequest) (response *models.GetKubernetesMasterNodeResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListKubernetesMasterNodeRequest{
		Spec: spec,
	}
	var result *models.ListKubernetesMasterNodeResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListKubernetesMasterNode(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.KubernetesMasterNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetKubernetesMasterNodeResponse{
		KubernetesMasterNode: result.KubernetesMasterNodes[0],
	}
	return response, nil
}

//RESTListKubernetesMasterNode handles a List REST service Request.
func (service *ContrailService) RESTListKubernetesMasterNode(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListKubernetesMasterNodeRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListKubernetesMasterNode(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListKubernetesMasterNode handles a List service Request.
func (service *ContrailService) ListKubernetesMasterNode(
	ctx context.Context,
	request *models.ListKubernetesMasterNodeRequest) (response *models.ListKubernetesMasterNodeResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListKubernetesMasterNode(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
