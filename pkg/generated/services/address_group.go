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

//RESTAddressGroupUpdateRequest for update request for REST.
type RESTAddressGroupUpdateRequest struct {
	Data map[string]interface{} `json:"address-group"`
}

//RESTCreateAddressGroup handle a Create REST service.
func (service *ContrailService) RESTCreateAddressGroup(c echo.Context) error {
	requestData := &models.CreateAddressGroupRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "address_group",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreateAddressGroup(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreateAddressGroup handle a Create API
func (service *ContrailService) CreateAddressGroup(
	ctx context.Context,
	request *models.CreateAddressGroupRequest) (*models.CreateAddressGroupResponse, error) {
	model := request.AddressGroup
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
			return db.CreateAddressGroup(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "address_group",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateAddressGroupResponse{
		AddressGroup: request.AddressGroup,
	}, nil
}

//RESTUpdateAddressGroup handles a REST Update request.
func (service *ContrailService) RESTUpdateAddressGroup(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdateAddressGroupRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "address_group",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdateAddressGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdateAddressGroup handles a Update request.
func (service *ContrailService) UpdateAddressGroup(
	ctx context.Context,
	request *models.UpdateAddressGroupRequest) (*models.UpdateAddressGroupResponse, error) {
	model := request.AddressGroup
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.UpdateAddressGroup(ctx, tx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "address_group",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateAddressGroupResponse{
		AddressGroup: model,
	}, nil
}

//RESTDeleteAddressGroup delete a resource using REST service.
func (service *ContrailService) RESTDeleteAddressGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeleteAddressGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeleteAddressGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//DeleteAddressGroup delete a resource.
func (service *ContrailService) DeleteAddressGroup(ctx context.Context, request *models.DeleteAddressGroupRequest) (*models.DeleteAddressGroupResponse, error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			return db.DeleteAddressGroup(ctx, tx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteAddressGroupResponse{
		ID: request.ID,
	}, nil
}

//RESTGetAddressGroup a REST Get request.
func (service *ContrailService) RESTGetAddressGroup(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetAddressGroupRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetAddressGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//GetAddressGroup a Get request.
func (service *ContrailService) GetAddressGroup(ctx context.Context, request *models.GetAddressGroupRequest) (response *models.GetAddressGroupResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListAddressGroupRequest{
		Spec: spec,
	}
	var result *models.ListAddressGroupResponse
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListAddressGroup(ctx, tx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.AddressGroups) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetAddressGroupResponse{
		AddressGroup: result.AddressGroups[0],
	}
	return response, nil
}

//RESTListAddressGroup handles a List REST service Request.
func (service *ContrailService) RESTListAddressGroup(c echo.Context) error {
	var err error
	spec := common.GetListSpec(c)
	request := &models.ListAddressGroupRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListAddressGroup(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//ListAddressGroup handles a List service Request.
func (service *ContrailService) ListAddressGroup(
	ctx context.Context,
	request *models.ListAddressGroupRequest) (response *models.ListAddressGroupResponse, err error) {
	if err := common.DoInTransaction(
		service.DB,
		func(tx *sql.Tx) error {
			response, err = db.ListAddressGroup(ctx, tx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
