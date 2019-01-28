package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// RESTFQNameToUUID is a REST handler for translating FQName to UUID
func (svc *ContrailService) RESTFQNameToUUID(c echo.Context) error {
	var request *FQNameToIDRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	response, err := svc.FQNameToID(c.Request().Context(), request)
	if err != nil {
		//TODO adding Project
		errMsg := fmt.Sprintf("Failed to retrieve metadata for FQName %v and Type %v", request.FQName, request.Type)
		return echo.NewHTTPError(http.StatusNotFound, errMsg)
	}
	return c.JSON(http.StatusOK, response)
}

// FQNameToID translates FQName to corresponding UUID stored in database
func (svc *ContrailService) FQNameToID(
	ctx context.Context,
	request *FQNameToIDRequest,
) (*FQNameToIDResponse, error) {

	// Try calling plugin first
	metadata, err := svc.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{Type: request.Type, FQName: request.FQName})
	if err != nil {
		errMsg := fmt.Sprintf("Failed to retrieve metadata for FQName %v and Type %v", request.FQName, request.Type)
		return nil, errors.Wrapf(err, errMsg)
	}

	//TODO permissions check

	return &FQNameToIDResponse{
		UUID: metadata.UUID,
	}, nil
}
