package neutron

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/services"
)

// Service implementation.
type Service struct {
	readService  services.ReadService
	writeService services.WriteService
}

type Request struct {
	Data    Data                 `json:"data" yaml:"data"`
	Context logic.RequestContext `json:"context" yaml:"context"`
}

type Data struct {
	Filters  logic.Filters  `json:"filters" yaml:"filters"`
	ID       string         `json:"id" yaml:"id"`
	Fields   logic.Fields   `json:"fields" yaml:"fields"`
	Resource logic.Resource `json:"resource" yaml:"resource"`
}

func NewRequest() *Request {
	return &Request{
		Data: Data{
			Filters: logic.Filters{},
			Fields:  logic.Fields{},
		},
		Context: logic.RequestContext{},
	}
}

func (r *Request) GetType() string {
	return r.Context.Type
}

func (s *Service) Process(r *Request) (logic.Response, error) {
	ctx := logic.Context{
		ReadService:    s.readService,
		WriteService:   s.writeService,
		RequestContext: r.Context,
	}
	switch r.Context.Operation {
	case "CREATE":
		return r.Data.Resource.Create(ctx)
	case "UPDATE":
		return r.Data.Resource.Update(ctx)
	case "DELETE":
		return r.Data.Resource.Delete(ctx, r.Data.ID)
	case "READ":
		return r.Data.Resource.Read(ctx, r.Data.ID)
	case "READALL":
		return r.Data.Resource.ReadAll(ctx, r.Data.Filters, r.Data.Fields)
	case "READCOUNT":
		return r.Data.Resource.ReadCount(ctx, r.Data.Filters)
	case "ADDINTERFACE":
		return r.Data.Resource.AddInterface(ctx)
	case "DELINTERFACE":
		return r.Data.Resource.DeleteInterface(ctx)
	default:
		return nil, fmt.Errorf("method %s not supported", r.Context.Operation)
	}
}

//RegisterNeutronAPI register neutron endpoints
func (s *Service) RegisterNeutronAPI(r routeRegistry) {
	r.POST("/neutron/:id", s.NeutronPost)
}

func (s *Service) NeutronPost(c echo.Context) error {
	r, err := logic.GetResource(c.Param("id"))
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
		e, ok := errors.Cause(err).(*logic.NeutronError)
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
