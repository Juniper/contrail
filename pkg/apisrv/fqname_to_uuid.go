package apisrv

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// FQNameToIDRequest defines fq_name to id request format
type FQNameToIDRequest struct {
	FQName []string `json:"fq_name"`
	Type   string   `json:"type"`
}

// FQNameToIDResponse defines FqNameToID response format.
type FQNameToIDResponse struct {
	UUID string `json:"uuid"`
}

func (s *Server) fqNameToUUIDHandler(c echo.Context) error {
	var request FQNameToIDRequest
	ctx := c.Request().Context()

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	metadata, err := s.DBService.GetMetadata(ctx, basemodels.Metadata{Type: request.Type, FQName: request.FQName})
	if err != nil {
		//TODO adding Project
		errMsg := fmt.Sprintf("Failed to retrieve metadata for FQName %v and Type %v", request.FQName, request.Type)
		return echo.NewHTTPError(http.StatusNotFound, errMsg)
	}

	//TODO permissions check

	fqNameToIDResponse := &FQNameToIDResponse{
		UUID: metadata.UUID,
	}
	return c.JSON(http.StatusOK, fqNameToIDResponse)
}
