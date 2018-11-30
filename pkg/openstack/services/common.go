package services

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

type Data struct {
	Filters  map[string][]string
	ID       string
	Fields   []string
	Resource Resource
}

type Request struct {
	data    Data
	context models.RequestContext
}

type Resource interface {
	Create(ctx models.RequestContext, r services.ReadService, w services.WriteService) Response
	Update(ctx models.RequestContext, r services.ReadService, w services.WriteService) Response
	Delete(ctx models.RequestContext, r services.ReadService, w services.WriteService) Response
	AddInterface(ctx models.RequestContext, r services.ReadService, w services.WriteService) Response
	RemoveInterface(ctx models.RequestContext, r services.ReadService, w services.WriteService) Response
}

type BaseResource struct{}

func (b *BaseResource) Create(
	ctx models.RequestContext,
	r services.ReadService,
	w services.WriteService,
) Response {
	return nil
}

func (b *BaseResource) Update(
	ctx models.RequestContext,
	r services.ReadService,
	w services.WriteService,
) Response {
	return nil
}

func (b *BaseResource) Delete(
	ctx models.RequestContext,
	r services.ReadService,
	w services.WriteService,
) Response {
	return nil
}

func (b *BaseResource) AddInterface(
	ctx models.RequestContext,
	r services.ReadService,
	w services.WriteService,
) Response {
	return nil
}

func (b *BaseResource) RemoveInterface(
	ctx models.RequestContext,
	r services.ReadService,
	w services.WriteService,
) Response {
	return nil
}

type Response interface{}

type routeRegistry interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}
