package cluster

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"os/exec"
	"testing"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/stretchr/testify/assert"
)

func TestStoppingSSHAgent(t *testing.T) {
	agent := &sshAgent{}
	createPrivateKey(t, "./test_data/test_pvt_key_real")
	assert.NoError(t, agent.Run("./test_data/test_pvt_key_real"))
	assert.True(t, processExists(agent.pid))
	assert.NoError(t, agent.Stop())
	assert.False(t, processExists(agent.pid))
}

func processExists(pid string) bool {
	return exec.Command("kill", "-0", pid).Run() == nil
}

func createDumbKey(t *testing.T, destination string) func() {
	assert.NoError(t, fileutil.WriteToFile(destination, []byte("Hi, I am a Private Key. Trust Me."), 0600),
		"couldn't create key file")
	return func() {
		assert.NoError(t, os.Remove(destination), "couldn't remove created key file")
	}
}

func createPrivateKey(t *testing.T, destination string) func() {
	key, err := rsa.GenerateKey(rand.Reader, 256)
	assert.NoError(t, err, "cannot generate key")
	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)
	assert.NoError(t, fileutil.WriteToFile(destination, pemdata, 0600), "couldn't create key file")
	return func() {
		assert.NoError(t, os.Remove(destination), "couldn't remove created key file")
	}
}

func TestRunSSHAgentWithKey(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		fails     bool
		keyExists bool
		createKey func(t *testing.T, destination string) (removeKey func())
	}{
		{
			name:  "Empty key",
			fails: true,
		},
		{
			name:  "Non Existing key",
			key:   "./non_existing_key",
			fails: true,
		},
		{
			name:      "Existing fake key",
			key:       "./test_data/test_pvt_key_fake",
			fails:     true,
			keyExists: true,
			createKey: createDumbKey,
		},
		{
			name:      "Existing real key",
			key:       "./test_data/test_pvt_key_real",
			keyExists: true,
			createKey: createPrivateKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := &sshAgent{}
			if tt.keyExists {
				defer tt.createKey(t, tt.key)()
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
