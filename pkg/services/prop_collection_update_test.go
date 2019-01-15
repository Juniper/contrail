package services

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models"
)

func Test_newCollectionItem(t *testing.T) {
	tests := []struct {
		name    string
		obj     interface{}
		field   string
		want    interface{}
		wantErr bool
	}{
		{
			name:  "resolve KeyValuePair type",
			obj:   models.VirtualNetwork{},
			field: "annotations",
			want:  (*models.KeyValuePair)(nil),
		},
		{
			name:  "pointer type",
			obj:   &models.VirtualNetwork{},
			field: "annotations",
			want:  (*models.KeyValuePair)(nil),
		},
		{
			name:    "non collection field",
			obj:     models.VirtualNetwork{},
			field:   "uuid",
			wantErr: true,
		},
		{
			name:    "invalid field",
			obj:     models.VirtualNetwork{},
			field:   "bad_field_that_does_not_exist",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newCollectionItem(tt.obj, tt.field)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_fieldByTag(t *testing.T) {
	tests := []struct {
		name   string
		t      reflect.Type
		key    string
		value  string
		want   reflect.StructField
		wantOk bool
	}{
		{
			name: "virtual network's annotations field",
			t:    reflect.TypeOf(models.VirtualNetwork{}),
			key:  "json", value: "annotations",
			want:   fieldByName(models.VirtualNetwork{}, "Annotations"),
			wantOk: true,
		},
		{
			name: "annotations field from pointer type",
			t:    reflect.TypeOf(&models.VirtualNetwork{}),
			key:  "json", value: "annotations",
			want:   fieldByName(models.VirtualNetwork{}, "Annotations"),
			wantOk: true,
		},
		{
			name: "field name typo",
			t:    reflect.TypeOf(models.VirtualNetwork{}),
			key:  "json", value: "annotatio",
			wantOk: false,
		},
		{
			name: "non existing tag",
			t:    reflect.TypeOf(models.VirtualNetwork{}),
			key:  "bad", value: "annotations",
			wantOk: false,
		},
		{
			name: "nil type",
			t:    reflect.TypeOf(nil),
			key:  "bad", value: "annotations",
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := fieldByTag(tt.t, tt.key, tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantOk, got1)
		})
	}
}

func fieldByName(x interface{}, name string) reflect.StructField {
	f, _ := reflect.TypeOf(x).FieldByName(name)
	return f
}

func TestPropCollectionUpdateRequest_resolveCollectionItems(t *testing.T) {
	tests := []struct {
		name     string
		updates  []*PropCollectionChange
		obj      interface{}
		expected PropCollectionUpdateRequest
		fails    bool
	}{
		{name: "empty"},
		{
			name:    "resolve annotations update",
			updates: []*PropCollectionChange{{Field: "annotations"}},
			obj:     models.VirtualNetwork{},
			expected: PropCollectionUpdateRequest{
				Updates: []*PropCollectionChange{{
					Field: "annotations",
					Value: &PropCollectionChange_KeyValuePairValue{},
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PropCollectionUpdateRequest{
				Updates: tt.updates,
			}
			err := p.resolveCollectionItems(tt.obj)
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, *p)
			}
		})
	}
}

func TestResolveCollectionItemsAndUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		request  PropCollectionUpdateRequest
		obj      interface{}
		expected PropCollectionUpdateRequest
	}{
		{name: "empty", data: []byte(`{}`)},
		{
			name: "annotations",
			data: []byte(`{"updates": [{"field": "annotations", "operation": "set", "value": {"key": "some-key", "value":"val"}}]}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.resolveCollectionItems(tt.obj)
			assert.NoError(t, err)

			err = json.Unmarshal(tt.data, &tt.request)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, tt.request)
		})
	}
}
