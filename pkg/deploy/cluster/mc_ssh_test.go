package cluster

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoppingSSHAgent(t *testing.T) {
	agent, err := startSSHAgentWithKey("./test_data/test_pvt_key_real")
	assert.NoError(t, err)

	assert.True(t, processExists(agent.pid))

	assert.NoError(t, agent.stop())

	assert.False(t, processExists(agent.pid))
}

func processExists(pid string) bool {
	return exec.Command("kill", "-0", pid).Run() == nil
}

func TestStartSSHAgentWithKey(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		fails bool
	}{
		{
			name:  "Empty key",
			fails: true,
		},
		{
			name:  "Non Existing key",
			key:   "./the_key_that_should_not_exist",
			fails: true,
		},
		{
			name:  "Existing fake key",
			key:   "./test_data/test_pvt_key_fake",
			fails: true,
		},
		{
			name: "Existing real key",
			key:  "./test_data/test_pvt_key_real",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent, err := startSSHAgentWithKey(tt.key)
			if tt.fails {
				assert.Error(t, err, "SSH Agent starting should fail")
			} else {
				assert.NoError(t, err, "SSH Agent did not start properly")
			}
			assert.NoError(t, agent.stop())
		})
	}
}
