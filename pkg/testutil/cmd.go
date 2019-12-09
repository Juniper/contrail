package testutil

import (
	"path/filepath"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/fileutil/template"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/flosch/pongo2"
)

const (
	defaultTestTemplate        = "./test_data/test_cmd.tmpl"
	defaultExecutedCmdTestFile = "executed_cmd.yml"
	rwOnlyPerm                 = 0600
)

// TestCommandExecutor helps to write cmd to a file (instead of executing)
type TestCommandExecutor struct {
	testTemplate        string
	executedCmdTestFile string
}

func NewTestCommandExecutor() *TestCommandExecutor {
	return &TestCommandExecutor{defaultTestTemplate, defaultExecutedCmdTestFile}
}

func (e *TestCommandExecutor) ExecuteAndWait(_ *report.Reporter, cmd string, args []string, dir string, _ ...string) error {
	content, err := template.Apply(e.testTemplate, pongo2.Context{
		"cmd":  cmd,
		"args": args,
	})
	if err != nil {
		return err
	}

	destPath := filepath.Join(dir, e.executedCmdTestFile)
	return fileutil.AppendToFile(destPath, content, rwOnlyPerm)
}
