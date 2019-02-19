package analytics

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/services"
)

func TestVncAPILatencyStatsLog(t *testing.T) {
	tests := []struct {
		name         string
		requestID    string
		operation    string
		application  string
		responseTime int64
	}{
		{
			name:         "SQL commit tracing",
			requestID:    "req-1",
			operation:    "COMMIT",
			application:  "SQL",
			responseTime: int64(37892743),
		},
		{
			name:         "Keystone request tracing",
			requestID:    "req-1",
			operation:    "VALIDATE",
			application:  "KEYSTONE",
			responseTime: int64(0),
		},
	}

	for _, tt := range tests {
		t.Run("VncAPILatencyStatsLog", func(t *testing.T) {
			message := VncAPILatencyStatsLog(
				services.WithRequestID(context.Background(), tt.requestID),
				tt.operation, tt.application, tt.responseTime).Build()

			assert.Equal(t, message.SandeshType, typeVncAPILatencyStatsLog)
			m, ok := message.Payload.(*payloadVncAPILatencyStatsLog)
			assert.True(t, ok)
			assert.Equal(t, m.NodeName, "issu-vm6")
			assert.Equal(t, m.Stats.Operation, tt.operation)
			assert.Equal(t, m.Stats.Application, tt.application)
			assert.Equal(t, m.Stats.ResponseTime, tt.responseTime)
			assert.Equal(t, m.Stats.ResponseSize, int64(0))
			assert.Equal(t, m.Stats.RequestID, tt.requestID)
		})
	}
}
