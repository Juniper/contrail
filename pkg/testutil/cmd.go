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

// FileWritingExecutor helps to write cmd to a file (instead of executing)
type FileWritingExecutor struct {
	executedCmdTestFile string
}

// NewFileWritingExecutor returns new CommandExecutor which records executed commands into provided executedCommandsFile
func NewFileWritingExecutor(executedCommandsFile string) *FileWritingExecutor {
	return &FileWritingExecutor{executedCommandsFile}
}

// ExecCmdAndWait record string command and args into file instead of executing
func (e *FileWritingExecutor) ExecCmdAndWait(
	_ *report.Reporter, cmd string, args []string, dir string, _ ...string) error {
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

// ExecAndWait record exec.Cmd command into file instead of executing
func (e *FileWritingExecutor) ExecAndWait(r *report.Reporter, cmd *exec.Cmd) error {
	return e.ExecCmdAndWait(r, cmd.Args[0], cmd.Args[:1], cmd.Path)
}
