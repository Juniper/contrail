package apisrv

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

func (s *Server) fqNameToUUIDHandler(c echo.Context) error {
	var request services.FQNameToIDRequest
	ctx := c.Request().Context()

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	metadata, err := s.dbService.GetMetadata(ctx, basemodels.Metadata{Type: request.Type, FQName: request.FQName})
	if err != nil {
		//TODO adding Project
		errMsg := fmt.Sprintf("Failed to retrieve metadata for FQName %v and Type %v", request.FQName, request.Type)
		return echo.NewHTTPError(http.StatusNotFound, errMsg)
	}

	//TODO permissions check

	fqNameToIDResponse := &services.FQNameToIDResponse{
		UUID: metadata.UUID,
	}
	return c.JSON(http.StatusOK, fqNameToIDResponse)
}
