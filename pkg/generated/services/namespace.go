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

//RESTNamespaceUpdateRequest for update request for REST.
type RESTNamespaceUpdateRequest struct {
	Data map[string]interface{} `json:"namespace"`
}

//RESTCreateNamespace handle a Create REST service.
func (service *ContrailService) RESTCreateNamespace(c echo.Context) error {
	requestData := &models.CreateNamespaceRequest{
		Namespace: models.MakeNamespace(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "namespace",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateNamespace(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateNamespace handle a Create API
func (service *ContrailService) CreateNamespace(
	ctx context.Context,
	request *models.CreateNamespaceRequest) (*models.CreateNamespaceResponse, error) {
	model := request.Namespace
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
			return db.CreateNamespace(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "namespace",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateNamespaceResponse{
		Namespace: request.Namespace,
	}, nil
}

//RESTUpdateNamespace handles a REST Update request.
func (service *ContrailService) RESTUpdateNamespace(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateNamespaceRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "namespace",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateNamespace(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateNamespace handles a Update request.
func (service *ContrailService) UpdateNamespace(
	ctx context.Context,
	request *models.UpdateNamespaceRequest) (*models.UpdateNamespaceResponse, error) {
	model := request.Namespace
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateNamespace(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "namespace",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateNamespaceResponse{
		Namespace: model,
	}, nil
}

//RESTDeleteNamespace delete a resource using REST service.
func (service *ContrailService) RESTDeleteNamespace(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteNamespaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteNamespace(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteNamespace delete a resource.
func (service *ContrailService) DeleteNamespace(ctx context.Context, request *models.DeleteNamespaceRequest) (*models.DeleteNamespaceResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteNamespace(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteNamespaceResponse{
		ID: request.ID,
	}, nil
}

//RESTGetNamespace a REST Get request.
func (service *ContrailService) RESTGetNamespace(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetNamespaceRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetNamespace(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetNamespace a Get request.
func (service *ContrailService) GetNamespace(ctx context.Context, request *models.GetNamespaceRequest) (response *models.GetNamespaceResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListNamespaceRequest{
		Spec: spec,
	}
	var result *models.ListNamespaceResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListNamespace(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Namespaces) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetNamespaceResponse{
		Namespace: result.Namespaces[0],
	}
	return response, nil
}

//RESTListNamespace handles a List REST service Request.
func (service *ContrailService) RESTListNamespace(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListNamespaceRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListNamespace(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListNamespace handles a List service Request.
func (service *ContrailService) ListNamespace(
	ctx context.Context,
	request *models.ListNamespaceRequest) (response *models.ListNamespaceResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListNamespace(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
