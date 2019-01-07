package logic

import (
	"testing"
)

// TODO make full unit tests to cover all cases of security group.

func TestSecurityGroup_Create(t *testing.T) {
	tests := []struct {
		name        string
		expectedErr bool
	}{
		{
			name:        "create security group",
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO implement it! Look into port_test.go as example.
		})
	}

}
