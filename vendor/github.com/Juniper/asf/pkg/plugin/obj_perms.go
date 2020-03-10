package plugin

import (
	"net/http"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/auth"
	"github.com/labstack/echo"
)

const (
	ObjPerms = "obj-perms"
)

type ObjPermsPlugin struct{}

// RESTGetObjPerms handles GET operation of obj-perms request.
func (p *ObjPermsPlugin) RESTGetObjPerms(c echo.Context) error {
	return c.JSON(http.StatusOK, auth.GetIdentity(c.Request().Context()).GetObjPerms())
}

// RegisterHTTPAPI registers HTTP services.
func (p *ObjPermsPlugin) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.GET(ObjPerms, p.RESTGetObjPerms)
}

// RegisterGRPCAPI registers GRPC services.
func (p *ObjPermsPlugin) RegisterGRPCAPI(r apiserver.GRPCRouter) {
	return
}
