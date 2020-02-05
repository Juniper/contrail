package cassandra

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
