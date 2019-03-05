package neutron

import (
	"context"
	"fmt"
	"net/http"

	google_protobuf3 "github.com/gogo/protobuf/types"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/services"
)

// Server implementation.
type Server struct {
	ReadService       services.ReadService
	WriteService      services.WriteService
	UserAgentKV       userAgentKVServer
	IDToFQNameService idToFQNameServer
	FQNameToIDService fqNameToIDServer
	InTransactionDoer services.InTransactionDoer
}

// RegisterNeutronAPI registers Neutron endpoints on given routeRegistry.
func (s *Server) RegisterNeutronAPI(r routeRegistry) {
	r.POST("/neutron/:type", s.handleNeutronPostRequest)
}

func (s *Server) handleNeutronPostRequest(c echo.Context) error {
	var requestMap map[string]interface{}
	if err := c.Bind(&requestMap); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid JSON format: '%s'", err))
	}
	var request *logic.Request
	if err := format.ApplyMap(requestMap, &request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("failed to apply map: '%s'", err))
	}
	if t := c.Param("type"); request.GetType() != t {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid Resource type: '%s'", t))
	}
	request.Data.FieldMask = basemodels.MapToFieldMask(requestMap)
	var response logic.Response
	if err := s.InTransactionDoer.DoInTransaction(c.Request().Context(), func(ctx context.Context) error {
		var err error
		response, err = s.handle(ctx, request)
		return err
	}); err != nil {
		e, ok := errors.Cause(err).(*logic.Error)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) handle(ctx context.Context, r *logic.Request) (logic.Response, error) {
	rp := logic.RequestParameters{
		ReadService:       s.ReadService,
		WriteService:      s.WriteService,
		UserAgentKV:       s.UserAgentKV,
		IDToFQNameService: s.IDToFQNameService,
		FQNameToIDService: s.FQNameToIDService,
		RequestContext:    r.Context,
		FieldMask:         r.Data.FieldMask,
	}
	switch r.Context.Operation {
	case logic.OperationCreate:
		return r.Data.Resource.Create(ctx, rp)
	case logic.OperationUpdate:
		return r.Data.Resource.Update(ctx, rp, r.Data.ID)
	case logic.OperationDelete:
		return r.Data.Resource.Delete(ctx, rp, r.Data.ID)
	case logic.OperationRead:
		return r.Data.Resource.Read(ctx, rp, r.Data.ID)
	case logic.OperationReadAll:
		return r.Data.Resource.ReadAll(ctx, rp, r.Data.Filters, r.Data.Fields)
	case logic.OperationReadCount:
		return r.Data.Resource.ReadCount(ctx, rp, r.Data.Filters)
	case logic.OperationAddInterface:
		return r.Data.Resource.AddInterface(ctx, rp)
	case logic.OperationDelInterface:
		return r.Data.Resource.DeleteInterface(ctx, rp)
	default:
		err := errors.Errorf("method '%s' is not supported", r.Context.Operation)
		logrus.WithError(err).WithField("request", r).Errorf("failed to handle")
		return nil, err
	}
}

type userAgentKVServer interface {
	StoreKeyValue(context.Context, *services.StoreKeyValueRequest) (*google_protobuf3.Empty, error)
	RetrieveValues(context.Context, *services.RetrieveValuesRequest) (*services.RetrieveValuesResponse, error)
	RetrieveKVPs(context.Context, *google_protobuf3.Empty) (*services.RetrieveKVPsResponse, error)
	DeleteKey(context.Context, *services.DeleteKeyRequest) (*google_protobuf3.Empty, error)
}

type idToFQNameServer interface {
	IDToFQName(context.Context, *services.IDToFQNameRequest) (*services.IDToFQNameResponse, error)
}

type fqNameToIDServer interface {
	FQNameToID(context.Context, *services.FQNameToIDRequest) (*services.FQNameToIDResponse, error)
}

type routeRegistry interface {
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}
