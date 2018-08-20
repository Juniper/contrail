package cassandra

import (
	"github.com/Juniper/contrail/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVirtualNetwork(t *testing.T) {

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

	m, err := VirtualMachineInterfaceToCassandraMap(vmi)

	assert.NoError(t, err)

	for k, v := range expected {
		s, found := m[k]
		assert.True(t, found)
		assert.Equal(t, v, s)
	}
}
