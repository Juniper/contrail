package cluster

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/osutil"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// SSHAgentConfig related to ssh-agent process
type SSHAgentConfig struct {
	AuthSock string `yaml:"auth_sock"`
	PID      string `yaml:"pid"`
}

// PubKeyConfig to read secret file
type PubKeyConfig struct {
	Info         map[string]string   `yaml:"public_key"`
	AwsAccessKey string              `yaml:"aws_access_key"`
	AwsSecretKey string              `yaml:"aws_secret_key"`
	AuthReg      []map[string]string `yaml:"authorized_registries"`
}

// nolint: gocyclo
func (m *multiCloudProvisioner) manageSSHAgent(action string) error {
	// To-Do: to use ssh/agent library to create agent unix process
	sshAgentPath := m.getSSHAgentFile()
	if _, err := os.Stat(sshAgentPath); err != nil {
		if action == deleteAction {
			return nil
		}
		if _, err = m.runSSHAgent(sshAgentPath); err != nil {
			return err
		}
	} else {
		process, err := isSSHAgentProcessRunning(sshAgentPath) // nolint: govet
		if err != nil {
			if _, err = m.runSSHAgent(sshAgentPath); err != nil {
				return err
			}
		}

		if action == deleteAction {
			// sends signal to pgid if pid is a negative number
			// this helps in killing all the child (zombie leftover process)
			return syscall.Kill(-process.Pid, syscall.SIGKILL)
		}
	}
	if m.cluster.config.Test {
		keypairDir := filepath.Join(m.getMCWorkingDir(m.getWorkingDir()), defaultSSHKeyRepo)
		if err := os.MkdirAll(keypairDir, 0755); err != nil {
			return err
		}
		pubKey, err := m.readPubKeyConfig()
		if err != nil {
			return err
		}
		sshKeyName, ok := pubKey.Info["name"]
		if !ok {
			return errors.New("secret file format is not valid")
		}
		if err = fileutil.WriteToFile(
			filepath.Join(keypairDir, sshKeyName), []byte("test pvt key"), defaultFilePermRWOnly,
		); err != nil {
			return err
		}
		if err = fileutil.WriteToFile(
			filepath.Join(keypairDir, sshKeyName+".pub"), []byte("test pub key"), defaultFilePermRWOnly,
		); err != nil {
			return err
		}
		return nil
	}

	sshAgentConf, err := readSSHAgentConfig(sshAgentPath)
	if err != nil {
		return err
	}

	if err = os.Setenv("SSH_AUTH_SOCK", sshAgentConf.AuthSock); err != nil {
		return err
	}

	if err = os.Setenv("SSH_AGENT_PID", sshAgentConf.PID); err != nil {
		return err
	}

	m.Log.Debugf("SSH_AUTH_SOCK: %s", os.Getenv("SSH_AUTH_SOCK"))
	m.Log.Debugf("SSH_AGENT_PID: %s", os.Getenv("SSH_AGENT_PID"))

	pubKey, err := m.readPubKeyConfig()
	if err != nil {
		return err
	}

	pubCloudID, _, err := m.getPubPvtCloudID()
	if err != nil {
		return err
	}

	pubCloudSSHKeyDir := filepath.Join(cloud.GetCloudDir(pubCloudID), defaultSSHKeyRepo)
	clusterSSHKeyDir := filepath.Join(m.getMCWorkingDir(m.getWorkingDir()), defaultSSHKeyRepo)
	sshKeyName, ok := pubKey.Info["name"]

	if !ok {
		return errors.New("secret file format is not valid")
	}

	if err = copySSHKeyPair(pubCloudSSHKeyDir, clusterSSHKeyDir, sshKeyName); err != nil {
		return err
	}

	sshPvtKeyPath := filepath.Join(clusterSSHKeyDir, sshKeyName)
	return m.addPvtKeyToSSHAgent(sshPvtKeyPath)
}

func (m *multiCloudProvisioner) runSSHAgent(sshAgentPath string) (*SSHAgentConfig, error) {
	cmd := "ssh-agent"
	args := []string{"-s"}
	cmdline := exec.Command(cmd, args...)

	cmdOutput := &bytes.Buffer{}
	cmdline.Stdout = cmdOutput
	// set pgid enables creation of all child process with pgid of parent process
	cmdline.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	if err := cmdline.Run(); err != nil {
		return nil, err
	}

	stdout := string(cmdOutput.Bytes())

	stdoutList := strings.SplitAfterN(stdout, "=", 3)
	sockIDList := strings.SplitAfterN(stdoutList[1], ";", 2)
	pidList := strings.SplitAfterN(stdoutList[2], ";", 2)

	sshAgentConf := &SSHAgentConfig{
		AuthSock: strings.TrimSuffix(sockIDList[0], ";"),
		PID:      strings.TrimSuffix(pidList[0], ";"),
	}

	if err := m.writeSSHAgentConfig(sshAgentConf, sshAgentPath); err != nil {
		return nil, err
	}
	return sshAgentConf, nil
}

func (m *multiCloudProvisioner) writeSSHAgentConfig(conf *SSHAgentConfig, dest string) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	return fileutil.WriteToFile(dest, data, defaultFilePermRWOnly)
}

func isSSHAgentProcessRunning(sshAgentPath string) (*os.Process, error) {
	sshAgentConf, err := readSSHAgentConfig(sshAgentPath)
	if err != nil {
		return nil, err
	}
	pid, err := strconv.Atoi(sshAgentConf.PID) // nolint: govet
	if err != nil {
		return nil, err
	}
	// check if process is running
	process, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}
	if err = process.Signal(syscall.Signal(0)); err != nil {
		return nil, err
	}
	return process, nil
}

func readSSHAgentConfig(path string) (*SSHAgentConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	agentConfig := &SSHAgentConfig{}
	if err = yaml.UnmarshalStrict(data, agentConfig); err != nil {
		return nil, err
	}
	return agentConfig, nil
}

func (m *multiCloudProvisioner) readPubKeyConfig() (*PubKeyConfig, error) {
	secretFilePath := m.getClusterSecretFile()
	data, err := ioutil.ReadFile(secretFilePath)
	if err != nil {
		return nil, err
	}

	pubKeyConfig := &PubKeyConfig{}
	if err = yaml.Unmarshal(data, pubKeyConfig); err != nil {
		return nil, err
	}
	return pubKeyConfig, nil
}

func (m *multiCloudProvisioner) addPvtKeyToSSHAgent(keyPath string) error {
	cmd := "ssh-add"
	args := []string{"-d", fmt.Sprintf("%s", keyPath)}
	m.Log.Debugf("Executing command: %s", fmt.Sprintf("%s %s", cmd, keyPath))

	if m.cluster.config.Test {
		return cloud.TestCmdHelper(cmd, args, m.workDir, testTemplate)
	}
	// ignore if there is an error while deleting key,
	// this step is to make sure that updated keys are loaded to agent
	_ = osutil.ExecCmdAndWait(m.Reporter, cmd, args, m.workDir) // nolint: errcheck

	// readd the key
	args = []string{fmt.Sprintf("%s", keyPath)}
	m.Log.Debugf("Executing command: %s", fmt.Sprintf("%s %s", cmd, keyPath))

	if m.cluster.config.Test {
		return cloud.TestCmdHelper(cmd, args, m.workDir, testTemplate)
	}
	return osutil.ExecCmdAndWait(m.Reporter, cmd, args, m.workDir)
}

func copySSHKeyPair(srcSSHKeyDir string, destSSHKeyDir string, keyName string) error {
	if err := fileutil.CopyFile(
		filepath.Join(srcSSHKeyDir, keyName), filepath.Join(destSSHKeyDir, keyName), true,
	); err != nil {
		return err
	}
	return fileutil.CopyFile(
		filepath.Join(srcSSHKeyDir, keyName+".pub"), filepath.Join(destSSHKeyDir, keyName+".pub"), true,
	)
}

func (m *multiCloudProvisioner) copySSHKeyPairToMC() error {
	pubKey, err := m.readPubKeyConfig()
	if err != nil {
		return err
	}
	sourceDir := filepath.Join(m.getMCWorkingDir(m.getWorkingDir()), defaultSSHKeyRepo)
	keyName, ok := pubKey.Info["name"]

	if !ok {
		return errors.New("secret file format is not valid")
	}

	var destinationDir string
	if m.cluster.config.Test {
		destinationDir = m.getWorkingDir()
	} else {
		destinationDir = m.getMCDeployerRepoDir()
	}
	return copySSHKeyPair(sourceDir, filepath.Join(destinationDir, "keys"), keyName)
}
