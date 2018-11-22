// Code generated by contrailschema tool from template service.tmpl; DO NOT EDIT.

package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	"github.com/Juniper/contrail/extension/pkg/models"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// This is needed to prevent an import error.
var _ = fmt.Println

type RESTUpdateSampleRequest struct {
	Sample map[string]interface{} `json:"sample"`
}

// RESTCreateSample handles a REST create request.
func (service *ContrailService) RESTCreateSample(c echo.Context) error {
	requestData := &CreateSampleRequest{}
	if err := c.Bind(requestData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %s", err))
	}
	ctx := c.Request().Context()
	response, err := service.CreateSample(ctx, requestData)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

// CreateSample handles a create request.
func (service *ContrailService) CreateSample(
	ctx context.Context, request *CreateSampleRequest,
) (*CreateSampleResponse, error) {
	model := request.Sample
	if model == nil {
		return nil, errutil.ErrorBadRequest("create body is empty")
	}
	auth := auth.GetAuthCTX(ctx)
	if auth == nil {
		return nil, errutil.ErrorUnauthenticated
	}
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}

	if model.Name == "" {
		if fqName := model.FQName; len(fqName) > 0 {
			model.Name = fqName[len(fqName)-1]
		} else {
			model.Name = "default-sample"
		}
	}

	if err := service.sanitizeFQNameForSample(ctx, request); err != nil {
		return nil, err
	}

	model.Perms2 = &models.PermType2{
		OwnerAccess: basemodels.PermsRWX,
	}
	model.Perms2.Owner = auth.ProjectID()

	if model.IDPerms == nil {
		model.IDPerms = models.NewIDPerms(model.UUID)
	}

	if model.IDPerms.UUID == nil {
		model.IDPerms.UUID = models.NewUUIDType(model.UUID)
	}

	err := service.TypeValidator.ValidateSample(request.Sample)
	if err != nil {
		return nil, errutil.ErrorBadRequestf(
			"validation failed for resource with UUID %v: %v",
			request.Sample.UUID,
			err,
		)
	}

	return service.Next().CreateSample(ctx, request)
}

func (service *ContrailService) sanitizeFQNameForSample(
	ctx context.Context,
	request *CreateSampleRequest,
) error {
	model := request.Sample
	if len(model.FQName) != 0 {
		return nil
	}

	model.FQName = []string{model.Name}

	return nil
}

// RESTUpdateSample handles a REST update request.
func (service *ContrailService) RESTUpdateSample(c echo.Context) error {
	var request RESTUpdateSampleRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %s", err))
	}

	model := models.InterfaceToSample(request.Sample)
	model.UUID = c.Param("id")

	response, err := service.UpdateSample(
		c.Request().Context(),
		&UpdateSampleRequest{
			Sample:    model,
			FieldMask: basemodels.MapToFieldMask(request.Sample),
		},
	)
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, response)
}

// UpdateSample handles an update request.
func (service *ContrailService) UpdateSample(
	ctx context.Context, request *UpdateSampleRequest,
) (*UpdateSampleResponse, error) {
	model := request.Sample
	if model == nil {
		return nil, errutil.ErrorBadRequest("update request body is empty")
	}

	model.IDPerms, request.FieldMask.Paths = sanitizeIDPermsUUID(model.GetIDPerms(), request.FieldMask.Paths)

	return service.Next().UpdateSample(ctx, request)
}

// RESTDeleteSample handles a REST delete request.
func (service *ContrailService) RESTDeleteSample(c echo.Context) error {
	request := &DeleteSampleRequest{
		ID: c.Param("id"),
	}
	ctx := c.Request().Context()
	_, err := service.DeleteSample(ctx, request)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.NoContent(http.StatusOK)
}

// RESTGetSample handles a REST get request.
func (service *ContrailService) RESTGetSample(c echo.Context) error {
	request := &GetSampleRequest{
		ID: c.Param("id"),
	}
	ctx := c.Request().Context()
	response, err := service.GetSample(ctx, request)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

// RESTListSample handles a REST list request.
func (service *ContrailService) RESTListSample(c echo.Context) error {
	var err error
	spec := baseservices.GetListSpec(c)
	request := &ListSampleRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListSample(ctx, request)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

type RESTUpdateTagRequest struct {
	Tag map[string]interface{} `json:"tag"`
}

// RESTCreateTag handles a REST create request.
func (service *ContrailService) RESTCreateTag(c echo.Context) error {
	requestData := &CreateTagRequest{}
	if err := c.Bind(requestData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %s", err))
	}
	ctx := c.Request().Context()
	response, err := service.CreateTag(ctx, requestData)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

// CreateTag handles a create request.
func (service *ContrailService) CreateTag(
	ctx context.Context, request *CreateTagRequest,
) (*CreateTagResponse, error) {
	model := request.Tag
	if model == nil {
		return nil, errutil.ErrorBadRequest("create body is empty")
	}
	auth := auth.GetAuthCTX(ctx)
	if auth == nil {
		return nil, errutil.ErrorUnauthenticated
	}
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}

	if model.Name == "" {
		if fqName := model.FQName; len(fqName) > 0 {
			model.Name = fqName[len(fqName)-1]
		} else {
			model.Name = "default-tag"
		}
	}

	if err := service.sanitizeFQNameForTag(ctx, request); err != nil {
		return nil, err
	}

	model.Perms2 = &models.PermType2{
		OwnerAccess: basemodels.PermsRWX,
	}
	model.Perms2.Owner = auth.ProjectID()

	if model.IDPerms == nil {
		model.IDPerms = models.NewIDPerms(model.UUID)
	}

	if model.IDPerms.UUID == nil {
		model.IDPerms.UUID = models.NewUUIDType(model.UUID)
	}

	err := service.TypeValidator.ValidateTag(request.Tag)
	if err != nil {
		return nil, errutil.ErrorBadRequestf(
			"validation failed for resource with UUID %v: %v",
			request.Tag.UUID,
			err,
		)
	}

	return service.Next().CreateTag(ctx, request)
}

func (service *ContrailService) sanitizeFQNameForTag(
	ctx context.Context,
	request *CreateTagRequest,
) error {
	model := request.Tag
	if len(model.FQName) != 0 {
		return nil
	}

	model.FQName = []string{model.Name}

	return nil
}

// RESTUpdateTag handles a REST update request.
func (service *ContrailService) RESTUpdateTag(c echo.Context) error {
	var request RESTUpdateTagRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %s", err))
	}

	model := models.InterfaceToTag(request.Tag)
	model.UUID = c.Param("id")

	response, err := service.UpdateTag(
		c.Request().Context(),
		&UpdateTagRequest{
			Tag:       model,
			FieldMask: basemodels.MapToFieldMask(request.Tag),
		},
	)
	if err != nil {
		return errutil.ToHTTPError(err)
	}

	return c.JSON(http.StatusOK, response)
}

// UpdateTag handles an update request.
func (service *ContrailService) UpdateTag(
	ctx context.Context, request *UpdateTagRequest,
) (*UpdateTagResponse, error) {
	model := request.Tag
	if model == nil {
		return nil, errutil.ErrorBadRequest("update request body is empty")
	}

	model.IDPerms, request.FieldMask.Paths = sanitizeIDPermsUUID(model.GetIDPerms(), request.FieldMask.Paths)

	return service.Next().UpdateTag(ctx, request)
}

// RESTDeleteTag handles a REST delete request.
func (service *ContrailService) RESTDeleteTag(c echo.Context) error {
	request := &DeleteTagRequest{
		ID: c.Param("id"),
	}
	ctx := c.Request().Context()
	_, err := service.DeleteTag(ctx, request)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.NoContent(http.StatusOK)
}

// RESTGetTag handles a REST get request.
func (service *ContrailService) RESTGetTag(c echo.Context) error {
	request := &GetTagRequest{
		ID: c.Param("id"),
	}
	ctx := c.Request().Context()
	response, err := service.GetTag(ctx, request)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

// RESTListTag handles a REST list request.
func (service *ContrailService) RESTListTag(c echo.Context) error {
	var err error
	spec := baseservices.GetListSpec(c)
	request := &ListTagRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListTag(ctx, request)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

func sanitizeIDPermsUUID(idPerms *models.IdPermsType, paths []string) (*models.IdPermsType, []string) {
	if idPerms.GetUUID() != nil {
		idPerms.UUID = nil
	}

	return idPerms, format.RemoveFromStringSlice(
		paths,
		map[string]struct{}{
			"id_perms.uuid":             {},
			"id_perms.uuid.uuid_mslong": {},
			"id_perms.uuid.uuid_lslong": {},
		},
	)
}
