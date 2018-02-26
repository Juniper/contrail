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

//RESTCreateContrailCluster handle a Create REST service.
func (service *ContrailService) RESTCreateContrailCluster(c echo.Context) error {
	requestData := &models.CreateContrailClusterRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_cluster",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateContrailCluster(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateContrailCluster handle a Create API
func (service *ContrailService) CreateContrailCluster(
	ctx context.Context,
	request *models.CreateContrailClusterRequest) (*models.CreateContrailClusterResponse, error) {
	model := request.ContrailCluster
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
			return db.CreateContrailCluster(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_cluster",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateContrailClusterResponse{
		ContrailCluster: request.ContrailCluster,
	}, nil
}

//RESTUpdateContrailCluster handles a REST Update request.
func (service *ContrailService) RESTUpdateContrailCluster(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateContrailClusterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_cluster",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateContrailCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateContrailCluster handles a Update request.
func (service *ContrailService) UpdateContrailCluster(
	ctx context.Context,
	request *models.UpdateContrailClusterRequest) (*models.UpdateContrailClusterResponse, error) {
	model := request.ContrailCluster
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateContrailCluster(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_cluster",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateContrailClusterResponse{
		ContrailCluster: model,
	}, nil
}

//RESTDeleteContrailCluster delete a resource using REST service.
func (service *ContrailService) RESTDeleteContrailCluster(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteContrailClusterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteContrailCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteContrailCluster delete a resource.
func (service *ContrailService) DeleteContrailCluster(ctx context.Context, request *models.DeleteContrailClusterRequest) (*models.DeleteContrailClusterResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteContrailCluster(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteContrailClusterResponse{
		ID: request.ID,
	}, nil
}

//RESTGetContrailCluster a REST Get request.
func (service *ContrailService) RESTGetContrailCluster(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetContrailClusterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetContrailCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetContrailCluster a Get request.
func (service *ContrailService) GetContrailCluster(ctx context.Context, request *models.GetContrailClusterRequest) (response *models.GetContrailClusterResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListContrailClusterRequest{
		Spec: spec,
	}
	var result *models.ListContrailClusterResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListContrailCluster(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ContrailClusters) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetContrailClusterResponse{
		ContrailCluster: result.ContrailClusters[0],
	}
	return response, nil
}

//RESTListContrailCluster handles a List REST service Request.
func (service *ContrailService) RESTListContrailCluster(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListContrailClusterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListContrailCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListContrailCluster handles a List service Request.
func (service *ContrailService) ListContrailCluster(
	ctx context.Context,
	request *models.ListContrailClusterRequest) (response *models.ListContrailClusterResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListContrailCluster(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
