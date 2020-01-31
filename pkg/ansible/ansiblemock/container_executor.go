package ansiblemock

import (
	"context"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/ansible"
	"github.com/stretchr/testify/assert"

	yaml "gopkg.in/yaml.v2"
)

type ContainerExecution struct {
	Cmd        []string
	Parameters *ansible.ContainerParameters
}

func (c *ContainerExecution) String() string {
	out, err := yaml.Marshal(c)
	if err != nil {
		return ""
	}
	return string(out)
}

// MockContainerExecutor collects all executions that were called via Container Executor.
type MockContainerExecutor struct {
	workingDirectory string
	t                *testing.T
	executions       []ContainerExecution
}

// NewMockContainerExecutor returns new testing container executor.
func NewMockContainerExecutor(t *testing.T) *MockContainerExecutor {
	return &MockContainerExecutor{t: t}
}

func (m *MockContainerExecutor) AssertAndClear(expectedExecutions []ContainerExecution) {
	if len(m.executions) != len(expectedExecutions) {
		assert.Fail(
			m.t, "invalid number of executions", "expected %v executions, got %v."+
				"\nExpected executions: \n%s\nActual executions: %s",
			len(expectedExecutions), len(m.executions), executions(expectedExecutions), executions(m.executions))
	} else {
		for i := range m.executions {
			assert.Equal(m.t, expectedExecutions[i].Cmd, m.executions[i].Cmd, "executed command do not match")
			assert.Equal(m.t, expectedExecutions[i].Parameters, m.executions[i].Parameters,
				"executed parameters do not match")
		}
	}
	m.executions = []ContainerExecution{}
}

func executions(executions []ContainerExecution) string {
	result := []string{}
	for _, e := range executions {
		result = append(result, e.String())
	}
	return strings.Join(result, "\n")
}

func (m *MockContainerExecutor) StartExecuteAndRemove(
	ctx context.Context, cp *ansible.ContainerParameters, cmd []string,
) error {
	m.executions = append(m.executions, ContainerExecution{Cmd: cmd, Parameters: cp})
	return nil
}
