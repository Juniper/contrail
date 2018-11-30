package neutron

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/openstack/models"
	"github.com/Juniper/contrail/pkg/openstack/neutronerrors"
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
		return s.Create(createRequest)
	case "UPDATE":
		updateRequest := &UpdateRequest{
			ctx:      r.Context,
			resource: r.Data.Resource,
			id:       r.Data.ID,
		}
		return s.Update(updateRequest)
	case "DELETE":
		deleteRequest := &DeleteRequest{
			ctx: r.Context,
			id:  r.Data.ID,
		}
		return s.Delete(deleteRequest)
	case "READ":
		readRequest := &ReadRequest{
			ctx: r.Context,
			id:  r.Data.ID,
		}
		return s.Read(readRequest)
	case "READALL":
		readAllRequest := &ReadAllRequest{
			ctx:     r.Context,
			filters: r.Data.Filters,
		}
		return s.ReadAll(readAllRequest)
	case "READCOUNT":
		readCountRequest := &ReadCountRequest{
			ctx:     r.Context,
			filters: r.Data.Filters,
		}
		return s.ReadCount(readCountRequest)
	case "ADDINTERFACE":
		addInterfaceRequest := &AddInterfaceRequest{
			ctx:      r.Context,
			resource: r.Data.Resource,
		}
		return s.AddInterface(addInterfaceRequest)
	case "DELINTERFACE":
		deleteInterfaceRequest := &DeleteInterfaceRequest{
			ctx:      r.Context,
			resource: r.Data.Resource,
		}
		return s.DeleteInterface(deleteInterfaceRequest)
	default:
		return nil, fmt.Errorf("method %s not supported", r.Context.Operation)
	}
}

//RegisterNeutronAPI register neutron endpoints
func (s *Service) RegisterNeutronAPI(r routeRegistry) {
	r.POST("/neutron/:id", s.NeutronPost)
}

func (s *Service) NeutronPost(c echo.Context) error {
	r, err := models.GetResource(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	request := &Request{
		Data: Data{
			Resource: r,
		},
	}
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %s", err))
	}
	response, err := s.Process(request)
	if err != nil {
		e, ok := errors.Cause(err).(*neutronerrors.NeutronError)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, response)
}

type routeRegistry interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}
