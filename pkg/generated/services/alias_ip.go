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

//RESTAliasIPUpdateRequest for update request for REST.
type RESTAliasIPUpdateRequest struct {
	Data map[string]interface{} `json:"alias-ip"`
}

//RESTCreateAliasIP handle a Create REST service.
func (service *ContrailService) RESTCreateAliasIP(c echo.Context) error {
	requestData := &models.CreateAliasIPRequest{
		AliasIP: models.MakeAliasIP(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAliasIP(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAliasIP handle a Create API
func (service *ContrailService) CreateAliasIP(
	ctx context.Context,
	request *models.CreateAliasIPRequest) (*models.CreateAliasIPResponse, error) {
	model := request.AliasIP
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
			return db.CreateAliasIP(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAliasIPResponse{
		AliasIP: request.AliasIP,
	}, nil
}

//RESTUpdateAliasIP handles a REST Update request.
func (service *ContrailService) RESTUpdateAliasIP(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAliasIPRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAliasIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAliasIP handles a Update request.
func (service *ContrailService) UpdateAliasIP(
	ctx context.Context,
	request *models.UpdateAliasIPRequest) (*models.UpdateAliasIPResponse, error) {
	model := request.AliasIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateAliasIP(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "alias_ip",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAliasIPResponse{
		AliasIP: model,
	}, nil
}

//RESTDeleteAliasIP delete a resource using REST service.
func (service *ContrailService) RESTDeleteAliasIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAliasIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAliasIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteAliasIP delete a resource.
func (service *ContrailService) DeleteAliasIP(ctx context.Context, request *models.DeleteAliasIPRequest) (*models.DeleteAliasIPResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteAliasIP(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAliasIPResponse{
		ID: request.ID,
	}, nil
}

//RESTGetAliasIP a REST Get request.
func (service *ContrailService) RESTGetAliasIP(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAliasIPRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAliasIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetAliasIP a Get request.
func (service *ContrailService) GetAliasIP(ctx context.Context, request *models.GetAliasIPRequest) (response *models.GetAliasIPResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filter: models.Filter{
			"uuid": []string{request.ID},
		},
	}
	listRequest := &models.ListAliasIPRequest{
		Spec: spec,
	}
	var result *models.ListAliasIPResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListAliasIP(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.AliasIPs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAliasIPResponse{
		AliasIP: result.AliasIPs[0],
	}
	return response, nil
}

//RESTListAliasIP handles a List REST service Request.
func (service *ContrailService) RESTListAliasIP(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAliasIPRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAliasIP(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListAliasIP handles a List service Request.
func (service *ContrailService) ListAliasIP(
	ctx context.Context,
	request *models.ListAliasIPRequest) (response *models.ListAliasIPResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListAliasIP(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
