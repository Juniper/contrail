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

//RESTCreateOpenstackCluster handle a Create REST service.
func (service *ContrailService) RESTCreateOpenstackCluster(c echo.Context) error {
	requestData := &models.CreateOpenstackClusterRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_cluster",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateOpenstackCluster(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateOpenstackCluster handle a Create API
func (service *ContrailService) CreateOpenstackCluster(
	ctx context.Context,
	request *models.CreateOpenstackClusterRequest) (*models.CreateOpenstackClusterResponse, error) {
	model := request.OpenstackCluster
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
			return db.CreateOpenstackCluster(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_cluster",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateOpenstackClusterResponse{
		OpenstackCluster: request.OpenstackCluster,
	}, nil
}

//RESTUpdateOpenstackCluster handles a REST Update request.
func (service *ContrailService) RESTUpdateOpenstackCluster(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateOpenstackClusterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_cluster",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateOpenstackCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateOpenstackCluster handles a Update request.
func (service *ContrailService) UpdateOpenstackCluster(
	ctx context.Context,
	request *models.UpdateOpenstackClusterRequest) (*models.UpdateOpenstackClusterResponse, error) {
	model := request.OpenstackCluster
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateOpenstackCluster(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "openstack_cluster",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateOpenstackClusterResponse{
		OpenstackCluster: model,
	}, nil
}

//RESTDeleteOpenstackCluster delete a resource using REST service.
func (service *ContrailService) RESTDeleteOpenstackCluster(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteOpenstackClusterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteOpenstackCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteOpenstackCluster delete a resource.
func (service *ContrailService) DeleteOpenstackCluster(ctx context.Context, request *models.DeleteOpenstackClusterRequest) (*models.DeleteOpenstackClusterResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteOpenstackCluster(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteOpenstackClusterResponse{
		ID: request.ID,
	}, nil
}

//RESTGetOpenstackCluster a REST Get request.
func (service *ContrailService) RESTGetOpenstackCluster(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetOpenstackClusterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetOpenstackCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetOpenstackCluster a Get request.
func (service *ContrailService) GetOpenstackCluster(ctx context.Context, request *models.GetOpenstackClusterRequest) (response *models.GetOpenstackClusterResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListOpenstackClusterRequest{
		Spec: spec,
	}
	var result *models.ListOpenstackClusterResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListOpenstackCluster(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.OpenstackClusters) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetOpenstackClusterResponse{
		OpenstackCluster: result.OpenstackClusters[0],
	}
	return response, nil
}

//RESTListOpenstackCluster handles a List REST service Request.
func (service *ContrailService) RESTListOpenstackCluster(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListOpenstackClusterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListOpenstackCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListOpenstackCluster handles a List service Request.
func (service *ContrailService) ListOpenstackCluster(
	ctx context.Context,
	request *models.ListOpenstackClusterRequest) (response *models.ListOpenstackClusterResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListOpenstackCluster(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
