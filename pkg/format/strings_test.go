package format

import (
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCamelToSnake(t *testing.T) {
	testDatas := [][]string{
		{"APIList", "api_list"},
		{"L4Policy", "l4_policy"},
		{"E2Service", "e2_service"},
		{"AppleBanana", "apple_banana"},
		{"AwsNode", "aws_node"},
	}
	for _, testData := range testDatas {
		camelToSnake := CamelToSnake(testData[0])
		snakeToCamel := SnakeToCamel(testData[1])
		logrus.Debug("hogehoge")
		if camelToSnake != testData[1] {
			t.Fatal("CamelToSnake failed expected", testData[1], " got ", camelToSnake)
		}
		if snakeToCamel != testData[0] {
			t.Fatal("SnakeToCamel failed expected", testData[0], " got ", snakeToCamel)
		}
	}
}

func TestCheckPath(t *testing.T) {
	testData := []struct {
		path      []string
		fieldMask *types.FieldMask
		fails     bool
	}{
		{
			path:      []string{"first", "test"},
			fieldMask: &types.FieldMask{Paths: []string{"first.test", "second.test", "third"}},
			fails:     false,
		},
		{
			path:      []string{"second", "test", "case"},
			fieldMask: &types.FieldMask{Paths: []string{"first.test", "second.test.case", "third"}},
			fails:     false,
		},
		{
			path:      []string{"third"},
			fieldMask: &types.FieldMask{Paths: []string{"first.test", "second.test", "third"}},
			fails:     false,
		},
		{
			path:      []string{"bad", "test"},
			fieldMask: &types.FieldMask{Paths: []string{"first.test", "second.test", "third"}},
			fails:     true,
		},
		{
			path:      []string{"it", "should", "fail"},
			fieldMask: &types.FieldMask{Paths: []string{"first.test", "second.test", "third"}},
			fails:     true,
		},
		{
			path:      []string{"fail"},
			fieldMask: &types.FieldMask{Paths: []string{"first.test", "second.test", "third"}},
			fails:     true,
		},
	}

	for _, tt := range testData {
		isInFieldMask := CheckPath(tt.fieldMask, tt.path)
		if tt.fails {
			assert.False(t, isInFieldMask)
		} else {
			assert.True(t, isInFieldMask)
		}
	}

}

func TestRemoveFromStringSlice(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		values   map[string]struct{}
		expected []string
	}{
		{
			name:     "does nothing when no values given",
			slice:    []string{"foo", "bar", "baz", "hoge"},
			values:   nil,
			expected: []string{"foo", "bar", "baz", "hoge"},
		},
		{
			name:  "removes single value from slice",
			slice: []string{"foo", "bar", "baz", "hoge"},
			values: map[string]struct{}{
				"bar": {},
			},
			expected: []string{"foo", "baz", "hoge"},
		},
		{
			name:  "removes multiple values from slice",
			slice: []string{"foo", "bar", "baz", "hoge"},
			values: map[string]struct{}{
				"bar": {},
				"baz": {},
			},
			expected: []string{"foo", "hoge"},
		},
		{
			name:  "removes all values from slice given all values",
			slice: []string{"foo", "bar", "baz", "hoge"},
			values: map[string]struct{}{
				"foo":  {},
				"bar":  {},
				"baz":  {},
				"hoge": {},
			},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RemoveFromStringSlice(tt.slice, tt.values)

			assert.Equal(t, tt.expected, s)
		})
	}
}
