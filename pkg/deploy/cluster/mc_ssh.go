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
	"regexp"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/osutil"
	"github.com/Juniper/contrail/pkg/cloud"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

const (
	SSH_AUTH_SOCK = "SSH_AUTH_SOCK"
	SSH_AGENT_PID = "SSH_AGENT_PID"
)

type sshAgent struct {
	pid      string
	authSock string
}

func StartSSHAgentWithKey(pvtKeyFile string) (*sshAgent, error) {
	agent := sshAgent{}
	err := agent.startSSHAgent()
	return agent, err
}

func (s *sshAgent) runSSHAgent(key string) error {
	cmdline := exec.Command("ssh-agent", "-s")
	cmdOutput := &bytes.Buffer{}
	cmdline.Stdout = cmdOutput
	if err := cmdline.Run(); err != nil {
		return err
	}
	agent.getPIDAndAuthSockFromSSHAgentOutput(string(cmdOutput.Bytes()))
	return addSSHKey(key)
}

func (s *sshAgent) getPIDAndAuthSockFromSSHAgentOutput(output string) {
	authSockRegex := fmt.Sprintf("%s=([^;]+)", SSH_AUTH_SOCK)
	pidRegex := fmt.Sprintf("%s=([^;]+)", SSH_AGENT_PID)
	s.authSock := regexp.MustCompile(authSockRegex).FindStringSubmatch(output)[1]
	s.pid := regexp.MustCompile(pidRegex).FindStringSubmatch(output)[1]
}

func (s *sshAgent) Kill() error {
	p, err := os.FindProcess(s.pid)
	if err != nil {
		return err
	}
	return p.Kill()
}

func (s *sshAgent) GetExportVars() []string {
	return []string{
		fmt.Sprintf("%s=%s", SSH_AUTH_SOCK, s.authSock),
		fmt.Sprintf("%s=%s", SSH_AGENT_PID, s.pid),
	}
}

func (s *sshAgent) addSSHKey(key string) error {
	cmd := exec.Command("ssh-add", key)
	cmd.Env = append(os.Environ(), s.GetExportVars()...)
	return cmd.Run()
}

// PubKeyConfig to read secret file
type PubKeyConfig struct {
	Info         map[string]string   `yaml:"public_key"`
	AwsAccessKey string              `yaml:"aws_access_key"`
	AwsSecretKey string              `yaml:"aws_secret_key"`
	AuthReg      []map[string]string `yaml:"authorized_registries"`
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

func copySSHKeyPair(srcSSHKeyDir string, destSSHKeyDir string, keyName string) error {
	srcKeyPath := filepath.Join(srcSSHKeyDir, keyName)
	destKeyPath := filepath.Join(destSSHKeyDir, keyName)
	if err := fileutil.CopyFile(srcKeyPath, destKeyPath, true); err != nil {
		return err
	}
	return fileutil.CopyFile(srcKeyPath+".pub", destKeyPath+".pub", true)
}
