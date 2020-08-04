package services

import (
	"net/http"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/auth"
	"github.com/labstack/echo"
)

// Endpoint path values.
const (
	ObjPermsPath = "/obj-perms"
)

// ObjPermsPlugin adds a obj-perms enpoint.
type ObjPermsPlugin struct{}

// RegisterHTTPAPI registers HTTP endpoints.
func (p *ObjPermsPlugin) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.GET(ObjPermsPath, p.RESTGetObjPerms)
}

// RegisterGRPCAPI does nothing.
func (p *ObjPermsPlugin) RegisterGRPCAPI(r apiserver.GRPCRouter) {}

// RESTGetObjPerms handles GET operation of obj-perms request.
func (p *ObjPermsPlugin) RESTGetObjPerms(c echo.Context) error {
	return c.JSON(http.StatusOK, auth.GetIdentity(c.Request().Context()).GetObjPerms())
}
