package analytics

import (
	"net/http"
	"time"

	"github.com/Juniper/contrail/pkg/collector"
)

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type LatencyLoggingDoer struct {
	Doer        doer
	Collector   collector.Collector
	Operation   string
	Application string
}

func (d LatencyLoggingDoer) Do(req *http.Request) (*http.Response, error) {
	startedAt := time.Now()
	resp, err := d.Doer.Do(req)
	elapsed := time.Since(startedAt)

	ctx := req.Context()
	d.Collector.Send(VncAPILatencyStatsLog(ctx, d.Operation, d.Application, int64(elapsed/time.Microsecond)))

	return resp, err
}
