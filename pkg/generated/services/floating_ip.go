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

//RESTFloatingIPUpdateRequest for update request for REST.
type RESTFloatingIPUpdateRequest struct {
	Data map[string]interface{} `json:"floating-ip"`
}

//RESTCreateFloatingIP handle a Create REST service.
func (service *ContrailService) RESTCreateFloatingIP(c echo.Context) error {
	requestData := &models.CreateFloatingIPRequest{
		FloatingIP: models.MakeFloatingIP(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateFloatingIP(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateFloatingIP handle a Create API
func (service *ContrailService) CreateFloatingIP(
	ctx context.Context,
	request *models.CreateFloatingIPRequest) (*models.CreateFloatingIPResponse, error) {
	model := request.FloatingIP
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
			return db.CreateFloatingIP(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateFloatingIPResponse{
		FloatingIP: request.FloatingIP,
	}, nil
}

//RESTUpdateFloatingIP handles a REST Update request.
func (service *ContrailService) RESTUpdateFloatingIP(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateFloatingIPRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateFloatingIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateFloatingIP handles a Update request.
func (service *ContrailService) UpdateFloatingIP(
	ctx context.Context,
	request *models.UpdateFloatingIPRequest) (*models.UpdateFloatingIPResponse, error) {
	model := request.FloatingIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateFloatingIP(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "floating_ip",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateFloatingIPResponse{
		FloatingIP: model,
	}, nil
}

//RESTDeleteFloatingIP delete a resource using REST service.
func (service *ContrailService) RESTDeleteFloatingIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteFloatingIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteFloatingIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteFloatingIP delete a resource.
func (service *ContrailService) DeleteFloatingIP(ctx context.Context, request *models.DeleteFloatingIPRequest) (*models.DeleteFloatingIPResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteFloatingIP(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteFloatingIPResponse{
		ID: request.ID,
	}, nil
}

//RESTGetFloatingIP a REST Get request.
func (service *ContrailService) RESTGetFloatingIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetFloatingIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetFloatingIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetFloatingIP a Get request.
func (service *ContrailService) GetFloatingIP(ctx context.Context, request *models.GetFloatingIPRequest) (response *models.GetFloatingIPResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListFloatingIPRequest{
		Spec: spec,
	}
	var result *models.ListFloatingIPResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListFloatingIP(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.FloatingIPs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetFloatingIPResponse{
		FloatingIP: result.FloatingIPs[0],
	}
	return response, nil
}

//RESTListFloatingIP handles a List REST service Request.
func (service *ContrailService) RESTListFloatingIP(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListFloatingIPRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListFloatingIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListFloatingIP handles a List service Request.
func (service *ContrailService) ListFloatingIP(
	ctx context.Context,
	request *models.ListFloatingIPRequest) (response *models.ListFloatingIPResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListFloatingIP(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
