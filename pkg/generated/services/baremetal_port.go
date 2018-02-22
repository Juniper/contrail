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

//RESTBaremetalPortUpdateRequest for update request for REST.
type RESTBaremetalPortUpdateRequest struct {
	Data map[string]interface{} `json:"baremetal-port"`
}

//RESTCreateBaremetalPort handle a Create REST service.
func (service *ContrailService) RESTCreateBaremetalPort(c echo.Context) error {
	requestData := &models.CreateBaremetalPortRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_port",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateBaremetalPort(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateBaremetalPort handle a Create API
func (service *ContrailService) CreateBaremetalPort(
	ctx context.Context,
	request *models.CreateBaremetalPortRequest) (*models.CreateBaremetalPortResponse, error) {
	model := request.BaremetalPort
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
			return db.CreateBaremetalPort(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_port",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBaremetalPortResponse{
		BaremetalPort: request.BaremetalPort,
	}, nil
}

//RESTUpdateBaremetalPort handles a REST Update request.
func (service *ContrailService) RESTUpdateBaremetalPort(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateBaremetalPortRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_port",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateBaremetalPort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateBaremetalPort handles a Update request.
func (service *ContrailService) UpdateBaremetalPort(
	ctx context.Context,
	request *models.UpdateBaremetalPortRequest) (*models.UpdateBaremetalPortResponse, error) {
	model := request.BaremetalPort
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateBaremetalPort(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_port",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBaremetalPortResponse{
		BaremetalPort: model,
	}, nil
}

//RESTDeleteBaremetalPort delete a resource using REST service.
func (service *ContrailService) RESTDeleteBaremetalPort(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteBaremetalPortRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteBaremetalPort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteBaremetalPort delete a resource.
func (service *ContrailService) DeleteBaremetalPort(ctx context.Context, request *models.DeleteBaremetalPortRequest) (*models.DeleteBaremetalPortResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteBaremetalPort(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBaremetalPortResponse{
		ID: request.ID,
	}, nil
}

//RESTGetBaremetalPort a REST Get request.
func (service *ContrailService) RESTGetBaremetalPort(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetBaremetalPortRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetBaremetalPort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetBaremetalPort a Get request.
func (service *ContrailService) GetBaremetalPort(ctx context.Context, request *models.GetBaremetalPortRequest) (response *models.GetBaremetalPortResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListBaremetalPortRequest{
		Spec: spec,
	}
	var result *models.ListBaremetalPortResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListBaremetalPort(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BaremetalPorts) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBaremetalPortResponse{
		BaremetalPort: result.BaremetalPorts[0],
	}
	return response, nil
}

//RESTListBaremetalPort handles a List REST service Request.
func (service *ContrailService) RESTListBaremetalPort(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListBaremetalPortRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListBaremetalPort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListBaremetalPort handles a List service Request.
func (service *ContrailService) ListBaremetalPort(
	ctx context.Context,
	request *models.ListBaremetalPortRequest) (response *models.ListBaremetalPortResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListBaremetalPort(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
