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
	keyPaths := services.NewKeyFileDefaults()
	registry, err := c.getRegistries(clusterUUID)
	if err != nil {
		return errors.Wrap(err, "cannot get registires username and password")
	}
	imageRef, err := c.getMultiCloudImageRef(clusterUUID)
	if err != nil {
		return errors.Wrap(err, "cannot get MultiCloud image reference")
	}

	return c.containerExecutor.StartExecuteAndRemove(
		c.ctx,
		&ansible.ContainerParameters{
			ImageRef:         imageRef,
			ImageRefUsername: registry.Username,
			ImageRefPassword: registry.Password,
			HostVolumes: []ansible.Volume{
				{
					Source: workDir,
					Target: workDir,
				},
				{
					Source: keyPaths.KeyHomeDir,
					Target: keyPaths.KeyHomeDir,
				},
			},
			ForceContainerRecreate: true,
			WorkingDirectory:       workDir,
			RemoveContainer:        true,
			OverwriteEntrypoint:    true,
			ContainerPrefix:        MultiCloudContainerPrefix,
			HostNetwork:            true,
			Privileged:             true,
		},
		append([]string{cmd}, args...),
	)
}

func (c *Cloud) getMultiCloudImageRef(clusterUUID string) (string, error) {
	clusterResp, err := c.APIServer.GetContrailCluster(c.ctx, &services.GetContrailClusterRequest{
		ID: clusterUUID,
	})
	if err != nil {
		return "", errors.Wrap(err, "cannot get contrail cluster")
	}
	return ansible.ImageReference(
		clusterResp.ContrailCluster.ContainerRegistry,
		MultiCloudDockerImage,
		ansible.GetContrailVersion(clusterResp.ContrailCluster, c.log),
	)
}
