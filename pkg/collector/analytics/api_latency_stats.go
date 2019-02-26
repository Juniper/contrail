package analytics

import (
	"context"

	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	typeVncAPILatencyStatsLog = "VncApiLatencyStatsLog"

	vncStatNodeName = "issu-vm6"
)

type payloadVncAPILatencyStatsLog struct {
	NodeName string             `json:"node_name"`
	Stats    *vncAPILatencyStat `json:"api_latency_stats"`
}

type vncAPILatencyStat struct {
	Operation    string `json:"operation_type"`
	Application  string `json:"application"`
	ResponseTime int64  `json:"response_time_in_usec"`
	ResponseSize int64  `json:"response_size"`
	RequestID    string `json:"identifier"`
}

func (p *payloadVncAPILatencyStatsLog) Build() *collector.Message {
	return &collector.Message{
		SandeshType: typeVncAPILatencyStatsLog,
		Payload:     p,
	}
}

// VncAPILatencyStatsLog sends message with type VncAPILatencyStatsLog
func VncAPILatencyStatsLog(
	ctx context.Context, operation, application string, responseTime int64,
) collector.MessageBuilder {
	requestID := services.GetRequestID(ctx)

	return &payloadVncAPILatencyStatsLog{
		NodeName: vncStatNodeName,
		Stats: &vncAPILatencyStat{
			Operation:    operation,
			Application:  application,
			ResponseTime: responseTime,
			RequestID:    requestID,
		},
	}
}
