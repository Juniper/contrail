package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// Endpoint paths and names.
const (
	FQNameToIDPath = "fqname-to-id"
	FQNameToIDName = "name-to-id"

	IDToFQNamePath = "id-to-fqname"
	IDToFQNameName = "id-to-name"
)

// FQNameTranslationPlugin is a plugin that adds `fqname-to-id` and `id-to-fqname` endpoints.
type FQNameTranslationPlugin struct {
	MetadataGetter MetadataGetter
}

// RegisterHTTPAPI registers HTTP endpoints.
func (p *FQNameTranslationPlugin) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.POST(FQNameToIDPath, p.RESTFQNameToUUID, apiserver.WithHomepageName(FQNameToIDName))
	r.POST(IDToFQNamePath, p.RESTIDToFQName, apiserver.WithHomepageName(IDToFQNameName))
}

// RegisterGRPCAPI registers GRPC endpoints.
func (p *FQNameTranslationPlugin) RegisterGRPCAPI(r apiserver.GRPCRouter) {
	r.RegisterService(&_FQNameToID_serviceDesc, p)
	r.RegisterService(&_IDToFQName_serviceDesc, p)
}

// RESTFQNameToUUID is a REST handler for translating FQName to UUID
func (p *FQNameTranslationPlugin) RESTFQNameToUUID(c echo.Context) error {
	var request *FQNameToIDRequest

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	response, err := p.FQNameToID(c.Request().Context(), request)
	if err != nil {
		//TODO adding Project
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

// FQNameToID translates FQName to corresponding UUID stored in database
func (p *FQNameTranslationPlugin) FQNameToID(
	ctx context.Context, request *FQNameToIDRequest,
) (*FQNameToIDResponse, error) {
	metadata, err := p.MetadataGetter.GetMetadata(ctx, models.FQNameMetadata(request.FQName, request.Type))
	if err != nil {
		//TODO adding Project
		errMsg := fmt.Sprintf("Failed to retrieve metadata for FQName %v and Type %v", request.FQName, request.Type)
		return nil, errors.Wrapf(err, errMsg)
	}

	//TODO permissions check

	return &FQNameToIDResponse{UUID: metadata.UUID}, nil
}

// RESTIDToFQName is a REST handler for translating UUID to FQName and Type:ta
func (p *FQNameTranslationPlugin) RESTIDToFQName(c echo.Context) error {
	var request *IDToFQNameRequest

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %v", err))
	}

	response, err := p.IDToFQName(c.Request().Context(), request)
	if err != nil {
		return errutil.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

// IDToFQName translates UUID to corresponding FQName and Type stored in database
func (p *FQNameTranslationPlugin) IDToFQName(
	ctx context.Context, request *IDToFQNameRequest,
) (*IDToFQNameResponse, error) {
	metadata, err := p.MetadataGetter.GetMetadata(ctx, models.UUIDMetadata(request.UUID))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve metadata for UUID %v", request.UUID)
	}

	return &IDToFQNameResponse{Type: metadata.Type, FQName: metadata.FQName}, nil
}
