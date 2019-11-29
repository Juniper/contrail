package analytics

import (
	"net/http"
	"time"

	"github.com/Juniper/contrail/pkg/collector"
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
