package cloud

import (
	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/ansible"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

func destroyTopology(c *Cloud, providers []string, clusterUUID string) error {
	if tfStateOutputExists(c.config.CloudID) {
		if err := fileutil.WriteToFile(GetTopoFile(c.config.CloudID), []byte{}, defaultRWOnlyPerm); err != nil {
			return err
		}
		if err := updateTopology(c, providers, clusterUUID); err != nil {
			return err
		}
	}
	return nil
}

func updateTopology(c *Cloud, providers []string, clusterUUID string) error {
	workDir := GetCloudDir(c.config.CloudID)
	cmd := "deployer"
	args := []string{
		"all", "topology",
		"--topology", GetTopoFile(c.config.CloudID),
		"--secret", GetSecretFile(c.config.CloudID),
		"--skip_validation",
	}
	for _, p := range providers {
		args = append(args, []string{"--limit", p}...)
	}
	volumes := []ansible.Volume{
		ansible.Volume{
			Source: workDir,
			Target: workDir,
		},
	}
	player, err := ansible.NewContainerPlayer(c.reporter, c.config.LogFile)
	if err != nil {
		return errors.Wrap(err, "cannot update topology without container player")
	}

	registry, err := c.getRegistries(clusterUUID)
	if err != nil {
		return errors.Wrap(err, "cannot get registry username and password")
	}
	imageRef, err := c.getMultiCloudImageRef(clusterUUID)
	if err != nil {
		return err
	}

	return player.StartExecuteAndRemove(
		c.ctx, imageRef, registry.Username, registry.Password, volumes, workDir,
		append([]string{cmd}, args...), []string{})
}

func (c *Cloud) getMultiCloudImageRef(clusterUUID string) (string, error) {
	clusterResp, err := c.APIServer.GetContrailCluster(c.ctx, &services.GetContrailClusterRequest{
		ID: clusterUUID,
	})
	if err != nil {
		return "", errors.Wrap(err, "cannot resolve MultiCloud Image Ref")
	}
	return ansible.GetContainerName(clusterResp.ContrailCluster.ContainerRegistry,
		MultiCloudContainer, ansible.GetContrailVersion(clusterResp.ContrailCluster, c.log))
}
