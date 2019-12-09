package cloud

import (
	"github.com/Juniper/asf/pkg/fileutil"
)

func destroyTopology(c *Cloud, providers []string) error {
	if tfStateOutputExists(c.config.CloudID) {
		if err := fileutil.WriteToFile(GetTopoFile(c.config.CloudID), []byte{}, defaultRWOnlyPerm); err != nil {
			return err
		}
		if err := updateTopology(c, providers); err != nil {
			return err
		}
	}
	return nil
}

func updateTopology(c *Cloud, providers []string) error {
	workDir := GetCloudDir(c.config.CloudID)
	cmd := "deployer"
	args := []string{
		"all", "topology",
		"--topology", GetTopoFile(c.config.CloudID),
		"--secret", GetSecretFile(c.config.CloudID),
		"--skip_validation",
	}
	for _, p := range providers {
		args = append(args, "--"+p)
	}
	return c.commandExecutor.ExecuteCmdAndWait(c.reporter, cmd, args, workDir)
}
