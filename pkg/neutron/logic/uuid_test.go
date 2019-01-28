package logic

import "testing"

func TestContrailUUIDToNeutronID(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"ff93b90f-74d0-4214-980a-86fbdf952ab8", "ff93b90f74d04214980a86fbdf952ab8"},
		{"abcdefgh-abcd-abcd-abcd-abcdefghijkl", "abcdefghabcdabcdabcdabcdefghijkl"},
		{"dde59c13739e46ac9f5cfe0d7f1c4567", "dde59c13739e46ac9f5cfe0d7f1c4567"},
		{"", ""},
	}

	for _, test := range tests {
		result := ContrailUUIDToNeutronID(test.input)
		if result != test.output {
			t.Errorf("Tranlating contrail uuid (%s) to neutron was incorrect, got: %s, want: %s.",
				test.input, result, test.output)
		}
	}
}

func TestNeutronIDToContrailUUID(t *testing.T) {
	tests := []struct {
		input       string
		output      string
		expectError bool
	}{
		{"ff93b90f74d04214980a86fbdf952ab8", "ff93b90f-74d0-4214-980a-86fbdf952ab8", false},
		{"abcdefghabcdabcdabcdabcdefghijkl", "", true},
		{"", "", true},
		{"\n", "", true},
		{"ff93b90f74d04214980a86fbdf952ab8ff93b90f74d04214980a86fbdf952ab8ff93b90f74d04214980a86fbdf952ab8ff93b" +
			"90f74d04214980a86fbdf952ab8ff93b90f74d04214980a86fbdf952ab8ff93b90f74d04214980a86fbdf952ab8ff93b90f74d04" +
			"214980a86fbdf952ab8ff93b90f74d04214980a86fbdf952ab8ff93b90f74d04214980a86fbdf952ab8ff93b90f74d04214980a8",
			"", true},
	}

	for _, test := range tests {
		result, err := neutronIDToContrailUUID(test.input)

		if err != nil && !test.expectError {
			t.Errorf("Expected no error but got: \" %+v\" while parsing \"%s\".", err, test.input)
		}
		if result != test.output {
			t.Errorf("Tranlating neutron id (%s) to contrail uuid was incorrect, got: %s, want: %s.",
				test.input, result, test.output)
		}
	}
}
