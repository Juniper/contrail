package cluster

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	nonExistingKeyPath = "./test_data/multicloud/non_existing_key"
	fakeKeyPath        = "./test_data/multicloud/test_pvt_key_fake"
	realKeyPath        = "./test_data/multicloud/test_pvt_key_real"
)

func TestStoppingSSHAgent(t *testing.T) {
	agent := &sshAgent{}
	setProperKeyPermissions(t, realKeyPath)
	assert.NoError(t, agent.Run(realKeyPath))
	assert.True(t, processExists(agent.pid))
	assert.NoError(t, agent.Stop())
	assert.False(t, processExists(agent.pid))
}

func processExists(pid string) bool {
	return exec.Command("kill", "-0", pid).Run() == nil
}

func setProperKeyPermissions(t *testing.T, file string) {
	assert.NoError(t, os.Chmod(file, 0600), "cannot set proper permissions for the key")
}

func TestRunSSHAgentWithKey(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		fails     bool
		keyExists bool
	}{
		{
			name:  "Empty key",
			fails: true,
		},
		{
			name:  "Non Existing key",
			key:   nonExistingKeyPath,
			fails: true,
		},
		{
			name:      "Existing fake key",
			key:       fakeKeyPath,
			fails:     true,
			keyExists: true,
		},
		{
			name:      "Existing real key",
			key:       realKeyPath,
			keyExists: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := &sshAgent{}
			if tt.keyExists {
				setProperKeyPermissions(t, tt.key)
			}
			err := agent.Run(tt.key)
			if tt.fails {
				assert.Error(t, err, "SSH Agent starting should fail")
			} else {
				assert.NoError(t, err, "SSH Agent did not start properly")
			}
			assert.NoError(t, agent.Stop())
		})
	}
}
