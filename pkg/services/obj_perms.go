package services

import (
	"net/http"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/labstack/echo"
)

func (service *ContrailService) RESTGetObjPerms(c echo.Context) error {
	// ctx := c.Request().Context()
	objPerms := auth.NewObjPerms(nil)
	// TODO: this is only example response for debugging. Replace it with the proper one.
	return c.JSON(http.StatusOK, objPerms)
}
