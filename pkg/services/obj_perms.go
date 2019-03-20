package services

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/auth"
)

func (service *ContrailService) RESTGetObjPerms(c echo.Context) error {
	return c.JSON(http.StatusOK, auth.GetContext(c).GetObjPerms())
}
