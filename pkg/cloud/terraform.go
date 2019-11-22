package cloud

import (
	"github.com/Juniper/contrail/pkg/fileutil"
)

func destroyTopology(c *Cloud) error {
	if tfStateOutputExists(c.config.CloudID) {
		if err := fileutil.WriteToFile(GetTopoFile(c.config.CloudID), []byte{}, defaultRWOnlyPerm); err != nil {
			return err
		}
		if err := updateTopology(c); err != nil {
			return err
		}
	}
	return nil
}

func updateTopology(c *Cloud) error {
	workDir := GetCloudDir(c.config.CloudID)
	cmd := "deployer"
	args := []string{
		"all", "topology",
		"--topology", GetTopoFile(c.config.CloudID),
		"--secret", GetSecretFile(c.config.CloudID),
		"--skip_validation",
	}
	return c.commandExecutor.ExecuteAndWait(c.reporter, cmd, args, workDir)
}
