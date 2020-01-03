package cloud

import (
	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/sirupsen/logrus"
)

type providers []string

type dryRunExecutor struct {
	log *logrus.Entry
}

func (d dryRunExecutor) ExecCmdAndWait(r *report.Reporter, cmd string, args []string, dir string, envVars ...string) error {
	d.log.Infof("Dry runing %s command with [%s] args in %s directory with [%s] env", cmd, args, dir, envVars)
	return nil
}

func newDeployer(cloudID string, dryRun bool, executor CommandExecutor, reporter *report.Reporter, log *logrus.Entry) deployer {
	if dryRun {
		executor = dryRunExecutor{log}
	}
	return deployer{cloudID, executor, reporter}
}

type deployer struct {
	cloudID         string
	commandExecutor CommandExecutor
	reporter        *report.Reporter
}

func (d deployer) destroyTopology(providers providers) error {
	if tfStateOutputExists(d.cloudID) {
		if err := fileutil.WriteToFile(GetTopoFile(d.cloudID), []byte{}, defaultRWOnlyPerm); err != nil {
			return err
		}
		if err := d.updateTopology(providers); err != nil {
			return err
		}
	}
	return nil
}

func (d deployer) updateTopology(providers providers) error {
	workDir := GetCloudDir(d.cloudID)
	cmd := "deployer"
	args := []string{
		"all", "topology",
		"--topology", GetTopoFile(d.cloudID),
		"--secret", GetSecretFile(d.cloudID),
		"--skip_validation",
	}
	for _, p := range providers {
		args = append(args, []string{"--limit", p}...)
	}
	return d.commandExecutor.ExecCmdAndWait(d.reporter, cmd, args, workDir)
}
