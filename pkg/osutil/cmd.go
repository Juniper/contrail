package osutil

import (
	"os/exec"

	"github.com/Juniper/contrail/pkg/logutil/report"
)

// ExecCmdAndWait execs cmd, reports the stdout & stderr and waits for cmd to complete.
func ExecCmdAndWait(r *report.Reporter, cmd string, args []string, dir string) error {
	cmdline := exec.Command(cmd, args...)
	if dir != "" {
		cmdline.Dir = dir
	}

	return ExecAndWait(r, cmdline)
}

// ExecAndWait execs cmd, reports the stdout & stderr and waits for cmd to complete.
func ExecAndWait(r *report.Reporter, cmd *exec.Cmd) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	// Report progress log periodically to stdout/db
	go r.ReportLog(stdout)
	go r.ReportLog(stderr)

	return cmd.Wait()
}
