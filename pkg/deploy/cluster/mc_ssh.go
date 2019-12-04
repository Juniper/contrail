package cluster

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"

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

func (s *sshAgent) runSSHAgent(key string) error {
	cmdline := exec.Command("ssh-agent", "-s")
	cmdOutput := &bytes.Buffer{}
	cmdline.Stdout = cmdOutput
	if err := cmdline.Run(); err != nil {
		return errors.Wrap(err, "cannot start ssh-agent")
	}
	if err := s.getPIDAndAuthSockFromSSHAgentOutput(string(cmdOutput.Bytes())); err != nil {
		return errors.Wrap(err, "ssh-agent started but its variables are not resolved")
	}
	return s.addSSHKey(key)
}

func (s *sshAgent) getPIDAndAuthSockFromSSHAgentOutput(output string) error {
	authSockRegex := fmt.Sprintf("%s=(.+?);", sshAuthSock)
	pidRegex := fmt.Sprintf("%s=(.+?);", sshAgentPid)
	authSockSubmatches := regexp.MustCompile(authSockRegex).FindStringSubmatch(output)
	pidSubmatches := regexp.MustCompile(pidRegex).FindStringSubmatch(output)
	if len(authSockSubmatches) != 2 || len(pidSubmatches) != 2 {
		return errors.Errorf("Cannot resolve SSH_AGENT_PID/SSH_AUTH_SOCK variables from the following output: %s",
			output)
	}
	s.authSock = authSockSubmatches[1]
	s.pid = pidSubmatches[1]
	if s.authSock == "" || s.pid == "" {
		return errors.Errorf("Cannot resolve SSH_AGENT_PID/SSH_AUTH_SOCK variables from the following output: %s",
			output)
	}
	return nil
}

func (s *sshAgent) addSSHKey(key string) error {
	if err := os.Chmod(key, 0600); err != nil {
		return errors.Wrap(err, "cannot set proper permissions for the key")
	}
	return s.runCommand("ssh-add", key)
}

func (s *sshAgent) getExportVars() []string {
	return []string{
		fmt.Sprintf("%s=%s", sshAuthSock, s.authSock),
		fmt.Sprintf("%s=%s", sshAgentPid, s.pid),
	}
}

func (s *sshAgent) runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Env = append(os.Environ(), s.getExportVars()...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrap(err, string(output))
	}
	return nil
}

func (s *sshAgent) stop() error {
	if s == nil {
		return nil
	}
	return s.runCommand("ssh-agent", "-k")
}
