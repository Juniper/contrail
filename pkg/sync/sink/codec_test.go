package sink

import (
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestUpdateResourceData(t *testing.T) {
	codecs := []Codec{
		JSONCodec,
	}

	tests := []struct {
		name string
		data []byte
		obj  db.Object

		expected []byte
		fails    bool
	}{
		{name: "empty"},
		{name: "nil data", obj: models.MakeLogicalInterface(), fails: true},
		{name: "malformed data", data: []byte("asd"), obj: models.MakeLogicalInterface(), fails: true},
		{name: "empty data", data: []byte("{}"), obj: models.MakeLogicalInterface(),
			expected: []byte(`{"id_perms":{"permissions":{},"uuid":{}},"annotations":{},"perms2":{}}`)},
		{
			name: "new data is set, refs and fq_name are not overwritten",
			data: []byte(`{
				"uuid":"old uuid",
				"description": "description should be removed",
				"fq_name": ["should", "persist"],
				"virtual_machine_interface_refs": [{"uuid":"some uuid"}]
			}`),
			obj: &models.LogicalInterface{
				UUID: "new uuid",
				VirtualMachineInterfaceRefs: nil, // should be unchanged
			},
			expected: []byte(`{"uuid":"new uuid","fq_name":["should","persist"],"virtual_machine_interface_refs":[{"uuid":"some uuid"}]}`),
		},
		{
			name: "refs and fq_name update",
			data: []byte(`{
				"fq_name": ["should", "persist"],
				"virtual_machine_interface_refs": [{"uuid":"some uuid"}]
			}`),
			obj: &models.LogicalInterface{
				FQName: []string{"new", "data"},
				VirtualMachineInterfaceRefs: nil,
			},
			expected: []byte(`{"fq_name":["new","data"],"virtual_machine_interface_refs":[{"uuid":"some uuid"}]}`),
		},
	}

	for _, codec := range codecs {
		t.Run(fmt.Sprintf("%T", codec), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					result, err := UpdateResourceData(codec, tt.data, tt.obj)
					if tt.fails {
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
						fmt.Println(string(result))
						assert.Equal(t, tt.expected, result)
					}
				})
			}
		})
	}

}
