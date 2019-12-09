package testutil

import (
	"os/exec"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/flosch/pongo2"
)

const (
	rwOnlyPerm          = 0600
	executedCmdTemplate = "{{ cmd }}{% for arg in args %} {{ arg }}{% endfor %}\n"
)

// TestCommandExecutor helps to write cmd to a file (instead of executing)
type TestCommandExecutor struct {
	executedCmdTestFile string
}

// NewTestCommandExecutor returns new CommandExecutor which records executed commands into provided executedCommandsFile
func NewTestCommandExecutor(executedCommandsFile string) *TestCommandExecutor {
	return &TestCommandExecutor{executedCommandsFile}
}

// ExecuteCmdAndWait record string command and args into file instead of executing
func (e *TestCommandExecutor) ExecuteCmdAndWait(_ *report.Reporter, cmd string, args []string, dir string, _ ...string) error {
	template, err := pongo2.FromString(executedCmdTemplate)
	if err != nil {
		return err
	}
	output, err := template.ExecuteBytes(pongo2.Context{
		"cmd":  cmd,
		"args": args,
	})
	if err != nil {
		return err
	}

	return fileutil.AppendToFile(e.executedCmdTestFile, output, rwOnlyPerm)
}

// ExecuteAndWait record exec.Cmd command into file instead of executing
func (e *TestCommandExecutor) ExecuteAndWait(r *report.Reporter, cmd *exec.Cmd) error {
	return e.ExecuteCmdAndWait(r, cmd.Args[0], cmd.Args[:1], cmd.Path)
}
