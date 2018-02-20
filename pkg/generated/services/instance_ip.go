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

//RESTInstanceIPUpdateRequest for update request for REST.
type RESTInstanceIPUpdateRequest struct {
	Data map[string]interface{} `json:"instance-ip"`
}

//RESTCreateInstanceIP handle a Create REST service.
func (service *ContrailService) RESTCreateInstanceIP(c echo.Context) error {
	requestData := &models.CreateInstanceIPRequest{
		InstanceIP: models.MakeInstanceIP(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "instance_ip",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateInstanceIP(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateInstanceIP handle a Create API
func (service *ContrailService) CreateInstanceIP(
	ctx context.Context,
	request *models.CreateInstanceIPRequest) (*models.CreateInstanceIPResponse, error) {
	model := request.InstanceIP
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
			return db.CreateInstanceIP(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "instance_ip",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateInstanceIPResponse{
		InstanceIP: request.InstanceIP,
	}, nil
}

//RESTUpdateInstanceIP handles a REST Update request.
func (service *ContrailService) RESTUpdateInstanceIP(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateInstanceIPRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "instance_ip",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateInstanceIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateInstanceIP handles a Update request.
func (service *ContrailService) UpdateInstanceIP(
	ctx context.Context,
	request *models.UpdateInstanceIPRequest) (*models.UpdateInstanceIPResponse, error) {
	model := request.InstanceIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateInstanceIP(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "instance_ip",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateInstanceIPResponse{
		InstanceIP: model,
	}, nil
}

//RESTDeleteInstanceIP delete a resource using REST service.
func (service *ContrailService) RESTDeleteInstanceIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteInstanceIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteInstanceIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteInstanceIP delete a resource.
func (service *ContrailService) DeleteInstanceIP(ctx context.Context, request *models.DeleteInstanceIPRequest) (*models.DeleteInstanceIPResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteInstanceIP(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteInstanceIPResponse{
		ID: request.ID,
	}, nil
}

//RESTGetInstanceIP a REST Get request.
func (service *ContrailService) RESTGetInstanceIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetInstanceIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetInstanceIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetInstanceIP a Get request.
func (service *ContrailService) GetInstanceIP(ctx context.Context, request *models.GetInstanceIPRequest) (response *models.GetInstanceIPResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListInstanceIPRequest{
		Spec: spec,
	}
	var result *models.ListInstanceIPResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListInstanceIP(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.InstanceIPs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetInstanceIPResponse{
		InstanceIP: result.InstanceIPs[0],
	}
	return response, nil
}

//RESTListInstanceIP handles a List REST service Request.
func (service *ContrailService) RESTListInstanceIP(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListInstanceIPRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListInstanceIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListInstanceIP handles a List service Request.
func (service *ContrailService) ListInstanceIP(
	ctx context.Context,
	request *models.ListInstanceIPRequest) (response *models.ListInstanceIPResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListInstanceIP(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
