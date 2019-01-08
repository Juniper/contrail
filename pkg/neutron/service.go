package neutron

import (
	"context"
	"fmt"
	"net/http"

	google_protobuf3 "github.com/gogo/protobuf/types"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/services"
)

// Service implementation.
type Service struct {
	ReadService     services.ReadService
	WriteService    services.WriteService
	UserAgentKV     userAgentKVServer
	FQNameService   services.FQNameToIDServer
	IDToTypeService idToTypeServer
}

// RegisterNeutronAPI registers Neutron endpoints on given routeRegistry.
func (s *Service) RegisterNeutronAPI(r routeRegistry) {
	r.POST("/neutron/:type", s.handleNeutronPostRequest)
}

func (s *Service) handleNeutronPostRequest(c echo.Context) error {
	request := &logic.Request{}
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: '%s'", err))
	}
	if t := c.Param("type"); request.GetType() != t {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid Resource type: '%s'", t))
	}
	response, err := s.handle(c.Request().Context(), request)
	if err != nil {
		e, ok := errors.Cause(err).(*logic.Error)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return echo.NewHTTPError(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Service) handle(ctx context.Context, r *logic.Request) (logic.Response, error) {
	rp := logic.RequestParameters{
		ReadService:     s.ReadService,
		WriteService:    s.WriteService,
		UserAgentKV:     s.UserAgentKV,
		IDToTypeService: s.IDToTypeService,
		RequestContext:  r.Context,
		FQNameService:   s.FQNameService,
	}
	switch r.Context.Operation {
	case "CREATE":
		return r.Data.Resource.Create(ctx, rp)
	case "UPDATE":
		return r.Data.Resource.Update(ctx, rp, r.Data.ID)
	case "DELETE":
		return r.Data.Resource.Delete(ctx, rp, r.Data.ID)
	case "READ":
		return r.Data.Resource.Read(ctx, rp, r.Data.ID)
	case "READALL":
		return r.Data.Resource.ReadAll(ctx, rp, r.Data.Filters, r.Data.Fields)
	case "READCOUNT":
		return r.Data.Resource.ReadCount(ctx, rp, r.Data.Filters)
	case "ADDINTERFACE":
		return r.Data.Resource.AddInterface(ctx, rp)
	case "DELINTERFACE":
		return r.Data.Resource.DeleteInterface(ctx, rp)
	default:
		err := errors.Errorf("method '%s' is not supported", r.Context.Operation)
		log.WithError(err).WithField("request", r).Errorf("failed to handle")
		return nil, err
	}
}

type userAgentKVServer interface {
	StoreKeyValue(context.Context, *services.StoreKeyValueRequest) (*google_protobuf3.Empty, error)
	RetrieveValues(context.Context, *services.RetrieveValuesRequest) (*services.RetrieveValuesResponse, error)
	RetrieveKVPs(context.Context, *google_protobuf3.Empty) (*services.RetrieveKVPsResponse, error)
	DeleteKey(context.Context, *services.DeleteKeyRequest) (*google_protobuf3.Empty, error)
}

type idToTypeServer interface {
	IDToType(context.Context, *services.IDToTypeRequest) (*services.IDToTypeResponse, error)
}

type routeRegistry interface {
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}
