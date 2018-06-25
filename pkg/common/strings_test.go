package common

import (
	"testing"

	"github.com/gogo/protobuf/types"
	log "github.com/sirupsen/logrus"
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
		log.Debug("hogehoge")
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
