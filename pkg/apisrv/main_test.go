package apisrv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	SetupAndRunTest(m)
}

func RunTest(t *testing.T, file string) {
	testScenario, err := LoadTest(file)
	assert.NoError(t, err, "failed to load test data")
	RunTestScenario(t, testScenario)
}

func LoadTest(file string) (*TestScenario, error) {
	var testScenario TestScenario
	err := LoadTestScenario(&testScenario, file)
	return &testScenario, err

}
