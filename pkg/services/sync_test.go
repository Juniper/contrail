package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderResourceList(t *testing.T) {
	resources := &ResourceList{
		Resources: []*ResourceEvent{
			&ResourceEvent{
				Kind: "b",
				Data: map[string]interface{}{
					"uuid": "b",
					"a_refs": []interface{}{
						map[string]interface{}{
							"uuid": "a",
						},
					},
				},
			},
			&ResourceEvent{
				Kind: "c",
				Data: map[string]interface{}{
					"uuid":        "c",
					"parent_uuid": "a",
				},
			},
			&ResourceEvent{
				Kind: "a",
				Data: map[string]interface{}{
					"uuid": "a",
				},
			},
		},
	}

	err := resources.Sort()
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, resources.Resources[0].Kind, "a")
	assert.Equal(t, resources.Resources[1].Kind, "b")
	assert.Equal(t, resources.Resources[2].Kind, "c")
}

func TestReorderLoopedResourceList(t *testing.T) {
	resources := &ResourceList{
		Resources: []*ResourceEvent{
			&ResourceEvent{
				Kind: "b",
				Data: map[string]interface{}{
					"uuid": "b",
					"a_refs": []interface{}{
						map[string]interface{}{
							"uuid": "a",
						},
					},
				},
			},
			&ResourceEvent{
				Kind: "c",
				Data: map[string]interface{}{
					"uuid":        "c",
					"parent_uuid": "a",
				},
			},
			&ResourceEvent{
				Kind: "a",
				Data: map[string]interface{}{
					"uuid": "a",
					"b_refs": map[string]interface{}{
						"uuid": "b",
					},
				},
			},
		},
	}

	err := resources.Sort()
	assert.Error(t, err, "loop error expected")
}
