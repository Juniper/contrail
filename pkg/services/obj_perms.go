package services

import (
	"net/http"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/labstack/echo"
)

// RESTGetObjPerms handles GET operation of obj-perms request.
func (service *ContrailService) RESTGetObjPerms(c echo.Context) error {
	return c.JSON(http.StatusOK, auth.GetIdentity(c.Request().Context()).GetObjPerms())
}
