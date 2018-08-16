package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceToVirtualMachineInterface(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  *VirtualMachineInterface
	}{
		{name: "nil"},
		{name: "empty map", input: map[string]interface{}{}, want: &VirtualMachineInterface{}},
		{
			name: "simple props",
			input: map[string]interface{}{
				"uuid": "some-uuid",
				"name": "some-name",
			},
			want: &VirtualMachineInterface{UUID: "some-uuid", Name: "some-name"},
		},
		{
			name: "annotations with subfield",
			input: map[string]interface{}{
				"annotations": map[string]interface{}{"key_value_pair": []interface{}{
					map[string]interface{}{"key": "k", "value": "v"},
				}},
			},
			want: &VirtualMachineInterface{
				Annotations: &KeyValuePairs{KeyValuePair: []*KeyValuePair{{Value: "v", Key: "k"}}},
			},
		},
		{
			name: "annotations provided as kv list",
			input: map[string]interface{}{
				"annotations": []interface{}{
					map[string]interface{}{"key": "k", "value": "v"},
				},
			},
			want: &VirtualMachineInterface{
				Annotations: &KeyValuePairs{KeyValuePair: []*KeyValuePair{{Value: "v", Key: "k"}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InterfaceToVirtualMachineInterface(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
