package baseservices

import (
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

func TestMapToFieldMask(t *testing.T) {
	tests := []struct {
		name     string
		request  map[string]interface{}
		expected types.FieldMask
	}{
		{
			name:     "returns nil paths given no data",
			expected: types.FieldMask{Paths: nil}},
		{
			name: "returns correct paths given data with maps",
			request: map[string]interface{}{
				"simple": 1,
				"nested": map[string]interface{}{"inner": 1},
			},
			expected: types.FieldMask{Paths: []string{"simple", "nested.inner"}},
		},
		{
			name: "returns correct paths given data with types implementing toMapper()",
			request: map[string]interface{}{
				"kvpairs": &models.KeyValuePairs{
					KeyValuePair: []*models.KeyValuePair{{Key: "key", Value: "val"}},
				},
				"qospairs": &models.QosIdForwardingClassPairs{
					QosIDForwardingClassPair: []*models.QosIdForwardingClassPair{{Key: 1, ForwardingClassID: 1234}},
				},
			},
			expected: types.FieldMask{Paths: []string{
				"kvpairs.key_value_pair",
				"qospairs.qos_id_forwarding_class_pair",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := MapToFieldMask(tt.request)

			assert.Len(t, fm.Paths, len(tt.expected.Paths))
			for _, p := range fm.Paths {
				assert.Contains(t, tt.expected.Paths, p)
			}
		})
	}
}

func TestFieldMaskContains(t *testing.T) {
	tests := []struct {
		name             string
		requestedFM      types.FieldMask
		requestedField   string
		expectedResponse bool
	}{
		{
			name:             "field mask contrains requested field",
			requestedFM:      types.FieldMask{Paths: []string{"first", "second"}},
			requestedField:   "first",
			expectedResponse: true,
		},
		{
			name:             "field mask doesn't contrain requested field",
			requestedFM:      types.FieldMask{Paths: []string{"first", "second"}},
			requestedField:   "third",
			expectedResponse: false,
		},
		{
			name:             "field mask is empty",
			requestedFM:      types.FieldMask{Paths: []string{}},
			requestedField:   "first",
			expectedResponse: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contains := FieldMaskContains(tt.requestedFM, tt.requestedField)
			assert.Equal(t, contains, tt.expectedResponse)
		})
	}
}
