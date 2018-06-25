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
	testDataCorrect := [][]string{
		{"first", "test"},
		{"second", "test"},
		{"third"},
		{"last", "more", "complex"},
	}
	testDataIncorrect := [][]string{
		{"bad", "test"},
		{"fail"},
		{"it", "should", "fail"},
	}
	fieldMask := types.FieldMask{Paths: []string{"first.test", "second.test", "third", "last.more.complex"}}
	for _, testData := range testDataCorrect {
		assert.True(t, CheckPath(&fieldMask, testData))
	}
	for _, testData := range testDataIncorrect {
		assert.False(t, CheckPath(&fieldMask, testData))
	}
}
