package services

import (
	"github.com/labstack/echo"
	"net/http"
)

func (service *ContrailService) RESTGetObjPerms(c echo.Context) error {
	ctx := c.Request().Context()
	// TODO: this is only example response for debugging. Replace it with the proper one.
	return c.JSON(http.StatusOK, ctx)
}
