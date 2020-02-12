package analytics

import (
	"net/http"
	"time"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// LatencyReportingDoer wraps doer and sends request execution times to collector.
type LatencyReportingDoer struct {
	Doer        doer
	Collector   collector.Collector
	Operation   string
	Application string
}

// Do executes the request and sends duration to collector.
func (d LatencyReportingDoer) Do(req *http.Request) (*http.Response, error) {
	startedAt := time.Now()
	resp, err := d.Doer.Do(req)
	elapsed := time.Since(startedAt)

	d.Collector.Send(
		VncAPILatencyStatsLog(req.Context(), d.Operation, d.Application, int64(elapsed/time.Microsecond)),
	)

	return resp, err
}

// BodyDumpPlugin sends HTTP request and response body to Collector.
type BodyDumpPlugin struct {
	collector.Collector
}

// RegisterHTTPAPI registers middleware for all endpoints.
func (p BodyDumpPlugin) RegisterHTTPAPI(r apiserver.HTTPRouter) {
	r.Use(middlewareFunc(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		p.Send(RESTAPITrace(ctx, reqBody, resBody))
	})))
}

// middlewareFunc makes an apiserver.MiddlewareFunc from echo.MiddlewareFunc.
func middlewareFunc(m echo.MiddlewareFunc) apiserver.MiddlewareFunc {
	return func(next apiserver.HandlerFunc) apiserver.HandlerFunc {
		return apiserver.HandlerFunc(m(echo.HandlerFunc(next)))
	}
}

// RegisterGRPCAPI does nothing.
func (BodyDumpPlugin) RegisterGRPCAPI(r apiserver.GRPCRouter) {
}
