package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	SetupAndRunTest(m)
}

func RunTest(t *testing.T, file string) {
	testScenario, err := LoadTest(file, nil)
	assert.NoError(t, err, "failed to load test data")
	RunCleanTestScenario(t, testScenario)
}
