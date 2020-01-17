package ansible

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/fileutil/template"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/asf/pkg/osutil"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	playbookCmd     = "ansible-playbook"
	hostPlaybookCmd = "/usr/bin/ansible-playbook"
	filePermRWOnly  = 0600
)

// Player runs Ansible playbooks.
type Player interface {
	Play()
}

// CLIClient allows to play Ansible playbooks via ansible-playbook CLI.
type CLIClient struct {
	reporter *report.Reporter
	log      *logrus.Entry

	// fields with testing purposes
	// TODO(Daniel): remove that and use provide ansible.MockClient instead
	workingDirectory string
	isTestInstance   bool
}

// NewCLIClient returns CLIClient.
func NewCLIClient(
	reporter *report.Reporter,
	logFilePath string,
	workingDirectory string,
	isTestInstance bool,
) *CLIClient {
	return &CLIClient{
		reporter:         reporter,
		log:              logutil.NewFileLogger("ansible-client", logFilePath),
		workingDirectory: workingDirectory,
		isTestInstance:   isTestInstance,
	}
}

// Play plays Ansible playbook via ansible-playbook CLI.
// It runs playbook from in give virtualenv directory if non-zero value virtualenvPath is given.
func (c *CLIClient) Play(repositoryPath string, ansibleArgs []string, virtualenvPath string) error {
	c.log.WithFields(logrus.Fields{
		"ansible-args":    ansibleArgs,
		"virtualenv-path": virtualenvPath,
	}).Info("Running playbook")

	if c.isTestInstance {
		return c.mockPlay(ansibleArgs)
	}

	cmd, err := prepareCLICmd(repositoryPath, ansibleArgs, virtualenvPath)
	if err != nil {
		return err
	}

	if err = osutil.ExecAndWait(c.reporter, cmd); err != nil {
		return errors.Wrap(err, "running playbook failed")
	}

	c.log.WithFields(logrus.Fields{
		"ansible-args":    ansibleArgs,
		"virtualenv-path": virtualenvPath,
	}).Info("Finished running playbook")
	return nil
}

func prepareCLICmd(repositoryPath string, ansibleArgs []string, virtualenvPath string) (*exec.Cmd, error) {
	var cmd *exec.Cmd
	if virtualenvPath != "" {
		var err error
		cmd, err = osutil.VenvCommand(virtualenvPath, playbookCmd, ansibleArgs...)
		if err != nil {
			return nil, err
		}
	} else {
		cmd = exec.Command(hostPlaybookCmd, ansibleArgs...)
	}

	cmd.Dir = repositoryPath
	return cmd, nil
}

// WorkingDirectory returns working directory.
func (c *CLIClient) WorkingDirectory() string {
	return c.workingDirectory
}

// IsTest returns true for testing instance.
// TODO: move test logic to separate Mock type and use dependency injection using ansible.Player interface.
func (c *CLIClient) IsTest() bool {
	return c.isTestInstance
}

func (c *CLIClient) mockPlay(ansibleArgs []string) error {
	playBookIndex := len(ansibleArgs) - 1
	content, err := template.Apply("./test_data/test_ansible_playbook.tmpl", pongo2.Context{
		"playBook":    ansibleArgs[playBookIndex],
		"ansibleArgs": strings.Join(ansibleArgs[:playBookIndex], " "),
	})
	if err != nil {
		return err
	}

	return fileutil.AppendToFile(
		filepath.Join(c.workingDirectory, "executed_ansible_playbook.yml"),
		content,
		filePermRWOnly,
	)
}
