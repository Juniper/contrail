package apisrv

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/services"
)

func (s *Server) fqNameToUUIDHandler(c echo.Context) error {
	var request *services.FQNameToIDRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}

	response, err := s.FQNameToIDServer.FQNameToID(c.Request().Context(), request)
	if err != nil {
		//TODO adding Project
		errMsg := fmt.Sprintf("Failed to retrieve metadata for FQName %v and Type %v", request.FQName, request.Type)
		return echo.NewHTTPError(http.StatusNotFound, errMsg)
	}
	return c.JSON(http.StatusOK, response)
}
