package services

import (
	"net/http"

	"github.com/labstack/echo"
	// TODO(dfurman): Decouple from below packages
	//"github.com/Juniper/asf/pkg/auth"
)

// RESTGetObjPerms handles GET operation of obj-perms request.
func (service *ContrailService) RESTGetObjPerms(c echo.Context) error {
	return c.JSON(http.StatusOK, auth.GetContext(c).GetObjPerms())
}
