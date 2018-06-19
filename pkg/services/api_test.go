package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFilter(t *testing.T) {
	filter := ParseFilter("check==a,check==b,name==Bob")
	assert.Equal(t, []*Filter{
		{
			Key:    "check",
			Values: []string{"a", "b"},
		},
		{
			Key:    "name",
			Values: []string{"Bob"},
		},
	}, filter, "parse filter correctly")
}
