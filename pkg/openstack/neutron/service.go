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
	Create(CreateRequest) (models.Response, error)
	Update(UpdateRequest) (models.Response, error)
	Delete(DeleteRequest) (models.Response, error)
	Read(ReadRequest) (models.Response, error)
	ReadAll(ReadAllRequest) (models.Response, error)
	ReadCount(ReadCountRequest) (models.Response, error)
	AddInterface(AddInterfaceRequest) (models.Response, error)
	DeleteInterface(DeleteInterfaceRequest) (models.Response, error)
}

func (s *Service) Create(r *CreateRequest) (models.Response, error) {
	return r.resource.Create(r.ctx, s.readService, s.writeService)
}

func (s *Service) Update(r *UpdateRequest) (models.Response, error) {
	return r.resource.Update(r.ctx, s.readService, s.writeService)
}

func (s *Service) Delete(r *DeleteRequest) (models.Response, error) {
	// TODO
	return nil, nil
}

func (s *Service) AddInterface(r *AddInterfaceRequest) (models.Response, error) {
	return r.resource.AddInterface(r.ctx, s.readService, s.writeService)
}

func (s *Service) DeleteInterface(r *DeleteInterfaceRequest) (models.Response, error) {
	return r.resource.DeleteInterface(r.ctx, s.readService, s.writeService)
}

func (s *Service) Read(r *ReadRequest) (models.Response, error) {
	// TODO
	return nil, nil
}

func (s *Service) ReadAll(r *ReadAllRequest) (models.Response, error) {
	// TODO
	return nil, nil
}

func (s *Service) ReadCount(r *ReadCountRequest) (models.Response, error) {
	// TODO
	return nil, nil
}

func (s *Service) Process(r *Request) (models.Response, error) {
	switch r.Context.Operation {
	case "CREATE":
		createRequest := &CreateRequest{
			ctx:      r.Context,
			resource: r.Data.Resource,
		}
		s.Create(createRequest)
	case "UPDATE":
		updateRequest := &UpdateRequest{
			ctx:      r.Context,
			resource: r.Data.Resource,
			id:       r.Data.ID,
		}
		s.Update(updateRequest)
	case "DELETE":
		deleteRequest := &DeleteRequest{
			ctx: r.Context,
			id:  r.Data.ID,
		}
		s.Delete(deleteRequest)
	case "READ":
		readRequest := &ReadRequest{
			ctx: r.Context,
			id:  r.Data.ID,
		}
		s.Read(readRequest)
	case "READALL":
		readAllRequest := &ReadAllRequest{
			ctx: r.Context,
			filters:  r.Data.Filters,
		}
		s.ReadAll(readAllRequest)
	case "READCOUNT":
		readCountRequest := &ReadCountRequest{
			ctx: r.Context,
			filters:  r.Data.Filters,
		}
		s.ReadCount(readCountRequest)
	case "ADDINTERFACE":
		addInterfaceRequest := &AddInterfaceRequest{
			ctx:      r.Context,
			resource: r.Data.Resource,
		}
		s.AddInterface(addInterfaceRequest)
	case "DELINTERFACE":
		deleteInterfaceRequest := &DeleteInterfaceRequest{
			ctx:      r.Context,
			resource: r.Data.Resource,
		}
		s.DeleteInterface(deleteInterfaceRequest)

	}
}

type routeRegistry interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}
