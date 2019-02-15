package collector

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

type mockBaseObject struct {
	models.Project
	typeName string
	uuid     string
	fqName   []string
}

func (o *mockBaseObject) Kind() string        { return o.typeName }
func (o *mockBaseObject) GetUUID() string     { return o.uuid }
func (o *mockBaseObject) GetFQName() []string { return o.fqName }

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
			operation: "create",
			typeName:  "project",
			uuid:      "created_project_uuid",
			fqName:    []string{"default-domain", "default-project"},
		},
		{
			name:      "update MessageBusNotifyTrace message",
			operation: "update",
			typeName:  "project",
			uuid:      "updated_project_uuid",
			fqName:    []string{"default-domain", "default-project"},
		},
		{
			name:      "delete MessageBusNotifyTrace message",
			operation: "delete",
			typeName:  "project",
			uuid:      "deleted_project_uuid",
			fqName:    []string{"default-domain", "default-project"},
		},
	}

	for _, tt := range tests {
		t.Run("RESTAPITrace", func(t *testing.T) {
			obj := &mockBaseObject{
				typeName: tt.typeName,
				uuid:     tt.uuid,
				fqName:   tt.fqName,
			}

			message := MessageBusNotifyTrace(tt.operation, obj).Build()

			assert.Equal(t, message.SandeshType, typeMessageBusNotifyTrace)
			m, ok := message.Payload.(*payloadMessageBusNotifyTrace)
			assert.True(t, ok)
			assert.Equal(t, m.Operation, tt.operation)
			assert.Equal(t, m.Body.Operation, tt.operation)
			assert.Equal(t, m.Body.Type, tt.typeName)
			assert.Equal(t, m.Body.UUID, tt.uuid)
			assert.Equal(t, m.Body.FQName, tt.fqName)
		})
	}
}
