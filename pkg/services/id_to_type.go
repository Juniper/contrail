package services

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// // RESTFQNameToUUID is a REST handler for translating FQName to UUID
// func (svc *ContrailService) RESTFQNameToUUID(c echo.Context) error {
// 	var request *FQNameToIDRequest

// 	err := c.Bind(&request)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
// 	}

// 	response, err := svc.FQNameToID(c.Request().Context(), request)
// 	if err != nil {
// 		//TODO adding Project
// 		errMsg := fmt.Sprintf("Failed to retrieve metadata for FQName %v and Type %v", request.FQName, request.Type)
// 		return echo.NewHTTPError(http.StatusNotFound, errMsg)
// 	}
// 	return c.JSON(http.StatusOK, response)
// }

// IDToType translates UUID to corresponding type stored in database
func (svc *ContrailService) IDToType(
	ctx context.Context,
	request *IDToTypeRequest,
) (*IDToTypeResponse, error) {
	metadata, err := svc.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{UUID: request.UUID})
	if err != nil {
		errMsg := fmt.Sprintf("Failed to retrieve metadata for UUID %v", request.UUID)
		return nil, errors.Wrapf(err, errMsg)
	}

	return &IDToTypeResponse{
		Type: metadata.Type,
	}, nil
}