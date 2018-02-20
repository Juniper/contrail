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

//RESTKubernetesClusterUpdateRequest for update request for REST.
type RESTKubernetesClusterUpdateRequest struct {
	Data map[string]interface{} `json:"kubernetes-cluster"`
}

//RESTCreateKubernetesCluster handle a Create REST service.
func (service *ContrailService) RESTCreateKubernetesCluster(c echo.Context) error {
	requestData := &models.CreateKubernetesClusterRequest{
		KubernetesCluster: models.MakeKubernetesCluster(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_cluster",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateKubernetesCluster(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateKubernetesCluster handle a Create API
func (service *ContrailService) CreateKubernetesCluster(
	ctx context.Context,
	request *models.CreateKubernetesClusterRequest) (*models.CreateKubernetesClusterResponse, error) {
	model := request.KubernetesCluster
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
			return db.CreateKubernetesCluster(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_cluster",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateKubernetesClusterResponse{
		KubernetesCluster: request.KubernetesCluster,
	}, nil
}

//RESTUpdateKubernetesCluster handles a REST Update request.
func (service *ContrailService) RESTUpdateKubernetesCluster(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateKubernetesClusterRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_cluster",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateKubernetesCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateKubernetesCluster handles a Update request.
func (service *ContrailService) UpdateKubernetesCluster(
	ctx context.Context,
	request *models.UpdateKubernetesClusterRequest) (*models.UpdateKubernetesClusterResponse, error) {
	model := request.KubernetesCluster
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateKubernetesCluster(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "kubernetes_cluster",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateKubernetesClusterResponse{
		KubernetesCluster: model,
	}, nil
}

//RESTDeleteKubernetesCluster delete a resource using REST service.
func (service *ContrailService) RESTDeleteKubernetesCluster(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteKubernetesClusterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteKubernetesCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteKubernetesCluster delete a resource.
func (service *ContrailService) DeleteKubernetesCluster(ctx context.Context, request *models.DeleteKubernetesClusterRequest) (*models.DeleteKubernetesClusterResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteKubernetesCluster(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteKubernetesClusterResponse{
		ID: request.ID,
	}, nil
}

//RESTGetKubernetesCluster a REST Get request.
func (service *ContrailService) RESTGetKubernetesCluster(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetKubernetesClusterRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetKubernetesCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetKubernetesCluster a Get request.
func (service *ContrailService) GetKubernetesCluster(ctx context.Context, request *models.GetKubernetesClusterRequest) (response *models.GetKubernetesClusterResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListKubernetesClusterRequest{
		Spec: spec,
	}
	var result *models.ListKubernetesClusterResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListKubernetesCluster(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.KubernetesClusters) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetKubernetesClusterResponse{
		KubernetesCluster: result.KubernetesClusters[0],
	}
	return response, nil
}

//RESTListKubernetesCluster handles a List REST service Request.
func (service *ContrailService) RESTListKubernetesCluster(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListKubernetesClusterRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListKubernetesCluster(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListKubernetesCluster handles a List service Request.
func (service *ContrailService) ListKubernetesCluster(
	ctx context.Context,
	request *models.ListKubernetesClusterRequest) (response *models.ListKubernetesClusterResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListKubernetesCluster(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
