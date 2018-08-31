package basemodels_test

import (
	"testing"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
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
			fm := basemodels.MapToFieldMask(tt.request)

			assert.Len(t, fm.Paths, len(tt.expected.Paths))
			for _, p := range fm.Paths {
				assert.Contains(t, tt.expected.Paths, p)
			}
		})
	}
}

func TestGetFromMapByPath(t *testing.T) {
	tests := []struct {
		name      string
		data      map[string]interface{}
		path      []string
		wantValue interface{}
		wantOk    bool
	}{
		{name: "nil"},
		{name: "nil map", path: []string{"asd", "bar"}},
		{name: "empty map", data: map[string]interface{}{}, path: []string{"asd"}},
		{name: "empty map nested path", data: map[string]interface{}{}, path: []string{"asd", "bar"}},
		{name: "first level exists", data: map[string]interface{}{"asd": 1}, path: []string{"asd", "bar"}},
		{name: "flat", data: map[string]interface{}{"asd": 1}, path: []string{"asd"}, wantValue: 1, wantOk: true},
		{
			name:      "nested",
			data:      map[string]interface{}{"asd": map[string]interface{}{"bar": "value"}},
			path:      []string{"asd", "bar"},
			wantValue: "value",
			wantOk:    true,
		},
		{name: "key exists value is nil", data: map[string]interface{}{"asd": nil}, path: []string{"asd"}, wantOk: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := basemodels.GetFromMapByPath(tt.data, tt.path)
			assert.Equal(t, tt.wantOk, gotOk)
			assert.Equal(t, tt.wantValue, gotValue)
		})
	}
}

func TestApplyFieldMask(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]interface{}
		fm   types.FieldMask
		want map[string]interface{}
	}{
		{name: "nil"},
		{name: "empty map empty fieldmask", m: map[string]interface{}{}, want: map[string]interface{}{}},
		{
			name: "empty map",
			m:    map[string]interface{}{},
			fm:   types.FieldMask{Paths: []string{"asd", "foo"}},
			want: map[string]interface{}{},
		},
		{name: "nonempty map empty fieldmask", m: map[string]interface{}{"key": "value"}, want: map[string]interface{}{}},
		{
			name: "fieldmask matching",
			m:    map[string]interface{}{"key": "value"},
			fm:   types.FieldMask{Paths: []string{"key"}},
			want: map[string]interface{}{"key": "value"},
		},
		{
			name: "mixed keys",
			m:    map[string]interface{}{"key": "value", "map": map[string]interface{}{"inner": 123}},
			fm:   types.FieldMask{Paths: []string{"key", "map.inner"}},
			want: map[string]interface{}{"key": "value", "map": map[string]interface{}{"inner": 123}},
		},
		{
			name: "three level nest",
			m:    map[string]interface{}{"map": map[string]interface{}{"inner": map[string]interface{}{"deep": true}}},
			fm:   types.FieldMask{Paths: []string{"key", "map.inner", "map.inner.deep"}},
			want: map[string]interface{}{"map": map[string]interface{}{"inner": map[string]interface{}{"deep": true}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := basemodels.ApplyFieldMask(tt.m, tt.fm)
			assert.Equal(t, tt.want, got)
		})
	}
}
