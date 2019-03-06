package basemodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChildFQName(t *testing.T) {
	tests := []struct {
		name         string
		parentFQName []string
		childName    string
		want         []string
	}{
		{name: "empty", want: []string{}},
		{name: "empty parentFQName", childName: "my-name", want: []string{"my-name"}},
		{
			name:         "empty childName",
			parentFQName: []string{"grandparent", "parent"},
			want:         []string{"grandparent", "parent"}},
		{
			name:         "both not empty",
			parentFQName: []string{"grandparent", "parent"},
			childName:    "name",
			want:         []string{"grandparent", "parent", "name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChildFQName(tt.parentFQName, tt.childName)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFQNameEquals(t *testing.T) {
	tests := []struct {
		name     string
		fqNameA  []string
		fqNameB  []string
		areEqual bool
	}{
		{
			name:     "Check if two FQNames (slices) with different length are equal",
			fqNameA:  []string{"a", "b", "c"},
			fqNameB:  []string{"a", "b", "c", "d"},
			areEqual: false,
		},
		{
			name:     "Check if two FQNames (slices) with the same length but diff values are equal",
			fqNameA:  []string{"a", "b", "c"},
			fqNameB:  []string{"a", "b", "d"},
			areEqual: false,
		},
		{
			name:     "Check if two FQNames (slices) with the same length and values but in diff order are equal",
			fqNameA:  []string{"a", "b", "c"},
			fqNameB:  []string{"c", "b", "a"},
			areEqual: false,
		},
		{
			name:     "Check if two FQNames (slices) with the same length, values and order are equal",
			fqNameA:  []string{"a", "b", "c"},
			fqNameB:  []string{"a", "b", "c"},
			areEqual: true,
		},
		{
			name:     "Check if two empty FQNames (slices) are equal",
			fqNameA:  []string{},
			fqNameB:  []string{},
			areEqual: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FQNameEquals(tt.fqNameA, tt.fqNameB)
			assert.Equal(t, tt.areEqual, got)
		})
	}
}

func TestKindToSchemaID(t *testing.T) {
	tests := []struct {
		kind     string
		schemaID string
	}{
		{schemaID: "api_list", kind: "api-list"},
		{schemaID: "l4_policy", kind: "l4-policy"},
		{schemaID: "e2_service", kind: "e2-service"},
		{schemaID: "apple_banana", kind: "apple-banana"},
		{schemaID: "aws_node", kind: "aws-node"},
		{schemaID: "kubernetes_master_node", kind: "kubernetes-master-node"},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			assert.Equal(t, tt.schemaID, KindToSchemaID(tt.kind))
		})
	}
}

func TestOmitEmpty(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]interface{}
		want map[string]interface{}
	}{
		{name: "empty"},
		{
			name: "map that don't have omitted keys",
			m:    map[string]interface{}{"key": "val"},
			want: map[string]interface{}{"key": "val"},
		},
		{
			name: "map that have omitted keys but not empty",
			m:    map[string]interface{}{"key": "val", "parent_uuid": "xyz"},
			want: map[string]interface{}{"key": "val", "parent_uuid": "xyz"},
		},
		{
			name: "map that have omitted keys empty",
			m:    map[string]interface{}{"key": "val", "parent_uuid": ""},
			want: map[string]interface{}{"key": "val"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OmitEmpty(tt.m)
			assert.Equal(t, tt.want, tt.m)
		})
	}
}
