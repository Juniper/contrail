package cluster

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const keyPermissions = 0600

func TestStoppingSSHAgent(t *testing.T) {
	agent := &sshAgent{}
	keyName, clean := createPrivateKey(t)
	defer clean()
	assert.NoError(t, agent.Run(keyName))
	assert.True(t, processExists(agent.pid))
	assert.NoError(t, agent.Stop())
	assert.False(t, processExists(agent.pid))
}

func createPrivateKey(t *testing.T) (string, func()) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	require.NoError(t, err, "cannot generate key")
	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)
	keyFile, clean := createTempKeyFile(t)
	_, err = keyFile.Write(pemdata)
	assert.NoError(t, err, "couldn't create key file")
	return keyFile.Name(), clean
}

func createTempKeyFile(t *testing.T) (*os.File, func()) {
	keyFile, err := ioutil.TempFile("", "testPvtKey")
	require.NoError(t, err, "private key was not created")
	require.NoError(t, keyFile.Chmod(keyPermissions), "cannot set proper permissions")
	return keyFile, func() {
		assert.NoError(t, keyFile.Close(), "couldn't close created key file")
		assert.NoError(t, os.Remove(keyFile.Name()), "couldn't remove created key file")
	}
}

func processExists(pid string) bool {
	return exec.Command("kill", "-0", pid).Run() == nil
}

func TestRunSSHAgentWithNoKey(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{
			name: "No Path",
		},
		{
			name: "Path to non existing file",
			path: "./non_existing_file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := &sshAgent{}
			assert.Error(t, agent.Run(tt.path), "SSH Agent starting should fail")
		})
	}
}

func TestRunSSHAgentWithKey(t *testing.T) {
	tests := []struct {
		name      string
		fails     bool
		createKey func(t *testing.T) (keyName string, removeKey func())
	}{
		{
			name:      "Existing fake key",
			fails:     true,
			createKey: createIncorrectKey,
		},
		{
			name:      "Existing real key",
			createKey: createPrivateKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := &sshAgent{}
			keyName, clean := tt.createKey(t)
			defer clean()
			err := agent.Run(keyName)
			if tt.fails {
				assert.Error(t, err, "SSH Agent starting should fail")
			} else {
				assert.NoError(t, err, "SSH Agent did not start properly")
			}
			assert.NoError(t, agent.Stop())
		})
	}
}

func createIncorrectKey(t *testing.T) (string, func()) {
	key, clean := createTempKeyFile(t)
	_, err := key.Write([]byte("Hi, I am a Private Key. Trust Me."))
	require.NoError(t, err, "couldn't create key file")
	return key.Name(), clean
}
