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

// Run creates ssh-agent process and adds a private key to it. AuthSockPath describes the path where agent
// authentication socket will be placed. If authSockPath is empty, default path is used.
func (s *sshAgent) Run(key string, authSockPath string) error {
	agentArgs := []string{"-s"}
	if authSockPath != "" {
		agentArgs = append(agentArgs, "-a", authSockPath)
	}
	cmdline := exec.Command("ssh-agent", agentArgs...)
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
		return errors.Errorf("cannot resolve SSH_AGENT_PID/SSH_AUTH_SOCK variables from the following output: %s",
			output)
	}
	s.authSock = authSockSubmatches[1]
	s.pid = pidSubmatches[1]
	if s.authSock == "" || s.pid == "" {
		return errors.Errorf("cannot resolve SSH_AGENT_PID/SSH_AUTH_SOCK variables from the following output: %s",
			output)
	}
	return nil
}

func (s *sshAgent) addSSHKey(key string) error {
	return s.RunCommand("ssh-add", key)
}

func (s *sshAgent) GetExportVars() []string {
	return []string{
		fmt.Sprintf("%s=%s", sshAuthSock, s.authSock),
		fmt.Sprintf("%s=%s", sshAgentPid, s.pid),
	}
}

func (s *sshAgent) AuthenticationSocket() string {
	return s.authSock
}

func (s *sshAgent) RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Env = append(os.Environ(), s.GetExportVars()...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrap(err, string(output))
	}
	return nil
}

func (s *sshAgent) Stop() error {
	if s == nil {
		return nil
	}
	return s.RunCommand("ssh-agent", "-k")
}
