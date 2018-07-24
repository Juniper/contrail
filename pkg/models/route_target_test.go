package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRouteTarget(t *testing.T) {

	tests := []struct {
		name        string
		routeTarget []string
		fails       bool
	}{
		{
			name:        "Try to parse route target with ip in name",
			routeTarget: []string{"target", "10.0.0.0", "8000000"},
			fails:       false,
		},
		{
			name:        "Try to parse route target with asn in name",
			routeTarget: []string{"target", "123456", "5123"},
			fails:       false,
		},
		{
			name:        "Try to parse route target with wrong prefix",
			routeTarget: []string{"rt", "10.0.0.0", "8000000"},
			fails:       true,
		},
	}

	for _, tt := range tests {
		_, _, _, err := ParseRouteTarget(tt.routeTarget)
		if tt.fails {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
