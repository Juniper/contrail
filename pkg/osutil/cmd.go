package osutil

import (
	"os/exec"

	"github.com/Juniper/contrail/pkg/logutil/report"
)

// ExecCmdAndWait execs cmd, reports the stdout & stderr and waits for cmd to complete
func ExecCmdAndWait(r *report.Reporter, cmd string,
	args []string, dir string) error {

	cmdline := exec.Command(cmd, args...)
	if dir != "" {
		cmdline.Dir = dir
	}
	stdout, err := cmdline.StdoutPipe()
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
