package models

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
