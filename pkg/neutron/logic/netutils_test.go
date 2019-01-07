package logic

import "testing"

func TestGetIPPrefixAndPrefixLen(t *testing.T) {

	tests := []struct {
		name                 string
		input                string
		outputIPPrefix       string
		outputIPNetworPrefix string
		outputIPPrefixLen    int64
		expectError          bool
	}{
		{"good ip 1", "137.105.215.54/23", "137.105.215.54", "137.105.214.0", 23, false},
		{"good ip 2", "3.37.8.253/13", "3.37.8.253", "3.32.0.0", 13, false},
		{"zero ipv4 address", "0.0.0.0/0", "0.0.0.0", "0.0.0.0", 0, false},
		{"zero ipv6 address", "::/0", "::", "::", 0, false},
		{"empty ip address", "", "", "", 0, true},
		{"wrong ip - invalid separator", "137.105.215.54:23", "", "", 0, true},
		{"wrong ip - invalid mask", "137.105.215.54/23a", "", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ipPrefix, ipNetwork, ipPrefixLen, err := getIPPrefixAndPrefixLen(tt.input)

			if err != nil && !tt.expectError {
				t.Errorf("Expected no error but got: '%+v'", err)
			}

			if ipPrefix != tt.outputIPPrefix {
				t.Errorf("Assertion error. Expected: '%+v', got: '%+v'", ipPrefix, tt.outputIPPrefix)
			}

			if ipNetwork != tt.outputIPNetworPrefix {
				t.Errorf("Assertion error. Expected: '%+v', got: '%+v'", ipNetwork, tt.outputIPNetworPrefix)
			}

			if ipPrefixLen != tt.outputIPPrefixLen {
				t.Errorf("Assertion error. Expected: '%+v', got: '%+v'", ipPrefixLen, tt.outputIPPrefixLen)
			}
		})
	}
}
