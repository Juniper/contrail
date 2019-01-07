package logic

import "testing"

func TestGetIPPrefixAndPrefixLen(t *testing.T) {

	tests := []struct {
		input             string
		outputIPPrefix    string
		outputIPPrefixLen int64
		expectError       bool
	}{
		{"137.105.215.54/23", "137.105.215.54", 23, false},
		{"3.37.8.253/13", "3.37.8.253", 13, false},
		{"0.0.0.0/0", "0.0.0.0", 0, false},
		{"", "", 0, true},
		{"137.105.215.54:23", "", 0, true},
		{"137.105.215.54/23a", "", 0, true},
	}

	for _, test := range tests {
		ipPrefix, ipPrefixLen, err := getIPPrefixAndPrefixLen(test.input)

		if err != nil && !test.expectError {
			t.Errorf("Expected no error but got: \" %+v\" while parsing \"%s\".", err, test.input)
		}

		if ipPrefix != test.outputIPPrefix {
			t.Errorf("Assertion error. Expected: \"+%v\", got: \"+%v\".", test.outputIPPrefix, ipPrefix)
		}

		if ipPrefixLen != test.outputIPPrefixLen {
			t.Errorf("Assertion error. Expected: \"+%v\", got: \"+%v\".", test.outputIPPrefixLen, ipPrefixLen)
		}
	}
}
