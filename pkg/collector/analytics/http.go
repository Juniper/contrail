package analytics

import (
	"net/http"
	"time"

	"github.com/Juniper/contrail/pkg/apisrv/baseapisrv"
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
func (p BodyDumpPlugin) RegisterHTTPAPI(r baseapisrv.HTTPRouter) error {
	r.Use(fromEchoMiddlewareFunc(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		p.Send(RESTAPITrace(ctx, reqBody, resBody))
	})))
	return nil
}

// fromEchoMiddlewareFunc makes a baseapisrv.MiddlewareFunc from echo.MiddlewareFunc.
func fromEchoMiddlewareFunc(m echo.MiddlewareFunc) baseapisrv.MiddlewareFunc {
	return func(next baseapisrv.HandlerFunc) baseapisrv.HandlerFunc {
		return baseapisrv.HandlerFunc(m(echo.HandlerFunc(next)))
	}
}

// RegisterGRPCAPI does nothing.
func (BodyDumpPlugin) RegisterGRPCAPI(r baseapisrv.GRPCRouter) error {
	return nil
}
