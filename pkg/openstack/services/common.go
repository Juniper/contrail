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

type Neutron interface {
	Create(CreateRequest) Response
	Update(UpdateRequest) Response
	Delete(DeleteRequest) Response
	Read(ReadRequest) Response
	ReadAll(ReadAllRequest) Response
	ReadCount(ReadCountRequest) Response
	AddInterface(AddInterfaceRequest) Response
	RemoveInterface(RemoveInterfaceRequest) Response
}

type CreateRequest struct {
	ctx models.RequestContext
	r   Resource
}

type UpdateRequest struct {
	ctx models.RequestContext
	r   Resource
}

type DeleteRequest struct {
	ctx models.RequestContext
	r   Resource
}

type ReadRequest struct {
	ctx    models.RequestContext
	fields Fields
	id     string
}

type ReadAllRequest struct {
	ctx     models.RequestContext
	filters Filters
}

type ReadCountRequest struct {
	ctx models.RequestContext
}

type AddInterfaceRequest struct {
	ctx models.RequestContext
	r   Resource
}

type RemoveInterfaceRequest struct {
	ctx models.RequestContext
	r   Resource
}

type Resource interface {
	Create(ctx models.RequestContext, r services.ReadService, w services.WriteService) Response
	Update(ctx models.RequestContext, r services.ReadService, w services.WriteService) Response
	Delete(ctx models.RequestContext, r services.ReadService, w services.WriteService) Response
	AddInterface(ctx models.RequestContext, r services.ReadService, w services.WriteService) Response
	RemoveInterface(ctx models.RequestContext, r services.ReadService, w services.WriteService) Response
}

type Filters map[string][]string
type Fields []string

type Data struct {
	Filters  Filters
	ID       string
	Fields   []string
	Resource Resource
}

type Request struct {
	Data    Data
	Context models.RequestContext
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
