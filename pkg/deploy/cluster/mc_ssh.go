package cluster

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

const (
	sshAuthSock = "SSH_AUTH_SOCK"
	sshAgentPid = "SSH_AGENT_PID"
)

type sshAgent struct {
	pid      string
	authSock string
}

func startSSHAgentWithKey(pvtKeyFile string) (*sshAgent, error) {
	agent := &sshAgent{}
	err := agent.runSSHAgent(pvtKeyFile)
	return agent, err
}

func (s *sshAgent) runSSHAgent(key string) error {
	cmdline := exec.Command("ssh-agent", "-s")
	cmdOutput := &bytes.Buffer{}
	cmdline.Stdout = cmdOutput
	if err := cmdline.Run(); err != nil {
		return errors.Wrap(err, "cannot start ssh-agent")
	}
	s.getPIDAndAuthSockFromSSHAgentOutput(string(cmdOutput.Bytes()))
	return s.addSSHKey(key)
}

func (s *sshAgent) getPIDAndAuthSockFromSSHAgentOutput(output string) {
	authSockRegex := fmt.Sprintf("%s=([^;]+)", sshAuthSock)
	pidRegex := fmt.Sprintf("%s=([^;]+)", sshAgentPid)
	s.authSock = regexp.MustCompile(authSockRegex).FindStringSubmatch(output)[1]
	s.pid = regexp.MustCompile(pidRegex).FindStringSubmatch(output)[1]
}

func (s *sshAgent) addSSHKey(key string) error {
	if err := os.Chmod(key, 0600); err != nil {
		return errors.Wrap(err, "cannot set proper permissions for the key")
	}
	cmd := exec.Command("ssh-add", key)
	cmd.Env = append(os.Environ(), s.getExportVars()...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrap(err, string(output))
	}
	return nil
}

func (s *sshAgent) getExportVars() []string {
	return []string{
		fmt.Sprintf("%s=%s", sshAuthSock, s.authSock),
		fmt.Sprintf("%s=%s", sshAgentPid, s.pid),
	}
}

func (s *sshAgent) kill() error {
	if s == nil {
		return nil
	}
	pid, err := strconv.Atoi(s.pid)
	if err != nil {
		return errors.Wrap(err, "cannot parse SSH Agent pid")
	}
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return p.Kill()
}
