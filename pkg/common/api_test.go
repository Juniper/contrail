package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFilter(t *testing.T) {
	filter := ParseFilter("check==a,check==b,name==Bob")
	assert.Equal(t, Filter{
		"check": []string{"a", "b"},
		"name":  []string{"Bob"},
	}, filter, "parse filter correctly")
}
