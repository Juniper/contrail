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
			name:         "VncAPILatencyStatsLog: SQL commit tracing",
			requestID:    "req-1",
			operation:    "COMMIT",
			application:  "SQL",
			responseTime: int64(37892743),
		},
		{
			name:         "VncAPILatencyStatsLog: Keystone request tracing",
			requestID:    "req-2",
			operation:    "VALIDATE",
			application:  "KEYSTONE",
			responseTime: int64(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageBuilder := VncAPILatencyStatsLog(
				services.WithRequestID(context.Background(), tt.requestID),
				tt.operation, tt.application, tt.responseTime)
			assert.NotNil(t, messageBuilder)
			message := messageBuilder.Build()
			assert.NotNil(t, message)

			assert.Equal(t, message.SandeshType, typeVncAPILatencyStatsLog)
			m, ok := message.Payload.(*payloadVncAPILatencyStatsLog)
			assert.True(t, ok)
			assert.Equal(t, m.NodeName, vncStatNodeName)
			assert.Equal(t, m.Stats.Operation, tt.operation)
			assert.Equal(t, m.Stats.Application, tt.application)
			assert.Equal(t, m.Stats.ResponseTime, tt.responseTime)
			assert.Equal(t, m.Stats.ResponseSize, int64(0))
			assert.Equal(t, m.Stats.RequestID, tt.requestID)
		})
	}
}
