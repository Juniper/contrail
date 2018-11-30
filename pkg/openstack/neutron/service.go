package neutron

import (
	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/openstack/models"
	"github.com/Juniper/contrail/pkg/services"
)

// Service implementation.
type Service struct {
	readService  services.ReadService
	writeService services.WriteService
}

type Neutron interface {
	Create(CreateRequest) models.Response
	Update(UpdateRequest) models.Response
	Delete(DeleteRequest) models.Response
	Read(ReadRequest) models.Response
	ReadAll(ReadAllRequest) models.Response
	ReadCount(ReadCountRequest) models.Response
	AddInterface(AddInterfaceRequest) models.Response
	RemoveInterface(RemoveInterfaceRequest) models.Response
}

type routeRegistry interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}
