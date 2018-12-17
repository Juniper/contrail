package osutil

import (
	"os"
	"os/exec"

	"github.com/Juniper/contrail/pkg/log/report"
)

// ExecCmdAndWait execs cmd, reports the stdout & stderr and waits for cmd to complete
func ExecCmdAndWait(r *report.Reporter, cmd string,
	args []string, dir string) error {

	// create dir, if it does not exist
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	cmdline := exec.Command(cmd, args...)
	if dir != "" {
		cmdline.Dir = dir
	}
	stdout, err = cmdline.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmdline.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmdline.Start(); err != nil {
		return err
	}
	// Report progress log periodically to stdout/db
	go r.ReportLog(stdout)
	go r.ReportLog(stderr)

	// wait for command to complete
	return cmdline.Wait()
}
