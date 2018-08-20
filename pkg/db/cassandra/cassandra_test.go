package cassandra

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

func TestParseProperty(t *testing.T) {
	tests := []struct {
		name     string
		data     map[string]interface{}
		property string
		value    interface{}
		expected map[string]interface{}
	}{
		{
			name:     "empty",
			expected: map[string]interface{}{"": (interface{})(nil)},
		},
		{
			name:     "unknown type",
			property: "key",
			value:    "value",
			expected: map[string]interface{}{"key": "value"},
		},
		{
			name:     "simple prop",
			property: "prop:foo",
			value:    1,
			expected: map[string]interface{}{"foo": 1},
		},
		{
			name:     "simple propl",
			property: "propl:foo",
			value:    1,
			expected: map[string]interface{}{"foo": []interface{}{1}},
		},
		{
			name:     "simple propm",
			property: "propm:foo",
			value:    1,
			expected: map[string]interface{}{"foo": []interface{}{1}},
		},
		{
			name:     "propl appends",
			data:     map[string]interface{}{"foo": []interface{}{1}},
			property: "propl:foo",
			value:    2,
			expected: map[string]interface{}{"foo": []interface{}{1, 2}},
		},
		{
			name:     "propm appends",
			data:     map[string]interface{}{"foo": []interface{}{1}},
			property: "propm:foo",
			value:    2,
			expected: map[string]interface{}{"foo": []interface{}{1, 2}},
		},
		{
			name:     "parent",
			property: "parent:foo:some-parent",
			expected: map[string]interface{}{"parent_uuid": "some-parent"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.data == nil {
				tt.data = map[string]interface{}{}
			}
			parseProperty(tt.data, tt.property, tt.value)
			assert.Equal(t, tt.expected, tt.data)
		})
	}
}

func TestVirtualNetworkToCassandra(t *testing.T) {
	vmi := &models.VirtualMachineInterface{
		UUID:       "hoge",
		ParentUUID: "hoge2",
		ParentType: "beef",
		Name:       "blue",
		RoutingInstanceRefs: []*models.VirtualMachineInterfaceRoutingInstanceRef{
			{
				UUID: "hoge3",
			},
		},
		VirtualMachineInterfaceBindings: &models.KeyValuePairs{
			KeyValuePair: []*models.KeyValuePair{
				{
					Key:   "a",
					Value: "b",
				},
				{
					Key:   "c",
					Value: "d",
				},
			},
		},
		VirtualMachineInterfaceFatFlowProtocols: &models.FatFlowProtocols{
			FatFlowProtocol: []*models.ProtocolType{
				{
					Port:     1,
					Protocol: "hoge",
				},
				{
					Port:     2,
					Protocol: "hoge2",
				},
			},
		},
	}

	expected := map[string]interface{}{
		"propl:virtual_machine_interface_fat_flow_protocols:0": `{"port":1,"protocol":"hoge"}`,
		"ref:routing_instance:hoge3":                           `{"attr":null}`,
		"parent:beef:hoge2":                                    ``,
		"parent_type":                                          `"beef"`,
		"propl:virtual_machine_interface_fat_flow_protocols:1": `{"port":2,"protocol":"hoge2"}`,
		"propm:virtual_machine_interface_bindings:c":           `{"value":"d","key":"c"}`,
		"propm:virtual_machine_interface_bindings:a":           `{"value":"b","key":"a"}`,
	}

	if m, err := VirtualMachineInterfaceToCassandraMap(vmi); assert.NoError(t, err) {
		for k, v := range expected {
			if s, found := m[k]; assert.True(t, found) {
				assert.Equal(t, v, s)
			}
		}
	}
}
