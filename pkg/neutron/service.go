package neutron

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/services"
)

// Service implementation.
type Service struct {
	ReadService  services.ReadService
	WriteService services.WriteService
}

// Request incoming data format
type Request struct {
	Data    Data                 `json:"data" yaml:"data"`
	Context logic.RequestContext `json:"context" yaml:"context"`
}

// Data structure
type Data struct {
	Filters  logic.Filters  `json:"filters" yaml:"filters"`
	ID       string         `json:"id" yaml:"id"`
	Fields   logic.Fields   `json:"fields" yaml:"fields"`
	Resource logic.Resource `json:"resource" yaml:"resource"`
}

// GetType returns resource type of a request
func (r *Request) GetType() string {
	return r.Context.Type
}

func (s *Service) process(r *Request) (logic.Response, error) {
	ctx := logic.Context{
		ReadService:    s.ReadService,
		WriteService:   s.WriteService,
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
		err := fmt.Errorf("method %s not supported", r.Context.Operation)
		log.WithError(err).Errorf("failed to process request %v", r)
		return nil, err
	}
}

//RegisterNeutronAPI register neutron endpoints
func (s *Service) RegisterNeutronAPI(r routeRegistry) {
	r.POST("/neutron/:id", s.neutronPost)
}

func (s *Service) neutronPost(c echo.Context) error {
	r, err := logic.GetResource(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	request := &Request{
		Data: Data{
			Resource: r,
		},
	}
	if err = c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: %s", err))
	}
	response, err := s.process(request)
	if err != nil {
		e, ok := errors.Cause(err).(*logic.NeutronError)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, e)
		}
		raw, err := json.Marshal(e)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError,
				errors.Wrapf(e, "failed to marshall error description."))
		}
		return echo.NewHTTPError(http.StatusBadRequest, raw)
	}
	return c.JSON(http.StatusOK, response)
}

type routeRegistry interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}
