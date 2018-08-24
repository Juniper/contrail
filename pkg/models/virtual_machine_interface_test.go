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
			name: "annotations as object",
			input: map[string]interface{}{
				"annotations": &KeyValuePairs{KeyValuePair: []*KeyValuePair{{Value: "v", Key: "k"}}},
			},
			want: &VirtualMachineInterface{
				Annotations: &KeyValuePairs{KeyValuePair: []*KeyValuePair{{Value: "v", Key: "k"}}},
			},
		},
		{
			name: "fat flow protocols provided as list of objects",
			input: map[string]interface{}{
				"virtual_machine_interface_fat_flow_protocols": []interface{}{
					map[string]interface{}{"ignore_address": "none", "port": 0, "protocol": "proto"},
				},
			},
			want: &VirtualMachineInterface{
				VirtualMachineInterfaceFatFlowProtocols: &FatFlowProtocols{FatFlowProtocol: []*ProtocolType{
					{IgnoreAddress: "none", Port: 0, Protocol: "proto"},
				}},
			},
		},
		{
			name: "fat flow protocols with subfield",
			input: map[string]interface{}{
				"virtual_machine_interface_fat_flow_protocols": map[string]interface{}{
					"fat_flow_protocol": []interface{}{
						map[string]interface{}{"ignore_address": "none", "port": 0, "protocol": "proto"},
					},
				},
			},
			want: &VirtualMachineInterface{
				VirtualMachineInterfaceFatFlowProtocols: &FatFlowProtocols{FatFlowProtocol: []*ProtocolType{
					{IgnoreAddress: "none", Port: 0, Protocol: "proto"},
				}},
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

func TestVirtualMachineInterfaceApplyMap(t *testing.T) {
	tests := []struct {
		name  string
		obj   *VirtualMachineInterface
		input map[string]interface{}
		want  *VirtualMachineInterface
	}{
		{name: "nil"},
		{name: "nil obj", input: map[string]interface{}{"uuid": "value"}},
		{name: "nil map", obj: &VirtualMachineInterface{}, want: &VirtualMachineInterface{}},
		{name: "empty map", obj: &VirtualMachineInterface{}, input: map[string]interface{}{}, want: &VirtualMachineInterface{}},
		{
			name: "simple props",
			obj:  &VirtualMachineInterface{UUID: "old-uuid", Name: "some-name"},
			input: map[string]interface{}{
				"uuid": "some-uuid",
			},
			want: &VirtualMachineInterface{UUID: "some-uuid", Name: "some-name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.obj.ApplyMap(tt.input)
			assert.Equal(t, tt.want, tt.obj)
		})
	}
}
