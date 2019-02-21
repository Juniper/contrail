package collector

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
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
			name:      "MessageBusNotifyTrace: create MessageBusNotifyTrace message",
			operation: "create",
			typeName:  "project",
			uuid:      "created_project_uuid",
			fqName:    []string{"default-domain", "default-project"},
		},
		{
			name:      "MessageBusNotifyTrace: update MessageBusNotifyTrace message",
			operation: "update",
			typeName:  "project",
			uuid:      "updated_project_uuid",
			fqName:    []string{"default-domain", "default-project"},
		},
		{
			name:      "MessageBusNotifyTrace: delete MessageBusNotifyTrace message",
			operation: "delete",
			typeName:  "project",
			uuid:      "deleted_project_uuid",
			fqName:    []string{"default-domain", "default-project"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &mockBaseObject{
				typeName: tt.typeName,
				uuid:     tt.uuid,
				fqName:   tt.fqName,
			}

			messageBuilder := MessageBusNotifyTrace(tt.operation, obj)
			assert.NotNil(t, messageBuilder)
			message := messageBuilder.Build()
			assert.Nil(t, message)

			/* TODO: Should be reverted as introspect service for Intent API will be introduced.
			messageBuilder := MessageBusNotifyTrace(tt.operation, obj)
			assert.NotNil(t, messageBuilder)
			message := messageBuilder.Build()
			assert.NotNil(t, message)
			assert.Equal(t, message.SandeshType, typeMessageBusNotifyTrace)
			m, ok := message.Payload.(*payloadMessageBusNotifyTrace)
			assert.True(t, ok)
			assert.Equal(t, m.Operation, tt.operation)
			assert.Equal(t, m.Body.Operation, tt.operation)
			assert.Equal(t, m.Body.Type, tt.typeName)
			assert.Equal(t, m.Body.UUID, tt.uuid)
			assert.Equal(t, m.Body.FQName, tt.fqName)
			assert.True(t, strings.HasPrefix(m.RequestID, "req-"))
			assert.Equal(t, m.RequestID, m.Body.RequestID)
			*/
		})
	}
}

type configTraceTestInfo struct {
	name       string
	kind       string
	operation  string
	table      string
	uuid       string
	fqName     []string
	resultName string
	deleted    bool
}

func TestContrailConfigTrace(t *testing.T) {
	var tests []configTraceTestInfo
	for kind, info := range uveTables {
		if info.IsGlobal == true {
			tests = append(tests, configTraceTestInfo{
				name:       "ContrailConfigTrace: check global type " + kind,
				kind:       kind,
				operation:  services.OperationCreate,
				table:      info.TableName,
				uuid:       kind + "_uuid",
				fqName:     []string{"default-domain", "default-project", kind + "_name"},
				resultName: kind + "_name",
				deleted:    false,
			})
		} else {
			tests = append(tests, configTraceTestInfo{
				name:       "ContrailConfigTrace: check type " + kind,
				kind:       kind,
				operation:  services.OperationDelete,
				table:      info.TableName,
				uuid:       kind + "_uuid",
				fqName:     []string{"default-domain", "default-project", kind + "_name"},
				resultName: "default-domain:default-project:" + kind + "_name",
				deleted:    true,
			})
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &mockBaseObject{
				typeName: tt.kind,
				uuid:     tt.uuid,
				fqName:   tt.fqName,
			}

			messageBuilder := ContrailConfigTrace(tt.kind, tt.operation, obj)
			assert.NotNil(t, messageBuilder)
			message := messageBuilder.Build()
			assert.NotNil(t, message)

			assert.Equal(t, message.SandeshType, typeContrailConfigTrace)
			m, ok := message.Payload.(*payloadContrailConfigTrace)
			assert.True(t, ok)
			assert.Equal(t, m.Table, tt.table)
			assert.Equal(t, m.Name, tt.resultName)
			assert.Equal(t, m.Deleted, tt.deleted)
		})
	}

	messageBuilder := ContrailConfigTrace("global_config", services.OperationUpdate, nil)
	assert.NotNil(t, messageBuilder)
	message := messageBuilder.Build()
	assert.Nil(t, message)
}
