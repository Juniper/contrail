package collector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageBusNotifyTrace(t *testing.T) {
	tests := []struct {
		name      string
		operation string
		typeName  string
		uuid      string
		fqName    []string
	}{
		{
			name:      "create MessageBusNotifyTrace message",
			operation: "CREATE",
			typeName:  "project",
			uuid:      "created_project_uuid",
			fqName:    []string{"default-domain", "default-project"},
		},
		{
			name:      "update MessageBusNotifyTrace message",
			operation: "UPDATE",
			typeName:  "project",
			uuid:      "updated_project_uuid",
			fqName:    []string{"default-domain", "default-project"},
		},
		{
			name:      "delete MessageBusNotifyTrace message",
			operation: "DELETE",
			typeName:  "project",
			uuid:      "deleted_project_uuid",
			fqName:    []string{"default-domain", "default-project"},
		},
	}

	c, err := NewCollector(&Config{})
	assert.NoError(t, err)
	s := &mockSender{}
	c.sender = s

	for _, tt := range tests {
		t.Run("RESTAPITrace", func(t *testing.T) {
			c.MessageBusNotifyTrace(tt.operation, tt.typeName, tt.uuid, tt.fqName)
			assert.Equal(t, s.message.SandeshType, typeMessageBusNotifyTrace)
			m, ok := s.message.Payload.(*payloadMessageBusNotifyTrace)
			assert.True(t, ok)
			assert.Equal(t, m.Operation, tt.operation)
			assert.Equal(t, m.Body.Operation, tt.operation)
			assert.Equal(t, m.Body.Type, tt.typeName)
			assert.Equal(t, m.Body.UUID, tt.uuid)
			assert.Equal(t, m.Body.FQName, tt.fqName)
		})
	}
}
