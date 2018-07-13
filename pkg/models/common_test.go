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
		{name: "empty childName", parentFQName: []string{"grandparent", "parent"},
			want: []string{"grandparent", "parent"}},
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
