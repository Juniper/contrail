package schema

import (
	"testing"

	log "github.com/sirupsen/logrus"
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
