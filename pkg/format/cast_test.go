package format

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var numericTests = []struct {
	name     string
	input    interface{}
	expected int
}{
	{
		name:     "int(4) to 4",
		input:    int(4),
		expected: 4,
	},
	{
		name:     "int64(4) to 4",
		input:    int64(4),
		expected: 4,
	},
	{
		name:     "float64(4.02) to 4",
		input:    float64(4.02),
		expected: 4,
	},
	{
		// not supported type
		name:     "int8(4) to 0",
		input:    int8(4),
		expected: 0,
	},
	{
		name:     `[]byte("4") to 4`,
		input:    []byte("4"),
		expected: 4,
	},
	{
		name:     `negative []byte("-4") to -4`,
		input:    []byte("-4"),
		expected: -4,
	},
	{
		name:     `wrong string escape []byte(\""-4"\") to 0`,
		input:    []byte(`\"-4\"`),
		expected: 0,
	},
	{
		name:     `big negative []byte("-9223372036854775808") to -9223372036854775808`,
		input:    []byte("-9223372036854775808"),
		expected: -9223372036854775808,
	},
}

func TestInterfaceToInt(t *testing.T) {
	assert.Equal(t, 0, InterfaceToInt(nil))
	for _, tt := range numericTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, InterfaceToInt(tt.input))
		})
	}
}

func TestInterfaceToInt64(t *testing.T) {
	assert.Equal(t, int64(0), InterfaceToInt64(nil))

	var jsonN json.Number
	jsonN = "9223372036854775807"
	assert.Equal(t, int64(9223372036854775807), InterfaceToInt64(jsonN))
	jsonN = "-9223372036854775808"
	assert.Equal(t, int64(-9223372036854775808), InterfaceToInt64(jsonN))

	for _, tt := range numericTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, int64(tt.expected), InterfaceToInt64(tt.input))
		})
	}
}

func TestInterfaceToUint64(t *testing.T) {
	assert.Equal(t, uint64(0), InterfaceToUint64(nil))

	var jsonN json.Number
	jsonN = "9223372036854775807"
	assert.Equal(t, uint64(9223372036854775807), InterfaceToUint64(jsonN))
	jsonN = "-9223372036854775808"
	assert.Equal(t, uint64(0), InterfaceToUint64(jsonN))

	for _, tt := range numericTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected < 0 {
				assert.Equal(t, uint64(0), InterfaceToUint64(tt.input))
			} else {
				assert.Equal(t, uint64(tt.expected), InterfaceToUint64(tt.input))
			}
		})
	}
}
