package apisrv

import (
	"fmt"
	"net/http"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/labstack/echo"
)

type fqNameToIDRequest struct {
	FQName []string `json:"fq_name"`
	Type   string   `json:"type"`
}

type fqNameToIDResponse struct {
	UUID string `json:"uuid"`
}

func (s *Server) fqNameToIDHandler(c echo.Context) error {
	fqNameToIDRequest := new(fqNameToIDRequest)
	ctx := c.Request().Context()

	err := c.Bind(fqNameToIDRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	fqName := fqNameToIDRequest.FQName
	metadata, err := s.dbService.TranslateBetweenFQNameUUID(ctx, "", fqName)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Name %s not found", models.FQNameToString(fqName)))
	}

	fqNameToIDResponse := &fqNameToIDResponse{
		UUID: metadata.UUID,
	}
	return c.JSON(http.StatusOK, fqNameToIDResponse)
}
