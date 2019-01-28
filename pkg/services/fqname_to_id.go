package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/errutil"
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
	metadata, err := svc.MetadataGetter.GetMetadata(ctx, basemodels.Metadata{Type: request.Type, FQName: request.FQName})
	if errutil.IsNotFound(err) {
		for _, p := range svc.Plugins.FqNameToIDPlugins {
			res, err := p.FQNameToID(ctx, request)
			if err != nil {
				log.WithError(err).Error("plugin returned error")
			}
			if res != nil {
				return res, nil
			}
		}
	}

	if err != nil {
		errMsg := fmt.Sprintf("Failed to retrieve metadata for FQName %v and Type %v", request.FQName, request.Type)
		return nil, errors.Wrapf(err, errMsg)
	}

	//TODO permissions check

	return &FQNameToIDResponse{
		UUID: metadata.UUID,
	}, nil
}
