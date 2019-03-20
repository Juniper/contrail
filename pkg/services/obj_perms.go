package services

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/auth"
)

func (service *ContrailService) RESTGetObjPerms(c echo.Context) error {
	ctx := c.Request().Context()
	var authKey interface{} = "auth"
	authContext := ctx.Value(authKey).(*auth.Context)
	return c.JSON(http.StatusOK, authContext.GetObjPerms())
}
