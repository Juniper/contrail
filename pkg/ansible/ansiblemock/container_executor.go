package ansiblemock

import (
	"context"
	"testing"

	"github.com/Juniper/contrail/pkg/ansible"
	"github.com/stretchr/testify/assert"
)

type ContainerExecution struct {
	Cmd        []string
	Parameters *ansible.ContainerParameters
}

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
		assert.Fail(m.t, "invalid number of executions", "expected %v executions, got %v",
			len(m.executions), len(expectedExecutions))
	} else {
		for i := range m.executions {
			assert.Equal(m.t, expectedExecutions[i].Cmd, m.executions[i].Cmd, "executed command do not match")
			assert.Equal(m.t, expectedExecutions[i].Parameters, m.executions[i].Parameters,
				"executed parameters do not match")
		}
	}
	m.executions = []ContainerExecution{}
}

func (m *MockContainerExecutor) StartExecuteAndRemove(
	ctx context.Context, cp *ansible.ContainerParameters, cmd []string,
) error {
	m.executions = append(m.executions, ContainerExecution{Cmd: cmd, Parameters: cp})
	return nil
}
