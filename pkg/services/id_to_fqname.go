package services

import (
	"context"

	"github.com/pkg/errors"

	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// RESTIDToFQName is a REST handler for translating UUID to FQName and Type
func (svc *ContrailService) RESTIDToFQName(c echo.Context) error {
	var request *IDToFQNameRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	response, err := svc.IDToFQName(c.Request().Context(), request)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to retrieve metadata for UUID %v", request.UUID)
		return echo.NewHTTPError(http.StatusNotFound, errMsg)
	}
	return c.JSON(http.StatusOK, response)
}

// IDToFQName translates UUID to corresponding FQName and Type stored in database
func (svc *ContrailService) IDToFQName(
	ctx context.Context,
	request *IDToFQNameRequest,
) (*IDToFQNameResponse, error) {
	metadata, err := svc.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{UUID: request.UUID})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to retrieve metadata for UUID %v", request.UUID)
	}

	return &IDToFQNameResponse{
		Type:   metadata.Type,
		FQName: metadata.FQName,
	}, nil
}
