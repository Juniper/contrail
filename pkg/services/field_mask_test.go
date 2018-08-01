package services

import (
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

func TestMapToFieldMask(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		request map[string]interface{}
		want    types.FieldMask
	}{
		{name: "empty", want: types.FieldMask{Paths: []string{}}},
		{
			name: "only maps",
			request: map[string]interface{}{
				"simple": 1,
				"nested": map[string]interface{}{"inner": 1},
			},
			want: types.FieldMask{Paths: []string{"simple", "nested.inner"}},
		},
		{
			name: "map with KeyValuePairs",
			request: map[string]interface{}{
				"kvpairs": &models.KeyValuePairs{
					KeyValuePair: []*models.KeyValuePair{{Value: "val", Key: "key"}},
				},
			},
			want: types.FieldMask{Paths: []string{"kvpairs.key_value_pair"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapToFieldMask(tt.request)
			assert.Equal(t, tt.want, got)
		})
	}
}
