package apisrv

import (
	"fmt"

	"net/http"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/labstack/echo"
)

type fqNameToIDRequest struct {
	FQName []string `json:"fq_name"`
}

type fqNameToIDResponse struct {
	UUID string `json:"uuid"`
}

func (s *Server) fqNameToUUIDHandler(c echo.Context) error {
	fqNameToIDRequest := new(fqNameToIDRequest)
	ctx := c.Request().Context()

	err := c.Bind(fqNameToIDRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	fqName := fqNameToIDRequest.FQName
	metadata, err := s.dbService.GetMetaData(ctx, "", fqName)
	if err != nil {
		//TODO adding Project
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Name %s not found", basemodels.FQNameToString(fqName)))
	}

	//TODO permissions check

	fqNameToIDResponse := &fqNameToIDResponse{
		UUID: metadata.UUID,
	}
	return c.JSON(http.StatusOK, fqNameToIDResponse)
}
