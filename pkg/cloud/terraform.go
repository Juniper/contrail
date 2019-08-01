package cloud

import (
	"github.com/Juniper/contrail/pkg/logutil/report"
	"github.com/Juniper/contrail/pkg/osutil"
)

func destroyTopology(c *Cloud, reporter *report.Reporter) error {
	workDir := GetCloudDir(c.config.CloudID)
	cmd := "deployer"
	args := []string{"topology", "destroy"}

	if c.config.Test {
		return TestCmdHelper(cmd, args, workDir, testTemplate)
	}

	return osutil.ExecCmdAndWait(reporter, cmd, args, workDir)
}

func createTopology(c *Cloud, reporter *report.Reporter) error {
	workDir := GetCloudDir(c.config.CloudID)
	cmd := "deployer"
	args := []string{
		"topology", "create",
		"--topology", GetTopoFile(c.config.CloudID),
		"--secret", GetSecretFile(c.config.CloudID),
		"--skip_validation",
	}

	if c.config.Test {
		return TestCmdHelper(cmd, args, workDir, testTemplate)
	}

	if err := osutil.ExecCmdAndWait(reporter, cmd, args, workDir); err != nil {
		return err
	}

	args = []string{"topology", "build", "--retry", "5"}
	return osutil.ExecCmdAndWait(reporter, cmd, args, workDir)
}
